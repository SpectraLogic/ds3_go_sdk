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

type VerifyS3TargetSpectraS3Request struct {
    FullDetails bool
    S3Target string
}

func NewVerifyS3TargetSpectraS3Request(s3Target string) *VerifyS3TargetSpectraS3Request {
    return &VerifyS3TargetSpectraS3Request{
        S3Target: s3Target,
    }
}

func (verifyS3TargetSpectraS3Request *VerifyS3TargetSpectraS3Request) WithFullDetails() *VerifyS3TargetSpectraS3Request {
    verifyS3TargetSpectraS3Request.FullDetails = true
    return verifyS3TargetSpectraS3Request
}

