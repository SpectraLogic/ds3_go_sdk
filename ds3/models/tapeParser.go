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

func (tape *Tape) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AssignedToStorageDomain":
            tape.AssignedToStorageDomain = parseBool(child.Content, aggErr)
        case "AvailableRawCapacity":
            tape.AvailableRawCapacity = parseNullableInt64(child.Content, aggErr)
        case "BarCode":
            tape.BarCode = parseNullableString(child.Content)
        case "BucketId":
            tape.BucketId = parseNullableString(child.Content)
        case "DescriptionForIdentification":
            tape.DescriptionForIdentification = parseNullableString(child.Content)
        case "EjectDate":
            tape.EjectDate = parseNullableString(child.Content)
        case "EjectLabel":
            tape.EjectLabel = parseNullableString(child.Content)
        case "EjectLocation":
            tape.EjectLocation = parseNullableString(child.Content)
        case "EjectPending":
            tape.EjectPending = parseNullableString(child.Content)
        case "FullOfData":
            tape.FullOfData = parseBool(child.Content, aggErr)
        case "Id":
            tape.Id = parseString(child.Content)
        case "LastAccessed":
            tape.LastAccessed = parseNullableString(child.Content)
        case "LastCheckpoint":
            tape.LastCheckpoint = parseNullableString(child.Content)
        case "LastModified":
            tape.LastModified = parseNullableString(child.Content)
        case "LastVerified":
            tape.LastVerified = parseNullableString(child.Content)
        case "PartiallyVerifiedEndOfTape":
            tape.PartiallyVerifiedEndOfTape = parseNullableString(child.Content)
        case "PartitionId":
            tape.PartitionId = parseNullableString(child.Content)
        case "PreviousState":
            parseNullableEnum(child.Content, tape.PreviousState, aggErr)
        case "SerialNumber":
            tape.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &tape.State, aggErr)
        case "StorageDomainMemberId":
            tape.StorageDomainMemberId = parseNullableString(child.Content)
        case "TakeOwnershipPending":
            tape.TakeOwnershipPending = parseBool(child.Content, aggErr)
        case "TotalRawCapacity":
            tape.TotalRawCapacity = parseNullableInt64(child.Content, aggErr)
        case "Type":
            tape.Type = parseString(child.Content)
        case "VerifyPending":
            parseNullableEnum(child.Content, tape.VerifyPending, aggErr)
        case "WriteProtected":
            tape.WriteProtected = parseBool(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Tape.", child.XMLName.Local)
        }
    }
}

func parseTapeSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []Tape {
    var result []Tape
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult Tape
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing Tape struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}
