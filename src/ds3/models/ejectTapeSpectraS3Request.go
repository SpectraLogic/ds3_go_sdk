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

type EjectTapeSpectraS3Request struct {
    EjectLabel *string
    EjectLocation *string
    TapeId string
}

func NewEjectTapeSpectraS3Request(tapeId string) *EjectTapeSpectraS3Request {
    return &EjectTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (ejectTapeSpectraS3Request *EjectTapeSpectraS3Request) WithEjectLabel(ejectLabel string) *EjectTapeSpectraS3Request {
    ejectTapeSpectraS3Request.EjectLabel = &ejectLabel
    return ejectTapeSpectraS3Request
}

func (ejectTapeSpectraS3Request *EjectTapeSpectraS3Request) WithEjectLocation(ejectLocation string) *EjectTapeSpectraS3Request {
    ejectTapeSpectraS3Request.EjectLocation = &ejectLocation
    return ejectTapeSpectraS3Request
}

