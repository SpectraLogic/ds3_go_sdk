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
    "net/url"
    "net/http"
    "ds3/networking"
)

type GetPhysicalPlacementForObjectsSpectraS3Request struct {
    bucketName string
    content networking.ReaderWithSizeDecorator
    storageDomainId string
    queryParams *url.Values
}

func NewGetPhysicalPlacementForObjectsSpectraS3Request(bucketName string, objects []Ds3Object) *GetPhysicalPlacementForObjectsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "get_physical_placement")

    return &GetPhysicalPlacementForObjectsSpectraS3Request{
        bucketName: bucketName,
        content: buildDs3ObjectListStream(objects),
        queryParams: queryParams,
    }
}

func (getPhysicalPlacementForObjectsSpectraS3Request *GetPhysicalPlacementForObjectsSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetPhysicalPlacementForObjectsSpectraS3Request {
    getPhysicalPlacementForObjectsSpectraS3Request.storageDomainId = storageDomainId
    getPhysicalPlacementForObjectsSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return getPhysicalPlacementForObjectsSpectraS3Request
}



func (GetPhysicalPlacementForObjectsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (getPhysicalPlacementForObjectsSpectraS3Request *GetPhysicalPlacementForObjectsSpectraS3Request) Path() string {
    return "/_rest_/bucket/" + getPhysicalPlacementForObjectsSpectraS3Request.bucketName
}

func (getPhysicalPlacementForObjectsSpectraS3Request *GetPhysicalPlacementForObjectsSpectraS3Request) QueryParams() *url.Values {
    return getPhysicalPlacementForObjectsSpectraS3Request.queryParams
}

func (GetPhysicalPlacementForObjectsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetPhysicalPlacementForObjectsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (getPhysicalPlacementForObjectsSpectraS3Request *GetPhysicalPlacementForObjectsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return getPhysicalPlacementForObjectsSpectraS3Request.content
}
