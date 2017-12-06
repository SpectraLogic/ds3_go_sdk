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

type GetSuspectBlobTapesSpectraS3Request struct {
    BlobId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TapeId *string
}

func NewGetSuspectBlobTapesSpectraS3Request() *GetSuspectBlobTapesSpectraS3Request {
    return &GetSuspectBlobTapesSpectraS3Request{
    }
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.BlobId = &blobId
    return getSuspectBlobTapesSpectraS3Request
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithLastPage() *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.LastPage = true
    return getSuspectBlobTapesSpectraS3Request
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.PageLength = &pageLength
    return getSuspectBlobTapesSpectraS3Request
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.PageOffset = &pageOffset
    return getSuspectBlobTapesSpectraS3Request
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectBlobTapesSpectraS3Request
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithTapeId(tapeId string) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.TapeId = &tapeId
    return getSuspectBlobTapesSpectraS3Request
}

