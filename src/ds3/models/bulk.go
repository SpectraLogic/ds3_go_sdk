package models

import (
    "encoding/xml"
    "net/http"
    "ds3/networking"
)

// Represents the XML document for both the bulk request and response data for
// bulk gets and bulk puts.
type masterobjectlist struct {
    ObjectLists []objects `xml:"objects"`
}

type objects struct {
    Objects []object `xml:"object"`
}

type object struct {
    Name string `xml:"name,attr"`
    Size int64 `xml:"size,attr"`
}

// Converts the ds3 object list to an object we can send in our request.
func buildObjectListStream(ds3Objects []Object) networking.SizedReadCloser {
    // Create an array of objects to put into the master object list.
    molObjects := make([]object, len(ds3Objects))

    // Copy the important ds3 object contents into the master object list objects.
    for i, obj := range ds3Objects {
        molObjects[i] = object{obj.Key, obj.Size}
    }

    // Build the mast object list entity.
    mol := objects{molObjects}

    // Create an xml document from the entity.
    xmlBytes, err := xml.Marshal(mol)
    if err != nil {
        panic(err)
    }

    // Create a SizedReadCloser which the network layer expects.
    return networking.BuildSizedReadCloser(xmlBytes)
}

// Parses the DS3 specific bulk command response.
func getObjectsFromBulkResponse(ds3Response networking.Ds3Response) ([][]Object, error) {
    // Parse the master object list response body.
    var mol masterobjectlist
    if err := readResponseBody(ds3Response, http.StatusOK, &mol); err != nil {
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

