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

type ClearSuspectBlobS3TargetsSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewClearSuspectBlobS3TargetsSpectraS3Request(ids []string) *ClearSuspectBlobS3TargetsSpectraS3Request {
    return &ClearSuspectBlobS3TargetsSpectraS3Request{
        Ids: ids,
    }
}

func (clearSuspectBlobS3TargetsSpectraS3Request *ClearSuspectBlobS3TargetsSpectraS3Request) WithForce() *ClearSuspectBlobS3TargetsSpectraS3Request {
    clearSuspectBlobS3TargetsSpectraS3Request.Force = true
    return clearSuspectBlobS3TargetsSpectraS3Request
}

