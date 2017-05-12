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

type PairBackRegisteredDs3TargetSpectraS3Request struct {
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
    replicatedUserDefaultDataPolicy *string
    queryParams *url.Values
}

func NewPairBackRegisteredDs3TargetSpectraS3Request(ds3Target string) *PairBackRegisteredDs3TargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "pair_back")

    return &PairBackRegisteredDs3TargetSpectraS3Request{
        ds3Target: ds3Target,
        queryParams: queryParams,
    }
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithAccessControlReplication(accessControlReplication Ds3TargetAccessControlReplication) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.accessControlReplication = accessControlReplication
    pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("access_control_replication", accessControlReplication.String())
    return pairBackRegisteredDs3TargetSpectraS3Request
}
func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.dataPathHttps = dataPathHttps
    pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("data_path_https", strconv.FormatBool(dataPathHttps))
    return pairBackRegisteredDs3TargetSpectraS3Request
}
func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.dataPathVerifyCertificate = dataPathVerifyCertificate
    pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("data_path_verify_certificate", strconv.FormatBool(dataPathVerifyCertificate))
    return pairBackRegisteredDs3TargetSpectraS3Request
}
func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.defaultReadPreference = defaultReadPreference
    pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("default_read_preference", defaultReadPreference.String())
    return pairBackRegisteredDs3TargetSpectraS3Request
}
func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.permitGoingOutOfSync = permitGoingOutOfSync
    pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("permit_going_out_of_sync", strconv.FormatBool(permitGoingOutOfSync))
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithAdminAuthId(adminAuthId *string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.adminAuthId = adminAuthId
    if adminAuthId != nil {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("admin_auth_id", *adminAuthId)
    } else {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("admin_auth_id", "")
    }
    return pairBackRegisteredDs3TargetSpectraS3Request
}
func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithAdminSecretKey(adminSecretKey *string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.adminSecretKey = adminSecretKey
    if adminSecretKey != nil {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("admin_secret_key", *adminSecretKey)
    } else {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("admin_secret_key", "")
    }
    return pairBackRegisteredDs3TargetSpectraS3Request
}
func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint *string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.dataPathEndPoint = dataPathEndPoint
    if dataPathEndPoint != nil {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("data_path_end_point", *dataPathEndPoint)
    } else {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("data_path_end_point", "")
    }
    return pairBackRegisteredDs3TargetSpectraS3Request
}
func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathPort(dataPathPort *int) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.dataPathPort = dataPathPort
    if dataPathPort != nil {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("data_path_port", strconv.Itoa(*dataPathPort))
    } else {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("data_path_port", "")
    }
    return pairBackRegisteredDs3TargetSpectraS3Request
}
func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathProxy(dataPathProxy *string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.dataPathProxy = dataPathProxy
    if dataPathProxy != nil {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("data_path_proxy", *dataPathProxy)
    } else {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("data_path_proxy", "")
    }
    return pairBackRegisteredDs3TargetSpectraS3Request
}
func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithName(name *string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.name = name
    if name != nil {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("name", *name)
    } else {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("name", "")
    }
    return pairBackRegisteredDs3TargetSpectraS3Request
}
func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithReplicatedUserDefaultDataPolicy(replicatedUserDefaultDataPolicy *string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.replicatedUserDefaultDataPolicy = replicatedUserDefaultDataPolicy
    if replicatedUserDefaultDataPolicy != nil {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("replicated_user_default_data_policy", *replicatedUserDefaultDataPolicy)
    } else {
        pairBackRegisteredDs3TargetSpectraS3Request.queryParams.Set("replicated_user_default_data_policy", "")
    }
    return pairBackRegisteredDs3TargetSpectraS3Request
}


func (PairBackRegisteredDs3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) Path() string {
    return "/_rest_/ds3_target/" + pairBackRegisteredDs3TargetSpectraS3Request.ds3Target
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) QueryParams() *url.Values {
    return pairBackRegisteredDs3TargetSpectraS3Request.queryParams
}

func (PairBackRegisteredDs3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PairBackRegisteredDs3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PairBackRegisteredDs3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
