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

func (masterObjectList *MasterObjectList) parse(xmlNode *XmlNode, aggErr *AggregateError) {
    // Parse Attributes
    for _, attr := range xmlNode.Attrs {
        switch attr.Name.Local {
        case "Aggregating":
            masterObjectList.Aggregating = parseBoolFromString(attr.Value, aggErr)
        case "BucketName":
            masterObjectList.BucketName = parseNullableStringFromString(attr.Value)
        case "CachedSizeInBytes":
            masterObjectList.CachedSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "ChunkClientProcessingOrderGuarantee":
            parseEnumFromString(attr.Value, &masterObjectList.ChunkClientProcessingOrderGuarantee, aggErr)
        case "CompletedSizeInBytes":
            masterObjectList.CompletedSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "EntirelyInCache":
            masterObjectList.EntirelyInCache = parseNullableBoolFromString(attr.Value, aggErr)
        case "JobId":
            masterObjectList.JobId = attr.Value
        case "Naked":
            masterObjectList.Naked = parseBoolFromString(attr.Value, aggErr)
        case "Name":
            masterObjectList.Name = parseNullableStringFromString(attr.Value)
        case "OriginalSizeInBytes":
            masterObjectList.OriginalSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "Priority":
            parseEnumFromString(attr.Value, &masterObjectList.Priority, aggErr)
        case "RequestType":
            parseEnumFromString(attr.Value, &masterObjectList.RequestType, aggErr)
        case "StartDate":
            masterObjectList.StartDate = attr.Value
        case "Status":
            parseEnumFromString(attr.Value, &masterObjectList.Status, aggErr)
        case "UserId":
            masterObjectList.UserId = attr.Value
        case "UserName":
            masterObjectList.UserName = parseNullableStringFromString(attr.Value)
        default:
            log.Printf("WARNING: unable to parse unknown attribute '%s' while parsing MasterObjectList.", attr.Name.Local)
        }
    }

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Nodes":
            masterObjectList.Nodes = parseJobNodeSlice("Node", child.Children, aggErr)
        case "Objects":
            var model Objects
            model.parse(&child, aggErr)
            masterObjectList.Objects = append(masterObjectList.Objects, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing MasterObjectList.", child.XMLName.Local)
        }
    }
}
