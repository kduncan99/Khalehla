// Khalehla Project
// Copyright © 2023-2024 by Kurt Duncan, BearSnake LLC
// All Rights Reserved

package mfdMgr

import (
	"fmt"
	"khalehla/kexec/types"
	"log"
)

type packFreeSpaceRegion struct {
	trackId    types.TrackId
	trackCount types.TrackCount
}

func newPackFreeSpaceRegion(trackId types.TrackId, trackCount types.TrackCount) *packFreeSpaceRegion {
	return &packFreeSpaceRegion{
		trackId:    trackId,
		trackCount: trackCount,
	}
}

type packFreeSpaceTable struct {
	capacity types.TrackCount
	content  []*packFreeSpaceRegion
}

func newPackFreeSpaceTable(capacity types.TrackCount) *packFreeSpaceTable {
	fst := &packFreeSpaceTable{}
	fst.capacity = capacity
	fsr := newPackFreeSpaceRegion(0, capacity)
	fst.content = []*packFreeSpaceRegion{fsr}
	return fst
}

// markTrackRegionUnallocated is a general-purpose function which manipulates the entries in a free space table
func (fst *packFreeSpaceTable) markTrackRegionAllocated(
	ldatIndex types.LDATIndex, // only for logging
	trackId types.TrackId,
	trackCount types.TrackCount) bool {

	if trackCount == 0 {
		log.Printf("markTrackRegionAllocated ldat:%v id:%v count:%v requested trackCount is zero",
			ldatIndex, trackId, trackCount)
	}

	// We're looking for a region of free space which contains the requested region
	reqTrackLimit := trackId + types.TrackId(trackCount) // track limit from specified id and count
	for fx, fsRegion := range fst.content {
		// Is requested region less than the current entry? If so, there's no point in continuing
		if trackId < fsRegion.trackId {
			break
		}

		// Does requested region begin within the current entry?
		entryLimit := types.TrackId(uint64(fsRegion.trackId) + uint64(fsRegion.trackCount))
		if trackId >= fsRegion.trackId && trackId <= entryLimit {
			// Quick check to ensure requested region does not exceed this entry.
			// If it does, something is bigly wrong
			if reqTrackLimit > entryLimit {
				log.Printf("markTrackRegionAllocated ldat:%v id:%v count:%v region too big",
					ldatIndex, trackId, trackCount)
				return false
			}

			// Does the requested region exactly match the current entry?
			// If so, just remove the current entry
			if fsRegion.trackCount == trackCount {
				fst.content = append(fst.content[:fx], fst.content[fx:]...)
				return true
			}

			// Is the region to be removed aligned with the front of the current entry?
			if trackId == fsRegion.trackId {
				fsRegion.trackId += types.TrackId(trackCount)
				fsRegion.trackCount -= trackCount
				return true
			}

			// Is the region to be removed aligned with the back of the current entry?
			if reqTrackLimit == entryLimit {
				fsRegion.trackCount -= trackCount
				return true
			}

			// Break the region into two sections. Messy. Don't like it.
			newRegion := newPackFreeSpaceRegion(entryLimit, types.TrackCount(entryLimit-reqTrackLimit))
			fsRegion.trackCount = types.TrackCount(trackId - fsRegion.trackId)
			newTable := append(fst.content[0:fx+1], newRegion)
			fst.content = append(newTable, fst.content[fx+1])
			return true
		}
	}

	log.Printf("markTrackRegionAllocated ldat:%v id:%v count:%v track not allocated",
		ldatIndex, trackId, trackCount)
	return false
}

// markTrackRegionUnallocated is a general-purpose function which manipulates the entries in a free space table
func (fst *packFreeSpaceTable) markTrackRegionUnallocated(
	ldatIndex types.LDATIndex, // only for logging
	trackId types.TrackId,
	trackCount types.TrackCount) bool {

	if trackCount == 0 {
		log.Printf("markTrackRegionUnallocated ldat:%v id:%v count:%v requested trackCount is zero",
			ldatIndex, trackId, trackCount)
	}

	// We are hoping that we do not find an entry which contains all or part of the requested region
	reqTrackLimit := trackId + types.TrackId(trackCount) // track limit from specified id and count
	for fx, fsRegion := range fst.content {
		// Does requested region overlap with this entry?
		entryTrackLimit := types.TrackId(uint64(fsRegion.trackId) + uint64(fsRegion.trackCount))
		if trackId >= fsRegion.trackId && trackId < entryTrackLimit {
			log.Printf("markTrackRegionUnallocated ldat%v id:%v count:%v region overlap",
				ldatIndex, trackId, trackCount)
			return false
		} else if reqTrackLimit > fsRegion.trackId && reqTrackLimit <= entryTrackLimit {
			log.Printf("markTrackRegionUnallocated ldat%v id:%v count:%v region overlap",
				ldatIndex, trackId, trackCount)
			return false
		}

		// Is requested region between this entry and the next?
		// If so, we need to coalesce this and the next region
		if fx < len(fst.content)-1 {
			fsNext := fst.content[fx+1]
			if trackId == entryTrackLimit && reqTrackLimit == fst.content[fx+1].trackId {
				fsRegion.trackCount += trackCount + fsNext.trackCount
				fst.content = append(fst.content[:fx+1], fst.content[fx+1:]...)
				return true
			}
		}

		// Is requested region aligned with the front of this entry?
		if reqTrackLimit == fsRegion.trackId {
			fsRegion.trackId = trackId
			fsRegion.trackCount += trackCount
			return true
		}

		// Is requested region aligned with the back of this entry?
		if trackId == entryTrackLimit {
			fsRegion.trackCount += trackCount
			return true
		}

		// Region is not aligned with the front or back of this entry, nor does it overlap.
		// If it is ahead of this entry, then we just need to insert a new entry for the requested region.
		if trackId < fsRegion.trackId {
			re := newPackFreeSpaceRegion(trackId, trackCount)
			newTable := append(fst.content[:fx], re)
			fst.content = append(newTable, fst.content[fx:]...)
			return true
		}
	}

	// Region is somewhere at the end of the pack. Create a new entry.
	re := newPackFreeSpaceRegion(trackId, trackCount)
	fst.content = append(fst.content, re)
	return true
}

// allocateSpecificTrackRegion is used only when it has been determined by some external means, that a particular
// track or range of tracks is not to be allocated otherwise (such as for VOL1 or directory tracks).
func (fst *packFreeSpaceTable) allocateSpecificTrackRegion(
	ldatIndex types.LDATIndex,
	trackId types.TrackId,
	trackCount types.TrackCount) error {

	ok := fst.markTrackRegionAllocated(ldatIndex, trackId, trackCount)
	if !ok {
		return fmt.Errorf("track not allocated")
	}
	return nil
}