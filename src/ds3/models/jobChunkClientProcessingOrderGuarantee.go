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

type JobChunkClientProcessingOrderGuarantee Enum

const (
    JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_NONE JobChunkClientProcessingOrderGuarantee = 1 + iota
    JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER JobChunkClientProcessingOrderGuarantee = 1 + iota
)

func (jobChunkClientProcessingOrderGuarantee *JobChunkClientProcessingOrderGuarantee) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "NONE": *jobChunkClientProcessingOrderGuarantee = JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_NONE
        case "IN_ORDER": *jobChunkClientProcessingOrderGuarantee = JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER
        default:
            *jobChunkClientProcessingOrderGuarantee = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into JobChunkClientProcessingOrderGuarantee", str))
    }
    return nil
}

func (jobChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) String() string {
    switch jobChunkClientProcessingOrderGuarantee {
        case JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_NONE: return "NONE"
        case JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER: return "IN_ORDER"
        default:
            log.Printf("Error: invalid JobChunkClientProcessingOrderGuarantee represented by '%d'", jobChunkClientProcessingOrderGuarantee)
            return ""
    }
}
