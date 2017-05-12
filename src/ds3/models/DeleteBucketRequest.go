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

type DeleteBucketRequest struct {
    bucketName string
    queryParams *url.Values
}

func NewDeleteBucketRequest(bucketName string) *DeleteBucketRequest {
    queryParams := &url.Values{}

    return &DeleteBucketRequest{
        bucketName: bucketName,
        queryParams: queryParams,
    }
}




func (DeleteBucketRequest) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteBucketRequest *DeleteBucketRequest) Path() string {
    return "/" + deleteBucketRequest.bucketName
}

func (deleteBucketRequest *DeleteBucketRequest) QueryParams() *url.Values {
    return deleteBucketRequest.queryParams
}

func (DeleteBucketRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteBucketRequest) Header() *http.Header {
    return &http.Header{}
}

func (DeleteBucketRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
