// Khalehla Project
// Copyright © 2023-2024 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package kexec

import "khalehla/hardware"

type INodeAttributes interface {
	GetFacNodeStatus() FacNodeStatus
	GetNodeCategoryType() hardware.NodeCategoryType
	GetNodeDeviceType() hardware.NodeDeviceType
	GetNodeIdentifier() hardware.NodeIdentifier
	GetNodeName() string
	SetFacNodeStatus(status FacNodeStatus)
}
