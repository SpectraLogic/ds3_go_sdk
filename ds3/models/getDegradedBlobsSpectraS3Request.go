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

type GetDegradedBlobsSpectraS3Request struct {
    BlobId *string
    BucketId *string
    Ds3ReplicationRuleId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PersistenceRuleId *string
}

func NewGetDegradedBlobsSpectraS3Request() *GetDegradedBlobsSpectraS3Request {
    return &GetDegradedBlobsSpectraS3Request{
    }
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithBlobId(blobId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.BlobId = &blobId
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithBucketId(bucketId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.BucketId = &bucketId
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithDs3ReplicationRuleId(ds3ReplicationRuleId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.Ds3ReplicationRuleId = &ds3ReplicationRuleId
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithLastPage() *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.LastPage = true
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPageLength(pageLength int) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.PageLength = &pageLength
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.PageOffset = &pageOffset
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPersistenceRuleId(persistenceRuleId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.PersistenceRuleId = &persistenceRuleId
    return getDegradedBlobsSpectraS3Request
}

