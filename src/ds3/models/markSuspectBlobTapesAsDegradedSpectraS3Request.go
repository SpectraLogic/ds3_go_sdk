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

type MarkSuspectBlobTapesAsDegradedSpectraS3Request struct {
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

func NewMarkSuspectBlobTapesAsDegradedSpectraS3Request(ids []string) *MarkSuspectBlobTapesAsDegradedSpectraS3Request {
    queryParams := &url.Values{}

    return &MarkSuspectBlobTapesAsDegradedSpectraS3Request{
        content: buildIdListPayload(ids),
        queryParams: queryParams,
    }
}



func (markSuspectBlobTapesAsDegradedSpectraS3Request *MarkSuspectBlobTapesAsDegradedSpectraS3Request) WithForce() *MarkSuspectBlobTapesAsDegradedSpectraS3Request {
    markSuspectBlobTapesAsDegradedSpectraS3Request.queryParams.Set("force", "")
    return markSuspectBlobTapesAsDegradedSpectraS3Request
}

func (MarkSuspectBlobTapesAsDegradedSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (markSuspectBlobTapesAsDegradedSpectraS3Request *MarkSuspectBlobTapesAsDegradedSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_tape"
}

func (markSuspectBlobTapesAsDegradedSpectraS3Request *MarkSuspectBlobTapesAsDegradedSpectraS3Request) QueryParams() *url.Values {
    return markSuspectBlobTapesAsDegradedSpectraS3Request.queryParams
}

func (MarkSuspectBlobTapesAsDegradedSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (MarkSuspectBlobTapesAsDegradedSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (markSuspectBlobTapesAsDegradedSpectraS3Request *MarkSuspectBlobTapesAsDegradedSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return markSuspectBlobTapesAsDegradedSpectraS3Request.content
}
