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

type ModifyS3TargetSpectraS3Request struct {
    AccessKey *string
    AutoVerifyFrequencyInDays *int
    CloudBucketPrefix *string
    CloudBucketSuffix *string
    DataPathEndPoint *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    Name *string
    OfflineDataStagingWindowInTb *int
    PermitGoingOutOfSync *bool
    ProxyDomain *string
    ProxyHost *string
    ProxyPassword *string
    ProxyPort *int
    ProxyUsername *string
    Quiesced Quiesced
    Region S3Region
    S3Target string
    SecretKey *string
    StagedDataExpirationInDays *int
}

func NewModifyS3TargetSpectraS3Request(s3Target string) *ModifyS3TargetSpectraS3Request {
    return &ModifyS3TargetSpectraS3Request{
        S3Target: s3Target,
    }
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithAccessKey(accessKey string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.AccessKey = &accessKey
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.AutoVerifyFrequencyInDays = &autoVerifyFrequencyInDays
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.CloudBucketPrefix = &cloudBucketPrefix
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.CloudBucketSuffix = &cloudBucketSuffix
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithHttps(https bool) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.Https = &https
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithName(name string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.Name = &name
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithOfflineDataStagingWindowInTb(offlineDataStagingWindowInTb int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.OfflineDataStagingWindowInTb = &offlineDataStagingWindowInTb
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyDomain(proxyDomain string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.ProxyDomain = &proxyDomain
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyHost(proxyHost string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.ProxyHost = &proxyHost
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyPassword(proxyPassword string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.ProxyPassword = &proxyPassword
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyPort(proxyPort int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.ProxyPort = &proxyPort
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyUsername(proxyUsername string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.ProxyUsername = &proxyUsername
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.Quiesced = quiesced
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithRegion(region S3Region) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.Region = region
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithSecretKey(secretKey string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.SecretKey = &secretKey
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithStagedDataExpirationInDays(stagedDataExpirationInDays int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.StagedDataExpirationInDays = &stagedDataExpirationInDays
    return modifyS3TargetSpectraS3Request
}

