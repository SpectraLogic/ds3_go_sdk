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

type VerifyDs3TargetSpectraS3Request struct {
    Ds3Target string
    FullDetails bool
}

func NewVerifyDs3TargetSpectraS3Request(ds3Target string) *VerifyDs3TargetSpectraS3Request {
    return &VerifyDs3TargetSpectraS3Request{
        Ds3Target: ds3Target,
    }
}

func (verifyDs3TargetSpectraS3Request *VerifyDs3TargetSpectraS3Request) WithFullDetails() *VerifyDs3TargetSpectraS3Request {
    verifyDs3TargetSpectraS3Request.FullDetails = true
    return verifyDs3TargetSpectraS3Request
}

