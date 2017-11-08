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

type GetSystemFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    SystemFailureType SystemFailureType
}

func NewGetSystemFailuresSpectraS3Request() *GetSystemFailuresSpectraS3Request {
    return &GetSystemFailuresSpectraS3Request{
    }
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithLastPage() *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.LastPage = true
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithPageLength(pageLength int) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.PageLength = &pageLength
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.PageOffset = &pageOffset
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithSystemFailureType(systemFailureType SystemFailureType) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.SystemFailureType = systemFailureType
    return getSystemFailuresSpectraS3Request
}

