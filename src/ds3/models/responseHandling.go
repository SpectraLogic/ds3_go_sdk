package models

import (
    "io"
    "io/ioutil"
    "encoding/xml"
    "ds3/network"
)

func checkStatusCode(response net.Response, expectedStatusCode int) error {
    if response.StatusCode() == expectedStatusCode {
        return nil
    } else {
        return buildBadStatusCodeError(response, expectedStatusCode)
    }
}

func readResponseBody(response net.Response, expectedStatusCode int, parsedBody interface{}) error {
    // Clean up the response body.
    body := response.Body()
    defer body.Close()

    // Check the status code.
    if err := checkStatusCode(response, expectedStatusCode); err != nil {
        return err
    }

    // Parse the response
    return parseResponseBody(body, parsedBody)
}

func parseResponseBody(reader io.ReadCloser, body interface{}) error {
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

