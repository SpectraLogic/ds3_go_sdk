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

type GetS3TargetFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetFailureType TargetFailureType
    TargetId *string
}

func NewGetS3TargetFailuresSpectraS3Request() *GetS3TargetFailuresSpectraS3Request {
    return &GetS3TargetFailuresSpectraS3Request{
    }
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithLastPage() *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.LastPage = true
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.PageLength = &pageLength
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.PageOffset = &pageOffset
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithTargetId(targetId string) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.TargetId = &targetId
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithTargetFailureType(targetFailureType TargetFailureType) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.TargetFailureType = targetFailureType
    return getS3TargetFailuresSpectraS3Request
}

