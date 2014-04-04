package models

import (
    "io/ioutil"
    "fmt"
    "net/http"
    "encoding/xml"
)

//TODO: read the response contents to make better error messages.
func checkStatusCode(response *http.Response, expectedStatusCode int) error {
    if response.StatusCode == expectedStatusCode {
        return nil
    } else {
        // Return a status code error
        return fmt.Errorf(
            "Expected HTTP status code %d but got %d.",
            expectedStatusCode,
            response.StatusCode,
        )
    }
}

func readResponseBody(response *http.Response, expectedStatusCode int, body interface{}) error {
    stream := response.Body
    defer stream.Close()

    // Check the status code.
    if err := checkStatusCode(response, expectedStatusCode); err != nil {
        return err
    }

    // Get the bytes or forward the error.
    bytes, err := ioutil.ReadAll(stream)
    if err != nil {
        return err
    }

    // Deserialize or forward the error.
    if err = xml.Unmarshal(bytes, body); err != nil {
        return err
    }

    // Return the result.
    return nil
}

