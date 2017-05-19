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

type DeleteDs3TargetSpectraS3Request struct {
    ds3Target string
    queryParams *url.Values
}

func NewDeleteDs3TargetSpectraS3Request(ds3Target string) *DeleteDs3TargetSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteDs3TargetSpectraS3Request{
        ds3Target: ds3Target,
        queryParams: queryParams,
    }
}




func (DeleteDs3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteDs3TargetSpectraS3Request *DeleteDs3TargetSpectraS3Request) Path() string {
    return "/_rest_/ds3_target/" + deleteDs3TargetSpectraS3Request.ds3Target
}

func (deleteDs3TargetSpectraS3Request *DeleteDs3TargetSpectraS3Request) QueryParams() *url.Values {
    return deleteDs3TargetSpectraS3Request.queryParams
}

func (DeleteDs3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteDs3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteDs3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
