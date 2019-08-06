package models

import (
    "github.com/SpectraLogic/core_go/test_helpers"
    "io"
    "strings"
    "testing"
)

type nopCloser struct {
    io.Reader
}

func (readCloser *nopCloser) Close() error {
    return nil
}

// GOSDK-25 unmarshaling optional enums in response payloads causes panic.
// This test is for validating the above jira is fixed and to catch future regressions.
func TestUnmarshalingTapeWithPreviousState(t *testing.T) {
    responseContent := "<Data><Tape><AssignedToStorageDomain>false</AssignedToStorageDomain>" +
        "<AvailableRawCapacity>10000</AvailableRawCapacity><BarCode>2f8f0dd0-c1db-4c72-9edc-cd23f7cccd2a</BarCode>" +
        "<BucketId/><DescriptionForIdentification/><EjectDate/><EjectLabel/><EjectLocation/><EjectPending/>" +
        "<FullOfData>false</FullOfData><Id>069ad1fb-8c46-4834-a8b1-f92bc9e66059</Id><LastAccessed/>" +
        "<LastCheckpoint>new</LastCheckpoint><LastModified/><LastVerified/><PartiallyVerifiedEndOfTape/>" +
        "<PartitionId>d6245cbe-b6ea-46e3-a974-87663016733c</PartitionId>" +
        "<PreviousState>eject_from_ee_pending</PreviousState><SerialNumber/><State>NORMAL</State>" +
        "<TakeOwnershipPending>false</TakeOwnershipPending><TotalRawCapacity>20000</TotalRawCapacity><Type>LTO5</Type>" +
        "<VerifyPending/><WriteProtected>false</WriteProtected></Tape></Data>"

    responseReadCloser := &nopCloser{Reader: strings.NewReader(responseContent)}

    // create the xml tree
    root, err := parseXmlTree(responseReadCloser)
    test_helpers.FatalOnError(t, err, "getting xml parse tree from response")

    // parse the response
    var aggErr AggregateError
    var tapeResponse GetTapesSpectraS3Response
    tapeResponse.TapeList.parse(root, &aggErr)

    test_helpers.AssertEqual(t, 0, len(aggErr.Errors), "expected no errors when parsing response")
    test_helpers.AssertEqual(t, 1, len(tapeResponse.TapeList.Tapes), "un-marshaled one tape")

    for _, err := range aggErr.Errors {
        t.Logf(err.Error())
    }
}
