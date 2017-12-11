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

type FeatureKeyType Enum

const (
    FEATURE_KEY_TYPE_AWS_S3_CLOUD_OUT FeatureKeyType = 1 + iota
    FEATURE_KEY_TYPE_MICROSOFT_AZURE_CLOUD_OUT FeatureKeyType = 1 + iota
)

func (featureKeyType *FeatureKeyType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *featureKeyType = UNDEFINED
        case "AWS_S3_CLOUD_OUT": *featureKeyType = FEATURE_KEY_TYPE_AWS_S3_CLOUD_OUT
        case "MICROSOFT_AZURE_CLOUD_OUT": *featureKeyType = FEATURE_KEY_TYPE_MICROSOFT_AZURE_CLOUD_OUT
        default:
            *featureKeyType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into FeatureKeyType", str))
    }
    return nil
}

func (featureKeyType FeatureKeyType) String() string {
    switch featureKeyType {
        case FEATURE_KEY_TYPE_AWS_S3_CLOUD_OUT: return "AWS_S3_CLOUD_OUT"
        case FEATURE_KEY_TYPE_MICROSOFT_AZURE_CLOUD_OUT: return "MICROSOFT_AZURE_CLOUD_OUT"
        default:
            log.Printf("Error: invalid FeatureKeyType represented by '%d'", featureKeyType)
            return ""
    }
}

func (featureKeyType FeatureKeyType) StringPtr() *string {
    if featureKeyType == UNDEFINED {
        return nil
    }
    result := featureKeyType.String()
    return &result
}