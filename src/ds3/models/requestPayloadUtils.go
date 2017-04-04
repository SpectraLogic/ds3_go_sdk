package models

import (
    "encoding/xml"
    "ds3/networking"
)

type CompleteMultipartUpload struct {
    Parts []Part `xml:"Part"`
}

type Part struct {
    PartNumber int
    ETag string
}

type ds3ObjectList struct {
    XMLName xml.Name
    Ds3Objects []Ds3Object `xml:"object"`
}

type Ds3Object struct {
    Name string `xml:"name,attr"`
    Size int64 `xml:"size,attr"`
}

func newDs3ObjectList(ds3objects []Ds3Object) *ds3ObjectList {
    return &ds3ObjectList{
        XMLName: xml.Name{Local:"objects"},
        Ds3Objects: ds3objects,
    }
}

// Converts the ds3 object list into a request payload stream.
func buildDs3ObjectListStream(ds3Objects []Ds3Object) networking.ReaderWithSizeDecorator {
    // Build the ds3 object list entity.
    objects := newDs3ObjectList(ds3Objects)

    // Create an xml document from the entity.
    xmlBytes, err := xml.Marshal(objects)
    if err != nil {
        //Should never happen
        panic(err)
    }

    // Create a ByteReaderWithSizeDecorator which the network layer expects.
    return networking.BuildByteReaderWithSizeDecorator(xmlBytes)
}

// Converts a string request payload into a request payload stream.
func buildStreamFromString(payload string) networking.ReaderWithSizeDecorator {
    return networking.BuildByteReaderWithSizeDecorator([]byte(payload))
}

// Converts the parts list into a request payload stream.
func buildPartsListStream(parts []Part) networking.ReaderWithSizeDecorator {
    // Create an xml document from the entity.
    xmlBytes, err := xml.Marshal(CompleteMultipartUpload{parts})
    if err != nil {
        //Should never happen
        panic(err)
    }

    // Create a ReaderWithSizeDecorator which the network layer expects.
    return networking.BuildByteReaderWithSizeDecorator(xmlBytes)
}
