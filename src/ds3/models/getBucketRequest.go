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

type GetBucketRequest struct {
    bucketName string
    delimiter *string
    marker *string
    maxKeys int
    prefix *string
    queryParams *url.Values
}

func NewGetBucketRequest(bucketName string) *GetBucketRequest {
    queryParams := &url.Values{}

    return &GetBucketRequest{
        bucketName: bucketName,
        queryParams: queryParams,
    }
}

func (getBucketRequest *GetBucketRequest) WithMaxKeys(maxKeys int) *GetBucketRequest {
    getBucketRequest.maxKeys = maxKeys
    getBucketRequest.queryParams.Set("max_keys", strconv.Itoa(maxKeys))
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithDelimiter(delimiter *string) *GetBucketRequest {
    getBucketRequest.delimiter = delimiter
    if delimiter != nil {
        getBucketRequest.queryParams.Set("delimiter", *delimiter)
    } else {
        getBucketRequest.queryParams.Set("delimiter", "")
    }
    return getBucketRequest
}
func (getBucketRequest *GetBucketRequest) WithMarker(marker *string) *GetBucketRequest {
    getBucketRequest.marker = marker
    if marker != nil {
        getBucketRequest.queryParams.Set("marker", *marker)
    } else {
        getBucketRequest.queryParams.Set("marker", "")
    }
    return getBucketRequest
}
func (getBucketRequest *GetBucketRequest) WithPrefix(prefix *string) *GetBucketRequest {
    getBucketRequest.prefix = prefix
    if prefix != nil {
        getBucketRequest.queryParams.Set("prefix", *prefix)
    } else {
        getBucketRequest.queryParams.Set("prefix", "")
    }
    return getBucketRequest
}


func (GetBucketRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (getBucketRequest *GetBucketRequest) Path() string {
    return "/" + getBucketRequest.bucketName
}

func (getBucketRequest *GetBucketRequest) QueryParams() *url.Values {
    return getBucketRequest.queryParams
}

func (GetBucketRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetBucketRequest) Header() *http.Header {
    return &http.Header{}
}

func (GetBucketRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
