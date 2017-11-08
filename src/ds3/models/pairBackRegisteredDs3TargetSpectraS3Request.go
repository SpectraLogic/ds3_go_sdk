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

type PairBackRegisteredDs3TargetSpectraS3Request struct {
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
    ReplicatedUserDefaultDataPolicy *string
}

func NewPairBackRegisteredDs3TargetSpectraS3Request(ds3Target string) *PairBackRegisteredDs3TargetSpectraS3Request {
    return &PairBackRegisteredDs3TargetSpectraS3Request{
        Ds3Target: ds3Target,
    }
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithAccessControlReplication(accessControlReplication Ds3TargetAccessControlReplication) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.AccessControlReplication = accessControlReplication
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithAdminAuthId(adminAuthId string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.AdminAuthId = &adminAuthId
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithAdminSecretKey(adminSecretKey string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.AdminSecretKey = &adminSecretKey
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DataPathHttps = &dataPathHttps
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathPort(dataPathPort int) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DataPathPort = &dataPathPort
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathProxy(dataPathProxy string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DataPathProxy = &dataPathProxy
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DataPathVerifyCertificate = &dataPathVerifyCertificate
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithName(name string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.Name = &name
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithReplicatedUserDefaultDataPolicy(replicatedUserDefaultDataPolicy string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.ReplicatedUserDefaultDataPolicy = &replicatedUserDefaultDataPolicy
    return pairBackRegisteredDs3TargetSpectraS3Request
}

