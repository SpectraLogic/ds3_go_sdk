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

type VerifyS3TargetSpectraS3Request struct {
    s3Target string
    queryParams *url.Values
}

func NewVerifyS3TargetSpectraS3Request(s3Target string) *VerifyS3TargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "verify")

    return &VerifyS3TargetSpectraS3Request{
        s3Target: s3Target,
        queryParams: queryParams,
    }
}



func (verifyS3TargetSpectraS3Request *VerifyS3TargetSpectraS3Request) WithFullDetails() *VerifyS3TargetSpectraS3Request {
    verifyS3TargetSpectraS3Request.queryParams.Set("full_details", "")
    return verifyS3TargetSpectraS3Request
}

func (VerifyS3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (verifyS3TargetSpectraS3Request *VerifyS3TargetSpectraS3Request) Path() string {
    return "/_rest_/s3_target/" + verifyS3TargetSpectraS3Request.s3Target
}

func (verifyS3TargetSpectraS3Request *VerifyS3TargetSpectraS3Request) QueryParams() *url.Values {
    return verifyS3TargetSpectraS3Request.queryParams
}

func (VerifyS3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifyS3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (VerifyS3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
