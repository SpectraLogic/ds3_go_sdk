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

type GetTapeDensityDirectivesSpectraS3Request struct {
    Density TapeDriveType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PartitionId *string
    TapeType TapeType
}

func NewGetTapeDensityDirectivesSpectraS3Request() *GetTapeDensityDirectivesSpectraS3Request {
    return &GetTapeDensityDirectivesSpectraS3Request{
    }
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithDensity(density TapeDriveType) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.Density = density
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithLastPage() *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.LastPage = true
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPageLength(pageLength int) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.PageLength = &pageLength
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.PageOffset = &pageOffset
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPartitionId(partitionId string) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.PartitionId = &partitionId
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithTapeType(tapeType TapeType) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.TapeType = tapeType
    return getTapeDensityDirectivesSpectraS3Request
}

