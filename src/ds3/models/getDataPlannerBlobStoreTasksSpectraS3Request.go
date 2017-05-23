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

type GetDataPlannerBlobStoreTasksSpectraS3Request struct {
    queryParams *url.Values
}

func NewGetDataPlannerBlobStoreTasksSpectraS3Request() *GetDataPlannerBlobStoreTasksSpectraS3Request {
    queryParams := &url.Values{}

    return &GetDataPlannerBlobStoreTasksSpectraS3Request{
        queryParams: queryParams,
    }
}



func (getDataPlannerBlobStoreTasksSpectraS3Request *GetDataPlannerBlobStoreTasksSpectraS3Request) WithFullDetails() *GetDataPlannerBlobStoreTasksSpectraS3Request {
    getDataPlannerBlobStoreTasksSpectraS3Request.queryParams.Set("full_details", "")
    return getDataPlannerBlobStoreTasksSpectraS3Request
}

func (GetDataPlannerBlobStoreTasksSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getDataPlannerBlobStoreTasksSpectraS3Request *GetDataPlannerBlobStoreTasksSpectraS3Request) Path() string {
    return "/_rest_/blob_store_task"
}

func (getDataPlannerBlobStoreTasksSpectraS3Request *GetDataPlannerBlobStoreTasksSpectraS3Request) QueryParams() *url.Values {
    return getDataPlannerBlobStoreTasksSpectraS3Request.queryParams
}

func (GetDataPlannerBlobStoreTasksSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetDataPlannerBlobStoreTasksSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetDataPlannerBlobStoreTasksSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
