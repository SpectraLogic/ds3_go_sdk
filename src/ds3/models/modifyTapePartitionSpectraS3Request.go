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

type ModifyTapePartitionSpectraS3Request struct {
    quiesced Quiesced
    tapePartition string
    queryParams *url.Values
}

func NewModifyTapePartitionSpectraS3Request(tapePartition string) *ModifyTapePartitionSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyTapePartitionSpectraS3Request{
        tapePartition: tapePartition,
        queryParams: queryParams,
    }
}

func (modifyTapePartitionSpectraS3Request *ModifyTapePartitionSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyTapePartitionSpectraS3Request {
    modifyTapePartitionSpectraS3Request.quiesced = quiesced
    modifyTapePartitionSpectraS3Request.queryParams.Set("quiesced", quiesced.String())
    return modifyTapePartitionSpectraS3Request
}



func (ModifyTapePartitionSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyTapePartitionSpectraS3Request *ModifyTapePartitionSpectraS3Request) Path() string {
    return "/_rest_/tape_partition/" + modifyTapePartitionSpectraS3Request.tapePartition
}

func (modifyTapePartitionSpectraS3Request *ModifyTapePartitionSpectraS3Request) QueryParams() *url.Values {
    return modifyTapePartitionSpectraS3Request.queryParams
}

func (ModifyTapePartitionSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyTapePartitionSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyTapePartitionSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
