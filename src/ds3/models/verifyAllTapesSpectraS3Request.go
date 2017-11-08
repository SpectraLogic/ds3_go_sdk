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

type VerifyAllTapesSpectraS3Request struct {
    TaskPriority Priority
}

func NewVerifyAllTapesSpectraS3Request() *VerifyAllTapesSpectraS3Request {
    return &VerifyAllTapesSpectraS3Request{
    }
}

func (verifyAllTapesSpectraS3Request *VerifyAllTapesSpectraS3Request) WithTaskPriority(taskPriority Priority) *VerifyAllTapesSpectraS3Request {
    verifyAllTapesSpectraS3Request.TaskPriority = taskPriority
    return verifyAllTapesSpectraS3Request
}

