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

type ModifyStorageDomainSpectraS3Request struct {
    AutoEjectMediaFullThreshold *int64
    AutoEjectUponCron *string
    AutoEjectUponJobCancellation *bool
    AutoEjectUponJobCompletion *bool
    AutoEjectUponMediaFull *bool
    LtfsFileNaming LtfsFileNamingMode
    MaximumAutoVerificationFrequencyInDays *int
    MaxTapeFragmentationPercent *int
    MediaEjectionAllowed *bool
    Name *string
    SecureMediaAllocation *bool
    StorageDomain string
    VerifyPriorToAutoEject Priority
    WriteOptimization WriteOptimization
}

func NewModifyStorageDomainSpectraS3Request(storageDomain string) *ModifyStorageDomainSpectraS3Request {
    return &ModifyStorageDomainSpectraS3Request{
        StorageDomain: storageDomain,
    }
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectMediaFullThreshold(autoEjectMediaFullThreshold int64) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.AutoEjectMediaFullThreshold = &autoEjectMediaFullThreshold
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponCron(autoEjectUponCron string) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.AutoEjectUponCron = &autoEjectUponCron
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponJobCancellation(autoEjectUponJobCancellation bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.AutoEjectUponJobCancellation = &autoEjectUponJobCancellation
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponJobCompletion(autoEjectUponJobCompletion bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.AutoEjectUponJobCompletion = &autoEjectUponJobCompletion
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponMediaFull(autoEjectUponMediaFull bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.AutoEjectUponMediaFull = &autoEjectUponMediaFull
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithLtfsFileNaming(ltfsFileNaming LtfsFileNamingMode) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.LtfsFileNaming = ltfsFileNaming
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithMaxTapeFragmentationPercent(maxTapeFragmentationPercent int) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.MaxTapeFragmentationPercent = &maxTapeFragmentationPercent
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithMaximumAutoVerificationFrequencyInDays(maximumAutoVerificationFrequencyInDays int) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.MaximumAutoVerificationFrequencyInDays = &maximumAutoVerificationFrequencyInDays
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithMediaEjectionAllowed(mediaEjectionAllowed bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.MediaEjectionAllowed = &mediaEjectionAllowed
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithName(name string) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.Name = &name
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithSecureMediaAllocation(secureMediaAllocation bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.SecureMediaAllocation = &secureMediaAllocation
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithVerifyPriorToAutoEject(verifyPriorToAutoEject Priority) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.VerifyPriorToAutoEject = verifyPriorToAutoEject
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithWriteOptimization(writeOptimization WriteOptimization) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.WriteOptimization = writeOptimization
    return modifyStorageDomainSpectraS3Request
}

