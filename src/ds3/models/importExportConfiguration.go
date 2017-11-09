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

type ImportExportConfiguration Enum

const (
    IMPORT_EXPORT_CONFIGURATION_SUPPORTED ImportExportConfiguration = 1 + iota
    IMPORT_EXPORT_CONFIGURATION_NOT_SUPPORTED ImportExportConfiguration = 1 + iota
)

func (importExportConfiguration *ImportExportConfiguration) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *importExportConfiguration = UNDEFINED
        case "SUPPORTED": *importExportConfiguration = IMPORT_EXPORT_CONFIGURATION_SUPPORTED
        case "NOT_SUPPORTED": *importExportConfiguration = IMPORT_EXPORT_CONFIGURATION_NOT_SUPPORTED
        default:
            *importExportConfiguration = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into ImportExportConfiguration", str))
    }
    return nil
}

func (importExportConfiguration ImportExportConfiguration) String() string {
    switch importExportConfiguration {
        case IMPORT_EXPORT_CONFIGURATION_SUPPORTED: return "SUPPORTED"
        case IMPORT_EXPORT_CONFIGURATION_NOT_SUPPORTED: return "NOT_SUPPORTED"
        default:
            log.Printf("Error: invalid ImportExportConfiguration represented by '%d'", importExportConfiguration)
            return ""
    }
}

func (importExportConfiguration ImportExportConfiguration) StringPtr() *string {
    if importExportConfiguration == UNDEFINED {
        return nil
    }
    result := importExportConfiguration.String()
    return &result
}