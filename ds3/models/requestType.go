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

type RequestType Enum

const (
    REQUEST_TYPE_DELETE RequestType = 1 + iota
    REQUEST_TYPE_GET RequestType = 1 + iota
    REQUEST_TYPE_HEAD RequestType = 1 + iota
    REQUEST_TYPE_POST RequestType = 1 + iota
    REQUEST_TYPE_PUT RequestType = 1 + iota
)

func (requestType *RequestType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *requestType = UNDEFINED
        case "DELETE": *requestType = REQUEST_TYPE_DELETE
        case "GET": *requestType = REQUEST_TYPE_GET
        case "HEAD": *requestType = REQUEST_TYPE_HEAD
        case "POST": *requestType = REQUEST_TYPE_POST
        case "PUT": *requestType = REQUEST_TYPE_PUT
        default:
            *requestType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into RequestType", str))
    }
    return nil
}

func (requestType RequestType) String() string {
    switch requestType {
        case REQUEST_TYPE_DELETE: return "DELETE"
        case REQUEST_TYPE_GET: return "GET"
        case REQUEST_TYPE_HEAD: return "HEAD"
        case REQUEST_TYPE_POST: return "POST"
        case REQUEST_TYPE_PUT: return "PUT"
        default:
            log.Printf("Error: invalid RequestType represented by '%d'", requestType)
            return ""
    }
}

func (requestType RequestType) StringPtr() *string {
    if requestType == UNDEFINED {
        return nil
    }
    result := requestType.String()
    return &result
}