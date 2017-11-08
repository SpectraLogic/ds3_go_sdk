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

type GetSuspectBlobS3TargetsSpectraS3Request struct {
    BlobId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetId *string
}

func NewGetSuspectBlobS3TargetsSpectraS3Request() *GetSuspectBlobS3TargetsSpectraS3Request {
    return &GetSuspectBlobS3TargetsSpectraS3Request{
    }
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.BlobId = &blobId
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithLastPage() *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.LastPage = true
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.PageLength = &pageLength
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.PageOffset = &pageOffset
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithTargetId(targetId string) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.TargetId = &targetId
    return getSuspectBlobS3TargetsSpectraS3Request
}

