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

type GetTapeFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TapeDriveId *string
    TapeFailureType TapeFailureType
    TapeId *string
}

func NewGetTapeFailuresSpectraS3Request() *GetTapeFailuresSpectraS3Request {
    return &GetTapeFailuresSpectraS3Request{
    }
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithLastPage() *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.LastPage = true
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithPageLength(pageLength int) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.PageLength = &pageLength
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.PageOffset = &pageOffset
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithTapeDriveId(tapeDriveId string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.TapeDriveId = &tapeDriveId
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithTapeId(tapeId string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.TapeId = &tapeId
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithTapeFailureType(tapeFailureType TapeFailureType) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.TapeFailureType = tapeFailureType
    return getTapeFailuresSpectraS3Request
}

