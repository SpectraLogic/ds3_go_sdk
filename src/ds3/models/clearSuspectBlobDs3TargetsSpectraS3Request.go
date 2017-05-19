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

type ClearSuspectBlobDs3TargetsSpectraS3Request struct {
    queryParams *url.Values
}

func NewClearSuspectBlobDs3TargetsSpectraS3Request() *ClearSuspectBlobDs3TargetsSpectraS3Request {
    queryParams := &url.Values{}

    return &ClearSuspectBlobDs3TargetsSpectraS3Request{
        queryParams: queryParams,
    }
}



func (clearSuspectBlobDs3TargetsSpectraS3Request *ClearSuspectBlobDs3TargetsSpectraS3Request) WithForce() *ClearSuspectBlobDs3TargetsSpectraS3Request {
    clearSuspectBlobDs3TargetsSpectraS3Request.queryParams.Set("force", "")
    return clearSuspectBlobDs3TargetsSpectraS3Request
}

func (ClearSuspectBlobDs3TargetsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (clearSuspectBlobDs3TargetsSpectraS3Request *ClearSuspectBlobDs3TargetsSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_ds3_target"
}

func (clearSuspectBlobDs3TargetsSpectraS3Request *ClearSuspectBlobDs3TargetsSpectraS3Request) QueryParams() *url.Values {
    return clearSuspectBlobDs3TargetsSpectraS3Request.queryParams
}

func (ClearSuspectBlobDs3TargetsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ClearSuspectBlobDs3TargetsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ClearSuspectBlobDs3TargetsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
