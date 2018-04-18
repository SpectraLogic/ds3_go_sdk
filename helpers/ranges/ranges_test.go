package ranges

import (
    "testing"
    "spectra/ds3_go_sdk/ds3/models"
    "sort"
    "spectra/ds3_go_sdk/ds3_utils/ds3Testing"
    "fmt"
)

func getTestRanges() []models.Range {
    return []models.Range {
        {1, 3},
        {0, 2}, // not in order regarding start position
        {5, 7},
        {5, 6}, // not in order regarding end position
        {10, 11},
        {10, 11}, // duplicate
        {12, 15},
        {13, 14},
        {20, 25},
        {21, 22}, // subset of other range
        {30, 31}, // shared boundary
        {31, 32},
    }
}

func verifyRangeLists(first []models.Range, second []models.Range) bool {
    if len(first) != len(second) {
        return false
    }

    for i := 0; i < len(first); i++ {
        if first[i] != second[i] {
            return false
        }
    }
    return true
}

func TestSortingRanges(t *testing.T) {
    ranges := getTestRanges()

    sort.Sort(Ranges(ranges))

    expected := []models.Range {
        {0, 2},
        {1, 3},
        {5, 6},
        {5, 7},
        {10, 11},
        {10, 11},
        {12, 15},
        {13, 14},
        {20, 25},
        {21, 22},
        {30, 31},
        {31, 32},
    }
    ds3Testing.AssertBool(t, "sorted list is correct", true, verifyRangeLists(ranges, expected))
}

func TestCollapsingRanges(t *testing.T) {
    ranges := Ranges(getTestRanges())

    collapser := RangeCollapserImpl{}
    newRanges := collapser.Collapse(ranges)

    ds3Testing.AssertBool(t, "original list is unmodified", true, verifyRangeLists(ranges, getTestRanges()))

    expected := []models.Range {
        {0, 3},
        {5, 7},
        {10, 15},
        {20, 25},
        {30, 32},
    }
    ds3Testing.AssertBool(t, "collapsed list is correct", true, verifyRangeLists(newRanges, expected))
}

func TestIsIntersection(t *testing.T) {
    rangeFinder := BlobRangeFinderImpl{}

    ds3Testing.AssertBool(t, "intersection 0-3, 5-9", false, rangeFinder.isIntersection(models.Range{0, 3}, 5, 5))
    ds3Testing.AssertBool(t, "intersection 3-4, 5-9", false, rangeFinder.isIntersection(models.Range{3, 4}, 5, 5))

    ds3Testing.AssertBool(t, "intersection 3-5, 5-9", true, rangeFinder.isIntersection(models.Range{3, 5}, 5, 5))
    ds3Testing.AssertBool(t, "intersection 5-7, 5-9", true, rangeFinder.isIntersection(models.Range{5, 7}, 5, 5))
    ds3Testing.AssertBool(t, "intersection 6-7, 5-9", true, rangeFinder.isIntersection(models.Range{6, 7}, 5, 5))
    ds3Testing.AssertBool(t, "intersection 8-10, 5-9", true, rangeFinder.isIntersection(models.Range{8, 10}, 5, 5))
    ds3Testing.AssertBool(t, "intersection 9-11, 5-9", true, rangeFinder.isIntersection(models.Range{9, 11}, 5, 5))
    ds3Testing.AssertBool(t, "intersection 3-11, 5-9", true, rangeFinder.isIntersection(models.Range{3, 11}, 5, 5))

    ds3Testing.AssertBool(t, "intersection 10-12, 5-9", false, rangeFinder.isIntersection(models.Range{10, 12}, 5, 5))
    ds3Testing.AssertBool(t, "intersection 11-15, 5-9", false, rangeFinder.isIntersection(models.Range{11, 15}, 5, 5))
}

func TestGetIntersection(t *testing.T) {
    input := []models.Range {
        {3, 7},
        {6, 7},
        {7, 12},
        {3, 12},
    }

    expected := []models.Range {
        {5, 7},
        {6, 7},
        {7, 9},
        {5, 9},
    }

    rangeFinder := BlobRangeFinderImpl{}

    for i := 0; i < len(input); i++ {
        result := rangeFinder.getIntersection(input[i], 5, 5)
        ds3Testing.AssertInt64(t, fmt.Sprintf("Input%v.Start", input[i]), expected[i].Start, result.Start)
        ds3Testing.AssertInt64(t, fmt.Sprintf("Input%v.End", input[i]), expected[i].End, result.End)
    }
}