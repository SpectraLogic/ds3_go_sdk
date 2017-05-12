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

type ImportS3TargetSpectraS3Request struct {
    cloudBucketName *string
    conflictResolutionMode ImportConflictResolutionMode
    dataPolicyId string
    priority Priority
    s3Target string
    userId string
    queryParams *url.Values
}

func NewImportS3TargetSpectraS3Request(cloudBucketName *string, s3Target string) *ImportS3TargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "import")
    queryParams.Set("cloud_bucket_name", *cloudBucketName)

    return &ImportS3TargetSpectraS3Request{
        s3Target: s3Target,
        cloudBucketName: cloudBucketName,
        queryParams: queryParams,
    }
}

func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.conflictResolutionMode = conflictResolutionMode
    importS3TargetSpectraS3Request.queryParams.Set("conflict_resolution_mode", conflictResolutionMode.String())
    return importS3TargetSpectraS3Request
}
func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.dataPolicyId = dataPolicyId
    importS3TargetSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return importS3TargetSpectraS3Request
}
func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithPriority(priority Priority) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.priority = priority
    importS3TargetSpectraS3Request.queryParams.Set("priority", priority.String())
    return importS3TargetSpectraS3Request
}
func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithUserId(userId string) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.userId = userId
    importS3TargetSpectraS3Request.queryParams.Set("user_id", userId)
    return importS3TargetSpectraS3Request
}



func (ImportS3TargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) Path() string {
    return "/_rest_/s3_target/" + importS3TargetSpectraS3Request.s3Target
}

func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) QueryParams() *url.Values {
    return importS3TargetSpectraS3Request.queryParams
}

func (ImportS3TargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ImportS3TargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ImportS3TargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
