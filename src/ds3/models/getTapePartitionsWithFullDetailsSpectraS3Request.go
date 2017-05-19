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

type GetTapePartitionsWithFullDetailsSpectraS3Request struct {
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

func NewGetTapePartitionsWithFullDetailsSpectraS3Request() *GetTapePartitionsWithFullDetailsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("full_details", "")

    return &GetTapePartitionsWithFullDetailsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithImportExportConfiguration(importExportConfiguration ImportExportConfiguration) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.importExportConfiguration = importExportConfiguration
    getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("import_export_configuration", importExportConfiguration.String())
    return getTapePartitionsWithFullDetailsSpectraS3Request
}
func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithLibraryId(libraryId string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.libraryId = libraryId
    getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("library_id", libraryId)
    return getTapePartitionsWithFullDetailsSpectraS3Request
}
func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.pageLength = pageLength
    getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getTapePartitionsWithFullDetailsSpectraS3Request
}
func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.pageOffset = pageOffset
    getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getTapePartitionsWithFullDetailsSpectraS3Request
}
func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.pageStartMarker = pageStartMarker
    getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getTapePartitionsWithFullDetailsSpectraS3Request
}
func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.quiesced = quiesced
    getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("quiesced", quiesced.String())
    return getTapePartitionsWithFullDetailsSpectraS3Request
}
func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithState(state TapePartitionState) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.state = state
    getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("state", state.String())
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithName(name *string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.name = name
    if name != nil {
        getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("name", "")
    }
    return getTapePartitionsWithFullDetailsSpectraS3Request
}
func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithSerialNumber(serialNumber *string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.serialNumber = serialNumber
    if serialNumber != nil {
        getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("serial_number", *serialNumber)
    } else {
        getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("serial_number", "")
    }
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithLastPage() *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.queryParams.Set("last_page", "")
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (GetTapePartitionsWithFullDetailsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) Path() string {
    return "/_rest_/tape_partition"
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) QueryParams() *url.Values {
    return getTapePartitionsWithFullDetailsSpectraS3Request.queryParams
}

func (GetTapePartitionsWithFullDetailsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapePartitionsWithFullDetailsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapePartitionsWithFullDetailsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
