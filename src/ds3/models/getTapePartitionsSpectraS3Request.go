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

type GetTapePartitionsSpectraS3Request struct {
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

func NewGetTapePartitionsSpectraS3Request() *GetTapePartitionsSpectraS3Request {
    return &GetTapePartitionsSpectraS3Request{
    }
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithImportExportConfiguration(importExportConfiguration ImportExportConfiguration) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.ImportExportConfiguration = importExportConfiguration
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithLastPage() *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.LastPage = true
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithLibraryId(libraryId string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.LibraryId = &libraryId
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithName(name string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.Name = &name
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.PageLength = &pageLength
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.PageOffset = &pageOffset
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.Quiesced = quiesced
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithSerialNumber(serialNumber string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.SerialNumber = &serialNumber
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithState(state TapePartitionState) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.State = state
    return getTapePartitionsSpectraS3Request
}

