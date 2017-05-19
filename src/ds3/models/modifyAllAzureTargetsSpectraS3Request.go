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

type ModifyAllAzureTargetsSpectraS3Request struct {
    quiesced Quiesced
    queryParams *url.Values
}

func NewModifyAllAzureTargetsSpectraS3Request(quiesced Quiesced) *ModifyAllAzureTargetsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("quiesced", quiesced.String())

    return &ModifyAllAzureTargetsSpectraS3Request{
        quiesced: quiesced,
        queryParams: queryParams,
    }
}




func (ModifyAllAzureTargetsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyAllAzureTargetsSpectraS3Request *ModifyAllAzureTargetsSpectraS3Request) Path() string {
    return "/_rest_/azure_target"
}

func (modifyAllAzureTargetsSpectraS3Request *ModifyAllAzureTargetsSpectraS3Request) QueryParams() *url.Values {
    return modifyAllAzureTargetsSpectraS3Request.queryParams
}

func (ModifyAllAzureTargetsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyAllAzureTargetsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyAllAzureTargetsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
