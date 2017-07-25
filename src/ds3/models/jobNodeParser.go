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

func (jobNode *JobNode) parse(xmlNode *XmlNode, aggErr *AggregateError) {
    // Parse Attributes
    for _, attr := range xmlNode.Attrs {
        switch attr.Name.Local {
        case "EndPoint":
            jobNode.EndPoint = parseNullableStringFromString(attr.Value)
        case "HttpPort":
            jobNode.HttpPort = parseNullableIntFromString(attr.Value, aggErr)
        case "HttpsPort":
            jobNode.HttpsPort = parseNullableIntFromString(attr.Value, aggErr)
        case "Id":
            jobNode.Id = attr.Value
        }
    }

}

func parseJobNodeSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []JobNode {
    var result []JobNode
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult JobNode
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing JobNode struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}
