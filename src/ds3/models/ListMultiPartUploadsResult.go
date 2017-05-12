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

type ListMultiPartUploadsResult struct {
    Bucket *string `xml:"Bucket"`
    CommonPrefixes []string `xml:"CommonPrefixes>Prefix"`
    Delimiter *string `xml:"Delimiter"`
    KeyMarker *string `xml:"KeyMarker"`
    MaxUploads int `xml:"MaxUploads"`
    NextKeyMarker *string `xml:"NextKeyMarker"`
    NextUploadIdMarker *string `xml:"NextUploadIdMarker"`
    Prefix *string `xml:"Prefix"`
    Truncated bool `xml:"IsTruncated"`
    UploadIdMarker *string `xml:"UploadIdMarker"`
    Uploads []MultiPartUpload `xml:"Upload"`
}