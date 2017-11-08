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

type RegisterS3TargetSpectraS3Request struct {
    AccessKey string
    AutoVerifyFrequencyInDays *int
    CloudBucketPrefix *string
    CloudBucketSuffix *string
    DataPathEndPoint *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    Name string
    OfflineDataStagingWindowInTb *int
    PermitGoingOutOfSync *bool
    ProxyDomain *string
    ProxyHost *string
    ProxyPassword *string
    ProxyPort *int
    ProxyUsername *string
    Region S3Region
    SecretKey string
    StagedDataExpirationInDays *int
}

func NewRegisterS3TargetSpectraS3Request(accessKey string, name string, secretKey string) *RegisterS3TargetSpectraS3Request {
    return &RegisterS3TargetSpectraS3Request{
        AccessKey: accessKey,
        Name: name,
        SecretKey: secretKey,
    }
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.AutoVerifyFrequencyInDays = &autoVerifyFrequencyInDays
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.CloudBucketPrefix = &cloudBucketPrefix
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.CloudBucketSuffix = &cloudBucketSuffix
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithHttps(https bool) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.Https = &https
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithOfflineDataStagingWindowInTb(offlineDataStagingWindowInTb int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.OfflineDataStagingWindowInTb = &offlineDataStagingWindowInTb
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyDomain(proxyDomain string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.ProxyDomain = &proxyDomain
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyHost(proxyHost string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.ProxyHost = &proxyHost
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyPassword(proxyPassword string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.ProxyPassword = &proxyPassword
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyPort(proxyPort int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.ProxyPort = &proxyPort
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyUsername(proxyUsername string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.ProxyUsername = &proxyUsername
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithRegion(region S3Region) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.Region = region
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithStagedDataExpirationInDays(stagedDataExpirationInDays int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.StagedDataExpirationInDays = &stagedDataExpirationInDays
    return registerS3TargetSpectraS3Request
}

