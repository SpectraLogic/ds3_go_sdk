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

type ModifyS3TargetSpectraS3Request struct {
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
    quiesced Quiesced
    region S3Region
    s3Target string
    secretKey *string
    stagedDataExpirationInDays int
    queryParams *url.Values
}

func NewModifyS3TargetSpectraS3Request(s3Target string) *ModifyS3TargetSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyS3TargetSpectraS3Request{
        s3Target: s3Target,
        queryParams: queryParams,
    }
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.defaultReadPreference = defaultReadPreference
    modifyS3TargetSpectraS3Request.queryParams.Set("default_read_preference", defaultReadPreference.String())
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithHttps(https bool) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.https = https
    modifyS3TargetSpectraS3Request.queryParams.Set("https", strconv.FormatBool(https))
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithOfflineDataStagingWindowInTb(offlineDataStagingWindowInTb int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.offlineDataStagingWindowInTb = offlineDataStagingWindowInTb
    modifyS3TargetSpectraS3Request.queryParams.Set("offline_data_staging_window_in_tb", strconv.Itoa(offlineDataStagingWindowInTb))
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.permitGoingOutOfSync = permitGoingOutOfSync
    modifyS3TargetSpectraS3Request.queryParams.Set("permit_going_out_of_sync", strconv.FormatBool(permitGoingOutOfSync))
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.quiesced = quiesced
    modifyS3TargetSpectraS3Request.queryParams.Set("quiesced", quiesced.String())
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithRegion(region S3Region) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.region = region
    modifyS3TargetSpectraS3Request.queryParams.Set("region", region.String())
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithStagedDataExpirationInDays(stagedDataExpirationInDays int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.stagedDataExpirationInDays = stagedDataExpirationInDays
    modifyS3TargetSpectraS3Request.queryParams.Set("staged_data_expiration_in_days", strconv.Itoa(stagedDataExpirationInDays))
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithAccessKey(accessKey *string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.accessKey = accessKey
    if accessKey != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("access_key", *accessKey)
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("access_key", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays *int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.autoVerifyFrequencyInDays = autoVerifyFrequencyInDays
    if autoVerifyFrequencyInDays != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("auto_verify_frequency_in_days", strconv.Itoa(*autoVerifyFrequencyInDays))
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("auto_verify_frequency_in_days", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix *string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.cloudBucketPrefix = cloudBucketPrefix
    if cloudBucketPrefix != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("cloud_bucket_prefix", *cloudBucketPrefix)
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("cloud_bucket_prefix", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix *string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.cloudBucketSuffix = cloudBucketSuffix
    if cloudBucketSuffix != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("cloud_bucket_suffix", *cloudBucketSuffix)
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("cloud_bucket_suffix", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint *string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.dataPathEndPoint = dataPathEndPoint
    if dataPathEndPoint != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("data_path_end_point", *dataPathEndPoint)
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("data_path_end_point", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithName(name *string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.name = name
    if name != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("name", *name)
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("name", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyDomain(proxyDomain *string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.proxyDomain = proxyDomain
    if proxyDomain != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("proxy_domain", *proxyDomain)
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("proxy_domain", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyHost(proxyHost *string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.proxyHost = proxyHost
    if proxyHost != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("proxy_host", *proxyHost)
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("proxy_host", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyPassword(proxyPassword *string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.proxyPassword = proxyPassword
    if proxyPassword != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("proxy_password", *proxyPassword)
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("proxy_password", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyPort(proxyPort *int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.proxyPort = proxyPort
    if proxyPort != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("proxy_port", strconv.Itoa(*proxyPort))
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("proxy_port", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyUsername(proxyUsername *string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.proxyUsername = proxyUsername
    if proxyUsername != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("proxy_username", *proxyUsername)
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("proxy_username", "")
    }
    return modifyS3TargetSpectraS3Request
}
func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithSecretKey(secretKey *string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.secretKey = secretKey
    if secretKey != nil {
        modifyS3TargetSpectraS3Request.queryParams.Set("secret_key", *secretKey)
    } else {
        modifyS3TargetSpectraS3Request.queryParams.Set("secret_key", "")
    }
    return modifyS3TargetSpectraS3Request
}


func (ModifyS3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) Path() string {
    return "/_rest_/s3_target/" + modifyS3TargetSpectraS3Request.s3Target
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) QueryParams() *url.Values {
    return modifyS3TargetSpectraS3Request.queryParams
}

func (ModifyS3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyS3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyS3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
