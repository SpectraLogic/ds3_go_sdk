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

type GetStorageDomainsSpectraS3Request struct {
    autoEjectUponCron *string
    autoEjectUponJobCancellation bool
    autoEjectUponJobCompletion bool
    autoEjectUponMediaFull bool
    mediaEjectionAllowed bool
    name *string
    pageLength int
    pageOffset int
    pageStartMarker string
    secureMediaAllocation bool
    writeOptimization WriteOptimization
    queryParams *url.Values
}

func NewGetStorageDomainsSpectraS3Request() *GetStorageDomainsSpectraS3Request {
    queryParams := &url.Values{}

    return &GetStorageDomainsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponJobCancellation(autoEjectUponJobCancellation bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.autoEjectUponJobCancellation = autoEjectUponJobCancellation
    getStorageDomainsSpectraS3Request.queryParams.Set("auto_eject_upon_job_cancellation", strconv.FormatBool(autoEjectUponJobCancellation))
    return getStorageDomainsSpectraS3Request
}
func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponJobCompletion(autoEjectUponJobCompletion bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.autoEjectUponJobCompletion = autoEjectUponJobCompletion
    getStorageDomainsSpectraS3Request.queryParams.Set("auto_eject_upon_job_completion", strconv.FormatBool(autoEjectUponJobCompletion))
    return getStorageDomainsSpectraS3Request
}
func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponMediaFull(autoEjectUponMediaFull bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.autoEjectUponMediaFull = autoEjectUponMediaFull
    getStorageDomainsSpectraS3Request.queryParams.Set("auto_eject_upon_media_full", strconv.FormatBool(autoEjectUponMediaFull))
    return getStorageDomainsSpectraS3Request
}
func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithMediaEjectionAllowed(mediaEjectionAllowed bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.mediaEjectionAllowed = mediaEjectionAllowed
    getStorageDomainsSpectraS3Request.queryParams.Set("media_ejection_allowed", strconv.FormatBool(mediaEjectionAllowed))
    return getStorageDomainsSpectraS3Request
}
func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.pageLength = pageLength
    getStorageDomainsSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getStorageDomainsSpectraS3Request
}
func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.pageOffset = pageOffset
    getStorageDomainsSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getStorageDomainsSpectraS3Request
}
func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.pageStartMarker = pageStartMarker
    getStorageDomainsSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getStorageDomainsSpectraS3Request
}
func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithSecureMediaAllocation(secureMediaAllocation bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.secureMediaAllocation = secureMediaAllocation
    getStorageDomainsSpectraS3Request.queryParams.Set("secure_media_allocation", strconv.FormatBool(secureMediaAllocation))
    return getStorageDomainsSpectraS3Request
}
func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithWriteOptimization(writeOptimization WriteOptimization) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.writeOptimization = writeOptimization
    getStorageDomainsSpectraS3Request.queryParams.Set("write_optimization", writeOptimization.String())
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponCron(autoEjectUponCron *string) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.autoEjectUponCron = autoEjectUponCron
    if autoEjectUponCron != nil {
        getStorageDomainsSpectraS3Request.queryParams.Set("auto_eject_upon_cron", *autoEjectUponCron)
    } else {
        getStorageDomainsSpectraS3Request.queryParams.Set("auto_eject_upon_cron", "")
    }
    return getStorageDomainsSpectraS3Request
}
func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithName(name *string) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.name = name
    if name != nil {
        getStorageDomainsSpectraS3Request.queryParams.Set("name", *name)
    } else {
        getStorageDomainsSpectraS3Request.queryParams.Set("name", "")
    }
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithLastPage() *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.queryParams.Set("last_page", "")
    return getStorageDomainsSpectraS3Request
}

func (GetStorageDomainsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) Path() string {
    return "/_rest_/storage_domain"
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) QueryParams() *url.Values {
    return getStorageDomainsSpectraS3Request.queryParams
}

func (GetStorageDomainsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetStorageDomainsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetStorageDomainsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
