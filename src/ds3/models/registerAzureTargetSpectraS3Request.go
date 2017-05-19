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

type RegisterAzureTargetSpectraS3Request struct {
    accountKey *string
    accountName *string
    autoVerifyFrequencyInDays *int
    cloudBucketPrefix *string
    cloudBucketSuffix *string
    defaultReadPreference TargetReadPreferenceType
    https bool
    name *string
    permitGoingOutOfSync bool
    queryParams *url.Values
}

func NewRegisterAzureTargetSpectraS3Request(accountKey *string, accountName *string, name *string) *RegisterAzureTargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("account_key", *accountKey)
    queryParams.Set("account_name", *accountName)
    queryParams.Set("name", *name)

    return &RegisterAzureTargetSpectraS3Request{
        accountKey: accountKey,
        accountName: accountName,
        name: name,
        queryParams: queryParams,
    }
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.defaultReadPreference = defaultReadPreference
    registerAzureTargetSpectraS3Request.queryParams.Set("default_read_preference", defaultReadPreference.String())
    return registerAzureTargetSpectraS3Request
}
func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithHttps(https bool) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.https = https
    registerAzureTargetSpectraS3Request.queryParams.Set("https", strconv.FormatBool(https))
    return registerAzureTargetSpectraS3Request
}
func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.permitGoingOutOfSync = permitGoingOutOfSync
    registerAzureTargetSpectraS3Request.queryParams.Set("permit_going_out_of_sync", strconv.FormatBool(permitGoingOutOfSync))
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays *int) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.autoVerifyFrequencyInDays = autoVerifyFrequencyInDays
    if autoVerifyFrequencyInDays != nil {
        registerAzureTargetSpectraS3Request.queryParams.Set("auto_verify_frequency_in_days", strconv.Itoa(*autoVerifyFrequencyInDays))
    } else {
        registerAzureTargetSpectraS3Request.queryParams.Set("auto_verify_frequency_in_days", "")
    }
    return registerAzureTargetSpectraS3Request
}
func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix *string) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.cloudBucketPrefix = cloudBucketPrefix
    if cloudBucketPrefix != nil {
        registerAzureTargetSpectraS3Request.queryParams.Set("cloud_bucket_prefix", *cloudBucketPrefix)
    } else {
        registerAzureTargetSpectraS3Request.queryParams.Set("cloud_bucket_prefix", "")
    }
    return registerAzureTargetSpectraS3Request
}
func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix *string) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.cloudBucketSuffix = cloudBucketSuffix
    if cloudBucketSuffix != nil {
        registerAzureTargetSpectraS3Request.queryParams.Set("cloud_bucket_suffix", *cloudBucketSuffix)
    } else {
        registerAzureTargetSpectraS3Request.queryParams.Set("cloud_bucket_suffix", "")
    }
    return registerAzureTargetSpectraS3Request
}


func (RegisterAzureTargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) Path() string {
    return "/_rest_/azure_target"
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) QueryParams() *url.Values {
    return registerAzureTargetSpectraS3Request.queryParams
}

func (RegisterAzureTargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (RegisterAzureTargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (RegisterAzureTargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
