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

type GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request struct {
    bucketName string
    content networking.ReaderWithSizeDecorator
    storageDomainId string
    queryParams *url.Values
}

func NewGetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request(bucketName string, objectNames []string) *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "get_physical_placement")
    queryParams.Set("full_details", "")

    return &GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request{
        bucketName: bucketName,
        content: buildDs3ObjectStreamFromNames(objectNames),
        queryParams: queryParams,
    }
}

func (getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request {
    getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.storageDomainId = storageDomainId
    getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request
}



func (GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) Path() string {
    return "/_rest_/bucket/" + getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.bucketName
}

func (getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) QueryParams() *url.Values {
    return getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.queryParams
}

func (GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.content
}
