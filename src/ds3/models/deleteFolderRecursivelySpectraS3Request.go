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

type DeleteFolderRecursivelySpectraS3Request struct {
    bucketId string
    folder string
    queryParams *url.Values
}

func NewDeleteFolderRecursivelySpectraS3Request(bucketId string, folder string) *DeleteFolderRecursivelySpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("bucket_id", bucketId)
    queryParams.Set("recursive", "")

    return &DeleteFolderRecursivelySpectraS3Request{
        folder: folder,
        bucketId: bucketId,
        queryParams: queryParams,
    }
}



func (deleteFolderRecursivelySpectraS3Request *DeleteFolderRecursivelySpectraS3Request) WithRollBack() *DeleteFolderRecursivelySpectraS3Request {
    deleteFolderRecursivelySpectraS3Request.queryParams.Set("roll_back", "")
    return deleteFolderRecursivelySpectraS3Request
}

func (DeleteFolderRecursivelySpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteFolderRecursivelySpectraS3Request *DeleteFolderRecursivelySpectraS3Request) Path() string {
    return "/_rest_/folder/" + deleteFolderRecursivelySpectraS3Request.folder
}

func (deleteFolderRecursivelySpectraS3Request *DeleteFolderRecursivelySpectraS3Request) QueryParams() *url.Values {
    return deleteFolderRecursivelySpectraS3Request.queryParams
}

func (DeleteFolderRecursivelySpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteFolderRecursivelySpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteFolderRecursivelySpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
