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

type ModifyTapeDriveSpectraS3Request struct {
    quiesced Quiesced
    tapeDriveId string
    queryParams *url.Values
}

func NewModifyTapeDriveSpectraS3Request(tapeDriveId string) *ModifyTapeDriveSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyTapeDriveSpectraS3Request{
        tapeDriveId: tapeDriveId,
        queryParams: queryParams,
    }
}

func (modifyTapeDriveSpectraS3Request *ModifyTapeDriveSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyTapeDriveSpectraS3Request {
    modifyTapeDriveSpectraS3Request.quiesced = quiesced
    modifyTapeDriveSpectraS3Request.queryParams.Set("quiesced", quiesced.String())
    return modifyTapeDriveSpectraS3Request
}



func (ModifyTapeDriveSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyTapeDriveSpectraS3Request *ModifyTapeDriveSpectraS3Request) Path() string {
    return "/_rest_/tape_drive/" + modifyTapeDriveSpectraS3Request.tapeDriveId
}

func (modifyTapeDriveSpectraS3Request *ModifyTapeDriveSpectraS3Request) QueryParams() *url.Values {
    return modifyTapeDriveSpectraS3Request.queryParams
}

func (ModifyTapeDriveSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyTapeDriveSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyTapeDriveSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
