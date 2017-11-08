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

type ModifyDs3TargetSpectraS3Request struct {
    AccessControlReplication Ds3TargetAccessControlReplication
    AdminAuthId *string
    AdminSecretKey *string
    DataPathEndPoint *string
    DataPathHttps *bool
    DataPathPort *int
    DataPathProxy *string
    DataPathVerifyCertificate *bool
    DefaultReadPreference TargetReadPreferenceType
    Ds3Target string
    Name *string
    PermitGoingOutOfSync *bool
    Quiesced Quiesced
    ReplicatedUserDefaultDataPolicy *string
}

func NewModifyDs3TargetSpectraS3Request(ds3Target string) *ModifyDs3TargetSpectraS3Request {
    return &ModifyDs3TargetSpectraS3Request{
        Ds3Target: ds3Target,
    }
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithAccessControlReplication(accessControlReplication Ds3TargetAccessControlReplication) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.AccessControlReplication = accessControlReplication
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithAdminAuthId(adminAuthId string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.AdminAuthId = &adminAuthId
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithAdminSecretKey(adminSecretKey string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.AdminSecretKey = &adminSecretKey
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DataPathHttps = &dataPathHttps
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathPort(dataPathPort int) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DataPathPort = &dataPathPort
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathProxy(dataPathProxy string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DataPathProxy = &dataPathProxy
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DataPathVerifyCertificate = &dataPathVerifyCertificate
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithName(name string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.Name = &name
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.Quiesced = quiesced
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithReplicatedUserDefaultDataPolicy(replicatedUserDefaultDataPolicy string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.ReplicatedUserDefaultDataPolicy = &replicatedUserDefaultDataPolicy
    return modifyDs3TargetSpectraS3Request
}

