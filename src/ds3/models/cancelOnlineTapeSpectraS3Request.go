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

type CancelOnlineTapeSpectraS3Request struct {
    tapeId string
    queryParams *url.Values
}

func NewCancelOnlineTapeSpectraS3Request(tapeId string) *CancelOnlineTapeSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "cancel_online")

    return &CancelOnlineTapeSpectraS3Request{
        tapeId: tapeId,
        queryParams: queryParams,
    }
}




func (CancelOnlineTapeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (cancelOnlineTapeSpectraS3Request *CancelOnlineTapeSpectraS3Request) Path() string {
    return "/_rest_/tape/" + cancelOnlineTapeSpectraS3Request.tapeId
}

func (cancelOnlineTapeSpectraS3Request *CancelOnlineTapeSpectraS3Request) QueryParams() *url.Values {
    return cancelOnlineTapeSpectraS3Request.queryParams
}

func (CancelOnlineTapeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (CancelOnlineTapeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (CancelOnlineTapeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
