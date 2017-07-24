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

func (job *Job) parse(xmlNode *XmlNode, aggErr *AggregateError) {
    // Parse Attributes
    for _, attr := range xmlNode.Attrs {
        switch attr.Name.Local {
        case "Aggregating":
            job.Aggregating = parseBoolFromString(attr.Value, aggErr)
        case "BucketName":
            job.BucketName = parseNullableStringFromString(attr.Value)
        case "CachedSizeInBytes":
            job.CachedSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "ChunkClientProcessingOrderGuarantee":
            parseEnumFromString(attr.Value, &job.ChunkClientProcessingOrderGuarantee, aggErr)
        case "CompletedSizeInBytes":
            job.CompletedSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "EntirelyInCache":
            job.EntirelyInCache = parseBoolFromString(attr.Value, aggErr)
        case "JobId":
            job.JobId = attr.Value
        case "Naked":
            job.Naked = parseBoolFromString(attr.Value, aggErr)
        case "Name":
            job.Name = parseNullableStringFromString(attr.Value)
        case "OriginalSizeInBytes":
            job.OriginalSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "Priority":
            parseEnumFromString(attr.Value, &job.Priority, aggErr)
        case "RequestType":
            parseEnumFromString(attr.Value, &job.RequestType, aggErr)
        case "StartDate":
            job.StartDate = attr.Value
        case "Status":
            parseEnumFromString(attr.Value, &job.Status, aggErr)
        case "UserId":
            job.UserId = attr.Value
        case "UserName":
            job.UserName = parseNullableStringFromString(attr.Value)
        }
    }

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Nodes":
            job.Nodes = parseJobNodeSlice("Node", child.Children, aggErr)
        }
    }
}

func parseJobSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []Job {
    var result []Job
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult Job
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing Job struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}
