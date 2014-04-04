package models

import (
    "bytes"
    "encoding/xml"
    "net/http"
)

func newBulkResponse(response *http.Response) ([][]Object, error) {
    var body masterobjectlist
    if err := readResponseBody(response, http.StatusOK, &body); err != nil {
        return nil, err
    }
    return convertMolToObjects(body), nil
}

func buildMolReader(objects []Object) SizedReadCloser {
    molObjects := make([]object, len(objects))
    for i, obj := range(objects) {
        molObjects[i] = object{obj.Key, obj.Size}
    }
    mol := masterobjectlist{[]objectList{objectList{molObjects}}}
    xmlBytes, err := xml.Marshal(mol)
    if err != nil {
        panic(err)
    }
    return &bytesSizedReadCloser{bytes.NewReader(xmlBytes), int64(len(xmlBytes))}
}

func convertMolToObjects(mol masterobjectlist) [][]Object {
    objectss := make([][]Object, len(mol.Objects))
    for i, objs := range mol.Objects {
        objectSlice := make([]Object, len(objs.Objects))
        for i, obj := range objs.Objects {
            objectSlice[i] = Object{Key: obj.Name, Size: obj.Size}
        }
        objectss[i] = objectSlice
    }
    return objectss
}

type bytesSizedReadCloser struct {
    reader *bytes.Reader
    size int64
}

func (self bytesSizedReadCloser) Read(b []byte) (int, error) {
    return self.reader.Read(b)
}

func (bytesSizedReadCloser) Close() error {
    return nil
}

func (self bytesSizedReadCloser) Size() int64 {
    return self.size
}

type masterobjectlist struct {
    Objects []objectList `xml:"objects"`
}

type objectList struct {
    Objects []object `xml:"object"`
}

type object struct {
    Name string `xml:"name,attr"`
    Size int64 `xml:"size,attr"`
}

