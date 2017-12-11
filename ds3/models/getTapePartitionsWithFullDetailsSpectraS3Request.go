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

type GetTapePartitionsWithFullDetailsSpectraS3Request struct {
    ImportExportConfiguration ImportExportConfiguration
    LastPage bool
    LibraryId *string
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    Quiesced Quiesced
    SerialNumber *string
    State TapePartitionState
}

func NewGetTapePartitionsWithFullDetailsSpectraS3Request() *GetTapePartitionsWithFullDetailsSpectraS3Request {
    return &GetTapePartitionsWithFullDetailsSpectraS3Request{
    }
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithImportExportConfiguration(importExportConfiguration ImportExportConfiguration) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.ImportExportConfiguration = importExportConfiguration
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithLastPage() *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.LastPage = true
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithLibraryId(libraryId string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.LibraryId = &libraryId
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithName(name string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.Name = &name
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.PageLength = &pageLength
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.PageOffset = &pageOffset
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.Quiesced = quiesced
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithSerialNumber(serialNumber string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.SerialNumber = &serialNumber
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithState(state TapePartitionState) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.State = state
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

