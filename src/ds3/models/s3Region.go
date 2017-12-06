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

type S3Region Enum

const (
    S3_REGION_GOV_CLOUD S3Region = 1 + iota
    S3_REGION_US_EAST_1 S3Region = 1 + iota
    S3_REGION_US_WEST_1 S3Region = 1 + iota
    S3_REGION_US_WEST_2 S3Region = 1 + iota
    S3_REGION_EU_WEST_1 S3Region = 1 + iota
    S3_REGION_EU_CENTRAL_1 S3Region = 1 + iota
    S3_REGION_AP_SOUTH_1 S3Region = 1 + iota
    S3_REGION_AP_SOUTHEAST_1 S3Region = 1 + iota
    S3_REGION_AP_SOUTHEAST_2 S3Region = 1 + iota
    S3_REGION_AP_NORTHEAST_1 S3Region = 1 + iota
    S3_REGION_AP_NORTHEAST_2 S3Region = 1 + iota
    S3_REGION_SA_EAST_1 S3Region = 1 + iota
    S3_REGION_CN_NORTH_1 S3Region = 1 + iota
)

func (s3Region *S3Region) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *s3Region = UNDEFINED
        case "GOV_CLOUD": *s3Region = S3_REGION_GOV_CLOUD
        case "US_EAST_1": *s3Region = S3_REGION_US_EAST_1
        case "US_WEST_1": *s3Region = S3_REGION_US_WEST_1
        case "US_WEST_2": *s3Region = S3_REGION_US_WEST_2
        case "EU_WEST_1": *s3Region = S3_REGION_EU_WEST_1
        case "EU_CENTRAL_1": *s3Region = S3_REGION_EU_CENTRAL_1
        case "AP_SOUTH_1": *s3Region = S3_REGION_AP_SOUTH_1
        case "AP_SOUTHEAST_1": *s3Region = S3_REGION_AP_SOUTHEAST_1
        case "AP_SOUTHEAST_2": *s3Region = S3_REGION_AP_SOUTHEAST_2
        case "AP_NORTHEAST_1": *s3Region = S3_REGION_AP_NORTHEAST_1
        case "AP_NORTHEAST_2": *s3Region = S3_REGION_AP_NORTHEAST_2
        case "SA_EAST_1": *s3Region = S3_REGION_SA_EAST_1
        case "CN_NORTH_1": *s3Region = S3_REGION_CN_NORTH_1
        default:
            *s3Region = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into S3Region", str))
    }
    return nil
}

func (s3Region S3Region) String() string {
    switch s3Region {
        case S3_REGION_GOV_CLOUD: return "GOV_CLOUD"
        case S3_REGION_US_EAST_1: return "US_EAST_1"
        case S3_REGION_US_WEST_1: return "US_WEST_1"
        case S3_REGION_US_WEST_2: return "US_WEST_2"
        case S3_REGION_EU_WEST_1: return "EU_WEST_1"
        case S3_REGION_EU_CENTRAL_1: return "EU_CENTRAL_1"
        case S3_REGION_AP_SOUTH_1: return "AP_SOUTH_1"
        case S3_REGION_AP_SOUTHEAST_1: return "AP_SOUTHEAST_1"
        case S3_REGION_AP_SOUTHEAST_2: return "AP_SOUTHEAST_2"
        case S3_REGION_AP_NORTHEAST_1: return "AP_NORTHEAST_1"
        case S3_REGION_AP_NORTHEAST_2: return "AP_NORTHEAST_2"
        case S3_REGION_SA_EAST_1: return "SA_EAST_1"
        case S3_REGION_CN_NORTH_1: return "CN_NORTH_1"
        default:
            log.Printf("Error: invalid S3Region represented by '%d'", s3Region)
            return ""
    }
}

func (s3Region S3Region) StringPtr() *string {
    if s3Region == UNDEFINED {
        return nil
    }
    result := s3Region.String()
    return &result
}