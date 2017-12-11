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

type GetPoolFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PoolFailureType PoolFailureType
    PoolId *string
}

func NewGetPoolFailuresSpectraS3Request() *GetPoolFailuresSpectraS3Request {
    return &GetPoolFailuresSpectraS3Request{
    }
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithLastPage() *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.LastPage = true
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPageLength(pageLength int) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.PageLength = &pageLength
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.PageOffset = &pageOffset
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPoolId(poolId string) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.PoolId = &poolId
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPoolFailureType(poolFailureType PoolFailureType) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.PoolFailureType = poolFailureType
    return getPoolFailuresSpectraS3Request
}

