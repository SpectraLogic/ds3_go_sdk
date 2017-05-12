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

type GetSystemFailuresSpectraS3Request struct {
    errorMessage *string
    pageLength int
    pageOffset int
    pageStartMarker string
    systemFailureType SystemFailureType
    queryParams *url.Values
}

func NewGetSystemFailuresSpectraS3Request() *GetSystemFailuresSpectraS3Request {
    queryParams := &url.Values{}

    return &GetSystemFailuresSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithPageLength(pageLength int) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.pageLength = pageLength
    getSystemFailuresSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getSystemFailuresSpectraS3Request
}
func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.pageOffset = pageOffset
    getSystemFailuresSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getSystemFailuresSpectraS3Request
}
func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.pageStartMarker = pageStartMarker
    getSystemFailuresSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getSystemFailuresSpectraS3Request
}
func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithSystemFailureType(systemFailureType SystemFailureType) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.systemFailureType = systemFailureType
    getSystemFailuresSpectraS3Request.queryParams.Set("type", systemFailureType.String())
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithErrorMessage(errorMessage *string) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.errorMessage = errorMessage
    if errorMessage != nil {
        getSystemFailuresSpectraS3Request.queryParams.Set("error_message", *errorMessage)
    } else {
        getSystemFailuresSpectraS3Request.queryParams.Set("error_message", "")
    }
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithLastPage() *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.queryParams.Set("last_page", "")
    return getSystemFailuresSpectraS3Request
}

func (GetSystemFailuresSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) Path() string {
    return "/_rest_/system_failure"
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) QueryParams() *url.Values {
    return getSystemFailuresSpectraS3Request.queryParams
}

func (GetSystemFailuresSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetSystemFailuresSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetSystemFailuresSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
