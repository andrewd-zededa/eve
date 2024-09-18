// Copyright (c) 2024 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

//go:build kubevirt

package zedkube

import (
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/lf-edge/eve/pkg/pillar/agentbase"
	"github.com/lf-edge/eve/pkg/pillar/agentlog"
	"github.com/lf-edge/eve/pkg/pillar/base"
	"github.com/lf-edge/eve/pkg/pillar/cipher"
	"github.com/lf-edge/eve/pkg/pillar/kubeapi"
	"github.com/lf-edge/eve/pkg/pillar/pubsub"
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
)

const (
	agentName = "zedkube"
	// Time limits for event loop handlers
	errorTime            = 3 * time.Minute
	warningTime          = 40 * time.Second
	stillRunningInterval = 25 * time.Second
	logcollectInterval   = 30
	// run VNC file
	vmiVNCFileName    = "/run/zedkube/vmiVNC.run"
	serverTLSDir      = "/persist/kube-save-var-lib/rancher/k3s/server/tls"
	pubServerCertFile = serverTLSDir + "/client-k3s-controller.crt"
	pubServerKeyFile  = serverTLSDir + "/client-k3s-controller.key"

	inlineCmdKubeClusterUpdateStatus = "pubKubeClusterUpdateStatus"
)

var (
	logger *logrus.Logger
	log    *base.LogObject
)

type ReceiveMap struct {
	mu sync.Mutex
	v  map[string]bool
}

type zedkubeContext struct {
	agentbase.AgentBase
	globalConfig             *types.ConfigItemValueMap
	subAppInstanceConfig     pubsub.Subscription
	subGlobalConfig          pubsub.Subscription
	subDeviceNetworkStatus   pubsub.Subscription
	subEdgeNodeClusterConfig pubsub.Subscription
	subNetworkInstanceConfig pubsub.Subscription
	subVolumeConfig          pubsub.Subscription
	subDatastoreConfig       pubsub.Subscription
	subContentTreeConfig     pubsub.Subscription
	subEdgeNodeInfo          pubsub.Subscription
	subZedAgentStatus        pubsub.Subscription

	subControllerCert    pubsub.Subscription
	subEdgeNodeCert      pubsub.Subscription
	cipherMetrics        *cipher.AgentMetrics
	pubCipherBlockStatus pubsub.Publication
	pubCipherMetrics     pubsub.Publication

	pubEncPubToRemoteData    pubsub.Publication
	pubEdgeNodeClusterStatus pubsub.Publication
	pubENClusterAppStatus    pubsub.Publication
	pubKubeClusterInfo       pubsub.Publication

	subNodeDrainRequestZA  pubsub.Subscription
	subNodeDrainRequestBoM pubsub.Subscription
	pubNodeDrainStatus     pubsub.Publication

	networkInstanceStatusMap sync.Map
	ioAdapterMap             sync.Map
	config                   *rest.Config
	appLogStarted            bool
	appContainerLogger       *logrus.Logger
	encNodeIPAddress         *net.IP
	nodeuuid                 string
	nodeName                 string
	isKubeStatsLeader        bool
	inKubeLeaderElection     bool
	electionStartCh          chan struct{}
	electionStopCh           chan struct{}
	pubResendTimer           *time.Timer
	drainOverrideTimer       *time.Timer
	receiveMap               *ReceiveMap
	stopMonitor              chan struct{}
	clusterPubSubStarted     bool
	quitServer               chan struct{}
	statusServer             *http.Server
	drainTimeoutHours        uint32
	pubServerCertFile        string
	pubServerKeyFile         string
	notifyPeerCount          int
}

func inlineUsage() int {
	log.Errorf("Usage: zedkube %s <node> <component> <status> <DestinationKubeUpdateVersion> <error>", inlineCmdKubeClusterUpdateStatus)
	return 1
}

