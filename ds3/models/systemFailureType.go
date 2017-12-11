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

type SystemFailureType Enum

const (
    SYSTEM_FAILURE_TYPE_RECONCILE_TAPE_ENVIRONMENT_FAILED SystemFailureType = 1 + iota
    SYSTEM_FAILURE_TYPE_RECONCILE_POOL_ENVIRONMENT_FAILED SystemFailureType = 1 + iota
    SYSTEM_FAILURE_TYPE_CRITICAL_DATA_VERIFICATION_ERROR_REQUIRES_USER_CONFIRMATION SystemFailureType = 1 + iota
    SYSTEM_FAILURE_TYPE_MICROSOFT_AZURE_WRITES_REQUIRE_FEATURE_LICENSE SystemFailureType = 1 + iota
    SYSTEM_FAILURE_TYPE_AWS_S3_WRITES_REQUIRE_FEATURE_LICENSE SystemFailureType = 1 + iota
    SYSTEM_FAILURE_TYPE_DATABASE_RUNNING_OUT_OF_SPACE SystemFailureType = 1 + iota
)

func (systemFailureType *SystemFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *systemFailureType = UNDEFINED
        case "RECONCILE_TAPE_ENVIRONMENT_FAILED": *systemFailureType = SYSTEM_FAILURE_TYPE_RECONCILE_TAPE_ENVIRONMENT_FAILED
        case "RECONCILE_POOL_ENVIRONMENT_FAILED": *systemFailureType = SYSTEM_FAILURE_TYPE_RECONCILE_POOL_ENVIRONMENT_FAILED
        case "CRITICAL_DATA_VERIFICATION_ERROR_REQUIRES_USER_CONFIRMATION": *systemFailureType = SYSTEM_FAILURE_TYPE_CRITICAL_DATA_VERIFICATION_ERROR_REQUIRES_USER_CONFIRMATION
        case "MICROSOFT_AZURE_WRITES_REQUIRE_FEATURE_LICENSE": *systemFailureType = SYSTEM_FAILURE_TYPE_MICROSOFT_AZURE_WRITES_REQUIRE_FEATURE_LICENSE
        case "AWS_S3_WRITES_REQUIRE_FEATURE_LICENSE": *systemFailureType = SYSTEM_FAILURE_TYPE_AWS_S3_WRITES_REQUIRE_FEATURE_LICENSE
        case "DATABASE_RUNNING_OUT_OF_SPACE": *systemFailureType = SYSTEM_FAILURE_TYPE_DATABASE_RUNNING_OUT_OF_SPACE
        default:
            *systemFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into SystemFailureType", str))
    }
    return nil
}

func (systemFailureType SystemFailureType) String() string {
    switch systemFailureType {
        case SYSTEM_FAILURE_TYPE_RECONCILE_TAPE_ENVIRONMENT_FAILED: return "RECONCILE_TAPE_ENVIRONMENT_FAILED"
        case SYSTEM_FAILURE_TYPE_RECONCILE_POOL_ENVIRONMENT_FAILED: return "RECONCILE_POOL_ENVIRONMENT_FAILED"
        case SYSTEM_FAILURE_TYPE_CRITICAL_DATA_VERIFICATION_ERROR_REQUIRES_USER_CONFIRMATION: return "CRITICAL_DATA_VERIFICATION_ERROR_REQUIRES_USER_CONFIRMATION"
        case SYSTEM_FAILURE_TYPE_MICROSOFT_AZURE_WRITES_REQUIRE_FEATURE_LICENSE: return "MICROSOFT_AZURE_WRITES_REQUIRE_FEATURE_LICENSE"
        case SYSTEM_FAILURE_TYPE_AWS_S3_WRITES_REQUIRE_FEATURE_LICENSE: return "AWS_S3_WRITES_REQUIRE_FEATURE_LICENSE"
        case SYSTEM_FAILURE_TYPE_DATABASE_RUNNING_OUT_OF_SPACE: return "DATABASE_RUNNING_OUT_OF_SPACE"
        default:
            log.Printf("Error: invalid SystemFailureType represented by '%d'", systemFailureType)
            return ""
    }
}

func (systemFailureType SystemFailureType) StringPtr() *string {
    if systemFailureType == UNDEFINED {
        return nil
    }
    result := systemFailureType.String()
    return &result
}