// Khalehla Project
// Copyright © 2023-2024 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package facilitiesMgr

import (
	"fmt"
	"io"
	"khalehla/kexec/types"
	"sync"
	"time"
)

type FacilitiesManager struct {
	exec            types.IExec
	mutex           sync.Mutex
	isInitialized   bool
	terminateThread bool
	threadStarted   bool
	threadStopped   bool

	deviceAssignments map[*types.DeviceInfo]*types.RunControlEntry
}

func NewFacilitiesManager(exec types.IExec) *FacilitiesManager {
	return &FacilitiesManager{
		exec:              exec,
		deviceAssignments: make(map[*types.DeviceInfo]*types.RunControlEntry),
	}
}

// CloseManager is invoked when the exec is stopping
func (mgr *FacilitiesManager) CloseManager() {
	// TODO
	mgr.threadStop()
	mgr.isInitialized = false
}

func (mgr *FacilitiesManager) InitializeManager() error {
	// TODO
	mgr.threadStart()
	mgr.isInitialized = true
	return nil
}

func (mgr *FacilitiesManager) IsInitialized() bool {
	return mgr.isInitialized
}

// ResetManager clears out any artifacts left over by a previous exec session,
// and prepares the console for normal operations
func (mgr *FacilitiesManager) ResetManager() error {
	// TODO

	mgr.threadStop()
	mgr.threadStart()
	mgr.isInitialized = true
	return nil
}

func (mgr *FacilitiesManager) NotifyDeviceReady(deviceInfo types.DeviceInfo, isReady bool) {
	// TODO
}

func (mgr *FacilitiesManager) AssignDevice(deviceName string, rce *types.RunControlEntry) {

}

func (mgr *FacilitiesManager) GetDeviceStatusDetail(deviceInfo types.DeviceInfo) string {
	str := ""
	if mgr.isInitialized {
		// TODO
	}
	return str
}

func (mgr *FacilitiesManager) thread() {
	mgr.threadStarted = true

	for !mgr.terminateThread {
		time.Sleep(25 * time.Millisecond)
		// TODO
	}

	mgr.threadStopped = true
}

func (mgr *FacilitiesManager) threadStart() {
	mgr.terminateThread = false
	if !mgr.threadStarted {
		go mgr.thread()
		for !mgr.threadStarted {
			time.Sleep(25 * time.Millisecond)
		}
	}
}

func (mgr *FacilitiesManager) threadStop() {
	if mgr.threadStarted {
		mgr.terminateThread = true
		for !mgr.threadStopped {
			time.Sleep(25 * time.Millisecond)
		}
	}
}

func (mgr *FacilitiesManager) Dump(dest io.Writer, indent string) {
	_, _ = fmt.Fprintf(dest, "%vFacilitiesManager ----------------------------------------------------\n", indent)

	_, _ = fmt.Fprintf(dest, "%v  initialized:     %v\n", indent, mgr.isInitialized)
	_, _ = fmt.Fprintf(dest, "%v  threadStarted:   %v\n", indent, mgr.threadStarted)
	_, _ = fmt.Fprintf(dest, "%v  threadStopped:   %v\n", indent, mgr.threadStopped)
	_, _ = fmt.Fprintf(dest, "%v  terminateThread: %v\n", indent, mgr.terminateThread)

	// TODO

}
