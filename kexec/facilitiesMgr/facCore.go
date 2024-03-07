// Khalehla Project
// Copyright © 2023-2024 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package facilitiesMgr

import (
	"khalehla/kexec"
	"khalehla/kexec/config"
	"khalehla/kexec/mfdMgr"
	"khalehla/kexec/nodeMgr"
	"log"
	"strconv"
	"strings"
)

type fieldSubfieldIndex struct {
	fieldIndex    int
	subFieldIndex int
	allSubfields  bool
}

type fieldSubfieldIndices struct {
	content []fieldSubfieldIndex
}

func newFieldSubfieldIndices() *fieldSubfieldIndices {
	return &fieldSubfieldIndices{
		content: make([]fieldSubfieldIndex, 0),
	}
}

func (fsi *fieldSubfieldIndices) add(fieldIndex int, subfieldIndex int) *fieldSubfieldIndices {
	index := fieldSubfieldIndex{
		fieldIndex:    fieldIndex,
		subFieldIndex: subfieldIndex,
	}
	fsi.content = append(fsi.content, index)
	return fsi
}

func (fsi *fieldSubfieldIndices) addAll(fieldIndex int) *fieldSubfieldIndices {
	index := fieldSubfieldIndex{
		fieldIndex:   fieldIndex,
		allSubfields: true,
	}
	fsi.content = append(fsi.content, index)
	return fsi
}

func (fsi *fieldSubfieldIndices) contains(fieldIndex int, subfieldIndex int) bool {
	for _, fsx := range fsi.content {
		if fieldIndex == fsx.fieldIndex && subfieldIndex == fsx.subFieldIndex {
			return true
		}
	}
	return false
}

// -----------------------------------------------------------------------------

var catFixedFSIs = newFieldSubfieldIndices().
	add(0, 0).
	add(1, 0).
	add(1, 1).
	add(1, 2).
	add(1, 3)

var catRemovableFSIs = newFieldSubfieldIndices().
	add(0, 0).
	add(1, 0).
	add(1, 1).
	add(1, 2).
	add(1, 3).
	addAll(2)

// -----------------------------------------------------------------------------

// checkSubFields
// Checks the user-provided operation fields against a list of accepted field/subfield combinations
// to see whether the user provided a subfield which is not acceptable.
// Returns true if all is well, else false
func (mgr *FacilitiesManager) checkSubFields(operandFields [][]string, accepted *fieldSubfieldIndices) bool {
	for fx := 0; fx < len(operandFields); fx++ {
		for fy := 0; fy < len(operandFields[fx]); fy++ {
			if len(operandFields[fx][fy]) > 0 && !accepted.contains(fx, fy) {
				return false
			}
		}
	}
	return true
}

// getField
// Retrieves the field indicated by the given field index as an array of strings.
// If the field was not specified, we return an empty array.
func (mgr *FacilitiesManager) getField(operandFields [][]string, fieldIndex int) []string {
	if fieldIndex < len(operandFields) {
		return operandFields[fieldIndex]
	} else {
		return []string{}
	}
}

// getSubField
// Retrieves the subfield indicated by the given field and subfield indicies.
// If the subfield was not specified, we return a blank string.
func (mgr *FacilitiesManager) getSubField(operandFields [][]string, fieldIndex int, subfieldIndex int) string {
	if fieldIndex < len(operandFields) && subfieldIndex < len(operandFields[fieldIndex]) {
		return operandFields[fieldIndex][subfieldIndex]
	} else {
		return ""
	}
}

// -----------------------------------------------------------------------------

func (mgr *FacilitiesManager) assignFixedFile(
	exec kexec.IExec,
	rce *kexec.RunControlEntry,
	fileSpecification *FileSpecification,
	optionWord uint64,
	operandFields [][]string,
	fileSetInfo *mfdMgr.FileSetInfo,
	mnemonic string,
	usage config.EquipmentUsage,
	sourceIsExecRequest bool,
) (facResult *FacStatusResult, resultCode uint64) {
	return nil, 0 // TODO
}

