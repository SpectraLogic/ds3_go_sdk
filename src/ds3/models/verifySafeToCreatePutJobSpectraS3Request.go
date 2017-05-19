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

type VerifySafeToCreatePutJobSpectraS3Request struct {
    bucketName string
    queryParams *url.Values
}

func NewVerifySafeToCreatePutJobSpectraS3Request(bucketName string) *VerifySafeToCreatePutJobSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "verify_safe_to_start_bulk_put")

    return &VerifySafeToCreatePutJobSpectraS3Request{
        bucketName: bucketName,
        queryParams: queryParams,
    }
}




func (VerifySafeToCreatePutJobSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (verifySafeToCreatePutJobSpectraS3Request *VerifySafeToCreatePutJobSpectraS3Request) Path() string {
    return "/_rest_/bucket/" + verifySafeToCreatePutJobSpectraS3Request.bucketName
}

func (verifySafeToCreatePutJobSpectraS3Request *VerifySafeToCreatePutJobSpectraS3Request) QueryParams() *url.Values {
    return verifySafeToCreatePutJobSpectraS3Request.queryParams
}

func (VerifySafeToCreatePutJobSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (VerifySafeToCreatePutJobSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (VerifySafeToCreatePutJobSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
