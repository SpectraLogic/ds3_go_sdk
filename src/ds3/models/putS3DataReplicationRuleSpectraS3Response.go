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

type PutS3DataReplicationRuleSpectraS3Response struct {
    S3DataReplicationRule S3DataReplicationRule
    Headers *http.Header
}

func (putS3DataReplicationRuleSpectraS3Response *PutS3DataReplicationRuleSpectraS3Response) parse(webResponse networking.WebResponse) error {
        return parseResponsePayload(webResponse, &putS3DataReplicationRuleSpectraS3Response.S3DataReplicationRule)
}

func NewPutS3DataReplicationRuleSpectraS3Response(webResponse networking.WebResponse) (*PutS3DataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutS3DataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
