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

type GetGroupsSpectraS3Request struct {
    BuiltIn *bool
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetGroupsSpectraS3Request() *GetGroupsSpectraS3Request {
    return &GetGroupsSpectraS3Request{
    }
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithBuiltIn(builtIn bool) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.BuiltIn = &builtIn
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithLastPage() *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.LastPage = true
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithName(name string) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.Name = &name
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithPageLength(pageLength int) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.PageLength = &pageLength
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithPageOffset(pageOffset int) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.PageOffset = &pageOffset
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getGroupsSpectraS3Request
}

