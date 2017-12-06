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

type ModifyBucketSpectraS3Request struct {
    BucketName string
    DataPolicyId *string
    UserId *string
}

func NewModifyBucketSpectraS3Request(bucketName string) *ModifyBucketSpectraS3Request {
    return &ModifyBucketSpectraS3Request{
        BucketName: bucketName,
    }
}

func (modifyBucketSpectraS3Request *ModifyBucketSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ModifyBucketSpectraS3Request {
    modifyBucketSpectraS3Request.DataPolicyId = &dataPolicyId
    return modifyBucketSpectraS3Request
}

func (modifyBucketSpectraS3Request *ModifyBucketSpectraS3Request) WithUserId(userId string) *ModifyBucketSpectraS3Request {
    modifyBucketSpectraS3Request.UserId = &userId
    return modifyBucketSpectraS3Request
}

