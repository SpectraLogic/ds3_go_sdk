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

type ModifyBucketSpectraS3Request struct {
    bucketName string
    dataPolicyId string
    userId string
    queryParams *url.Values
}

func NewModifyBucketSpectraS3Request(bucketName string) *ModifyBucketSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyBucketSpectraS3Request{
        bucketName: bucketName,
        queryParams: queryParams,
    }
}

func (modifyBucketSpectraS3Request *ModifyBucketSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ModifyBucketSpectraS3Request {
    modifyBucketSpectraS3Request.dataPolicyId = dataPolicyId
    modifyBucketSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return modifyBucketSpectraS3Request
}
func (modifyBucketSpectraS3Request *ModifyBucketSpectraS3Request) WithUserId(userId string) *ModifyBucketSpectraS3Request {
    modifyBucketSpectraS3Request.userId = userId
    modifyBucketSpectraS3Request.queryParams.Set("user_id", userId)
    return modifyBucketSpectraS3Request
}



func (ModifyBucketSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyBucketSpectraS3Request *ModifyBucketSpectraS3Request) Path() string {
    return "/_rest_/bucket/" + modifyBucketSpectraS3Request.bucketName
}

func (modifyBucketSpectraS3Request *ModifyBucketSpectraS3Request) QueryParams() *url.Values {
    return modifyBucketSpectraS3Request.queryParams
}

func (ModifyBucketSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyBucketSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyBucketSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
