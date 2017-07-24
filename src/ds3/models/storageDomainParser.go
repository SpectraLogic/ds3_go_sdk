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

func (storageDomain *StorageDomain) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AutoEjectMediaFullThreshold":
            storageDomain.AutoEjectMediaFullThreshold = parseNullableInt64(child.Content, aggErr)
        case "AutoEjectUponCron":
            storageDomain.AutoEjectUponCron = parseNullableString(child.Content)
        case "AutoEjectUponJobCancellation":
            storageDomain.AutoEjectUponJobCancellation = parseBool(child.Content, aggErr)
        case "AutoEjectUponJobCompletion":
            storageDomain.AutoEjectUponJobCompletion = parseBool(child.Content, aggErr)
        case "AutoEjectUponMediaFull":
            storageDomain.AutoEjectUponMediaFull = parseBool(child.Content, aggErr)
        case "Id":
            storageDomain.Id = parseString(child.Content)
        case "LtfsFileNaming":
            parseEnum(child.Content, &storageDomain.LtfsFileNaming, aggErr)
        case "MaxTapeFragmentationPercent":
            storageDomain.MaxTapeFragmentationPercent = parseInt(child.Content, aggErr)
        case "MaximumAutoVerificationFrequencyInDays":
            storageDomain.MaximumAutoVerificationFrequencyInDays = parseNullableInt(child.Content, aggErr)
        case "MediaEjectionAllowed":
            storageDomain.MediaEjectionAllowed = parseBool(child.Content, aggErr)
        case "Name":
            storageDomain.Name = parseNullableString(child.Content)
        case "SecureMediaAllocation":
            storageDomain.SecureMediaAllocation = parseBool(child.Content, aggErr)
        case "VerifyPriorToAutoEject":
            parseNullableEnum(child.Content, storageDomain.VerifyPriorToAutoEject, aggErr)
        case "WriteOptimization":
            parseEnum(child.Content, &storageDomain.WriteOptimization, aggErr)
        }
    }
}
