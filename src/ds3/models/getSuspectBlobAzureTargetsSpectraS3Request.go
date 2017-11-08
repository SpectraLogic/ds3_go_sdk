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

type GetSuspectBlobAzureTargetsSpectraS3Request struct {
    BlobId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetId *string
}

func NewGetSuspectBlobAzureTargetsSpectraS3Request() *GetSuspectBlobAzureTargetsSpectraS3Request {
    return &GetSuspectBlobAzureTargetsSpectraS3Request{
    }
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.BlobId = &blobId
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithLastPage() *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.LastPage = true
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.PageLength = &pageLength
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.PageOffset = &pageOffset
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithTargetId(targetId string) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.TargetId = &targetId
    return getSuspectBlobAzureTargetsSpectraS3Request
}

