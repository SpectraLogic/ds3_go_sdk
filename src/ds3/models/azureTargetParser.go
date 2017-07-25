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

func (azureTarget *AzureTarget) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AccountKey":
            azureTarget.AccountKey = parseNullableString(child.Content)
        case "AccountName":
            azureTarget.AccountName = parseNullableString(child.Content)
        case "AutoVerifyFrequencyInDays":
            azureTarget.AutoVerifyFrequencyInDays = parseNullableInt(child.Content, aggErr)
        case "CloudBucketPrefix":
            azureTarget.CloudBucketPrefix = parseNullableString(child.Content)
        case "CloudBucketSuffix":
            azureTarget.CloudBucketSuffix = parseNullableString(child.Content)
        case "DefaultReadPreference":
            parseEnum(child.Content, &azureTarget.DefaultReadPreference, aggErr)
        case "Https":
            azureTarget.Https = parseBool(child.Content, aggErr)
        case "Id":
            azureTarget.Id = parseString(child.Content)
        case "LastFullyVerified":
            azureTarget.LastFullyVerified = parseNullableString(child.Content)
        case "Name":
            azureTarget.Name = parseNullableString(child.Content)
        case "PermitGoingOutOfSync":
            azureTarget.PermitGoingOutOfSync = parseBool(child.Content, aggErr)
        case "Quiesced":
            parseEnum(child.Content, &azureTarget.Quiesced, aggErr)
        case "State":
            parseEnum(child.Content, &azureTarget.State, aggErr)
        }
    }
}

func parseAzureTargetSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []AzureTarget {
    var result []AzureTarget
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult AzureTarget
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing AzureTarget struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}
