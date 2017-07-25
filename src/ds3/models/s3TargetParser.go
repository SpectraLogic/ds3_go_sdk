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

func (s3Target *S3Target) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AccessKey":
            s3Target.AccessKey = parseNullableString(child.Content)
        case "AutoVerifyFrequencyInDays":
            s3Target.AutoVerifyFrequencyInDays = parseNullableInt(child.Content, aggErr)
        case "CloudBucketPrefix":
            s3Target.CloudBucketPrefix = parseNullableString(child.Content)
        case "CloudBucketSuffix":
            s3Target.CloudBucketSuffix = parseNullableString(child.Content)
        case "DataPathEndPoint":
            s3Target.DataPathEndPoint = parseNullableString(child.Content)
        case "DefaultReadPreference":
            parseEnum(child.Content, &s3Target.DefaultReadPreference, aggErr)
        case "Https":
            s3Target.Https = parseBool(child.Content, aggErr)
        case "Id":
            s3Target.Id = parseString(child.Content)
        case "LastFullyVerified":
            s3Target.LastFullyVerified = parseNullableString(child.Content)
        case "Name":
            s3Target.Name = parseNullableString(child.Content)
        case "OfflineDataStagingWindowInTb":
            s3Target.OfflineDataStagingWindowInTb = parseInt(child.Content, aggErr)
        case "PermitGoingOutOfSync":
            s3Target.PermitGoingOutOfSync = parseBool(child.Content, aggErr)
        case "ProxyDomain":
            s3Target.ProxyDomain = parseNullableString(child.Content)
        case "ProxyHost":
            s3Target.ProxyHost = parseNullableString(child.Content)
        case "ProxyPassword":
            s3Target.ProxyPassword = parseNullableString(child.Content)
        case "ProxyPort":
            s3Target.ProxyPort = parseNullableInt(child.Content, aggErr)
        case "ProxyUsername":
            s3Target.ProxyUsername = parseNullableString(child.Content)
        case "Quiesced":
            parseEnum(child.Content, &s3Target.Quiesced, aggErr)
        case "Region":
            parseEnum(child.Content, &s3Target.Region, aggErr)
        case "SecretKey":
            s3Target.SecretKey = parseNullableString(child.Content)
        case "StagedDataExpirationInDays":
            s3Target.StagedDataExpirationInDays = parseInt(child.Content, aggErr)
        case "State":
            parseEnum(child.Content, &s3Target.State, aggErr)
        }
    }
}

func parseS3TargetSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []S3Target {
    var result []S3Target
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult S3Target
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing S3Target struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}
