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

type GetDs3TargetsSpectraS3Request struct {
    AdminAuthId *string
    DataPathEndPoint *string
    DataPathHttps *bool
    DataPathPort *int
    DataPathProxy *string
    DataPathVerifyCertificate *bool
    DefaultReadPreference TargetReadPreferenceType
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PermitGoingOutOfSync *bool
    Quiesced Quiesced
    State TargetState
}

func NewGetDs3TargetsSpectraS3Request() *GetDs3TargetsSpectraS3Request {
    return &GetDs3TargetsSpectraS3Request{
    }
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithAdminAuthId(adminAuthId string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.AdminAuthId = &adminAuthId
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DataPathHttps = &dataPathHttps
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathPort(dataPathPort int) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DataPathPort = &dataPathPort
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathProxy(dataPathProxy string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DataPathProxy = &dataPathProxy
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DataPathVerifyCertificate = &dataPathVerifyCertificate
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithLastPage() *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.LastPage = true
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithName(name string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.Name = &name
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.PageLength = &pageLength
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.PageOffset = &pageOffset
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.Quiesced = quiesced
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithState(state TargetState) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.State = state
    return getDs3TargetsSpectraS3Request
}