func (mgr *FacilitiesManager) assignRemovableFile(
	exec kexec.IExec,
	rce *kexec.RunControlEntry,
	fileSpecification *FileSpecification,
	optionWord uint64,
	operandFields [][]string,
	fileSetInfo *mfdMgr.FileSetInfo,
	mnemonic string,
	usage config.EquipmentUsage,
	sourceIsExecRequest bool,
) (facResult *FacStatusResult, resultCode uint64) {
	return nil, 0 // TODO
}

func (mgr *FacilitiesManager) assignTapeFile(
	exec kexec.IExec,
	rce *kexec.RunControlEntry,
	fileSpecification *FileSpecification,
	optionWord uint64,
	operandFields [][]string,
	fileSetInfo *mfdMgr.FileSetInfo,
	mnemonic string,
	usage config.EquipmentUsage,
	sourceIsExecRequest bool,
) (facResult *FacStatusResult, resultCode uint64) {
	return nil, 0 // TODO
}

// catalogFixedFile takes the various inputs, validates them, and then invokes
// mfd services to create the appropriate file cycle (and file set, if necessary).
// Caller should immediately check whether the exec has stopped upon return.
func (mgr *FacilitiesManager) catalogFixedFile(
	exec kexec.IExec,
	rce kexec.RunControlEntry,
	fileSpecification *FileSpecification,
	optionWord uint64,
	operandFields [][]string,
	fileSetInfo *mfdMgr.FileSetInfo,
	mnemonic string,
	usage config.EquipmentUsage,
	sourceIsExecRequest bool,
) (facResult *FacStatusResult, resultCode uint64) {
	//	For Mass Storage Files
	//		@CAT[,options] filename[,type/reserve/granule/maximum,pack-id-1/.../pack-id-n,,,ACR-name]
	//	options include
	//		B: save on checkpoint
	//		G: guarded file
	//		P: make the file public (not private)
	//		R: make the file read-only
	//		V: file will not be unloaded
	//		W: make the file write-only
	//		Z: run should not be held (probably only happens on removable when the pack is not mounted)
	//			I'm unaware of any situation where cataloging a fixed file would result in a hold.
	facResult = NewFacResult()
	resultCode = 0

	allowedOpts := uint64(kexec.BOption | kexec.GOption | kexec.POption |
		kexec.ROption | kexec.VOption | kexec.WOption | kexec.ZOption)
	if !checkIllegalOptions(rce, optionWord, allowedOpts, facResult, rce.IsExec()) {
		resultCode |= 0_600000_000000
		return
	}

	if !mgr.checkSubFields(operandFields, catFixedFSIs) {
		if len(mgr.getSubField(operandFields, 1, 4)) > 0 {
			facResult.PostMessage(FacStatusPlacementFieldNotAllowed, nil)
		}
		facResult.PostMessage(FacStatusUndefinedFieldOrSubfield, nil)
		resultCode |= 0_600000_000000
		return
	}

	saveOnCheckpoint := optionWord&kexec.BOption != 0
	guardedFile := optionWord&kexec.GOption != 0
	publicFile := optionWord&kexec.POption != 0
	readOnly := optionWord&kexec.ROption != 0
	inhibitUnload := optionWord&kexec.VOption != 0
	writeOnly := optionWord&kexec.WOption != 0
	wordAddressable := usage == config.EquipmentUsageWordAddressableMassStorage

	// ensure initial reserve <= max allocations (means words or granules, depending on word/sector addressable)
	initStr := mgr.getSubField(operandFields, 1, 1)
	granStr := strings.ToUpper(mgr.getSubField(operandFields, 1, 2))
	maxStr := mgr.getSubField(operandFields, 1, 3)

	var granularity kexec.Granularity
	if len(granStr) == 0 || granStr == "TRK" {
		granularity = kexec.TrackGranularity
	} else if granStr == "POS" {
		granularity = kexec.PositionGranularity
	} else {
		facResult.PostMessage(FacStatusIllegalValueForGranularity, nil)
		resultCode |= 0_600000_000000
		return
	}

	var initReserve uint64
	if len(initStr) > 12 {
		facResult.PostMessage(FacStatusIllegalInitialReserve, nil)
		resultCode |= 0_600000_000000
		return
	} else if len(initStr) > 0 {
		initReserve, err := strconv.Atoi(initStr)
		if err != nil || initReserve < 0 {
			facResult.PostMessage(FacStatusIllegalInitialReserve, nil)
			resultCode |= 0_600000_000000
			return
		}
	}

	maxGranules := exec.GetConfiguration().MaxGranules
	if len(maxStr) > 12 {
		facResult.PostMessage(FacStatusIllegalMaxGranules, nil)
		resultCode |= 0_600000_000000
		return
	} else if len(maxStr) > 0 {
		iMaxGran, err := strconv.Atoi(maxStr)
		maxGranules = uint64(iMaxGran)
		if err != nil || maxGranules < 0 || maxGranules > 262143 {
			facResult.PostMessage(FacStatusIllegalMaxGranules, nil)
			resultCode |= 0_600000_000000
			return
		} else if maxGranules < initReserve {
			facResult.PostMessage(FacStatusMaximumIsLessThanInitialReserve, nil)
			resultCode |= 0_600000_000000
			return
		}
	}

	// If there isn't an existing fileset, create one.
	mm := exec.GetMFDManager().(*mfdMgr.MFDManager)
	if fileSetInfo == nil {
		_, result := mm.CreateFileSet(
			mfdMgr.FileTypeFixed,
			fileSpecification.Qualifier,
			fileSpecification.Filename,
			rce.GetProjectId(),
			fileSpecification.ReadKey,
			fileSpecification.WriteKey)
		if result == mfdMgr.MFDInternalError {
			return
		} else if result != mfdMgr.MFDSuccessful {
			log.Printf("FacMgr:MFD failed to create file set")
			exec.Stop(kexec.StopFacilitiesComplex)
			return
		}
	} else {
		// Otherwise, do sanity checking on the file cycle and privacy settings on the existing file set.
		if fileSetInfo.PlusOneExists {
			facResult.PostMessage(FacStatusRelativeFCycleConflict, nil)
			resultCode |= 0_600000_000040
			return
		}

		gaveReadKey := len(fileSpecification.ReadKey) > 0
		gaveWriteKey := len(fileSpecification.WriteKey) > 0
		hasReadKey := len(fileSetInfo.ReadKey) > 0
		hasWriteKey := len(fileSetInfo.WriteKey) > 0
		needsMsg := false
		if hasReadKey {
			if !gaveReadKey {
				facResult.PostMessage(FacStatusReadWriteKeysNeeded, nil)
				needsMsg = true
				resultCode |= 0_600000_000000
				return
			} else if fileSetInfo.ReadKey != fileSpecification.ReadKey {
				facResult.PostMessage(FacStatusIncorrectReadKey, nil)
				resultCode |= 0_401000_000000
				if sourceIsExecRequest {
					rce.PostContingencyWithAuxiliary(017, 0, 0, 015)
				}
				return
			}
		} else {
			if gaveReadKey {
				facResult.PostMessage(FacStatusFileNotCatalogedWithReadKey, nil)
				resultCode |= 0_400040_000000
				if sourceIsExecRequest {
					rce.PostContingencyWithAuxiliary(017, 0, 0, 015)
				}
				return
			}
		}

		if hasWriteKey {
			if !gaveWriteKey && !needsMsg {
				facResult.PostMessage(FacStatusReadWriteKeysNeeded, nil)
				resultCode |= 0_600000_000000
				return
			} else if fileSetInfo.WriteKey != fileSpecification.WriteKey {
				facResult.PostMessage(FacStatusIncorrectWriteKey, nil)
				resultCode |= 0_400400_000000
				if sourceIsExecRequest {
					rce.PostContingencyWithAuxiliary(017, 0, 0, 015)
				}
				return
			}
		} else {
			if gaveWriteKey {
				facResult.PostMessage(FacStatusFileNotCatalogedWithWriteKey, nil)
				resultCode |= 0_400020_000000
				if sourceIsExecRequest {
					rce.PostContingencyWithAuxiliary(017, 0, 0, 015)
				}
				return
			}
		}

		/*
			E:242533 File cycle out of range.
			E:242633 Cannot catalog file because read or write access not allowed.
			E:243233 Creation of file would require illegal dropping of private file.
			E:244433 File is already catalogued.
			E:253733 Relative F-cycle conflict.
		*/

		dropOldest := false
		if fileSpecification.AbsoluteCycle != nil {
			// Caller has given us a specific absolute cycle.
			// We need to ensure that it is not so far below the highest existing cycle,
			// and that it is not so far above the second-lowest existing cycle,
			// that it would violate the max cycle range. It *can* be sufficiently above the lowest cycle that
			// the lowest cycle needs to be dropped... and if that is the case, we need to ensure
			// that we can actually drop that cycle.
			// Also we need to ensure the actual absolute cycle doesn't already exist.
			highest := fileSetInfo.HighestAbsolute
			lowest := highest - fileSetInfo.CurrentRange + 1
			var highestAllowed uint
			var lowestAllowed uint
			if highest == lowest {
				// only one cycle exists
				lowestAllowed =
			}
			if lowest <= highest {
				// things are nicely normal
				lowestAllowed := highest - fileSetInfo.MaxCycleRange + 1
				highestAllowed := lowest + fileSetInfo.MaxCycleRange
			} else {
				// wraparound exists
			}
			// TODO
		} else if fileSpecification.RelativeCycle != nil {
			// If we get here, we've already limited relative cycle to only +1.
			// Is there already a +1?
			if fileSetInfo.PlusOneExists {
				facResult.PostMessage(FacStatusRelativeFCycleConflict, nil)
				resultCode |= 0_600000_000040
				return
			}

			// If the highest absolute cycle file is deleted, then that file is (re)cataloged.
			// Otherwise, a new highest absolute cycle is created and all the other files are shifted
			// downward by one cycle, which might cause the lowest cycle to be deleted
			// (depending upon the current and max cycle ranges).
			// At this point, we only care about the latter case, in that we are checking whether
			// it is necessary and possible to delete an oldest cycle.
			if fileSetInfo.CycleInfo[0] != nil && fileSetInfo.CycleInfo[fileSetInfo.CurrentRange-1] != nil {
				dropOldest = true
			}
		} else {
			// We're here with a file set but no cycle spec on a @CAT request. That won't fly.
			facResult.PostMessage(FacStatusFileAlreadyCataloged, nil)
			resultCode |= 0_500000_000000
			return
		}

		if dropOldest {
			oldestFCIdent := fileSetInfo.CycleInfo[fileSetInfo.CurrentRange-1].FileCycleIdentifier
			fcInfo, mfdResult := mm.GetFileCycleInfo(oldestFCIdent)
			if mfdResult != mfdMgr.MFDSuccessful {
				log.Printf("FacMgr:MFD can't find a filecycle which should exist")
				mgr.exec.Stop(kexec.StopDirectoryErrors)
				return
			}

			if fcInfo.AssignedIndicator {
				facResult.PostMessage(FacStatusRelativeFCycleConflict, nil)
				resultCode |= 0_600000_000040
				return
			}

			// TODO check the following
			//  should we deal with Guarded and private earlier? Can we?
			//fcInfo.InhibitFlags.IsGuarded
			//fcInfo.InhibitFlags.IsPrivate
			//fcInfo.InhibitFlags.IsReadOnly
			//fcInfo.InhibitFlags.IsWriteOnly

			// TODO invoke MFD services to do so
			mfdResult = mm.DropFileCycle(oldestFCIdent)
			if mfdResult != mfdMgr.MFDSuccessful {
				return
			}
		}
	}

	// Create the file cycle for the file set.
	var absCycle *uint // TODO
	var relCycle *int  // TODO
	descriptorFlags := mfdMgr.DescriptorFlags{
		SaveOnCheckpoint:    saveOnCheckpoint,
		IsTapeFile:          false,
		IsRemovableDiskFile: false,
	}
	fileFlags := mfdMgr.FileFlags{
		IsLargeFile: false, // TODO should we set this here, and how do we know?
	}
	pcharFlags := mfdMgr.PCHARFlags{
		Granularity:       granularity,
		IsWordAddressable: wordAddressable,
	}
	inhibitFlags := mfdMgr.InhibitFlags{
		IsGuarded:         guardedFile,
		IsUnloadInhibited: inhibitUnload,
		IsPrivate:         publicFile,
		IsWriteOnly:       writeOnly,
		IsReadOnly:        readOnly,
	}
	unitSelection := mfdMgr.UnitSelectionIndicators{}
	// TODO change the service below slightly... ?
	//   we don't need variance b/w abs and rel cycles, because in the end, we're not creating
	//   rel cycle (we're @CAT, so... yeah). BUT... maybe @ASG still needs that ability?
	_, mfdResult := mm.CreateFixedDiskFileCycle(
		fileSetInfo.FileSetIdentifier,
		absCycle,
		relCycle,
		rce.GetAccountId(),
		mnemonic,
		descriptorFlags,
		fileFlags,
		pcharFlags,
		inhibitFlags,
		initReserve,
		maxGranules,
		unitSelection,
		make([]mfdMgr.DiskPackEntry, 0))

	if mfdResult != mfdMgr.MFDSuccessful {
		// TODO what various things should we look for here?
	}

	return
}

