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

type GetStorageDomainsSpectraS3Request struct {
    AutoEjectUponCron *string
    AutoEjectUponJobCancellation *bool
    AutoEjectUponJobCompletion *bool
    AutoEjectUponMediaFull *bool
    LastPage bool
    MediaEjectionAllowed *bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    SecureMediaAllocation *bool
    WriteOptimization WriteOptimization
}

func NewGetStorageDomainsSpectraS3Request() *GetStorageDomainsSpectraS3Request {
    return &GetStorageDomainsSpectraS3Request{
    }
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponCron(autoEjectUponCron string) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.AutoEjectUponCron = &autoEjectUponCron
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponJobCancellation(autoEjectUponJobCancellation bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.AutoEjectUponJobCancellation = &autoEjectUponJobCancellation
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponJobCompletion(autoEjectUponJobCompletion bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.AutoEjectUponJobCompletion = &autoEjectUponJobCompletion
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponMediaFull(autoEjectUponMediaFull bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.AutoEjectUponMediaFull = &autoEjectUponMediaFull
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithLastPage() *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.LastPage = true
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithMediaEjectionAllowed(mediaEjectionAllowed bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.MediaEjectionAllowed = &mediaEjectionAllowed
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithName(name string) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.Name = &name
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.PageLength = &pageLength
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.PageOffset = &pageOffset
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithSecureMediaAllocation(secureMediaAllocation bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.SecureMediaAllocation = &secureMediaAllocation
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithWriteOptimization(writeOptimization WriteOptimization) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.WriteOptimization = writeOptimization
    return getStorageDomainsSpectraS3Request
}

