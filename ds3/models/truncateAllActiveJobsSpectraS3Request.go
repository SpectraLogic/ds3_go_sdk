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

type TruncateAllActiveJobsSpectraS3Request struct {
    BucketId *string
    RequestType JobRequestType
}

func NewTruncateAllActiveJobsSpectraS3Request() *TruncateAllActiveJobsSpectraS3Request {
    return &TruncateAllActiveJobsSpectraS3Request{
    }
}

func (truncateAllActiveJobsSpectraS3Request *TruncateAllActiveJobsSpectraS3Request) WithBucketId(bucketId string) *TruncateAllActiveJobsSpectraS3Request {
    truncateAllActiveJobsSpectraS3Request.BucketId = &bucketId
    return truncateAllActiveJobsSpectraS3Request
}

func (truncateAllActiveJobsSpectraS3Request *TruncateAllActiveJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *TruncateAllActiveJobsSpectraS3Request {
    truncateAllActiveJobsSpectraS3Request.RequestType = requestType
    return truncateAllActiveJobsSpectraS3Request
}

