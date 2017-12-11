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

func (dataPathBackend *DataPathBackend) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Activated":
            dataPathBackend.Activated = parseBool(child.Content, aggErr)
        case "AutoActivateTimeoutInMins":
            dataPathBackend.AutoActivateTimeoutInMins = parseNullableInt(child.Content, aggErr)
        case "AutoInspect":
            parseEnum(child.Content, &dataPathBackend.AutoInspect, aggErr)
        case "DefaultImportConflictResolutionMode":
            parseEnum(child.Content, &dataPathBackend.DefaultImportConflictResolutionMode, aggErr)
        case "DefaultVerifyDataAfterImport":
            parseNullableEnum(child.Content, dataPathBackend.DefaultVerifyDataAfterImport, aggErr)
        case "DefaultVerifyDataPriorToImport":
            dataPathBackend.DefaultVerifyDataPriorToImport = parseBool(child.Content, aggErr)
        case "Id":
            dataPathBackend.Id = parseString(child.Content)
        case "InstanceId":
            dataPathBackend.InstanceId = parseString(child.Content)
        case "LastHeartbeat":
            dataPathBackend.LastHeartbeat = parseString(child.Content)
        case "PartiallyVerifyLastPercentOfTapes":
            dataPathBackend.PartiallyVerifyLastPercentOfTapes = parseNullableInt(child.Content, aggErr)
        case "UnavailableMediaPolicy":
            parseEnum(child.Content, &dataPathBackend.UnavailableMediaPolicy, aggErr)
        case "UnavailablePoolMaxJobRetryInMins":
            dataPathBackend.UnavailablePoolMaxJobRetryInMins = parseInt(child.Content, aggErr)
        case "UnavailableTapePartitionMaxJobRetryInMins":
            dataPathBackend.UnavailableTapePartitionMaxJobRetryInMins = parseInt(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DataPathBackend.", child.XMLName.Local)
        }
    }
}
