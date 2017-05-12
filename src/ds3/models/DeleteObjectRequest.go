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

type DeleteObjectRequest struct {
    bucketName string
    objectName string
    queryParams *url.Values
}

func NewDeleteObjectRequest(bucketName string, objectName string) *DeleteObjectRequest {
    queryParams := &url.Values{}

    return &DeleteObjectRequest{
        bucketName: bucketName,
        objectName: objectName,
        queryParams: queryParams,
    }
}



func (deleteObjectRequest *DeleteObjectRequest) WithRollBack() *DeleteObjectRequest {
    deleteObjectRequest.queryParams.Set("roll_back", "")
    return deleteObjectRequest
}

func (DeleteObjectRequest) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteObjectRequest *DeleteObjectRequest) Path() string {
    return "/" + deleteObjectRequest.bucketName + "/" + deleteObjectRequest.objectName
}

func (deleteObjectRequest *DeleteObjectRequest) QueryParams() *url.Values {
    return deleteObjectRequest.queryParams
}

func (DeleteObjectRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteObjectRequest) Header() *http.Header {
    return &http.Header{}
}

func (DeleteObjectRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
