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
)

type PutTapeDensityDirectiveSpectraS3Request struct {
    density TapeDriveType
    partitionId string
    tapeType TapeType
    queryParams *url.Values
}

func NewPutTapeDensityDirectiveSpectraS3Request(density TapeDriveType, partitionId string, tapeType TapeType) *PutTapeDensityDirectiveSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("density", density.String())
    queryParams.Set("partition_id", partitionId)
    queryParams.Set("tape_type", tapeType.String())

    return &PutTapeDensityDirectiveSpectraS3Request{
        density: density,
        partitionId: partitionId,
        tapeType: tapeType,
        queryParams: queryParams,
    }
}




func (PutTapeDensityDirectiveSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putTapeDensityDirectiveSpectraS3Request *PutTapeDensityDirectiveSpectraS3Request) Path() string {
    return "/_rest_/tape_density_directive"
}

func (putTapeDensityDirectiveSpectraS3Request *PutTapeDensityDirectiveSpectraS3Request) QueryParams() *url.Values {
    return putTapeDensityDirectiveSpectraS3Request.queryParams
}

func (PutTapeDensityDirectiveSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutTapeDensityDirectiveSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutTapeDensityDirectiveSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
