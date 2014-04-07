package models

import (
    "io/ioutil"
    "fmt"
    "net/http"
    "encoding/xml"
)

func checkStatusCode(response *http.Response, expectedStatusCode int) error {
    if response.StatusCode == expectedStatusCode {
        return nil
    } else {
        return buildBadStatusCodeError(response, expectedStatusCode)
    }
}

func readResponseBody(response *http.Response, expectedStatusCode int, body interface{}) error {
    if response.Body != nil {
        defer response.Body.Close()
    }

    // Check the status code.
    if err := checkStatusCode(response, expectedStatusCode); err != nil {
        return err
    }

    return parseResponseBody(response, body)
}

func parseResponseBody(response *http.Response, body interface{}) error {
    // Get the bytes or forward the error.
    bytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return err
    }

    // Deserialize or forward the error.
    if err = xml.Unmarshal(bytes, body); err != nil {
        return err
    }

    return nil
}

func buildBadStatusCodeError(response *http.Response, expectedStatusCode int) *BadStatusCodeError {
    var errorBody Error
    var errorBodyPtr *Error

    // Parse the body and if it worked then use the structure.
    err := parseResponseBody(response, &errorBody)
    if err == nil {
        errorBodyPtr = &errorBody
    }

    // Return the bad status code entity.
    return &BadStatusCodeError{
        expectedStatusCode,
        response.StatusCode,
        errorBodyPtr,
    }
}

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

