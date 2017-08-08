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

func (jobCreationFailedNotificationRegistration *JobCreationFailedNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            jobCreationFailedNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &jobCreationFailedNotificationRegistration.Format, aggErr)
        case "Id":
            jobCreationFailedNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            jobCreationFailedNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            jobCreationFailedNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            jobCreationFailedNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &jobCreationFailedNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            jobCreationFailedNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &jobCreationFailedNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            jobCreationFailedNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            jobCreationFailedNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing JobCreationFailedNotificationRegistration.", child.XMLName.Local)
        }
    }
}
