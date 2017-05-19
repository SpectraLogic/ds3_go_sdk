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
    "errors"
    "fmt"
    "bytes"
    "log"
)

type JobRequestType Enum

const (
    JOB_REQUEST_TYPE_PUT JobRequestType = 1 + iota
    JOB_REQUEST_TYPE_GET JobRequestType = 1 + iota
    JOB_REQUEST_TYPE_VERIFY JobRequestType = 1 + iota
)

func (jobRequestType *JobRequestType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "PUT": *jobRequestType = JOB_REQUEST_TYPE_PUT
        case "GET": *jobRequestType = JOB_REQUEST_TYPE_GET
        case "VERIFY": *jobRequestType = JOB_REQUEST_TYPE_VERIFY
        default:
            *jobRequestType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into JobRequestType", str))
    }
    return nil
}

func (jobRequestType JobRequestType) String() string {
    switch jobRequestType {
        case JOB_REQUEST_TYPE_PUT: return "PUT"
        case JOB_REQUEST_TYPE_GET: return "GET"
        case JOB_REQUEST_TYPE_VERIFY: return "VERIFY"
        default:
            log.Printf("Error: invalid JobRequestType represented by '%d'", jobRequestType)
            return ""
    }
}