func (mgr *FacilitiesManager) catalogRemovableFile(
	exec kexec.IExec,
	rce kexec.RunControlEntry,
	fileSpecification *FileSpecification,
	optionWord uint64,
	operandFields [][]string,
	fileSetInfo *mfdMgr.FileSetInfo,
	mnemonic string,
	usage config.EquipmentUsage,
	sourceIsExecRequest bool,
) (facResult *FacStatusResult, resultCode uint64) {
	//	For Mass Storage Files
	//		@CAT[,options] filename[,type/reserve/granule/maximum,pack-id-1/.../pack-id-n,,,ACR-name]
	//	options include
	//		B: save on checkpoint
	//		G: guarded file
	//		P: make the file public (not private)
	//		R: make the file read-only
	//		V: file will not be unloaded
	//		W: make the file write-only
	//		Z: run should not be held (probably only happens on removable when the pack is not mounted)
	allowedOpts := uint64(kexec.BOption | kexec.GOption | kexec.POption |
		kexec.ROption | kexec.VOption | kexec.WOption | kexec.ZOption)
	if !checkIllegalOptions(rce, optionWord, allowedOpts, facResult, rce.IsExec()) {
		// TODO
	}

	if !mgr.checkSubFields(operandFields, catRemovableFSIs) {
		// TODO
	}

	// saveOnCheckpoint := optionWord&kexec.BOption != 0
	// guardedFile := optionWord&kexec.GOption != 0
	// publicFile := optionWord&kexec.POption != 0
	// readOnly := optionWord&kexec.ROption != 0
	// inhibitUnload := optionWord&kexec.VOption != 0
	// writeOnly := optionWord&kexec.WOption != 0
	// doNotHold := optionWord&kexec.ZOption != 0
	// wordAddressable := usage == config.EquipmentUsageWordAddressableMassStorage

	// TODO granularity, initial-reserve, max-granules

	// Ensure the pack list is compatible with the files in the fileset (if there is a fileset)
	// Is it okay to just use the highest cycle?
	// TODO

	// If we are removable ensure each pack name is known and mounted.
	// Do not wait for mount if Z option is set
	// TODO

	return nil, 0 // TODO
}

