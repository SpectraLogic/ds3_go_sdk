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

type ModifyGroupSpectraS3Request struct {
    Group string
    Name *string
}

func NewModifyGroupSpectraS3Request(group string) *ModifyGroupSpectraS3Request {
    return &ModifyGroupSpectraS3Request{
        Group: group,
    }
}

func (modifyGroupSpectraS3Request *ModifyGroupSpectraS3Request) WithName(name string) *ModifyGroupSpectraS3Request {
    modifyGroupSpectraS3Request.Name = &name
    return modifyGroupSpectraS3Request
}