func runCommand(ps *pubsub.PubSub, command string, args []string) int {
	if args == nil {
		return inlineUsage()
	}
	switch command {
	case inlineCmdKubeClusterUpdateStatus:
		if args == nil {
			return inlineUsage()
		}
		node := args[0]
		comp := kubeapi.KubeCompFromStr(args[1])
		status := kubeapi.KubeCompUpdateStatusFromStr(args[2])

		dest_kube_version := uint32(0)
		val, err := strconv.ParseInt(args[3], 10, 32)
		if err != nil {
			log.Errorf("zedkube %s unable to parse dest_version:%s err:%v", inlineCmdKubeClusterUpdateStatus, args[3], err)
			return 1
		}
		dest_kube_version = uint32(val)

		error_str := ""
		if len(args) == 5 {
			error_str = args[4]
		}

		pubKubeClusterUpdateStatus, err := ps.NewPublication(
			pubsub.PublicationOptions{
				AgentName:  "zedagent",
				TopicType:  kubeapi.KubeClusterUpdateStatus{},
				Persistent: true,
			})
		if err != nil {
			log.Fatal(err)
			return 2
		}
		if (comp == kubeapi.COMP_UNKNOWN) && (status == kubeapi.COMP_STATUS_UNKNOWN) {
			if _, err := pubKubeClusterUpdateStatus.Get("global"); err == nil {
				pubKubeClusterUpdateStatus.Unpublish("global")
			}
		} else {
			upStatusObj := kubeapi.KubeClusterUpdateStatus{
				CurrentNode:                  node,
				Component:                    comp,
				Status:                       status,
				DestinationKubeUpdateVersion: dest_kube_version,
			}
			if status == kubeapi.COMP_STATUS_FAILED {
				upStatusObj.SetError(error_str, time.Now())
			}
			pubKubeClusterUpdateStatus.Publish("global", upStatusObj)
		}
	default:
		log.Errorf("Unknown command %s", command)
		return 99
	}

	ps.StillRunning("zedkube", warningTime, errorTime)
	time.Sleep(time.Second * 1)
	return 0
}

