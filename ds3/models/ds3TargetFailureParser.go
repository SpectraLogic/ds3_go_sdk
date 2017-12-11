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

func (ds3TargetFailure *Ds3TargetFailure) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Date":
            ds3TargetFailure.Date = parseString(child.Content)
        case "ErrorMessage":
            ds3TargetFailure.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            ds3TargetFailure.Id = parseString(child.Content)
        case "TargetId":
            ds3TargetFailure.TargetId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &ds3TargetFailure.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3TargetFailure.", child.XMLName.Local)
        }
    }
}
