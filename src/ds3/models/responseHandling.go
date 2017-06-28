//TODO delete once replaced with responseHandlingUtil

package models

import (
    "io"
    "io/ioutil"
    "encoding/xml"
    "ds3/networking"
)

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

func getResponseBodyAsString(webResponse networking.WebResponse) (string, error) {
    // Clean up the response body.
    body := webResponse.Body()
    defer body.Close()

    // Get the bytes or forward the error
    bytes, err := ioutil.ReadAll(body)
    if err != nil {
        return "", err
    }

    // Convert the bytes into a string
    return string(bytes), nil
}
