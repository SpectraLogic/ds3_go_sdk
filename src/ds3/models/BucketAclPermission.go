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

type BucketAclPermission Enum

const (
    BUCKET_ACL_PERMISSION_LIST BucketAclPermission = 1 + iota
    BUCKET_ACL_PERMISSION_READ BucketAclPermission = 1 + iota
    BUCKET_ACL_PERMISSION_WRITE BucketAclPermission = 1 + iota
    BUCKET_ACL_PERMISSION_DELETE BucketAclPermission = 1 + iota
    BUCKET_ACL_PERMISSION_JOB BucketAclPermission = 1 + iota
    BUCKET_ACL_PERMISSION_OWNER BucketAclPermission = 1 + iota
)

func (bucketAclPermission *BucketAclPermission) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "LIST": *bucketAclPermission = BUCKET_ACL_PERMISSION_LIST
        case "READ": *bucketAclPermission = BUCKET_ACL_PERMISSION_READ
        case "WRITE": *bucketAclPermission = BUCKET_ACL_PERMISSION_WRITE
        case "DELETE": *bucketAclPermission = BUCKET_ACL_PERMISSION_DELETE
        case "JOB": *bucketAclPermission = BUCKET_ACL_PERMISSION_JOB
        case "OWNER": *bucketAclPermission = BUCKET_ACL_PERMISSION_OWNER
        default:
            *bucketAclPermission = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into BucketAclPermission", str))
    }
    return nil
}

func (bucketAclPermission BucketAclPermission) String() string {
    switch bucketAclPermission {
        case BUCKET_ACL_PERMISSION_LIST: return "LIST"
        case BUCKET_ACL_PERMISSION_READ: return "READ"
        case BUCKET_ACL_PERMISSION_WRITE: return "WRITE"
        case BUCKET_ACL_PERMISSION_DELETE: return "DELETE"
        case BUCKET_ACL_PERMISSION_JOB: return "JOB"
        case BUCKET_ACL_PERMISSION_OWNER: return "OWNER"
        default:
            log.Printf("Error: invalid BucketAclPermission represented by '%d'", bucketAclPermission)
            return ""
    }
}
