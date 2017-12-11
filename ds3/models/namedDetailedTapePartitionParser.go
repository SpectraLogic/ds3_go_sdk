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

func (namedDetailedTapePartition *NamedDetailedTapePartition) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DriveType":
            parseNullableEnum(child.Content, namedDetailedTapePartition.DriveType, aggErr)
        case "DriveTypes":
            var model TapeDriveType
            parseEnum(child.Content, &model, aggErr)
            namedDetailedTapePartition.DriveTypes = append(namedDetailedTapePartition.DriveTypes, model)
        case "ErrorMessage":
            namedDetailedTapePartition.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            namedDetailedTapePartition.Id = parseString(child.Content)
        case "ImportExportConfiguration":
            parseEnum(child.Content, &namedDetailedTapePartition.ImportExportConfiguration, aggErr)
        case "LibraryId":
            namedDetailedTapePartition.LibraryId = parseString(child.Content)
        case "Name":
            namedDetailedTapePartition.Name = parseNullableString(child.Content)
        case "Quiesced":
            parseEnum(child.Content, &namedDetailedTapePartition.Quiesced, aggErr)
        case "SerialId":
            namedDetailedTapePartition.SerialId = parseNullableString(child.Content)
        case "SerialNumber":
            namedDetailedTapePartition.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &namedDetailedTapePartition.State, aggErr)
        case "TapeTypes":
            var model TapeType
            parseEnum(child.Content, &model, aggErr)
            namedDetailedTapePartition.TapeTypes = append(namedDetailedTapePartition.TapeTypes, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing NamedDetailedTapePartition.", child.XMLName.Local)
        }
    }
}
