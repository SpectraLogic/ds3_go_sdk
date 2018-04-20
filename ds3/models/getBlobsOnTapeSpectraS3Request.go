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

type GetBlobsOnTapeSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TapeId string
}

func NewGetBlobsOnTapeSpectraS3Request(tapeId string) *GetBlobsOnTapeSpectraS3Request {
    return &GetBlobsOnTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (getBlobsOnTapeSpectraS3Request *GetBlobsOnTapeSpectraS3Request) WithLastPage() *GetBlobsOnTapeSpectraS3Request {
    getBlobsOnTapeSpectraS3Request.LastPage = true
    return getBlobsOnTapeSpectraS3Request
}

func (getBlobsOnTapeSpectraS3Request *GetBlobsOnTapeSpectraS3Request) WithPageLength(pageLength int) *GetBlobsOnTapeSpectraS3Request {
    getBlobsOnTapeSpectraS3Request.PageLength = &pageLength
    return getBlobsOnTapeSpectraS3Request
}

func (getBlobsOnTapeSpectraS3Request *GetBlobsOnTapeSpectraS3Request) WithPageOffset(pageOffset int) *GetBlobsOnTapeSpectraS3Request {
    getBlobsOnTapeSpectraS3Request.PageOffset = &pageOffset
    return getBlobsOnTapeSpectraS3Request
}

func (getBlobsOnTapeSpectraS3Request *GetBlobsOnTapeSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetBlobsOnTapeSpectraS3Request {
    getBlobsOnTapeSpectraS3Request.PageStartMarker = &pageStartMarker
    return getBlobsOnTapeSpectraS3Request
}

