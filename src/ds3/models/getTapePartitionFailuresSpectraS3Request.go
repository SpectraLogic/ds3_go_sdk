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

type GetTapePartitionFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PartitionId *string
    TapePartitionFailureType TapePartitionFailureType
}

func NewGetTapePartitionFailuresSpectraS3Request() *GetTapePartitionFailuresSpectraS3Request {
    return &GetTapePartitionFailuresSpectraS3Request{
    }
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithLastPage() *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.LastPage = true
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.PageLength = &pageLength
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.PageOffset = &pageOffset
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPartitionId(partitionId string) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.PartitionId = &partitionId
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithTapePartitionFailureType(tapePartitionFailureType TapePartitionFailureType) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.TapePartitionFailureType = tapePartitionFailureType
    return getTapePartitionFailuresSpectraS3Request
}

