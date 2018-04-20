package channels

import (
    ds3Models "spectra/ds3_go_sdk/ds3/models"
    "spectra/ds3_go_sdk/helpers/ranges"
    "reflect"
    "fmt"
    "errors"
)

// Manages the mapping of object ranges to end destination offsets
type partialObjectRangeMap struct {
    // map of object ranges to destination offset for start of range
    rangeMap map[ds3Models.Range]int64
}

func newPartialObjectRangeMap(objRanges []ds3Models.Range) partialObjectRangeMap {
    collapser := ranges.RangeCollapserImpl{}
    collapsedRanges := collapser.Collapse(objRanges)

    rangeMap := make(map[ds3Models.Range]int64, len(collapsedRanges))

    var prevStartOffset int64 = 0
    var prevLength int64 = 0
    for _, r := range collapsedRanges {
        curStartOffset := prevStartOffset + prevLength
        rangeMap[r] = curStartOffset
        prevStartOffset = curStartOffset
        prevLength = r.End - r.Start + 1
    }
    return partialObjectRangeMap{rangeMap:rangeMap}
}

// Retrieves the destination offset that correlates to the object offset based on the ranges.
func (rangeMap *partialObjectRangeMap) getDestinationOffset(objectOffset int64) (int64, error) {
    for curRange, destinationOffset := range rangeMap.rangeMap {
        if objectOffset >= curRange.Start && objectOffset <= curRange.End {
            return destinationOffset + (objectOffset - curRange.Start), nil
        }
    }

    ranges := reflect.ValueOf(rangeMap.rangeMap).MapKeys()
    return -1, errors.New(fmt.Sprintf("Partial Object Channel Builder: specified object offset '%d' is outside of expected ranges '%v'", objectOffset, ranges))
}