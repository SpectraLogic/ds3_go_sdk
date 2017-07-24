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

type ActiveJob struct {
    Aggregating bool `xml:"Aggregating"`
    BucketId string `xml:"BucketId"`
    CachedSizeInBytes int64 `xml:"CachedSizeInBytes"`
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee `xml:"ChunkClientProcessingOrderGuarantee"`
    CompletedSizeInBytes int64 `xml:"CompletedSizeInBytes"`
    CreatedAt string `xml:"CreatedAt"`
    DeadJobCleanupAllowed bool `xml:"DeadJobCleanupAllowed"`
    ErrorMessage *string `xml:"ErrorMessage"`
    Id string `xml:"Id"`
    ImplicitJobIdResolution bool `xml:"ImplicitJobIdResolution"`
    MinimizeSpanningAcrossMedia bool `xml:"MinimizeSpanningAcrossMedia"`
    Naked bool `xml:"Naked"`
    Name *string `xml:"Name"`
    OriginalSizeInBytes int64 `xml:"OriginalSizeInBytes"`
    Priority Priority `xml:"Priority"`
    Rechunked *string `xml:"Rechunked"`
    Replicating bool `xml:"Replicating"`
    RequestType JobRequestType `xml:"RequestType"`
    Truncated bool `xml:"Truncated"`
    TruncatedDueToTimeout bool `xml:"TruncatedDueToTimeout"`
    UserId string `xml:"UserId"`
    VerifyAfterWrite bool `xml:"VerifyAfterWrite"`
}