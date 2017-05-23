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

type ModifyTapeSpectraS3Request struct {
    ejectLabel *string
    ejectLocation *string
    state TapeState
    tapeId string
    queryParams *url.Values
}

func NewModifyTapeSpectraS3Request(tapeId string) *ModifyTapeSpectraS3Request {
    queryParams := &url.Values{}

    return &ModifyTapeSpectraS3Request{
        tapeId: tapeId,
        queryParams: queryParams,
    }
}

func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) WithState(state TapeState) *ModifyTapeSpectraS3Request {
    modifyTapeSpectraS3Request.state = state
    modifyTapeSpectraS3Request.queryParams.Set("state", state.String())
    return modifyTapeSpectraS3Request
}

func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) WithEjectLabel(ejectLabel *string) *ModifyTapeSpectraS3Request {
    modifyTapeSpectraS3Request.ejectLabel = ejectLabel
    if ejectLabel != nil {
        modifyTapeSpectraS3Request.queryParams.Set("eject_label", *ejectLabel)
    } else {
        modifyTapeSpectraS3Request.queryParams.Set("eject_label", "")
    }
    return modifyTapeSpectraS3Request
}
func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) WithEjectLocation(ejectLocation *string) *ModifyTapeSpectraS3Request {
    modifyTapeSpectraS3Request.ejectLocation = ejectLocation
    if ejectLocation != nil {
        modifyTapeSpectraS3Request.queryParams.Set("eject_location", *ejectLocation)
    } else {
        modifyTapeSpectraS3Request.queryParams.Set("eject_location", "")
    }
    return modifyTapeSpectraS3Request
}


func (ModifyTapeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) Path() string {
    return "/_rest_/tape/" + modifyTapeSpectraS3Request.tapeId
}

func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) QueryParams() *url.Values {
    return modifyTapeSpectraS3Request.queryParams
}

func (ModifyTapeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (ModifyTapeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (ModifyTapeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
