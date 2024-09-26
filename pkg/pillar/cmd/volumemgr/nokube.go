// Copyright (c) 2024 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

//go:build !kubevirt

package volumemgr

// createOrUpdatePvcDiskMetrics has no work in non kubevirt builds
func createOrUpdatePvcDiskMetrics(*volumemgrContext, string) {
	return
}
