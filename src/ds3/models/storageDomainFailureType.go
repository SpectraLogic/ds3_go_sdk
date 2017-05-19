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

type StorageDomainFailureType Enum

const (
    STORAGE_DOMAIN_FAILURE_TYPE_ILLEGAL_EJECTION_OCCURRED StorageDomainFailureType = 1 + iota
    STORAGE_DOMAIN_FAILURE_TYPE_MEMBER_BECAME_READ_ONLY StorageDomainFailureType = 1 + iota
    STORAGE_DOMAIN_FAILURE_TYPE_WRITES_STALLED_DUE_TO_NO_FREE_MEDIA_REMAINING StorageDomainFailureType = 1 + iota
)

func (storageDomainFailureType *StorageDomainFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "ILLEGAL_EJECTION_OCCURRED": *storageDomainFailureType = STORAGE_DOMAIN_FAILURE_TYPE_ILLEGAL_EJECTION_OCCURRED
        case "MEMBER_BECAME_READ_ONLY": *storageDomainFailureType = STORAGE_DOMAIN_FAILURE_TYPE_MEMBER_BECAME_READ_ONLY
        case "WRITES_STALLED_DUE_TO_NO_FREE_MEDIA_REMAINING": *storageDomainFailureType = STORAGE_DOMAIN_FAILURE_TYPE_WRITES_STALLED_DUE_TO_NO_FREE_MEDIA_REMAINING
        default:
            *storageDomainFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into StorageDomainFailureType", str))
    }
    return nil
}

func (storageDomainFailureType StorageDomainFailureType) String() string {
    switch storageDomainFailureType {
        case STORAGE_DOMAIN_FAILURE_TYPE_ILLEGAL_EJECTION_OCCURRED: return "ILLEGAL_EJECTION_OCCURRED"
        case STORAGE_DOMAIN_FAILURE_TYPE_MEMBER_BECAME_READ_ONLY: return "MEMBER_BECAME_READ_ONLY"
        case STORAGE_DOMAIN_FAILURE_TYPE_WRITES_STALLED_DUE_TO_NO_FREE_MEDIA_REMAINING: return "WRITES_STALLED_DUE_TO_NO_FREE_MEDIA_REMAINING"
        default:
            log.Printf("Error: invalid StorageDomainFailureType represented by '%d'", storageDomainFailureType)
            return ""
    }
}
