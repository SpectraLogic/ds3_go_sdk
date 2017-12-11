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

package ds3

import (
    "encoding/xml"
    "bytes"
    "spectra/ds3_go_sdk/ds3/models"
)

// Converts the parts list into a request payload stream of format:
// <CompleteMultipartUpload><Part><PartNumber>partNumber</PartNumber><ETag>eTag</ETag></Part>...</CompleteMultipartUpload>
func buildPartsListStream(parts []models.Part) models.ReadCloserWithSizeDecorator {
    completeMultipartUpload := models.CompleteMultipartUpload{ Parts:parts }
    return marshalRequestPayload(completeMultipartUpload)
}

type ds3PutObjectList struct {
    XMLName xml.Name
    Ds3PutObjects []models.Ds3PutObject `xml:"Object"`
}

// Used to encapsulate a slice of Ds3PutObjects with xml marshaling information
func newDs3PutObjectList(ds3PutObjects []models.Ds3PutObject) *ds3PutObjectList {
    return &ds3PutObjectList{
        XMLName: xml.Name{Local:"Objects"},
        Ds3PutObjects: ds3PutObjects,
    }
}

// Converts the ds3 put object list into a request payload stream of format:
// <Objects><Object Name="o1" Size="2048"></Object><Object Name="o2" Size="2048"></Object>...</Objects>
func buildDs3PutObjectListStream(ds3PutObjects []models.Ds3PutObject) models.ReadCloserWithSizeDecorator {
    // Build the ds3 put object list entity.
    objects := newDs3PutObjectList(ds3PutObjects)
    return marshalRequestPayload(objects)
}



type ds3GetObjectList struct {
    XMLName xml.Name
    Ds3GetObjects []models.Ds3GetObject `xml:"Object"`
}

// Used to encapsulate a slice of Ds3GetObjects with xml marshaling information
func newDs3GetObjectList(ds3GetObjects []models.Ds3GetObject) *ds3GetObjectList {
    return &ds3GetObjectList{
        XMLName: xml.Name{Local:"Objects"},
        Ds3GetObjects: ds3GetObjects,
    }
}

// Converts string list into Ds3Object and marshals to request payload stream of format:
// <Objects><Object Name="o1"></Object><Object Name="o2"></Object>...</Objects>
func buildDs3ObjectStreamFromNames(objectNames []string) models.ReadCloserWithSizeDecorator {
    // Build the ds3 get object list entity.
    var objects []models.Ds3GetObject
    for _, name := range objectNames {
        objects = append(objects, models.NewDs3GetObject(name))
    }

    objectList := newDs3GetObjectList(objects)

    return marshalRequestPayload(objectList)
}

// Converts the ds3 get object list into a request payload stream of format:
// <Objects><Object Name="o1" Length="2" Offset="3"></Object><Object Name="o2" Length="3" offset="4"></Object>...</Objects>
func buildDs3GetObjectListStream(ds3GetObjects []models.Ds3GetObject) models.ReadCloserWithSizeDecorator {
    // Build the ds3 get object list entity.
    objects := newDs3GetObjectList(ds3GetObjects)
    return marshalRequestPayload(objects)
}

type deleteObjectList struct {
    XMLName xml.Name
    Objects []deleteObject `xml:"Object"`
}

type deleteObject struct {
    Key string `xml:"Key"`
}

// Used to encapsulate a list of object names with xml marshaling information
// for creating a delete objects request payload
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

// Converts a list of object names into a request payload stream for delete objects of format:
// <Delete><Object><Key>o1</Key></Object><Object><Key>o2</Key></Object>...</Delete>
func buildDeleteObjectsPayload(objectNames []string) models.ReadCloserWithSizeDecorator {
    deleteObjects := newDeleteObjectList(objectNames)
    return marshalRequestPayload(deleteObjects)
}

type idList struct {
    XMLName xml.Name
    IdList []string `xml:"Id"`
}

// Used to encapsulate a list of ids with xml marshaling information
func newIdList(objectNames []string) *idList {
    return &idList{
        XMLName: xml.Name{Local:"Ids"},
        IdList:objectNames,
    }
}

// Converts a list of ids into a request payload stream of format:
// <Ids><Id>id1</Id><Id>id2</Id>...</Ids>
func buildIdListPayload(ids []string) models.ReadCloserWithSizeDecorator {
    idList := newIdList(ids)
    return marshalRequestPayload(idList)
}

// Converts a string request payload into a request payload stream.
func buildStreamFromString(payload string) models.ReadCloserWithSizeDecorator {
    return BuildByteReaderWithSizeDecorator([]byte(payload))
}

func marshalRequestPayload(model interface{}) models.ReadCloserWithSizeDecorator {
    xmlBytes, err := xml.Marshal(model)
    if err != nil {
        // Should never happen
        panic(err)
    }
    return BuildByteReaderWithSizeDecorator(xmlBytes)
}

// Defines a ReaderWithSizeDecorator based on an array of bytes.
type byteReaderWithSizeDecorator struct {
    reader *bytes.Reader
    size int64
}

func BuildByteReaderWithSizeDecorator(contentBytes []byte) models.ReadCloserWithSizeDecorator {
    return &byteReaderWithSizeDecorator{
        bytes.NewReader(contentBytes),
        int64(len(contentBytes)),
    }
}

func (byteReaderWithSizeDecorator *byteReaderWithSizeDecorator) Read(b []byte) (int, error) {
    return byteReaderWithSizeDecorator.reader.Read(b)
}

func (byteReaderWithSizeDecorator) Close() error {
    return nil
}

func (byteReaderWithSizeDecorator *byteReaderWithSizeDecorator) Seek(offset int64, whence int) (int64, error) {
    return byteReaderWithSizeDecorator.reader.Seek(offset, whence)
}

func (byteReaderWithSizeDecorator *byteReaderWithSizeDecorator) Size() (int64, error) {
    return byteReaderWithSizeDecorator.size, nil
}