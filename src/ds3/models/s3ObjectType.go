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

type S3ObjectType Enum

const (
    S3_OBJECT_TYPE_DATA S3ObjectType = 1 + iota
    S3_OBJECT_TYPE_FOLDER S3ObjectType = 1 + iota
)

func (s3ObjectType *S3ObjectType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "DATA": *s3ObjectType = S3_OBJECT_TYPE_DATA
        case "FOLDER": *s3ObjectType = S3_OBJECT_TYPE_FOLDER
        default:
            *s3ObjectType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into S3ObjectType", str))
    }
    return nil
}

func (s3ObjectType S3ObjectType) String() string {
    switch s3ObjectType {
        case S3_OBJECT_TYPE_DATA: return "DATA"
        case S3_OBJECT_TYPE_FOLDER: return "FOLDER"
        default:
            log.Printf("Error: invalid S3ObjectType represented by '%d'", s3ObjectType)
            return ""
    }
}
