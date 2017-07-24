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

type Job struct {
    Aggregating bool `xml:"Aggregating,attr"`
    BucketName *string `xml:"BucketName,attr"`
    CachedSizeInBytes int64 `xml:"CachedSizeInBytes,attr"`
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee `xml:"ChunkClientProcessingOrderGuarantee,attr"`
    CompletedSizeInBytes int64 `xml:"CompletedSizeInBytes,attr"`
    EntirelyInCache *bool `xml:"EntirelyInCache,attr"`
    JobId string `xml:"JobId,attr"`
    Naked bool `xml:"Naked,attr"`
    Name *string `xml:"Name,attr"`
    Nodes []JobNode `xml:"Nodes>Node"`
    OriginalSizeInBytes int64 `xml:"OriginalSizeInBytes,attr"`
    Priority Priority `xml:"Priority,attr"`
    RequestType JobRequestType `xml:"RequestType,attr"`
    StartDate string `xml:"StartDate,attr"`
    Status JobStatus `xml:"Status,attr"`
    UserId string `xml:"UserId,attr"`
    UserName *string `xml:"UserName,attr"`
}