func (mgr *FacilitiesManager) catalogTapeFile(
	exec kexec.IExec,
	rce kexec.RunControlEntry,
	fileSpecification *FileSpecification,
	optionWord uint64,
	operandFields [][]string,
	fileSetInfo *mfdMgr.FileSetInfo,
	mnemonic string,
	usage config.EquipmentUsage,
	sourceIsExecRequest bool,
) (facResult *FacStatusResult, resultCode uint64) {
	//	For Tape Files
	//		@CAT,options filename,type[/units/log/noise/processor/tape/
	//			format/data-converter/block-numbering/data-compression/
	//			buffered-write/expanded-buffer,reel-1/reel-2/.../reel-n,
	//			expiration-period/mmspec,,ACR-name,CTL-pool]
	//	options include
	//		E: even parity (not supported)
	//		G: guarded file
	//		H: density selection (not supported)
	//		J: tape is to be unlabeled
	//		L: density selection (not supported)
	//		M: density selection (not supported)
	//		O: odd parity (supported but ignored)
	//		P: make the file public
	//		R: make the file read-only
	//		S: 6250 BPI (only for SCSI 9-track - future)
	//		V: 1600 BPI (only for SCSI 9-track - future)
	//		W: make the file write-only
	//		Z: run should not be held (probably only happens on removable when the pack is not mounted)
	allowedOpts := uint64(kexec.EOption|kexec.GOption|kexec.HOption|kexec.JOption|
		kexec.LOption|kexec.MOption|kexec.OOption) | kexec.POption | kexec.ROption |
		kexec.SOption | kexec.VOption | kexec.WOption | kexec.ZOption
	if !checkIllegalOptions(rce, optionWord, allowedOpts, facResult, rce.IsExec()) {
		// TODO
	}

	return nil, 0 // TODO
}

