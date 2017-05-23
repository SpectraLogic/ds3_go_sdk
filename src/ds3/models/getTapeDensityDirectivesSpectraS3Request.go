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

type GetTapeDensityDirectivesSpectraS3Request struct {
    density TapeDriveType
    pageLength int
    pageOffset int
    pageStartMarker string
    partitionId string
    tapeType TapeType
    queryParams *url.Values
}

func NewGetTapeDensityDirectivesSpectraS3Request() *GetTapeDensityDirectivesSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapeDensityDirectivesSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithDensity(density TapeDriveType) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.density = density
    getTapeDensityDirectivesSpectraS3Request.queryParams.Set("density", density.String())
    return getTapeDensityDirectivesSpectraS3Request
}
func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPageLength(pageLength int) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.pageLength = pageLength
    getTapeDensityDirectivesSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getTapeDensityDirectivesSpectraS3Request
}
func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.pageOffset = pageOffset
    getTapeDensityDirectivesSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getTapeDensityDirectivesSpectraS3Request
}
func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.pageStartMarker = pageStartMarker
    getTapeDensityDirectivesSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getTapeDensityDirectivesSpectraS3Request
}
func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPartitionId(partitionId string) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.partitionId = partitionId
    getTapeDensityDirectivesSpectraS3Request.queryParams.Set("partition_id", partitionId)
    return getTapeDensityDirectivesSpectraS3Request
}
func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithTapeType(tapeType TapeType) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.tapeType = tapeType
    getTapeDensityDirectivesSpectraS3Request.queryParams.Set("tape_type", tapeType.String())
    return getTapeDensityDirectivesSpectraS3Request
}


func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithLastPage() *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.queryParams.Set("last_page", "")
    return getTapeDensityDirectivesSpectraS3Request
}

func (GetTapeDensityDirectivesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) Path() string {
    return "/_rest_/tape_density_directive"
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) QueryParams() *url.Values {
    return getTapeDensityDirectivesSpectraS3Request.queryParams
}

func (GetTapeDensityDirectivesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapeDensityDirectivesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapeDensityDirectivesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
