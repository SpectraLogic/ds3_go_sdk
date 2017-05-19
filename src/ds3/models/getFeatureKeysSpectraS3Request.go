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

type GetFeatureKeysSpectraS3Request struct {
    errorMessage *string
    expirationDate string
    key FeatureKeyType
    pageLength int
    pageOffset int
    pageStartMarker string
    queryParams *url.Values
}

func NewGetFeatureKeysSpectraS3Request() *GetFeatureKeysSpectraS3Request {
    queryParams := &url.Values{}

    return &GetFeatureKeysSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithExpirationDate(expirationDate string) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.expirationDate = expirationDate
    getFeatureKeysSpectraS3Request.queryParams.Set("expiration_date", expirationDate)
    return getFeatureKeysSpectraS3Request
}
func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithKey(key FeatureKeyType) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.key = key
    getFeatureKeysSpectraS3Request.queryParams.Set("key", key.String())
    return getFeatureKeysSpectraS3Request
}
func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithPageLength(pageLength int) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.pageLength = pageLength
    getFeatureKeysSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getFeatureKeysSpectraS3Request
}
func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithPageOffset(pageOffset int) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.pageOffset = pageOffset
    getFeatureKeysSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getFeatureKeysSpectraS3Request
}
func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.pageStartMarker = pageStartMarker
    getFeatureKeysSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithErrorMessage(errorMessage *string) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.errorMessage = errorMessage
    if errorMessage != nil {
        getFeatureKeysSpectraS3Request.queryParams.Set("error_message", *errorMessage)
    } else {
        getFeatureKeysSpectraS3Request.queryParams.Set("error_message", "")
    }
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithLastPage() *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.queryParams.Set("last_page", "")
    return getFeatureKeysSpectraS3Request
}

func (GetFeatureKeysSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) Path() string {
    return "/_rest_/feature_key"
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) QueryParams() *url.Values {
    return getFeatureKeysSpectraS3Request.queryParams
}

func (GetFeatureKeysSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetFeatureKeysSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetFeatureKeysSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
