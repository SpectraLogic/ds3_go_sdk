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

type ModifyAzureTargetSpectraS3Request struct {
    accountKey *string
    accountName *string
    autoVerifyFrequencyInDays *int
    azureTarget string
    cloudBucketPrefix *string
    cloudBucketSuffix *string
    defaultReadPreference TargetReadPreferenceType
    https bool
    name *string
    permitGoingOutOfSync bool
    quiesced Quiesced
    queryParams *url.Values
}

func NewModifyAzureTargetSpectraS3Request(azureTarget string) *ModifyAzureTargetSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyAzureTargetSpectraS3Request{
        azureTarget: azureTarget,
        queryParams: queryParams,
    }
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.defaultReadPreference = defaultReadPreference
    modifyAzureTargetSpectraS3Request.queryParams.Set("default_read_preference", defaultReadPreference.String())
    return modifyAzureTargetSpectraS3Request
}
func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithHttps(https bool) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.https = https
    modifyAzureTargetSpectraS3Request.queryParams.Set("https", strconv.FormatBool(https))
    return modifyAzureTargetSpectraS3Request
}
func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.permitGoingOutOfSync = permitGoingOutOfSync
    modifyAzureTargetSpectraS3Request.queryParams.Set("permit_going_out_of_sync", strconv.FormatBool(permitGoingOutOfSync))
    return modifyAzureTargetSpectraS3Request
}
func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.quiesced = quiesced
    modifyAzureTargetSpectraS3Request.queryParams.Set("quiesced", quiesced.String())
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithAccountKey(accountKey *string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.accountKey = accountKey
    if accountKey != nil {
        modifyAzureTargetSpectraS3Request.queryParams.Set("account_key", *accountKey)
    } else {
        modifyAzureTargetSpectraS3Request.queryParams.Set("account_key", "")
    }
    return modifyAzureTargetSpectraS3Request
}
func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithAccountName(accountName *string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.accountName = accountName
    if accountName != nil {
        modifyAzureTargetSpectraS3Request.queryParams.Set("account_name", *accountName)
    } else {
        modifyAzureTargetSpectraS3Request.queryParams.Set("account_name", "")
    }
    return modifyAzureTargetSpectraS3Request
}
func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays *int) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.autoVerifyFrequencyInDays = autoVerifyFrequencyInDays
    if autoVerifyFrequencyInDays != nil {
        modifyAzureTargetSpectraS3Request.queryParams.Set("auto_verify_frequency_in_days", strconv.Itoa(*autoVerifyFrequencyInDays))
    } else {
        modifyAzureTargetSpectraS3Request.queryParams.Set("auto_verify_frequency_in_days", "")
    }
    return modifyAzureTargetSpectraS3Request
}
func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix *string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.cloudBucketPrefix = cloudBucketPrefix
    if cloudBucketPrefix != nil {
        modifyAzureTargetSpectraS3Request.queryParams.Set("cloud_bucket_prefix", *cloudBucketPrefix)
    } else {
        modifyAzureTargetSpectraS3Request.queryParams.Set("cloud_bucket_prefix", "")
    }
    return modifyAzureTargetSpectraS3Request
}
func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix *string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.cloudBucketSuffix = cloudBucketSuffix
    if cloudBucketSuffix != nil {
        modifyAzureTargetSpectraS3Request.queryParams.Set("cloud_bucket_suffix", *cloudBucketSuffix)
    } else {
        modifyAzureTargetSpectraS3Request.queryParams.Set("cloud_bucket_suffix", "")
    }
    return modifyAzureTargetSpectraS3Request
}
func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithName(name *string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.name = name
    if name != nil {
        modifyAzureTargetSpectraS3Request.queryParams.Set("name", *name)
    } else {
        modifyAzureTargetSpectraS3Request.queryParams.Set("name", "")
    }
    return modifyAzureTargetSpectraS3Request
}


func (ModifyAzureTargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) Path() string {
    return "/_rest_/azure_target/" + modifyAzureTargetSpectraS3Request.azureTarget
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) QueryParams() *url.Values {
    return modifyAzureTargetSpectraS3Request.queryParams
}

func (ModifyAzureTargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyAzureTargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyAzureTargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
