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

type CancelActiveJobSpectraS3Request struct {
    activeJobId string
    queryParams *url.Values
}

func NewCancelActiveJobSpectraS3Request(activeJobId string) *CancelActiveJobSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("force", "")

    return &CancelActiveJobSpectraS3Request{
        activeJobId: activeJobId,
        queryParams: queryParams,
    }
}




func (CancelActiveJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (cancelActiveJobSpectraS3Request *CancelActiveJobSpectraS3Request) Path() string {
    return "/_rest_/active_job/" + cancelActiveJobSpectraS3Request.activeJobId
}

func (cancelActiveJobSpectraS3Request *CancelActiveJobSpectraS3Request) QueryParams() *url.Values {
    return cancelActiveJobSpectraS3Request.queryParams
}

func (CancelActiveJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (CancelActiveJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (CancelActiveJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
