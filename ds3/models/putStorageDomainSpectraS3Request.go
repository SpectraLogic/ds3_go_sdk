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

type PutStorageDomainSpectraS3Request struct {
    AutoEjectMediaFullThreshold *int64
    AutoEjectUponCron *string
    AutoEjectUponJobCancellation *bool
    AutoEjectUponJobCompletion *bool
    AutoEjectUponMediaFull *bool
    LtfsFileNaming LtfsFileNamingMode
    MaximumAutoVerificationFrequencyInDays *int
    MaxTapeFragmentationPercent *int
    MediaEjectionAllowed *bool
    Name string
    SecureMediaAllocation *bool
    VerifyPriorToAutoEject Priority
    WriteOptimization WriteOptimization
}

func NewPutStorageDomainSpectraS3Request(name string) *PutStorageDomainSpectraS3Request {
    return &PutStorageDomainSpectraS3Request{
        Name: name,
    }
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectMediaFullThreshold(autoEjectMediaFullThreshold int64) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.AutoEjectMediaFullThreshold = &autoEjectMediaFullThreshold
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponCron(autoEjectUponCron string) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.AutoEjectUponCron = &autoEjectUponCron
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponJobCancellation(autoEjectUponJobCancellation bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.AutoEjectUponJobCancellation = &autoEjectUponJobCancellation
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponJobCompletion(autoEjectUponJobCompletion bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.AutoEjectUponJobCompletion = &autoEjectUponJobCompletion
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponMediaFull(autoEjectUponMediaFull bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.AutoEjectUponMediaFull = &autoEjectUponMediaFull
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithLtfsFileNaming(ltfsFileNaming LtfsFileNamingMode) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.LtfsFileNaming = ltfsFileNaming
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithMaxTapeFragmentationPercent(maxTapeFragmentationPercent int) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.MaxTapeFragmentationPercent = &maxTapeFragmentationPercent
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithMaximumAutoVerificationFrequencyInDays(maximumAutoVerificationFrequencyInDays int) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.MaximumAutoVerificationFrequencyInDays = &maximumAutoVerificationFrequencyInDays
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithMediaEjectionAllowed(mediaEjectionAllowed bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.MediaEjectionAllowed = &mediaEjectionAllowed
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithSecureMediaAllocation(secureMediaAllocation bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.SecureMediaAllocation = &secureMediaAllocation
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithVerifyPriorToAutoEject(verifyPriorToAutoEject Priority) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.VerifyPriorToAutoEject = verifyPriorToAutoEject
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithWriteOptimization(writeOptimization WriteOptimization) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.WriteOptimization = writeOptimization
    return putStorageDomainSpectraS3Request
}

