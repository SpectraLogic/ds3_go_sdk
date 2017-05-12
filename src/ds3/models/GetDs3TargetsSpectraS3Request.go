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

type GetDs3TargetsSpectraS3Request struct {
    adminAuthId *string
    dataPathEndPoint *string
    dataPathHttps bool
    dataPathPort *int
    dataPathProxy *string
    dataPathVerifyCertificate bool
    defaultReadPreference TargetReadPreferenceType
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    permitGoingOutOfSync bool
    quiesced Quiesced
    state TargetState
    queryParams *url.Values
}

func NewGetDs3TargetsSpectraS3Request() *GetDs3TargetsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDs3TargetsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.dataPathHttps = dataPathHttps
    getDs3TargetsSpectraS3Request.queryParams.Set("data_path_https", strconv.FormatBool(dataPathHttps))
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.dataPathVerifyCertificate = dataPathVerifyCertificate
    getDs3TargetsSpectraS3Request.queryParams.Set("data_path_verify_certificate", strconv.FormatBool(dataPathVerifyCertificate))
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.defaultReadPreference = defaultReadPreference
    getDs3TargetsSpectraS3Request.queryParams.Set("default_read_preference", defaultReadPreference.String())
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.pageLength = pageLength
    getDs3TargetsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.pageOffset = pageOffset
    getDs3TargetsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.pageStartMarker = pageStartMarker
    getDs3TargetsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.permitGoingOutOfSync = permitGoingOutOfSync
    getDs3TargetsSpectraS3Request.queryParams.Set("permit_going_out_of_sync", strconv.FormatBool(permitGoingOutOfSync))
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.quiesced = quiesced
    getDs3TargetsSpectraS3Request.queryParams.Set("quiesced", quiesced.String())
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithState(state TargetState) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.state = state
    getDs3TargetsSpectraS3Request.queryParams.Set("state", state.String())
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithAdminAuthId(adminAuthId *string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.adminAuthId = adminAuthId
    if adminAuthId != nil {
        getDs3TargetsSpectraS3Request.queryParams.Set("admin_auth_id", *adminAuthId)
    } else {
        getDs3TargetsSpectraS3Request.queryParams.Set("admin_auth_id", "")
    }
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint *string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.dataPathEndPoint = dataPathEndPoint
    if dataPathEndPoint != nil {
        getDs3TargetsSpectraS3Request.queryParams.Set("data_path_end_point", *dataPathEndPoint)
    } else {
        getDs3TargetsSpectraS3Request.queryParams.Set("data_path_end_point", "")
    }
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathPort(dataPathPort *int) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.dataPathPort = dataPathPort
    if dataPathPort != nil {
        getDs3TargetsSpectraS3Request.queryParams.Set("data_path_port", strconv.Itoa(*dataPathPort))
    } else {
        getDs3TargetsSpectraS3Request.queryParams.Set("data_path_port", "")
    }
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathProxy(dataPathProxy *string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.dataPathProxy = dataPathProxy
    if dataPathProxy != nil {
        getDs3TargetsSpectraS3Request.queryParams.Set("data_path_proxy", *dataPathProxy)
    } else {
        getDs3TargetsSpectraS3Request.queryParams.Set("data_path_proxy", "")
    }
    return getDs3TargetsSpectraS3Request
}
func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithName(name *string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.name = name
    if name != nil {
        getDs3TargetsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getDs3TargetsSpectraS3Request.queryParams.Set("name", "")
    }
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithLastPage() *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.queryParams.Set("last_page", "")
    return getDs3TargetsSpectraS3Request
}

func (GetDs3TargetsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) Path() string {
    return "/_rest_/ds3_target"
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) QueryParams() *url.Values {
    return getDs3TargetsSpectraS3Request.queryParams
}

func (GetDs3TargetsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDs3TargetsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDs3TargetsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
