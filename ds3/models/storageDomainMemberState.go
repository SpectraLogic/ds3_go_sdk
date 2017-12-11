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

type StorageDomainMemberState Enum

const (
    STORAGE_DOMAIN_MEMBER_STATE_NORMAL StorageDomainMemberState = 1 + iota
    STORAGE_DOMAIN_MEMBER_STATE_EXCLUSION_IN_PROGRESS StorageDomainMemberState = 1 + iota
)

func (storageDomainMemberState *StorageDomainMemberState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *storageDomainMemberState = UNDEFINED
        case "NORMAL": *storageDomainMemberState = STORAGE_DOMAIN_MEMBER_STATE_NORMAL
        case "EXCLUSION_IN_PROGRESS": *storageDomainMemberState = STORAGE_DOMAIN_MEMBER_STATE_EXCLUSION_IN_PROGRESS
        default:
            *storageDomainMemberState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into StorageDomainMemberState", str))
    }
    return nil
}

func (storageDomainMemberState StorageDomainMemberState) String() string {
    switch storageDomainMemberState {
        case STORAGE_DOMAIN_MEMBER_STATE_NORMAL: return "NORMAL"
        case STORAGE_DOMAIN_MEMBER_STATE_EXCLUSION_IN_PROGRESS: return "EXCLUSION_IN_PROGRESS"
        default:
            log.Printf("Error: invalid StorageDomainMemberState represented by '%d'", storageDomainMemberState)
            return ""
    }
}

func (storageDomainMemberState StorageDomainMemberState) StringPtr() *string {
    if storageDomainMemberState == UNDEFINED {
        return nil
    }
    result := storageDomainMemberState.String()
    return &result
}