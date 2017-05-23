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

type CompleteMultiPartUploadRequest struct {
    bucketName string
    objectName string
    content networking.ReaderWithSizeDecorator
    uploadId string
    queryParams *url.Values
}

func NewCompleteMultiPartUploadRequest(bucketName string, objectName string, parts []Part, uploadId string) *CompleteMultiPartUploadRequest {
    queryParams := &url.Values{}
    queryParams.Set("upload_id", uploadId)

    return &CompleteMultiPartUploadRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
        content: buildPartsListStream(parts),
        queryParams: queryParams,
    }
}




func (CompleteMultiPartUploadRequest) Verb() networking.HttpVerb {
    return networking.POST
}

func (completeMultiPartUploadRequest *CompleteMultiPartUploadRequest) Path() string {
    return "/" + completeMultiPartUploadRequest.bucketName + "/" + completeMultiPartUploadRequest.objectName
}

func (completeMultiPartUploadRequest *CompleteMultiPartUploadRequest) QueryParams() *url.Values {
    return completeMultiPartUploadRequest.queryParams
}

func (CompleteMultiPartUploadRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (CompleteMultiPartUploadRequest) Header() *http.Header {
    return &http.Header{}
}

func (completeMultiPartUploadRequest *CompleteMultiPartUploadRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return completeMultiPartUploadRequest.content
}
