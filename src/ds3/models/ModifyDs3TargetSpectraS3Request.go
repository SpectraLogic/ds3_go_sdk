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

type ModifyDs3TargetSpectraS3Request struct {
    accessControlReplication Ds3TargetAccessControlReplication
    adminAuthId *string
    adminSecretKey *string
    dataPathEndPoint *string
    dataPathHttps bool
    dataPathPort *int
    dataPathProxy *string
    dataPathVerifyCertificate bool
    defaultReadPreference TargetReadPreferenceType
    ds3Target string
    name *string
    permitGoingOutOfSync bool
    quiesced Quiesced
    replicatedUserDefaultDataPolicy *string
    queryParams *url.Values
}

func NewModifyDs3TargetSpectraS3Request(ds3Target string) *ModifyDs3TargetSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyDs3TargetSpectraS3Request{
        ds3Target: ds3Target,
        queryParams: queryParams,
    }
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithAccessControlReplication(accessControlReplication Ds3TargetAccessControlReplication) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.accessControlReplication = accessControlReplication
    modifyDs3TargetSpectraS3Request.queryParams.Set("access_control_replication", accessControlReplication.String())
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.dataPathHttps = dataPathHttps
    modifyDs3TargetSpectraS3Request.queryParams.Set("data_path_https", strconv.FormatBool(dataPathHttps))
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.dataPathVerifyCertificate = dataPathVerifyCertificate
    modifyDs3TargetSpectraS3Request.queryParams.Set("data_path_verify_certificate", strconv.FormatBool(dataPathVerifyCertificate))
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.defaultReadPreference = defaultReadPreference
    modifyDs3TargetSpectraS3Request.queryParams.Set("default_read_preference", defaultReadPreference.String())
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.permitGoingOutOfSync = permitGoingOutOfSync
    modifyDs3TargetSpectraS3Request.queryParams.Set("permit_going_out_of_sync", strconv.FormatBool(permitGoingOutOfSync))
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.quiesced = quiesced
    modifyDs3TargetSpectraS3Request.queryParams.Set("quiesced", quiesced.String())
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithAdminAuthId(adminAuthId *string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.adminAuthId = adminAuthId
    if adminAuthId != nil {
        modifyDs3TargetSpectraS3Request.queryParams.Set("admin_auth_id", *adminAuthId)
    } else {
        modifyDs3TargetSpectraS3Request.queryParams.Set("admin_auth_id", "")
    }
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithAdminSecretKey(adminSecretKey *string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.adminSecretKey = adminSecretKey
    if adminSecretKey != nil {
        modifyDs3TargetSpectraS3Request.queryParams.Set("admin_secret_key", *adminSecretKey)
    } else {
        modifyDs3TargetSpectraS3Request.queryParams.Set("admin_secret_key", "")
    }
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint *string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.dataPathEndPoint = dataPathEndPoint
    if dataPathEndPoint != nil {
        modifyDs3TargetSpectraS3Request.queryParams.Set("data_path_end_point", *dataPathEndPoint)
    } else {
        modifyDs3TargetSpectraS3Request.queryParams.Set("data_path_end_point", "")
    }
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathPort(dataPathPort *int) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.dataPathPort = dataPathPort
    if dataPathPort != nil {
        modifyDs3TargetSpectraS3Request.queryParams.Set("data_path_port", strconv.Itoa(*dataPathPort))
    } else {
        modifyDs3TargetSpectraS3Request.queryParams.Set("data_path_port", "")
    }
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathProxy(dataPathProxy *string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.dataPathProxy = dataPathProxy
    if dataPathProxy != nil {
        modifyDs3TargetSpectraS3Request.queryParams.Set("data_path_proxy", *dataPathProxy)
    } else {
        modifyDs3TargetSpectraS3Request.queryParams.Set("data_path_proxy", "")
    }
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithName(name *string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.name = name
    if name != nil {
        modifyDs3TargetSpectraS3Request.queryParams.Set("name", *name)
    } else {
        modifyDs3TargetSpectraS3Request.queryParams.Set("name", "")
    }
    return modifyDs3TargetSpectraS3Request
}
func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithReplicatedUserDefaultDataPolicy(replicatedUserDefaultDataPolicy *string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.replicatedUserDefaultDataPolicy = replicatedUserDefaultDataPolicy
    if replicatedUserDefaultDataPolicy != nil {
        modifyDs3TargetSpectraS3Request.queryParams.Set("replicated_user_default_data_policy", *replicatedUserDefaultDataPolicy)
    } else {
        modifyDs3TargetSpectraS3Request.queryParams.Set("replicated_user_default_data_policy", "")
    }
    return modifyDs3TargetSpectraS3Request
}


func (ModifyDs3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) Path() string {
    return "/_rest_/ds3_target/" + modifyDs3TargetSpectraS3Request.ds3Target
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) QueryParams() *url.Values {
    return modifyDs3TargetSpectraS3Request.queryParams
}

func (ModifyDs3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyDs3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyDs3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
