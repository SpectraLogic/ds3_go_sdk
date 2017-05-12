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
)

type PutAzureTargetReadPreferenceSpectraS3Request struct {
    bucketId string
    readPreference TargetReadPreferenceType
    targetId string
    queryParams *url.Values
}

func NewPutAzureTargetReadPreferenceSpectraS3Request(bucketId string, readPreference TargetReadPreferenceType, targetId string) *PutAzureTargetReadPreferenceSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("bucket_id", bucketId)
    queryParams.Set("read_preference", readPreference.String())
    queryParams.Set("target_id", targetId)

    return &PutAzureTargetReadPreferenceSpectraS3Request{
        bucketId: bucketId,
        readPreference: readPreference,
        targetId: targetId,
        queryParams: queryParams,
    }
}




func (PutAzureTargetReadPreferenceSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putAzureTargetReadPreferenceSpectraS3Request *PutAzureTargetReadPreferenceSpectraS3Request) Path() string {
    return "/_rest_/azure_target_read_preference"
}

func (putAzureTargetReadPreferenceSpectraS3Request *PutAzureTargetReadPreferenceSpectraS3Request) QueryParams() *url.Values {
    return putAzureTargetReadPreferenceSpectraS3Request.queryParams
}

func (PutAzureTargetReadPreferenceSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutAzureTargetReadPreferenceSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutAzureTargetReadPreferenceSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
