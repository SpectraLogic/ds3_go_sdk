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

type ListMultiPartUploadPartsRequest struct {
    BucketName string
    ObjectName string
    MaxParts *int
    PartNumberMarker *int
    UploadId string
}

func NewListMultiPartUploadPartsRequest(bucketName string, objectName string, uploadId string) *ListMultiPartUploadPartsRequest {
    return &ListMultiPartUploadPartsRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        UploadId: uploadId,
    }
}

func (listMultiPartUploadPartsRequest *ListMultiPartUploadPartsRequest) WithMaxParts(maxParts int) *ListMultiPartUploadPartsRequest {
    listMultiPartUploadPartsRequest.MaxParts = &maxParts
    return listMultiPartUploadPartsRequest
}

func (listMultiPartUploadPartsRequest *ListMultiPartUploadPartsRequest) WithPartNumberMarker(partNumberMarker int) *ListMultiPartUploadPartsRequest {
    listMultiPartUploadPartsRequest.PartNumberMarker = &partNumberMarker
    return listMultiPartUploadPartsRequest
}

