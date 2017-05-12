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

type PutS3TargetBucketNameSpectraS3Request struct {
    bucketId string
    name *string
    targetId string
    queryParams *url.Values
}

func NewPutS3TargetBucketNameSpectraS3Request(bucketId string, name *string, targetId string) *PutS3TargetBucketNameSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("bucket_id", bucketId)
    queryParams.Set("name", *name)
    queryParams.Set("target_id", targetId)

    return &PutS3TargetBucketNameSpectraS3Request{
        bucketId: bucketId,
        name: name,
        targetId: targetId,
        queryParams: queryParams,
    }
}




func (PutS3TargetBucketNameSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putS3TargetBucketNameSpectraS3Request *PutS3TargetBucketNameSpectraS3Request) Path() string {
    return "/_rest_/s3_target_bucket_name"
}

func (putS3TargetBucketNameSpectraS3Request *PutS3TargetBucketNameSpectraS3Request) QueryParams() *url.Values {
    return putS3TargetBucketNameSpectraS3Request.queryParams
}

func (PutS3TargetBucketNameSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutS3TargetBucketNameSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutS3TargetBucketNameSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
