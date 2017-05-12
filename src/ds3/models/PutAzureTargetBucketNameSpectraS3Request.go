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

type PutAzureTargetBucketNameSpectraS3Request struct {
    bucketId string
    name *string
    targetId string
    queryParams *url.Values
}

func NewPutAzureTargetBucketNameSpectraS3Request(bucketId string, name *string, targetId string) *PutAzureTargetBucketNameSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("bucket_id", bucketId)
    queryParams.Set("name", *name)
    queryParams.Set("target_id", targetId)

    return &PutAzureTargetBucketNameSpectraS3Request{
        bucketId: bucketId,
        name: name,
        targetId: targetId,
        queryParams: queryParams,
    }
}




func (PutAzureTargetBucketNameSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putAzureTargetBucketNameSpectraS3Request *PutAzureTargetBucketNameSpectraS3Request) Path() string {
    return "/_rest_/azure_target_bucket_name"
}

func (putAzureTargetBucketNameSpectraS3Request *PutAzureTargetBucketNameSpectraS3Request) QueryParams() *url.Values {
    return putAzureTargetBucketNameSpectraS3Request.queryParams
}

func (PutAzureTargetBucketNameSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutAzureTargetBucketNameSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutAzureTargetBucketNameSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
