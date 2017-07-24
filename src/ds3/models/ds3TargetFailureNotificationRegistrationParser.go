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

func (ds3TargetFailureNotificationRegistration *Ds3TargetFailureNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            ds3TargetFailureNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &ds3TargetFailureNotificationRegistration.Format, aggErr)
        case "Id":
            ds3TargetFailureNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            ds3TargetFailureNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            ds3TargetFailureNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            ds3TargetFailureNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &ds3TargetFailureNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            ds3TargetFailureNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &ds3TargetFailureNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            ds3TargetFailureNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            ds3TargetFailureNotificationRegistration.UserId = parseNullableString(child.Content)
        }
    }
}
