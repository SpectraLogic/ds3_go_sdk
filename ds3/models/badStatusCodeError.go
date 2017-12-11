package models

import (
    "fmt"
    "io"
    "io/ioutil"
    "encoding/xml"
)

type BadStatusCodeError struct {
    ExpectedStatusCode []int
    ActualStatusCode int
    ErrorBody *Error
}

func buildBadStatusCodeError(webResponse WebResponse, expectedStatusCodes []int) *BadStatusCodeError {
    var errorBody Error
    var errorBodyPtr *Error

    // Parse the body and if it worked then use the structure.
    err := parseErrorResponseBody(webResponse.Body(), &errorBody)
    if err == nil {
        errorBodyPtr = &errorBody
    }

    // Return the bad status code entity.
    return &BadStatusCodeError{
        expectedStatusCodes,
        webResponse.StatusCode(),
        errorBodyPtr,
    }
}

func parseErrorResponseBody(reader io.ReadCloser, body interface{}) error {
    // Get the bytes or forward the error.
    bytes, err := ioutil.ReadAll(reader)
    if err != nil {
        return err
    }

    // Deserialize or forward the error.
    if err = xml.Unmarshal(bytes, body); err != nil {
        return err
    }

    return nil
}

func (err BadStatusCodeError) Error() string {
    if err.ErrorBody != nil && err.ErrorBody.Message != nil {
        return fmt.Sprintf(
            "Received a status code of %d when %v was expected. Error message: \"%s\"",
            err.ActualStatusCode,
            err.ExpectedStatusCode,
            *err.ErrorBody.Message,
        )
    } else {
        return fmt.Sprintf(
            "Received a status code of %d when %v was expected. Could not parse the response for additional information.",
            err.ActualStatusCode,
            err.ExpectedStatusCode,
        )
    }
}
