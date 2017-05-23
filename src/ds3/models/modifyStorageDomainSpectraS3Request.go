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

type ModifyStorageDomainSpectraS3Request struct {
    autoEjectMediaFullThreshold *int64
    autoEjectUponCron *string
    autoEjectUponJobCancellation bool
    autoEjectUponJobCompletion bool
    autoEjectUponMediaFull bool
    ltfsFileNaming LtfsFileNamingMode
    maximumAutoVerificationFrequencyInDays *int
    maxTapeFragmentationPercent int
    mediaEjectionAllowed bool
    name *string
    secureMediaAllocation bool
    storageDomain string
    verifyPriorToAutoEject Priority
    writeOptimization WriteOptimization
    queryParams *url.Values
}

func NewModifyStorageDomainSpectraS3Request(storageDomain string) *ModifyStorageDomainSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyStorageDomainSpectraS3Request{
        storageDomain: storageDomain,
        queryParams: queryParams,
    }
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponJobCancellation(autoEjectUponJobCancellation bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.autoEjectUponJobCancellation = autoEjectUponJobCancellation
    modifyStorageDomainSpectraS3Request.queryParams.Set("auto_eject_upon_job_cancellation", strconv.FormatBool(autoEjectUponJobCancellation))
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponJobCompletion(autoEjectUponJobCompletion bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.autoEjectUponJobCompletion = autoEjectUponJobCompletion
    modifyStorageDomainSpectraS3Request.queryParams.Set("auto_eject_upon_job_completion", strconv.FormatBool(autoEjectUponJobCompletion))
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponMediaFull(autoEjectUponMediaFull bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.autoEjectUponMediaFull = autoEjectUponMediaFull
    modifyStorageDomainSpectraS3Request.queryParams.Set("auto_eject_upon_media_full", strconv.FormatBool(autoEjectUponMediaFull))
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithLtfsFileNaming(ltfsFileNaming LtfsFileNamingMode) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.ltfsFileNaming = ltfsFileNaming
    modifyStorageDomainSpectraS3Request.queryParams.Set("ltfs_file_naming", ltfsFileNaming.String())
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithMaxTapeFragmentationPercent(maxTapeFragmentationPercent int) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.maxTapeFragmentationPercent = maxTapeFragmentationPercent
    modifyStorageDomainSpectraS3Request.queryParams.Set("max_tape_fragmentation_percent", strconv.Itoa(maxTapeFragmentationPercent))
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithMediaEjectionAllowed(mediaEjectionAllowed bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.mediaEjectionAllowed = mediaEjectionAllowed
    modifyStorageDomainSpectraS3Request.queryParams.Set("media_ejection_allowed", strconv.FormatBool(mediaEjectionAllowed))
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithSecureMediaAllocation(secureMediaAllocation bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.secureMediaAllocation = secureMediaAllocation
    modifyStorageDomainSpectraS3Request.queryParams.Set("secure_media_allocation", strconv.FormatBool(secureMediaAllocation))
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithVerifyPriorToAutoEject(verifyPriorToAutoEject Priority) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.verifyPriorToAutoEject = verifyPriorToAutoEject
    modifyStorageDomainSpectraS3Request.queryParams.Set("verify_prior_to_auto_eject", verifyPriorToAutoEject.String())
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithWriteOptimization(writeOptimization WriteOptimization) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.writeOptimization = writeOptimization
    modifyStorageDomainSpectraS3Request.queryParams.Set("write_optimization", writeOptimization.String())
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectMediaFullThreshold(autoEjectMediaFullThreshold *int64) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.autoEjectMediaFullThreshold = autoEjectMediaFullThreshold
    if autoEjectMediaFullThreshold != nil {
        modifyStorageDomainSpectraS3Request.queryParams.Set("auto_eject_media_full_threshold", strconv.FormatInt(*autoEjectMediaFullThreshold, 10))
    } else {
        modifyStorageDomainSpectraS3Request.queryParams.Set("auto_eject_media_full_threshold", "")
    }
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponCron(autoEjectUponCron *string) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.autoEjectUponCron = autoEjectUponCron
    if autoEjectUponCron != nil {
        modifyStorageDomainSpectraS3Request.queryParams.Set("auto_eject_upon_cron", *autoEjectUponCron)
    } else {
        modifyStorageDomainSpectraS3Request.queryParams.Set("auto_eject_upon_cron", "")
    }
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithMaximumAutoVerificationFrequencyInDays(maximumAutoVerificationFrequencyInDays *int) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.maximumAutoVerificationFrequencyInDays = maximumAutoVerificationFrequencyInDays
    if maximumAutoVerificationFrequencyInDays != nil {
        modifyStorageDomainSpectraS3Request.queryParams.Set("maximum_auto_verification_frequency_in_days", strconv.Itoa(*maximumAutoVerificationFrequencyInDays))
    } else {
        modifyStorageDomainSpectraS3Request.queryParams.Set("maximum_auto_verification_frequency_in_days", "")
    }
    return modifyStorageDomainSpectraS3Request
}
func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithName(name *string) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.name = name
    if name != nil {
        modifyStorageDomainSpectraS3Request.queryParams.Set("name", *name)
    } else {
        modifyStorageDomainSpectraS3Request.queryParams.Set("name", "")
    }
    return modifyStorageDomainSpectraS3Request
}


func (ModifyStorageDomainSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) Path() string {
    return "/_rest_/storage_domain/" + modifyStorageDomainSpectraS3Request.storageDomain
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) QueryParams() *url.Values {
    return modifyStorageDomainSpectraS3Request.queryParams
}

func (ModifyStorageDomainSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyStorageDomainSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyStorageDomainSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
