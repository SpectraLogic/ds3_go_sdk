package channels

import (
    "testing"
    ds3Models "spectra/ds3_go_sdk/ds3/models"
    "reflect"
    "spectra/ds3_go_sdk/ds3_utils/ds3Testing"
    "fmt"
)

func getTestRanges() []ds3Models.Range {
    return []ds3Models.Range {
        {1, 5},
        {4, 7},
        {10, 11},
        {12, 13},
        {15, 16},
    }
}

func TestCreatingRangeMap(t *testing.T) {
    result := newPartialObjectRangeMap(getTestRanges())

    // expected map
    expected := make(map[ds3Models.Range]int64)
    expected[ds3Models.Range{Start: 1, End: 7}] = 0
    expected[ds3Models.Range{Start: 10, End: 13}] = 7
    expected[ds3Models.Range{Start: 15, End: 16}] = 11

    if !reflect.DeepEqual(result.rangeMap, expected) {
        t.Errorf("Expected map '%v' but got '%v'", expected, result)
    }
}

func TestGettingDestinationOffset(t *testing.T) {
    mapper := newPartialObjectRangeMap(getTestRanges())

    type testResponse struct {
        input    int64
        expected int64
        error    bool
    }

    expected := []testResponse {
        { 0, -1, true },
        { 1, 0, false },
        { 2, 1, false },
        { 3, 2, false },
        { 4, 3, false },
        { 5, 4, false },
        { 6, 5, false },
        { 7, 6, false },
        { 8, -1, true },
        { 9, -1, true },
        { 10, 7, false },
        { 11, 8, false },
        { 12, 9, false },
        { 13, 10, false },
        { 14, -1, true },
        { 15, 11, false },
        { 16, 12, false },
        { 17, -1, true },
    }

    for _, exp := range expected {
        result, err := mapper.getDestinationOffset(exp.input)
        ds3Testing.AssertBool(t, fmt.Sprintf("input '%d' errored", exp.input), exp.error, err != nil)
        ds3Testing.AssertInt64(t, "destination offset", exp.expected, result)
    }
}

func TestGettingDestinationOffsetWithUnorderedRanges(t *testing.T) {
    ranges := []ds3Models.Range {
        {12, 13},
        {1, 5},
        {10, 11},
        {4, 7},
        {15, 16},
    }

    mapper := newPartialObjectRangeMap(ranges)

    type testResponse struct {
        input    int64
        expected int64
        error    bool
    }

    expected := []testResponse {
        { 0, -1, true },
        { 1, 0, false },
        { 2, 1, false },
        { 3, 2, false },
        { 4, 3, false },
        { 5, 4, false },
        { 6, 5, false },
        { 7, 6, false },
        { 8, -1, true },
        { 9, -1, true },
        { 10, 7, false },
        { 11, 8, false },
        { 12, 9, false },
        { 13, 10, false },
        { 14, -1, true },
        { 15, 11, false },
        { 16, 12, false },
        { 17, -1, true },
    }

    for _, exp := range expected {
        result, err := mapper.getDestinationOffset(exp.input)
        ds3Testing.AssertBool(t, fmt.Sprintf("input '%d' errored", exp.input), exp.error, err != nil)
        ds3Testing.AssertInt64(t, "destination offset", exp.expected, result)
    }
}