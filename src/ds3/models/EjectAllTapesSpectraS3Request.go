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

type EjectAllTapesSpectraS3Request struct {
    ejectLabel *string
    ejectLocation *string
    queryParams *url.Values
}

func NewEjectAllTapesSpectraS3Request() *EjectAllTapesSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "eject")

    return &EjectAllTapesSpectraS3Request{
        queryParams: queryParams,
    }
}


func (ejectAllTapesSpectraS3Request *EjectAllTapesSpectraS3Request) WithEjectLabel(ejectLabel *string) *EjectAllTapesSpectraS3Request {
    ejectAllTapesSpectraS3Request.ejectLabel = ejectLabel
    if ejectLabel != nil {
        ejectAllTapesSpectraS3Request.queryParams.Set("eject_label", *ejectLabel)
    } else {
        ejectAllTapesSpectraS3Request.queryParams.Set("eject_label", "")
    }
    return ejectAllTapesSpectraS3Request
}
func (ejectAllTapesSpectraS3Request *EjectAllTapesSpectraS3Request) WithEjectLocation(ejectLocation *string) *EjectAllTapesSpectraS3Request {
    ejectAllTapesSpectraS3Request.ejectLocation = ejectLocation
    if ejectLocation != nil {
        ejectAllTapesSpectraS3Request.queryParams.Set("eject_location", *ejectLocation)
    } else {
        ejectAllTapesSpectraS3Request.queryParams.Set("eject_location", "")
    }
    return ejectAllTapesSpectraS3Request
}


func (EjectAllTapesSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (ejectAllTapesSpectraS3Request *EjectAllTapesSpectraS3Request) Path() string {
    return "/_rest_/tape"
}

func (ejectAllTapesSpectraS3Request *EjectAllTapesSpectraS3Request) QueryParams() *url.Values {
    return ejectAllTapesSpectraS3Request.queryParams
}

func (EjectAllTapesSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (EjectAllTapesSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (EjectAllTapesSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
