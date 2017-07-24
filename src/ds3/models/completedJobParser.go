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

func (completedJob *CompletedJob) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            completedJob.BucketId = parseString(child.Content)
        case "CachedSizeInBytes":
            completedJob.CachedSizeInBytes = parseInt64(child.Content, aggErr)
        case "ChunkClientProcessingOrderGuarantee":
            parseEnum(child.Content, &completedJob.ChunkClientProcessingOrderGuarantee, aggErr)
        case "CompletedSizeInBytes":
            completedJob.CompletedSizeInBytes = parseInt64(child.Content, aggErr)
        case "CreatedAt":
            completedJob.CreatedAt = parseString(child.Content)
        case "DateCompleted":
            completedJob.DateCompleted = parseString(child.Content)
        case "ErrorMessage":
            completedJob.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            completedJob.Id = parseString(child.Content)
        case "Naked":
            completedJob.Naked = parseBool(child.Content, aggErr)
        case "Name":
            completedJob.Name = parseNullableString(child.Content)
        case "OriginalSizeInBytes":
            completedJob.OriginalSizeInBytes = parseInt64(child.Content, aggErr)
        case "Priority":
            parseEnum(child.Content, &completedJob.Priority, aggErr)
        case "Rechunked":
            completedJob.Rechunked = parseNullableString(child.Content)
        case "RequestType":
            parseEnum(child.Content, &completedJob.RequestType, aggErr)
        case "Truncated":
            completedJob.Truncated = parseBool(child.Content, aggErr)
        case "UserId":
            completedJob.UserId = parseString(child.Content)
        }
    }
}
