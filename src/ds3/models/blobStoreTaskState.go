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

type BlobStoreTaskState Enum

const (
    BLOB_STORE_TASK_STATE_NOT_READY BlobStoreTaskState = 1 + iota
    BLOB_STORE_TASK_STATE_READY BlobStoreTaskState = 1 + iota
    BLOB_STORE_TASK_STATE_PENDING_EXECUTION BlobStoreTaskState = 1 + iota
    BLOB_STORE_TASK_STATE_IN_PROGRESS BlobStoreTaskState = 1 + iota
    BLOB_STORE_TASK_STATE_COMPLETED BlobStoreTaskState = 1 + iota
)

func (blobStoreTaskState *BlobStoreTaskState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *blobStoreTaskState = UNDEFINED
        case "NOT_READY": *blobStoreTaskState = BLOB_STORE_TASK_STATE_NOT_READY
        case "READY": *blobStoreTaskState = BLOB_STORE_TASK_STATE_READY
        case "PENDING_EXECUTION": *blobStoreTaskState = BLOB_STORE_TASK_STATE_PENDING_EXECUTION
        case "IN_PROGRESS": *blobStoreTaskState = BLOB_STORE_TASK_STATE_IN_PROGRESS
        case "COMPLETED": *blobStoreTaskState = BLOB_STORE_TASK_STATE_COMPLETED
        default:
            *blobStoreTaskState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into BlobStoreTaskState", str))
    }
    return nil
}

func (blobStoreTaskState BlobStoreTaskState) String() string {
    switch blobStoreTaskState {
        case BLOB_STORE_TASK_STATE_NOT_READY: return "NOT_READY"
        case BLOB_STORE_TASK_STATE_READY: return "READY"
        case BLOB_STORE_TASK_STATE_PENDING_EXECUTION: return "PENDING_EXECUTION"
        case BLOB_STORE_TASK_STATE_IN_PROGRESS: return "IN_PROGRESS"
        case BLOB_STORE_TASK_STATE_COMPLETED: return "COMPLETED"
        default:
            log.Printf("Error: invalid BlobStoreTaskState represented by '%d'", blobStoreTaskState)
            return ""
    }
}
