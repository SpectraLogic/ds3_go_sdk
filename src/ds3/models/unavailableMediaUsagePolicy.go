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

type UnavailableMediaUsagePolicy Enum

const (
    UNAVAILABLE_MEDIA_USAGE_POLICY_ALLOW UnavailableMediaUsagePolicy = 1 + iota
    UNAVAILABLE_MEDIA_USAGE_POLICY_DISCOURAGED UnavailableMediaUsagePolicy = 1 + iota
    UNAVAILABLE_MEDIA_USAGE_POLICY_DISALLOW UnavailableMediaUsagePolicy = 1 + iota
)

func (unavailableMediaUsagePolicy *UnavailableMediaUsagePolicy) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "ALLOW": *unavailableMediaUsagePolicy = UNAVAILABLE_MEDIA_USAGE_POLICY_ALLOW
        case "DISCOURAGED": *unavailableMediaUsagePolicy = UNAVAILABLE_MEDIA_USAGE_POLICY_DISCOURAGED
        case "DISALLOW": *unavailableMediaUsagePolicy = UNAVAILABLE_MEDIA_USAGE_POLICY_DISALLOW
        default:
            *unavailableMediaUsagePolicy = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into UnavailableMediaUsagePolicy", str))
    }
    return nil
}

func (unavailableMediaUsagePolicy UnavailableMediaUsagePolicy) String() string {
    switch unavailableMediaUsagePolicy {
        case UNAVAILABLE_MEDIA_USAGE_POLICY_ALLOW: return "ALLOW"
        case UNAVAILABLE_MEDIA_USAGE_POLICY_DISCOURAGED: return "DISCOURAGED"
        case UNAVAILABLE_MEDIA_USAGE_POLICY_DISALLOW: return "DISALLOW"
        default:
            log.Printf("Error: invalid UnavailableMediaUsagePolicy represented by '%d'", unavailableMediaUsagePolicy)
            return ""
    }
}
