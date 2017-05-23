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

type PutS3TargetReadPreferenceSpectraS3Request struct {
    bucketId string
    readPreference TargetReadPreferenceType
    targetId string
    queryParams *url.Values
}

func NewPutS3TargetReadPreferenceSpectraS3Request(bucketId string, readPreference TargetReadPreferenceType, targetId string) *PutS3TargetReadPreferenceSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("bucket_id", bucketId)
    queryParams.Set("read_preference", readPreference.String())
    queryParams.Set("target_id", targetId)

    return &PutS3TargetReadPreferenceSpectraS3Request{
        bucketId: bucketId,
        readPreference: readPreference,
        targetId: targetId,
        queryParams: queryParams,
    }
}




func (PutS3TargetReadPreferenceSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putS3TargetReadPreferenceSpectraS3Request *PutS3TargetReadPreferenceSpectraS3Request) Path() string {
    return "/_rest_/s3_target_read_preference"
}

func (putS3TargetReadPreferenceSpectraS3Request *PutS3TargetReadPreferenceSpectraS3Request) QueryParams() *url.Values {
    return putS3TargetReadPreferenceSpectraS3Request.queryParams
}

func (PutS3TargetReadPreferenceSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutS3TargetReadPreferenceSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutS3TargetReadPreferenceSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
