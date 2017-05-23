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
    "strconv"
)

type VerifyBulkJobSpectraS3Request struct {
    bucketName string
    aggregating bool
    content networking.ReaderWithSizeDecorator
    name *string
    priority Priority
    queryParams *url.Values
}

func NewVerifyBulkJobSpectraS3Request(bucketName string, objects []Ds3Object) *VerifyBulkJobSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "start_bulk_verify")

    return &VerifyBulkJobSpectraS3Request{
        bucketName: bucketName,
        content: buildDs3ObjectListStream(objects),
        queryParams: queryParams,
    }
}

func (verifyBulkJobSpectraS3Request *VerifyBulkJobSpectraS3Request) WithAggregating(aggregating bool) *VerifyBulkJobSpectraS3Request {
    verifyBulkJobSpectraS3Request.aggregating = aggregating
    verifyBulkJobSpectraS3Request.queryParams.Set("aggregating", strconv.FormatBool(aggregating))
    return verifyBulkJobSpectraS3Request
}
func (verifyBulkJobSpectraS3Request *VerifyBulkJobSpectraS3Request) WithPriority(priority Priority) *VerifyBulkJobSpectraS3Request {
    verifyBulkJobSpectraS3Request.priority = priority
    verifyBulkJobSpectraS3Request.queryParams.Set("priority", priority.String())
    return verifyBulkJobSpectraS3Request
}

func (verifyBulkJobSpectraS3Request *VerifyBulkJobSpectraS3Request) WithName(name *string) *VerifyBulkJobSpectraS3Request {
    verifyBulkJobSpectraS3Request.name = name
    if name != nil {
        verifyBulkJobSpectraS3Request.queryParams.Set("name", *name)
    } else {
        verifyBulkJobSpectraS3Request.queryParams.Set("name", "")
    }
    return verifyBulkJobSpectraS3Request
}


func (VerifyBulkJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (verifyBulkJobSpectraS3Request *VerifyBulkJobSpectraS3Request) Path() string {
    return "/_rest_/bucket/" + verifyBulkJobSpectraS3Request.bucketName
}

func (verifyBulkJobSpectraS3Request *VerifyBulkJobSpectraS3Request) QueryParams() *url.Values {
    return verifyBulkJobSpectraS3Request.queryParams
}

func (VerifyBulkJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifyBulkJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (verifyBulkJobSpectraS3Request *VerifyBulkJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return verifyBulkJobSpectraS3Request.content
}
