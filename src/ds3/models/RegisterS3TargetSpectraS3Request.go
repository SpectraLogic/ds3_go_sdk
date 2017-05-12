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

type RegisterS3TargetSpectraS3Request struct {
    accessKey *string
    autoVerifyFrequencyInDays *int
    cloudBucketPrefix *string
    cloudBucketSuffix *string
    dataPathEndPoint *string
    defaultReadPreference TargetReadPreferenceType
    https bool
    name *string
    offlineDataStagingWindowInTb int
    permitGoingOutOfSync bool
    proxyDomain *string
    proxyHost *string
    proxyPassword *string
    proxyPort *int
    proxyUsername *string
    region S3Region
    secretKey *string
    stagedDataExpirationInDays int
    queryParams *url.Values
}

func NewRegisterS3TargetSpectraS3Request(accessKey *string, name *string, secretKey *string) *RegisterS3TargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("access_key", *accessKey)
    queryParams.Set("name", *name)
    queryParams.Set("secret_key", *secretKey)

    return &RegisterS3TargetSpectraS3Request{
        accessKey: accessKey,
        name: name,
        secretKey: secretKey,
        queryParams: queryParams,
    }
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.defaultReadPreference = defaultReadPreference
    registerS3TargetSpectraS3Request.queryParams.Set("default_read_preference", defaultReadPreference.String())
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithHttps(https bool) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.https = https
    registerS3TargetSpectraS3Request.queryParams.Set("https", strconv.FormatBool(https))
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithOfflineDataStagingWindowInTb(offlineDataStagingWindowInTb int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.offlineDataStagingWindowInTb = offlineDataStagingWindowInTb
    registerS3TargetSpectraS3Request.queryParams.Set("offline_data_staging_window_in_tb", strconv.Itoa(offlineDataStagingWindowInTb))
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.permitGoingOutOfSync = permitGoingOutOfSync
    registerS3TargetSpectraS3Request.queryParams.Set("permit_going_out_of_sync", strconv.FormatBool(permitGoingOutOfSync))
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithRegion(region S3Region) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.region = region
    registerS3TargetSpectraS3Request.queryParams.Set("region", region.String())
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithStagedDataExpirationInDays(stagedDataExpirationInDays int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.stagedDataExpirationInDays = stagedDataExpirationInDays
    registerS3TargetSpectraS3Request.queryParams.Set("staged_data_expiration_in_days", strconv.Itoa(stagedDataExpirationInDays))
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays *int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.autoVerifyFrequencyInDays = autoVerifyFrequencyInDays
    if autoVerifyFrequencyInDays != nil {
        registerS3TargetSpectraS3Request.queryParams.Set("auto_verify_frequency_in_days", strconv.Itoa(*autoVerifyFrequencyInDays))
    } else {
        registerS3TargetSpectraS3Request.queryParams.Set("auto_verify_frequency_in_days", "")
    }
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix *string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.cloudBucketPrefix = cloudBucketPrefix
    if cloudBucketPrefix != nil {
        registerS3TargetSpectraS3Request.queryParams.Set("cloud_bucket_prefix", *cloudBucketPrefix)
    } else {
        registerS3TargetSpectraS3Request.queryParams.Set("cloud_bucket_prefix", "")
    }
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix *string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.cloudBucketSuffix = cloudBucketSuffix
    if cloudBucketSuffix != nil {
        registerS3TargetSpectraS3Request.queryParams.Set("cloud_bucket_suffix", *cloudBucketSuffix)
    } else {
        registerS3TargetSpectraS3Request.queryParams.Set("cloud_bucket_suffix", "")
    }
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint *string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.dataPathEndPoint = dataPathEndPoint
    if dataPathEndPoint != nil {
        registerS3TargetSpectraS3Request.queryParams.Set("data_path_end_point", *dataPathEndPoint)
    } else {
        registerS3TargetSpectraS3Request.queryParams.Set("data_path_end_point", "")
    }
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyDomain(proxyDomain *string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.proxyDomain = proxyDomain
    if proxyDomain != nil {
        registerS3TargetSpectraS3Request.queryParams.Set("proxy_domain", *proxyDomain)
    } else {
        registerS3TargetSpectraS3Request.queryParams.Set("proxy_domain", "")
    }
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyHost(proxyHost *string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.proxyHost = proxyHost
    if proxyHost != nil {
        registerS3TargetSpectraS3Request.queryParams.Set("proxy_host", *proxyHost)
    } else {
        registerS3TargetSpectraS3Request.queryParams.Set("proxy_host", "")
    }
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyPassword(proxyPassword *string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.proxyPassword = proxyPassword
    if proxyPassword != nil {
        registerS3TargetSpectraS3Request.queryParams.Set("proxy_password", *proxyPassword)
    } else {
        registerS3TargetSpectraS3Request.queryParams.Set("proxy_password", "")
    }
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyPort(proxyPort *int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.proxyPort = proxyPort
    if proxyPort != nil {
        registerS3TargetSpectraS3Request.queryParams.Set("proxy_port", strconv.Itoa(*proxyPort))
    } else {
        registerS3TargetSpectraS3Request.queryParams.Set("proxy_port", "")
    }
    return registerS3TargetSpectraS3Request
}
func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyUsername(proxyUsername *string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.proxyUsername = proxyUsername
    if proxyUsername != nil {
        registerS3TargetSpectraS3Request.queryParams.Set("proxy_username", *proxyUsername)
    } else {
        registerS3TargetSpectraS3Request.queryParams.Set("proxy_username", "")
    }
    return registerS3TargetSpectraS3Request
}


func (RegisterS3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) Path() string {
    return "/_rest_/s3_target"
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) QueryParams() *url.Values {
    return registerS3TargetSpectraS3Request.queryParams
}

func (RegisterS3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (RegisterS3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (RegisterS3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
