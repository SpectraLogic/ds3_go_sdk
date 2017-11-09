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

type NamingConventionType Enum

const (
    NAMING_CONVENTION_TYPE_CONCAT_LOWERCASE NamingConventionType = 1 + iota
    NAMING_CONVENTION_TYPE_CONSTANT NamingConventionType = 1 + iota
    NAMING_CONVENTION_TYPE_UNDERSCORED NamingConventionType = 1 + iota
    NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_UPPERCASE NamingConventionType = 1 + iota
    NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_LOWERCASE NamingConventionType = 1 + iota
)

func (namingConventionType *NamingConventionType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *namingConventionType = UNDEFINED
        case "CONCAT_LOWERCASE": *namingConventionType = NAMING_CONVENTION_TYPE_CONCAT_LOWERCASE
        case "CONSTANT": *namingConventionType = NAMING_CONVENTION_TYPE_CONSTANT
        case "UNDERSCORED": *namingConventionType = NAMING_CONVENTION_TYPE_UNDERSCORED
        case "CAMEL_CASE_WITH_FIRST_LETTER_UPPERCASE": *namingConventionType = NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_UPPERCASE
        case "CAMEL_CASE_WITH_FIRST_LETTER_LOWERCASE": *namingConventionType = NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_LOWERCASE
        default:
            *namingConventionType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into NamingConventionType", str))
    }
    return nil
}

func (namingConventionType NamingConventionType) String() string {
    switch namingConventionType {
        case NAMING_CONVENTION_TYPE_CONCAT_LOWERCASE: return "CONCAT_LOWERCASE"
        case NAMING_CONVENTION_TYPE_CONSTANT: return "CONSTANT"
        case NAMING_CONVENTION_TYPE_UNDERSCORED: return "UNDERSCORED"
        case NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_UPPERCASE: return "CAMEL_CASE_WITH_FIRST_LETTER_UPPERCASE"
        case NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_LOWERCASE: return "CAMEL_CASE_WITH_FIRST_LETTER_LOWERCASE"
        default:
            log.Printf("Error: invalid NamingConventionType represented by '%d'", namingConventionType)
            return ""
    }
}

func (namingConventionType NamingConventionType) StringPtr() *string {
    if namingConventionType == UNDEFINED {
        return nil
    }
    result := namingConventionType.String()
    return &result
}