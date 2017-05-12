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

type GetBlobsOnDs3TargetSpectraS3Request struct {
    content networking.ReaderWithSizeDecorator
    ds3Target string
    queryParams *url.Values
}

func NewGetBlobsOnDs3TargetSpectraS3Request(ds3Target string, objects []Ds3Object) *GetBlobsOnDs3TargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "get_physical_placement")

    return &GetBlobsOnDs3TargetSpectraS3Request{
        ds3Target: ds3Target,
        content: buildDs3ObjectListStream(objects),
        queryParams: queryParams,
    }
}




func (GetBlobsOnDs3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getBlobsOnDs3TargetSpectraS3Request *GetBlobsOnDs3TargetSpectraS3Request) Path() string {
    return "/_rest_/ds3_target/" + getBlobsOnDs3TargetSpectraS3Request.ds3Target
}

func (getBlobsOnDs3TargetSpectraS3Request *GetBlobsOnDs3TargetSpectraS3Request) QueryParams() *url.Values {
    return getBlobsOnDs3TargetSpectraS3Request.queryParams
}

func (GetBlobsOnDs3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetBlobsOnDs3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (getBlobsOnDs3TargetSpectraS3Request *GetBlobsOnDs3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return getBlobsOnDs3TargetSpectraS3Request.content
}
