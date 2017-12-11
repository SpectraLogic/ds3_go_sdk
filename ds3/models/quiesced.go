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

type Quiesced Enum

const (
    QUIESCED_NO Quiesced = 1 + iota
    QUIESCED_PENDING Quiesced = 1 + iota
    QUIESCED_YES Quiesced = 1 + iota
)

func (quiesced *Quiesced) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *quiesced = UNDEFINED
        case "NO": *quiesced = QUIESCED_NO
        case "PENDING": *quiesced = QUIESCED_PENDING
        case "YES": *quiesced = QUIESCED_YES
        default:
            *quiesced = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into Quiesced", str))
    }
    return nil
}

func (quiesced Quiesced) String() string {
    switch quiesced {
        case QUIESCED_NO: return "NO"
        case QUIESCED_PENDING: return "PENDING"
        case QUIESCED_YES: return "YES"
        default:
            log.Printf("Error: invalid Quiesced represented by '%d'", quiesced)
            return ""
    }
}

func (quiesced Quiesced) StringPtr() *string {
    if quiesced == UNDEFINED {
        return nil
    }
    result := quiesced.String()
    return &result
}