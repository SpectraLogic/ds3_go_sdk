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

type GetBlobsOnPoolSpectraS3Request struct {
    content networking.ReaderWithSizeDecorator
    pool string
    queryParams *url.Values
}

func NewGetBlobsOnPoolSpectraS3Request(objects []Ds3Object, pool string) *GetBlobsOnPoolSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "get_physical_placement")

    return &GetBlobsOnPoolSpectraS3Request{
        pool: pool,
        content: buildDs3ObjectListStream(objects),
        queryParams: queryParams,
    }
}




func (GetBlobsOnPoolSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getBlobsOnPoolSpectraS3Request *GetBlobsOnPoolSpectraS3Request) Path() string {
    return "/_rest_/pool/" + getBlobsOnPoolSpectraS3Request.pool
}

func (getBlobsOnPoolSpectraS3Request *GetBlobsOnPoolSpectraS3Request) QueryParams() *url.Values {
    return getBlobsOnPoolSpectraS3Request.queryParams
}

func (GetBlobsOnPoolSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetBlobsOnPoolSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (getBlobsOnPoolSpectraS3Request *GetBlobsOnPoolSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return getBlobsOnPoolSpectraS3Request.content
}
