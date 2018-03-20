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

func (dataPolicy *DataPolicy) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AlwaysForcePutJobCreation":
            dataPolicy.AlwaysForcePutJobCreation = parseBool(child.Content, aggErr)
        case "AlwaysMinimizeSpanningAcrossMedia":
            dataPolicy.AlwaysMinimizeSpanningAcrossMedia = parseBool(child.Content, aggErr)
        case "BlobbingEnabled":
            dataPolicy.BlobbingEnabled = parseBool(child.Content, aggErr)
        case "ChecksumType":
            parseEnum(child.Content, &dataPolicy.ChecksumType, aggErr)
        case "CreationDate":
            dataPolicy.CreationDate = parseString(child.Content)
        case "DefaultBlobSize":
            dataPolicy.DefaultBlobSize = parseNullableInt64(child.Content, aggErr)
        case "DefaultGetJobPriority":
            parseEnum(child.Content, &dataPolicy.DefaultGetJobPriority, aggErr)
        case "DefaultPutJobPriority":
            parseEnum(child.Content, &dataPolicy.DefaultPutJobPriority, aggErr)
        case "DefaultVerifyAfterWrite":
            dataPolicy.DefaultVerifyAfterWrite = parseBool(child.Content, aggErr)
        case "DefaultVerifyJobPriority":
            parseEnum(child.Content, &dataPolicy.DefaultVerifyJobPriority, aggErr)
        case "EndToEndCrcRequired":
            dataPolicy.EndToEndCrcRequired = parseBool(child.Content, aggErr)
        case "Id":
            dataPolicy.Id = parseString(child.Content)
        case "MaxVersionsToKeep":
            dataPolicy.MaxVersionsToKeep = parseInt(child.Content, aggErr)
        case "Name":
            dataPolicy.Name = parseNullableString(child.Content)
        case "RebuildPriority":
            parseEnum(child.Content, &dataPolicy.RebuildPriority, aggErr)
        case "Versioning":
            parseEnum(child.Content, &dataPolicy.Versioning, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DataPolicy.", child.XMLName.Local)
        }
    }
}
