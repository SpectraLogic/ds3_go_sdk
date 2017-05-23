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

type PutBucketSpectraS3Request struct {
    dataPolicyId string
    id string
    name *string
    userId string
    queryParams *url.Values
}

func NewPutBucketSpectraS3Request(name *string) *PutBucketSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("name", *name)

    return &PutBucketSpectraS3Request{
        name: name,
        queryParams: queryParams,
    }
}

func (putBucketSpectraS3Request *PutBucketSpectraS3Request) WithDataPolicyId(dataPolicyId string) *PutBucketSpectraS3Request {
    putBucketSpectraS3Request.dataPolicyId = dataPolicyId
    putBucketSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return putBucketSpectraS3Request
}
func (putBucketSpectraS3Request *PutBucketSpectraS3Request) WithId(id string) *PutBucketSpectraS3Request {
    putBucketSpectraS3Request.id = id
    putBucketSpectraS3Request.queryParams.Set("id", id)
    return putBucketSpectraS3Request
}
func (putBucketSpectraS3Request *PutBucketSpectraS3Request) WithUserId(userId string) *PutBucketSpectraS3Request {
    putBucketSpectraS3Request.userId = userId
    putBucketSpectraS3Request.queryParams.Set("user_id", userId)
    return putBucketSpectraS3Request
}



func (PutBucketSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putBucketSpectraS3Request *PutBucketSpectraS3Request) Path() string {
    return "/_rest_/bucket"
}

func (putBucketSpectraS3Request *PutBucketSpectraS3Request) QueryParams() *url.Values {
    return putBucketSpectraS3Request.queryParams
}

func (PutBucketSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutBucketSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutBucketSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
