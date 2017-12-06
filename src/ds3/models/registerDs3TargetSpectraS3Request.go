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

type RegisterDs3TargetSpectraS3Request struct {
    AccessControlReplication Ds3TargetAccessControlReplication
    AdminAuthId string
    AdminSecretKey string
    DataPathEndPoint string
    DataPathHttps *bool
    DataPathPort *int
    DataPathProxy *string
    DataPathVerifyCertificate *bool
    DefaultReadPreference TargetReadPreferenceType
    Name string
    PermitGoingOutOfSync *bool
    ReplicatedUserDefaultDataPolicy *string
}

func NewRegisterDs3TargetSpectraS3Request(adminAuthId string, adminSecretKey string, dataPathEndPoint string, name string) *RegisterDs3TargetSpectraS3Request {
    return &RegisterDs3TargetSpectraS3Request{
        AdminAuthId: adminAuthId,
        AdminSecretKey: adminSecretKey,
        DataPathEndPoint: dataPathEndPoint,
        Name: name,
    }
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithAccessControlReplication(accessControlReplication Ds3TargetAccessControlReplication) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.AccessControlReplication = accessControlReplication
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.DataPathHttps = &dataPathHttps
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathPort(dataPathPort int) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.DataPathPort = &dataPathPort
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathProxy(dataPathProxy string) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.DataPathProxy = &dataPathProxy
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.DataPathVerifyCertificate = &dataPathVerifyCertificate
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithReplicatedUserDefaultDataPolicy(replicatedUserDefaultDataPolicy string) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.ReplicatedUserDefaultDataPolicy = &replicatedUserDefaultDataPolicy
    return registerDs3TargetSpectraS3Request
}

