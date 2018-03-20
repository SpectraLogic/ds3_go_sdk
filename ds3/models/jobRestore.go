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

type JobRestore Enum

const (
    JOB_RESTORE_NO JobRestore = 1 + iota
    JOB_RESTORE_YES JobRestore = 1 + iota
    JOB_RESTORE_PERMANENT_ONLY JobRestore = 1 + iota
)

func (jobRestore *JobRestore) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *jobRestore = UNDEFINED
        case "NO": *jobRestore = JOB_RESTORE_NO
        case "YES": *jobRestore = JOB_RESTORE_YES
        case "PERMANENT_ONLY": *jobRestore = JOB_RESTORE_PERMANENT_ONLY
        default:
            *jobRestore = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into JobRestore", str))
    }
    return nil
}

func (jobRestore JobRestore) String() string {
    switch jobRestore {
        case JOB_RESTORE_NO: return "NO"
        case JOB_RESTORE_YES: return "YES"
        case JOB_RESTORE_PERMANENT_ONLY: return "PERMANENT_ONLY"
        default:
            log.Printf("Error: invalid JobRestore represented by '%d'", jobRestore)
            return ""
    }
}

func (jobRestore JobRestore) StringPtr() *string {
    if jobRestore == UNDEFINED {
        return nil
    }
    result := jobRestore.String()
    return &result
}