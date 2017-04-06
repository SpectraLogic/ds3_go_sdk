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

type deleteObjectList struct {
    XMLName xml.Name
    Objects []deleteObject `xml:"Object"`
}

type deleteObject struct {
    Key string `xml:"Key"`
}

func newDeleteObjectList(objectNames []string) *deleteObjectList {
    objects := make([]deleteObject, len(objectNames))
    for index, name := range objectNames {
        objects[index].Key = name
    }
    return &deleteObjectList{
        XMLName: xml.Name{Local:"Delete"},
        Objects: objects,
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

// Converts a list of object names into a request payload stream for delete objects
func buildDeleteObjectsPayload(objectNames []string) networking.ReaderWithSizeDecorator {
    deleteObjects := newDeleteObjectList(objectNames)
    xmlBytes, err := xml.Marshal(deleteObjects)
    if err != nil {
        // Should never happen
        panic(err)
    }
    return networking.BuildByteReaderWithSizeDecorator(xmlBytes)
}
