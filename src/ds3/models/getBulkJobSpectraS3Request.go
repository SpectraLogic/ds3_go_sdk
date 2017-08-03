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

type GetBulkJobSpectraS3Request struct {
    bucketName string
    aggregating bool
    chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    content networking.ReaderWithSizeDecorator
    implicitJobIdResolution bool
    name *string
    priority Priority
    queryParams *url.Values
}

//TODO update autogen to add second constructor
func NewGetBulkJobSpectraS3Request(bucketName string, objectNames []string) *GetBulkJobSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "start_bulk_get")

    return &GetBulkJobSpectraS3Request{
        bucketName: bucketName,
        content: buildDs3ObjectStreamFromNames(objectNames),
        queryParams: queryParams,
    }
}

func NewGetBulkJobSpectraS3RequestWithPartialObjects(bucketName string, objects []Ds3GetObject) *GetBulkJobSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "start_bulk_get")

    return &GetBulkJobSpectraS3Request{
        bucketName: bucketName,
        content: buildDs3GetObjectListStream(objects),
        queryParams: queryParams,
    }
}

func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) WithAggregating(aggregating bool) *GetBulkJobSpectraS3Request {
    getBulkJobSpectraS3Request.aggregating = aggregating
    getBulkJobSpectraS3Request.queryParams.Set("aggregating", strconv.FormatBool(aggregating))
    return getBulkJobSpectraS3Request
}
func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetBulkJobSpectraS3Request {
    getBulkJobSpectraS3Request.chunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    getBulkJobSpectraS3Request.queryParams.Set("chunk_client_processing_order_guarantee", chunkClientProcessingOrderGuarantee.String())
    return getBulkJobSpectraS3Request
}
func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) WithImplicitJobIdResolution(implicitJobIdResolution bool) *GetBulkJobSpectraS3Request {
    getBulkJobSpectraS3Request.implicitJobIdResolution = implicitJobIdResolution
    getBulkJobSpectraS3Request.queryParams.Set("implicit_job_id_resolution", strconv.FormatBool(implicitJobIdResolution))
    return getBulkJobSpectraS3Request
}
func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) WithPriority(priority Priority) *GetBulkJobSpectraS3Request {
    getBulkJobSpectraS3Request.priority = priority
    getBulkJobSpectraS3Request.queryParams.Set("priority", priority.String())
    return getBulkJobSpectraS3Request
}

func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) WithName(name *string) *GetBulkJobSpectraS3Request {
    getBulkJobSpectraS3Request.name = name
    if name != nil {
        getBulkJobSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getBulkJobSpectraS3Request.queryParams.Set("name", "")
    }
    return getBulkJobSpectraS3Request
}


func (GetBulkJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) Path() string {
    return "/_rest_/bucket/" + getBulkJobSpectraS3Request.bucketName
}

func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) QueryParams() *url.Values {
    return getBulkJobSpectraS3Request.queryParams
}

func (GetBulkJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetBulkJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return getBulkJobSpectraS3Request.content
}
