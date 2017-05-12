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

type ModifyGroupSpectraS3Request struct {
    group string
    name *string
    queryParams *url.Values
}

func NewModifyGroupSpectraS3Request(group string) *ModifyGroupSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyGroupSpectraS3Request{
        group: group,
        queryParams: queryParams,
    }
}


func (modifyGroupSpectraS3Request *ModifyGroupSpectraS3Request) WithName(name *string) *ModifyGroupSpectraS3Request {
    modifyGroupSpectraS3Request.name = name
    if name != nil {
        modifyGroupSpectraS3Request.queryParams.Set("name", *name)
    } else {
        modifyGroupSpectraS3Request.queryParams.Set("name", "")
    }
    return modifyGroupSpectraS3Request
}


func (ModifyGroupSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyGroupSpectraS3Request *ModifyGroupSpectraS3Request) Path() string {
    return "/_rest_/group/" + modifyGroupSpectraS3Request.group
}

func (modifyGroupSpectraS3Request *ModifyGroupSpectraS3Request) QueryParams() *url.Values {
    return modifyGroupSpectraS3Request.queryParams
}

func (ModifyGroupSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyGroupSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyGroupSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
