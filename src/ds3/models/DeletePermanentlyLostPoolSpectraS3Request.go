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

type DeletePermanentlyLostPoolSpectraS3Request struct {
    pool string
    queryParams *url.Values
}

func NewDeletePermanentlyLostPoolSpectraS3Request(pool string) *DeletePermanentlyLostPoolSpectraS3Request {
    queryParams := &url.Values{}

    return &DeletePermanentlyLostPoolSpectraS3Request{
        pool: pool,
        queryParams: queryParams,
    }
}




func (DeletePermanentlyLostPoolSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deletePermanentlyLostPoolSpectraS3Request *DeletePermanentlyLostPoolSpectraS3Request) Path() string {
    return "/_rest_/pool/" + deletePermanentlyLostPoolSpectraS3Request.pool
}

func (deletePermanentlyLostPoolSpectraS3Request *DeletePermanentlyLostPoolSpectraS3Request) QueryParams() *url.Values {
    return deletePermanentlyLostPoolSpectraS3Request.queryParams
}

func (DeletePermanentlyLostPoolSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeletePermanentlyLostPoolSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeletePermanentlyLostPoolSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
