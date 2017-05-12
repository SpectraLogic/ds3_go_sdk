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

type LtfsFileNamingMode Enum

const (
    LTFS_FILE_NAMING_MODE_OBJECT_NAME LtfsFileNamingMode = 1 + iota
    LTFS_FILE_NAMING_MODE_OBJECT_ID LtfsFileNamingMode = 1 + iota
)

func (ltfsFileNamingMode *LtfsFileNamingMode) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "OBJECT_NAME": *ltfsFileNamingMode = LTFS_FILE_NAMING_MODE_OBJECT_NAME
        case "OBJECT_ID": *ltfsFileNamingMode = LTFS_FILE_NAMING_MODE_OBJECT_ID
        default:
            *ltfsFileNamingMode = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into LtfsFileNamingMode", str))
    }
    return nil
}

func (ltfsFileNamingMode LtfsFileNamingMode) String() string {
    switch ltfsFileNamingMode {
        case LTFS_FILE_NAMING_MODE_OBJECT_NAME: return "OBJECT_NAME"
        case LTFS_FILE_NAMING_MODE_OBJECT_ID: return "OBJECT_ID"
        default:
            log.Printf("Error: invalid LtfsFileNamingMode represented by '%d'", ltfsFileNamingMode)
            return ""
    }
}
