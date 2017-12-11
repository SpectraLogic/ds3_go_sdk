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

type GetS3TargetReadPreferencesSpectraS3Request struct {
    BucketId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    ReadPreference TargetReadPreferenceType
    TargetId *string
}

func NewGetS3TargetReadPreferencesSpectraS3Request() *GetS3TargetReadPreferencesSpectraS3Request {
    return &GetS3TargetReadPreferencesSpectraS3Request{
    }
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithBucketId(bucketId string) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.BucketId = &bucketId
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithLastPage() *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.LastPage = true
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.PageLength = &pageLength
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.PageOffset = &pageOffset
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithReadPreference(readPreference TargetReadPreferenceType) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.ReadPreference = readPreference
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithTargetId(targetId string) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.TargetId = &targetId
    return getS3TargetReadPreferencesSpectraS3Request
}

