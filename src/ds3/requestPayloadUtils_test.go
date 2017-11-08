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
    "testing"
    "io/ioutil"
    "ds3/networking"
    "ds3/models"
)

func verifyStreamContent(t *testing.T, contentStream networking.ReaderWithSizeDecorator, expected string) {
    bs, readErr := ioutil.ReadAll(contentStream)
    if readErr != nil {
        t.Fatalf("Unexpected error: '%s'.", readErr.Error())
    }
    result := string(bs)
    if result != expected {
        t.Fatalf("Expected '%s' but got '%s'.", expected, result)
    }
}

func TestBuildPartsListStream(t *testing.T) {
    expected := "<CompleteMultipartUpload><Part><PartNumber>1</PartNumber><ETag>eTag1</ETag></Part><Part><PartNumber>2</PartNumber><ETag>eTag2</ETag></Part><Part><PartNumber>3</PartNumber><ETag>eTag3</ETag></Part></CompleteMultipartUpload>"

    parts := []models.Part{
        {PartNumber:1, ETag:"eTag1"},
        {PartNumber:2, ETag:"eTag2"},
        {PartNumber:3, ETag:"eTag3"},
    }

    contentStream := buildPartsListStream(parts)

    defer contentStream.Close()
    verifyStreamContent(t, contentStream, expected)
}

func TestBuildDs3PutObjectListStream(t *testing.T) {
    expected := "<Objects><Object Name=\"o1\" Size=\"1\"></Object><Object Name=\"o2\" Size=\"2\"></Object><Object Name=\"o3\" Size=\"3\"></Object></Objects>"

    ds3PutObjects := []models.Ds3PutObject{
        {Name:"o1", Size:1},
        {Name:"o2", Size:2},
        {Name:"o3", Size:3},
    }

    contentStream := buildDs3PutObjectListStream(ds3PutObjects)

    defer contentStream.Close()
    verifyStreamContent(t, contentStream, expected)
}

func TestBuildDs3ObjectStreamFromNames(t *testing.T) {
    expected := "<Objects><Object Name=\"o1\"></Object><Object Name=\"o2\"></Object><Object Name=\"o3\"></Object></Objects>"

    names := []string {"o1", "o2", "o3"}

    contentStream := buildDs3ObjectStreamFromNames(names)

    defer contentStream.Close()
    verifyStreamContent(t, contentStream, expected)
}

func TestBuildDs3GetObjectListStream(t *testing.T) {
    expected := "<Objects><Object Name=\"o1\" Length=\"10\" Offset=\"1\"></Object><Object Name=\"o2\" Length=\"20\" Offset=\"2\"></Object><Object Name=\"o3\" Length=\"30\" Offset=\"3\"></Object></Objects>"

    ds3GetObjects := []models.Ds3GetObject {
        models.NewPartialDs3GetObject("o1", 10, 1),
        models.NewPartialDs3GetObject("o2", 20, 2),
        models.NewPartialDs3GetObject("o3", 30, 3),
    }

    contentStream := buildDs3GetObjectListStream(ds3GetObjects)

    defer contentStream.Close()
    verifyStreamContent(t, contentStream, expected)
}

func TestBuildDeleteObjectsPayload(t *testing.T) {
    expected := "<Delete><Object><Key>o1</Key></Object><Object><Key>o2</Key></Object><Object><Key>o3</Key></Object></Delete>"

    names := []string {"o1", "o2", "o3"}

    contentStream := buildDeleteObjectsPayload(names)

    defer contentStream.Close()
    verifyStreamContent(t, contentStream, expected)
}

func TestBuildIdListPayload(t *testing.T) {
    expected := "<Ids><Id>id1</Id><Id>id2</Id><Id>id3</Id></Ids>"

    ids := []string{"id1", "id2", "id3"}

    contentStream := buildIdListPayload(ids)

    defer contentStream.Close()
    verifyStreamContent(t, contentStream, expected)
}
