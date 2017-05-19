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

type TruncateActiveJobSpectraS3Request struct {
    activeJobId string
    queryParams *url.Values
}

func NewTruncateActiveJobSpectraS3Request(activeJobId string) *TruncateActiveJobSpectraS3Request {
    queryParams := &url.Values{}

    return &TruncateActiveJobSpectraS3Request{
        activeJobId: activeJobId,
        queryParams: queryParams,
    }
}




func (TruncateActiveJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (truncateActiveJobSpectraS3Request *TruncateActiveJobSpectraS3Request) Path() string {
    return "/_rest_/active_job/" + truncateActiveJobSpectraS3Request.activeJobId
}

func (truncateActiveJobSpectraS3Request *TruncateActiveJobSpectraS3Request) QueryParams() *url.Values {
    return truncateActiveJobSpectraS3Request.queryParams
}

func (TruncateActiveJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (TruncateActiveJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (TruncateActiveJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
