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

type ModifyPoolPartitionSpectraS3Request struct {
    name *string
    poolPartition string
    queryParams *url.Values
}

func NewModifyPoolPartitionSpectraS3Request(poolPartition string) *ModifyPoolPartitionSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyPoolPartitionSpectraS3Request{
        poolPartition: poolPartition,
        queryParams: queryParams,
    }
}


func (modifyPoolPartitionSpectraS3Request *ModifyPoolPartitionSpectraS3Request) WithName(name *string) *ModifyPoolPartitionSpectraS3Request {
    modifyPoolPartitionSpectraS3Request.name = name
    if name != nil {
        modifyPoolPartitionSpectraS3Request.queryParams.Set("name", *name)
    } else {
        modifyPoolPartitionSpectraS3Request.queryParams.Set("name", "")
    }
    return modifyPoolPartitionSpectraS3Request
}


func (ModifyPoolPartitionSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyPoolPartitionSpectraS3Request *ModifyPoolPartitionSpectraS3Request) Path() string {
    return "/_rest_/pool_partition/" + modifyPoolPartitionSpectraS3Request.poolPartition
}

func (modifyPoolPartitionSpectraS3Request *ModifyPoolPartitionSpectraS3Request) QueryParams() *url.Values {
    return modifyPoolPartitionSpectraS3Request.queryParams
}

func (ModifyPoolPartitionSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyPoolPartitionSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyPoolPartitionSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
