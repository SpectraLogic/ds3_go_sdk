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

type CompactPoolSpectraS3Request struct {
    pool string
    priority Priority
    queryParams *url.Values
}

func NewCompactPoolSpectraS3Request(pool string) *CompactPoolSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "compact")

    return &CompactPoolSpectraS3Request{
        pool: pool,
        queryParams: queryParams,
    }
}

func (compactPoolSpectraS3Request *CompactPoolSpectraS3Request) WithPriority(priority Priority) *CompactPoolSpectraS3Request {
    compactPoolSpectraS3Request.priority = priority
    compactPoolSpectraS3Request.queryParams.Set("priority", priority.String())
    return compactPoolSpectraS3Request
}



func (CompactPoolSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (compactPoolSpectraS3Request *CompactPoolSpectraS3Request) Path() string {
    return "/_rest_/pool/" + compactPoolSpectraS3Request.pool
}

func (compactPoolSpectraS3Request *CompactPoolSpectraS3Request) QueryParams() *url.Values {
    return compactPoolSpectraS3Request.queryParams
}

func (CompactPoolSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (CompactPoolSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (CompactPoolSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
