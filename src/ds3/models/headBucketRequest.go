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

type HeadBucketRequest struct {
    bucketName string
    queryParams *url.Values
}

func NewHeadBucketRequest(bucketName string) *HeadBucketRequest {
    queryParams := &url.Values{}

    return &HeadBucketRequest{
        bucketName: bucketName,
        queryParams: queryParams,
    }
}




func (HeadBucketRequest) Verb() networking.HttpVerb {
    return networking.HEAD
}

func (headBucketRequest *HeadBucketRequest) Path() string {
    return "/" + headBucketRequest.bucketName
}

func (headBucketRequest *HeadBucketRequest) QueryParams() *url.Values {
    return headBucketRequest.queryParams
}

func (HeadBucketRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (HeadBucketRequest) Header() *http.Header {
    return &http.Header{}
}

func (HeadBucketRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