// checkIllegalOptions compares the given options word to the allowed options word,
// producing a fac message for each option set in the given word which does not appear in the allowed word.
// Returns true if no such instances were found, else false
// If not ok and the source is an ER CSF$/ACSF$/CSI$, we post a contingency
func checkIllegalOptions(
	rce kexec.RunControlEntry,
	givenOptions uint64,
	allowedOptions uint64,
	facResult *FacStatusResult,
	sourceIsExec bool,
) bool {
	bit := uint64(kexec.AOption)
	letter := 'A'
	ok := true

	for {
		if bit&givenOptions != 0 && bit&allowedOptions == 0 {
			param := string(letter)
			facResult.PostMessage(FacStatusIllegalOption, []string{param})
			ok = false
		}

		if bit == kexec.ZOption {
			break
		} else {
			letter++
			bit >>= 1
		}
	}

	if !ok {
		if sourceIsExec {
			rce.PostContingency(012, 04, 040)
		}
	}

	return ok
}

func getEffectiveQualifier(rce kexec.RunControlEntry, fileSpec *FileSpecification) string {
	if fileSpec.HasAsterisk {
		if len(fileSpec.Qualifier) == 0 {
			return fileSpec.Qualifier
		} else {
			return rce.GetImpliedQualifier()
		}
	} else {
		return rce.GetDefaultQualifier()
	}
}

