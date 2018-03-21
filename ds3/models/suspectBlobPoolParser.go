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

import "log"

func (suspectBlobPool *SuspectBlobPool) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BlobId":
            suspectBlobPool.BlobId = parseString(child.Content)
        case "BucketId":
            suspectBlobPool.BucketId = parseString(child.Content)
        case "DateWritten":
            suspectBlobPool.DateWritten = parseString(child.Content)
        case "Id":
            suspectBlobPool.Id = parseString(child.Content)
        case "LastAccessed":
            suspectBlobPool.LastAccessed = parseString(child.Content)
        case "ObsoletionId":
            suspectBlobPool.ObsoletionId = parseNullableString(child.Content)
        case "PoolId":
            suspectBlobPool.PoolId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobPool.", child.XMLName.Local)
        }
    }
}
