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
    "ds3/networking"
    "net/http"
)

type PutTapeDensityDirectiveSpectraS3Response struct {
    TapeDensityDirective TapeDensityDirective
    Headers *http.Header
}

func (putTapeDensityDirectiveSpectraS3Response *PutTapeDensityDirectiveSpectraS3Response) parse(webResponse networking.WebResponse) error {
        return parseResponsePayload(webResponse, &putTapeDensityDirectiveSpectraS3Response.TapeDensityDirective)
}

func NewPutTapeDensityDirectiveSpectraS3Response(webResponse networking.WebResponse) (*PutTapeDensityDirectiveSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutTapeDensityDirectiveSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}