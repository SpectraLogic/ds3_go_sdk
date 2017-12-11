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

type GetAzureTargetFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetFailureType TargetFailureType
    TargetId *string
}

func NewGetAzureTargetFailuresSpectraS3Request() *GetAzureTargetFailuresSpectraS3Request {
    return &GetAzureTargetFailuresSpectraS3Request{
    }
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithLastPage() *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.LastPage = true
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.PageLength = &pageLength
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.PageOffset = &pageOffset
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithTargetId(targetId string) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.TargetId = &targetId
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithTargetFailureType(targetFailureType TargetFailureType) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.TargetFailureType = targetFailureType
    return getAzureTargetFailuresSpectraS3Request
}

