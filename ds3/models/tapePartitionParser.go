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

func (tapePartition *TapePartition) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AutoCompactionEnabled":
            tapePartition.AutoCompactionEnabled = parseBool(child.Content, aggErr)
        case "DriveType":
            parseNullableEnum(child.Content, tapePartition.DriveType, aggErr)
        case "ErrorMessage":
            tapePartition.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            tapePartition.Id = parseString(child.Content)
        case "ImportExportConfiguration":
            parseEnum(child.Content, &tapePartition.ImportExportConfiguration, aggErr)
        case "LibraryId":
            tapePartition.LibraryId = parseString(child.Content)
        case "MinimumReadReservedDrives":
            tapePartition.MinimumReadReservedDrives = parseInt(child.Content, aggErr)
        case "MinimumWriteReservedDrives":
            tapePartition.MinimumWriteReservedDrives = parseInt(child.Content, aggErr)
        case "Name":
            tapePartition.Name = parseNullableString(child.Content)
        case "Quiesced":
            parseEnum(child.Content, &tapePartition.Quiesced, aggErr)
        case "SerialNumber":
            tapePartition.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &tapePartition.State, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapePartition.", child.XMLName.Local)
        }
    }
}
