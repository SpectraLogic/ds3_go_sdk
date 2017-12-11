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

type JobStatus Enum

const (
    JOB_STATUS_IN_PROGRESS JobStatus = 1 + iota
    JOB_STATUS_COMPLETED JobStatus = 1 + iota
    JOB_STATUS_CANCELED JobStatus = 1 + iota
)

func (jobStatus *JobStatus) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *jobStatus = UNDEFINED
        case "IN_PROGRESS": *jobStatus = JOB_STATUS_IN_PROGRESS
        case "COMPLETED": *jobStatus = JOB_STATUS_COMPLETED
        case "CANCELED": *jobStatus = JOB_STATUS_CANCELED
        default:
            *jobStatus = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into JobStatus", str))
    }
    return nil
}

func (jobStatus JobStatus) String() string {
    switch jobStatus {
        case JOB_STATUS_IN_PROGRESS: return "IN_PROGRESS"
        case JOB_STATUS_COMPLETED: return "COMPLETED"
        case JOB_STATUS_CANCELED: return "CANCELED"
        default:
            log.Printf("Error: invalid JobStatus represented by '%d'", jobStatus)
            return ""
    }
}

func (jobStatus JobStatus) StringPtr() *string {
    if jobStatus == UNDEFINED {
        return nil
    }
    result := jobStatus.String()
    return &result
}