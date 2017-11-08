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

type GetAzureTargetsSpectraS3Request struct {
    AccountName *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PermitGoingOutOfSync *bool
    Quiesced Quiesced
    State TargetState
}

func NewGetAzureTargetsSpectraS3Request() *GetAzureTargetsSpectraS3Request {
    return &GetAzureTargetsSpectraS3Request{
    }
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithAccountName(accountName string) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.AccountName = &accountName
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithHttps(https bool) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.Https = &https
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithLastPage() *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.LastPage = true
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithName(name string) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.Name = &name
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.PageLength = &pageLength
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.PageOffset = &pageOffset
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.Quiesced = quiesced
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithState(state TargetState) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.State = state
    return getAzureTargetsSpectraS3Request
}

