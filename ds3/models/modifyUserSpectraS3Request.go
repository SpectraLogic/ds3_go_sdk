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

type ModifyUserSpectraS3Request struct {
    DefaultDataPolicyId *string
    Name *string
    SecretKey *string
    UserId string
}

func NewModifyUserSpectraS3Request(userId string) *ModifyUserSpectraS3Request {
    return &ModifyUserSpectraS3Request{
        UserId: userId,
    }
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) WithDefaultDataPolicyId(defaultDataPolicyId string) *ModifyUserSpectraS3Request {
    modifyUserSpectraS3Request.DefaultDataPolicyId = &defaultDataPolicyId
    return modifyUserSpectraS3Request
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) WithName(name string) *ModifyUserSpectraS3Request {
    modifyUserSpectraS3Request.Name = &name
    return modifyUserSpectraS3Request
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) WithSecretKey(secretKey string) *ModifyUserSpectraS3Request {
    modifyUserSpectraS3Request.SecretKey = &secretKey
    return modifyUserSpectraS3Request
}

