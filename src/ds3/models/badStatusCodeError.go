package models

import (
    "fmt"
    "ds3/networking"
)

type BadStatusCodeError struct {
    ExpectedStatusCode int
    ActualStatusCode int
    ErrorBody *Error
}

type Error struct {
    Code string
    Message string
    Resource string
    RequestId string
}

func buildBadStatusCodeError(ds3Response networking.Ds3Response, expectedStatusCode int) *BadStatusCodeError {
    var errorBody Error
    var errorBodyPtr *Error

    // Parse the body and if it worked then use the structure.
    err := parseResponseBody(ds3Response.Body(), &errorBody)
    if err == nil {
        errorBodyPtr = &errorBody
    }

    // Return the bad status code entity.
    return &BadStatusCodeError{
        expectedStatusCode,
        ds3Response.StatusCode(),
        errorBodyPtr,
    }
}

func (err BadStatusCodeError) Error() string {
    if err.ErrorBody != nil {
        return fmt.Sprintf(
            "Received a status code of %d when %d was expected. Error message: \"%s\"",
            err.ActualStatusCode,
            err.ExpectedStatusCode,
            err.ErrorBody.Message,
        )
    } else {
        return fmt.Sprintf(
            "Received a status code of %d when %d was expected. Could not parse the response for additional information.",
            err.ActualStatusCode,
            err.ExpectedStatusCode,
        )
    }
}

