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

func (blobStoreTaskInformation *BlobStoreTaskInformation) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DateScheduled":
            blobStoreTaskInformation.DateScheduled = parseString(child.Content)
        case "DateStarted":
            blobStoreTaskInformation.DateStarted = parseNullableString(child.Content)
        case "Description":
            blobStoreTaskInformation.Description = parseNullableString(child.Content)
        case "DriveId":
            blobStoreTaskInformation.DriveId = parseNullableString(child.Content)
        case "Id":
            blobStoreTaskInformation.Id = parseInt64(child.Content, aggErr)
        case "Name":
            blobStoreTaskInformation.Name = parseNullableString(child.Content)
        case "PoolId":
            blobStoreTaskInformation.PoolId = parseNullableString(child.Content)
        case "Priority":
            parseEnum(child.Content, &blobStoreTaskInformation.Priority, aggErr)
        case "State":
            parseEnum(child.Content, &blobStoreTaskInformation.State, aggErr)
        case "TapeId":
            blobStoreTaskInformation.TapeId = parseNullableString(child.Content)
        case "TargetId":
            blobStoreTaskInformation.TargetId = parseNullableString(child.Content)
        case "TargetType":
            blobStoreTaskInformation.TargetType = parseNullableString(child.Content)
        }
    }
}
