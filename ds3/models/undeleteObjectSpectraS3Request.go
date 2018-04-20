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

type UndeleteObjectSpectraS3Request struct {
    BucketId string
    Name string
    VersionId *string
}

func NewUndeleteObjectSpectraS3Request(bucketId string, name string) *UndeleteObjectSpectraS3Request {
    return &UndeleteObjectSpectraS3Request{
        BucketId: bucketId,
        Name: name,
    }
}

func (undeleteObjectSpectraS3Request *UndeleteObjectSpectraS3Request) WithVersionId(versionId string) *UndeleteObjectSpectraS3Request {
    undeleteObjectSpectraS3Request.VersionId = &versionId
    return undeleteObjectSpectraS3Request
}

