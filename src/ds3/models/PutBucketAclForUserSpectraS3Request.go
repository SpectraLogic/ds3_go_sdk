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

type PutBucketAclForUserSpectraS3Request struct {
    bucketId string
    permission BucketAclPermission
    userId string
    queryParams *url.Values
}

func NewPutBucketAclForUserSpectraS3Request(bucketId string, permission BucketAclPermission, userId string) *PutBucketAclForUserSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("bucket_id", bucketId)
    queryParams.Set("permission", permission.String())
    queryParams.Set("user_id", userId)

    return &PutBucketAclForUserSpectraS3Request{
        bucketId: bucketId,
        permission: permission,
        userId: userId,
        queryParams: queryParams,
    }
}




func (PutBucketAclForUserSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (putBucketAclForUserSpectraS3Request *PutBucketAclForUserSpectraS3Request) Path() string {
    return "/_rest_/bucket_acl"
}

func (putBucketAclForUserSpectraS3Request *PutBucketAclForUserSpectraS3Request) QueryParams() *url.Values {
    return putBucketAclForUserSpectraS3Request.queryParams
}

func (PutBucketAclForUserSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (PutBucketAclForUserSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (PutBucketAclForUserSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
