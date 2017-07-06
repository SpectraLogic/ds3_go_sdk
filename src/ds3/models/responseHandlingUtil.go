package models

import (
    "ds3/networking"
    "encoding/xml"
    "io"
)

type modelParser interface {
    parse(node *XmlNode, aggErr *AggregateError)
}

//todo update name and test
func parseResponsePayload2(webResponse networking.WebResponse, parsedBody modelParser) error {
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