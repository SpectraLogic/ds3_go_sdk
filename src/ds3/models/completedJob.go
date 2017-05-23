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

type CompletedJob struct {
    BucketId string `xml:"BucketId"`
    CachedSizeInBytes int64 `xml:"CachedSizeInBytes"`
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee `xml:"ChunkClientProcessingOrderGuarantee"`
    CompletedSizeInBytes int64 `xml:"CompletedSizeInBytes"`
    CreatedAt string `xml:"CreatedAt"`
    DateCompleted string `xml:"DateCompleted"`
    ErrorMessage *string `xml:"ErrorMessage"`
    Id string `xml:"Id"`
    Naked bool `xml:"Naked"`
    Name *string `xml:"Name"`
    OriginalSizeInBytes int64 `xml:"OriginalSizeInBytes"`
    Priority Priority `xml:"Priority"`
    Rechunked *string `xml:"Rechunked"`
    RequestType JobRequestType `xml:"RequestType"`
    Truncated bool `xml:"Truncated"`
    UserId string `xml:"UserId"`
}