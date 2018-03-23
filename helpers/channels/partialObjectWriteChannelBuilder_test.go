package channels

import (
    "testing"
    "io/ioutil"
    "os"
    "spectra/ds3_go_sdk/ds3_utils/ds3Testing"
    ds3Models "spectra/ds3_go_sdk/ds3/models"
    "reflect"
)

func TestPartialObjectWriteChannel(t *testing.T) {
    file, err := ioutil.TempFile(os.TempDir(), "goTest")
    ds3Testing.AssertNilError(t, err)
    defer file.Close()
    defer os.Remove(file.Name())

    ranges := []ds3Models.Range {
        {1, 3},
        {4, 5},
        {8, 10},
        {15, 20},
    }

    builder := NewPartialObjectChannelBuilder(file.Name(), ranges)
    originalContent := []byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

    for _, curRange := range ranges {
        channel, err := builder.GetChannel(curRange.Start)
        if err == nil {
            _, err := channel.Write(originalContent[curRange.Start:curRange.End+1])
            ds3Testing.AssertNilError(t, err)
        }
    }

    result, err := ioutil.ReadAll(file)
    expected := []byte {'b', 'c', 'd', 'e', 'f', 'i', 'j', 'k', 'p', 'q', 'r', 's', 't', 'u'}
    if !reflect.DeepEqual(expected, result) {
        t.Errorf("Expected '%s' but got '%s'", string(expected), string(result))
    }
}
