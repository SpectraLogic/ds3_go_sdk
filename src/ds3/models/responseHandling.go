package models

import (
    "io"
    "io/ioutil"
    "encoding/xml"
    "ds3/networking"
)

func checkStatusCode(webResponse networking.WebResponse, expectedStatusCodes []int) error {
    code := webResponse.StatusCode()
    for _, expectedCode := range expectedStatusCodes {
        if code == expectedCode {
            return nil
        }
    }
    return buildBadStatusCodeError(webResponse, expectedStatusCodes)
}

func readResponseBody(webResponse networking.WebResponse, parsedBody interface{}) error {
    // Clean up the response body.
    body := webResponse.Body()
    defer body.Close()

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

