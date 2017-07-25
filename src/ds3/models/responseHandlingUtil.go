// Copyright 2014-2017 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package models

import (
    "ds3/networking"
    "encoding/xml"
    "io"
    "io/ioutil"
)

type modelParser interface {
    parse(node *XmlNode, aggErr *AggregateError)
}

// Parses a response payload into the specified model parser. A best effort parsing
// is performed, and all errors that occur during parsing are captured within an aggregate error.
func parseResponsePayload(webResponse networking.WebResponse, parsedBody modelParser) error {
    // Clean up the response body.
    body := webResponse.Body()
    defer body.Close()

    // Create the xml tree
    root, err := parseXmlTree(body)
    if err != nil {
        return err
    }

    // Parse the response
    var aggErr AggregateError
    parsedBody.parse(root, &aggErr)
    return aggErr.GetErrors()
}

// Converts the contents of the reader into an xml node tree
func parseXmlTree(reader io.ReadCloser) (*XmlNode, error) {

    var xmlNode XmlNode
    dec := xml.NewDecoder(reader)
    xmlTreeErr := dec.Decode(&xmlNode)
    if xmlTreeErr != nil {
        return nil, xmlTreeErr
    }

    return &xmlNode, nil
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