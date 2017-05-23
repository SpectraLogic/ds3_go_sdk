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

type DeleteAzureTargetBucketNameSpectraS3Request struct {
    azureTargetBucketName string
    queryParams *url.Values
}

func NewDeleteAzureTargetBucketNameSpectraS3Request(azureTargetBucketName string) *DeleteAzureTargetBucketNameSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteAzureTargetBucketNameSpectraS3Request{
        azureTargetBucketName: azureTargetBucketName,
        queryParams: queryParams,
    }
}




func (DeleteAzureTargetBucketNameSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteAzureTargetBucketNameSpectraS3Request *DeleteAzureTargetBucketNameSpectraS3Request) Path() string {
    return "/_rest_/azure_target_bucket_name/" + deleteAzureTargetBucketNameSpectraS3Request.azureTargetBucketName
}

func (deleteAzureTargetBucketNameSpectraS3Request *DeleteAzureTargetBucketNameSpectraS3Request) QueryParams() *url.Values {
    return deleteAzureTargetBucketNameSpectraS3Request.queryParams
}

func (DeleteAzureTargetBucketNameSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteAzureTargetBucketNameSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteAzureTargetBucketNameSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
