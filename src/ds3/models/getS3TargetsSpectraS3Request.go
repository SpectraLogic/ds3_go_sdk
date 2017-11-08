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

type GetS3TargetsSpectraS3Request struct {
    AccessKey *string
    DataPathEndPoint *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PermitGoingOutOfSync *bool
    Quiesced Quiesced
    Region S3Region
    State TargetState
}

func NewGetS3TargetsSpectraS3Request() *GetS3TargetsSpectraS3Request {
    return &GetS3TargetsSpectraS3Request{
    }
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithAccessKey(accessKey string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.AccessKey = &accessKey
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithHttps(https bool) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.Https = &https
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithLastPage() *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.LastPage = true
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithName(name string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.Name = &name
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.PageLength = &pageLength
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.PageOffset = &pageOffset
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.Quiesced = quiesced
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithRegion(region S3Region) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.Region = region
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithState(state TargetState) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.State = state
    return getS3TargetsSpectraS3Request
}

