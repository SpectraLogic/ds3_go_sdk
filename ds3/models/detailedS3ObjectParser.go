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

func (detailedS3Object *DetailedS3Object) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Blobs":
            var model BulkObjectList
            model.parse(&child, aggErr)
            detailedS3Object.Blobs = &model
        case "BlobsBeingPersisted":
            detailedS3Object.BlobsBeingPersisted = parseNullableInt(child.Content, aggErr)
        case "BlobsDegraded":
            detailedS3Object.BlobsDegraded = parseNullableInt(child.Content, aggErr)
        case "BlobsInCache":
            detailedS3Object.BlobsInCache = parseNullableInt(child.Content, aggErr)
        case "BlobsTotal":
            detailedS3Object.BlobsTotal = parseNullableInt(child.Content, aggErr)
        case "BucketId":
            detailedS3Object.BucketId = parseString(child.Content)
        case "CreationDate":
            detailedS3Object.CreationDate = parseNullableString(child.Content)
        case "ETag":
            detailedS3Object.ETag = parseNullableString(child.Content)
        case "Id":
            detailedS3Object.Id = parseString(child.Content)
        case "Latest":
            detailedS3Object.Latest = parseBool(child.Content, aggErr)
        case "Name":
            detailedS3Object.Name = parseNullableString(child.Content)
        case "Owner":
            detailedS3Object.Owner = parseNullableString(child.Content)
        case "Size":
            detailedS3Object.Size = parseInt64(child.Content, aggErr)
        case "Type":
            parseEnum(child.Content, &detailedS3Object.Type, aggErr)
        case "Version":
            detailedS3Object.Version = parseInt64(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DetailedS3Object.", child.XMLName.Local)
        }
    }
}
