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

type RegisterAzureTargetSpectraS3Request struct {
    AccountKey string
    AccountName string
    AutoVerifyFrequencyInDays *int
    CloudBucketPrefix *string
    CloudBucketSuffix *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    Name string
    PermitGoingOutOfSync *bool
}

func NewRegisterAzureTargetSpectraS3Request(accountKey string, accountName string, name string) *RegisterAzureTargetSpectraS3Request {
    return &RegisterAzureTargetSpectraS3Request{
        AccountKey: accountKey,
        AccountName: accountName,
        Name: name,
    }
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays int) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.AutoVerifyFrequencyInDays = &autoVerifyFrequencyInDays
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix string) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.CloudBucketPrefix = &cloudBucketPrefix
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix string) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.CloudBucketSuffix = &cloudBucketSuffix
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithHttps(https bool) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.Https = &https
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return registerAzureTargetSpectraS3Request
}

