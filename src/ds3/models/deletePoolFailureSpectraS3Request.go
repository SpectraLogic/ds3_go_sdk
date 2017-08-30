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

type DeletePoolFailureSpectraS3Request struct {
    poolFailure string
    queryParams *url.Values
}

func NewDeletePoolFailureSpectraS3Request(poolFailure string) *DeletePoolFailureSpectraS3Request {
    queryParams := &url.Values{}

    return &DeletePoolFailureSpectraS3Request{
        poolFailure: poolFailure,
        queryParams: queryParams,
    }
}




func (DeletePoolFailureSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deletePoolFailureSpectraS3Request *DeletePoolFailureSpectraS3Request) Path() string {
    return "/_rest_/pool_failure/" + deletePoolFailureSpectraS3Request.poolFailure
}

func (deletePoolFailureSpectraS3Request *DeletePoolFailureSpectraS3Request) QueryParams() *url.Values {
    return deletePoolFailureSpectraS3Request.queryParams
}

func (DeletePoolFailureSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeletePoolFailureSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeletePoolFailureSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}