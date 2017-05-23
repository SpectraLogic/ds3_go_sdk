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

type VerifyDs3TargetSpectraS3Request struct {
    ds3Target string
    queryParams *url.Values
}

func NewVerifyDs3TargetSpectraS3Request(ds3Target string) *VerifyDs3TargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "verify")

    return &VerifyDs3TargetSpectraS3Request{
        ds3Target: ds3Target,
        queryParams: queryParams,
    }
}



func (verifyDs3TargetSpectraS3Request *VerifyDs3TargetSpectraS3Request) WithFullDetails() *VerifyDs3TargetSpectraS3Request {
    verifyDs3TargetSpectraS3Request.queryParams.Set("full_details", "")
    return verifyDs3TargetSpectraS3Request
}

func (VerifyDs3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (verifyDs3TargetSpectraS3Request *VerifyDs3TargetSpectraS3Request) Path() string {
    return "/_rest_/ds3_target/" + verifyDs3TargetSpectraS3Request.ds3Target
}

func (verifyDs3TargetSpectraS3Request *VerifyDs3TargetSpectraS3Request) QueryParams() *url.Values {
    return verifyDs3TargetSpectraS3Request.queryParams
}

func (VerifyDs3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifyDs3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (VerifyDs3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
