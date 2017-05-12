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

type VerifyPoolSpectraS3Request struct {
    pool string
    priority Priority
    queryParams *url.Values
}

func NewVerifyPoolSpectraS3Request(pool string) *VerifyPoolSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "verify")

    return &VerifyPoolSpectraS3Request{
        pool: pool,
        queryParams: queryParams,
    }
}

func (verifyPoolSpectraS3Request *VerifyPoolSpectraS3Request) WithPriority(priority Priority) *VerifyPoolSpectraS3Request {
    verifyPoolSpectraS3Request.priority = priority
    verifyPoolSpectraS3Request.queryParams.Set("priority", priority.String())
    return verifyPoolSpectraS3Request
}



func (VerifyPoolSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (verifyPoolSpectraS3Request *VerifyPoolSpectraS3Request) Path() string {
    return "/_rest_/pool/" + verifyPoolSpectraS3Request.pool
}

func (verifyPoolSpectraS3Request *VerifyPoolSpectraS3Request) QueryParams() *url.Values {
    return verifyPoolSpectraS3Request.queryParams
}

func (VerifyPoolSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifyPoolSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (VerifyPoolSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
