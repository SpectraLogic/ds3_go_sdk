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

type VerifyUserIsMemberOfGroupSpectraS3Request struct {
    group string
    userId string
    queryParams *url.Values
}

func NewVerifyUserIsMemberOfGroupSpectraS3Request(group string) *VerifyUserIsMemberOfGroupSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "verify")

    return &VerifyUserIsMemberOfGroupSpectraS3Request{
        group: group,
        queryParams: queryParams,
    }
}

func (verifyUserIsMemberOfGroupSpectraS3Request *VerifyUserIsMemberOfGroupSpectraS3Request) WithUserId(userId string) *VerifyUserIsMemberOfGroupSpectraS3Request {
    verifyUserIsMemberOfGroupSpectraS3Request.userId = userId
    verifyUserIsMemberOfGroupSpectraS3Request.queryParams.Set("user_id", userId)
    return verifyUserIsMemberOfGroupSpectraS3Request
}



func (VerifyUserIsMemberOfGroupSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (verifyUserIsMemberOfGroupSpectraS3Request *VerifyUserIsMemberOfGroupSpectraS3Request) Path() string {
    return "/_rest_/group/" + verifyUserIsMemberOfGroupSpectraS3Request.group
}

func (verifyUserIsMemberOfGroupSpectraS3Request *VerifyUserIsMemberOfGroupSpectraS3Request) QueryParams() *url.Values {
    return verifyUserIsMemberOfGroupSpectraS3Request.queryParams
}

func (VerifyUserIsMemberOfGroupSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifyUserIsMemberOfGroupSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (VerifyUserIsMemberOfGroupSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
