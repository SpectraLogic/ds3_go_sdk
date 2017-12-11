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

type ModifyTapeSpectraS3Request struct {
    EjectLabel *string
    EjectLocation *string
    State TapeState
    TapeId string
}

func NewModifyTapeSpectraS3Request(tapeId string) *ModifyTapeSpectraS3Request {
    return &ModifyTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) WithEjectLabel(ejectLabel string) *ModifyTapeSpectraS3Request {
    modifyTapeSpectraS3Request.EjectLabel = &ejectLabel
    return modifyTapeSpectraS3Request
}

func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) WithEjectLocation(ejectLocation string) *ModifyTapeSpectraS3Request {
    modifyTapeSpectraS3Request.EjectLocation = &ejectLocation
    return modifyTapeSpectraS3Request
}

func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) WithState(state TapeState) *ModifyTapeSpectraS3Request {
    modifyTapeSpectraS3Request.State = state
    return modifyTapeSpectraS3Request
}

