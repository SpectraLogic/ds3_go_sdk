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

type HeadObjectRequest struct {
    bucketName string
    objectName string
    queryParams *url.Values
}

func NewHeadObjectRequest(bucketName string, objectName string) *HeadObjectRequest {
    queryParams := &url.Values{}

    return &HeadObjectRequest{
        bucketName: bucketName,
        objectName: objectName,
        queryParams: queryParams,
    }
}




func (HeadObjectRequest) Verb() networking.HttpVerb {
    return networking.HEAD
}

func (headObjectRequest *HeadObjectRequest) Path() string {
    return "/" + headObjectRequest.bucketName + "/" + headObjectRequest.objectName
}

func (headObjectRequest *HeadObjectRequest) QueryParams() *url.Values {
    return headObjectRequest.queryParams
}

func (HeadObjectRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (HeadObjectRequest) Header() *http.Header {
    return &http.Header{}
}

func (HeadObjectRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
