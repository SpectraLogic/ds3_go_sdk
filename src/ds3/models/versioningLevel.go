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

type VersioningLevel Enum

const (
    VERSIONING_LEVEL_NONE VersioningLevel = 1 + iota
    VERSIONING_LEVEL_KEEP_LATEST VersioningLevel = 1 + iota
)

func (versioningLevel *VersioningLevel) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "NONE": *versioningLevel = VERSIONING_LEVEL_NONE
        case "KEEP_LATEST": *versioningLevel = VERSIONING_LEVEL_KEEP_LATEST
        default:
            *versioningLevel = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into VersioningLevel", str))
    }
    return nil
}

func (versioningLevel VersioningLevel) String() string {
    switch versioningLevel {
        case VERSIONING_LEVEL_NONE: return "NONE"
        case VERSIONING_LEVEL_KEEP_LATEST: return "KEEP_LATEST"
        default:
            log.Printf("Error: invalid VersioningLevel represented by '%d'", versioningLevel)
            return ""
    }
}
