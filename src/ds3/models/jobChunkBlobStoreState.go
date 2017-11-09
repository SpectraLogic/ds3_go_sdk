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

type JobChunkBlobStoreState Enum

const (
    JOB_CHUNK_BLOB_STORE_STATE_PENDING JobChunkBlobStoreState = 1 + iota
    JOB_CHUNK_BLOB_STORE_STATE_IN_PROGRESS JobChunkBlobStoreState = 1 + iota
    JOB_CHUNK_BLOB_STORE_STATE_COMPLETED JobChunkBlobStoreState = 1 + iota
)

func (jobChunkBlobStoreState *JobChunkBlobStoreState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *jobChunkBlobStoreState = UNDEFINED
        case "PENDING": *jobChunkBlobStoreState = JOB_CHUNK_BLOB_STORE_STATE_PENDING
        case "IN_PROGRESS": *jobChunkBlobStoreState = JOB_CHUNK_BLOB_STORE_STATE_IN_PROGRESS
        case "COMPLETED": *jobChunkBlobStoreState = JOB_CHUNK_BLOB_STORE_STATE_COMPLETED
        default:
            *jobChunkBlobStoreState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into JobChunkBlobStoreState", str))
    }
    return nil
}

func (jobChunkBlobStoreState JobChunkBlobStoreState) String() string {
    switch jobChunkBlobStoreState {
        case JOB_CHUNK_BLOB_STORE_STATE_PENDING: return "PENDING"
        case JOB_CHUNK_BLOB_STORE_STATE_IN_PROGRESS: return "IN_PROGRESS"
        case JOB_CHUNK_BLOB_STORE_STATE_COMPLETED: return "COMPLETED"
        default:
            log.Printf("Error: invalid JobChunkBlobStoreState represented by '%d'", jobChunkBlobStoreState)
            return ""
    }
}

func (jobChunkBlobStoreState JobChunkBlobStoreState) StringPtr() *string {
    if jobChunkBlobStoreState == UNDEFINED {
        return nil
    }
    result := jobChunkBlobStoreState.String()
    return &result
}