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

type PutStorageDomainSpectraS3Request struct {
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
    verifyPriorToAutoEject Priority
    writeOptimization WriteOptimization
    queryParams *url.Values
}

func NewPutStorageDomainSpectraS3Request(name *string) *PutStorageDomainSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("name", *name)

    return &PutStorageDomainSpectraS3Request{
        name: name,
        queryParams: queryParams,
    }
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponJobCancellation(autoEjectUponJobCancellation bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.autoEjectUponJobCancellation = autoEjectUponJobCancellation
    putStorageDomainSpectraS3Request.queryParams.Set("auto_eject_upon_job_cancellation", strconv.FormatBool(autoEjectUponJobCancellation))
    return putStorageDomainSpectraS3Request
}
func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponJobCompletion(autoEjectUponJobCompletion bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.autoEjectUponJobCompletion = autoEjectUponJobCompletion
    putStorageDomainSpectraS3Request.queryParams.Set("auto_eject_upon_job_completion", strconv.FormatBool(autoEjectUponJobCompletion))
    return putStorageDomainSpectraS3Request
}
func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponMediaFull(autoEjectUponMediaFull bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.autoEjectUponMediaFull = autoEjectUponMediaFull
    putStorageDomainSpectraS3Request.queryParams.Set("auto_eject_upon_media_full", strconv.FormatBool(autoEjectUponMediaFull))
    return putStorageDomainSpectraS3Request
}
func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithLtfsFileNaming(ltfsFileNaming LtfsFileNamingMode) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.ltfsFileNaming = ltfsFileNaming
    putStorageDomainSpectraS3Request.queryParams.Set("ltfs_file_naming", ltfsFileNaming.String())
    return putStorageDomainSpectraS3Request
}
func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithMaxTapeFragmentationPercent(maxTapeFragmentationPercent int) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.maxTapeFragmentationPercent = maxTapeFragmentationPercent
    putStorageDomainSpectraS3Request.queryParams.Set("max_tape_fragmentation_percent", strconv.Itoa(maxTapeFragmentationPercent))
    return putStorageDomainSpectraS3Request
}
func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithMediaEjectionAllowed(mediaEjectionAllowed bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.mediaEjectionAllowed = mediaEjectionAllowed
    putStorageDomainSpectraS3Request.queryParams.Set("media_ejection_allowed", strconv.FormatBool(mediaEjectionAllowed))
    return putStorageDomainSpectraS3Request
}
func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithSecureMediaAllocation(secureMediaAllocation bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.secureMediaAllocation = secureMediaAllocation
    putStorageDomainSpectraS3Request.queryParams.Set("secure_media_allocation", strconv.FormatBool(secureMediaAllocation))
    return putStorageDomainSpectraS3Request
}
func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithVerifyPriorToAutoEject(verifyPriorToAutoEject Priority) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.verifyPriorToAutoEject = verifyPriorToAutoEject
    putStorageDomainSpectraS3Request.queryParams.Set("verify_prior_to_auto_eject", verifyPriorToAutoEject.String())
    return putStorageDomainSpectraS3Request
}
func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithWriteOptimization(writeOptimization WriteOptimization) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.writeOptimization = writeOptimization
    putStorageDomainSpectraS3Request.queryParams.Set("write_optimization", writeOptimization.String())
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectMediaFullThreshold(autoEjectMediaFullThreshold *int64) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.autoEjectMediaFullThreshold = autoEjectMediaFullThreshold
    if autoEjectMediaFullThreshold != nil {
        putStorageDomainSpectraS3Request.queryParams.Set("auto_eject_media_full_threshold", strconv.FormatInt(*autoEjectMediaFullThreshold, 10))
    } else {
        putStorageDomainSpectraS3Request.queryParams.Set("auto_eject_media_full_threshold", "")
    }
    return putStorageDomainSpectraS3Request
}
func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponCron(autoEjectUponCron *string) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.autoEjectUponCron = autoEjectUponCron
    if autoEjectUponCron != nil {
        putStorageDomainSpectraS3Request.queryParams.Set("auto_eject_upon_cron", *autoEjectUponCron)
    } else {
        putStorageDomainSpectraS3Request.queryParams.Set("auto_eject_upon_cron", "")
    }
    return putStorageDomainSpectraS3Request
}
func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithMaximumAutoVerificationFrequencyInDays(maximumAutoVerificationFrequencyInDays *int) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.maximumAutoVerificationFrequencyInDays = maximumAutoVerificationFrequencyInDays
    if maximumAutoVerificationFrequencyInDays != nil {
        putStorageDomainSpectraS3Request.queryParams.Set("maximum_auto_verification_frequency_in_days", strconv.Itoa(*maximumAutoVerificationFrequencyInDays))
    } else {
        putStorageDomainSpectraS3Request.queryParams.Set("maximum_auto_verification_frequency_in_days", "")
    }
    return putStorageDomainSpectraS3Request
}


func (PutStorageDomainSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) Path() string {
    return "/_rest_/storage_domain"
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) QueryParams() *url.Values {
    return putStorageDomainSpectraS3Request.queryParams
}

func (PutStorageDomainSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutStorageDomainSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutStorageDomainSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
