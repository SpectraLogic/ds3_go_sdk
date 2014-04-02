package models

import (
    "io"
    "io/ioutil"
    "encoding/xml"
)


//TODO: improve error handling
func readResponseBody(stream io.ReadCloser, body interface{}) (error) {
    defer stream.Close()

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

