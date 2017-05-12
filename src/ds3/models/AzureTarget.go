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

type AzureTarget struct {
    AccountKey *string `xml:"AccountKey"`
    AccountName *string `xml:"AccountName"`
    AutoVerifyFrequencyInDays *int `xml:"AutoVerifyFrequencyInDays"`
    CloudBucketPrefix *string `xml:"CloudBucketPrefix"`
    CloudBucketSuffix *string `xml:"CloudBucketSuffix"`
    DefaultReadPreference TargetReadPreferenceType `xml:"DefaultReadPreference"`
    Https bool `xml:"Https"`
    Id string `xml:"Id"`
    LastFullyVerified *string `xml:"LastFullyVerified"`
    Name *string `xml:"Name"`
    PermitGoingOutOfSync bool `xml:"PermitGoingOutOfSync"`
    Quiesced Quiesced `xml:"Quiesced"`
    State TargetState `xml:"State"`
}