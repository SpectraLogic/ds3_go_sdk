package ranges

import (
    ds3Models "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    helperModels "github.com/SpectraLogic/ds3_go_sdk/helpers/models"
    "sort"
)

// Used to sort a slice of Range objects in ascending order.
// Original slice is modified on sort.
type Ranges []ds3Models.Range

func (ranges Ranges) Len() int {
    return len(ranges)
}

func (ranges Ranges) Swap(i, j int) {
    ranges[i], ranges[j] = ranges[j], ranges[i]
}

// compares Start and then End
func (ranges Ranges) Less(i, j int) bool {
    if ranges[i].Start == ranges[j].Start {
        return ranges[i].End < ranges[j].End
    }
    return ranges[i].Start < ranges[j].Start
}

type RangeCollapser interface {
    // Collapses a set of ranges. Any duplicates or sub-ranges are removed.
    // The original range slice must remain unmodified.
    Collapse(ranges []ds3Models.Range) []ds3Models.Range
}

type RangeCollapserImpl struct { }

// Creates a slice of ranges that contain the collapsed ranges provided.
// The original range slice remains unmodified.
func (RangeCollapserImpl) Collapse(ranges []ds3Models.Range) []ds3Models.Range {
    rangeCopy := make([]ds3Models.Range, len(ranges))
    copy(rangeCopy, ranges)

    // sorts ranges in ascending order
    sort.Sort(Ranges(rangeCopy))

    // collapse ranges
    for i, j := 0, 1; j < len(rangeCopy); {
        if rangeCopy[i] == rangeCopy[j] {
            // if they are the same delete the second item
            rangeCopy = append(rangeCopy[:j], rangeCopy[j+1:]...)
        } else if rangeCopy[i].End + 1 >= rangeCopy[j].Start {
            // if overlapping ranges, collapse into single entry
            if rangeCopy[i].End < rangeCopy[j].End {
                rangeCopy[i].End = rangeCopy[j].End //update i entry to encompass both
            }
            rangeCopy = append(rangeCopy[:j], rangeCopy[j+1:]...) //delete j entry
        } else {
            i++
            j++
        }
    }
    return rangeCopy
}

type BlobRangeFinder interface {
    // Retrieves the ranges associated with a given blob.
    GetRanges(name string, offset int64, length int64) []ds3Models.Range
}

type BlobRangeFinderImpl struct {
    // Map of objects to ranges to be retrieved from object.
    // Ranges are collapsed.
    rangeMap map[string][]ds3Models.Range
    collapser RangeCollapser
}

// Creates a BlobRangeFinderImpl which is populated with the provided getObjects
func NewBlobRangeFinder(getObjects *[]helperModels.GetObject) BlobRangeFinder {
    rangeMap := toRangeMap(getObjects)
    return &BlobRangeFinderImpl{
        rangeMap: *rangeMap,
        collapser: RangeCollapserImpl{},
    }
}

func toRangeMap(getObjects *[]helperModels.GetObject) *map[string][]ds3Models.Range {
    rangeMap := make(map[string][]ds3Models.Range)

    if getObjects == nil {
        return &rangeMap
    }

    collapser := RangeCollapserImpl{}
    for _, obj := range *getObjects {
        rangeMap[obj.Name] = collapser.Collapse(obj.Ranges)
    }
    return &rangeMap
}

func (rangeFinder *BlobRangeFinderImpl) GetRanges(name string, offset int64, length int64) []ds3Models.Range {
    objectRanges := rangeFinder.rangeMap[name]

    // get the subset of ranges that intersect the blob boundary
    var blobRanges []ds3Models.Range
    for _, r := range objectRanges {
        if rangeFinder.isIntersection(r, offset, length) {
            blobRanges = append(blobRanges, rangeFinder.getIntersection(r, offset, length))
        }
    }

    return blobRanges
}

func (BlobRangeFinderImpl) isIntersection(r ds3Models.Range, offset int64, length int64) bool {
    return !(r.Start > offset + length - 1 || r.End < offset)
}

// Returns the intersection of the range and the portion of the file assuming that said range exists
func (BlobRangeFinderImpl) getIntersection(r ds3Models.Range, offset int64, length int64) ds3Models.Range {
    var intersection ds3Models.Range
    if r.Start < offset {
        intersection.Start = offset
    } else {
        intersection.Start = r.Start
    }

    if r.End < offset + length - 1 {
        intersection.End = r.End
    } else {
        intersection.End = offset + length - 1
    }
    return intersection
}
