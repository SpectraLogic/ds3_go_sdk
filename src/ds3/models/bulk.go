package models

import (
    "encoding/xml"
    "net/http"
    "ds3/networking"
)

// Represents the XML document for the bulk response data for
// bulk gets and bulk puts.
//TODO delete: will be auto-generated
type masterobjectlist struct {
    ObjectLists []objects `xml:"objects"`
}

//TODO delete: will be auto-generated
type objects struct {
    Objects []bulkObject `xml:"object"`
}

//TODO delete: will be auto-generated
type bulkObject struct {
    Name string `xml:"name,attr"`
    Size int64 `xml:"size,attr"`
}

type ds3ObjectList struct {
    XMLName xml.Name
    Ds3Objects []Ds3Object `xml:"object"`
}

func newDs3ObjectList(ds3objects []Ds3Object) *ds3ObjectList {
    return &ds3ObjectList{
        XMLName: xml.Name{Local:"objects"},
        Ds3Objects: ds3objects,
    }
}

type Ds3Object struct {
    Name string `xml:"name,attr"`
    Size int64 `xml:"size,attr"`
}

// Converts the ds3 object list to an object we can send in our request.
func buildDs3ObjectListStream(ds3Objects []Ds3Object) networking.ReaderWithSizeDecorator {
    // Build the ds3 object list entity.
    objects := newDs3ObjectList(ds3Objects)

    // Create an xml document from the entity.
    xmlBytes, err := xml.Marshal(objects)
    if err != nil {
        panic(err)
    }

    // Create a ByteReaderWithSizeDecorator which the network layer expects.
    return networking.BuildByteReaderWithSizeDecorator(xmlBytes)
}

// Parses the DS3 specific bulk command response.
func getObjectsFromBulkResponse(webResponse networking.WebResponse) ([][]Object, error) {
    // Parse the master object list response body.
    var mol masterobjectlist
    if err := readResponseBody(webResponse, http.StatusOK, &mol); err != nil {
        return nil, err
    }

    // Create an array of ds3 object arrays.
    ds3ObjectArrayArray := make([][]Object, len(mol.ObjectLists))

    // For each object list in the deserialized XML...
    for i, molObjs := range mol.ObjectLists {

        // Create an array of ds3 objects.
        ds3ObjectArray := make([]Object, len(molObjs.Objects))

        // Copy the relevant contents of each deserialized object into new ds3 objects.
        for i, obj := range molObjs.Objects {
            ds3ObjectArray[i] = Object{Key: obj.Name, Size: obj.Size}
        }

        // Put the result into the array of arrays.
        ds3ObjectArrayArray[i] = ds3ObjectArray
    }
    return ds3ObjectArrayArray, nil
}

