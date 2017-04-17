package models

import (
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

// Parses the DS3 specific bulk command response.
func getObjectsFromBulkResponse(webResponse networking.WebResponse) ([][]Object, error) {
    // Parse the master object list response body.
    var mol masterobjectlist
    if err := readResponseBody(webResponse, &mol); err != nil {
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

