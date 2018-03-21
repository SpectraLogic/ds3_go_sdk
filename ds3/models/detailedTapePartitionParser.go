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

func (detailedTapePartition *DetailedTapePartition) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AutoCompactionEnabled":
            detailedTapePartition.AutoCompactionEnabled = parseBool(child.Content, aggErr)
        case "DriveType":
            parseNullableEnum(child.Content, detailedTapePartition.DriveType, aggErr)
        case "DriveTypes":
            var model TapeDriveType
            parseEnum(child.Content, &model, aggErr)
            detailedTapePartition.DriveTypes = append(detailedTapePartition.DriveTypes, model)
        case "ErrorMessage":
            detailedTapePartition.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            detailedTapePartition.Id = parseString(child.Content)
        case "ImportExportConfiguration":
            parseEnum(child.Content, &detailedTapePartition.ImportExportConfiguration, aggErr)
        case "LibraryId":
            detailedTapePartition.LibraryId = parseString(child.Content)
        case "MinimumReadReservedDrives":
            detailedTapePartition.MinimumReadReservedDrives = parseInt(child.Content, aggErr)
        case "MinimumWriteReservedDrives":
            detailedTapePartition.MinimumWriteReservedDrives = parseInt(child.Content, aggErr)
        case "Name":
            detailedTapePartition.Name = parseNullableString(child.Content)
        case "Quiesced":
            parseEnum(child.Content, &detailedTapePartition.Quiesced, aggErr)
        case "SerialId":
            detailedTapePartition.SerialId = parseNullableString(child.Content)
        case "SerialNumber":
            detailedTapePartition.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &detailedTapePartition.State, aggErr)
        case "TapeTypes":
            var str = parseString(child.Content)
            detailedTapePartition.TapeTypes = append(detailedTapePartition.TapeTypes, str)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DetailedTapePartition.", child.XMLName.Local)
        }
    }
}
