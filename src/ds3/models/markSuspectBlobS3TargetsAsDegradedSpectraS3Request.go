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

type MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request struct {
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

//TODO special case request payload and add unit test
func NewMarkSuspectBlobS3TargetsAsDegradedSpectraS3Request(ids []string) *MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request {
    queryParams := &url.Values{}

    return &MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request{
        content: buildIdListPayload(ids),
        queryParams: queryParams,
    }
}



func (markSuspectBlobS3TargetsAsDegradedSpectraS3Request *MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request) WithForce() *MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request {
    markSuspectBlobS3TargetsAsDegradedSpectraS3Request.queryParams.Set("force", "")
    return markSuspectBlobS3TargetsAsDegradedSpectraS3Request
}

func (MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (markSuspectBlobS3TargetsAsDegradedSpectraS3Request *MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request) Path() string {
    return "/_rest_/suspect_blob_s3_target"
}

func (markSuspectBlobS3TargetsAsDegradedSpectraS3Request *MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request) QueryParams() *url.Values {
    return markSuspectBlobS3TargetsAsDegradedSpectraS3Request.queryParams
}

func (MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (markSuspectBlobS3TargetsAsDegradedSpectraS3Request *MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return markSuspectBlobS3TargetsAsDegradedSpectraS3Request.content
}
