// Copyright (c) 2023 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (
	// EveVirtTypeFile contains the virtualization type, ie kvm, xen or kubevirt
	EveVirtTypeFile = "/run/eve-hv-type"
	// KubeAppNameMaxLen limits the length of the app name for Kubernetes.
	// This also includes the appended UUID prefix.
	KubeAppNameMaxLen = 32
	// KubeAppNameUUIDSuffixLen : number of characters taken from the app UUID and appended
	// to the app name for Kubernetes (to avoid name collisions between apps of the same
	// DisplayName, see GetAppKubeName).
	KubeAppNameUUIDSuffixLen = 5
	// VMIPodNamePrefix : prefix added to name of every pod created to run VM.
	VMIPodNamePrefix = "virt-launcher-"
	// InstallOptionEtcdSizeGB grub option at install time.  Size of etcd volume in GB.
	InstallOptionEtcdSizeGB = "eve_install_kubevirt_etcd_sizeGB"
	// DefaultEtcdSizeGB default for InstallOptionEtcdSizeGB
	DefaultEtcdSizeGB uint32 = 10
	// EtcdVolBlockSizeBytes is the block size for the etcd volume
	EtcdVolBlockSizeBytes = uint64(4 * 1024)
)

// IsHVTypeKube - return true if the EVE image is kube cluster type.
func IsHVTypeKube() bool {
	retbytes, err := os.ReadFile(EveVirtTypeFile)
	if err != nil {
		return false
	}

	if strings.Contains(string(retbytes), "kubevirt") {
		return true
	}
	return false
}

// XXX get local cluster ip address
func GetLocalClusterIP(isHVKube bool) (string, error) {
	if !isHVKube {
		return "", nil
	}
	file, err := os.Open("/config/server")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstLine := scanner.Text()

	if strings.Contains(firstLine, "zedcloud.local.zededa.net") {
		hostsFile, err := os.Open("/config/hosts")
		if err != nil {
			return "", nil // return empty string if hosts file is not there
		}
		defer hostsFile.Close()

		hostsScanner := bufio.NewScanner(hostsFile)
		for hostsScanner.Scan() {
			line := hostsScanner.Text()
			if strings.Contains(line, "zedcloud.local.zededa.net") {
				ipAddress := strings.Fields(line)[0] // get the first field (IP address)
				return ipAddress, nil
			}
		}
	}

	return "", nil // return empty string if no matching line is found
}

var (
	kubeNameForbiddenChars = regexp.MustCompile("[^a-zA-Z0-9-.]")
	kubeNameSeparators     = regexp.MustCompile("[.-]+")
)

// GetAppKubeName returns name of the application used inside Kubernetes (for Pod or VMI).
func GetAppKubeName(displayName string, uuid uuid.UUID) string {
	appKubeName := displayName
	// Replace underscores with dashes for Kubernetes
	appKubeName = strings.ReplaceAll(appKubeName, "_", "-")
	// Remove special characters using regular expressions
	appKubeName = kubeNameForbiddenChars.ReplaceAllString(appKubeName, "")
	// Reduce combinations like '-.-' or '.-.' to a single dash
	appKubeName = kubeNameSeparators.ReplaceAllString(appKubeName, "-")
	appKubeName = strings.ToLower(appKubeName)
	const maxLen = KubeAppNameMaxLen - 1 - KubeAppNameUUIDSuffixLen
	if len(appKubeName) > maxLen {
		appKubeName = appKubeName[:maxLen]
	}
	return appKubeName + "-" + uuid.String()[:KubeAppNameUUIDSuffixLen]
}

// GetVMINameFromVirtLauncher : get VMI name from the corresponding Kubevirt
// launcher pod name.
func GetVMINameFromVirtLauncher(podName string) (vmiName string, isVirtLauncher bool) {
	if !strings.HasPrefix(podName, VMIPodNamePrefix) {
		return "", false
	}
	vmiName = strings.TrimPrefix(podName, VMIPodNamePrefix)
	lastSep := strings.LastIndex(vmiName, "-")
	if lastSep != -1 {
		vmiName = vmiName[:lastSep]
	}
	return vmiName, true
}

func GetReplicaPodName(displayName, podName string, uuid uuid.UUID) (kubeName string, isReplicaPod bool) {
	kubeName = GetAppKubeName(displayName, uuid)
	if !strings.HasPrefix(podName, kubeName) {
		return "", false
	}
	suffix := strings.TrimPrefix(podName, kubeName)
	if strings.HasPrefix(suffix, "-") && len(suffix[1:]) == 5 {
		return kubeName, true
	}
	return "", false
}
