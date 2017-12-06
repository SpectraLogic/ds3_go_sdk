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

type S3InitialDataPlacementPolicy Enum

const (
    S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD S3InitialDataPlacementPolicy = 1 + iota
    S3_INITIAL_DATA_PLACEMENT_POLICY_REDUCED_REDUNDANCY S3InitialDataPlacementPolicy = 1 + iota
    S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD_IA S3InitialDataPlacementPolicy = 1 + iota
    S3_INITIAL_DATA_PLACEMENT_POLICY_GLACIER S3InitialDataPlacementPolicy = 1 + iota
)

func (s3InitialDataPlacementPolicy *S3InitialDataPlacementPolicy) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *s3InitialDataPlacementPolicy = UNDEFINED
        case "STANDARD": *s3InitialDataPlacementPolicy = S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD
        case "REDUCED_REDUNDANCY": *s3InitialDataPlacementPolicy = S3_INITIAL_DATA_PLACEMENT_POLICY_REDUCED_REDUNDANCY
        case "STANDARD_IA": *s3InitialDataPlacementPolicy = S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD_IA
        case "GLACIER": *s3InitialDataPlacementPolicy = S3_INITIAL_DATA_PLACEMENT_POLICY_GLACIER
        default:
            *s3InitialDataPlacementPolicy = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into S3InitialDataPlacementPolicy", str))
    }
    return nil
}

func (s3InitialDataPlacementPolicy S3InitialDataPlacementPolicy) String() string {
    switch s3InitialDataPlacementPolicy {
        case S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD: return "STANDARD"
        case S3_INITIAL_DATA_PLACEMENT_POLICY_REDUCED_REDUNDANCY: return "REDUCED_REDUNDANCY"
        case S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD_IA: return "STANDARD_IA"
        case S3_INITIAL_DATA_PLACEMENT_POLICY_GLACIER: return "GLACIER"
        default:
            log.Printf("Error: invalid S3InitialDataPlacementPolicy represented by '%d'", s3InitialDataPlacementPolicy)
            return ""
    }
}

func (s3InitialDataPlacementPolicy S3InitialDataPlacementPolicy) StringPtr() *string {
    if s3InitialDataPlacementPolicy == UNDEFINED {
        return nil
    }
    result := s3InitialDataPlacementPolicy.String()
    return &result
}