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

// Converts the parts list into a request payload stream of format:
// <CompleteMultipartUpload><Part><PartNumber>partNumber</PartNumber><ETag>eTag</ETag></Part>...</CompleteMultipartUpload>
func buildPartsListStream(parts []Part) networking.ReaderWithSizeDecorator {
    completeMultipartUpload := CompleteMultipartUpload{ Parts:parts }
    return marshalRequestPayload(completeMultipartUpload)
}


type ds3PutObjectList struct {
    XMLName xml.Name
    Ds3PutObjects []Ds3PutObject `xml:"Object"`
}

type Ds3PutObject struct {
    Name string `xml:"Name,attr"`
    Size int64 `xml:"Size,attr"`
}

func newDs3PutObjectList(ds3PutObjects []Ds3PutObject) *ds3PutObjectList {
    return &ds3PutObjectList{
        XMLName: xml.Name{Local:"Objects"},
        Ds3PutObjects: ds3PutObjects,
    }
}

// Converts the ds3 put object list into a request payload stream of format:
// <Objects><Object Name="o1" Size="2048"></Object><Object Name="o2" Size="2048"></Object>...</Objects>
func buildDs3PutObjectListStream(ds3PutObjects []Ds3PutObject) networking.ReaderWithSizeDecorator {
    // Build the ds3 put object list entity.
    objects := newDs3PutObjectList(ds3PutObjects)
    return marshalRequestPayload(objects)
}

type Ds3GetObject struct {
    Name string `xml:"Name,attr"`
    Length *int64 `xml:"Length,attr,omitempty"`
    Offset *int64 `xml:"Offset,attr,omitempty"`
}

type ds3GetObjectList struct {
    XMLName xml.Name
    Ds3GetObjects []Ds3GetObject `xml:"Object"`
}

func NewDs3GetObject(name string) Ds3GetObject {
    return Ds3GetObject{Name:name}
}

func newDs3GetObjectList(ds3GetObjects []Ds3GetObject) *ds3GetObjectList {
    return &ds3GetObjectList{
        XMLName: xml.Name{Local:"Objects"},
        Ds3GetObjects: ds3GetObjects,
    }
}

// Creates a Ds3GetObject used for partial objects
func NewPartialDs3GetObject(name string, length int64, offset int64) Ds3GetObject {
    return Ds3GetObject{
        Name:name,
        Length:&length,
        Offset:&offset,
    }
}

// Converts string list into Ds3Object and marshals to request payload stream of format:
// <Objects><Object Name="o1"></Object><Object Name="o2"></Object>...</Objects>
func buildDs3ObjectStreamFromNames(objectNames []string) networking.ReaderWithSizeDecorator {
    // Build the ds3 get object list entity.
    var objects []Ds3GetObject
    for _, name := range objectNames {
        objects = append(objects, NewDs3GetObject(name))
    }

    objectList := newDs3GetObjectList(objects)

    return marshalRequestPayload(objectList)
}

// Converts the ds3 get object list into a request payload stream of format:
// <Objects><Object Name="o1" Length="2" Offset="3"></Object><Object Name="o2" Length="3" offset="4"></Object>...</Objects>
func buildDs3GetObjectListStream(ds3GetObjects []Ds3GetObject) networking.ReaderWithSizeDecorator {
    // Build the ds3 get object list entity.
    objects := newDs3GetObjectList(ds3GetObjects)
    return marshalRequestPayload(objects)
}

type Ds3VerifyObject struct {
    Name string `xml:"Name,attr"`
    Length int64 `xml:"Length,attr"`
}

type ds3VerifyObjectList struct {
    XMLName xml.Name
    Ds3VerifyObjects []Ds3VerifyObject `xml:"Object"`
}

func newDs3VerifyObjectList(ds3VerifyObjects []Ds3VerifyObject) *ds3VerifyObjectList {
    return &ds3VerifyObjectList{
        XMLName: xml.Name{Local:"Objects"},
        Ds3VerifyObjects: ds3VerifyObjects,
    }
}

// Converts the ds3 put object list into a request payload stream of format:
// <Objects><Object Name="o1" Length="2048"></Object><Object Name="o2" Length="2048"></Object>...</Objects>
func buildDs3VerifyObjectListStream(ds3VerifyObjects []Ds3VerifyObject) networking.ReaderWithSizeDecorator {
    // Build the ds3 put object list entity.
    objects := newDs3VerifyObjectList(ds3VerifyObjects)
    return marshalRequestPayload(objects)
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

// Converts a list of object names into a request payload stream for delete objects of format:
// <Delete><Object><Key>o1</Key></Object><Object><Key>o2</Key></Object>...</Delete>
func buildDeleteObjectsPayload(objectNames []string) networking.ReaderWithSizeDecorator {
    deleteObjects := newDeleteObjectList(objectNames)
    return marshalRequestPayload(deleteObjects)
}

type idList struct {
    XMLName xml.Name
    IdList []string `xml:"Id"`
}

func newIdList(objectNames []string) *idList {
    return &idList{
        XMLName: xml.Name{Local:"Ids"},
        IdList:objectNames,
    }
}

// Converts a list of ids into a request payload stream of format:
// <Ids><Id>id1</Id><Id>id2</Id>...</Ids>
func buildIdListPayload(ids []string) networking.ReaderWithSizeDecorator {
    idList := newIdList(ids)
    return marshalRequestPayload(idList)
}

// Converts a string request payload into a request payload stream.
func buildStreamFromString(payload string) networking.ReaderWithSizeDecorator {
    return networking.BuildByteReaderWithSizeDecorator([]byte(payload))
}

func marshalRequestPayload(model interface{}) networking.ReaderWithSizeDecorator {
    xmlBytes, err := xml.Marshal(model)
    if err != nil {
        // Should never happen
        panic(err)
    }
    return networking.BuildByteReaderWithSizeDecorator(xmlBytes)
}
