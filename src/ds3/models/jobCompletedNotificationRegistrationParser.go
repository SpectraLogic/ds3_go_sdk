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

func (jobCompletedNotificationRegistration *JobCompletedNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            jobCompletedNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &jobCompletedNotificationRegistration.Format, aggErr)
        case "Id":
            jobCompletedNotificationRegistration.Id = parseString(child.Content)
        case "JobId":
            jobCompletedNotificationRegistration.JobId = parseNullableString(child.Content)
        case "LastFailure":
            jobCompletedNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            jobCompletedNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            jobCompletedNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &jobCompletedNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            jobCompletedNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &jobCompletedNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            jobCompletedNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            jobCompletedNotificationRegistration.UserId = parseNullableString(child.Content)
        }
    }
}
