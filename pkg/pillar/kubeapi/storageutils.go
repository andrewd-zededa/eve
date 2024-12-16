// Copyright (c) 2024 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

//go:build kubevirt

package kubeapi

import (
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/lf-edge/eve/pkg/pillar/pubsub"
	"github.com/lf-edge/eve/pkg/pillar/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	longhornDevPath = "/dev/longhorn"
)

func deviceMajor(stat syscall.Stat_t) int32 {
	return int32((stat.Rdev >> 8) & 0xfff)
}
func deviceMinor(stat syscall.Stat_t) int32 {
	return int32((stat.Rdev & 0xff) | ((stat.Rdev >> 12) & 0xfff00))
}
func getMajorMinor(stat syscall.Stat_t) string {
	major := deviceMajor(stat)
	minor := deviceMinor(stat)
	return fmt.Sprintf("%d:%d", major, minor)
}

func isPvMountedOnThisNode(pv string) bool {
	if _, err := os.Stat(longhornDevPath + "/" + pv); err != nil {
		return false
	}
	return true
}
func isBlockDeviceMountedOnThisNode(dev string) bool {
	if _, err := os.Stat("/dev/" + dev); err != nil {
		return false
	}
	return true
}

// CleanupUnmountedDiskMetrics loops over existing DiskMetric
// objects and unpublishes them if the device no longer exists
func CleanupUnmountedDiskMetrics(pubDiskMetric pubsub.Publication, pvcToPvMap map[string]string) {
	existingMetrics := pubDiskMetric.GetAll()

	for id, metric := range existingMetrics {
		if id == "" {
			continue
		}

		dm, ok := metric.(types.DiskMetric)
		if ok && dm.IsDir {
			continue
		}

		if strings.Contains(id, "pvc-") {
			pvName, ok := pvcToPvMap[id]
			// Could be PVC deleted or just not mounted locally
			if !ok || !isPvMountedOnThisNode(pvName) {
				pubDiskMetric.Unpublish(id)
			}
		} else {
			// Look for sdX devices which used to exist
			// These would have been the block device
			// which shared major:minor with the longhorn device
			if !isBlockDeviceMountedOnThisNode(id) {
				pubDiskMetric.Unpublish(id)
			}
		}
	}
}

// LonghornGetMajorMinorMaps builds two maps between
// device major:minor -> kube-pv-name/lh-volume-name
// and kube-pv-name/lh-volume-name -> maj:min to
// help callers find a PV/PVC in /proc/diskstats
// which only shows the sdX path.
func LonghornGetMajorMinorMaps() (map[string]string, map[string]string, error) {
	lhMajMinToNameMap := make(map[string]string) // maj:min -> kube-pv-name/lh-volume-name
	lhNameToMajMinMap := make(map[string]string) // kube-pv-name/lh-volume-name -> maj:min

	if _, err := os.Stat(longhornDevPath); err != nil {
		return lhMajMinToNameMap, lhNameToMajMinMap, fmt.Errorf("longhorn dev path missing")
	}

	lhPvcList, err := os.ReadDir(longhornDevPath)
	if err != nil {
		return lhMajMinToNameMap, lhNameToMajMinMap, fmt.Errorf("unable to read longhorn devs")
	}

	for _, lhDirEnt := range lhPvcList {
		var lhStat syscall.Stat_t
		err := syscall.Stat(longhornDevPath+"/"+lhDirEnt.Name(), &lhStat)

		if err != nil {
			continue
		}
		majMinKey := getMajorMinor(lhStat)
		lhMajMinToNameMap[majMinKey] = lhDirEnt.Name()
		lhNameToMajMinMap[lhDirEnt.Name()] = majMinKey

	}
	return lhMajMinToNameMap, lhNameToMajMinMap, nil
}

// SCSIGetMajMinMaps builds two maps to assist linking with other devices
// First map: maj:min -> sdX
// Second map: sdX -> maj:min
func SCSIGetMajMinMaps() (map[string]string, map[string]string, error) {
	sdMajMinToNameMap := make(map[string]string) // maj:min -> sdX
	sdNameToMajMinMap := make(map[string]string) // sdX -> maj:min

	blockDevs, err := os.ReadDir("/sys/class/block/")
	if err != nil {
		return sdMajMinToNameMap, sdNameToMajMinMap, fmt.Errorf("unable to read sd devs")
	}

	for _, devEnt := range blockDevs {
		if !strings.HasPrefix(devEnt.Name(), "sd") {
			continue
		}

		var blockStat syscall.Stat_t
		err := syscall.Stat("/dev/"+devEnt.Name(), &blockStat)
		if err != nil {
			continue
		}
		majMinVal := getMajorMinor(blockStat)
		sdMajMinToNameMap[majMinVal] = devEnt.Name()
		sdNameToMajMinMap[devEnt.Name()] = majMinVal
	}
	return sdMajMinToNameMap, sdNameToMajMinMap, nil
}

// PvPvcMaps returns two maps of pv-name/longhorn-name -> pvc-name
// and pvc-name -> pv-name/longhorn-name
func PvPvcMaps() (map[string]string, map[string]string, error) {
	pvsMap := make(map[string]string)
	pvcsMap := make(map[string]string)

	clientset, err := GetClientSet()
	if err != nil {
		return pvsMap, pvcsMap, fmt.Errorf("PvToPvc_Map: can't get clientset %v", err)
	}

	pvcs, err := clientset.CoreV1().PersistentVolumeClaims(EVEKubeNameSpace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return pvsMap, pvcsMap, fmt.Errorf("get pvcs:%v", err)
	}
	for _, pvc := range pvcs.Items {
		pvsMap[pvc.Spec.VolumeName] = pvc.ObjectMeta.Name
		pvcsMap[pvc.ObjectMeta.Name] = pvc.Spec.VolumeName
	}
	return pvsMap, pvcsMap, nil
}