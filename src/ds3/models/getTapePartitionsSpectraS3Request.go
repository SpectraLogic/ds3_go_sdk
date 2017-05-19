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
    "strconv"
)

type GetTapePartitionsSpectraS3Request struct {
    importExportConfiguration ImportExportConfiguration
    libraryId string
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    quiesced Quiesced
    serialNumber *string
    state TapePartitionState
    queryParams *url.Values
}

func NewGetTapePartitionsSpectraS3Request() *GetTapePartitionsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapePartitionsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithImportExportConfiguration(importExportConfiguration ImportExportConfiguration) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.importExportConfiguration = importExportConfiguration
    getTapePartitionsSpectraS3Request.queryParams.Set("import_export_configuration", importExportConfiguration.String())
    return getTapePartitionsSpectraS3Request
}
func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithLibraryId(libraryId string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.libraryId = libraryId
    getTapePartitionsSpectraS3Request.queryParams.Set("library_id", libraryId)
    return getTapePartitionsSpectraS3Request
}
func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.pageLength = pageLength
    getTapePartitionsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getTapePartitionsSpectraS3Request
}
func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.pageOffset = pageOffset
    getTapePartitionsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getTapePartitionsSpectraS3Request
}
func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.pageStartMarker = pageStartMarker
    getTapePartitionsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getTapePartitionsSpectraS3Request
}
func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.quiesced = quiesced
    getTapePartitionsSpectraS3Request.queryParams.Set("quiesced", quiesced.String())
    return getTapePartitionsSpectraS3Request
}
func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithState(state TapePartitionState) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.state = state
    getTapePartitionsSpectraS3Request.queryParams.Set("state", state.String())
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithName(name *string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.name = name
    if name != nil {
        getTapePartitionsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getTapePartitionsSpectraS3Request.queryParams.Set("name", "")
    }
    return getTapePartitionsSpectraS3Request
}
func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithSerialNumber(serialNumber *string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.serialNumber = serialNumber
    if serialNumber != nil {
        getTapePartitionsSpectraS3Request.queryParams.Set("serial_number", *serialNumber)
    } else {
        getTapePartitionsSpectraS3Request.queryParams.Set("serial_number", "")
    }
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithLastPage() *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.queryParams.Set("last_page", "")
    return getTapePartitionsSpectraS3Request
}

func (GetTapePartitionsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) Path() string {
    return "/_rest_/tape_partition"
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) QueryParams() *url.Values {
    return getTapePartitionsSpectraS3Request.queryParams
}

func (GetTapePartitionsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapePartitionsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapePartitionsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
