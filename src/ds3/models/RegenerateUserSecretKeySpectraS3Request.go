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

type RegenerateUserSecretKeySpectraS3Request struct {
    userId string
    queryParams *url.Values
}

func NewRegenerateUserSecretKeySpectraS3Request(userId string) *RegenerateUserSecretKeySpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "regenerate_secret_key")

    return &RegenerateUserSecretKeySpectraS3Request{
        userId: userId,
        queryParams: queryParams,
    }
}




func (RegenerateUserSecretKeySpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (regenerateUserSecretKeySpectraS3Request *RegenerateUserSecretKeySpectraS3Request) Path() string {
    return "/_rest_/user/" + regenerateUserSecretKeySpectraS3Request.userId
}

func (regenerateUserSecretKeySpectraS3Request *RegenerateUserSecretKeySpectraS3Request) QueryParams() *url.Values {
    return regenerateUserSecretKeySpectraS3Request.queryParams
}

func (RegenerateUserSecretKeySpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (RegenerateUserSecretKeySpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (RegenerateUserSecretKeySpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
