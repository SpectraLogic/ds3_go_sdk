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

type GetDs3TargetFailuresSpectraS3Request struct {
    errorMessage *string
    pageLength int
    pageOffset int
    pageStartMarker string
    targetFailureType TargetFailureType
    targetId string
    queryParams *url.Values
}

func NewGetDs3TargetFailuresSpectraS3Request() *GetDs3TargetFailuresSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDs3TargetFailuresSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.pageLength = pageLength
    getDs3TargetFailuresSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getDs3TargetFailuresSpectraS3Request
}
func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.pageOffset = pageOffset
    getDs3TargetFailuresSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getDs3TargetFailuresSpectraS3Request
}
func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.pageStartMarker = pageStartMarker
    getDs3TargetFailuresSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getDs3TargetFailuresSpectraS3Request
}
func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithTargetId(targetId string) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.targetId = targetId
    getDs3TargetFailuresSpectraS3Request.queryParams.Set("target_id", targetId)
    return getDs3TargetFailuresSpectraS3Request
}
func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithTargetFailureType(targetFailureType TargetFailureType) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.targetFailureType = targetFailureType
    getDs3TargetFailuresSpectraS3Request.queryParams.Set("type", targetFailureType.String())
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithErrorMessage(errorMessage *string) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.errorMessage = errorMessage
    if errorMessage != nil {
        getDs3TargetFailuresSpectraS3Request.queryParams.Set("error_message", *errorMessage)
    } else {
        getDs3TargetFailuresSpectraS3Request.queryParams.Set("error_message", "")
    }
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithLastPage() *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.queryParams.Set("last_page", "")
    return getDs3TargetFailuresSpectraS3Request
}

func (GetDs3TargetFailuresSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) Path() string {
    return "/_rest_/ds3_target_failure"
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) QueryParams() *url.Values {
    return getDs3TargetFailuresSpectraS3Request.queryParams
}

func (GetDs3TargetFailuresSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDs3TargetFailuresSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDs3TargetFailuresSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
