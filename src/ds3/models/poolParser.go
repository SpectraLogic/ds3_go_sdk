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

func (pool *Pool) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AssignedToStorageDomain":
            pool.AssignedToStorageDomain = parseBool(child.Content, aggErr)
        case "AvailableCapacity":
            pool.AvailableCapacity = parseInt64(child.Content, aggErr)
        case "BucketId":
            pool.BucketId = parseNullableString(child.Content)
        case "Guid":
            pool.Guid = parseNullableString(child.Content)
        case "Health":
            parseEnum(child.Content, &pool.Health, aggErr)
        case "Id":
            pool.Id = parseString(child.Content)
        case "LastAccessed":
            pool.LastAccessed = parseNullableString(child.Content)
        case "LastModified":
            pool.LastModified = parseNullableString(child.Content)
        case "LastVerified":
            pool.LastVerified = parseNullableString(child.Content)
        case "Mountpoint":
            pool.Mountpoint = parseNullableString(child.Content)
        case "Name":
            pool.Name = parseNullableString(child.Content)
        case "PartitionId":
            pool.PartitionId = parseNullableString(child.Content)
        case "PoweredOn":
            pool.PoweredOn = parseBool(child.Content, aggErr)
        case "Quiesced":
            parseEnum(child.Content, &pool.Quiesced, aggErr)
        case "ReservedCapacity":
            pool.ReservedCapacity = parseInt64(child.Content, aggErr)
        case "State":
            parseEnum(child.Content, &pool.State, aggErr)
        case "StorageDomainId":
            pool.StorageDomainId = parseNullableString(child.Content)
        case "TotalCapacity":
            pool.TotalCapacity = parseInt64(child.Content, aggErr)
        case "Type":
            parseEnum(child.Content, &pool.Type, aggErr)
        case "UsedCapacity":
            pool.UsedCapacity = parseInt64(child.Content, aggErr)
        }
    }
}

func parsePoolSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []Pool {
    var result []Pool
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult Pool
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing Pool struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}
