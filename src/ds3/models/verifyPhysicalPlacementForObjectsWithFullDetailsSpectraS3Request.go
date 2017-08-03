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

type VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request struct {
    bucketName string
    content networking.ReaderWithSizeDecorator
    storageDomainId string
    queryParams *url.Values
}

//TODO update request payload type to []string and conversion to stream
func NewVerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request(bucketName string, objectNames []string) *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "verify_physical_placement")
    queryParams.Set("full_details", "")

    return &VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request{
        bucketName: bucketName,
        content: buildDs3ObjectStreamFromNames(objectNames),
        queryParams: queryParams,
    }
}

func (verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) WithStorageDomainId(storageDomainId string) *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request {
    verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.storageDomainId = storageDomainId
    verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.queryParams.Set("storage_domain_id", storageDomainId)
    return verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request
}



func (VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) Path() string {
    return "/_rest_/bucket/" + verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.bucketName
}

func (verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) QueryParams() *url.Values {
    return verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.queryParams
}

func (VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.content
}
