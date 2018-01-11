package helpers

type Range struct {
    start int64
    end int64
}

/*
type Ranges []Range

func (ranges Ranges) Len() int {
    return len(ranges)
}

func (ranges Ranges) Swap(i int, j int) {
    ranges[i], ranges[j] = ranges[j], ranges[i]
}

// Determines if element with index i should sort before the element with index j
func (ranges Ranges) Less(i int, j int) bool {
    return ranges[i].start < ranges[j].start
}

type RangeHelper struct {
    ranges Ranges
}

*/

type Ranges struct {
    rangesSlice []Range
}

// inserts the new range into the slice
func (ranges *Ranges) insert(r Range) {
    //todo
}

type RangeHelper struct {
    rangeMap map[string]*Ranges
}

func NewRangeHelper() RangeHelper {
    return RangeHelper{rangeMap:make(map[string]*Ranges)}
}

func (rangeHelper *RangeHelper) add(key string, r Range) {
    if rangeHelper.rangeMap[key] == nil {
        rangeHelper.rangeMap[key] = &Ranges{ rangesSlice: make([]Range, 1)}
    }

    rangeHelper.rangeMap[key].insert(r)
}