// Khalehla Project
// Copyright © 2023-2024 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package kexec

import "khalehla/hardware"

// Because this as FacNodeStatus, it needs to be somewhere in kexec

type NodeAttributes interface {
	GetFacNodeStatus() FacNodeStatus
	GetNodeCategoryType() hardware.NodeCategoryType
	GetNodeDeviceType() hardware.NodeDeviceType
	GetNodeIdentifier() hardware.NodeIdentifier
	GetNodeName() string
	SetFacNodeStatus(status FacNodeStatus)
}
