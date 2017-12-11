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

func (bulkObject *BulkObject) parse(xmlNode *XmlNode, aggErr *AggregateError) {
    // Parse Attributes
    for _, attr := range xmlNode.Attrs {
        switch attr.Name.Local {
        case "Bucket":
            bulkObject.Bucket = parseNullableStringFromString(attr.Value)
        case "Id":
            bulkObject.Id = parseNullableStringFromString(attr.Value)
        case "InCache":
            bulkObject.InCache = parseNullableBoolFromString(attr.Value, aggErr)
        case "Latest":
            bulkObject.Latest = parseBoolFromString(attr.Value, aggErr)
        case "Length":
            bulkObject.Length = parseInt64FromString(attr.Value, aggErr)
        case "Name":
            bulkObject.Name = parseNullableStringFromString(attr.Value)
        case "Offset":
            bulkObject.Offset = parseInt64FromString(attr.Value, aggErr)
        case "Version":
            bulkObject.Version = parseInt64FromString(attr.Value, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown attribute '%s' while parsing BulkObject.", attr.Name.Local)
        }
    }

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "PhysicalPlacement":
            var model PhysicalPlacement
            model.parse(&child, aggErr)
            bulkObject.PhysicalPlacement = &model
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BulkObject.", child.XMLName.Local)
        }
    }
}
