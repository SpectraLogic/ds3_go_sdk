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

type PutBulkJobSpectraS3Request struct {
    bucketName string
    aggregating bool
    content networking.ReaderWithSizeDecorator
    implicitJobIdResolution bool
    maxUploadSize int64
    minimizeSpanningAcrossMedia bool
    name *string
    priority Priority
    verifyAfterWrite bool
    queryParams *url.Values
}

func NewPutBulkJobSpectraS3Request(bucketName string, objects []Ds3Object) *PutBulkJobSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "start_bulk_put")

    return &PutBulkJobSpectraS3Request{
        bucketName: bucketName,
        content: buildDs3ObjectListStream(objects),
        queryParams: queryParams,
    }
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithAggregating(aggregating bool) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.aggregating = aggregating
    putBulkJobSpectraS3Request.queryParams.Set("aggregating", strconv.FormatBool(aggregating))
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithImplicitJobIdResolution(implicitJobIdResolution bool) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.implicitJobIdResolution = implicitJobIdResolution
    putBulkJobSpectraS3Request.queryParams.Set("implicit_job_id_resolution", strconv.FormatBool(implicitJobIdResolution))
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithMaxUploadSize(maxUploadSize int64) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.maxUploadSize = maxUploadSize
    putBulkJobSpectraS3Request.queryParams.Set("max_upload_size", strconv.FormatInt(maxUploadSize, 10))
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithMinimizeSpanningAcrossMedia(minimizeSpanningAcrossMedia bool) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.minimizeSpanningAcrossMedia = minimizeSpanningAcrossMedia
    putBulkJobSpectraS3Request.queryParams.Set("minimize_spanning_across_media", strconv.FormatBool(minimizeSpanningAcrossMedia))
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithPriority(priority Priority) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.priority = priority
    putBulkJobSpectraS3Request.queryParams.Set("priority", priority.String())
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithVerifyAfterWrite(verifyAfterWrite bool) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.verifyAfterWrite = verifyAfterWrite
    putBulkJobSpectraS3Request.queryParams.Set("verify_after_write", strconv.FormatBool(verifyAfterWrite))
    return putBulkJobSpectraS3Request
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithName(name *string) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.name = name
    if name != nil {
        putBulkJobSpectraS3Request.queryParams.Set("name", *name)
    } else {
        putBulkJobSpectraS3Request.queryParams.Set("name", "")
    }
    return putBulkJobSpectraS3Request
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithForce() *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.queryParams.Set("force", "")
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithIgnoreNamingConflicts() *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.queryParams.Set("ignore_naming_conflicts", "")
    return putBulkJobSpectraS3Request
}

func (PutBulkJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) Path() string {
    return "/_rest_/bucket/" + putBulkJobSpectraS3Request.bucketName
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) QueryParams() *url.Values {
    return putBulkJobSpectraS3Request.queryParams
}

func (PutBulkJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutBulkJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return putBulkJobSpectraS3Request.content
}
