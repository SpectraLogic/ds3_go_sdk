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

type EjectTapeSpectraS3Request struct {
    ejectLabel *string
    ejectLocation *string
    tapeId string
    queryParams *url.Values
}

func NewEjectTapeSpectraS3Request(tapeId string) *EjectTapeSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "eject")

    return &EjectTapeSpectraS3Request{
        tapeId: tapeId,
        queryParams: queryParams,
    }
}


func (ejectTapeSpectraS3Request *EjectTapeSpectraS3Request) WithEjectLabel(ejectLabel *string) *EjectTapeSpectraS3Request {
    ejectTapeSpectraS3Request.ejectLabel = ejectLabel
    if ejectLabel != nil {
        ejectTapeSpectraS3Request.queryParams.Set("eject_label", *ejectLabel)
    } else {
        ejectTapeSpectraS3Request.queryParams.Set("eject_label", "")
    }
    return ejectTapeSpectraS3Request
}
func (ejectTapeSpectraS3Request *EjectTapeSpectraS3Request) WithEjectLocation(ejectLocation *string) *EjectTapeSpectraS3Request {
    ejectTapeSpectraS3Request.ejectLocation = ejectLocation
    if ejectLocation != nil {
        ejectTapeSpectraS3Request.queryParams.Set("eject_location", *ejectLocation)
    } else {
        ejectTapeSpectraS3Request.queryParams.Set("eject_location", "")
    }
    return ejectTapeSpectraS3Request
}


func (EjectTapeSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (ejectTapeSpectraS3Request *EjectTapeSpectraS3Request) Path() string {
    return "/_rest_/tape/" + ejectTapeSpectraS3Request.tapeId
}

func (ejectTapeSpectraS3Request *EjectTapeSpectraS3Request) QueryParams() *url.Values {
    return ejectTapeSpectraS3Request.queryParams
}

func (EjectTapeSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (EjectTapeSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (EjectTapeSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
