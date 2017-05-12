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

type ReplicatePutJobSpectraS3Request struct {
    bucketName string
    content networking.ReaderWithSizeDecorator
    priority Priority
    queryParams *url.Values
}

func NewReplicatePutJobSpectraS3Request(bucketName string, requestPayload string) *ReplicatePutJobSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "start_bulk_put")
    queryParams.Set("replicate", "")

    return &ReplicatePutJobSpectraS3Request{
        bucketName: bucketName,
        content: buildStreamFromString(requestPayload),
        queryParams: queryParams,
    }
}

func (replicatePutJobSpectraS3Request *ReplicatePutJobSpectraS3Request) WithPriority(priority Priority) *ReplicatePutJobSpectraS3Request {
    replicatePutJobSpectraS3Request.priority = priority
    replicatePutJobSpectraS3Request.queryParams.Set("priority", priority.String())
    return replicatePutJobSpectraS3Request
}



func (ReplicatePutJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (replicatePutJobSpectraS3Request *ReplicatePutJobSpectraS3Request) Path() string {
    return "/_rest_/bucket/" + replicatePutJobSpectraS3Request.bucketName
}

func (replicatePutJobSpectraS3Request *ReplicatePutJobSpectraS3Request) QueryParams() *url.Values {
    return replicatePutJobSpectraS3Request.queryParams
}

func (ReplicatePutJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ReplicatePutJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (replicatePutJobSpectraS3Request *ReplicatePutJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return replicatePutJobSpectraS3Request.content
}
