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

func (listBucketResult *ListBucketResult) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CommonPrefixes":
            var prefixes []string
            prefixes = parseStringSlice("Prefix", child.Children, aggErr)
            listBucketResult.CommonPrefixes = append(listBucketResult.CommonPrefixes, prefixes...)
        case "CreationDate":
            listBucketResult.CreationDate = parseNullableString(child.Content)
        case "Delimiter":
            listBucketResult.Delimiter = parseNullableString(child.Content)
        case "Marker":
            listBucketResult.Marker = parseNullableString(child.Content)
        case "MaxKeys":
            listBucketResult.MaxKeys = parseInt(child.Content, aggErr)
        case "Name":
            listBucketResult.Name = parseNullableString(child.Content)
        case "NextMarker":
            listBucketResult.NextMarker = parseNullableString(child.Content)
        case "Contents":
            var model Contents
            model.parse(&child, aggErr)
            listBucketResult.Objects = append(listBucketResult.Objects, model)
        case "Prefix":
            listBucketResult.Prefix = parseNullableString(child.Content)
        case "IsTruncated":
            listBucketResult.Truncated = parseBool(child.Content, aggErr)
        case "Version":
            var model Contents
            model.parse(&child, aggErr)
            listBucketResult.VersionedObjects = append(listBucketResult.VersionedObjects, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing ListBucketResult.", child.XMLName.Local)
        }
    }
}
