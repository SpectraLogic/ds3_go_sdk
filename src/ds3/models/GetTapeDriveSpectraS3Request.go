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

type GetTapeDriveSpectraS3Request struct {
    tapeDriveId string
    queryParams *url.Values
}

func NewGetTapeDriveSpectraS3Request(tapeDriveId string) *GetTapeDriveSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapeDriveSpectraS3Request{
        tapeDriveId: tapeDriveId,
        queryParams: queryParams,
    }
}




func (GetTapeDriveSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapeDriveSpectraS3Request *GetTapeDriveSpectraS3Request) Path() string {
    return "/_rest_/tape_drive/" + getTapeDriveSpectraS3Request.tapeDriveId
}

func (getTapeDriveSpectraS3Request *GetTapeDriveSpectraS3Request) QueryParams() *url.Values {
    return getTapeDriveSpectraS3Request.queryParams
}

func (GetTapeDriveSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapeDriveSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapeDriveSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
