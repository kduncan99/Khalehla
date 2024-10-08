// Khalehla Project
// Copyright © 2023-2024 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package kexec

import "khalehla/hardware"

type DiskAttributes struct {
	Identifier    hardware.NodeIdentifier
	Name          string
	Status        FacNodeStatus
	PackLabelInfo *PackLabelInfo
	IsPrepped     bool
	IsFixed       bool
	IsRemovable   bool
}

func (da *DiskAttributes) GetFacNodeStatus() FacNodeStatus {
	return da.Status
}

func (da *DiskAttributes) GetNodeCategoryType() hardware.NodeCategoryType {
	return hardware.NodeCategoryDevice
}

func (da *DiskAttributes) GetNodeDeviceType() hardware.NodeDeviceType {
	return hardware.NodeDeviceDisk
}

func (da *DiskAttributes) GetNodeIdentifier() hardware.NodeIdentifier {
	return da.Identifier
}

func (da *DiskAttributes) GetNodeName() string {
	return da.Name
}

func (da *DiskAttributes) SetFacNodeStatus(status FacNodeStatus) {
	da.Status = status
}
