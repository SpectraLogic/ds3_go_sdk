package models

import (
    "io/ioutil"
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
    // Clean up the response body if it exists.
    if response.Body != nil {
        defer response.Body.Close()
    }

    // Check the status code.
    if err := checkStatusCode(response, expectedStatusCode); err != nil {
        return err
    }

    // Parse the response
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

