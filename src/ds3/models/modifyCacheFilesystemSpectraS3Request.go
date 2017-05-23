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

type ModifyCacheFilesystemSpectraS3Request struct {
    autoReclaimInitiateThreshold float64
    autoReclaimTerminateThreshold float64
    burstThreshold float64
    cacheFilesystem string
    maxCapacityInBytes *int64
    queryParams *url.Values
}

func NewModifyCacheFilesystemSpectraS3Request(cacheFilesystem string) *ModifyCacheFilesystemSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyCacheFilesystemSpectraS3Request{
        cacheFilesystem: cacheFilesystem,
        queryParams: queryParams,
    }
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithAutoReclaimInitiateThreshold(autoReclaimInitiateThreshold float64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.autoReclaimInitiateThreshold = autoReclaimInitiateThreshold
    modifyCacheFilesystemSpectraS3Request.queryParams.Set("auto_reclaim_initiate_threshold", strconv.FormatFloat(autoReclaimInitiateThreshold, 'f', -1, 64))
    return modifyCacheFilesystemSpectraS3Request
}
func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithAutoReclaimTerminateThreshold(autoReclaimTerminateThreshold float64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.autoReclaimTerminateThreshold = autoReclaimTerminateThreshold
    modifyCacheFilesystemSpectraS3Request.queryParams.Set("auto_reclaim_terminate_threshold", strconv.FormatFloat(autoReclaimTerminateThreshold, 'f', -1, 64))
    return modifyCacheFilesystemSpectraS3Request
}
func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithBurstThreshold(burstThreshold float64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.burstThreshold = burstThreshold
    modifyCacheFilesystemSpectraS3Request.queryParams.Set("burst_threshold", strconv.FormatFloat(burstThreshold, 'f', -1, 64))
    return modifyCacheFilesystemSpectraS3Request
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithMaxCapacityInBytes(maxCapacityInBytes *int64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.maxCapacityInBytes = maxCapacityInBytes
    if maxCapacityInBytes != nil {
        modifyCacheFilesystemSpectraS3Request.queryParams.Set("max_capacity_in_bytes", strconv.FormatInt(*maxCapacityInBytes, 10))
    } else {
        modifyCacheFilesystemSpectraS3Request.queryParams.Set("max_capacity_in_bytes", "")
    }
    return modifyCacheFilesystemSpectraS3Request
}


func (ModifyCacheFilesystemSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) Path() string {
    return "/_rest_/cache_filesystem/" + modifyCacheFilesystemSpectraS3Request.cacheFilesystem
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) QueryParams() *url.Values {
    return modifyCacheFilesystemSpectraS3Request.queryParams
}

func (ModifyCacheFilesystemSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyCacheFilesystemSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyCacheFilesystemSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
