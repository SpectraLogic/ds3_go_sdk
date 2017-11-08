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

type GetDataPolicyAclsSpectraS3Request struct {
    DataPolicyId *string
    GroupId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetDataPolicyAclsSpectraS3Request() *GetDataPolicyAclsSpectraS3Request {
    return &GetDataPolicyAclsSpectraS3Request{
    }
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithGroupId(groupId string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.GroupId = &groupId
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithLastPage() *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.LastPage = true
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithPageLength(pageLength int) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.PageLength = &pageLength
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithPageOffset(pageOffset int) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.PageOffset = &pageOffset
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithUserId(userId string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.UserId = &userId
    return getDataPolicyAclsSpectraS3Request
}

