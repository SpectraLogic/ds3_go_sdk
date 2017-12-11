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

type ModifyNodeSpectraS3Request struct {
    DnsName *string
    Name *string
    Node string
}

func NewModifyNodeSpectraS3Request(node string) *ModifyNodeSpectraS3Request {
    return &ModifyNodeSpectraS3Request{
        Node: node,
    }
}

func (modifyNodeSpectraS3Request *ModifyNodeSpectraS3Request) WithDnsName(dnsName string) *ModifyNodeSpectraS3Request {
    modifyNodeSpectraS3Request.DnsName = &dnsName
    return modifyNodeSpectraS3Request
}

func (modifyNodeSpectraS3Request *ModifyNodeSpectraS3Request) WithName(name string) *ModifyNodeSpectraS3Request {
    modifyNodeSpectraS3Request.Name = &name
    return modifyNodeSpectraS3Request
}

