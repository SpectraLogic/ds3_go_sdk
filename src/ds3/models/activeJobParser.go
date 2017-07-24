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

func (activeJob *ActiveJob) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Aggregating":
            activeJob.Aggregating = parseBool(child.Content, aggErr)
        case "BucketId":
            activeJob.BucketId = parseString(child.Content)
        case "CachedSizeInBytes":
            activeJob.CachedSizeInBytes = parseInt64(child.Content, aggErr)
        case "ChunkClientProcessingOrderGuarantee":
            parseEnum(child.Content, &activeJob.ChunkClientProcessingOrderGuarantee, aggErr)
        case "CompletedSizeInBytes":
            activeJob.CompletedSizeInBytes = parseInt64(child.Content, aggErr)
        case "CreatedAt":
            activeJob.CreatedAt = parseString(child.Content)
        case "ErrorMessage":
            activeJob.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            activeJob.Id = parseString(child.Content)
        case "ImplicitJobIdResolution":
            activeJob.ImplicitJobIdResolution = parseBool(child.Content, aggErr)
        case "MinimizeSpanningAcrossMedia":
            activeJob.MinimizeSpanningAcrossMedia = parseBool(child.Content, aggErr)
        case "Naked":
            activeJob.Naked = parseBool(child.Content, aggErr)
        case "Name":
            activeJob.Name = parseNullableString(child.Content)
        case "OriginalSizeInBytes":
            activeJob.OriginalSizeInBytes = parseInt64(child.Content, aggErr)
        case "Priority":
            parseEnum(child.Content, &activeJob.Priority, aggErr)
        case "Rechunked":
            activeJob.Rechunked = parseNullableString(child.Content)
        case "Replicating":
            activeJob.Replicating = parseBool(child.Content, aggErr)
        case "RequestType":
            parseEnum(child.Content, &activeJob.RequestType, aggErr)
        case "Truncated":
            activeJob.Truncated = parseBool(child.Content, aggErr)
        case "TruncatedDueToTimeout":
            activeJob.TruncatedDueToTimeout = parseBool(child.Content, aggErr)
        case "UserId":
            activeJob.UserId = parseString(child.Content)
        case "VerifyAfterWrite":
            activeJob.VerifyAfterWrite = parseBool(child.Content, aggErr)
        }
    }
}
