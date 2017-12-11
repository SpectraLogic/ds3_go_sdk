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

type ModifyAzureTargetSpectraS3Request struct {
    AccountKey *string
    AccountName *string
    AutoVerifyFrequencyInDays *int
    AzureTarget string
    CloudBucketPrefix *string
    CloudBucketSuffix *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    Name *string
    PermitGoingOutOfSync *bool
    Quiesced Quiesced
}

func NewModifyAzureTargetSpectraS3Request(azureTarget string) *ModifyAzureTargetSpectraS3Request {
    return &ModifyAzureTargetSpectraS3Request{
        AzureTarget: azureTarget,
    }
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithAccountKey(accountKey string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.AccountKey = &accountKey
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithAccountName(accountName string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.AccountName = &accountName
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays int) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.AutoVerifyFrequencyInDays = &autoVerifyFrequencyInDays
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.CloudBucketPrefix = &cloudBucketPrefix
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.CloudBucketSuffix = &cloudBucketSuffix
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithHttps(https bool) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.Https = &https
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithName(name string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.Name = &name
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.Quiesced = quiesced
    return modifyAzureTargetSpectraS3Request
}

