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

type ClearSuspectBlobS3TargetsSpectraS3Request struct {
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

//TODO special case in autogen and add unit test
func NewClearSuspectBlobS3TargetsSpectraS3Request(ids []string) *ClearSuspectBlobS3TargetsSpectraS3Request {
    queryParams := &url.Values{}

    return &ClearSuspectBlobS3TargetsSpectraS3Request{
        content: buildIdListPayload(ids),
        queryParams: queryParams,
    }
}



func (clearSuspectBlobS3TargetsSpectraS3Request *ClearSuspectBlobS3TargetsSpectraS3Request) WithForce() *ClearSuspectBlobS3TargetsSpectraS3Request {
    clearSuspectBlobS3TargetsSpectraS3Request.queryParams.Set("force", "")
    return clearSuspectBlobS3TargetsSpectraS3Request
}

func (ClearSuspectBlobS3TargetsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (clearSuspectBlobS3TargetsSpectraS3Request *ClearSuspectBlobS3TargetsSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_s3_target"
}

func (clearSuspectBlobS3TargetsSpectraS3Request *ClearSuspectBlobS3TargetsSpectraS3Request) QueryParams() *url.Values {
    return clearSuspectBlobS3TargetsSpectraS3Request.queryParams
}

func (ClearSuspectBlobS3TargetsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ClearSuspectBlobS3TargetsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (clearSuspectBlobS3TargetsSpectraS3Request *ClearSuspectBlobS3TargetsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return clearSuspectBlobS3TargetsSpectraS3Request.content
}
