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

type GetUsersSpectraS3Request struct {
    AuthId *string
    DefaultDataPolicyId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetUsersSpectraS3Request() *GetUsersSpectraS3Request {
    return &GetUsersSpectraS3Request{
    }
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithAuthId(authId string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.AuthId = &authId
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithDefaultDataPolicyId(defaultDataPolicyId string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.DefaultDataPolicyId = &defaultDataPolicyId
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithLastPage() *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.LastPage = true
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithName(name string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.Name = &name
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithPageLength(pageLength int) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.PageLength = &pageLength
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithPageOffset(pageOffset int) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.PageOffset = &pageOffset
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.PageStartMarker = &pageStartMarker
    return getUsersSpectraS3Request
}

