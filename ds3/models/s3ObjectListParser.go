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

func (s3ObjectList *S3ObjectList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "S3Object":
            var model S3Object
            model.parse(&child, aggErr)
            s3ObjectList.S3Objects = append(s3ObjectList.S3Objects, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3ObjectList.", child.XMLName.Local)
        }
    }
}
