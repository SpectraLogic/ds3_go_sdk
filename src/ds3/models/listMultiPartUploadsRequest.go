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

type ListMultiPartUploadsRequest struct {
    BucketName string
    Delimiter *string
    KeyMarker *string
    MaxUploads *int
    Prefix *string
    UploadIdMarker *string
}

func NewListMultiPartUploadsRequest(bucketName string) *ListMultiPartUploadsRequest {
    return &ListMultiPartUploadsRequest{
        BucketName: bucketName,
    }
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithDelimiter(delimiter string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.Delimiter = &delimiter
    return listMultiPartUploadsRequest
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithKeyMarker(keyMarker string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.KeyMarker = &keyMarker
    return listMultiPartUploadsRequest
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithMaxUploads(maxUploads int) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.MaxUploads = &maxUploads
    return listMultiPartUploadsRequest
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithPrefix(prefix string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.Prefix = &prefix
    return listMultiPartUploadsRequest
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithUploadIdMarker(uploadIdMarker string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.UploadIdMarker = &uploadIdMarker
    return listMultiPartUploadsRequest
}

