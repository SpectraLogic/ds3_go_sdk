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

type ClearSuspectBlobPoolsSpectraS3Request struct {
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

func NewClearSuspectBlobPoolsSpectraS3Request(ids []string) *ClearSuspectBlobPoolsSpectraS3Request {
    queryParams := &url.Values{}

    return &ClearSuspectBlobPoolsSpectraS3Request{
        content: buildIdListPayload(ids),
        queryParams: queryParams,
    }
}



func (clearSuspectBlobPoolsSpectraS3Request *ClearSuspectBlobPoolsSpectraS3Request) WithForce() *ClearSuspectBlobPoolsSpectraS3Request {
    clearSuspectBlobPoolsSpectraS3Request.queryParams.Set("force", "")
    return clearSuspectBlobPoolsSpectraS3Request
}

func (ClearSuspectBlobPoolsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (clearSuspectBlobPoolsSpectraS3Request *ClearSuspectBlobPoolsSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_pool"
}

func (clearSuspectBlobPoolsSpectraS3Request *ClearSuspectBlobPoolsSpectraS3Request) QueryParams() *url.Values {
    return clearSuspectBlobPoolsSpectraS3Request.queryParams
}

func (ClearSuspectBlobPoolsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ClearSuspectBlobPoolsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (clearSuspectBlobPoolsSpectraS3Request *ClearSuspectBlobPoolsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return clearSuspectBlobPoolsSpectraS3Request.content
}
