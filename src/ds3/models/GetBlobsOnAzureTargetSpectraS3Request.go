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

type GetBlobsOnAzureTargetSpectraS3Request struct {
    azureTarget string
    content networking.ReaderWithSizeDecorator
    queryParams *url.Values
}

func NewGetBlobsOnAzureTargetSpectraS3Request(azureTarget string, objects []Ds3Object) *GetBlobsOnAzureTargetSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "get_physical_placement")

    return &GetBlobsOnAzureTargetSpectraS3Request{
        azureTarget: azureTarget,
        content: buildDs3ObjectListStream(objects),
        queryParams: queryParams,
    }
}




func (GetBlobsOnAzureTargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getBlobsOnAzureTargetSpectraS3Request *GetBlobsOnAzureTargetSpectraS3Request) Path() string {
    return "/_rest_/azure_target/" + getBlobsOnAzureTargetSpectraS3Request.azureTarget
}

func (getBlobsOnAzureTargetSpectraS3Request *GetBlobsOnAzureTargetSpectraS3Request) QueryParams() *url.Values {
    return getBlobsOnAzureTargetSpectraS3Request.queryParams
}

func (GetBlobsOnAzureTargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetBlobsOnAzureTargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (getBlobsOnAzureTargetSpectraS3Request *GetBlobsOnAzureTargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return getBlobsOnAzureTargetSpectraS3Request.content
}
