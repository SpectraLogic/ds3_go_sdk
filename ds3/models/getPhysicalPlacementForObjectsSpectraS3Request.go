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

type GetPhysicalPlacementForObjectsSpectraS3Request struct {
    BucketName string
    ObjectNames []string
    StorageDomainId *string
}

func NewGetPhysicalPlacementForObjectsSpectraS3Request(bucketName string, objectNames []string) *GetPhysicalPlacementForObjectsSpectraS3Request {
    return &GetPhysicalPlacementForObjectsSpectraS3Request{
        BucketName: bucketName,
        ObjectNames: objectNames,
    }
}

func (getPhysicalPlacementForObjectsSpectraS3Request *GetPhysicalPlacementForObjectsSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetPhysicalPlacementForObjectsSpectraS3Request {
    getPhysicalPlacementForObjectsSpectraS3Request.StorageDomainId = &storageDomainId
    return getPhysicalPlacementForObjectsSpectraS3Request
}

