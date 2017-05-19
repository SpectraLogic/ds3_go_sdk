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

type GetAzureTargetFailuresSpectraS3Request struct {
    errorMessage *string
    pageLength int
    pageOffset int
    pageStartMarker string
    targetFailureType TargetFailureType
    targetId string
    queryParams *url.Values
}

func NewGetAzureTargetFailuresSpectraS3Request() *GetAzureTargetFailuresSpectraS3Request {
    queryParams := &url.Values{}

    return &GetAzureTargetFailuresSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.pageLength = pageLength
    getAzureTargetFailuresSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getAzureTargetFailuresSpectraS3Request
}
func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.pageOffset = pageOffset
    getAzureTargetFailuresSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getAzureTargetFailuresSpectraS3Request
}
func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.pageStartMarker = pageStartMarker
    getAzureTargetFailuresSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getAzureTargetFailuresSpectraS3Request
}
func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithTargetId(targetId string) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.targetId = targetId
    getAzureTargetFailuresSpectraS3Request.queryParams.Set("target_id", targetId)
    return getAzureTargetFailuresSpectraS3Request
}
func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithTargetFailureType(targetFailureType TargetFailureType) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.targetFailureType = targetFailureType
    getAzureTargetFailuresSpectraS3Request.queryParams.Set("type", targetFailureType.String())
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithErrorMessage(errorMessage *string) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.errorMessage = errorMessage
    if errorMessage != nil {
        getAzureTargetFailuresSpectraS3Request.queryParams.Set("error_message", *errorMessage)
    } else {
        getAzureTargetFailuresSpectraS3Request.queryParams.Set("error_message", "")
    }
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithLastPage() *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.queryParams.Set("last_page", "")
    return getAzureTargetFailuresSpectraS3Request
}

func (GetAzureTargetFailuresSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) Path() string {
    return "/_rest_/azure_target_failure"
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) QueryParams() *url.Values {
    return getAzureTargetFailuresSpectraS3Request.queryParams
}

func (GetAzureTargetFailuresSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetAzureTargetFailuresSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetAzureTargetFailuresSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
