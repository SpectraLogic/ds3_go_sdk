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

type ImportS3TargetSpectraS3Request struct {
    CloudBucketName string
    ConflictResolutionMode ImportConflictResolutionMode
    DataPolicyId *string
    Priority Priority
    S3Target string
    UserId *string
}

func NewImportS3TargetSpectraS3Request(cloudBucketName string, s3Target string) *ImportS3TargetSpectraS3Request {
    return &ImportS3TargetSpectraS3Request{
        S3Target: s3Target,
        CloudBucketName: cloudBucketName,
    }
}

func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.ConflictResolutionMode = conflictResolutionMode
    return importS3TargetSpectraS3Request
}

func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.DataPolicyId = &dataPolicyId
    return importS3TargetSpectraS3Request
}

func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithPriority(priority Priority) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.Priority = priority
    return importS3TargetSpectraS3Request
}

func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithUserId(userId string) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.UserId = &userId
    return importS3TargetSpectraS3Request
}