// selectEquipmentModel accepts an equipment mnemonic (likely from a control statement)
// and an optional FileSetInfo struct, and returns a list of NodeModel structs
// representing the various equipment models which can be used to satisfy the mnemonic.
// If the mnemonic is an @ASG or @CAT for a file cycle of an existing file set,
// the corresponding FileSetInfo struct must be specified.
// A false return indicates that the mnemonic is not found.
func (mgr *FacilitiesManager) selectEquipmentModel(
	mnemonic string,
	fileSetInfo *mfdMgr.FileSetInfo,
) ([]nodeMgr.NodeModel, config.EquipmentUsage, bool) {

	effectiveMnemonic := mnemonic

	// If we do not have a given mnemonic but we *do* have a fileSetInfo...
	if len(effectiveMnemonic) == 0 && fileSetInfo != nil {
		// Use the equipment type from the highest absolute fcycle entry of a not to-be file cycle
		//	(an existing file cycle which is not to-be-cataloged or to-be-deleted)... if there is one.
		// Otherwise, use the equipment type from the highest absolute fcycle entry of a to-be file cycle
		for _, preventToBe := range []bool{true, false} {
			for _, fsCycleInfo := range fileSetInfo.CycleInfo {
				if !preventToBe || (!fsCycleInfo.ToBeCataloged && !fsCycleInfo.ToBeDropped) {
					mm := mgr.exec.GetMFDManager().(*mfdMgr.MFDManager)
					fcInfo, mfdResult := mm.GetFileCycleInfo(fsCycleInfo.FileCycleIdentifier)
					if mfdResult != mfdMgr.MFDSuccessful {
						mgr.exec.Stop(kexec.StopFacilitiesComplex)
						return nil, 0, false
					}
					effectiveMnemonic = fcInfo.GetAssignMnemonic()
				}
			}
		}
	}

	// If we still do not have an effective mnemonic use the default sector-formatted mass storage mnemonic.
	if len(effectiveMnemonic) == 0 {
		effectiveMnemonic = mgr.exec.GetConfiguration().MassStorageDefaultMnemonic
	}

	// Now go look for the mnemonic in the configured equipment entry table.
	entry, ok := mgr.exec.GetConfiguration().EquipmentTable[mnemonic]
	if !ok {
		return nil, 0, false
	}

	models := make([]nodeMgr.NodeModel, 0)
	usage := entry.Usage
	for _, modelName := range entry.SelectableEquipment {
		model, ok := nodeMgr.NodeModelTable[modelName]
		if ok {
			models = append(models, model)
		}
	}
	return models, usage, true
}
