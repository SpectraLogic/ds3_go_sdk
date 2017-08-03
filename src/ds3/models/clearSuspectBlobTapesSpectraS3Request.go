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

type ClearSuspectBlobTapesSpectraS3Request struct {
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

//TODO special case in autogen and add unit test
func NewClearSuspectBlobTapesSpectraS3Request(ids []string) *ClearSuspectBlobTapesSpectraS3Request {
    queryParams := &url.Values{}

    return &ClearSuspectBlobTapesSpectraS3Request{
        content: buildIdListPayload(ids),
        queryParams: queryParams,
    }
}



func (clearSuspectBlobTapesSpectraS3Request *ClearSuspectBlobTapesSpectraS3Request) WithForce() *ClearSuspectBlobTapesSpectraS3Request {
    clearSuspectBlobTapesSpectraS3Request.queryParams.Set("force", "")
    return clearSuspectBlobTapesSpectraS3Request
}

func (ClearSuspectBlobTapesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (clearSuspectBlobTapesSpectraS3Request *ClearSuspectBlobTapesSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_tape"
}

func (clearSuspectBlobTapesSpectraS3Request *ClearSuspectBlobTapesSpectraS3Request) QueryParams() *url.Values {
    return clearSuspectBlobTapesSpectraS3Request.queryParams
}

func (ClearSuspectBlobTapesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ClearSuspectBlobTapesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ClearSuspectBlobTapesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