// Run - an zedkube run
func Run(ps *pubsub.PubSub, loggerArg *logrus.Logger, logArg *base.LogObject, arguments []string, baseDir string) int {
	logger = loggerArg
	log = logArg

	zedkubeCtx := zedkubeContext{
		globalConfig: types.DefaultConfigItemValueMap(),
	}

	// do we run a single command, or long-running service?
	// if any args defined, will run that single command and exit.
	// otherwise, will run the agent
	var (
		command string
		args    []string
	)
	if len(arguments) > 0 {
		command = arguments[0]
	}
	if len(arguments) > 1 {
		args = arguments[1:]
	}

	// if an explicit command was given, run that command and return, else run the agent
	if command != "" {
		return runCommand(ps, command, args)
	}

	agentbase.Init(&zedkubeCtx, logger, log, agentName,
		agentbase.WithPidFile(),
		agentbase.WithWatchdog(ps, warningTime, errorTime),
		agentbase.WithArguments(arguments))

	// Run a periodic timer so we always update StillRunning
	stillRunning := time.NewTicker(stillRunningInterval)

	zedkubeCtx.stopMonitor = make(chan struct{})
	zedkubeCtx.quitServer = make(chan struct{})
	zedkubeCtx.appContainerLogger = agentlog.CustomLogInit(logrus.InfoLevel)

	// Get AppInstanceConfig from zedagent
	subAppInstanceConfig, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		AgentName:     "zedagent",
		MyAgentName:   agentName,
		TopicImpl:     types.AppInstanceConfig{},
		Activate:      false,
		Ctx:           &zedkubeCtx,
		CreateHandler: handleAppInstanceConfigCreate,
		ModifyHandler: handleAppInstanceConfigModify,
		DeleteHandler: handleAppInstanceConfigDelete,
		WarningTime:   warningTime,
		ErrorTime:     errorTime,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subAppInstanceConfig = subAppInstanceConfig
	subAppInstanceConfig.Activate()

	// Look for controller certs which will be used for decryption.
	subControllerCert, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		AgentName:   "zedagent",
		MyAgentName: agentName,
		TopicImpl:   types.ControllerCert{},
		Persistent:  true,
		Activate:    true,
		WarningTime: warningTime,
		ErrorTime:   errorTime,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subControllerCert = subControllerCert

	// Look for edge node certs which will be used for decryption
	subEdgeNodeCert, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		AgentName:   "tpmmgr",
		MyAgentName: agentName,
		TopicImpl:   types.EdgeNodeCert{},
		Persistent:  true,
		Activate:    true,
		WarningTime: warningTime,
		ErrorTime:   errorTime,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subEdgeNodeCert = subEdgeNodeCert

	pubCipherBlockStatus, err := ps.NewPublication(
		pubsub.PublicationOptions{
			AgentName: agentName,
			TopicType: types.CipherBlockStatus{},
		})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.pubCipherBlockStatus = pubCipherBlockStatus
	pubCipherBlockStatus.ClearRestarted()

	pubEdgeNodeClusterStatus, err := ps.NewPublication(
		pubsub.PublicationOptions{
			AgentName: agentName,
			TopicType: types.EdgeNodeClusterStatus{},
		})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.pubEdgeNodeClusterStatus = pubEdgeNodeClusterStatus

	pubENClusterAppStatus, err := ps.NewPublication(
		pubsub.PublicationOptions{
			AgentName: agentName,
			TopicType: types.ENClusterAppStatus{},
		})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.pubENClusterAppStatus = pubENClusterAppStatus

	pubKubeClusterInfo, err := ps.NewPublication(
		pubsub.PublicationOptions{
			AgentName: agentName,
			TopicType: types.KubeClusterInfo{},
		})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.pubKubeClusterInfo = pubKubeClusterInfo

	// Look for global config such as log levels
	subGlobalConfig, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		AgentName:     "zedagent",
		MyAgentName:   agentName,
		TopicImpl:     types.ConfigItemValueMap{},
		Persistent:    true,
		Activate:      false,
		Ctx:           &zedkubeCtx,
		CreateHandler: handleGlobalConfigCreate,
		ModifyHandler: handleGlobalConfigModify,
		WarningTime:   warningTime,
		ErrorTime:     errorTime,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subGlobalConfig = subGlobalConfig
	subGlobalConfig.Activate()

	// Watch DNS to learn which ports are used for management.
	subDeviceNetworkStatus, err := ps.NewSubscription(
		pubsub.SubscriptionOptions{
			AgentName:   "nim",
			MyAgentName: agentName,
			TopicImpl:   types.DeviceNetworkStatus{},
			Activate:    false,
			Ctx:         &zedkubeCtx,
			WarningTime: warningTime,
			ErrorTime:   errorTime,
		})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subDeviceNetworkStatus = subDeviceNetworkStatus
	subDeviceNetworkStatus.Activate()

	// setup a map to keep track of received encPubToRemoteData
	// so we don't send something's publication out as ours
	zedkubeCtx.receiveMap = newReceiveMap()

	// For cluster publication
	subNetworkInstanceConfig, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		AgentName:     "zedagent",
		MyAgentName:   agentName,
		TopicImpl:     types.NetworkInstanceConfig{},
		Activate:      false,
		Ctx:           &zedkubeCtx,
		CreateHandler: handleNetworkInstanceCreate,
		ModifyHandler: handleNetworkInstanceModify,
		DeleteHandler: handleNetworkInstanceDelete,
		WarningTime:   warningTime,
		ErrorTime:     errorTime,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subNetworkInstanceConfig = subNetworkInstanceConfig
	subNetworkInstanceConfig.Activate()

	subVolumeConfig, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		CreateHandler: handleVolumeCreate,
		ModifyHandler: handleVolumeModify,
		DeleteHandler: handleVolumeDelete,
		//RestartHandler: handleVolumeRestart, // XXX
		WarningTime: warningTime,
		ErrorTime:   errorTime,
		AgentName:   "zedagent",
		MyAgentName: agentName,
		TopicImpl:   types.VolumeConfig{},
		Ctx:         &zedkubeCtx,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subVolumeConfig = subVolumeConfig
	subVolumeConfig.Activate()

	pubEncPubToRemoteData, err := ps.NewPublication(
		pubsub.PublicationOptions{
			AgentName: agentName,
			TopicType: types.EncPubToRemoteData{},
		})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.pubEncPubToRemoteData = pubEncPubToRemoteData

	zedkubeCtx.cipherMetrics = cipher.NewAgentMetrics(agentName)
	pubCipherMetrics, err := ps.NewPublication(
		pubsub.PublicationOptions{
			AgentName: agentName,
			TopicType: types.CipherMetrics{},
		})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.pubCipherMetrics = pubCipherMetrics

	subDatastoreConfig, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		CreateHandler: handleDatastoreConfigCreate,
		ModifyHandler: handleDatastoreConfigModify,
		DeleteHandler: handleDatastoreConfigDelete,
		WarningTime:   warningTime,
		ErrorTime:     errorTime,
		AgentName:     "zedagent",
		TopicImpl:     types.DatastoreConfig{},
		Ctx:           &zedkubeCtx,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subDatastoreConfig = subDatastoreConfig
	subDatastoreConfig.Activate()

	subContentTreeConfig, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		CreateHandler: handleContentTreeCreate,
		ModifyHandler: handleContentTreeModify,
		DeleteHandler: handleContentTreeDelete,
		//RestartHandler: handleContentTreeRestart,
		WarningTime: warningTime,
		ErrorTime:   errorTime,
		AgentName:   "zedagent",
		MyAgentName: agentName,
		TopicImpl:   types.ContentTreeConfig{},
		Ctx:         &zedkubeCtx,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subContentTreeConfig = subContentTreeConfig
	subContentTreeConfig.Activate()

	// start the leader election
	zedkubeCtx.electionStartCh = make(chan struct{})
	zedkubeCtx.electionStopCh = make(chan struct{})
	go handleLeaderElection(&zedkubeCtx)

	// Wait for device network status to be initialized, this is need for
	// provisionting the 2nd ip address on the cluster interface, otherwise
	// we'll add ip prefix onto kethX interface instead of ethX interface
	// and wait for the certs, which cluster config need to decrypt the token
	var deviceNetStatusInitialized, controllerCertInitiazlized, edgenodeCertInitiazlized bool
	for !deviceNetStatusInitialized || !controllerCertInitiazlized || !edgenodeCertInitiazlized {
		log.Noticef("zedkube run: waiting for device network status, net %v, controller %v, edgenode %v",
			deviceNetStatusInitialized, controllerCertInitiazlized, edgenodeCertInitiazlized)
		select {
		case change := <-subDeviceNetworkStatus.MsgChan():
			subDeviceNetworkStatus.ProcessChange(change)
			deviceNetStatusInitialized = true

		case change := <-subControllerCert.MsgChan():
			subControllerCert.ProcessChange(change)
			controllerCertInitiazlized = true

		case change := <-subEdgeNodeCert.MsgChan():
			subEdgeNodeCert.ProcessChange(change)
			edgenodeCertInitiazlized = true

		case <-stillRunning.C:
		}
		ps.StillRunning(agentName, warningTime, errorTime)
	}
	log.Noticef("zedkube run: device network status initialized")

	//
	// NodeDrainRequest subscriber and NodeDrainStatus publisher
	//
	// Sub the request
	zedkubeCtx.subNodeDrainRequestZA, err = ps.NewSubscription(pubsub.SubscriptionOptions{
		CreateHandler: handleNodeDrainRequestCreate,
		ModifyHandler: handleNodeDrainRequestModify,
		DeleteHandler: handleNodeDrainRequestDelete,
		WarningTime:   warningTime,
		ErrorTime:     errorTime,
		AgentName:     "zedagent",
		MyAgentName:   agentName,
		TopicImpl:     kubeapi.NodeDrainRequest{},
		Ctx:           &zedkubeCtx,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subNodeDrainRequestZA.Activate()

	// Sub the request
	zedkubeCtx.subNodeDrainRequestBoM, err = ps.NewSubscription(pubsub.SubscriptionOptions{
		CreateHandler: handleNodeDrainRequestCreate,
		ModifyHandler: handleNodeDrainRequestModify,
		DeleteHandler: handleNodeDrainRequestDelete,
		WarningTime:   warningTime,
		ErrorTime:     errorTime,
		AgentName:     "baseosmgr",
		MyAgentName:   agentName,
		TopicImpl:     kubeapi.NodeDrainRequest{},
		Ctx:           &zedkubeCtx,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subNodeDrainRequestBoM.Activate()

	//Pub the status
	zedkubeCtx.pubNodeDrainStatus, err = ps.NewPublication(
		pubsub.PublicationOptions{
			AgentName: agentName,
			TopicType: kubeapi.NodeDrainStatus{},
		})
	if err != nil {
		log.Fatal(err)
	}

	zedkubeCtx.drainOverrideTimer = time.NewTimer(1 * time.Minute)
	zedkubeCtx.drainOverrideTimer.Stop()
	// Until we hear otherwise that we are in a cluster
	publishNodeDrainStatus(&zedkubeCtx, kubeapi.NOTSUPPORTED)

	// EdgeNodeClusterConfig create needs to publish NodeDrainStatus, so wait to activate it.
	time.Sleep(5 * time.Second)

	// EdgeNodeClusterConfig subscription
	subEdgeNodeClusterConfig, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		AgentName:     "zedagent",
		MyAgentName:   agentName,
		TopicImpl:     types.EdgeNodeClusterConfig{},
		Persistent:    true,
		Activate:      false,
		Ctx:           &zedkubeCtx,
		CreateHandler: handleEdgeNodeClusterConfigCreate,
		ModifyHandler: handleEdgeNodeClusterConfigModify,
		DeleteHandler: handleEdgeNodeClusterConfigDelete,
		WarningTime:   warningTime,
		ErrorTime:     errorTime,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subEdgeNodeClusterConfig = subEdgeNodeClusterConfig
	subEdgeNodeClusterConfig.Activate()

	// Look for edge node info
	subEdgeNodeInfo, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		AgentName:   "zedagent",
		MyAgentName: agentName,
		TopicImpl:   types.EdgeNodeInfo{},
		Persistent:  true,
		Activate:    true,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subEdgeNodeInfo = subEdgeNodeInfo

	// subscribe to zedagent status events, for controller connection status
	subZedAgentStatus, err := ps.NewSubscription(pubsub.SubscriptionOptions{
		AgentName:     "zedagent",
		MyAgentName:   agentName,
		TopicImpl:     types.ZedAgentStatus{},
		Activate:      false,
		Ctx:           &zedkubeCtx,
		CreateHandler: handleZedAgentStatusCreate,
		ModifyHandler: handleZedAgentStatusModify,
		DeleteHandler: handleZedAgentStatusDelete,
		WarningTime:   warningTime,
		ErrorTime:     errorTime,
	})
	if err != nil {
		log.Fatal(err)
	}
	zedkubeCtx.subZedAgentStatus = subZedAgentStatus
	subZedAgentStatus.Activate()

	if len(subEdgeNodeClusterConfig.GetAll()) != 0 {
		// Handle persistent existing cluster config
		publishNodeDrainStatus(&zedkubeCtx, kubeapi.NOTREQUESTED)
	}

	err = kubeapi.WaitForKubernetes(agentName, ps, subEdgeNodeClusterConfig, stillRunning)
	if err != nil {
		log.Errorf("zedkube: WaitForKubenetes %v", err)
	}
	zedkubeCtx.config, err = kubeapi.GetKubeConfig()
	if err != nil {
		log.Errorf("zedkube: GetKubeConfig %v", err)
	} else {
		log.Noticef("zedkube: running")
	}

	zedkubeCtx.pubResendTimer = time.NewTimer(60 * time.Second)
	zedkubeCtx.pubResendTimer.Stop()

	//Re-enable local node
	log.Noticef("zedkube re-enable-node/uncordon+")
	cordoned, err := isNodeCordoned(&zedkubeCtx)
	if err != nil {
		log.Errorf("zedkube can't read local node cordon state, err:%v", err)
	} else {
		log.Noticef("zedkube isNodeCordoned cordoned:%v", cordoned)
		if cordoned {
			if err := cordonNode(&zedkubeCtx, false); err != nil {
				log.Errorf("zedkube Unable to uncordon local node: %v", err)
			}
		}
	}
	log.Noticef("zedkube re-enable-node/uncordon-")

	// notify peer nodes we are up, if there is any pubs, resend them
	if zedkubeCtx.clusterPubSubStarted {
		startupNotifyPeers(&zedkubeCtx)
	}

	appLogTimer := time.NewTimer(logcollectInterval * time.Second)

	for {
		select {
		case change := <-subAppInstanceConfig.MsgChan():
			subAppInstanceConfig.ProcessChange(change)

		case <-appLogTimer.C:
			collectAppLogs(&zedkubeCtx)
			checkAppsStatus(&zedkubeCtx)
			collectKubeStats(&zedkubeCtx)
			checkPubServerStatus(&zedkubeCtx)
			checkNotifyPeer(&zedkubeCtx)
			appLogTimer = time.NewTimer(logcollectInterval * time.Second)

		case change := <-subGlobalConfig.MsgChan():
			subGlobalConfig.ProcessChange(change)

		case change := <-subEdgeNodeClusterConfig.MsgChan():
			subEdgeNodeClusterConfig.ProcessChange(change)

		case change := <-subNetworkInstanceConfig.MsgChan():
			subNetworkInstanceConfig.ProcessChange(change)

		case change := <-subVolumeConfig.MsgChan():
			subVolumeConfig.ProcessChange(change)

		case change := <-subDatastoreConfig.MsgChan():
			subDatastoreConfig.ProcessChange(change)

		case change := <-subContentTreeConfig.MsgChan():
			subContentTreeConfig.ProcessChange(change)

		case <-zedkubeCtx.pubResendTimer.C:
			// Resend the cluster pub info
			resendPubsToRemoteNodes(&zedkubeCtx)

		case change := <-subControllerCert.MsgChan():
			subControllerCert.ProcessChange(change)

		case change := <-subEdgeNodeCert.MsgChan():
			subEdgeNodeCert.ProcessChange(change)

		case change := <-subEdgeNodeInfo.MsgChan():
			subEdgeNodeInfo.ProcessChange(change)

		case change := <-subZedAgentStatus.MsgChan():
			subZedAgentStatus.ProcessChange(change)

		case change := <-zedkubeCtx.subNodeDrainRequestZA.MsgChan():
			zedkubeCtx.subNodeDrainRequestZA.ProcessChange(change)

		case change := <-zedkubeCtx.subNodeDrainRequestBoM.MsgChan():
			zedkubeCtx.subNodeDrainRequestBoM.ProcessChange(change)

		case <-zedkubeCtx.drainOverrideTimer.C:
			override := kubeapi.GetDrainStatusOverride()
			if override != nil {
				zedkubeCtx.pubNodeDrainStatus.Publish("global", override)
			}
			zedkubeCtx.drainOverrideTimer.Reset(5 * time.Minute)

		case <-stillRunning.C:
		}
		ps.StillRunning(agentName, warningTime, errorTime)
	}
}

func handleAppInstanceConfigCreate(ctxArg interface{}, key string,
	configArg interface{}) {
	ctx := ctxArg.(*zedkubeContext)
	config := configArg.(types.AppInstanceConfig)

	log.Noticef("handleAppInstanceConfigCreate(%v) spec for %s",
		config.UUIDandVersion, config.DisplayName) // XXX

	err := checkIoAdapterEthernet(ctx, &config)
	log.Functionf("handleAppInstancConfigModify: genAISpec %v", err)

	sendAndPubEncAppInstConfig(ctx, &config, key, types.EncPubOpCreate)
}

func handleAppInstanceConfigModify(ctxArg interface{}, key string,
	configArg interface{}, oldConfigArg interface{}) {
	ctx := ctxArg.(*zedkubeContext)
	config := configArg.(types.AppInstanceConfig)
	oldconfig := oldConfigArg.(types.AppInstanceConfig)

	log.Noticef("handleAppInstancConfigModify(%v) spec for %s",
		config.UUIDandVersion, config.DisplayName) // XXX

	err := checkIoAdapterEthernet(ctx, &config)

	if oldconfig.RemoteConsole != config.RemoteConsole {
		log.Functionf("handleAppInstancConfigModify: new remote console %v", config.RemoteConsole)
		go runAppVNC(ctx, &config)
	}
	log.Functionf("handleAppInstancConfigModify: genAISpec %v", err)

	sendAndPubEncAppInstConfig(ctx, &config, key, types.EncPubOpModify)
}

func handleAppInstanceConfigDelete(ctxArg interface{}, key string,
	configArg interface{}) {

	log.Functionf("handleAppInstanceConfigDelete(%s)", key)
	ctx := ctxArg.(*zedkubeContext)
	config := configArg.(types.AppInstanceConfig)

	checkDelIoAdapterEthernet(ctx, &config)
	log.Functionf("handleAppInstanceConfigDelete(%s) done", key)

	sendAndPubEncAppInstConfig(ctx, nil, key, types.EncPubOpDelete)

	// remove the cluster app status publication
	pub := ctx.pubENClusterAppStatus
	stItmes := pub.GetAll()
	for _, st := range stItmes {
		aiStatus := st.(types.ENClusterAppStatus)
		if aiStatus.AppUUID == config.UUIDandVersion.UUID {
			ctx.pubENClusterAppStatus.Unpublish(config.UUIDandVersion.UUID.String())
			break
		}
	}
}

func handleGlobalConfigCreate(ctxArg interface{}, key string,
	statusArg interface{}) {
	handleGlobalConfigImpl(ctxArg, key, statusArg)
}

func handleGlobalConfigModify(ctxArg interface{}, key string,
	statusArg interface{}, oldStatusArg interface{}) {
	handleGlobalConfigImpl(ctxArg, key, statusArg)
}

func handleGlobalConfigImpl(ctxArg interface{}, key string,
	statusArg interface{}) {

	ctx := ctxArg.(*zedkubeContext)
	if key != "global" {
		log.Functionf("handleGlobalConfigImpl: ignoring %s", key)
		return
	}
	log.Functionf("handleGlobalConfigImpl for %s", key)
	gcp := agentlog.HandleGlobalConfig(log, ctx.subGlobalConfig, agentName,
		ctx.CLIParams().DebugOverride, ctx.Logger())
	if gcp != nil {
		allowClusterPubSub := gcp.GlobalValueBool(types.ENClusterPubSub)
		if allowClusterPubSub && !ctx.clusterPubSubStarted {
			log.Noticef("handleGlobalConfigImpl: starting cluster pubsub")

			// Start the cluster pubsub server
			go runClusterPubSubServer(ctx)
		}

		currentConfigItemValueMap := ctx.globalConfig
		newConfigItemValueMap := gcp
		// Handle Drain Timeout Change
		if newConfigItemValueMap.GlobalValueInt(types.KubevirtDrainTimeout) != 0 &&
			newConfigItemValueMap.GlobalValueInt(types.KubevirtDrainTimeout) !=
				currentConfigItemValueMap.GlobalValueInt(types.KubevirtDrainTimeout) {
			log.Functionf("handleGlobalConfigImpl: Updating drainTimeoutHours from %d to %d",
				currentConfigItemValueMap.GlobalValueInt(types.KubevirtDrainTimeout),
				newConfigItemValueMap.GlobalValueInt(types.KubevirtDrainTimeout))
			ctx.drainTimeoutHours = newConfigItemValueMap.GlobalValueInt(types.KubevirtDrainTimeout)
		}
	}
	log.Functionf("handleGlobalConfigImpl(%s): done", key)
}

func handleEdgeNodeClusterConfigCreate(ctxArg interface{}, key string,
	configArg interface{}) {
	log.Noticef("handleEdgeNodeClusterConfigCreate: %s", key)
	handleEdgeNodeClusterConfigImpl(ctxArg, key, configArg, nil)
}

func handleEdgeNodeClusterConfigModify(ctxArg interface{}, key string,
	configArg interface{}, oldConfigArg interface{}) {
	log.Noticef("handleEdgeNodeClusterConfigModify: %s", key)
	handleEdgeNodeClusterConfigImpl(ctxArg, key, configArg, oldConfigArg)
}

func handleEdgeNodeClusterConfigImpl(ctxArg interface{}, key string,
	configArg interface{}, oldConfigArg interface{}) {

	var config, oldconfig types.EdgeNodeClusterConfig
	var oldConfigPtr *types.EdgeNodeClusterConfig
	config = configArg.(types.EdgeNodeClusterConfig)
	if oldConfigArg != nil {
		oldconfig = oldConfigArg.(types.EdgeNodeClusterConfig)
		oldConfigPtr = &oldconfig
	}

	ctx := ctxArg.(*zedkubeContext)
	log.Noticef("handleEdgeNodeClusterConfigImpl for %s, config %+v, oldconfig %+v",
		key, config, oldconfig)

	runKubeConfig(ctx, &config, oldConfigPtr, false)

	publishNodeDrainStatus(ctx, kubeapi.NOTREQUESTED)
}

func handleEdgeNodeClusterConfigDelete(ctxArg interface{}, key string,
	statusArg interface{}) {
	ctx := ctxArg.(*zedkubeContext)
	log.Noticef("handleEdgeNodeClusterConfigDelete for %s", key)
	config := statusArg.(types.EdgeNodeClusterConfig)
	runKubeConfig(ctx, &config, nil, true)
	ctx.pubEdgeNodeClusterStatus.Unpublish("global")

	publishNodeDrainStatus(ctx, kubeapi.NOTSUPPORTED)
}

// handle zedagent status events, for cloud connectivity
func handleZedAgentStatusCreate(ctxArg interface{}, key string,
	statusArg interface{}) {
	handleZedAgentStatusImpl(ctxArg, key, statusArg)
}

func handleZedAgentStatusModify(ctxArg interface{}, key string,
	statusArg interface{}, oldStatusArg interface{}) {
	handleZedAgentStatusImpl(ctxArg, key, statusArg)
}

func handleZedAgentStatusImpl(ctxArg interface{}, key string,
	statusArg interface{}) {

	ctxPtr := ctxArg.(*zedkubeContext)
	status := statusArg.(types.ZedAgentStatus)
	handleControllerStatusChange(ctxPtr, &status)
	log.Functionf("handleZedAgentStatusImpl: for Leader status %v, done", status)
}

func handleZedAgentStatusDelete(ctxArg interface{}, key string,
	statusArg interface{}) {
	// do nothing
	log.Functionf("handleZedAgentStatusDelete(%s) done", key)
}

func newReceiveMap() *ReceiveMap {
	return &ReceiveMap{v: make(map[string]bool)}
}

func (s *ReceiveMap) Insert(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.v[key] = true
}

func (s *ReceiveMap) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.v, key)
}

func (s *ReceiveMap) Find(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.v[key]
	return ok
}
