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

func (ds3Target *Ds3Target) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AccessControlReplication":
            parseEnum(child.Content, &ds3Target.AccessControlReplication, aggErr)
        case "AdminAuthId":
            ds3Target.AdminAuthId = parseNullableString(child.Content)
        case "AdminSecretKey":
            ds3Target.AdminSecretKey = parseNullableString(child.Content)
        case "DataPathEndPoint":
            ds3Target.DataPathEndPoint = parseNullableString(child.Content)
        case "DataPathHttps":
            ds3Target.DataPathHttps = parseBool(child.Content, aggErr)
        case "DataPathPort":
            ds3Target.DataPathPort = parseNullableInt(child.Content, aggErr)
        case "DataPathProxy":
            ds3Target.DataPathProxy = parseNullableString(child.Content)
        case "DataPathVerifyCertificate":
            ds3Target.DataPathVerifyCertificate = parseBool(child.Content, aggErr)
        case "DefaultReadPreference":
            parseEnum(child.Content, &ds3Target.DefaultReadPreference, aggErr)
        case "Id":
            ds3Target.Id = parseString(child.Content)
        case "Name":
            ds3Target.Name = parseNullableString(child.Content)
        case "PermitGoingOutOfSync":
            ds3Target.PermitGoingOutOfSync = parseBool(child.Content, aggErr)
        case "Quiesced":
            parseEnum(child.Content, &ds3Target.Quiesced, aggErr)
        case "ReplicatedUserDefaultDataPolicy":
            ds3Target.ReplicatedUserDefaultDataPolicy = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &ds3Target.State, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3Target.", child.XMLName.Local)
        }
    }
}

func parseDs3TargetSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []Ds3Target {
    var result []Ds3Target
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult Ds3Target
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing Ds3Target struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}
