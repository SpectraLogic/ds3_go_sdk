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

type DeletePermanentlyLostTapeSpectraS3Request struct {
    tapeId string
    queryParams *url.Values
}

func NewDeletePermanentlyLostTapeSpectraS3Request(tapeId string) *DeletePermanentlyLostTapeSpectraS3Request {
    queryParams := &url.Values{}

    return &DeletePermanentlyLostTapeSpectraS3Request{
        tapeId: tapeId,
        queryParams: queryParams,
    }
}




func (DeletePermanentlyLostTapeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deletePermanentlyLostTapeSpectraS3Request *DeletePermanentlyLostTapeSpectraS3Request) Path() string {
    return "/_rest_/tape/" + deletePermanentlyLostTapeSpectraS3Request.tapeId
}

func (deletePermanentlyLostTapeSpectraS3Request *DeletePermanentlyLostTapeSpectraS3Request) QueryParams() *url.Values {
    return deletePermanentlyLostTapeSpectraS3Request.queryParams
}

func (DeletePermanentlyLostTapeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeletePermanentlyLostTapeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeletePermanentlyLostTapeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
