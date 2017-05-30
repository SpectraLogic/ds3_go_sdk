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

type PoolFailureType Enum

const (
    POOL_FAILURE_TYPE_BLOB_READ_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_DATA_CHECKPOINT_MISSING PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_FORMAT_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_IMPORT_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_INSPECT_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_QUIESCED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_READ_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_VERIFY_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_WRITE_FAILED PoolFailureType = 1 + iota
)

func (poolFailureType *PoolFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *poolFailureType = UNDEFINED
        case "BLOB_READ_FAILED": *poolFailureType = POOL_FAILURE_TYPE_BLOB_READ_FAILED
        case "DATA_CHECKPOINT_FAILURE": *poolFailureType = POOL_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE
        case "DATA_CHECKPOINT_MISSING": *poolFailureType = POOL_FAILURE_TYPE_DATA_CHECKPOINT_MISSING
        case "FORMAT_FAILED": *poolFailureType = POOL_FAILURE_TYPE_FORMAT_FAILED
        case "IMPORT_FAILED": *poolFailureType = POOL_FAILURE_TYPE_IMPORT_FAILED
        case "IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE": *poolFailureType = POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE
        case "IMPORT_FAILED_DUE_TO_DATA_INTEGRITY": *poolFailureType = POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY
        case "INSPECT_FAILED": *poolFailureType = POOL_FAILURE_TYPE_INSPECT_FAILED
        case "QUIESCED": *poolFailureType = POOL_FAILURE_TYPE_QUIESCED
        case "READ_FAILED": *poolFailureType = POOL_FAILURE_TYPE_READ_FAILED
        case "VERIFY_FAILED": *poolFailureType = POOL_FAILURE_TYPE_VERIFY_FAILED
        case "WRITE_FAILED": *poolFailureType = POOL_FAILURE_TYPE_WRITE_FAILED
        default:
            *poolFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into PoolFailureType", str))
    }
    return nil
}

func (poolFailureType PoolFailureType) String() string {
    switch poolFailureType {
        case POOL_FAILURE_TYPE_BLOB_READ_FAILED: return "BLOB_READ_FAILED"
        case POOL_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE: return "DATA_CHECKPOINT_FAILURE"
        case POOL_FAILURE_TYPE_DATA_CHECKPOINT_MISSING: return "DATA_CHECKPOINT_MISSING"
        case POOL_FAILURE_TYPE_FORMAT_FAILED: return "FORMAT_FAILED"
        case POOL_FAILURE_TYPE_IMPORT_FAILED: return "IMPORT_FAILED"
        case POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE: return "IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE"
        case POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY: return "IMPORT_FAILED_DUE_TO_DATA_INTEGRITY"
        case POOL_FAILURE_TYPE_INSPECT_FAILED: return "INSPECT_FAILED"
        case POOL_FAILURE_TYPE_QUIESCED: return "QUIESCED"
        case POOL_FAILURE_TYPE_READ_FAILED: return "READ_FAILED"
        case POOL_FAILURE_TYPE_VERIFY_FAILED: return "VERIFY_FAILED"
        case POOL_FAILURE_TYPE_WRITE_FAILED: return "WRITE_FAILED"
        default:
            log.Printf("Error: invalid PoolFailureType represented by '%d'", poolFailureType)
            return ""
    }
}
