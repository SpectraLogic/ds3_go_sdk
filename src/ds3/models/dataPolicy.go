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

type DataPolicy struct {
    AlwaysForcePutJobCreation bool `xml:"AlwaysForcePutJobCreation"`
    AlwaysMinimizeSpanningAcrossMedia bool `xml:"AlwaysMinimizeSpanningAcrossMedia"`
    AlwaysReplicateDeletes bool `xml:"AlwaysReplicateDeletes"`
    BlobbingEnabled bool `xml:"BlobbingEnabled"`
    ChecksumType ChecksumType `xml:"ChecksumType"`
    CreationDate string `xml:"CreationDate"`
    DefaultBlobSize *int64 `xml:"DefaultBlobSize"`
    DefaultGetJobPriority Priority `xml:"DefaultGetJobPriority"`
    DefaultPutJobPriority Priority `xml:"DefaultPutJobPriority"`
    DefaultVerifyAfterWrite bool `xml:"DefaultVerifyAfterWrite"`
    DefaultVerifyJobPriority Priority `xml:"DefaultVerifyJobPriority"`
    EndToEndCrcRequired bool `xml:"EndToEndCrcRequired"`
    Id string `xml:"Id"`
    LtfsObjectNamingAllowed bool `xml:"LtfsObjectNamingAllowed"`
    Name *string `xml:"Name"`
    RebuildPriority Priority `xml:"RebuildPriority"`
    Versioning VersioningLevel `xml:"Versioning"`
}