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

type VerifyAzureTargetSpectraS3Request struct {
    azureTarget string
    queryParams *url.Values
}

func NewVerifyAzureTargetSpectraS3Request(azureTarget string) *VerifyAzureTargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "verify")

    return &VerifyAzureTargetSpectraS3Request{
        azureTarget: azureTarget,
        queryParams: queryParams,
    }
}



func (verifyAzureTargetSpectraS3Request *VerifyAzureTargetSpectraS3Request) WithFullDetails() *VerifyAzureTargetSpectraS3Request {
    verifyAzureTargetSpectraS3Request.queryParams.Set("full_details", "")
    return verifyAzureTargetSpectraS3Request
}

func (VerifyAzureTargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (verifyAzureTargetSpectraS3Request *VerifyAzureTargetSpectraS3Request) Path() string {
    return "/_rest_/azure_target/" + verifyAzureTargetSpectraS3Request.azureTarget
}

func (verifyAzureTargetSpectraS3Request *VerifyAzureTargetSpectraS3Request) QueryParams() *url.Values {
    return verifyAzureTargetSpectraS3Request.queryParams
}

func (VerifyAzureTargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifyAzureTargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (VerifyAzureTargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
