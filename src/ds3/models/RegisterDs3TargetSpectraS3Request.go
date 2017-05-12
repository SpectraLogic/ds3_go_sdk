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

import (
    "net/url"
    "net/http"
    "ds3/networking"
    "strconv"
)

type RegisterDs3TargetSpectraS3Request struct {
    accessControlReplication Ds3TargetAccessControlReplication
    adminAuthId *string
    adminSecretKey *string
    dataPathEndPoint *string
    dataPathHttps bool
    dataPathPort *int
    dataPathProxy *string
    dataPathVerifyCertificate bool
    defaultReadPreference TargetReadPreferenceType
    name *string
    permitGoingOutOfSync bool
    replicatedUserDefaultDataPolicy *string
    queryParams *url.Values
}

func NewRegisterDs3TargetSpectraS3Request(adminAuthId *string, adminSecretKey *string, dataPathEndPoint *string, name *string) *RegisterDs3TargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("admin_auth_id", *adminAuthId)
    queryParams.Set("admin_secret_key", *adminSecretKey)
    queryParams.Set("data_path_end_point", *dataPathEndPoint)
    queryParams.Set("name", *name)

    return &RegisterDs3TargetSpectraS3Request{
        adminAuthId: adminAuthId,
        adminSecretKey: adminSecretKey,
        dataPathEndPoint: dataPathEndPoint,
        name: name,
        queryParams: queryParams,
    }
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithAccessControlReplication(accessControlReplication Ds3TargetAccessControlReplication) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.accessControlReplication = accessControlReplication
    registerDs3TargetSpectraS3Request.queryParams.Set("access_control_replication", accessControlReplication.String())
    return registerDs3TargetSpectraS3Request
}
func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.dataPathHttps = dataPathHttps
    registerDs3TargetSpectraS3Request.queryParams.Set("data_path_https", strconv.FormatBool(dataPathHttps))
    return registerDs3TargetSpectraS3Request
}
func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.dataPathVerifyCertificate = dataPathVerifyCertificate
    registerDs3TargetSpectraS3Request.queryParams.Set("data_path_verify_certificate", strconv.FormatBool(dataPathVerifyCertificate))
    return registerDs3TargetSpectraS3Request
}
func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.defaultReadPreference = defaultReadPreference
    registerDs3TargetSpectraS3Request.queryParams.Set("default_read_preference", defaultReadPreference.String())
    return registerDs3TargetSpectraS3Request
}
func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.permitGoingOutOfSync = permitGoingOutOfSync
    registerDs3TargetSpectraS3Request.queryParams.Set("permit_going_out_of_sync", strconv.FormatBool(permitGoingOutOfSync))
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathPort(dataPathPort *int) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.dataPathPort = dataPathPort
    if dataPathPort != nil {
        registerDs3TargetSpectraS3Request.queryParams.Set("data_path_port", strconv.Itoa(*dataPathPort))
    } else {
        registerDs3TargetSpectraS3Request.queryParams.Set("data_path_port", "")
    }
    return registerDs3TargetSpectraS3Request
}
func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathProxy(dataPathProxy *string) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.dataPathProxy = dataPathProxy
    if dataPathProxy != nil {
        registerDs3TargetSpectraS3Request.queryParams.Set("data_path_proxy", *dataPathProxy)
    } else {
        registerDs3TargetSpectraS3Request.queryParams.Set("data_path_proxy", "")
    }
    return registerDs3TargetSpectraS3Request
}
func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithReplicatedUserDefaultDataPolicy(replicatedUserDefaultDataPolicy *string) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.replicatedUserDefaultDataPolicy = replicatedUserDefaultDataPolicy
    if replicatedUserDefaultDataPolicy != nil {
        registerDs3TargetSpectraS3Request.queryParams.Set("replicated_user_default_data_policy", *replicatedUserDefaultDataPolicy)
    } else {
        registerDs3TargetSpectraS3Request.queryParams.Set("replicated_user_default_data_policy", "")
    }
    return registerDs3TargetSpectraS3Request
}


func (RegisterDs3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) Path() string {
    return "/_rest_/ds3_target"
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) QueryParams() *url.Values {
    return registerDs3TargetSpectraS3Request.queryParams
}

func (RegisterDs3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (RegisterDs3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (RegisterDs3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
