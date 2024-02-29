// Khalehla Project
// Copyright © 2023-2024 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package nodeMgr

import (
	"io"
	"khalehla/kexec"
)

// DeviceInfo is intended primarily as a means of documenting the use of a more generic NodeInfo
type DeviceInfo interface {
	CreateNode()
	Dump(destination io.Writer, indent string)
	GetChannelInfos() []ChannelInfo
	GetDevice() Device
	GetNodeCategoryType() kexec.NodeCategoryType
	GetNodeDeviceType() kexec.NodeDeviceType
	GetNodeIdentifier() kexec.NodeIdentifier
	GetNodeName() string
	IsAccessible() bool
	IsReady() bool
	SetIsAccessible(bool)
	SetIsReady(flag bool)
}
