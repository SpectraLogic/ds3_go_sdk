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

type GetNodesSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetNodesSpectraS3Request() *GetNodesSpectraS3Request {
    return &GetNodesSpectraS3Request{
    }
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithLastPage() *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.LastPage = true
    return getNodesSpectraS3Request
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithPageLength(pageLength int) *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.PageLength = &pageLength
    return getNodesSpectraS3Request
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithPageOffset(pageOffset int) *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.PageOffset = &pageOffset
    return getNodesSpectraS3Request
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getNodesSpectraS3Request
}

