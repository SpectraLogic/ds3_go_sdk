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

type HttpResponseFormatType Enum

const (
    HTTP_RESPONSE_FORMAT_TYPE_DEFAULT HttpResponseFormatType = 1 + iota
    HTTP_RESPONSE_FORMAT_TYPE_JSON HttpResponseFormatType = 1 + iota
    HTTP_RESPONSE_FORMAT_TYPE_XML HttpResponseFormatType = 1 + iota
)

func (httpResponseFormatType *HttpResponseFormatType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *httpResponseFormatType = UNDEFINED
        case "DEFAULT": *httpResponseFormatType = HTTP_RESPONSE_FORMAT_TYPE_DEFAULT
        case "JSON": *httpResponseFormatType = HTTP_RESPONSE_FORMAT_TYPE_JSON
        case "XML": *httpResponseFormatType = HTTP_RESPONSE_FORMAT_TYPE_XML
        default:
            *httpResponseFormatType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into HttpResponseFormatType", str))
    }
    return nil
}

func (httpResponseFormatType HttpResponseFormatType) String() string {
    switch httpResponseFormatType {
        case HTTP_RESPONSE_FORMAT_TYPE_DEFAULT: return "DEFAULT"
        case HTTP_RESPONSE_FORMAT_TYPE_JSON: return "JSON"
        case HTTP_RESPONSE_FORMAT_TYPE_XML: return "XML"
        default:
            log.Printf("Error: invalid HttpResponseFormatType represented by '%d'", httpResponseFormatType)
            return ""
    }
}
