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
)

type DeleteObjectsRequest struct {
    bucketName string
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

func NewDeleteObjectsRequest(bucketName string, objectNames []string) *DeleteObjectsRequest {
    queryParams := &url.Values{}
    queryParams.Set("delete", "")

    return &DeleteObjectsRequest{
        bucketName: bucketName,
        content: buildDeleteObjectsPayload(objectNames),
        queryParams: queryParams,
    }
}



func (deleteObjectsRequest *DeleteObjectsRequest) WithRollBack() *DeleteObjectsRequest {
    deleteObjectsRequest.queryParams.Set("roll_back", "")
    return deleteObjectsRequest
}

func (DeleteObjectsRequest) Verb() networking.HttpVerb {
    return networking.POST
}

func (deleteObjectsRequest *DeleteObjectsRequest) Path() string {
    return "/" + deleteObjectsRequest.bucketName
}

func (deleteObjectsRequest *DeleteObjectsRequest) QueryParams() *url.Values {
    return deleteObjectsRequest.queryParams
}

func (DeleteObjectsRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteObjectsRequest) Header() *http.Header {
    return &http.Header{}
}

func (deleteObjectsRequest *DeleteObjectsRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return deleteObjectsRequest.content
}
