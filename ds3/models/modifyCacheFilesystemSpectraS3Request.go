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

type ModifyCacheFilesystemSpectraS3Request struct {
    AutoReclaimInitiateThreshold *float64
    AutoReclaimTerminateThreshold *float64
    BurstThreshold *float64
    CacheFilesystem string
    MaxCapacityInBytes *int64
}

func NewModifyCacheFilesystemSpectraS3Request(cacheFilesystem string) *ModifyCacheFilesystemSpectraS3Request {
    return &ModifyCacheFilesystemSpectraS3Request{
        CacheFilesystem: cacheFilesystem,
    }
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithAutoReclaimInitiateThreshold(autoReclaimInitiateThreshold float64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.AutoReclaimInitiateThreshold = &autoReclaimInitiateThreshold
    return modifyCacheFilesystemSpectraS3Request
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithAutoReclaimTerminateThreshold(autoReclaimTerminateThreshold float64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.AutoReclaimTerminateThreshold = &autoReclaimTerminateThreshold
    return modifyCacheFilesystemSpectraS3Request
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithBurstThreshold(burstThreshold float64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.BurstThreshold = &burstThreshold
    return modifyCacheFilesystemSpectraS3Request
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithMaxCapacityInBytes(maxCapacityInBytes int64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.MaxCapacityInBytes = &maxCapacityInBytes
    return modifyCacheFilesystemSpectraS3Request
}

