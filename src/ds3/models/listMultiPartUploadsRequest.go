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

type ListMultiPartUploadsRequest struct {
    bucketName string
    delimiter *string
    keyMarker *string
    maxUploads int
    prefix *string
    uploadIdMarker *string
    queryParams *url.Values
}

func NewListMultiPartUploadsRequest(bucketName string) *ListMultiPartUploadsRequest {
    queryParams := &url.Values{}
    queryParams.Set("uploads", "")

    return &ListMultiPartUploadsRequest{
        bucketName: bucketName,
        queryParams: queryParams,
    }
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithMaxUploads(maxUploads int) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.maxUploads = maxUploads
    listMultiPartUploadsRequest.queryParams.Set("max_uploads", strconv.Itoa(maxUploads))
    return listMultiPartUploadsRequest
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithDelimiter(delimiter *string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.delimiter = delimiter
    if delimiter != nil {
        listMultiPartUploadsRequest.queryParams.Set("delimiter", *delimiter)
    } else {
        listMultiPartUploadsRequest.queryParams.Set("delimiter", "")
    }
    return listMultiPartUploadsRequest
}
func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithKeyMarker(keyMarker *string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.keyMarker = keyMarker
    if keyMarker != nil {
        listMultiPartUploadsRequest.queryParams.Set("key_marker", *keyMarker)
    } else {
        listMultiPartUploadsRequest.queryParams.Set("key_marker", "")
    }
    return listMultiPartUploadsRequest
}
func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithPrefix(prefix *string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.prefix = prefix
    if prefix != nil {
        listMultiPartUploadsRequest.queryParams.Set("prefix", *prefix)
    } else {
        listMultiPartUploadsRequest.queryParams.Set("prefix", "")
    }
    return listMultiPartUploadsRequest
}
func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithUploadIdMarker(uploadIdMarker *string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.uploadIdMarker = uploadIdMarker
    if uploadIdMarker != nil {
        listMultiPartUploadsRequest.queryParams.Set("upload_id_marker", *uploadIdMarker)
    } else {
        listMultiPartUploadsRequest.queryParams.Set("upload_id_marker", "")
    }
    return listMultiPartUploadsRequest
}


func (ListMultiPartUploadsRequest) Verb() networking.HttpVerb {
    return networking.GET
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) Path() string {
    return "/" + listMultiPartUploadsRequest.bucketName
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) QueryParams() *url.Values {
    return listMultiPartUploadsRequest.queryParams
}

func (ListMultiPartUploadsRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ListMultiPartUploadsRequest) Header() *http.Header {
    return &http.Header{}
}

func (ListMultiPartUploadsRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
