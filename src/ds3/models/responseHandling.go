package models

import (
    "io"
    "io/ioutil"
    "encoding/xml"
    "ds3/networking"
)

func checkStatusCode(webResponse networking.WebResponse, expectedStatusCode int) error {
    if webResponse.StatusCode() == expectedStatusCode {
        return nil
    } else {
        return buildBadStatusCodeError(webResponse, expectedStatusCode)
    }
}

func readResponseBody(webResponse networking.WebResponse, expectedStatusCode int, parsedBody interface{}) error {
    // Clean up the response body.
    body := webResponse.Body()
    defer body.Close()

    // Check the status code.
    if err := checkStatusCode(webResponse, expectedStatusCode); err != nil {
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

