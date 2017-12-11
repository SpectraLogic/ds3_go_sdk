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

type GetFeatureKeysSpectraS3Request struct {
    ErrorMessage *string
    ExpirationDate *string
    Key FeatureKeyType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetFeatureKeysSpectraS3Request() *GetFeatureKeysSpectraS3Request {
    return &GetFeatureKeysSpectraS3Request{
    }
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithErrorMessage(errorMessage string) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.ErrorMessage = &errorMessage
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithExpirationDate(expirationDate string) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.ExpirationDate = &expirationDate
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithKey(key FeatureKeyType) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.Key = key
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithLastPage() *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.LastPage = true
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithPageLength(pageLength int) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.PageLength = &pageLength
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithPageOffset(pageOffset int) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.PageOffset = &pageOffset
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.PageStartMarker = &pageStartMarker
    return getFeatureKeysSpectraS3Request
}

