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

func (bucketDetails *BucketDetails) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            bucketDetails.CreationDate = parseNullableString(child.Content)
        case "Name":
            bucketDetails.Name = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BucketDetails.", child.XMLName.Local)
        }
    }
}

func parseBucketDetailsSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []BucketDetails {
    var result []BucketDetails
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult BucketDetails
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing BucketDetails struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}
