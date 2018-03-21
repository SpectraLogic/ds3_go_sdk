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

type RestOperationType Enum

const (
    REST_OPERATION_TYPE_ALLOCATE RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CANCEL_EJECT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CANCEL_FORMAT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CANCEL_IMPORT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CANCEL_ONLINE RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CANCEL_VERIFY RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CLEAN RestOperationType = 1 + iota
    REST_OPERATION_TYPE_COMPACT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_DEALLOCATE RestOperationType = 1 + iota
    REST_OPERATION_TYPE_EJECT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_FORMAT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_GET_PHYSICAL_PLACEMENT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_IMPORT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_INSPECT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_ONLINE RestOperationType = 1 + iota
    REST_OPERATION_TYPE_PAIR_BACK RestOperationType = 1 + iota
    REST_OPERATION_TYPE_REGENERATE_SECRET_KEY RestOperationType = 1 + iota
    REST_OPERATION_TYPE_START_BULK_GET RestOperationType = 1 + iota
    REST_OPERATION_TYPE_START_BULK_PUT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_START_BULK_STAGE RestOperationType = 1 + iota
    REST_OPERATION_TYPE_START_BULK_VERIFY RestOperationType = 1 + iota
    REST_OPERATION_TYPE_VERIFY RestOperationType = 1 + iota
    REST_OPERATION_TYPE_VERIFY_SAFE_TO_START_BULK_PUT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_VERIFY_PHYSICAL_PLACEMENT RestOperationType = 1 + iota
)

func (restOperationType *RestOperationType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *restOperationType = UNDEFINED
        case "ALLOCATE": *restOperationType = REST_OPERATION_TYPE_ALLOCATE
        case "CANCEL_EJECT": *restOperationType = REST_OPERATION_TYPE_CANCEL_EJECT
        case "CANCEL_FORMAT": *restOperationType = REST_OPERATION_TYPE_CANCEL_FORMAT
        case "CANCEL_IMPORT": *restOperationType = REST_OPERATION_TYPE_CANCEL_IMPORT
        case "CANCEL_ONLINE": *restOperationType = REST_OPERATION_TYPE_CANCEL_ONLINE
        case "CANCEL_VERIFY": *restOperationType = REST_OPERATION_TYPE_CANCEL_VERIFY
        case "CLEAN": *restOperationType = REST_OPERATION_TYPE_CLEAN
        case "COMPACT": *restOperationType = REST_OPERATION_TYPE_COMPACT
        case "DEALLOCATE": *restOperationType = REST_OPERATION_TYPE_DEALLOCATE
        case "EJECT": *restOperationType = REST_OPERATION_TYPE_EJECT
        case "FORMAT": *restOperationType = REST_OPERATION_TYPE_FORMAT
        case "GET_PHYSICAL_PLACEMENT": *restOperationType = REST_OPERATION_TYPE_GET_PHYSICAL_PLACEMENT
        case "IMPORT": *restOperationType = REST_OPERATION_TYPE_IMPORT
        case "INSPECT": *restOperationType = REST_OPERATION_TYPE_INSPECT
        case "ONLINE": *restOperationType = REST_OPERATION_TYPE_ONLINE
        case "PAIR_BACK": *restOperationType = REST_OPERATION_TYPE_PAIR_BACK
        case "REGENERATE_SECRET_KEY": *restOperationType = REST_OPERATION_TYPE_REGENERATE_SECRET_KEY
        case "START_BULK_GET": *restOperationType = REST_OPERATION_TYPE_START_BULK_GET
        case "START_BULK_PUT": *restOperationType = REST_OPERATION_TYPE_START_BULK_PUT
        case "START_BULK_STAGE": *restOperationType = REST_OPERATION_TYPE_START_BULK_STAGE
        case "START_BULK_VERIFY": *restOperationType = REST_OPERATION_TYPE_START_BULK_VERIFY
        case "VERIFY": *restOperationType = REST_OPERATION_TYPE_VERIFY
        case "VERIFY_SAFE_TO_START_BULK_PUT": *restOperationType = REST_OPERATION_TYPE_VERIFY_SAFE_TO_START_BULK_PUT
        case "VERIFY_PHYSICAL_PLACEMENT": *restOperationType = REST_OPERATION_TYPE_VERIFY_PHYSICAL_PLACEMENT
        default:
            *restOperationType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into RestOperationType", str))
    }
    return nil
}

func (restOperationType RestOperationType) String() string {
    switch restOperationType {
        case REST_OPERATION_TYPE_ALLOCATE: return "ALLOCATE"
        case REST_OPERATION_TYPE_CANCEL_EJECT: return "CANCEL_EJECT"
        case REST_OPERATION_TYPE_CANCEL_FORMAT: return "CANCEL_FORMAT"
        case REST_OPERATION_TYPE_CANCEL_IMPORT: return "CANCEL_IMPORT"
        case REST_OPERATION_TYPE_CANCEL_ONLINE: return "CANCEL_ONLINE"
        case REST_OPERATION_TYPE_CANCEL_VERIFY: return "CANCEL_VERIFY"
        case REST_OPERATION_TYPE_CLEAN: return "CLEAN"
        case REST_OPERATION_TYPE_COMPACT: return "COMPACT"
        case REST_OPERATION_TYPE_DEALLOCATE: return "DEALLOCATE"
        case REST_OPERATION_TYPE_EJECT: return "EJECT"
        case REST_OPERATION_TYPE_FORMAT: return "FORMAT"
        case REST_OPERATION_TYPE_GET_PHYSICAL_PLACEMENT: return "GET_PHYSICAL_PLACEMENT"
        case REST_OPERATION_TYPE_IMPORT: return "IMPORT"
        case REST_OPERATION_TYPE_INSPECT: return "INSPECT"
        case REST_OPERATION_TYPE_ONLINE: return "ONLINE"
        case REST_OPERATION_TYPE_PAIR_BACK: return "PAIR_BACK"
        case REST_OPERATION_TYPE_REGENERATE_SECRET_KEY: return "REGENERATE_SECRET_KEY"
        case REST_OPERATION_TYPE_START_BULK_GET: return "START_BULK_GET"
        case REST_OPERATION_TYPE_START_BULK_PUT: return "START_BULK_PUT"
        case REST_OPERATION_TYPE_START_BULK_STAGE: return "START_BULK_STAGE"
        case REST_OPERATION_TYPE_START_BULK_VERIFY: return "START_BULK_VERIFY"
        case REST_OPERATION_TYPE_VERIFY: return "VERIFY"
        case REST_OPERATION_TYPE_VERIFY_SAFE_TO_START_BULK_PUT: return "VERIFY_SAFE_TO_START_BULK_PUT"
        case REST_OPERATION_TYPE_VERIFY_PHYSICAL_PLACEMENT: return "VERIFY_PHYSICAL_PLACEMENT"
        default:
            log.Printf("Error: invalid RestOperationType represented by '%d'", restOperationType)
            return ""
    }
}

func (restOperationType RestOperationType) StringPtr() *string {
    if restOperationType == UNDEFINED {
        return nil
    }
    result := restOperationType.String()
    return &result
}