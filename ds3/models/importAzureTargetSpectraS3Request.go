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

type ImportAzureTargetSpectraS3Request struct {
    AzureTarget string
    CloudBucketName string
    ConflictResolutionMode ImportConflictResolutionMode
    DataPolicyId *string
    Priority Priority
    UserId *string
}

func NewImportAzureTargetSpectraS3Request(azureTarget string, cloudBucketName string) *ImportAzureTargetSpectraS3Request {
    return &ImportAzureTargetSpectraS3Request{
        AzureTarget: azureTarget,
        CloudBucketName: cloudBucketName,
    }
}

func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.ConflictResolutionMode = conflictResolutionMode
    return importAzureTargetSpectraS3Request
}

func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.DataPolicyId = &dataPolicyId
    return importAzureTargetSpectraS3Request
}

func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithPriority(priority Priority) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.Priority = priority
    return importAzureTargetSpectraS3Request
}

func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithUserId(userId string) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.UserId = &userId
    return importAzureTargetSpectraS3Request
}

