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

func (listPartsResult *ListPartsResult) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Bucket":
            listPartsResult.Bucket = parseNullableString(child.Content)
        case "Key":
            listPartsResult.Key = parseNullableString(child.Content)
        case "MaxParts":
            listPartsResult.MaxParts = parseInt(child.Content, aggErr)
        case "NextPartNumberMarker":
            listPartsResult.NextPartNumberMarker = parseInt(child.Content, aggErr)
        case "Owner":
            listPartsResult.Owner.parse(&child, aggErr)
        case "PartNumberMarker":
            listPartsResult.PartNumberMarker = parseNullableInt(child.Content, aggErr)
        case "Part":
            var model MultiPartUploadPart
            model.parse(&child, aggErr)
            listPartsResult.Parts = append(listPartsResult.Parts, model)
        case "IsTruncated":
            listPartsResult.Truncated = parseBool(child.Content, aggErr)
        case "UploadId":
            listPartsResult.UploadId = parseString(child.Content)
        }
    }
}
