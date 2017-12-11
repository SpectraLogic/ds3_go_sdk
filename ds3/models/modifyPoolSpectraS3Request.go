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

type ModifyPoolSpectraS3Request struct {
    PartitionId *string
    Pool string
    Quiesced Quiesced
}

func NewModifyPoolSpectraS3Request(pool string) *ModifyPoolSpectraS3Request {
    return &ModifyPoolSpectraS3Request{
        Pool: pool,
    }
}

func (modifyPoolSpectraS3Request *ModifyPoolSpectraS3Request) WithPartitionId(partitionId string) *ModifyPoolSpectraS3Request {
    modifyPoolSpectraS3Request.PartitionId = &partitionId
    return modifyPoolSpectraS3Request
}

func (modifyPoolSpectraS3Request *ModifyPoolSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyPoolSpectraS3Request {
    modifyPoolSpectraS3Request.Quiesced = quiesced
    return modifyPoolSpectraS3Request
}

