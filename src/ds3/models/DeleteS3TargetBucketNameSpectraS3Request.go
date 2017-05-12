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

type DeleteS3TargetBucketNameSpectraS3Request struct {
    s3TargetBucketName string
    queryParams *url.Values
}

func NewDeleteS3TargetBucketNameSpectraS3Request(s3TargetBucketName string) *DeleteS3TargetBucketNameSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteS3TargetBucketNameSpectraS3Request{
        s3TargetBucketName: s3TargetBucketName,
        queryParams: queryParams,
    }
}




func (DeleteS3TargetBucketNameSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteS3TargetBucketNameSpectraS3Request *DeleteS3TargetBucketNameSpectraS3Request) Path() string {
    return "/_rest_/s3_target_bucket_name/" + deleteS3TargetBucketNameSpectraS3Request.s3TargetBucketName
}

func (deleteS3TargetBucketNameSpectraS3Request *DeleteS3TargetBucketNameSpectraS3Request) QueryParams() *url.Values {
    return deleteS3TargetBucketNameSpectraS3Request.queryParams
}

func (DeleteS3TargetBucketNameSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteS3TargetBucketNameSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteS3TargetBucketNameSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
