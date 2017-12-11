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

type GetDs3TargetReadPreferencesSpectraS3Request struct {
    BucketId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    ReadPreference TargetReadPreferenceType
    TargetId *string
}

func NewGetDs3TargetReadPreferencesSpectraS3Request() *GetDs3TargetReadPreferencesSpectraS3Request {
    return &GetDs3TargetReadPreferencesSpectraS3Request{
    }
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithBucketId(bucketId string) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.BucketId = &bucketId
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithLastPage() *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.LastPage = true
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.PageLength = &pageLength
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.PageOffset = &pageOffset
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithReadPreference(readPreference TargetReadPreferenceType) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.ReadPreference = readPreference
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithTargetId(targetId string) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.TargetId = &targetId
    return getDs3TargetReadPreferencesSpectraS3Request
}

