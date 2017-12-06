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

type GetPoolPartitionsSpectraS3Request struct {
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PoolType PoolType
}

func NewGetPoolPartitionsSpectraS3Request() *GetPoolPartitionsSpectraS3Request {
    return &GetPoolPartitionsSpectraS3Request{
    }
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithLastPage() *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.LastPage = true
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithName(name string) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.Name = &name
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPageLength(pageLength int) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.PageLength = &pageLength
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.PageOffset = &pageOffset
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPoolType(poolType PoolType) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.PoolType = poolType
    return getPoolPartitionsSpectraS3Request
}

