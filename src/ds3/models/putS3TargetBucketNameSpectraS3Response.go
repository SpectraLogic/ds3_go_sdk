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
    "ds3/networking"
)

type PutS3TargetBucketNameSpectraS3Response struct {
    S3TargetBucketName S3TargetBucketName `xml:"S3TargetBucketName"`
}

func NewPutS3TargetBucketNameSpectraS3Response(webResponse networking.WebResponse) (*PutS3TargetBucketNameSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutS3TargetBucketNameSpectraS3Response
        if err := readResponseBody(webResponse, &body); err != nil {
            return nil, err
        }
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
