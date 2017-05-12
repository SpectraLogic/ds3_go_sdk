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

type DelegateCreateUserSpectraS3Request struct {
    id string
    name *string
    secretKey *string
    queryParams *url.Values
}

func NewDelegateCreateUserSpectraS3Request(name *string) *DelegateCreateUserSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("name", *name)

    return &DelegateCreateUserSpectraS3Request{
        name: name,
        queryParams: queryParams,
    }
}

func (delegateCreateUserSpectraS3Request *DelegateCreateUserSpectraS3Request) WithId(id string) *DelegateCreateUserSpectraS3Request {
    delegateCreateUserSpectraS3Request.id = id
    delegateCreateUserSpectraS3Request.queryParams.Set("id", id)
    return delegateCreateUserSpectraS3Request
}

func (delegateCreateUserSpectraS3Request *DelegateCreateUserSpectraS3Request) WithSecretKey(secretKey *string) *DelegateCreateUserSpectraS3Request {
    delegateCreateUserSpectraS3Request.secretKey = secretKey
    if secretKey != nil {
        delegateCreateUserSpectraS3Request.queryParams.Set("secret_key", *secretKey)
    } else {
        delegateCreateUserSpectraS3Request.queryParams.Set("secret_key", "")
    }
    return delegateCreateUserSpectraS3Request
}


func (DelegateCreateUserSpectraS3Request) Verb() networking.HttpVerb {
    return networking.POST
}

func (delegateCreateUserSpectraS3Request *DelegateCreateUserSpectraS3Request) Path() string {
    return "/_rest_/user"
}

func (delegateCreateUserSpectraS3Request *DelegateCreateUserSpectraS3Request) QueryParams() *url.Values {
    return delegateCreateUserSpectraS3Request.queryParams
}

func (DelegateCreateUserSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DelegateCreateUserSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DelegateCreateUserSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
