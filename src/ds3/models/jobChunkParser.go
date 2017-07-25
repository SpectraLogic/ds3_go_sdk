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

func (jobChunk *JobChunk) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BlobStoreState":
            parseEnum(child.Content, &jobChunk.BlobStoreState, aggErr)
        case "ChunkNumber":
            jobChunk.ChunkNumber = parseInt(child.Content, aggErr)
        case "Id":
            jobChunk.Id = parseString(child.Content)
        case "JobCreationDate":
            jobChunk.JobCreationDate = parseString(child.Content)
        case "JobId":
            jobChunk.JobId = parseString(child.Content)
        case "NodeId":
            jobChunk.NodeId = parseNullableString(child.Content)
        case "PendingTargetCommit":
            jobChunk.PendingTargetCommit = parseBool(child.Content, aggErr)
        case "ReadFromAzureTargetId":
            jobChunk.ReadFromAzureTargetId = parseNullableString(child.Content)
        case "ReadFromDs3TargetId":
            jobChunk.ReadFromDs3TargetId = parseNullableString(child.Content)
        case "ReadFromPoolId":
            jobChunk.ReadFromPoolId = parseNullableString(child.Content)
        case "ReadFromS3TargetId":
            jobChunk.ReadFromS3TargetId = parseNullableString(child.Content)
        case "ReadFromTapeId":
            jobChunk.ReadFromTapeId = parseNullableString(child.Content)
        }
    }
}
