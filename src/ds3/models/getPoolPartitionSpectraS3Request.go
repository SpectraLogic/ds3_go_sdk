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

type GetPoolPartitionSpectraS3Request struct {
    poolPartition string
    queryParams *url.Values
}

func NewGetPoolPartitionSpectraS3Request(poolPartition string) *GetPoolPartitionSpectraS3Request {
    queryParams := &url.Values{}

    return &GetPoolPartitionSpectraS3Request{
        poolPartition: poolPartition,
        queryParams: queryParams,
    }
}




func (GetPoolPartitionSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getPoolPartitionSpectraS3Request *GetPoolPartitionSpectraS3Request) Path() string {
    return "/_rest_/pool_partition/" + getPoolPartitionSpectraS3Request.poolPartition
}

func (getPoolPartitionSpectraS3Request *GetPoolPartitionSpectraS3Request) QueryParams() *url.Values {
    return getPoolPartitionSpectraS3Request.queryParams
}

func (GetPoolPartitionSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetPoolPartitionSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetPoolPartitionSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
