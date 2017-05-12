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

type VerifyAllPoolsSpectraS3Request struct {
    priority Priority
    queryParams *url.Values
}

func NewVerifyAllPoolsSpectraS3Request() *VerifyAllPoolsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "verify")

    return &VerifyAllPoolsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (verifyAllPoolsSpectraS3Request *VerifyAllPoolsSpectraS3Request) WithPriority(priority Priority) *VerifyAllPoolsSpectraS3Request {
    verifyAllPoolsSpectraS3Request.priority = priority
    verifyAllPoolsSpectraS3Request.queryParams.Set("priority", priority.String())
    return verifyAllPoolsSpectraS3Request
}



func (VerifyAllPoolsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (verifyAllPoolsSpectraS3Request *VerifyAllPoolsSpectraS3Request) Path() string {
    return "/_rest_/pool"
}

func (verifyAllPoolsSpectraS3Request *VerifyAllPoolsSpectraS3Request) QueryParams() *url.Values {
    return verifyAllPoolsSpectraS3Request.queryParams
}

func (VerifyAllPoolsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifyAllPoolsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (VerifyAllPoolsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
