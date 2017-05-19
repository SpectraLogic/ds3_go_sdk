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

type GetNodeSpectraS3Request struct {
    node string
    queryParams *url.Values
}

func NewGetNodeSpectraS3Request(node string) *GetNodeSpectraS3Request {
    queryParams := &url.Values{}

    return &GetNodeSpectraS3Request{
        node: node,
        queryParams: queryParams,
    }
}




func (GetNodeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getNodeSpectraS3Request *GetNodeSpectraS3Request) Path() string {
    return "/_rest_/node/" + getNodeSpectraS3Request.node
}

func (getNodeSpectraS3Request *GetNodeSpectraS3Request) QueryParams() *url.Values {
    return getNodeSpectraS3Request.queryParams
}

func (GetNodeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetNodeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetNodeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
