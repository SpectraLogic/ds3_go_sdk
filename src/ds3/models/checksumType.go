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

type ChecksumType Enum

const (
    CHECKSUM_TYPE_CRC_32 ChecksumType = 1 + iota
    CHECKSUM_TYPE_CRC_32C ChecksumType = 1 + iota
    CHECKSUM_TYPE_MD5 ChecksumType = 1 + iota
    CHECKSUM_TYPE_SHA_256 ChecksumType = 1 + iota
    CHECKSUM_TYPE_SHA_512 ChecksumType = 1 + iota
)

func (checksumType *ChecksumType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *checksumType = UNDEFINED
        case "CRC_32": *checksumType = CHECKSUM_TYPE_CRC_32
        case "CRC_32C": *checksumType = CHECKSUM_TYPE_CRC_32C
        case "MD5": *checksumType = CHECKSUM_TYPE_MD5
        case "SHA_256": *checksumType = CHECKSUM_TYPE_SHA_256
        case "SHA_512": *checksumType = CHECKSUM_TYPE_SHA_512
        default:
            *checksumType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into ChecksumType", str))
    }
    return nil
}

func (checksumType ChecksumType) String() string {
    switch checksumType {
        case CHECKSUM_TYPE_CRC_32: return "CRC_32"
        case CHECKSUM_TYPE_CRC_32C: return "CRC_32C"
        case CHECKSUM_TYPE_MD5: return "MD5"
        case CHECKSUM_TYPE_SHA_256: return "SHA_256"
        case CHECKSUM_TYPE_SHA_512: return "SHA_512"
        default:
            log.Printf("Error: invalid ChecksumType represented by '%d'", checksumType)
            return ""
    }
}
