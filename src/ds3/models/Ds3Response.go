package models

import (
    "io/ioutil"
    "fmt"
    "net/http"
    "encoding/xml"
)

func readResponseBody(response *http.Response, expectedStatusCode int, body interface{}) (error) {
    stream := response.Body
    defer stream.Close()

    if response.StatusCode == expectedStatusCode {
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
    } else {
        // Return a status code error
        return fmt.Errorf(
            "Expected HTTP status code %d but got %d.",
            expectedStatusCode,
            response.StatusCode,
        )
    }
}

