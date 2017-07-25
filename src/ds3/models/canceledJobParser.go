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

func (canceledJob *CanceledJob) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            canceledJob.BucketId = parseString(child.Content)
        case "CachedSizeInBytes":
            canceledJob.CachedSizeInBytes = parseInt64(child.Content, aggErr)
        case "CanceledDueToTimeout":
            canceledJob.CanceledDueToTimeout = parseBool(child.Content, aggErr)
        case "ChunkClientProcessingOrderGuarantee":
            parseEnum(child.Content, &canceledJob.ChunkClientProcessingOrderGuarantee, aggErr)
        case "CompletedSizeInBytes":
            canceledJob.CompletedSizeInBytes = parseInt64(child.Content, aggErr)
        case "CreatedAt":
            canceledJob.CreatedAt = parseString(child.Content)
        case "DateCanceled":
            canceledJob.DateCanceled = parseString(child.Content)
        case "ErrorMessage":
            canceledJob.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            canceledJob.Id = parseString(child.Content)
        case "Naked":
            canceledJob.Naked = parseBool(child.Content, aggErr)
        case "Name":
            canceledJob.Name = parseNullableString(child.Content)
        case "OriginalSizeInBytes":
            canceledJob.OriginalSizeInBytes = parseInt64(child.Content, aggErr)
        case "Priority":
            parseEnum(child.Content, &canceledJob.Priority, aggErr)
        case "Rechunked":
            canceledJob.Rechunked = parseNullableString(child.Content)
        case "RequestType":
            parseEnum(child.Content, &canceledJob.RequestType, aggErr)
        case "Truncated":
            canceledJob.Truncated = parseBool(child.Content, aggErr)
        case "UserId":
            canceledJob.UserId = parseString(child.Content)
        }
    }
}
