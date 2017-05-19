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

type ImportAzureTargetSpectraS3Request struct {
    azureTarget string
    cloudBucketName *string
    conflictResolutionMode ImportConflictResolutionMode
    dataPolicyId string
    priority Priority
    userId string
    queryParams *url.Values
}

func NewImportAzureTargetSpectraS3Request(azureTarget string, cloudBucketName *string) *ImportAzureTargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "import")
    queryParams.Set("cloud_bucket_name", *cloudBucketName)

    return &ImportAzureTargetSpectraS3Request{
        azureTarget: azureTarget,
        cloudBucketName: cloudBucketName,
        queryParams: queryParams,
    }
}

func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithConflictResolutionMode(conflictResolutionMode ImportConflictResolutionMode) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.conflictResolutionMode = conflictResolutionMode
    importAzureTargetSpectraS3Request.queryParams.Set("conflict_resolution_mode", conflictResolutionMode.String())
    return importAzureTargetSpectraS3Request
}
func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.dataPolicyId = dataPolicyId
    importAzureTargetSpectraS3Request.queryParams.Set("data_policy_id", dataPolicyId)
    return importAzureTargetSpectraS3Request
}
func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithPriority(priority Priority) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.priority = priority
    importAzureTargetSpectraS3Request.queryParams.Set("priority", priority.String())
    return importAzureTargetSpectraS3Request
}
func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithUserId(userId string) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.userId = userId
    importAzureTargetSpectraS3Request.queryParams.Set("user_id", userId)
    return importAzureTargetSpectraS3Request
}



func (ImportAzureTargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) Path() string {
    return "/_rest_/azure_target/" + importAzureTargetSpectraS3Request.azureTarget
}

func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) QueryParams() *url.Values {
    return importAzureTargetSpectraS3Request.queryParams
}

func (ImportAzureTargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ImportAzureTargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ImportAzureTargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
