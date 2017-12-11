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

func (degradedBlob *DegradedBlob) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AzureReplicationRuleId":
            degradedBlob.AzureReplicationRuleId = parseNullableString(child.Content)
        case "BlobId":
            degradedBlob.BlobId = parseString(child.Content)
        case "BucketId":
            degradedBlob.BucketId = parseString(child.Content)
        case "Ds3ReplicationRuleId":
            degradedBlob.Ds3ReplicationRuleId = parseNullableString(child.Content)
        case "Id":
            degradedBlob.Id = parseString(child.Content)
        case "PersistenceRuleId":
            degradedBlob.PersistenceRuleId = parseNullableString(child.Content)
        case "S3ReplicationRuleId":
            degradedBlob.S3ReplicationRuleId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DegradedBlob.", child.XMLName.Local)
        }
    }
}
