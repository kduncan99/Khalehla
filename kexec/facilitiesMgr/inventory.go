// Khalehla Project
// Copyright © 2023-2024 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package facilitiesMgr

import (
	"khalehla/hardware"
	"khalehla/kexec"
	"khalehla/kexec/nodeMgr"
)

type inventory struct {
	nodes map[hardware.NodeIdentifier]kexec.INodeAttributes
	disks map[hardware.NodeIdentifier]*kexec.DiskAttributes
	tapes map[hardware.NodeIdentifier]*kexec.TapeAttributes
}

func newInventory() *inventory {
	i := &inventory{
		nodes: make(map[hardware.NodeIdentifier]kexec.INodeAttributes),
		disks: make(map[hardware.NodeIdentifier]*kexec.DiskAttributes),
		tapes: make(map[hardware.NodeIdentifier]*kexec.TapeAttributes),
	}
	return i
}

func (i *inventory) injectNode(nodeInfo nodeMgr.NodeInfo) {
	if nodeInfo.GetNodeCategoryType() == hardware.NodeCategoryDevice {
		devInfo := nodeInfo.(nodeMgr.DeviceInfo)
		devId := devInfo.GetNodeIdentifier()
		switch devInfo.GetNodeDeviceType() {
		case hardware.NodeDeviceDisk:
			attr := &kexec.DiskAttributes{
				Identifier: devInfo.GetNodeIdentifier(),
				Name:       devInfo.GetNodeName(),
				Status:     kexec.FacNodeStatusUp,
			}
			i.nodes[devId] = attr
			i.disks[devId] = attr
		case hardware.NodeDeviceTape:
			attr := &kexec.TapeAttributes{
				Identifier: devInfo.GetNodeIdentifier(),
				Name:       devInfo.GetNodeName(),
				Status:     kexec.FacNodeStatusUp,
			}
			i.nodes[devId] = attr
			i.tapes[devId] = attr
		}
	}
}
