// Khalehla Project
// Copyright © 2023-2024 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package csi

import (
	"khalehla/kexec/facilitiesMgr"
	"khalehla/klog"
)

func handleSymcn(pkt *handlerPacket) (*facilitiesMgr.FacStatusResult, uint64) {
	klog.LogTraceF("CSI", "handleSymcn:%v", *pkt.pcs)

	/*
		@SYMCN,L
		@SYMCN,N [{n}]
	*/
	// TODO implement
	return nil, 0
}
