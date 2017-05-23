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

// This code is auto-generated, do not modify

package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
    "strconv"
)

type ListMultiPartUploadPartsRequest struct {
    bucketName string
    objectName string
    maxParts int
    partNumberMarker *int
    uploadId string
    queryParams *url.Values
}

func NewListMultiPartUploadPartsRequest(bucketName string, objectName string, uploadId string) *ListMultiPartUploadPartsRequest {
    queryParams := &url.Values{}
    queryParams.Set("upload_id", uploadId)

    return &ListMultiPartUploadPartsRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
        queryParams: queryParams,
    }
}

func (listMultiPartUploadPartsRequest *ListMultiPartUploadPartsRequest) WithMaxParts(maxParts int) *ListMultiPartUploadPartsRequest {
    listMultiPartUploadPartsRequest.maxParts = maxParts
    listMultiPartUploadPartsRequest.queryParams.Set("max_parts", strconv.Itoa(maxParts))
    return listMultiPartUploadPartsRequest
}

func (listMultiPartUploadPartsRequest *ListMultiPartUploadPartsRequest) WithPartNumberMarker(partNumberMarker *int) *ListMultiPartUploadPartsRequest {
    listMultiPartUploadPartsRequest.partNumberMarker = partNumberMarker
    if partNumberMarker != nil {
        listMultiPartUploadPartsRequest.queryParams.Set("part_number_marker", strconv.Itoa(*partNumberMarker))
    } else {
        listMultiPartUploadPartsRequest.queryParams.Set("part_number_marker", "")
    }
    return listMultiPartUploadPartsRequest
}


func (ListMultiPartUploadPartsRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (listMultiPartUploadPartsRequest *ListMultiPartUploadPartsRequest) Path() string {
    return "/" + listMultiPartUploadPartsRequest.bucketName + "/" + listMultiPartUploadPartsRequest.objectName
}

func (listMultiPartUploadPartsRequest *ListMultiPartUploadPartsRequest) QueryParams() *url.Values {
    return listMultiPartUploadPartsRequest.queryParams
}

func (ListMultiPartUploadPartsRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ListMultiPartUploadPartsRequest) Header() *http.Header {
    return &http.Header{}
}

func (ListMultiPartUploadPartsRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
