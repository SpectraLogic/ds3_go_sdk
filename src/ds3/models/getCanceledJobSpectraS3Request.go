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

type GetCanceledJobSpectraS3Request struct {
    canceledJob string
    queryParams *url.Values
}

func NewGetCanceledJobSpectraS3Request(canceledJob string) *GetCanceledJobSpectraS3Request {
    queryParams := &url.Values{}

    return &GetCanceledJobSpectraS3Request{
        canceledJob: canceledJob,
        queryParams: queryParams,
    }
}




func (GetCanceledJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getCanceledJobSpectraS3Request *GetCanceledJobSpectraS3Request) Path() string {
    return "/_rest_/canceled_job/" + getCanceledJobSpectraS3Request.canceledJob
}

func (getCanceledJobSpectraS3Request *GetCanceledJobSpectraS3Request) QueryParams() *url.Values {
    return getCanceledJobSpectraS3Request.queryParams
}

func (GetCanceledJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetCanceledJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetCanceledJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
