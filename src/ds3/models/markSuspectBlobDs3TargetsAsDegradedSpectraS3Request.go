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

type MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request struct {
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

//TODO special case request payload and add unit test
func NewMarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request(ids []string) *MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request {
    queryParams := &url.Values{}

    return &MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request{
        content: buildIdListPayload(ids),
        queryParams: queryParams,
    }
}



func (markSuspectBlobDs3TargetsAsDegradedSpectraS3Request *MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request) WithForce() *MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request {
    markSuspectBlobDs3TargetsAsDegradedSpectraS3Request.queryParams.Set("force", "")
    return markSuspectBlobDs3TargetsAsDegradedSpectraS3Request
}

func (MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (markSuspectBlobDs3TargetsAsDegradedSpectraS3Request *MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_ds3_target"
}

func (markSuspectBlobDs3TargetsAsDegradedSpectraS3Request *MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request) QueryParams() *url.Values {
    return markSuspectBlobDs3TargetsAsDegradedSpectraS3Request.queryParams
}

func (MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (markSuspectBlobDs3TargetsAsDegradedSpectraS3Request *MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return markSuspectBlobDs3TargetsAsDegradedSpectraS3Request.content
}
