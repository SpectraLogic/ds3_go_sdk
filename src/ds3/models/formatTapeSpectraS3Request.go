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

type FormatTapeSpectraS3Request struct {
    tapeId string
    queryParams *url.Values
}

func NewFormatTapeSpectraS3Request(tapeId string) *FormatTapeSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "format")

    return &FormatTapeSpectraS3Request{
        tapeId: tapeId,
        queryParams: queryParams,
    }
}



func (formatTapeSpectraS3Request *FormatTapeSpectraS3Request) WithForce() *FormatTapeSpectraS3Request {
    formatTapeSpectraS3Request.queryParams.Set("force", "")
    return formatTapeSpectraS3Request
}

func (FormatTapeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (formatTapeSpectraS3Request *FormatTapeSpectraS3Request) Path() string {
    return "/_rest_/tape/" + formatTapeSpectraS3Request.tapeId
}

func (formatTapeSpectraS3Request *FormatTapeSpectraS3Request) QueryParams() *url.Values {
    return formatTapeSpectraS3Request.queryParams
}

func (FormatTapeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (FormatTapeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (FormatTapeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}