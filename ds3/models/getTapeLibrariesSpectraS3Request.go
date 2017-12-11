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

type GetTapeLibrariesSpectraS3Request struct {
    LastPage bool
    ManagementUrl *string
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    SerialNumber *string
}

func NewGetTapeLibrariesSpectraS3Request() *GetTapeLibrariesSpectraS3Request {
    return &GetTapeLibrariesSpectraS3Request{
    }
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithLastPage() *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.LastPage = true
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithManagementUrl(managementUrl string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.ManagementUrl = &managementUrl
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithName(name string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.Name = &name
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithPageLength(pageLength int) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.PageLength = &pageLength
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.PageOffset = &pageOffset
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithSerialNumber(serialNumber string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.SerialNumber = &serialNumber
    return getTapeLibrariesSpectraS3Request
}

