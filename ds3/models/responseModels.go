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

import (
    "errors"
    "fmt"
    "bytes"
    "log"
)

type PhysicalPlacement struct {
    AzureTargets []AzureTarget
    Ds3Targets []Ds3Target
    Pools []Pool
    S3Targets []S3Target
    Tapes []Tape
}

func (physicalPlacement *PhysicalPlacement) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AzureTargets":
            physicalPlacement.AzureTargets = parseAzureTargetSlice("AzureTarget", child.Children, aggErr)
        case "Ds3Targets":
            physicalPlacement.Ds3Targets = parseDs3TargetSlice("Ds3Target", child.Children, aggErr)
        case "Pools":
            physicalPlacement.Pools = parsePoolSlice("Pool", child.Children, aggErr)
        case "S3Targets":
            physicalPlacement.S3Targets = parseS3TargetSlice("S3Target", child.Children, aggErr)
        case "Tapes":
            physicalPlacement.Tapes = parseTapeSlice("Tape", child.Children, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing PhysicalPlacement.", child.XMLName.Local)
        }
    }
}

type AutoInspectMode Enum

const (
    AUTO_INSPECT_MODE_NEVER AutoInspectMode = 1 + iota
    AUTO_INSPECT_MODE_MINIMAL AutoInspectMode = 1 + iota
    AUTO_INSPECT_MODE_FULL AutoInspectMode = 1 + iota
)

func (autoInspectMode *AutoInspectMode) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *autoInspectMode = UNDEFINED
        case "NEVER": *autoInspectMode = AUTO_INSPECT_MODE_NEVER
        case "MINIMAL": *autoInspectMode = AUTO_INSPECT_MODE_MINIMAL
        case "FULL": *autoInspectMode = AUTO_INSPECT_MODE_FULL
        default:
            *autoInspectMode = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into AutoInspectMode", str))
    }
    return nil
}

func (autoInspectMode AutoInspectMode) String() string {
    switch autoInspectMode {
        case AUTO_INSPECT_MODE_NEVER: return "NEVER"
        case AUTO_INSPECT_MODE_MINIMAL: return "MINIMAL"
        case AUTO_INSPECT_MODE_FULL: return "FULL"
        default:
            log.Printf("Error: invalid AutoInspectMode represented by '%d'", autoInspectMode)
            return ""
    }
}

func (autoInspectMode AutoInspectMode) StringPtr() *string {
    if autoInspectMode == UNDEFINED {
        return nil
    }
    result := autoInspectMode.String()
    return &result
}

func newAutoInspectModeFromContent(content []byte, aggErr *AggregateError) *AutoInspectMode {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(AutoInspectMode)
    parseEnum(content, result, aggErr)
    return result
}
type AzureDataReplicationRule struct {
    DataPolicyId string
    Id string
    MaxBlobPartSizeInBytes int64
    ReplicateDeletes bool
    State DataPlacementRuleState
    TargetId string
    Type DataReplicationRuleType
}

func (azureDataReplicationRule *AzureDataReplicationRule) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DataPolicyId":
            azureDataReplicationRule.DataPolicyId = parseString(child.Content)
        case "Id":
            azureDataReplicationRule.Id = parseString(child.Content)
        case "MaxBlobPartSizeInBytes":
            azureDataReplicationRule.MaxBlobPartSizeInBytes = parseInt64(child.Content, aggErr)
        case "ReplicateDeletes":
            azureDataReplicationRule.ReplicateDeletes = parseBool(child.Content, aggErr)
        case "State":
            parseEnum(child.Content, &azureDataReplicationRule.State, aggErr)
        case "TargetId":
            azureDataReplicationRule.TargetId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &azureDataReplicationRule.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureDataReplicationRule.", child.XMLName.Local)
        }
    }
}

type Blob struct {
    ByteOffset int64
    Checksum *string
    ChecksumType *ChecksumType
    Id string
    Length int64
    ObjectId string
}

func (blob *Blob) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "ByteOffset":
            blob.ByteOffset = parseInt64(child.Content, aggErr)
        case "Checksum":
            blob.Checksum = parseNullableString(child.Content)
        case "ChecksumType":
            blob.ChecksumType = newChecksumTypeFromContent(child.Content, aggErr)
        case "Id":
            blob.Id = parseString(child.Content)
        case "Length":
            blob.Length = parseInt64(child.Content, aggErr)
        case "ObjectId":
            blob.ObjectId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Blob.", child.XMLName.Local)
        }
    }
}

type Priority Enum

const (
    PRIORITY_CRITICAL Priority = 1 + iota
    PRIORITY_URGENT Priority = 1 + iota
    PRIORITY_HIGH Priority = 1 + iota
    PRIORITY_NORMAL Priority = 1 + iota
    PRIORITY_LOW Priority = 1 + iota
    PRIORITY_BACKGROUND Priority = 1 + iota
)

func (priority *Priority) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *priority = UNDEFINED
        case "CRITICAL": *priority = PRIORITY_CRITICAL
        case "URGENT": *priority = PRIORITY_URGENT
        case "HIGH": *priority = PRIORITY_HIGH
        case "NORMAL": *priority = PRIORITY_NORMAL
        case "LOW": *priority = PRIORITY_LOW
        case "BACKGROUND": *priority = PRIORITY_BACKGROUND
        default:
            *priority = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into Priority", str))
    }
    return nil
}

func (priority Priority) String() string {
    switch priority {
        case PRIORITY_CRITICAL: return "CRITICAL"
        case PRIORITY_URGENT: return "URGENT"
        case PRIORITY_HIGH: return "HIGH"
        case PRIORITY_NORMAL: return "NORMAL"
        case PRIORITY_LOW: return "LOW"
        case PRIORITY_BACKGROUND: return "BACKGROUND"
        default:
            log.Printf("Error: invalid Priority represented by '%d'", priority)
            return ""
    }
}

func (priority Priority) StringPtr() *string {
    if priority == UNDEFINED {
        return nil
    }
    result := priority.String()
    return &result
}

func newPriorityFromContent(content []byte, aggErr *AggregateError) *Priority {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(Priority)
    parseEnum(content, result, aggErr)
    return result
}
type Bucket struct {
    CreationDate string
    DataPolicyId string
    Empty *bool
    Id string
    LastPreferredChunkSizeInBytes *int64
    LogicalUsedCapacity *int64
    Name *string
    UserId string
}

func (bucket *Bucket) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            bucket.CreationDate = parseString(child.Content)
        case "DataPolicyId":
            bucket.DataPolicyId = parseString(child.Content)
        case "Empty":
            bucket.Empty = parseNullableBool(child.Content, aggErr)
        case "Id":
            bucket.Id = parseString(child.Content)
        case "LastPreferredChunkSizeInBytes":
            bucket.LastPreferredChunkSizeInBytes = parseNullableInt64(child.Content, aggErr)
        case "LogicalUsedCapacity":
            bucket.LogicalUsedCapacity = parseNullableInt64(child.Content, aggErr)
        case "Name":
            bucket.Name = parseNullableString(child.Content)
        case "UserId":
            bucket.UserId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Bucket.", child.XMLName.Local)
        }
    }
}

type BucketAcl struct {
    BucketId *string
    GroupId *string
    Id string
    Permission BucketAclPermission
    UserId *string
}

func (bucketAcl *BucketAcl) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            bucketAcl.BucketId = parseNullableString(child.Content)
        case "GroupId":
            bucketAcl.GroupId = parseNullableString(child.Content)
        case "Id":
            bucketAcl.Id = parseString(child.Content)
        case "Permission":
            parseEnum(child.Content, &bucketAcl.Permission, aggErr)
        case "UserId":
            bucketAcl.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BucketAcl.", child.XMLName.Local)
        }
    }
}

type BucketAclPermission Enum

const (
    BUCKET_ACL_PERMISSION_LIST BucketAclPermission = 1 + iota
    BUCKET_ACL_PERMISSION_READ BucketAclPermission = 1 + iota
    BUCKET_ACL_PERMISSION_WRITE BucketAclPermission = 1 + iota
    BUCKET_ACL_PERMISSION_DELETE BucketAclPermission = 1 + iota
    BUCKET_ACL_PERMISSION_JOB BucketAclPermission = 1 + iota
    BUCKET_ACL_PERMISSION_OWNER BucketAclPermission = 1 + iota
)

func (bucketAclPermission *BucketAclPermission) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *bucketAclPermission = UNDEFINED
        case "LIST": *bucketAclPermission = BUCKET_ACL_PERMISSION_LIST
        case "READ": *bucketAclPermission = BUCKET_ACL_PERMISSION_READ
        case "WRITE": *bucketAclPermission = BUCKET_ACL_PERMISSION_WRITE
        case "DELETE": *bucketAclPermission = BUCKET_ACL_PERMISSION_DELETE
        case "JOB": *bucketAclPermission = BUCKET_ACL_PERMISSION_JOB
        case "OWNER": *bucketAclPermission = BUCKET_ACL_PERMISSION_OWNER
        default:
            *bucketAclPermission = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into BucketAclPermission", str))
    }
    return nil
}

func (bucketAclPermission BucketAclPermission) String() string {
    switch bucketAclPermission {
        case BUCKET_ACL_PERMISSION_LIST: return "LIST"
        case BUCKET_ACL_PERMISSION_READ: return "READ"
        case BUCKET_ACL_PERMISSION_WRITE: return "WRITE"
        case BUCKET_ACL_PERMISSION_DELETE: return "DELETE"
        case BUCKET_ACL_PERMISSION_JOB: return "JOB"
        case BUCKET_ACL_PERMISSION_OWNER: return "OWNER"
        default:
            log.Printf("Error: invalid BucketAclPermission represented by '%d'", bucketAclPermission)
            return ""
    }
}

func (bucketAclPermission BucketAclPermission) StringPtr() *string {
    if bucketAclPermission == UNDEFINED {
        return nil
    }
    result := bucketAclPermission.String()
    return &result
}

func newBucketAclPermissionFromContent(content []byte, aggErr *AggregateError) *BucketAclPermission {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(BucketAclPermission)
    parseEnum(content, result, aggErr)
    return result
}
type CanceledJob struct {
    BucketId string
    CachedSizeInBytes int64
    CanceledDueToTimeout bool
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    CompletedSizeInBytes int64
    CreatedAt string
    DateCanceled string
    ErrorMessage *string
    Id string
    Naked bool
    Name *string
    OriginalSizeInBytes int64
    Priority Priority
    Rechunked *string
    RequestType JobRequestType
    Truncated bool
    UserId string
}

func (canceledJob *CanceledJob) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            canceledJob.BucketId = parseString(child.Content)
        case "CachedSizeInBytes":
            canceledJob.CachedSizeInBytes = parseInt64(child.Content, aggErr)
        case "CanceledDueToTimeout":
            canceledJob.CanceledDueToTimeout = parseBool(child.Content, aggErr)
        case "ChunkClientProcessingOrderGuarantee":
            parseEnum(child.Content, &canceledJob.ChunkClientProcessingOrderGuarantee, aggErr)
        case "CompletedSizeInBytes":
            canceledJob.CompletedSizeInBytes = parseInt64(child.Content, aggErr)
        case "CreatedAt":
            canceledJob.CreatedAt = parseString(child.Content)
        case "DateCanceled":
            canceledJob.DateCanceled = parseString(child.Content)
        case "ErrorMessage":
            canceledJob.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            canceledJob.Id = parseString(child.Content)
        case "Naked":
            canceledJob.Naked = parseBool(child.Content, aggErr)
        case "Name":
            canceledJob.Name = parseNullableString(child.Content)
        case "OriginalSizeInBytes":
            canceledJob.OriginalSizeInBytes = parseInt64(child.Content, aggErr)
        case "Priority":
            parseEnum(child.Content, &canceledJob.Priority, aggErr)
        case "Rechunked":
            canceledJob.Rechunked = parseNullableString(child.Content)
        case "RequestType":
            parseEnum(child.Content, &canceledJob.RequestType, aggErr)
        case "Truncated":
            canceledJob.Truncated = parseBool(child.Content, aggErr)
        case "UserId":
            canceledJob.UserId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CanceledJob.", child.XMLName.Local)
        }
    }
}

type CapacitySummaryContainer struct {
    Pool StorageDomainCapacitySummary
    Tape StorageDomainCapacitySummary
}

func (capacitySummaryContainer *CapacitySummaryContainer) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Pool":
            capacitySummaryContainer.Pool.parse(&child, aggErr)
        case "Tape":
            capacitySummaryContainer.Tape.parse(&child, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CapacitySummaryContainer.", child.XMLName.Local)
        }
    }
}

type CloudNamingMode Enum

const (
    CLOUD_NAMING_MODE_BLACK_PEARL CloudNamingMode = 1 + iota
    CLOUD_NAMING_MODE_AWS_S3 CloudNamingMode = 1 + iota
)

func (cloudNamingMode *CloudNamingMode) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *cloudNamingMode = UNDEFINED
        case "BLACK_PEARL": *cloudNamingMode = CLOUD_NAMING_MODE_BLACK_PEARL
        case "AWS_S3": *cloudNamingMode = CLOUD_NAMING_MODE_AWS_S3
        default:
            *cloudNamingMode = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into CloudNamingMode", str))
    }
    return nil
}

func (cloudNamingMode CloudNamingMode) String() string {
    switch cloudNamingMode {
        case CLOUD_NAMING_MODE_BLACK_PEARL: return "BLACK_PEARL"
        case CLOUD_NAMING_MODE_AWS_S3: return "AWS_S3"
        default:
            log.Printf("Error: invalid CloudNamingMode represented by '%d'", cloudNamingMode)
            return ""
    }
}

func (cloudNamingMode CloudNamingMode) StringPtr() *string {
    if cloudNamingMode == UNDEFINED {
        return nil
    }
    result := cloudNamingMode.String()
    return &result
}

func newCloudNamingModeFromContent(content []byte, aggErr *AggregateError) *CloudNamingMode {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(CloudNamingMode)
    parseEnum(content, result, aggErr)
    return result
}
type CompletedJob struct {
    BucketId string
    CachedSizeInBytes int64
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    CompletedSizeInBytes int64
    CreatedAt string
    DateCompleted string
    ErrorMessage *string
    Id string
    Naked bool
    Name *string
    OriginalSizeInBytes int64
    Priority Priority
    Rechunked *string
    RequestType JobRequestType
    Truncated bool
    UserId string
}

func (completedJob *CompletedJob) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            completedJob.BucketId = parseString(child.Content)
        case "CachedSizeInBytes":
            completedJob.CachedSizeInBytes = parseInt64(child.Content, aggErr)
        case "ChunkClientProcessingOrderGuarantee":
            parseEnum(child.Content, &completedJob.ChunkClientProcessingOrderGuarantee, aggErr)
        case "CompletedSizeInBytes":
            completedJob.CompletedSizeInBytes = parseInt64(child.Content, aggErr)
        case "CreatedAt":
            completedJob.CreatedAt = parseString(child.Content)
        case "DateCompleted":
            completedJob.DateCompleted = parseString(child.Content)
        case "ErrorMessage":
            completedJob.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            completedJob.Id = parseString(child.Content)
        case "Naked":
            completedJob.Naked = parseBool(child.Content, aggErr)
        case "Name":
            completedJob.Name = parseNullableString(child.Content)
        case "OriginalSizeInBytes":
            completedJob.OriginalSizeInBytes = parseInt64(child.Content, aggErr)
        case "Priority":
            parseEnum(child.Content, &completedJob.Priority, aggErr)
        case "Rechunked":
            completedJob.Rechunked = parseNullableString(child.Content)
        case "RequestType":
            parseEnum(child.Content, &completedJob.RequestType, aggErr)
        case "Truncated":
            completedJob.Truncated = parseBool(child.Content, aggErr)
        case "UserId":
            completedJob.UserId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CompletedJob.", child.XMLName.Local)
        }
    }
}

type DataIsolationLevel Enum

const (
    DATA_ISOLATION_LEVEL_STANDARD DataIsolationLevel = 1 + iota
    DATA_ISOLATION_LEVEL_BUCKET_ISOLATED DataIsolationLevel = 1 + iota
)

func (dataIsolationLevel *DataIsolationLevel) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *dataIsolationLevel = UNDEFINED
        case "STANDARD": *dataIsolationLevel = DATA_ISOLATION_LEVEL_STANDARD
        case "BUCKET_ISOLATED": *dataIsolationLevel = DATA_ISOLATION_LEVEL_BUCKET_ISOLATED
        default:
            *dataIsolationLevel = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into DataIsolationLevel", str))
    }
    return nil
}

func (dataIsolationLevel DataIsolationLevel) String() string {
    switch dataIsolationLevel {
        case DATA_ISOLATION_LEVEL_STANDARD: return "STANDARD"
        case DATA_ISOLATION_LEVEL_BUCKET_ISOLATED: return "BUCKET_ISOLATED"
        default:
            log.Printf("Error: invalid DataIsolationLevel represented by '%d'", dataIsolationLevel)
            return ""
    }
}

func (dataIsolationLevel DataIsolationLevel) StringPtr() *string {
    if dataIsolationLevel == UNDEFINED {
        return nil
    }
    result := dataIsolationLevel.String()
    return &result
}

func newDataIsolationLevelFromContent(content []byte, aggErr *AggregateError) *DataIsolationLevel {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(DataIsolationLevel)
    parseEnum(content, result, aggErr)
    return result
}
type DataPathBackend struct {
    Activated bool
    AllowNewJobRequests bool
    AutoActivateTimeoutInMins *int
    AutoInspect AutoInspectMode
    CacheAvailableRetryAfterInSeconds int
    DefaultVerifyDataAfterImport *Priority
    DefaultVerifyDataPriorToImport bool
    Id string
    InstanceId string
    IomEnabled bool
    LastHeartbeat string
    PartiallyVerifyLastPercentOfTapes *int
    UnavailableMediaPolicy UnavailableMediaUsagePolicy
    UnavailablePoolMaxJobRetryInMins int
    UnavailableTapePartitionMaxJobRetryInMins int
}

func (dataPathBackend *DataPathBackend) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Activated":
            dataPathBackend.Activated = parseBool(child.Content, aggErr)
        case "AllowNewJobRequests":
            dataPathBackend.AllowNewJobRequests = parseBool(child.Content, aggErr)
        case "AutoActivateTimeoutInMins":
            dataPathBackend.AutoActivateTimeoutInMins = parseNullableInt(child.Content, aggErr)
        case "AutoInspect":
            parseEnum(child.Content, &dataPathBackend.AutoInspect, aggErr)
        case "CacheAvailableRetryAfterInSeconds":
            dataPathBackend.CacheAvailableRetryAfterInSeconds = parseInt(child.Content, aggErr)
        case "DefaultVerifyDataAfterImport":
            dataPathBackend.DefaultVerifyDataAfterImport = newPriorityFromContent(child.Content, aggErr)
        case "DefaultVerifyDataPriorToImport":
            dataPathBackend.DefaultVerifyDataPriorToImport = parseBool(child.Content, aggErr)
        case "Id":
            dataPathBackend.Id = parseString(child.Content)
        case "InstanceId":
            dataPathBackend.InstanceId = parseString(child.Content)
        case "IomEnabled":
            dataPathBackend.IomEnabled = parseBool(child.Content, aggErr)
        case "LastHeartbeat":
            dataPathBackend.LastHeartbeat = parseString(child.Content)
        case "PartiallyVerifyLastPercentOfTapes":
            dataPathBackend.PartiallyVerifyLastPercentOfTapes = parseNullableInt(child.Content, aggErr)
        case "UnavailableMediaPolicy":
            parseEnum(child.Content, &dataPathBackend.UnavailableMediaPolicy, aggErr)
        case "UnavailablePoolMaxJobRetryInMins":
            dataPathBackend.UnavailablePoolMaxJobRetryInMins = parseInt(child.Content, aggErr)
        case "UnavailableTapePartitionMaxJobRetryInMins":
            dataPathBackend.UnavailableTapePartitionMaxJobRetryInMins = parseInt(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DataPathBackend.", child.XMLName.Local)
        }
    }
}

type DataPersistenceRule struct {
    DataPolicyId string
    Id string
    IsolationLevel DataIsolationLevel
    MinimumDaysToRetain *int
    State DataPlacementRuleState
    StorageDomainId string
    Type DataPersistenceRuleType
}

func (dataPersistenceRule *DataPersistenceRule) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DataPolicyId":
            dataPersistenceRule.DataPolicyId = parseString(child.Content)
        case "Id":
            dataPersistenceRule.Id = parseString(child.Content)
        case "IsolationLevel":
            parseEnum(child.Content, &dataPersistenceRule.IsolationLevel, aggErr)
        case "MinimumDaysToRetain":
            dataPersistenceRule.MinimumDaysToRetain = parseNullableInt(child.Content, aggErr)
        case "State":
            parseEnum(child.Content, &dataPersistenceRule.State, aggErr)
        case "StorageDomainId":
            dataPersistenceRule.StorageDomainId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &dataPersistenceRule.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DataPersistenceRule.", child.XMLName.Local)
        }
    }
}

type DataPersistenceRuleType Enum

const (
    DATA_PERSISTENCE_RULE_TYPE_PERMANENT DataPersistenceRuleType = 1 + iota
    DATA_PERSISTENCE_RULE_TYPE_TEMPORARY DataPersistenceRuleType = 1 + iota
    DATA_PERSISTENCE_RULE_TYPE_RETIRED DataPersistenceRuleType = 1 + iota
)

func (dataPersistenceRuleType *DataPersistenceRuleType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *dataPersistenceRuleType = UNDEFINED
        case "PERMANENT": *dataPersistenceRuleType = DATA_PERSISTENCE_RULE_TYPE_PERMANENT
        case "TEMPORARY": *dataPersistenceRuleType = DATA_PERSISTENCE_RULE_TYPE_TEMPORARY
        case "RETIRED": *dataPersistenceRuleType = DATA_PERSISTENCE_RULE_TYPE_RETIRED
        default:
            *dataPersistenceRuleType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into DataPersistenceRuleType", str))
    }
    return nil
}

func (dataPersistenceRuleType DataPersistenceRuleType) String() string {
    switch dataPersistenceRuleType {
        case DATA_PERSISTENCE_RULE_TYPE_PERMANENT: return "PERMANENT"
        case DATA_PERSISTENCE_RULE_TYPE_TEMPORARY: return "TEMPORARY"
        case DATA_PERSISTENCE_RULE_TYPE_RETIRED: return "RETIRED"
        default:
            log.Printf("Error: invalid DataPersistenceRuleType represented by '%d'", dataPersistenceRuleType)
            return ""
    }
}

func (dataPersistenceRuleType DataPersistenceRuleType) StringPtr() *string {
    if dataPersistenceRuleType == UNDEFINED {
        return nil
    }
    result := dataPersistenceRuleType.String()
    return &result
}

func newDataPersistenceRuleTypeFromContent(content []byte, aggErr *AggregateError) *DataPersistenceRuleType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(DataPersistenceRuleType)
    parseEnum(content, result, aggErr)
    return result
}
type DataPlacementRuleState Enum

const (
    DATA_PLACEMENT_RULE_STATE_NORMAL DataPlacementRuleState = 1 + iota
    DATA_PLACEMENT_RULE_STATE_INCLUSION_IN_PROGRESS DataPlacementRuleState = 1 + iota
)

func (dataPlacementRuleState *DataPlacementRuleState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *dataPlacementRuleState = UNDEFINED
        case "NORMAL": *dataPlacementRuleState = DATA_PLACEMENT_RULE_STATE_NORMAL
        case "INCLUSION_IN_PROGRESS": *dataPlacementRuleState = DATA_PLACEMENT_RULE_STATE_INCLUSION_IN_PROGRESS
        default:
            *dataPlacementRuleState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into DataPlacementRuleState", str))
    }
    return nil
}

func (dataPlacementRuleState DataPlacementRuleState) String() string {
    switch dataPlacementRuleState {
        case DATA_PLACEMENT_RULE_STATE_NORMAL: return "NORMAL"
        case DATA_PLACEMENT_RULE_STATE_INCLUSION_IN_PROGRESS: return "INCLUSION_IN_PROGRESS"
        default:
            log.Printf("Error: invalid DataPlacementRuleState represented by '%d'", dataPlacementRuleState)
            return ""
    }
}

func (dataPlacementRuleState DataPlacementRuleState) StringPtr() *string {
    if dataPlacementRuleState == UNDEFINED {
        return nil
    }
    result := dataPlacementRuleState.String()
    return &result
}

func newDataPlacementRuleStateFromContent(content []byte, aggErr *AggregateError) *DataPlacementRuleState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(DataPlacementRuleState)
    parseEnum(content, result, aggErr)
    return result
}
type DataPolicy struct {
    AlwaysForcePutJobCreation bool
    AlwaysMinimizeSpanningAcrossMedia bool
    BlobbingEnabled bool
    ChecksumType ChecksumType
    CreationDate string
    DefaultBlobSize *int64
    DefaultGetJobPriority Priority
    DefaultPutJobPriority Priority
    DefaultVerifyAfterWrite bool
    DefaultVerifyJobPriority Priority
    EndToEndCrcRequired bool
    Id string
    MaxVersionsToKeep int
    Name *string
    RebuildPriority Priority
    Versioning VersioningLevel
}

func (dataPolicy *DataPolicy) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AlwaysForcePutJobCreation":
            dataPolicy.AlwaysForcePutJobCreation = parseBool(child.Content, aggErr)
        case "AlwaysMinimizeSpanningAcrossMedia":
            dataPolicy.AlwaysMinimizeSpanningAcrossMedia = parseBool(child.Content, aggErr)
        case "BlobbingEnabled":
            dataPolicy.BlobbingEnabled = parseBool(child.Content, aggErr)
        case "ChecksumType":
            parseEnum(child.Content, &dataPolicy.ChecksumType, aggErr)
        case "CreationDate":
            dataPolicy.CreationDate = parseString(child.Content)
        case "DefaultBlobSize":
            dataPolicy.DefaultBlobSize = parseNullableInt64(child.Content, aggErr)
        case "DefaultGetJobPriority":
            parseEnum(child.Content, &dataPolicy.DefaultGetJobPriority, aggErr)
        case "DefaultPutJobPriority":
            parseEnum(child.Content, &dataPolicy.DefaultPutJobPriority, aggErr)
        case "DefaultVerifyAfterWrite":
            dataPolicy.DefaultVerifyAfterWrite = parseBool(child.Content, aggErr)
        case "DefaultVerifyJobPriority":
            parseEnum(child.Content, &dataPolicy.DefaultVerifyJobPriority, aggErr)
        case "EndToEndCrcRequired":
            dataPolicy.EndToEndCrcRequired = parseBool(child.Content, aggErr)
        case "Id":
            dataPolicy.Id = parseString(child.Content)
        case "MaxVersionsToKeep":
            dataPolicy.MaxVersionsToKeep = parseInt(child.Content, aggErr)
        case "Name":
            dataPolicy.Name = parseNullableString(child.Content)
        case "RebuildPriority":
            parseEnum(child.Content, &dataPolicy.RebuildPriority, aggErr)
        case "Versioning":
            parseEnum(child.Content, &dataPolicy.Versioning, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DataPolicy.", child.XMLName.Local)
        }
    }
}

type DataPolicyAcl struct {
    DataPolicyId *string
    GroupId *string
    Id string
    UserId *string
}

func (dataPolicyAcl *DataPolicyAcl) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DataPolicyId":
            dataPolicyAcl.DataPolicyId = parseNullableString(child.Content)
        case "GroupId":
            dataPolicyAcl.GroupId = parseNullableString(child.Content)
        case "Id":
            dataPolicyAcl.Id = parseString(child.Content)
        case "UserId":
            dataPolicyAcl.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DataPolicyAcl.", child.XMLName.Local)
        }
    }
}

type DataReplicationRuleType Enum

const (
    DATA_REPLICATION_RULE_TYPE_PERMANENT DataReplicationRuleType = 1 + iota
    DATA_REPLICATION_RULE_TYPE_RETIRED DataReplicationRuleType = 1 + iota
)

func (dataReplicationRuleType *DataReplicationRuleType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *dataReplicationRuleType = UNDEFINED
        case "PERMANENT": *dataReplicationRuleType = DATA_REPLICATION_RULE_TYPE_PERMANENT
        case "RETIRED": *dataReplicationRuleType = DATA_REPLICATION_RULE_TYPE_RETIRED
        default:
            *dataReplicationRuleType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into DataReplicationRuleType", str))
    }
    return nil
}

func (dataReplicationRuleType DataReplicationRuleType) String() string {
    switch dataReplicationRuleType {
        case DATA_REPLICATION_RULE_TYPE_PERMANENT: return "PERMANENT"
        case DATA_REPLICATION_RULE_TYPE_RETIRED: return "RETIRED"
        default:
            log.Printf("Error: invalid DataReplicationRuleType represented by '%d'", dataReplicationRuleType)
            return ""
    }
}

func (dataReplicationRuleType DataReplicationRuleType) StringPtr() *string {
    if dataReplicationRuleType == UNDEFINED {
        return nil
    }
    result := dataReplicationRuleType.String()
    return &result
}

func newDataReplicationRuleTypeFromContent(content []byte, aggErr *AggregateError) *DataReplicationRuleType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(DataReplicationRuleType)
    parseEnum(content, result, aggErr)
    return result
}
type DegradedBlob struct {
    AzureReplicationRuleId *string
    BlobId string
    BucketId string
    Ds3ReplicationRuleId *string
    Id string
    PersistenceRuleId *string
    S3ReplicationRuleId *string
}

func (degradedBlob *DegradedBlob) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AzureReplicationRuleId":
            degradedBlob.AzureReplicationRuleId = parseNullableString(child.Content)
        case "BlobId":
            degradedBlob.BlobId = parseString(child.Content)
        case "BucketId":
            degradedBlob.BucketId = parseString(child.Content)
        case "Ds3ReplicationRuleId":
            degradedBlob.Ds3ReplicationRuleId = parseNullableString(child.Content)
        case "Id":
            degradedBlob.Id = parseString(child.Content)
        case "PersistenceRuleId":
            degradedBlob.PersistenceRuleId = parseNullableString(child.Content)
        case "S3ReplicationRuleId":
            degradedBlob.S3ReplicationRuleId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DegradedBlob.", child.XMLName.Local)
        }
    }
}

type Ds3DataReplicationRule struct {
    DataPolicyId string
    Id string
    ReplicateDeletes bool
    State DataPlacementRuleState
    TargetDataPolicy *string
    TargetId string
    Type DataReplicationRuleType
}

func (ds3DataReplicationRule *Ds3DataReplicationRule) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DataPolicyId":
            ds3DataReplicationRule.DataPolicyId = parseString(child.Content)
        case "Id":
            ds3DataReplicationRule.Id = parseString(child.Content)
        case "ReplicateDeletes":
            ds3DataReplicationRule.ReplicateDeletes = parseBool(child.Content, aggErr)
        case "State":
            parseEnum(child.Content, &ds3DataReplicationRule.State, aggErr)
        case "TargetDataPolicy":
            ds3DataReplicationRule.TargetDataPolicy = parseNullableString(child.Content)
        case "TargetId":
            ds3DataReplicationRule.TargetId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &ds3DataReplicationRule.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3DataReplicationRule.", child.XMLName.Local)
        }
    }
}

type FeatureKey struct {
    CurrentValue *int64
    ErrorMessage *string
    ExpirationDate *string
    Id string
    Key FeatureKeyType
    LimitValue *int64
}

func (featureKey *FeatureKey) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CurrentValue":
            featureKey.CurrentValue = parseNullableInt64(child.Content, aggErr)
        case "ErrorMessage":
            featureKey.ErrorMessage = parseNullableString(child.Content)
        case "ExpirationDate":
            featureKey.ExpirationDate = parseNullableString(child.Content)
        case "Id":
            featureKey.Id = parseString(child.Content)
        case "Key":
            parseEnum(child.Content, &featureKey.Key, aggErr)
        case "LimitValue":
            featureKey.LimitValue = parseNullableInt64(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing FeatureKey.", child.XMLName.Local)
        }
    }
}

type FeatureKeyType Enum

const (
    FEATURE_KEY_TYPE_AWS_S3_CLOUD_OUT FeatureKeyType = 1 + iota
    FEATURE_KEY_TYPE_MICROSOFT_AZURE_CLOUD_OUT FeatureKeyType = 1 + iota
)

func (featureKeyType *FeatureKeyType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *featureKeyType = UNDEFINED
        case "AWS_S3_CLOUD_OUT": *featureKeyType = FEATURE_KEY_TYPE_AWS_S3_CLOUD_OUT
        case "MICROSOFT_AZURE_CLOUD_OUT": *featureKeyType = FEATURE_KEY_TYPE_MICROSOFT_AZURE_CLOUD_OUT
        default:
            *featureKeyType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into FeatureKeyType", str))
    }
    return nil
}

func (featureKeyType FeatureKeyType) String() string {
    switch featureKeyType {
        case FEATURE_KEY_TYPE_AWS_S3_CLOUD_OUT: return "AWS_S3_CLOUD_OUT"
        case FEATURE_KEY_TYPE_MICROSOFT_AZURE_CLOUD_OUT: return "MICROSOFT_AZURE_CLOUD_OUT"
        default:
            log.Printf("Error: invalid FeatureKeyType represented by '%d'", featureKeyType)
            return ""
    }
}

func (featureKeyType FeatureKeyType) StringPtr() *string {
    if featureKeyType == UNDEFINED {
        return nil
    }
    result := featureKeyType.String()
    return &result
}

func newFeatureKeyTypeFromContent(content []byte, aggErr *AggregateError) *FeatureKeyType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(FeatureKeyType)
    parseEnum(content, result, aggErr)
    return result
}
type Group struct {
    BuiltIn bool
    Id string
    Name *string
}

func (group *Group) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BuiltIn":
            group.BuiltIn = parseBool(child.Content, aggErr)
        case "Id":
            group.Id = parseString(child.Content)
        case "Name":
            group.Name = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Group.", child.XMLName.Local)
        }
    }
}

type GroupMember struct {
    GroupId string
    Id string
    MemberGroupId *string
    MemberUserId *string
}

func (groupMember *GroupMember) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "GroupId":
            groupMember.GroupId = parseString(child.Content)
        case "Id":
            groupMember.Id = parseString(child.Content)
        case "MemberGroupId":
            groupMember.MemberGroupId = parseNullableString(child.Content)
        case "MemberUserId":
            groupMember.MemberUserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing GroupMember.", child.XMLName.Local)
        }
    }
}

type ActiveJob struct {
    Aggregating bool
    BucketId string
    CachedSizeInBytes int64
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    CompletedSizeInBytes int64
    CreatedAt string
    DeadJobCleanupAllowed bool
    ErrorMessage *string
    Id string
    ImplicitJobIdResolution bool
    MinimizeSpanningAcrossMedia bool
    Naked bool
    Name *string
    OriginalSizeInBytes int64
    Priority Priority
    Rechunked *string
    Replicating bool
    RequestType JobRequestType
    Restore JobRestore
    Truncated bool
    TruncatedDueToTimeout bool
    UserId string
    VerifyAfterWrite bool
}

func (activeJob *ActiveJob) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Aggregating":
            activeJob.Aggregating = parseBool(child.Content, aggErr)
        case "BucketId":
            activeJob.BucketId = parseString(child.Content)
        case "CachedSizeInBytes":
            activeJob.CachedSizeInBytes = parseInt64(child.Content, aggErr)
        case "ChunkClientProcessingOrderGuarantee":
            parseEnum(child.Content, &activeJob.ChunkClientProcessingOrderGuarantee, aggErr)
        case "CompletedSizeInBytes":
            activeJob.CompletedSizeInBytes = parseInt64(child.Content, aggErr)
        case "CreatedAt":
            activeJob.CreatedAt = parseString(child.Content)
        case "DeadJobCleanupAllowed":
            activeJob.DeadJobCleanupAllowed = parseBool(child.Content, aggErr)
        case "ErrorMessage":
            activeJob.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            activeJob.Id = parseString(child.Content)
        case "ImplicitJobIdResolution":
            activeJob.ImplicitJobIdResolution = parseBool(child.Content, aggErr)
        case "MinimizeSpanningAcrossMedia":
            activeJob.MinimizeSpanningAcrossMedia = parseBool(child.Content, aggErr)
        case "Naked":
            activeJob.Naked = parseBool(child.Content, aggErr)
        case "Name":
            activeJob.Name = parseNullableString(child.Content)
        case "OriginalSizeInBytes":
            activeJob.OriginalSizeInBytes = parseInt64(child.Content, aggErr)
        case "Priority":
            parseEnum(child.Content, &activeJob.Priority, aggErr)
        case "Rechunked":
            activeJob.Rechunked = parseNullableString(child.Content)
        case "Replicating":
            activeJob.Replicating = parseBool(child.Content, aggErr)
        case "RequestType":
            parseEnum(child.Content, &activeJob.RequestType, aggErr)
        case "Restore":
            parseEnum(child.Content, &activeJob.Restore, aggErr)
        case "Truncated":
            activeJob.Truncated = parseBool(child.Content, aggErr)
        case "TruncatedDueToTimeout":
            activeJob.TruncatedDueToTimeout = parseBool(child.Content, aggErr)
        case "UserId":
            activeJob.UserId = parseString(child.Content)
        case "VerifyAfterWrite":
            activeJob.VerifyAfterWrite = parseBool(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing ActiveJob.", child.XMLName.Local)
        }
    }
}

type JobChunk struct {
    BlobStoreState JobChunkBlobStoreState
    ChunkNumber int
    Id string
    JobCreationDate string
    JobId string
    NodeId *string
    PendingTargetCommit bool
    ReadFromAzureTargetId *string
    ReadFromDs3TargetId *string
    ReadFromPoolId *string
    ReadFromS3TargetId *string
    ReadFromTapeId *string
}

func (jobChunk *JobChunk) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BlobStoreState":
            parseEnum(child.Content, &jobChunk.BlobStoreState, aggErr)
        case "ChunkNumber":
            jobChunk.ChunkNumber = parseInt(child.Content, aggErr)
        case "Id":
            jobChunk.Id = parseString(child.Content)
        case "JobCreationDate":
            jobChunk.JobCreationDate = parseString(child.Content)
        case "JobId":
            jobChunk.JobId = parseString(child.Content)
        case "NodeId":
            jobChunk.NodeId = parseNullableString(child.Content)
        case "PendingTargetCommit":
            jobChunk.PendingTargetCommit = parseBool(child.Content, aggErr)
        case "ReadFromAzureTargetId":
            jobChunk.ReadFromAzureTargetId = parseNullableString(child.Content)
        case "ReadFromDs3TargetId":
            jobChunk.ReadFromDs3TargetId = parseNullableString(child.Content)
        case "ReadFromPoolId":
            jobChunk.ReadFromPoolId = parseNullableString(child.Content)
        case "ReadFromS3TargetId":
            jobChunk.ReadFromS3TargetId = parseNullableString(child.Content)
        case "ReadFromTapeId":
            jobChunk.ReadFromTapeId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing JobChunk.", child.XMLName.Local)
        }
    }
}

type JobChunkBlobStoreState Enum

const (
    JOB_CHUNK_BLOB_STORE_STATE_PENDING JobChunkBlobStoreState = 1 + iota
    JOB_CHUNK_BLOB_STORE_STATE_IN_PROGRESS JobChunkBlobStoreState = 1 + iota
    JOB_CHUNK_BLOB_STORE_STATE_COMPLETED JobChunkBlobStoreState = 1 + iota
)

func (jobChunkBlobStoreState *JobChunkBlobStoreState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *jobChunkBlobStoreState = UNDEFINED
        case "PENDING": *jobChunkBlobStoreState = JOB_CHUNK_BLOB_STORE_STATE_PENDING
        case "IN_PROGRESS": *jobChunkBlobStoreState = JOB_CHUNK_BLOB_STORE_STATE_IN_PROGRESS
        case "COMPLETED": *jobChunkBlobStoreState = JOB_CHUNK_BLOB_STORE_STATE_COMPLETED
        default:
            *jobChunkBlobStoreState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into JobChunkBlobStoreState", str))
    }
    return nil
}

func (jobChunkBlobStoreState JobChunkBlobStoreState) String() string {
    switch jobChunkBlobStoreState {
        case JOB_CHUNK_BLOB_STORE_STATE_PENDING: return "PENDING"
        case JOB_CHUNK_BLOB_STORE_STATE_IN_PROGRESS: return "IN_PROGRESS"
        case JOB_CHUNK_BLOB_STORE_STATE_COMPLETED: return "COMPLETED"
        default:
            log.Printf("Error: invalid JobChunkBlobStoreState represented by '%d'", jobChunkBlobStoreState)
            return ""
    }
}

func (jobChunkBlobStoreState JobChunkBlobStoreState) StringPtr() *string {
    if jobChunkBlobStoreState == UNDEFINED {
        return nil
    }
    result := jobChunkBlobStoreState.String()
    return &result
}

func newJobChunkBlobStoreStateFromContent(content []byte, aggErr *AggregateError) *JobChunkBlobStoreState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(JobChunkBlobStoreState)
    parseEnum(content, result, aggErr)
    return result
}
type JobChunkClientProcessingOrderGuarantee Enum

const (
    JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_NONE JobChunkClientProcessingOrderGuarantee = 1 + iota
    JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER JobChunkClientProcessingOrderGuarantee = 1 + iota
)

func (jobChunkClientProcessingOrderGuarantee *JobChunkClientProcessingOrderGuarantee) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *jobChunkClientProcessingOrderGuarantee = UNDEFINED
        case "NONE": *jobChunkClientProcessingOrderGuarantee = JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_NONE
        case "IN_ORDER": *jobChunkClientProcessingOrderGuarantee = JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER
        default:
            *jobChunkClientProcessingOrderGuarantee = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into JobChunkClientProcessingOrderGuarantee", str))
    }
    return nil
}

func (jobChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) String() string {
    switch jobChunkClientProcessingOrderGuarantee {
        case JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_NONE: return "NONE"
        case JOB_CHUNK_CLIENT_PROCESSING_ORDER_GUARANTEE_IN_ORDER: return "IN_ORDER"
        default:
            log.Printf("Error: invalid JobChunkClientProcessingOrderGuarantee represented by '%d'", jobChunkClientProcessingOrderGuarantee)
            return ""
    }
}

func (jobChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) StringPtr() *string {
    if jobChunkClientProcessingOrderGuarantee == UNDEFINED {
        return nil
    }
    result := jobChunkClientProcessingOrderGuarantee.String()
    return &result
}

func newJobChunkClientProcessingOrderGuaranteeFromContent(content []byte, aggErr *AggregateError) *JobChunkClientProcessingOrderGuarantee {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(JobChunkClientProcessingOrderGuarantee)
    parseEnum(content, result, aggErr)
    return result
}
type JobRequestType Enum

const (
    JOB_REQUEST_TYPE_PUT JobRequestType = 1 + iota
    JOB_REQUEST_TYPE_GET JobRequestType = 1 + iota
    JOB_REQUEST_TYPE_VERIFY JobRequestType = 1 + iota
)

func (jobRequestType *JobRequestType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *jobRequestType = UNDEFINED
        case "PUT": *jobRequestType = JOB_REQUEST_TYPE_PUT
        case "GET": *jobRequestType = JOB_REQUEST_TYPE_GET
        case "VERIFY": *jobRequestType = JOB_REQUEST_TYPE_VERIFY
        default:
            *jobRequestType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into JobRequestType", str))
    }
    return nil
}

func (jobRequestType JobRequestType) String() string {
    switch jobRequestType {
        case JOB_REQUEST_TYPE_PUT: return "PUT"
        case JOB_REQUEST_TYPE_GET: return "GET"
        case JOB_REQUEST_TYPE_VERIFY: return "VERIFY"
        default:
            log.Printf("Error: invalid JobRequestType represented by '%d'", jobRequestType)
            return ""
    }
}

func (jobRequestType JobRequestType) StringPtr() *string {
    if jobRequestType == UNDEFINED {
        return nil
    }
    result := jobRequestType.String()
    return &result
}

func newJobRequestTypeFromContent(content []byte, aggErr *AggregateError) *JobRequestType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(JobRequestType)
    parseEnum(content, result, aggErr)
    return result
}
type JobRestore Enum

const (
    JOB_RESTORE_NO JobRestore = 1 + iota
    JOB_RESTORE_YES JobRestore = 1 + iota
    JOB_RESTORE_PERMANENT_ONLY JobRestore = 1 + iota
)

func (jobRestore *JobRestore) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *jobRestore = UNDEFINED
        case "NO": *jobRestore = JOB_RESTORE_NO
        case "YES": *jobRestore = JOB_RESTORE_YES
        case "PERMANENT_ONLY": *jobRestore = JOB_RESTORE_PERMANENT_ONLY
        default:
            *jobRestore = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into JobRestore", str))
    }
    return nil
}

func (jobRestore JobRestore) String() string {
    switch jobRestore {
        case JOB_RESTORE_NO: return "NO"
        case JOB_RESTORE_YES: return "YES"
        case JOB_RESTORE_PERMANENT_ONLY: return "PERMANENT_ONLY"
        default:
            log.Printf("Error: invalid JobRestore represented by '%d'", jobRestore)
            return ""
    }
}

func (jobRestore JobRestore) StringPtr() *string {
    if jobRestore == UNDEFINED {
        return nil
    }
    result := jobRestore.String()
    return &result
}

func newJobRestoreFromContent(content []byte, aggErr *AggregateError) *JobRestore {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(JobRestore)
    parseEnum(content, result, aggErr)
    return result
}
type LtfsFileNamingMode Enum

const (
    LTFS_FILE_NAMING_MODE_OBJECT_NAME LtfsFileNamingMode = 1 + iota
    LTFS_FILE_NAMING_MODE_OBJECT_ID LtfsFileNamingMode = 1 + iota
)

func (ltfsFileNamingMode *LtfsFileNamingMode) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *ltfsFileNamingMode = UNDEFINED
        case "OBJECT_NAME": *ltfsFileNamingMode = LTFS_FILE_NAMING_MODE_OBJECT_NAME
        case "OBJECT_ID": *ltfsFileNamingMode = LTFS_FILE_NAMING_MODE_OBJECT_ID
        default:
            *ltfsFileNamingMode = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into LtfsFileNamingMode", str))
    }
    return nil
}

func (ltfsFileNamingMode LtfsFileNamingMode) String() string {
    switch ltfsFileNamingMode {
        case LTFS_FILE_NAMING_MODE_OBJECT_NAME: return "OBJECT_NAME"
        case LTFS_FILE_NAMING_MODE_OBJECT_ID: return "OBJECT_ID"
        default:
            log.Printf("Error: invalid LtfsFileNamingMode represented by '%d'", ltfsFileNamingMode)
            return ""
    }
}

func (ltfsFileNamingMode LtfsFileNamingMode) StringPtr() *string {
    if ltfsFileNamingMode == UNDEFINED {
        return nil
    }
    result := ltfsFileNamingMode.String()
    return &result
}

func newLtfsFileNamingModeFromContent(content []byte, aggErr *AggregateError) *LtfsFileNamingMode {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(LtfsFileNamingMode)
    parseEnum(content, result, aggErr)
    return result
}
type Node struct {
    DataPathHttpPort *int
    DataPathHttpsPort *int
    DataPathIpAddress *string
    DnsName *string
    Id string
    LastHeartbeat string
    Name *string
    SerialNumber *string
}

func (node *Node) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DataPathHttpPort":
            node.DataPathHttpPort = parseNullableInt(child.Content, aggErr)
        case "DataPathHttpsPort":
            node.DataPathHttpsPort = parseNullableInt(child.Content, aggErr)
        case "DataPathIpAddress":
            node.DataPathIpAddress = parseNullableString(child.Content)
        case "DnsName":
            node.DnsName = parseNullableString(child.Content)
        case "Id":
            node.Id = parseString(child.Content)
        case "LastHeartbeat":
            node.LastHeartbeat = parseString(child.Content)
        case "Name":
            node.Name = parseNullableString(child.Content)
        case "SerialNumber":
            node.SerialNumber = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Node.", child.XMLName.Local)
        }
    }
}

type S3DataReplicationRule struct {
    DataPolicyId string
    Id string
    InitialDataPlacement S3InitialDataPlacementPolicy
    MaxBlobPartSizeInBytes int64
    ReplicateDeletes bool
    State DataPlacementRuleState
    TargetId string
    Type DataReplicationRuleType
}

func (s3DataReplicationRule *S3DataReplicationRule) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DataPolicyId":
            s3DataReplicationRule.DataPolicyId = parseString(child.Content)
        case "Id":
            s3DataReplicationRule.Id = parseString(child.Content)
        case "InitialDataPlacement":
            parseEnum(child.Content, &s3DataReplicationRule.InitialDataPlacement, aggErr)
        case "MaxBlobPartSizeInBytes":
            s3DataReplicationRule.MaxBlobPartSizeInBytes = parseInt64(child.Content, aggErr)
        case "ReplicateDeletes":
            s3DataReplicationRule.ReplicateDeletes = parseBool(child.Content, aggErr)
        case "State":
            parseEnum(child.Content, &s3DataReplicationRule.State, aggErr)
        case "TargetId":
            s3DataReplicationRule.TargetId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &s3DataReplicationRule.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3DataReplicationRule.", child.XMLName.Local)
        }
    }
}

type S3InitialDataPlacementPolicy Enum

const (
    S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD S3InitialDataPlacementPolicy = 1 + iota
    S3_INITIAL_DATA_PLACEMENT_POLICY_REDUCED_REDUNDANCY S3InitialDataPlacementPolicy = 1 + iota
    S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD_IA S3InitialDataPlacementPolicy = 1 + iota
    S3_INITIAL_DATA_PLACEMENT_POLICY_GLACIER S3InitialDataPlacementPolicy = 1 + iota
    S3_INITIAL_DATA_PLACEMENT_POLICY_DEEP_ARCHIVE S3InitialDataPlacementPolicy = 1 + iota
)

func (s3InitialDataPlacementPolicy *S3InitialDataPlacementPolicy) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *s3InitialDataPlacementPolicy = UNDEFINED
        case "STANDARD": *s3InitialDataPlacementPolicy = S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD
        case "REDUCED_REDUNDANCY": *s3InitialDataPlacementPolicy = S3_INITIAL_DATA_PLACEMENT_POLICY_REDUCED_REDUNDANCY
        case "STANDARD_IA": *s3InitialDataPlacementPolicy = S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD_IA
        case "GLACIER": *s3InitialDataPlacementPolicy = S3_INITIAL_DATA_PLACEMENT_POLICY_GLACIER
        case "DEEP_ARCHIVE": *s3InitialDataPlacementPolicy = S3_INITIAL_DATA_PLACEMENT_POLICY_DEEP_ARCHIVE
        default:
            *s3InitialDataPlacementPolicy = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into S3InitialDataPlacementPolicy", str))
    }
    return nil
}

func (s3InitialDataPlacementPolicy S3InitialDataPlacementPolicy) String() string {
    switch s3InitialDataPlacementPolicy {
        case S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD: return "STANDARD"
        case S3_INITIAL_DATA_PLACEMENT_POLICY_REDUCED_REDUNDANCY: return "REDUCED_REDUNDANCY"
        case S3_INITIAL_DATA_PLACEMENT_POLICY_STANDARD_IA: return "STANDARD_IA"
        case S3_INITIAL_DATA_PLACEMENT_POLICY_GLACIER: return "GLACIER"
        case S3_INITIAL_DATA_PLACEMENT_POLICY_DEEP_ARCHIVE: return "DEEP_ARCHIVE"
        default:
            log.Printf("Error: invalid S3InitialDataPlacementPolicy represented by '%d'", s3InitialDataPlacementPolicy)
            return ""
    }
}

func (s3InitialDataPlacementPolicy S3InitialDataPlacementPolicy) StringPtr() *string {
    if s3InitialDataPlacementPolicy == UNDEFINED {
        return nil
    }
    result := s3InitialDataPlacementPolicy.String()
    return &result
}

func newS3InitialDataPlacementPolicyFromContent(content []byte, aggErr *AggregateError) *S3InitialDataPlacementPolicy {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(S3InitialDataPlacementPolicy)
    parseEnum(content, result, aggErr)
    return result
}
type S3Object struct {
    BucketId string
    CreationDate *string
    Id string
    Latest bool
    Name *string
    Type S3ObjectType
}

func (s3Object *S3Object) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            s3Object.BucketId = parseString(child.Content)
        case "CreationDate":
            s3Object.CreationDate = parseNullableString(child.Content)
        case "Id":
            s3Object.Id = parseString(child.Content)
        case "Latest":
            s3Object.Latest = parseBool(child.Content, aggErr)
        case "Name":
            s3Object.Name = parseNullableString(child.Content)
        case "Type":
            parseEnum(child.Content, &s3Object.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3Object.", child.XMLName.Local)
        }
    }
}

type S3ObjectType Enum

const (
    S3_OBJECT_TYPE_DATA S3ObjectType = 1 + iota
    S3_OBJECT_TYPE_FOLDER S3ObjectType = 1 + iota
)

func (s3ObjectType *S3ObjectType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *s3ObjectType = UNDEFINED
        case "DATA": *s3ObjectType = S3_OBJECT_TYPE_DATA
        case "FOLDER": *s3ObjectType = S3_OBJECT_TYPE_FOLDER
        default:
            *s3ObjectType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into S3ObjectType", str))
    }
    return nil
}

func (s3ObjectType S3ObjectType) String() string {
    switch s3ObjectType {
        case S3_OBJECT_TYPE_DATA: return "DATA"
        case S3_OBJECT_TYPE_FOLDER: return "FOLDER"
        default:
            log.Printf("Error: invalid S3ObjectType represented by '%d'", s3ObjectType)
            return ""
    }
}

func (s3ObjectType S3ObjectType) StringPtr() *string {
    if s3ObjectType == UNDEFINED {
        return nil
    }
    result := s3ObjectType.String()
    return &result
}

func newS3ObjectTypeFromContent(content []byte, aggErr *AggregateError) *S3ObjectType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(S3ObjectType)
    parseEnum(content, result, aggErr)
    return result
}
type S3Region Enum

const (
    S3_REGION_GOV_CLOUD S3Region = 1 + iota
    S3_REGION_US_EAST_1 S3Region = 1 + iota
    S3_REGION_US_EAST_2 S3Region = 1 + iota
    S3_REGION_US_WEST_1 S3Region = 1 + iota
    S3_REGION_US_WEST_2 S3Region = 1 + iota
    S3_REGION_EU_WEST_1 S3Region = 1 + iota
    S3_REGION_EU_WEST_2 S3Region = 1 + iota
    S3_REGION_EU_CENTRAL_1 S3Region = 1 + iota
    S3_REGION_AP_SOUTH_1 S3Region = 1 + iota
    S3_REGION_AP_SOUTHEAST_1 S3Region = 1 + iota
    S3_REGION_AP_SOUTHEAST_2 S3Region = 1 + iota
    S3_REGION_AP_NORTHEAST_1 S3Region = 1 + iota
    S3_REGION_AP_NORTHEAST_2 S3Region = 1 + iota
    S3_REGION_SA_EAST_1 S3Region = 1 + iota
    S3_REGION_CN_NORTH_1 S3Region = 1 + iota
    S3_REGION_CA_CENTRAL_1 S3Region = 1 + iota
)

func (s3Region *S3Region) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *s3Region = UNDEFINED
        case "GOV_CLOUD": *s3Region = S3_REGION_GOV_CLOUD
        case "US_EAST_1": *s3Region = S3_REGION_US_EAST_1
        case "US_EAST_2": *s3Region = S3_REGION_US_EAST_2
        case "US_WEST_1": *s3Region = S3_REGION_US_WEST_1
        case "US_WEST_2": *s3Region = S3_REGION_US_WEST_2
        case "EU_WEST_1": *s3Region = S3_REGION_EU_WEST_1
        case "EU_WEST_2": *s3Region = S3_REGION_EU_WEST_2
        case "EU_CENTRAL_1": *s3Region = S3_REGION_EU_CENTRAL_1
        case "AP_SOUTH_1": *s3Region = S3_REGION_AP_SOUTH_1
        case "AP_SOUTHEAST_1": *s3Region = S3_REGION_AP_SOUTHEAST_1
        case "AP_SOUTHEAST_2": *s3Region = S3_REGION_AP_SOUTHEAST_2
        case "AP_NORTHEAST_1": *s3Region = S3_REGION_AP_NORTHEAST_1
        case "AP_NORTHEAST_2": *s3Region = S3_REGION_AP_NORTHEAST_2
        case "SA_EAST_1": *s3Region = S3_REGION_SA_EAST_1
        case "CN_NORTH_1": *s3Region = S3_REGION_CN_NORTH_1
        case "CA_CENTRAL_1": *s3Region = S3_REGION_CA_CENTRAL_1
        default:
            *s3Region = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into S3Region", str))
    }
    return nil
}

func (s3Region S3Region) String() string {
    switch s3Region {
        case S3_REGION_GOV_CLOUD: return "GOV_CLOUD"
        case S3_REGION_US_EAST_1: return "US_EAST_1"
        case S3_REGION_US_EAST_2: return "US_EAST_2"
        case S3_REGION_US_WEST_1: return "US_WEST_1"
        case S3_REGION_US_WEST_2: return "US_WEST_2"
        case S3_REGION_EU_WEST_1: return "EU_WEST_1"
        case S3_REGION_EU_WEST_2: return "EU_WEST_2"
        case S3_REGION_EU_CENTRAL_1: return "EU_CENTRAL_1"
        case S3_REGION_AP_SOUTH_1: return "AP_SOUTH_1"
        case S3_REGION_AP_SOUTHEAST_1: return "AP_SOUTHEAST_1"
        case S3_REGION_AP_SOUTHEAST_2: return "AP_SOUTHEAST_2"
        case S3_REGION_AP_NORTHEAST_1: return "AP_NORTHEAST_1"
        case S3_REGION_AP_NORTHEAST_2: return "AP_NORTHEAST_2"
        case S3_REGION_SA_EAST_1: return "SA_EAST_1"
        case S3_REGION_CN_NORTH_1: return "CN_NORTH_1"
        case S3_REGION_CA_CENTRAL_1: return "CA_CENTRAL_1"
        default:
            log.Printf("Error: invalid S3Region represented by '%d'", s3Region)
            return ""
    }
}

func (s3Region S3Region) StringPtr() *string {
    if s3Region == UNDEFINED {
        return nil
    }
    result := s3Region.String()
    return &result
}

func newS3RegionFromContent(content []byte, aggErr *AggregateError) *S3Region {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(S3Region)
    parseEnum(content, result, aggErr)
    return result
}
type StorageDomain struct {
    AutoEjectMediaFullThreshold *int64
    AutoEjectUponCron *string
    AutoEjectUponJobCancellation bool
    AutoEjectUponJobCompletion bool
    AutoEjectUponMediaFull bool
    Id string
    LtfsFileNaming LtfsFileNamingMode
    MaxTapeFragmentationPercent int
    MaximumAutoVerificationFrequencyInDays *int
    MediaEjectionAllowed bool
    Name *string
    SecureMediaAllocation bool
    VerifyPriorToAutoEject *Priority
    WriteOptimization WriteOptimization
}

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
            storageDomain.VerifyPriorToAutoEject = newPriorityFromContent(child.Content, aggErr)
        case "WriteOptimization":
            parseEnum(child.Content, &storageDomain.WriteOptimization, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing StorageDomain.", child.XMLName.Local)
        }
    }
}

type StorageDomainCapacitySummary struct {
    PhysicalAllocated int64
    PhysicalFree int64
    PhysicalUsed int64
}

func (storageDomainCapacitySummary *StorageDomainCapacitySummary) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "PhysicalAllocated":
            storageDomainCapacitySummary.PhysicalAllocated = parseInt64(child.Content, aggErr)
        case "PhysicalFree":
            storageDomainCapacitySummary.PhysicalFree = parseInt64(child.Content, aggErr)
        case "PhysicalUsed":
            storageDomainCapacitySummary.PhysicalUsed = parseInt64(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing StorageDomainCapacitySummary.", child.XMLName.Local)
        }
    }
}

type StorageDomainFailure struct {
    Date string
    ErrorMessage *string
    Id string
    StorageDomainId string
    Type StorageDomainFailureType
}

func (storageDomainFailure *StorageDomainFailure) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Date":
            storageDomainFailure.Date = parseString(child.Content)
        case "ErrorMessage":
            storageDomainFailure.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            storageDomainFailure.Id = parseString(child.Content)
        case "StorageDomainId":
            storageDomainFailure.StorageDomainId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &storageDomainFailure.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing StorageDomainFailure.", child.XMLName.Local)
        }
    }
}

type StorageDomainFailureType Enum

const (
    STORAGE_DOMAIN_FAILURE_TYPE_ILLEGAL_EJECTION_OCCURRED StorageDomainFailureType = 1 + iota
    STORAGE_DOMAIN_FAILURE_TYPE_MEMBER_BECAME_READ_ONLY StorageDomainFailureType = 1 + iota
    STORAGE_DOMAIN_FAILURE_TYPE_WRITES_STALLED_DUE_TO_NO_FREE_MEDIA_REMAINING StorageDomainFailureType = 1 + iota
)

func (storageDomainFailureType *StorageDomainFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *storageDomainFailureType = UNDEFINED
        case "ILLEGAL_EJECTION_OCCURRED": *storageDomainFailureType = STORAGE_DOMAIN_FAILURE_TYPE_ILLEGAL_EJECTION_OCCURRED
        case "MEMBER_BECAME_READ_ONLY": *storageDomainFailureType = STORAGE_DOMAIN_FAILURE_TYPE_MEMBER_BECAME_READ_ONLY
        case "WRITES_STALLED_DUE_TO_NO_FREE_MEDIA_REMAINING": *storageDomainFailureType = STORAGE_DOMAIN_FAILURE_TYPE_WRITES_STALLED_DUE_TO_NO_FREE_MEDIA_REMAINING
        default:
            *storageDomainFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into StorageDomainFailureType", str))
    }
    return nil
}

func (storageDomainFailureType StorageDomainFailureType) String() string {
    switch storageDomainFailureType {
        case STORAGE_DOMAIN_FAILURE_TYPE_ILLEGAL_EJECTION_OCCURRED: return "ILLEGAL_EJECTION_OCCURRED"
        case STORAGE_DOMAIN_FAILURE_TYPE_MEMBER_BECAME_READ_ONLY: return "MEMBER_BECAME_READ_ONLY"
        case STORAGE_DOMAIN_FAILURE_TYPE_WRITES_STALLED_DUE_TO_NO_FREE_MEDIA_REMAINING: return "WRITES_STALLED_DUE_TO_NO_FREE_MEDIA_REMAINING"
        default:
            log.Printf("Error: invalid StorageDomainFailureType represented by '%d'", storageDomainFailureType)
            return ""
    }
}

func (storageDomainFailureType StorageDomainFailureType) StringPtr() *string {
    if storageDomainFailureType == UNDEFINED {
        return nil
    }
    result := storageDomainFailureType.String()
    return &result
}

func newStorageDomainFailureTypeFromContent(content []byte, aggErr *AggregateError) *StorageDomainFailureType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(StorageDomainFailureType)
    parseEnum(content, result, aggErr)
    return result
}
type StorageDomainMember struct {
    AutoCompactionThreshold *int
    Id string
    PoolPartitionId *string
    State StorageDomainMemberState
    StorageDomainId string
    TapePartitionId *string
    TapeType *string
    WritePreference WritePreferenceLevel
}

func (storageDomainMember *StorageDomainMember) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AutoCompactionThreshold":
            storageDomainMember.AutoCompactionThreshold = parseNullableInt(child.Content, aggErr)
        case "Id":
            storageDomainMember.Id = parseString(child.Content)
        case "PoolPartitionId":
            storageDomainMember.PoolPartitionId = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &storageDomainMember.State, aggErr)
        case "StorageDomainId":
            storageDomainMember.StorageDomainId = parseString(child.Content)
        case "TapePartitionId":
            storageDomainMember.TapePartitionId = parseNullableString(child.Content)
        case "TapeType":
            storageDomainMember.TapeType = parseNullableString(child.Content)
        case "WritePreference":
            parseEnum(child.Content, &storageDomainMember.WritePreference, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing StorageDomainMember.", child.XMLName.Local)
        }
    }
}

type StorageDomainMemberState Enum

const (
    STORAGE_DOMAIN_MEMBER_STATE_NORMAL StorageDomainMemberState = 1 + iota
    STORAGE_DOMAIN_MEMBER_STATE_EXCLUSION_IN_PROGRESS StorageDomainMemberState = 1 + iota
)

func (storageDomainMemberState *StorageDomainMemberState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *storageDomainMemberState = UNDEFINED
        case "NORMAL": *storageDomainMemberState = STORAGE_DOMAIN_MEMBER_STATE_NORMAL
        case "EXCLUSION_IN_PROGRESS": *storageDomainMemberState = STORAGE_DOMAIN_MEMBER_STATE_EXCLUSION_IN_PROGRESS
        default:
            *storageDomainMemberState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into StorageDomainMemberState", str))
    }
    return nil
}

func (storageDomainMemberState StorageDomainMemberState) String() string {
    switch storageDomainMemberState {
        case STORAGE_DOMAIN_MEMBER_STATE_NORMAL: return "NORMAL"
        case STORAGE_DOMAIN_MEMBER_STATE_EXCLUSION_IN_PROGRESS: return "EXCLUSION_IN_PROGRESS"
        default:
            log.Printf("Error: invalid StorageDomainMemberState represented by '%d'", storageDomainMemberState)
            return ""
    }
}

func (storageDomainMemberState StorageDomainMemberState) StringPtr() *string {
    if storageDomainMemberState == UNDEFINED {
        return nil
    }
    result := storageDomainMemberState.String()
    return &result
}

func newStorageDomainMemberStateFromContent(content []byte, aggErr *AggregateError) *StorageDomainMemberState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(StorageDomainMemberState)
    parseEnum(content, result, aggErr)
    return result
}
type SystemFailure struct {
    Date string
    ErrorMessage *string
    Id string
    Type SystemFailureType
}

func (systemFailure *SystemFailure) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Date":
            systemFailure.Date = parseString(child.Content)
        case "ErrorMessage":
            systemFailure.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            systemFailure.Id = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &systemFailure.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SystemFailure.", child.XMLName.Local)
        }
    }
}

type SystemFailureType Enum

const (
    SYSTEM_FAILURE_TYPE_RECONCILE_TAPE_ENVIRONMENT_FAILED SystemFailureType = 1 + iota
    SYSTEM_FAILURE_TYPE_RECONCILE_POOL_ENVIRONMENT_FAILED SystemFailureType = 1 + iota
    SYSTEM_FAILURE_TYPE_CRITICAL_DATA_VERIFICATION_ERROR_REQUIRES_USER_CONFIRMATION SystemFailureType = 1 + iota
    SYSTEM_FAILURE_TYPE_MICROSOFT_AZURE_WRITES_REQUIRE_FEATURE_LICENSE SystemFailureType = 1 + iota
    SYSTEM_FAILURE_TYPE_AWS_S3_WRITES_REQUIRE_FEATURE_LICENSE SystemFailureType = 1 + iota
    SYSTEM_FAILURE_TYPE_DATABASE_RUNNING_OUT_OF_SPACE SystemFailureType = 1 + iota
)

func (systemFailureType *SystemFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *systemFailureType = UNDEFINED
        case "RECONCILE_TAPE_ENVIRONMENT_FAILED": *systemFailureType = SYSTEM_FAILURE_TYPE_RECONCILE_TAPE_ENVIRONMENT_FAILED
        case "RECONCILE_POOL_ENVIRONMENT_FAILED": *systemFailureType = SYSTEM_FAILURE_TYPE_RECONCILE_POOL_ENVIRONMENT_FAILED
        case "CRITICAL_DATA_VERIFICATION_ERROR_REQUIRES_USER_CONFIRMATION": *systemFailureType = SYSTEM_FAILURE_TYPE_CRITICAL_DATA_VERIFICATION_ERROR_REQUIRES_USER_CONFIRMATION
        case "MICROSOFT_AZURE_WRITES_REQUIRE_FEATURE_LICENSE": *systemFailureType = SYSTEM_FAILURE_TYPE_MICROSOFT_AZURE_WRITES_REQUIRE_FEATURE_LICENSE
        case "AWS_S3_WRITES_REQUIRE_FEATURE_LICENSE": *systemFailureType = SYSTEM_FAILURE_TYPE_AWS_S3_WRITES_REQUIRE_FEATURE_LICENSE
        case "DATABASE_RUNNING_OUT_OF_SPACE": *systemFailureType = SYSTEM_FAILURE_TYPE_DATABASE_RUNNING_OUT_OF_SPACE
        default:
            *systemFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into SystemFailureType", str))
    }
    return nil
}

func (systemFailureType SystemFailureType) String() string {
    switch systemFailureType {
        case SYSTEM_FAILURE_TYPE_RECONCILE_TAPE_ENVIRONMENT_FAILED: return "RECONCILE_TAPE_ENVIRONMENT_FAILED"
        case SYSTEM_FAILURE_TYPE_RECONCILE_POOL_ENVIRONMENT_FAILED: return "RECONCILE_POOL_ENVIRONMENT_FAILED"
        case SYSTEM_FAILURE_TYPE_CRITICAL_DATA_VERIFICATION_ERROR_REQUIRES_USER_CONFIRMATION: return "CRITICAL_DATA_VERIFICATION_ERROR_REQUIRES_USER_CONFIRMATION"
        case SYSTEM_FAILURE_TYPE_MICROSOFT_AZURE_WRITES_REQUIRE_FEATURE_LICENSE: return "MICROSOFT_AZURE_WRITES_REQUIRE_FEATURE_LICENSE"
        case SYSTEM_FAILURE_TYPE_AWS_S3_WRITES_REQUIRE_FEATURE_LICENSE: return "AWS_S3_WRITES_REQUIRE_FEATURE_LICENSE"
        case SYSTEM_FAILURE_TYPE_DATABASE_RUNNING_OUT_OF_SPACE: return "DATABASE_RUNNING_OUT_OF_SPACE"
        default:
            log.Printf("Error: invalid SystemFailureType represented by '%d'", systemFailureType)
            return ""
    }
}

func (systemFailureType SystemFailureType) StringPtr() *string {
    if systemFailureType == UNDEFINED {
        return nil
    }
    result := systemFailureType.String()
    return &result
}

func newSystemFailureTypeFromContent(content []byte, aggErr *AggregateError) *SystemFailureType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(SystemFailureType)
    parseEnum(content, result, aggErr)
    return result
}
type UnavailableMediaUsagePolicy Enum

const (
    UNAVAILABLE_MEDIA_USAGE_POLICY_ALLOW UnavailableMediaUsagePolicy = 1 + iota
    UNAVAILABLE_MEDIA_USAGE_POLICY_DISCOURAGED UnavailableMediaUsagePolicy = 1 + iota
    UNAVAILABLE_MEDIA_USAGE_POLICY_DISALLOW UnavailableMediaUsagePolicy = 1 + iota
)

func (unavailableMediaUsagePolicy *UnavailableMediaUsagePolicy) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *unavailableMediaUsagePolicy = UNDEFINED
        case "ALLOW": *unavailableMediaUsagePolicy = UNAVAILABLE_MEDIA_USAGE_POLICY_ALLOW
        case "DISCOURAGED": *unavailableMediaUsagePolicy = UNAVAILABLE_MEDIA_USAGE_POLICY_DISCOURAGED
        case "DISALLOW": *unavailableMediaUsagePolicy = UNAVAILABLE_MEDIA_USAGE_POLICY_DISALLOW
        default:
            *unavailableMediaUsagePolicy = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into UnavailableMediaUsagePolicy", str))
    }
    return nil
}

func (unavailableMediaUsagePolicy UnavailableMediaUsagePolicy) String() string {
    switch unavailableMediaUsagePolicy {
        case UNAVAILABLE_MEDIA_USAGE_POLICY_ALLOW: return "ALLOW"
        case UNAVAILABLE_MEDIA_USAGE_POLICY_DISCOURAGED: return "DISCOURAGED"
        case UNAVAILABLE_MEDIA_USAGE_POLICY_DISALLOW: return "DISALLOW"
        default:
            log.Printf("Error: invalid UnavailableMediaUsagePolicy represented by '%d'", unavailableMediaUsagePolicy)
            return ""
    }
}

func (unavailableMediaUsagePolicy UnavailableMediaUsagePolicy) StringPtr() *string {
    if unavailableMediaUsagePolicy == UNDEFINED {
        return nil
    }
    result := unavailableMediaUsagePolicy.String()
    return &result
}

func newUnavailableMediaUsagePolicyFromContent(content []byte, aggErr *AggregateError) *UnavailableMediaUsagePolicy {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(UnavailableMediaUsagePolicy)
    parseEnum(content, result, aggErr)
    return result
}
type SpectraUser struct {
    AuthId *string
    DefaultDataPolicyId *string
    Id string
    MaxBuckets int
    Name *string
    SecretKey *string
}

func (spectraUser *SpectraUser) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AuthId":
            spectraUser.AuthId = parseNullableString(child.Content)
        case "DefaultDataPolicyId":
            spectraUser.DefaultDataPolicyId = parseNullableString(child.Content)
        case "Id":
            spectraUser.Id = parseString(child.Content)
        case "MaxBuckets":
            spectraUser.MaxBuckets = parseInt(child.Content, aggErr)
        case "Name":
            spectraUser.Name = parseNullableString(child.Content)
        case "SecretKey":
            spectraUser.SecretKey = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SpectraUser.", child.XMLName.Local)
        }
    }
}

type VersioningLevel Enum

const (
    VERSIONING_LEVEL_NONE VersioningLevel = 1 + iota
    VERSIONING_LEVEL_KEEP_LATEST VersioningLevel = 1 + iota
    VERSIONING_LEVEL_KEEP_MULTIPLE_VERSIONS VersioningLevel = 1 + iota
)

func (versioningLevel *VersioningLevel) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *versioningLevel = UNDEFINED
        case "NONE": *versioningLevel = VERSIONING_LEVEL_NONE
        case "KEEP_LATEST": *versioningLevel = VERSIONING_LEVEL_KEEP_LATEST
        case "KEEP_MULTIPLE_VERSIONS": *versioningLevel = VERSIONING_LEVEL_KEEP_MULTIPLE_VERSIONS
        default:
            *versioningLevel = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into VersioningLevel", str))
    }
    return nil
}

func (versioningLevel VersioningLevel) String() string {
    switch versioningLevel {
        case VERSIONING_LEVEL_NONE: return "NONE"
        case VERSIONING_LEVEL_KEEP_LATEST: return "KEEP_LATEST"
        case VERSIONING_LEVEL_KEEP_MULTIPLE_VERSIONS: return "KEEP_MULTIPLE_VERSIONS"
        default:
            log.Printf("Error: invalid VersioningLevel represented by '%d'", versioningLevel)
            return ""
    }
}

func (versioningLevel VersioningLevel) StringPtr() *string {
    if versioningLevel == UNDEFINED {
        return nil
    }
    result := versioningLevel.String()
    return &result
}

func newVersioningLevelFromContent(content []byte, aggErr *AggregateError) *VersioningLevel {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(VersioningLevel)
    parseEnum(content, result, aggErr)
    return result
}
type WriteOptimization Enum

const (
    WRITE_OPTIMIZATION_CAPACITY WriteOptimization = 1 + iota
    WRITE_OPTIMIZATION_PERFORMANCE WriteOptimization = 1 + iota
)

func (writeOptimization *WriteOptimization) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *writeOptimization = UNDEFINED
        case "CAPACITY": *writeOptimization = WRITE_OPTIMIZATION_CAPACITY
        case "PERFORMANCE": *writeOptimization = WRITE_OPTIMIZATION_PERFORMANCE
        default:
            *writeOptimization = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into WriteOptimization", str))
    }
    return nil
}

func (writeOptimization WriteOptimization) String() string {
    switch writeOptimization {
        case WRITE_OPTIMIZATION_CAPACITY: return "CAPACITY"
        case WRITE_OPTIMIZATION_PERFORMANCE: return "PERFORMANCE"
        default:
            log.Printf("Error: invalid WriteOptimization represented by '%d'", writeOptimization)
            return ""
    }
}

func (writeOptimization WriteOptimization) StringPtr() *string {
    if writeOptimization == UNDEFINED {
        return nil
    }
    result := writeOptimization.String()
    return &result
}

func newWriteOptimizationFromContent(content []byte, aggErr *AggregateError) *WriteOptimization {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(WriteOptimization)
    parseEnum(content, result, aggErr)
    return result
}
type WritePreferenceLevel Enum

const (
    WRITE_PREFERENCE_LEVEL_HIGH WritePreferenceLevel = 1 + iota
    WRITE_PREFERENCE_LEVEL_NORMAL WritePreferenceLevel = 1 + iota
    WRITE_PREFERENCE_LEVEL_LOW WritePreferenceLevel = 1 + iota
    WRITE_PREFERENCE_LEVEL_NEVER_SELECT WritePreferenceLevel = 1 + iota
)

func (writePreferenceLevel *WritePreferenceLevel) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *writePreferenceLevel = UNDEFINED
        case "HIGH": *writePreferenceLevel = WRITE_PREFERENCE_LEVEL_HIGH
        case "NORMAL": *writePreferenceLevel = WRITE_PREFERENCE_LEVEL_NORMAL
        case "LOW": *writePreferenceLevel = WRITE_PREFERENCE_LEVEL_LOW
        case "NEVER_SELECT": *writePreferenceLevel = WRITE_PREFERENCE_LEVEL_NEVER_SELECT
        default:
            *writePreferenceLevel = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into WritePreferenceLevel", str))
    }
    return nil
}

func (writePreferenceLevel WritePreferenceLevel) String() string {
    switch writePreferenceLevel {
        case WRITE_PREFERENCE_LEVEL_HIGH: return "HIGH"
        case WRITE_PREFERENCE_LEVEL_NORMAL: return "NORMAL"
        case WRITE_PREFERENCE_LEVEL_LOW: return "LOW"
        case WRITE_PREFERENCE_LEVEL_NEVER_SELECT: return "NEVER_SELECT"
        default:
            log.Printf("Error: invalid WritePreferenceLevel represented by '%d'", writePreferenceLevel)
            return ""
    }
}

func (writePreferenceLevel WritePreferenceLevel) StringPtr() *string {
    if writePreferenceLevel == UNDEFINED {
        return nil
    }
    result := writePreferenceLevel.String()
    return &result
}

func newWritePreferenceLevelFromContent(content []byte, aggErr *AggregateError) *WritePreferenceLevel {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(WritePreferenceLevel)
    parseEnum(content, result, aggErr)
    return result
}
type AzureTargetFailureNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (azureTargetFailureNotificationRegistration *AzureTargetFailureNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            azureTargetFailureNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &azureTargetFailureNotificationRegistration.Format, aggErr)
        case "Id":
            azureTargetFailureNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            azureTargetFailureNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            azureTargetFailureNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            azureTargetFailureNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &azureTargetFailureNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            azureTargetFailureNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &azureTargetFailureNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            azureTargetFailureNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            azureTargetFailureNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureTargetFailureNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type BucketChangesNotificationRegistration struct {
    BucketId *string
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    LastSequenceNumber *int64
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (bucketChangesNotificationRegistration *BucketChangesNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            bucketChangesNotificationRegistration.BucketId = parseNullableString(child.Content)
        case "CreationDate":
            bucketChangesNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &bucketChangesNotificationRegistration.Format, aggErr)
        case "Id":
            bucketChangesNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            bucketChangesNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            bucketChangesNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            bucketChangesNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "LastSequenceNumber":
            bucketChangesNotificationRegistration.LastSequenceNumber = parseNullableInt64(child.Content, aggErr)
        case "NamingConvention":
            parseEnum(child.Content, &bucketChangesNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            bucketChangesNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &bucketChangesNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            bucketChangesNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            bucketChangesNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BucketChangesNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type BucketHistoryEvent struct {
    BucketId string
    Id string
    ObjectCreationDate *string
    ObjectName *string
    SequenceNumber *int64
    Type BucketHistoryEventType
    VersionId string
}

func (bucketHistoryEvent *BucketHistoryEvent) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            bucketHistoryEvent.BucketId = parseString(child.Content)
        case "Id":
            bucketHistoryEvent.Id = parseString(child.Content)
        case "ObjectCreationDate":
            bucketHistoryEvent.ObjectCreationDate = parseNullableString(child.Content)
        case "ObjectName":
            bucketHistoryEvent.ObjectName = parseNullableString(child.Content)
        case "SequenceNumber":
            bucketHistoryEvent.SequenceNumber = parseNullableInt64(child.Content, aggErr)
        case "Type":
            parseEnum(child.Content, &bucketHistoryEvent.Type, aggErr)
        case "VersionId":
            bucketHistoryEvent.VersionId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BucketHistoryEvent.", child.XMLName.Local)
        }
    }
}

type BucketHistoryEventType Enum

const (
    BUCKET_HISTORY_EVENT_TYPE_DELETE BucketHistoryEventType = 1 + iota
    BUCKET_HISTORY_EVENT_TYPE_MARK_LATEST BucketHistoryEventType = 1 + iota
    BUCKET_HISTORY_EVENT_TYPE_UNMARK_LATEST BucketHistoryEventType = 1 + iota
    BUCKET_HISTORY_EVENT_TYPE_CREATE BucketHistoryEventType = 1 + iota
)

func (bucketHistoryEventType *BucketHistoryEventType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *bucketHistoryEventType = UNDEFINED
        case "DELETE": *bucketHistoryEventType = BUCKET_HISTORY_EVENT_TYPE_DELETE
        case "MARK_LATEST": *bucketHistoryEventType = BUCKET_HISTORY_EVENT_TYPE_MARK_LATEST
        case "UNMARK_LATEST": *bucketHistoryEventType = BUCKET_HISTORY_EVENT_TYPE_UNMARK_LATEST
        case "CREATE": *bucketHistoryEventType = BUCKET_HISTORY_EVENT_TYPE_CREATE
        default:
            *bucketHistoryEventType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into BucketHistoryEventType", str))
    }
    return nil
}

func (bucketHistoryEventType BucketHistoryEventType) String() string {
    switch bucketHistoryEventType {
        case BUCKET_HISTORY_EVENT_TYPE_DELETE: return "DELETE"
        case BUCKET_HISTORY_EVENT_TYPE_MARK_LATEST: return "MARK_LATEST"
        case BUCKET_HISTORY_EVENT_TYPE_UNMARK_LATEST: return "UNMARK_LATEST"
        case BUCKET_HISTORY_EVENT_TYPE_CREATE: return "CREATE"
        default:
            log.Printf("Error: invalid BucketHistoryEventType represented by '%d'", bucketHistoryEventType)
            return ""
    }
}

func (bucketHistoryEventType BucketHistoryEventType) StringPtr() *string {
    if bucketHistoryEventType == UNDEFINED {
        return nil
    }
    result := bucketHistoryEventType.String()
    return &result
}

func newBucketHistoryEventTypeFromContent(content []byte, aggErr *AggregateError) *BucketHistoryEventType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(BucketHistoryEventType)
    parseEnum(content, result, aggErr)
    return result
}
type Ds3TargetFailureNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

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
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3TargetFailureNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type JobCompletedNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    JobId *string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

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
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing JobCompletedNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type JobCreatedNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (jobCreatedNotificationRegistration *JobCreatedNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            jobCreatedNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &jobCreatedNotificationRegistration.Format, aggErr)
        case "Id":
            jobCreatedNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            jobCreatedNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            jobCreatedNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            jobCreatedNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &jobCreatedNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            jobCreatedNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &jobCreatedNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            jobCreatedNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            jobCreatedNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing JobCreatedNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type JobCreationFailedNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

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

type PoolFailureNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (poolFailureNotificationRegistration *PoolFailureNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            poolFailureNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &poolFailureNotificationRegistration.Format, aggErr)
        case "Id":
            poolFailureNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            poolFailureNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            poolFailureNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            poolFailureNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &poolFailureNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            poolFailureNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &poolFailureNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            poolFailureNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            poolFailureNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing PoolFailureNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type S3ObjectCachedNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    JobId *string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (s3ObjectCachedNotificationRegistration *S3ObjectCachedNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            s3ObjectCachedNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &s3ObjectCachedNotificationRegistration.Format, aggErr)
        case "Id":
            s3ObjectCachedNotificationRegistration.Id = parseString(child.Content)
        case "JobId":
            s3ObjectCachedNotificationRegistration.JobId = parseNullableString(child.Content)
        case "LastFailure":
            s3ObjectCachedNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            s3ObjectCachedNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            s3ObjectCachedNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &s3ObjectCachedNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            s3ObjectCachedNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &s3ObjectCachedNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            s3ObjectCachedNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            s3ObjectCachedNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3ObjectCachedNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type S3ObjectLostNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (s3ObjectLostNotificationRegistration *S3ObjectLostNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            s3ObjectLostNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &s3ObjectLostNotificationRegistration.Format, aggErr)
        case "Id":
            s3ObjectLostNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            s3ObjectLostNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            s3ObjectLostNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            s3ObjectLostNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &s3ObjectLostNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            s3ObjectLostNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &s3ObjectLostNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            s3ObjectLostNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            s3ObjectLostNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3ObjectLostNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type S3ObjectPersistedNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    JobId *string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (s3ObjectPersistedNotificationRegistration *S3ObjectPersistedNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            s3ObjectPersistedNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &s3ObjectPersistedNotificationRegistration.Format, aggErr)
        case "Id":
            s3ObjectPersistedNotificationRegistration.Id = parseString(child.Content)
        case "JobId":
            s3ObjectPersistedNotificationRegistration.JobId = parseNullableString(child.Content)
        case "LastFailure":
            s3ObjectPersistedNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            s3ObjectPersistedNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            s3ObjectPersistedNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &s3ObjectPersistedNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            s3ObjectPersistedNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &s3ObjectPersistedNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            s3ObjectPersistedNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            s3ObjectPersistedNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3ObjectPersistedNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type S3TargetFailureNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (s3TargetFailureNotificationRegistration *S3TargetFailureNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            s3TargetFailureNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &s3TargetFailureNotificationRegistration.Format, aggErr)
        case "Id":
            s3TargetFailureNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            s3TargetFailureNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            s3TargetFailureNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            s3TargetFailureNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &s3TargetFailureNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            s3TargetFailureNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &s3TargetFailureNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            s3TargetFailureNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            s3TargetFailureNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3TargetFailureNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type StorageDomainFailureNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (storageDomainFailureNotificationRegistration *StorageDomainFailureNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            storageDomainFailureNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &storageDomainFailureNotificationRegistration.Format, aggErr)
        case "Id":
            storageDomainFailureNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            storageDomainFailureNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            storageDomainFailureNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            storageDomainFailureNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &storageDomainFailureNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            storageDomainFailureNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &storageDomainFailureNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            storageDomainFailureNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            storageDomainFailureNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing StorageDomainFailureNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type SystemFailureNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (systemFailureNotificationRegistration *SystemFailureNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            systemFailureNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &systemFailureNotificationRegistration.Format, aggErr)
        case "Id":
            systemFailureNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            systemFailureNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            systemFailureNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            systemFailureNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &systemFailureNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            systemFailureNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &systemFailureNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            systemFailureNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            systemFailureNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SystemFailureNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type TapeFailureNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (tapeFailureNotificationRegistration *TapeFailureNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            tapeFailureNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &tapeFailureNotificationRegistration.Format, aggErr)
        case "Id":
            tapeFailureNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            tapeFailureNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            tapeFailureNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            tapeFailureNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &tapeFailureNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            tapeFailureNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &tapeFailureNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            tapeFailureNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            tapeFailureNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeFailureNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type TapePartitionFailureNotificationRegistration struct {
    CreationDate string
    Format HttpResponseFormatType
    Id string
    LastFailure *string
    LastHttpResponseCode *int
    LastNotification *string
    NamingConvention NamingConventionType
    NotificationEndPoint *string
    NotificationHttpMethod RequestType
    NumberOfFailuresSinceLastSuccess int
    UserId *string
}

func (tapePartitionFailureNotificationRegistration *TapePartitionFailureNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            tapePartitionFailureNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &tapePartitionFailureNotificationRegistration.Format, aggErr)
        case "Id":
            tapePartitionFailureNotificationRegistration.Id = parseString(child.Content)
        case "LastFailure":
            tapePartitionFailureNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            tapePartitionFailureNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            tapePartitionFailureNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &tapePartitionFailureNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            tapePartitionFailureNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &tapePartitionFailureNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            tapePartitionFailureNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            tapePartitionFailureNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapePartitionFailureNotificationRegistration.", child.XMLName.Local)
        }
    }
}

type CacheFilesystem struct {
    AutoReclaimInitiateThreshold float64
    AutoReclaimTerminateThreshold float64
    BurstThreshold float64
    Id string
    MaxCapacityInBytes *int64
    MaxPercentUtilizationOfFilesystem *float64
    NodeId string
    Path *string
}

func (cacheFilesystem *CacheFilesystem) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AutoReclaimInitiateThreshold":
            cacheFilesystem.AutoReclaimInitiateThreshold = parseFloat64(child.Content, aggErr)
        case "AutoReclaimTerminateThreshold":
            cacheFilesystem.AutoReclaimTerminateThreshold = parseFloat64(child.Content, aggErr)
        case "BurstThreshold":
            cacheFilesystem.BurstThreshold = parseFloat64(child.Content, aggErr)
        case "Id":
            cacheFilesystem.Id = parseString(child.Content)
        case "MaxCapacityInBytes":
            cacheFilesystem.MaxCapacityInBytes = parseNullableInt64(child.Content, aggErr)
        case "MaxPercentUtilizationOfFilesystem":
            cacheFilesystem.MaxPercentUtilizationOfFilesystem = parseNullableFloat64(child.Content, aggErr)
        case "NodeId":
            cacheFilesystem.NodeId = parseString(child.Content)
        case "Path":
            cacheFilesystem.Path = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CacheFilesystem.", child.XMLName.Local)
        }
    }
}

type Pool struct {
    AssignedToStorageDomain bool
    AvailableCapacity int64
    BucketId *string
    Guid *string
    Health PoolHealth
    Id string
    LastAccessed *string
    LastModified *string
    LastVerified *string
    Mountpoint *string
    Name *string
    PartitionId *string
    PoweredOn bool
    Quiesced Quiesced
    ReservedCapacity int64
    State PoolState
    StorageDomainMemberId *string
    TotalCapacity int64
    Type PoolType
    UsedCapacity int64
}

func (pool *Pool) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AssignedToStorageDomain":
            pool.AssignedToStorageDomain = parseBool(child.Content, aggErr)
        case "AvailableCapacity":
            pool.AvailableCapacity = parseInt64(child.Content, aggErr)
        case "BucketId":
            pool.BucketId = parseNullableString(child.Content)
        case "Guid":
            pool.Guid = parseNullableString(child.Content)
        case "Health":
            parseEnum(child.Content, &pool.Health, aggErr)
        case "Id":
            pool.Id = parseString(child.Content)
        case "LastAccessed":
            pool.LastAccessed = parseNullableString(child.Content)
        case "LastModified":
            pool.LastModified = parseNullableString(child.Content)
        case "LastVerified":
            pool.LastVerified = parseNullableString(child.Content)
        case "Mountpoint":
            pool.Mountpoint = parseNullableString(child.Content)
        case "Name":
            pool.Name = parseNullableString(child.Content)
        case "PartitionId":
            pool.PartitionId = parseNullableString(child.Content)
        case "PoweredOn":
            pool.PoweredOn = parseBool(child.Content, aggErr)
        case "Quiesced":
            parseEnum(child.Content, &pool.Quiesced, aggErr)
        case "ReservedCapacity":
            pool.ReservedCapacity = parseInt64(child.Content, aggErr)
        case "State":
            parseEnum(child.Content, &pool.State, aggErr)
        case "StorageDomainMemberId":
            pool.StorageDomainMemberId = parseNullableString(child.Content)
        case "TotalCapacity":
            pool.TotalCapacity = parseInt64(child.Content, aggErr)
        case "Type":
            parseEnum(child.Content, &pool.Type, aggErr)
        case "UsedCapacity":
            pool.UsedCapacity = parseInt64(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Pool.", child.XMLName.Local)
        }
    }
}


func parsePoolSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []Pool {
    var result []Pool
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult Pool
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing Pool struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}

type PoolFailure struct {
    Date string
    ErrorMessage *string
    Id string
    PoolId string
    Type PoolFailureType
}

func (poolFailure *PoolFailure) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Date":
            poolFailure.Date = parseString(child.Content)
        case "ErrorMessage":
            poolFailure.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            poolFailure.Id = parseString(child.Content)
        case "PoolId":
            poolFailure.PoolId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &poolFailure.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing PoolFailure.", child.XMLName.Local)
        }
    }
}

type PoolFailureType Enum

const (
    POOL_FAILURE_TYPE_BLOB_READ_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_DATA_CHECKPOINT_MISSING PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_FORMAT_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_IMPORT_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_IMPORT_INCOMPLETE PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_INSPECT_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_QUIESCED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_READ_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_VERIFY_FAILED PoolFailureType = 1 + iota
    POOL_FAILURE_TYPE_WRITE_FAILED PoolFailureType = 1 + iota
)

func (poolFailureType *PoolFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *poolFailureType = UNDEFINED
        case "BLOB_READ_FAILED": *poolFailureType = POOL_FAILURE_TYPE_BLOB_READ_FAILED
        case "DATA_CHECKPOINT_FAILURE": *poolFailureType = POOL_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE
        case "DATA_CHECKPOINT_MISSING": *poolFailureType = POOL_FAILURE_TYPE_DATA_CHECKPOINT_MISSING
        case "FORMAT_FAILED": *poolFailureType = POOL_FAILURE_TYPE_FORMAT_FAILED
        case "IMPORT_FAILED": *poolFailureType = POOL_FAILURE_TYPE_IMPORT_FAILED
        case "IMPORT_INCOMPLETE": *poolFailureType = POOL_FAILURE_TYPE_IMPORT_INCOMPLETE
        case "IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE": *poolFailureType = POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE
        case "IMPORT_FAILED_DUE_TO_DATA_INTEGRITY": *poolFailureType = POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY
        case "INSPECT_FAILED": *poolFailureType = POOL_FAILURE_TYPE_INSPECT_FAILED
        case "QUIESCED": *poolFailureType = POOL_FAILURE_TYPE_QUIESCED
        case "READ_FAILED": *poolFailureType = POOL_FAILURE_TYPE_READ_FAILED
        case "VERIFY_FAILED": *poolFailureType = POOL_FAILURE_TYPE_VERIFY_FAILED
        case "WRITE_FAILED": *poolFailureType = POOL_FAILURE_TYPE_WRITE_FAILED
        default:
            *poolFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into PoolFailureType", str))
    }
    return nil
}

func (poolFailureType PoolFailureType) String() string {
    switch poolFailureType {
        case POOL_FAILURE_TYPE_BLOB_READ_FAILED: return "BLOB_READ_FAILED"
        case POOL_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE: return "DATA_CHECKPOINT_FAILURE"
        case POOL_FAILURE_TYPE_DATA_CHECKPOINT_MISSING: return "DATA_CHECKPOINT_MISSING"
        case POOL_FAILURE_TYPE_FORMAT_FAILED: return "FORMAT_FAILED"
        case POOL_FAILURE_TYPE_IMPORT_FAILED: return "IMPORT_FAILED"
        case POOL_FAILURE_TYPE_IMPORT_INCOMPLETE: return "IMPORT_INCOMPLETE"
        case POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE: return "IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE"
        case POOL_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY: return "IMPORT_FAILED_DUE_TO_DATA_INTEGRITY"
        case POOL_FAILURE_TYPE_INSPECT_FAILED: return "INSPECT_FAILED"
        case POOL_FAILURE_TYPE_QUIESCED: return "QUIESCED"
        case POOL_FAILURE_TYPE_READ_FAILED: return "READ_FAILED"
        case POOL_FAILURE_TYPE_VERIFY_FAILED: return "VERIFY_FAILED"
        case POOL_FAILURE_TYPE_WRITE_FAILED: return "WRITE_FAILED"
        default:
            log.Printf("Error: invalid PoolFailureType represented by '%d'", poolFailureType)
            return ""
    }
}

func (poolFailureType PoolFailureType) StringPtr() *string {
    if poolFailureType == UNDEFINED {
        return nil
    }
    result := poolFailureType.String()
    return &result
}

func newPoolFailureTypeFromContent(content []byte, aggErr *AggregateError) *PoolFailureType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(PoolFailureType)
    parseEnum(content, result, aggErr)
    return result
}
type PoolHealth Enum

const (
    POOL_HEALTH_OK PoolHealth = 1 + iota
    POOL_HEALTH_DEGRADED PoolHealth = 1 + iota
)

func (poolHealth *PoolHealth) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *poolHealth = UNDEFINED
        case "OK": *poolHealth = POOL_HEALTH_OK
        case "DEGRADED": *poolHealth = POOL_HEALTH_DEGRADED
        default:
            *poolHealth = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into PoolHealth", str))
    }
    return nil
}

func (poolHealth PoolHealth) String() string {
    switch poolHealth {
        case POOL_HEALTH_OK: return "OK"
        case POOL_HEALTH_DEGRADED: return "DEGRADED"
        default:
            log.Printf("Error: invalid PoolHealth represented by '%d'", poolHealth)
            return ""
    }
}

func (poolHealth PoolHealth) StringPtr() *string {
    if poolHealth == UNDEFINED {
        return nil
    }
    result := poolHealth.String()
    return &result
}

func newPoolHealthFromContent(content []byte, aggErr *AggregateError) *PoolHealth {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(PoolHealth)
    parseEnum(content, result, aggErr)
    return result
}
type PoolPartition struct {
    Id string
    Name *string
    Type PoolType
}

func (poolPartition *PoolPartition) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Id":
            poolPartition.Id = parseString(child.Content)
        case "Name":
            poolPartition.Name = parseNullableString(child.Content)
        case "Type":
            parseEnum(child.Content, &poolPartition.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing PoolPartition.", child.XMLName.Local)
        }
    }
}

type PoolState Enum

const (
    POOL_STATE_BLANK PoolState = 1 + iota
    POOL_STATE_NORMAL PoolState = 1 + iota
    POOL_STATE_LOST PoolState = 1 + iota
    POOL_STATE_FOREIGN PoolState = 1 + iota
    POOL_STATE_IMPORT_PENDING PoolState = 1 + iota
    POOL_STATE_IMPORT_IN_PROGRESS PoolState = 1 + iota
)

func (poolState *PoolState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *poolState = UNDEFINED
        case "BLANK": *poolState = POOL_STATE_BLANK
        case "NORMAL": *poolState = POOL_STATE_NORMAL
        case "LOST": *poolState = POOL_STATE_LOST
        case "FOREIGN": *poolState = POOL_STATE_FOREIGN
        case "IMPORT_PENDING": *poolState = POOL_STATE_IMPORT_PENDING
        case "IMPORT_IN_PROGRESS": *poolState = POOL_STATE_IMPORT_IN_PROGRESS
        default:
            *poolState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into PoolState", str))
    }
    return nil
}

func (poolState PoolState) String() string {
    switch poolState {
        case POOL_STATE_BLANK: return "BLANK"
        case POOL_STATE_NORMAL: return "NORMAL"
        case POOL_STATE_LOST: return "LOST"
        case POOL_STATE_FOREIGN: return "FOREIGN"
        case POOL_STATE_IMPORT_PENDING: return "IMPORT_PENDING"
        case POOL_STATE_IMPORT_IN_PROGRESS: return "IMPORT_IN_PROGRESS"
        default:
            log.Printf("Error: invalid PoolState represented by '%d'", poolState)
            return ""
    }
}

func (poolState PoolState) StringPtr() *string {
    if poolState == UNDEFINED {
        return nil
    }
    result := poolState.String()
    return &result
}

func newPoolStateFromContent(content []byte, aggErr *AggregateError) *PoolState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(PoolState)
    parseEnum(content, result, aggErr)
    return result
}
type PoolType Enum

const (
    POOL_TYPE_NEARLINE PoolType = 1 + iota
    POOL_TYPE_ONLINE PoolType = 1 + iota
)

func (poolType *PoolType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *poolType = UNDEFINED
        case "NEARLINE": *poolType = POOL_TYPE_NEARLINE
        case "ONLINE": *poolType = POOL_TYPE_ONLINE
        default:
            *poolType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into PoolType", str))
    }
    return nil
}

func (poolType PoolType) String() string {
    switch poolType {
        case POOL_TYPE_NEARLINE: return "NEARLINE"
        case POOL_TYPE_ONLINE: return "ONLINE"
        default:
            log.Printf("Error: invalid PoolType represented by '%d'", poolType)
            return ""
    }
}

func (poolType PoolType) StringPtr() *string {
    if poolType == UNDEFINED {
        return nil
    }
    result := poolType.String()
    return &result
}

func newPoolTypeFromContent(content []byte, aggErr *AggregateError) *PoolType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(PoolType)
    parseEnum(content, result, aggErr)
    return result
}
type SuspectBlobPool struct {
    BlobId string
    BucketId string
    DateWritten string
    Id string
    LastAccessed string
    PoolId string
}

func (suspectBlobPool *SuspectBlobPool) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BlobId":
            suspectBlobPool.BlobId = parseString(child.Content)
        case "BucketId":
            suspectBlobPool.BucketId = parseString(child.Content)
        case "DateWritten":
            suspectBlobPool.DateWritten = parseString(child.Content)
        case "Id":
            suspectBlobPool.Id = parseString(child.Content)
        case "LastAccessed":
            suspectBlobPool.LastAccessed = parseString(child.Content)
        case "PoolId":
            suspectBlobPool.PoolId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobPool.", child.XMLName.Local)
        }
    }
}

type Quiesced Enum

const (
    QUIESCED_NO Quiesced = 1 + iota
    QUIESCED_PENDING Quiesced = 1 + iota
    QUIESCED_YES Quiesced = 1 + iota
)

func (quiesced *Quiesced) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *quiesced = UNDEFINED
        case "NO": *quiesced = QUIESCED_NO
        case "PENDING": *quiesced = QUIESCED_PENDING
        case "YES": *quiesced = QUIESCED_YES
        default:
            *quiesced = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into Quiesced", str))
    }
    return nil
}

func (quiesced Quiesced) String() string {
    switch quiesced {
        case QUIESCED_NO: return "NO"
        case QUIESCED_PENDING: return "PENDING"
        case QUIESCED_YES: return "YES"
        default:
            log.Printf("Error: invalid Quiesced represented by '%d'", quiesced)
            return ""
    }
}

func (quiesced Quiesced) StringPtr() *string {
    if quiesced == UNDEFINED {
        return nil
    }
    result := quiesced.String()
    return &result
}

func newQuiescedFromContent(content []byte, aggErr *AggregateError) *Quiesced {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(Quiesced)
    parseEnum(content, result, aggErr)
    return result
}
type ReservedTaskType Enum

const (
    RESERVED_TASK_TYPE_ANY ReservedTaskType = 1 + iota
    RESERVED_TASK_TYPE_READ ReservedTaskType = 1 + iota
    RESERVED_TASK_TYPE_WRITE ReservedTaskType = 1 + iota
)

func (reservedTaskType *ReservedTaskType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *reservedTaskType = UNDEFINED
        case "ANY": *reservedTaskType = RESERVED_TASK_TYPE_ANY
        case "READ": *reservedTaskType = RESERVED_TASK_TYPE_READ
        case "WRITE": *reservedTaskType = RESERVED_TASK_TYPE_WRITE
        default:
            *reservedTaskType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into ReservedTaskType", str))
    }
    return nil
}

func (reservedTaskType ReservedTaskType) String() string {
    switch reservedTaskType {
        case RESERVED_TASK_TYPE_ANY: return "ANY"
        case RESERVED_TASK_TYPE_READ: return "READ"
        case RESERVED_TASK_TYPE_WRITE: return "WRITE"
        default:
            log.Printf("Error: invalid ReservedTaskType represented by '%d'", reservedTaskType)
            return ""
    }
}

func (reservedTaskType ReservedTaskType) StringPtr() *string {
    if reservedTaskType == UNDEFINED {
        return nil
    }
    result := reservedTaskType.String()
    return &result
}

func newReservedTaskTypeFromContent(content []byte, aggErr *AggregateError) *ReservedTaskType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(ReservedTaskType)
    parseEnum(content, result, aggErr)
    return result
}
type ImportExportConfiguration Enum

const (
    IMPORT_EXPORT_CONFIGURATION_SUPPORTED ImportExportConfiguration = 1 + iota
    IMPORT_EXPORT_CONFIGURATION_NOT_SUPPORTED ImportExportConfiguration = 1 + iota
)

func (importExportConfiguration *ImportExportConfiguration) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *importExportConfiguration = UNDEFINED
        case "SUPPORTED": *importExportConfiguration = IMPORT_EXPORT_CONFIGURATION_SUPPORTED
        case "NOT_SUPPORTED": *importExportConfiguration = IMPORT_EXPORT_CONFIGURATION_NOT_SUPPORTED
        default:
            *importExportConfiguration = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into ImportExportConfiguration", str))
    }
    return nil
}

func (importExportConfiguration ImportExportConfiguration) String() string {
    switch importExportConfiguration {
        case IMPORT_EXPORT_CONFIGURATION_SUPPORTED: return "SUPPORTED"
        case IMPORT_EXPORT_CONFIGURATION_NOT_SUPPORTED: return "NOT_SUPPORTED"
        default:
            log.Printf("Error: invalid ImportExportConfiguration represented by '%d'", importExportConfiguration)
            return ""
    }
}

func (importExportConfiguration ImportExportConfiguration) StringPtr() *string {
    if importExportConfiguration == UNDEFINED {
        return nil
    }
    result := importExportConfiguration.String()
    return &result
}

func newImportExportConfigurationFromContent(content []byte, aggErr *AggregateError) *ImportExportConfiguration {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(ImportExportConfiguration)
    parseEnum(content, result, aggErr)
    return result
}
type SuspectBlobTape struct {
    BlobId string
    Id string
    OrderIndex int
    TapeId string
}

func (suspectBlobTape *SuspectBlobTape) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BlobId":
            suspectBlobTape.BlobId = parseString(child.Content)
        case "Id":
            suspectBlobTape.Id = parseString(child.Content)
        case "OrderIndex":
            suspectBlobTape.OrderIndex = parseInt(child.Content, aggErr)
        case "TapeId":
            suspectBlobTape.TapeId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobTape.", child.XMLName.Local)
        }
    }
}

type Tape struct {
    AssignedToStorageDomain bool
    AvailableRawCapacity *int64
    BarCode *string
    BucketId *string
    DescriptionForIdentification *string
    EjectDate *string
    EjectLabel *string
    EjectLocation *string
    EjectPending *string
    FullOfData bool
    Id string
    LastAccessed *string
    LastCheckpoint *string
    LastModified *string
    LastVerified *string
    PartiallyVerifiedEndOfTape *string
    PartitionId *string
    PreviousState *TapeState
    SerialNumber *string
    State TapeState
    StorageDomainMemberId *string
    TakeOwnershipPending bool
    TotalRawCapacity *int64
    Type string
    VerifyPending *Priority
    WriteProtected bool
}

func (tape *Tape) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AssignedToStorageDomain":
            tape.AssignedToStorageDomain = parseBool(child.Content, aggErr)
        case "AvailableRawCapacity":
            tape.AvailableRawCapacity = parseNullableInt64(child.Content, aggErr)
        case "BarCode":
            tape.BarCode = parseNullableString(child.Content)
        case "BucketId":
            tape.BucketId = parseNullableString(child.Content)
        case "DescriptionForIdentification":
            tape.DescriptionForIdentification = parseNullableString(child.Content)
        case "EjectDate":
            tape.EjectDate = parseNullableString(child.Content)
        case "EjectLabel":
            tape.EjectLabel = parseNullableString(child.Content)
        case "EjectLocation":
            tape.EjectLocation = parseNullableString(child.Content)
        case "EjectPending":
            tape.EjectPending = parseNullableString(child.Content)
        case "FullOfData":
            tape.FullOfData = parseBool(child.Content, aggErr)
        case "Id":
            tape.Id = parseString(child.Content)
        case "LastAccessed":
            tape.LastAccessed = parseNullableString(child.Content)
        case "LastCheckpoint":
            tape.LastCheckpoint = parseNullableString(child.Content)
        case "LastModified":
            tape.LastModified = parseNullableString(child.Content)
        case "LastVerified":
            tape.LastVerified = parseNullableString(child.Content)
        case "PartiallyVerifiedEndOfTape":
            tape.PartiallyVerifiedEndOfTape = parseNullableString(child.Content)
        case "PartitionId":
            tape.PartitionId = parseNullableString(child.Content)
        case "PreviousState":
            tape.PreviousState = newTapeStateFromContent(child.Content, aggErr)
        case "SerialNumber":
            tape.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &tape.State, aggErr)
        case "StorageDomainMemberId":
            tape.StorageDomainMemberId = parseNullableString(child.Content)
        case "TakeOwnershipPending":
            tape.TakeOwnershipPending = parseBool(child.Content, aggErr)
        case "TotalRawCapacity":
            tape.TotalRawCapacity = parseNullableInt64(child.Content, aggErr)
        case "Type":
            tape.Type = parseString(child.Content)
        case "VerifyPending":
            tape.VerifyPending = newPriorityFromContent(child.Content, aggErr)
        case "WriteProtected":
            tape.WriteProtected = parseBool(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Tape.", child.XMLName.Local)
        }
    }
}


func parseTapeSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []Tape {
    var result []Tape
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult Tape
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing Tape struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}

type TapeDensityDirective struct {
    Density TapeDriveType
    Id string
    PartitionId string
    TapeType string
}

func (tapeDensityDirective *TapeDensityDirective) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Density":
            parseEnum(child.Content, &tapeDensityDirective.Density, aggErr)
        case "Id":
            tapeDensityDirective.Id = parseString(child.Content)
        case "PartitionId":
            tapeDensityDirective.PartitionId = parseString(child.Content)
        case "TapeType":
            tapeDensityDirective.TapeType = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeDensityDirective.", child.XMLName.Local)
        }
    }
}

type TapeDrive struct {
    CleaningRequired bool
    ErrorMessage *string
    ForceTapeRemoval bool
    Id string
    LastCleaned *string
    MfgSerialNumber *string
    MinimumTaskPriority *Priority
    PartitionId string
    Quiesced Quiesced
    ReservedTaskType ReservedTaskType
    SerialNumber *string
    State TapeDriveState
    TapeId *string
    Type TapeDriveType
}

func (tapeDrive *TapeDrive) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CleaningRequired":
            tapeDrive.CleaningRequired = parseBool(child.Content, aggErr)
        case "ErrorMessage":
            tapeDrive.ErrorMessage = parseNullableString(child.Content)
        case "ForceTapeRemoval":
            tapeDrive.ForceTapeRemoval = parseBool(child.Content, aggErr)
        case "Id":
            tapeDrive.Id = parseString(child.Content)
        case "LastCleaned":
            tapeDrive.LastCleaned = parseNullableString(child.Content)
        case "MfgSerialNumber":
            tapeDrive.MfgSerialNumber = parseNullableString(child.Content)
        case "MinimumTaskPriority":
            tapeDrive.MinimumTaskPriority = newPriorityFromContent(child.Content, aggErr)
        case "PartitionId":
            tapeDrive.PartitionId = parseString(child.Content)
        case "Quiesced":
            parseEnum(child.Content, &tapeDrive.Quiesced, aggErr)
        case "ReservedTaskType":
            parseEnum(child.Content, &tapeDrive.ReservedTaskType, aggErr)
        case "SerialNumber":
            tapeDrive.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &tapeDrive.State, aggErr)
        case "TapeId":
            tapeDrive.TapeId = parseNullableString(child.Content)
        case "Type":
            parseEnum(child.Content, &tapeDrive.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeDrive.", child.XMLName.Local)
        }
    }
}

type TapeDriveState Enum

const (
    TAPE_DRIVE_STATE_OFFLINE TapeDriveState = 1 + iota
    TAPE_DRIVE_STATE_NORMAL TapeDriveState = 1 + iota
    TAPE_DRIVE_STATE_ERROR TapeDriveState = 1 + iota
    TAPE_DRIVE_STATE_NOT_COMPATIBLE_IN_PARTITION_DUE_TO_NEWER_TAPE_DRIVES TapeDriveState = 1 + iota
)

func (tapeDriveState *TapeDriveState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapeDriveState = UNDEFINED
        case "OFFLINE": *tapeDriveState = TAPE_DRIVE_STATE_OFFLINE
        case "NORMAL": *tapeDriveState = TAPE_DRIVE_STATE_NORMAL
        case "ERROR": *tapeDriveState = TAPE_DRIVE_STATE_ERROR
        case "NOT_COMPATIBLE_IN_PARTITION_DUE_TO_NEWER_TAPE_DRIVES": *tapeDriveState = TAPE_DRIVE_STATE_NOT_COMPATIBLE_IN_PARTITION_DUE_TO_NEWER_TAPE_DRIVES
        default:
            *tapeDriveState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapeDriveState", str))
    }
    return nil
}

func (tapeDriveState TapeDriveState) String() string {
    switch tapeDriveState {
        case TAPE_DRIVE_STATE_OFFLINE: return "OFFLINE"
        case TAPE_DRIVE_STATE_NORMAL: return "NORMAL"
        case TAPE_DRIVE_STATE_ERROR: return "ERROR"
        case TAPE_DRIVE_STATE_NOT_COMPATIBLE_IN_PARTITION_DUE_TO_NEWER_TAPE_DRIVES: return "NOT_COMPATIBLE_IN_PARTITION_DUE_TO_NEWER_TAPE_DRIVES"
        default:
            log.Printf("Error: invalid TapeDriveState represented by '%d'", tapeDriveState)
            return ""
    }
}

func (tapeDriveState TapeDriveState) StringPtr() *string {
    if tapeDriveState == UNDEFINED {
        return nil
    }
    result := tapeDriveState.String()
    return &result
}

func newTapeDriveStateFromContent(content []byte, aggErr *AggregateError) *TapeDriveState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(TapeDriveState)
    parseEnum(content, result, aggErr)
    return result
}
type TapeDriveType Enum

const (
    TAPE_DRIVE_TYPE_UNKNOWN TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_LTO5 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_LTO6 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_LTO7 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_LTO8 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_LTO9 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_TS1140 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_TS1150 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_TS1155 TapeDriveType = 1 + iota
    TAPE_DRIVE_TYPE_TS1160 TapeDriveType = 1 + iota
)

func (tapeDriveType *TapeDriveType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapeDriveType = UNDEFINED
        case "UNKNOWN": *tapeDriveType = TAPE_DRIVE_TYPE_UNKNOWN
        case "LTO5": *tapeDriveType = TAPE_DRIVE_TYPE_LTO5
        case "LTO6": *tapeDriveType = TAPE_DRIVE_TYPE_LTO6
        case "LTO7": *tapeDriveType = TAPE_DRIVE_TYPE_LTO7
        case "LTO8": *tapeDriveType = TAPE_DRIVE_TYPE_LTO8
        case "LTO9": *tapeDriveType = TAPE_DRIVE_TYPE_LTO9
        case "TS1140": *tapeDriveType = TAPE_DRIVE_TYPE_TS1140
        case "TS1150": *tapeDriveType = TAPE_DRIVE_TYPE_TS1150
        case "TS1155": *tapeDriveType = TAPE_DRIVE_TYPE_TS1155
        case "TS1160": *tapeDriveType = TAPE_DRIVE_TYPE_TS1160
        default:
            *tapeDriveType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapeDriveType", str))
    }
    return nil
}

func (tapeDriveType TapeDriveType) String() string {
    switch tapeDriveType {
        case TAPE_DRIVE_TYPE_UNKNOWN: return "UNKNOWN"
        case TAPE_DRIVE_TYPE_LTO5: return "LTO5"
        case TAPE_DRIVE_TYPE_LTO6: return "LTO6"
        case TAPE_DRIVE_TYPE_LTO7: return "LTO7"
        case TAPE_DRIVE_TYPE_LTO8: return "LTO8"
        case TAPE_DRIVE_TYPE_LTO9: return "LTO9"
        case TAPE_DRIVE_TYPE_TS1140: return "TS1140"
        case TAPE_DRIVE_TYPE_TS1150: return "TS1150"
        case TAPE_DRIVE_TYPE_TS1155: return "TS1155"
        case TAPE_DRIVE_TYPE_TS1160: return "TS1160"
        default:
            log.Printf("Error: invalid TapeDriveType represented by '%d'", tapeDriveType)
            return ""
    }
}

func (tapeDriveType TapeDriveType) StringPtr() *string {
    if tapeDriveType == UNDEFINED {
        return nil
    }
    result := tapeDriveType.String()
    return &result
}

func newTapeDriveTypeFromContent(content []byte, aggErr *AggregateError) *TapeDriveType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(TapeDriveType)
    parseEnum(content, result, aggErr)
    return result
}
type DetailedTapeFailure struct {
    Date string
    ErrorMessage *string
    Id string
    TapeDriveId string
    TapeId string
    Type TapeFailureType
}

func (detailedTapeFailure *DetailedTapeFailure) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Date":
            detailedTapeFailure.Date = parseString(child.Content)
        case "ErrorMessage":
            detailedTapeFailure.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            detailedTapeFailure.Id = parseString(child.Content)
        case "TapeDriveId":
            detailedTapeFailure.TapeDriveId = parseString(child.Content)
        case "TapeId":
            detailedTapeFailure.TapeId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &detailedTapeFailure.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DetailedTapeFailure.", child.XMLName.Local)
        }
    }
}

type TapeFailureType Enum

const (
    TAPE_FAILURE_TYPE_BAR_CODE_CHANGED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_BAR_CODE_DUPLICATE TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_BLOB_READ_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DATA_CHECKPOINT_MISSING TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DELAYED_OWNERSHIP_FAILURE TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DRIVE_CLEAN_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_DRIVE_CLEANED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_FORMAT_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_GET_TAPE_INFORMATION_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_IMPORT_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_IMPORT_INCOMPLETE TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_INSPECT_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_READ_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_REIMPORT_REQUIRED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_SERIAL_NUMBER_MISMATCH TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_VERIFY_FAILED TapeFailureType = 1 + iota
    TAPE_FAILURE_TYPE_WRITE_FAILED TapeFailureType = 1 + iota
)

func (tapeFailureType *TapeFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapeFailureType = UNDEFINED
        case "BAR_CODE_CHANGED": *tapeFailureType = TAPE_FAILURE_TYPE_BAR_CODE_CHANGED
        case "BAR_CODE_DUPLICATE": *tapeFailureType = TAPE_FAILURE_TYPE_BAR_CODE_DUPLICATE
        case "BLOB_READ_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_BLOB_READ_FAILED
        case "DATA_CHECKPOINT_FAILURE": *tapeFailureType = TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE
        case "DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY": *tapeFailureType = TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY
        case "DATA_CHECKPOINT_MISSING": *tapeFailureType = TAPE_FAILURE_TYPE_DATA_CHECKPOINT_MISSING
        case "DELAYED_OWNERSHIP_FAILURE": *tapeFailureType = TAPE_FAILURE_TYPE_DELAYED_OWNERSHIP_FAILURE
        case "DRIVE_CLEAN_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_DRIVE_CLEAN_FAILED
        case "DRIVE_CLEANED": *tapeFailureType = TAPE_FAILURE_TYPE_DRIVE_CLEANED
        case "FORMAT_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_FORMAT_FAILED
        case "GET_TAPE_INFORMATION_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_GET_TAPE_INFORMATION_FAILED
        case "IMPORT_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_IMPORT_FAILED
        case "IMPORT_INCOMPLETE": *tapeFailureType = TAPE_FAILURE_TYPE_IMPORT_INCOMPLETE
        case "IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE": *tapeFailureType = TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE
        case "IMPORT_FAILED_DUE_TO_DATA_INTEGRITY": *tapeFailureType = TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY
        case "INSPECT_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_INSPECT_FAILED
        case "READ_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_READ_FAILED
        case "REIMPORT_REQUIRED": *tapeFailureType = TAPE_FAILURE_TYPE_REIMPORT_REQUIRED
        case "SERIAL_NUMBER_MISMATCH": *tapeFailureType = TAPE_FAILURE_TYPE_SERIAL_NUMBER_MISMATCH
        case "VERIFY_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_VERIFY_FAILED
        case "WRITE_FAILED": *tapeFailureType = TAPE_FAILURE_TYPE_WRITE_FAILED
        default:
            *tapeFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapeFailureType", str))
    }
    return nil
}

func (tapeFailureType TapeFailureType) String() string {
    switch tapeFailureType {
        case TAPE_FAILURE_TYPE_BAR_CODE_CHANGED: return "BAR_CODE_CHANGED"
        case TAPE_FAILURE_TYPE_BAR_CODE_DUPLICATE: return "BAR_CODE_DUPLICATE"
        case TAPE_FAILURE_TYPE_BLOB_READ_FAILED: return "BLOB_READ_FAILED"
        case TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE: return "DATA_CHECKPOINT_FAILURE"
        case TAPE_FAILURE_TYPE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY: return "DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY"
        case TAPE_FAILURE_TYPE_DATA_CHECKPOINT_MISSING: return "DATA_CHECKPOINT_MISSING"
        case TAPE_FAILURE_TYPE_DELAYED_OWNERSHIP_FAILURE: return "DELAYED_OWNERSHIP_FAILURE"
        case TAPE_FAILURE_TYPE_DRIVE_CLEAN_FAILED: return "DRIVE_CLEAN_FAILED"
        case TAPE_FAILURE_TYPE_DRIVE_CLEANED: return "DRIVE_CLEANED"
        case TAPE_FAILURE_TYPE_FORMAT_FAILED: return "FORMAT_FAILED"
        case TAPE_FAILURE_TYPE_GET_TAPE_INFORMATION_FAILED: return "GET_TAPE_INFORMATION_FAILED"
        case TAPE_FAILURE_TYPE_IMPORT_FAILED: return "IMPORT_FAILED"
        case TAPE_FAILURE_TYPE_IMPORT_INCOMPLETE: return "IMPORT_INCOMPLETE"
        case TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE: return "IMPORT_FAILED_DUE_TO_TAKE_OWNERSHIP_FAILURE"
        case TAPE_FAILURE_TYPE_IMPORT_FAILED_DUE_TO_DATA_INTEGRITY: return "IMPORT_FAILED_DUE_TO_DATA_INTEGRITY"
        case TAPE_FAILURE_TYPE_INSPECT_FAILED: return "INSPECT_FAILED"
        case TAPE_FAILURE_TYPE_READ_FAILED: return "READ_FAILED"
        case TAPE_FAILURE_TYPE_REIMPORT_REQUIRED: return "REIMPORT_REQUIRED"
        case TAPE_FAILURE_TYPE_SERIAL_NUMBER_MISMATCH: return "SERIAL_NUMBER_MISMATCH"
        case TAPE_FAILURE_TYPE_VERIFY_FAILED: return "VERIFY_FAILED"
        case TAPE_FAILURE_TYPE_WRITE_FAILED: return "WRITE_FAILED"
        default:
            log.Printf("Error: invalid TapeFailureType represented by '%d'", tapeFailureType)
            return ""
    }
}

func (tapeFailureType TapeFailureType) StringPtr() *string {
    if tapeFailureType == UNDEFINED {
        return nil
    }
    result := tapeFailureType.String()
    return &result
}

func newTapeFailureTypeFromContent(content []byte, aggErr *AggregateError) *TapeFailureType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(TapeFailureType)
    parseEnum(content, result, aggErr)
    return result
}
type TapeLibrary struct {
    Id string
    ManagementUrl *string
    Name *string
    SerialNumber *string
}

func (tapeLibrary *TapeLibrary) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Id":
            tapeLibrary.Id = parseString(child.Content)
        case "ManagementUrl":
            tapeLibrary.ManagementUrl = parseNullableString(child.Content)
        case "Name":
            tapeLibrary.Name = parseNullableString(child.Content)
        case "SerialNumber":
            tapeLibrary.SerialNumber = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeLibrary.", child.XMLName.Local)
        }
    }
}

type TapePartition struct {
    AutoCompactionEnabled bool
    DriveType *TapeDriveType
    ErrorMessage *string
    Id string
    ImportExportConfiguration ImportExportConfiguration
    LibraryId string
    MinimumReadReservedDrives int
    MinimumWriteReservedDrives int
    Name *string
    Quiesced Quiesced
    SerialNumber *string
    State TapePartitionState
}

func (tapePartition *TapePartition) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AutoCompactionEnabled":
            tapePartition.AutoCompactionEnabled = parseBool(child.Content, aggErr)
        case "DriveType":
            tapePartition.DriveType = newTapeDriveTypeFromContent(child.Content, aggErr)
        case "ErrorMessage":
            tapePartition.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            tapePartition.Id = parseString(child.Content)
        case "ImportExportConfiguration":
            parseEnum(child.Content, &tapePartition.ImportExportConfiguration, aggErr)
        case "LibraryId":
            tapePartition.LibraryId = parseString(child.Content)
        case "MinimumReadReservedDrives":
            tapePartition.MinimumReadReservedDrives = parseInt(child.Content, aggErr)
        case "MinimumWriteReservedDrives":
            tapePartition.MinimumWriteReservedDrives = parseInt(child.Content, aggErr)
        case "Name":
            tapePartition.Name = parseNullableString(child.Content)
        case "Quiesced":
            parseEnum(child.Content, &tapePartition.Quiesced, aggErr)
        case "SerialNumber":
            tapePartition.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &tapePartition.State, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapePartition.", child.XMLName.Local)
        }
    }
}

type TapePartitionFailure struct {
    Date string
    ErrorMessage *string
    Id string
    PartitionId string
    Type TapePartitionFailureType
}

func (tapePartitionFailure *TapePartitionFailure) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Date":
            tapePartitionFailure.Date = parseString(child.Content)
        case "ErrorMessage":
            tapePartitionFailure.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            tapePartitionFailure.Id = parseString(child.Content)
        case "PartitionId":
            tapePartitionFailure.PartitionId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &tapePartitionFailure.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapePartitionFailure.", child.XMLName.Local)
        }
    }
}

type TapePartitionFailureType Enum

const (
    TAPE_PARTITION_FAILURE_TYPE_CLEANING_TAPE_REQUIRED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_DUPLICATE_TAPE_BAR_CODES_DETECTED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_EJECT_STALLED_DUE_TO_OFFLINE_TAPES TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_MINIMUM_DRIVE_COUNT_NOT_MET TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED_DUE_TO_PREPARE_TAPE_FOR_REMOVAL_FAILURE TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_NO_USABLE_DRIVES TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_ONLINE_STALLED_DUE_TO_NO_STORAGE_SLOTS TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_IN_ERROR TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_MISSING TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_QUIESCED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_TYPE_MISMATCH TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_EJECTION_BY_OPERATOR_REQUIRED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_MEDIA_TYPE_INCOMPATIBLE TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_REMOVAL_UNEXPECTED TapePartitionFailureType = 1 + iota
    TAPE_PARTITION_FAILURE_TYPE_TAPE_IN_INVALID_PARTITION TapePartitionFailureType = 1 + iota
)

func (tapePartitionFailureType *TapePartitionFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapePartitionFailureType = UNDEFINED
        case "CLEANING_TAPE_REQUIRED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_CLEANING_TAPE_REQUIRED
        case "DUPLICATE_TAPE_BAR_CODES_DETECTED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_DUPLICATE_TAPE_BAR_CODES_DETECTED
        case "EJECT_STALLED_DUE_TO_OFFLINE_TAPES": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_EJECT_STALLED_DUE_TO_OFFLINE_TAPES
        case "MINIMUM_DRIVE_COUNT_NOT_MET": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_MINIMUM_DRIVE_COUNT_NOT_MET
        case "MOVE_FAILED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED
        case "MOVE_FAILED_DUE_TO_PREPARE_TAPE_FOR_REMOVAL_FAILURE": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED_DUE_TO_PREPARE_TAPE_FOR_REMOVAL_FAILURE
        case "NO_USABLE_DRIVES": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_NO_USABLE_DRIVES
        case "ONLINE_STALLED_DUE_TO_NO_STORAGE_SLOTS": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_ONLINE_STALLED_DUE_TO_NO_STORAGE_SLOTS
        case "TAPE_DRIVE_IN_ERROR": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_IN_ERROR
        case "TAPE_DRIVE_MISSING": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_MISSING
        case "TAPE_DRIVE_QUIESCED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_QUIESCED
        case "TAPE_DRIVE_TYPE_MISMATCH": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_TYPE_MISMATCH
        case "TAPE_EJECTION_BY_OPERATOR_REQUIRED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_EJECTION_BY_OPERATOR_REQUIRED
        case "TAPE_MEDIA_TYPE_INCOMPATIBLE": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_MEDIA_TYPE_INCOMPATIBLE
        case "TAPE_REMOVAL_UNEXPECTED": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_REMOVAL_UNEXPECTED
        case "TAPE_IN_INVALID_PARTITION": *tapePartitionFailureType = TAPE_PARTITION_FAILURE_TYPE_TAPE_IN_INVALID_PARTITION
        default:
            *tapePartitionFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapePartitionFailureType", str))
    }
    return nil
}

func (tapePartitionFailureType TapePartitionFailureType) String() string {
    switch tapePartitionFailureType {
        case TAPE_PARTITION_FAILURE_TYPE_CLEANING_TAPE_REQUIRED: return "CLEANING_TAPE_REQUIRED"
        case TAPE_PARTITION_FAILURE_TYPE_DUPLICATE_TAPE_BAR_CODES_DETECTED: return "DUPLICATE_TAPE_BAR_CODES_DETECTED"
        case TAPE_PARTITION_FAILURE_TYPE_EJECT_STALLED_DUE_TO_OFFLINE_TAPES: return "EJECT_STALLED_DUE_TO_OFFLINE_TAPES"
        case TAPE_PARTITION_FAILURE_TYPE_MINIMUM_DRIVE_COUNT_NOT_MET: return "MINIMUM_DRIVE_COUNT_NOT_MET"
        case TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED: return "MOVE_FAILED"
        case TAPE_PARTITION_FAILURE_TYPE_MOVE_FAILED_DUE_TO_PREPARE_TAPE_FOR_REMOVAL_FAILURE: return "MOVE_FAILED_DUE_TO_PREPARE_TAPE_FOR_REMOVAL_FAILURE"
        case TAPE_PARTITION_FAILURE_TYPE_NO_USABLE_DRIVES: return "NO_USABLE_DRIVES"
        case TAPE_PARTITION_FAILURE_TYPE_ONLINE_STALLED_DUE_TO_NO_STORAGE_SLOTS: return "ONLINE_STALLED_DUE_TO_NO_STORAGE_SLOTS"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_IN_ERROR: return "TAPE_DRIVE_IN_ERROR"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_MISSING: return "TAPE_DRIVE_MISSING"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_QUIESCED: return "TAPE_DRIVE_QUIESCED"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_DRIVE_TYPE_MISMATCH: return "TAPE_DRIVE_TYPE_MISMATCH"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_EJECTION_BY_OPERATOR_REQUIRED: return "TAPE_EJECTION_BY_OPERATOR_REQUIRED"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_MEDIA_TYPE_INCOMPATIBLE: return "TAPE_MEDIA_TYPE_INCOMPATIBLE"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_REMOVAL_UNEXPECTED: return "TAPE_REMOVAL_UNEXPECTED"
        case TAPE_PARTITION_FAILURE_TYPE_TAPE_IN_INVALID_PARTITION: return "TAPE_IN_INVALID_PARTITION"
        default:
            log.Printf("Error: invalid TapePartitionFailureType represented by '%d'", tapePartitionFailureType)
            return ""
    }
}

func (tapePartitionFailureType TapePartitionFailureType) StringPtr() *string {
    if tapePartitionFailureType == UNDEFINED {
        return nil
    }
    result := tapePartitionFailureType.String()
    return &result
}

func newTapePartitionFailureTypeFromContent(content []byte, aggErr *AggregateError) *TapePartitionFailureType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(TapePartitionFailureType)
    parseEnum(content, result, aggErr)
    return result
}
type TapePartitionState Enum

const (
    TAPE_PARTITION_STATE_ONLINE TapePartitionState = 1 + iota
    TAPE_PARTITION_STATE_OFFLINE TapePartitionState = 1 + iota
    TAPE_PARTITION_STATE_ERROR TapePartitionState = 1 + iota
)

func (tapePartitionState *TapePartitionState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapePartitionState = UNDEFINED
        case "ONLINE": *tapePartitionState = TAPE_PARTITION_STATE_ONLINE
        case "OFFLINE": *tapePartitionState = TAPE_PARTITION_STATE_OFFLINE
        case "ERROR": *tapePartitionState = TAPE_PARTITION_STATE_ERROR
        default:
            *tapePartitionState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapePartitionState", str))
    }
    return nil
}

func (tapePartitionState TapePartitionState) String() string {
    switch tapePartitionState {
        case TAPE_PARTITION_STATE_ONLINE: return "ONLINE"
        case TAPE_PARTITION_STATE_OFFLINE: return "OFFLINE"
        case TAPE_PARTITION_STATE_ERROR: return "ERROR"
        default:
            log.Printf("Error: invalid TapePartitionState represented by '%d'", tapePartitionState)
            return ""
    }
}

func (tapePartitionState TapePartitionState) StringPtr() *string {
    if tapePartitionState == UNDEFINED {
        return nil
    }
    result := tapePartitionState.String()
    return &result
}

func newTapePartitionStateFromContent(content []byte, aggErr *AggregateError) *TapePartitionState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(TapePartitionState)
    parseEnum(content, result, aggErr)
    return result
}
type TapeState Enum

const (
    TAPE_STATE_NORMAL TapeState = 1 + iota
    TAPE_STATE_OFFLINE TapeState = 1 + iota
    TAPE_STATE_ONLINE_PENDING TapeState = 1 + iota
    TAPE_STATE_ONLINE_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_PENDING_INSPECTION TapeState = 1 + iota
    TAPE_STATE_UNKNOWN TapeState = 1 + iota
    TAPE_STATE_DATA_CHECKPOINT_FAILURE TapeState = 1 + iota
    TAPE_STATE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY TapeState = 1 + iota
    TAPE_STATE_DATA_CHECKPOINT_MISSING TapeState = 1 + iota
    TAPE_STATE_LTFS_WITH_FOREIGN_DATA TapeState = 1 + iota
    TAPE_STATE_RAW_IMPORT_PENDING TapeState = 1 + iota
    TAPE_STATE_RAW_IMPORT_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_FOREIGN TapeState = 1 + iota
    TAPE_STATE_IMPORT_PENDING TapeState = 1 + iota
    TAPE_STATE_IMPORT_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_INCOMPATIBLE TapeState = 1 + iota
    TAPE_STATE_LOST TapeState = 1 + iota
    TAPE_STATE_BAD TapeState = 1 + iota
    TAPE_STATE_CANNOT_FORMAT_DUE_TO_WRITE_PROTECTION TapeState = 1 + iota
    TAPE_STATE_SERIAL_NUMBER_MISMATCH TapeState = 1 + iota
    TAPE_STATE_BAR_CODE_MISSING TapeState = 1 + iota
    TAPE_STATE_AUTO_COMPACTION_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_FORMAT_PENDING TapeState = 1 + iota
    TAPE_STATE_FORMAT_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_EJECT_TO_EE_IN_PROGRESS TapeState = 1 + iota
    TAPE_STATE_EJECT_FROM_EE_PENDING TapeState = 1 + iota
    TAPE_STATE_EJECTED TapeState = 1 + iota
)

func (tapeState *TapeState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *tapeState = UNDEFINED
        case "NORMAL": *tapeState = TAPE_STATE_NORMAL
        case "OFFLINE": *tapeState = TAPE_STATE_OFFLINE
        case "ONLINE_PENDING": *tapeState = TAPE_STATE_ONLINE_PENDING
        case "ONLINE_IN_PROGRESS": *tapeState = TAPE_STATE_ONLINE_IN_PROGRESS
        case "PENDING_INSPECTION": *tapeState = TAPE_STATE_PENDING_INSPECTION
        case "UNKNOWN": *tapeState = TAPE_STATE_UNKNOWN
        case "DATA_CHECKPOINT_FAILURE": *tapeState = TAPE_STATE_DATA_CHECKPOINT_FAILURE
        case "DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY": *tapeState = TAPE_STATE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY
        case "DATA_CHECKPOINT_MISSING": *tapeState = TAPE_STATE_DATA_CHECKPOINT_MISSING
        case "LTFS_WITH_FOREIGN_DATA": *tapeState = TAPE_STATE_LTFS_WITH_FOREIGN_DATA
        case "RAW_IMPORT_PENDING": *tapeState = TAPE_STATE_RAW_IMPORT_PENDING
        case "RAW_IMPORT_IN_PROGRESS": *tapeState = TAPE_STATE_RAW_IMPORT_IN_PROGRESS
        case "FOREIGN": *tapeState = TAPE_STATE_FOREIGN
        case "IMPORT_PENDING": *tapeState = TAPE_STATE_IMPORT_PENDING
        case "IMPORT_IN_PROGRESS": *tapeState = TAPE_STATE_IMPORT_IN_PROGRESS
        case "INCOMPATIBLE": *tapeState = TAPE_STATE_INCOMPATIBLE
        case "LOST": *tapeState = TAPE_STATE_LOST
        case "BAD": *tapeState = TAPE_STATE_BAD
        case "CANNOT_FORMAT_DUE_TO_WRITE_PROTECTION": *tapeState = TAPE_STATE_CANNOT_FORMAT_DUE_TO_WRITE_PROTECTION
        case "SERIAL_NUMBER_MISMATCH": *tapeState = TAPE_STATE_SERIAL_NUMBER_MISMATCH
        case "BAR_CODE_MISSING": *tapeState = TAPE_STATE_BAR_CODE_MISSING
        case "AUTO_COMPACTION_IN_PROGRESS": *tapeState = TAPE_STATE_AUTO_COMPACTION_IN_PROGRESS
        case "FORMAT_PENDING": *tapeState = TAPE_STATE_FORMAT_PENDING
        case "FORMAT_IN_PROGRESS": *tapeState = TAPE_STATE_FORMAT_IN_PROGRESS
        case "EJECT_TO_EE_IN_PROGRESS": *tapeState = TAPE_STATE_EJECT_TO_EE_IN_PROGRESS
        case "EJECT_FROM_EE_PENDING": *tapeState = TAPE_STATE_EJECT_FROM_EE_PENDING
        case "EJECTED": *tapeState = TAPE_STATE_EJECTED
        default:
            *tapeState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TapeState", str))
    }
    return nil
}

func (tapeState TapeState) String() string {
    switch tapeState {
        case TAPE_STATE_NORMAL: return "NORMAL"
        case TAPE_STATE_OFFLINE: return "OFFLINE"
        case TAPE_STATE_ONLINE_PENDING: return "ONLINE_PENDING"
        case TAPE_STATE_ONLINE_IN_PROGRESS: return "ONLINE_IN_PROGRESS"
        case TAPE_STATE_PENDING_INSPECTION: return "PENDING_INSPECTION"
        case TAPE_STATE_UNKNOWN: return "UNKNOWN"
        case TAPE_STATE_DATA_CHECKPOINT_FAILURE: return "DATA_CHECKPOINT_FAILURE"
        case TAPE_STATE_DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY: return "DATA_CHECKPOINT_FAILURE_DUE_TO_READ_ONLY"
        case TAPE_STATE_DATA_CHECKPOINT_MISSING: return "DATA_CHECKPOINT_MISSING"
        case TAPE_STATE_LTFS_WITH_FOREIGN_DATA: return "LTFS_WITH_FOREIGN_DATA"
        case TAPE_STATE_RAW_IMPORT_PENDING: return "RAW_IMPORT_PENDING"
        case TAPE_STATE_RAW_IMPORT_IN_PROGRESS: return "RAW_IMPORT_IN_PROGRESS"
        case TAPE_STATE_FOREIGN: return "FOREIGN"
        case TAPE_STATE_IMPORT_PENDING: return "IMPORT_PENDING"
        case TAPE_STATE_IMPORT_IN_PROGRESS: return "IMPORT_IN_PROGRESS"
        case TAPE_STATE_INCOMPATIBLE: return "INCOMPATIBLE"
        case TAPE_STATE_LOST: return "LOST"
        case TAPE_STATE_BAD: return "BAD"
        case TAPE_STATE_CANNOT_FORMAT_DUE_TO_WRITE_PROTECTION: return "CANNOT_FORMAT_DUE_TO_WRITE_PROTECTION"
        case TAPE_STATE_SERIAL_NUMBER_MISMATCH: return "SERIAL_NUMBER_MISMATCH"
        case TAPE_STATE_BAR_CODE_MISSING: return "BAR_CODE_MISSING"
        case TAPE_STATE_AUTO_COMPACTION_IN_PROGRESS: return "AUTO_COMPACTION_IN_PROGRESS"
        case TAPE_STATE_FORMAT_PENDING: return "FORMAT_PENDING"
        case TAPE_STATE_FORMAT_IN_PROGRESS: return "FORMAT_IN_PROGRESS"
        case TAPE_STATE_EJECT_TO_EE_IN_PROGRESS: return "EJECT_TO_EE_IN_PROGRESS"
        case TAPE_STATE_EJECT_FROM_EE_PENDING: return "EJECT_FROM_EE_PENDING"
        case TAPE_STATE_EJECTED: return "EJECTED"
        default:
            log.Printf("Error: invalid TapeState represented by '%d'", tapeState)
            return ""
    }
}

func (tapeState TapeState) StringPtr() *string {
    if tapeState == UNDEFINED {
        return nil
    }
    result := tapeState.String()
    return &result
}

func newTapeStateFromContent(content []byte, aggErr *AggregateError) *TapeState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(TapeState)
    parseEnum(content, result, aggErr)
    return result
}
type AzureTarget struct {
    AccountKey *string
    AccountName *string
    AutoVerifyFrequencyInDays *int
    CloudBucketPrefix *string
    CloudBucketSuffix *string
    DefaultReadPreference TargetReadPreferenceType
    Https bool
    Id string
    LastFullyVerified *string
    Name *string
    NamingMode CloudNamingMode
    PermitGoingOutOfSync bool
    Quiesced Quiesced
    State TargetState
}

func (azureTarget *AzureTarget) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AccountKey":
            azureTarget.AccountKey = parseNullableString(child.Content)
        case "AccountName":
            azureTarget.AccountName = parseNullableString(child.Content)
        case "AutoVerifyFrequencyInDays":
            azureTarget.AutoVerifyFrequencyInDays = parseNullableInt(child.Content, aggErr)
        case "CloudBucketPrefix":
            azureTarget.CloudBucketPrefix = parseNullableString(child.Content)
        case "CloudBucketSuffix":
            azureTarget.CloudBucketSuffix = parseNullableString(child.Content)
        case "DefaultReadPreference":
            parseEnum(child.Content, &azureTarget.DefaultReadPreference, aggErr)
        case "Https":
            azureTarget.Https = parseBool(child.Content, aggErr)
        case "Id":
            azureTarget.Id = parseString(child.Content)
        case "LastFullyVerified":
            azureTarget.LastFullyVerified = parseNullableString(child.Content)
        case "Name":
            azureTarget.Name = parseNullableString(child.Content)
        case "NamingMode":
            parseEnum(child.Content, &azureTarget.NamingMode, aggErr)
        case "PermitGoingOutOfSync":
            azureTarget.PermitGoingOutOfSync = parseBool(child.Content, aggErr)
        case "Quiesced":
            parseEnum(child.Content, &azureTarget.Quiesced, aggErr)
        case "State":
            parseEnum(child.Content, &azureTarget.State, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureTarget.", child.XMLName.Local)
        }
    }
}


func parseAzureTargetSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []AzureTarget {
    var result []AzureTarget
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult AzureTarget
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing AzureTarget struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}

type AzureTargetBucketName struct {
    BucketId string
    Id string
    Name *string
    TargetId string
}

func (azureTargetBucketName *AzureTargetBucketName) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            azureTargetBucketName.BucketId = parseString(child.Content)
        case "Id":
            azureTargetBucketName.Id = parseString(child.Content)
        case "Name":
            azureTargetBucketName.Name = parseNullableString(child.Content)
        case "TargetId":
            azureTargetBucketName.TargetId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureTargetBucketName.", child.XMLName.Local)
        }
    }
}

type AzureTargetFailure struct {
    Date string
    ErrorMessage *string
    Id string
    TargetId string
    Type TargetFailureType
}

func (azureTargetFailure *AzureTargetFailure) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Date":
            azureTargetFailure.Date = parseString(child.Content)
        case "ErrorMessage":
            azureTargetFailure.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            azureTargetFailure.Id = parseString(child.Content)
        case "TargetId":
            azureTargetFailure.TargetId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &azureTargetFailure.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureTargetFailure.", child.XMLName.Local)
        }
    }
}

type AzureTargetReadPreference struct {
    BucketId string
    Id string
    ReadPreference TargetReadPreferenceType
    TargetId string
}

func (azureTargetReadPreference *AzureTargetReadPreference) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            azureTargetReadPreference.BucketId = parseString(child.Content)
        case "Id":
            azureTargetReadPreference.Id = parseString(child.Content)
        case "ReadPreference":
            parseEnum(child.Content, &azureTargetReadPreference.ReadPreference, aggErr)
        case "TargetId":
            azureTargetReadPreference.TargetId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureTargetReadPreference.", child.XMLName.Local)
        }
    }
}

type Ds3Target struct {
    AccessControlReplication Ds3TargetAccessControlReplication
    AdminAuthId *string
    AdminSecretKey *string
    DataPathEndPoint *string
    DataPathHttps bool
    DataPathPort *int
    DataPathProxy *string
    DataPathVerifyCertificate bool
    DefaultReadPreference TargetReadPreferenceType
    Id string
    Name *string
    PermitGoingOutOfSync bool
    Quiesced Quiesced
    ReplicatedUserDefaultDataPolicy *string
    State TargetState
}

func (ds3Target *Ds3Target) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AccessControlReplication":
            parseEnum(child.Content, &ds3Target.AccessControlReplication, aggErr)
        case "AdminAuthId":
            ds3Target.AdminAuthId = parseNullableString(child.Content)
        case "AdminSecretKey":
            ds3Target.AdminSecretKey = parseNullableString(child.Content)
        case "DataPathEndPoint":
            ds3Target.DataPathEndPoint = parseNullableString(child.Content)
        case "DataPathHttps":
            ds3Target.DataPathHttps = parseBool(child.Content, aggErr)
        case "DataPathPort":
            ds3Target.DataPathPort = parseNullableInt(child.Content, aggErr)
        case "DataPathProxy":
            ds3Target.DataPathProxy = parseNullableString(child.Content)
        case "DataPathVerifyCertificate":
            ds3Target.DataPathVerifyCertificate = parseBool(child.Content, aggErr)
        case "DefaultReadPreference":
            parseEnum(child.Content, &ds3Target.DefaultReadPreference, aggErr)
        case "Id":
            ds3Target.Id = parseString(child.Content)
        case "Name":
            ds3Target.Name = parseNullableString(child.Content)
        case "PermitGoingOutOfSync":
            ds3Target.PermitGoingOutOfSync = parseBool(child.Content, aggErr)
        case "Quiesced":
            parseEnum(child.Content, &ds3Target.Quiesced, aggErr)
        case "ReplicatedUserDefaultDataPolicy":
            ds3Target.ReplicatedUserDefaultDataPolicy = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &ds3Target.State, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3Target.", child.XMLName.Local)
        }
    }
}


func parseDs3TargetSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []Ds3Target {
    var result []Ds3Target
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult Ds3Target
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing Ds3Target struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}

type Ds3TargetAccessControlReplication Enum

const (
    DS3_TARGET_ACCESS_CONTROL_REPLICATION_NONE Ds3TargetAccessControlReplication = 1 + iota
    DS3_TARGET_ACCESS_CONTROL_REPLICATION_USERS Ds3TargetAccessControlReplication = 1 + iota
)

func (ds3TargetAccessControlReplication *Ds3TargetAccessControlReplication) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *ds3TargetAccessControlReplication = UNDEFINED
        case "NONE": *ds3TargetAccessControlReplication = DS3_TARGET_ACCESS_CONTROL_REPLICATION_NONE
        case "USERS": *ds3TargetAccessControlReplication = DS3_TARGET_ACCESS_CONTROL_REPLICATION_USERS
        default:
            *ds3TargetAccessControlReplication = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into Ds3TargetAccessControlReplication", str))
    }
    return nil
}

func (ds3TargetAccessControlReplication Ds3TargetAccessControlReplication) String() string {
    switch ds3TargetAccessControlReplication {
        case DS3_TARGET_ACCESS_CONTROL_REPLICATION_NONE: return "NONE"
        case DS3_TARGET_ACCESS_CONTROL_REPLICATION_USERS: return "USERS"
        default:
            log.Printf("Error: invalid Ds3TargetAccessControlReplication represented by '%d'", ds3TargetAccessControlReplication)
            return ""
    }
}

func (ds3TargetAccessControlReplication Ds3TargetAccessControlReplication) StringPtr() *string {
    if ds3TargetAccessControlReplication == UNDEFINED {
        return nil
    }
    result := ds3TargetAccessControlReplication.String()
    return &result
}

func newDs3TargetAccessControlReplicationFromContent(content []byte, aggErr *AggregateError) *Ds3TargetAccessControlReplication {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(Ds3TargetAccessControlReplication)
    parseEnum(content, result, aggErr)
    return result
}
type Ds3TargetFailure struct {
    Date string
    ErrorMessage *string
    Id string
    TargetId string
    Type TargetFailureType
}

func (ds3TargetFailure *Ds3TargetFailure) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Date":
            ds3TargetFailure.Date = parseString(child.Content)
        case "ErrorMessage":
            ds3TargetFailure.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            ds3TargetFailure.Id = parseString(child.Content)
        case "TargetId":
            ds3TargetFailure.TargetId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &ds3TargetFailure.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3TargetFailure.", child.XMLName.Local)
        }
    }
}

type Ds3TargetReadPreference struct {
    BucketId string
    Id string
    ReadPreference TargetReadPreferenceType
    TargetId string
}

func (ds3TargetReadPreference *Ds3TargetReadPreference) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            ds3TargetReadPreference.BucketId = parseString(child.Content)
        case "Id":
            ds3TargetReadPreference.Id = parseString(child.Content)
        case "ReadPreference":
            parseEnum(child.Content, &ds3TargetReadPreference.ReadPreference, aggErr)
        case "TargetId":
            ds3TargetReadPreference.TargetId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3TargetReadPreference.", child.XMLName.Local)
        }
    }
}

type S3Target struct {
    AccessKey *string
    AutoVerifyFrequencyInDays *int
    CloudBucketPrefix *string
    CloudBucketSuffix *string
    DataPathEndPoint *string
    DefaultReadPreference TargetReadPreferenceType
    Https bool
    Id string
    LastFullyVerified *string
    Name *string
    NamingMode CloudNamingMode
    OfflineDataStagingWindowInTb int
    PermitGoingOutOfSync bool
    ProxyDomain *string
    ProxyHost *string
    ProxyPassword *string
    ProxyPort *int
    ProxyUsername *string
    Quiesced Quiesced
    Region *S3Region
    RestrictedAccess bool
    SecretKey *string
    StagedDataExpirationInDays int
    State TargetState
}

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
        case "NamingMode":
            parseEnum(child.Content, &s3Target.NamingMode, aggErr)
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
            s3Target.Region = newS3RegionFromContent(child.Content, aggErr)
        case "RestrictedAccess":
            s3Target.RestrictedAccess = parseBool(child.Content, aggErr)
        case "SecretKey":
            s3Target.SecretKey = parseNullableString(child.Content)
        case "StagedDataExpirationInDays":
            s3Target.StagedDataExpirationInDays = parseInt(child.Content, aggErr)
        case "State":
            parseEnum(child.Content, &s3Target.State, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3Target.", child.XMLName.Local)
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

type S3TargetBucketName struct {
    BucketId string
    Id string
    Name *string
    TargetId string
}

func (s3TargetBucketName *S3TargetBucketName) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            s3TargetBucketName.BucketId = parseString(child.Content)
        case "Id":
            s3TargetBucketName.Id = parseString(child.Content)
        case "Name":
            s3TargetBucketName.Name = parseNullableString(child.Content)
        case "TargetId":
            s3TargetBucketName.TargetId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3TargetBucketName.", child.XMLName.Local)
        }
    }
}

type S3TargetFailure struct {
    Date string
    ErrorMessage *string
    Id string
    TargetId string
    Type TargetFailureType
}

func (s3TargetFailure *S3TargetFailure) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Date":
            s3TargetFailure.Date = parseString(child.Content)
        case "ErrorMessage":
            s3TargetFailure.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            s3TargetFailure.Id = parseString(child.Content)
        case "TargetId":
            s3TargetFailure.TargetId = parseString(child.Content)
        case "Type":
            parseEnum(child.Content, &s3TargetFailure.Type, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3TargetFailure.", child.XMLName.Local)
        }
    }
}

type S3TargetReadPreference struct {
    BucketId string
    Id string
    ReadPreference TargetReadPreferenceType
    TargetId string
}

func (s3TargetReadPreference *S3TargetReadPreference) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketId":
            s3TargetReadPreference.BucketId = parseString(child.Content)
        case "Id":
            s3TargetReadPreference.Id = parseString(child.Content)
        case "ReadPreference":
            parseEnum(child.Content, &s3TargetReadPreference.ReadPreference, aggErr)
        case "TargetId":
            s3TargetReadPreference.TargetId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3TargetReadPreference.", child.XMLName.Local)
        }
    }
}

type SuspectBlobAzureTarget struct {
    BlobId string
    Id string
    TargetId string
}

func (suspectBlobAzureTarget *SuspectBlobAzureTarget) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BlobId":
            suspectBlobAzureTarget.BlobId = parseString(child.Content)
        case "Id":
            suspectBlobAzureTarget.Id = parseString(child.Content)
        case "TargetId":
            suspectBlobAzureTarget.TargetId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobAzureTarget.", child.XMLName.Local)
        }
    }
}

type SuspectBlobDs3Target struct {
    BlobId string
    Id string
    TargetId string
}

func (suspectBlobDs3Target *SuspectBlobDs3Target) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BlobId":
            suspectBlobDs3Target.BlobId = parseString(child.Content)
        case "Id":
            suspectBlobDs3Target.Id = parseString(child.Content)
        case "TargetId":
            suspectBlobDs3Target.TargetId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobDs3Target.", child.XMLName.Local)
        }
    }
}

type SuspectBlobS3Target struct {
    BlobId string
    Id string
    TargetId string
}

func (suspectBlobS3Target *SuspectBlobS3Target) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BlobId":
            suspectBlobS3Target.BlobId = parseString(child.Content)
        case "Id":
            suspectBlobS3Target.Id = parseString(child.Content)
        case "TargetId":
            suspectBlobS3Target.TargetId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobS3Target.", child.XMLName.Local)
        }
    }
}

type TargetFailureType Enum

const (
    TARGET_FAILURE_TYPE_IMPORT_FAILED TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_IMPORT_INCOMPLETE TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_NOT_ONLINE TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_WRITE_FAILED TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_WRITE_INITIATE_FAILED TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_READ_FAILED TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_READ_INITIATE_FAILED TargetFailureType = 1 + iota
    TARGET_FAILURE_TYPE_VERIFY_FAILED TargetFailureType = 1 + iota
)

func (targetFailureType *TargetFailureType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *targetFailureType = UNDEFINED
        case "IMPORT_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_IMPORT_FAILED
        case "IMPORT_INCOMPLETE": *targetFailureType = TARGET_FAILURE_TYPE_IMPORT_INCOMPLETE
        case "NOT_ONLINE": *targetFailureType = TARGET_FAILURE_TYPE_NOT_ONLINE
        case "WRITE_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_WRITE_FAILED
        case "WRITE_INITIATE_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_WRITE_INITIATE_FAILED
        case "READ_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_READ_FAILED
        case "READ_INITIATE_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_READ_INITIATE_FAILED
        case "VERIFY_FAILED": *targetFailureType = TARGET_FAILURE_TYPE_VERIFY_FAILED
        default:
            *targetFailureType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TargetFailureType", str))
    }
    return nil
}

func (targetFailureType TargetFailureType) String() string {
    switch targetFailureType {
        case TARGET_FAILURE_TYPE_IMPORT_FAILED: return "IMPORT_FAILED"
        case TARGET_FAILURE_TYPE_IMPORT_INCOMPLETE: return "IMPORT_INCOMPLETE"
        case TARGET_FAILURE_TYPE_NOT_ONLINE: return "NOT_ONLINE"
        case TARGET_FAILURE_TYPE_WRITE_FAILED: return "WRITE_FAILED"
        case TARGET_FAILURE_TYPE_WRITE_INITIATE_FAILED: return "WRITE_INITIATE_FAILED"
        case TARGET_FAILURE_TYPE_READ_FAILED: return "READ_FAILED"
        case TARGET_FAILURE_TYPE_READ_INITIATE_FAILED: return "READ_INITIATE_FAILED"
        case TARGET_FAILURE_TYPE_VERIFY_FAILED: return "VERIFY_FAILED"
        default:
            log.Printf("Error: invalid TargetFailureType represented by '%d'", targetFailureType)
            return ""
    }
}

func (targetFailureType TargetFailureType) StringPtr() *string {
    if targetFailureType == UNDEFINED {
        return nil
    }
    result := targetFailureType.String()
    return &result
}

func newTargetFailureTypeFromContent(content []byte, aggErr *AggregateError) *TargetFailureType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(TargetFailureType)
    parseEnum(content, result, aggErr)
    return result
}
type TargetReadPreferenceType Enum

const (
    TARGET_READ_PREFERENCE_TYPE_MINIMUM_LATENCY TargetReadPreferenceType = 1 + iota
    TARGET_READ_PREFERENCE_TYPE_AFTER_ONLINE_POOL TargetReadPreferenceType = 1 + iota
    TARGET_READ_PREFERENCE_TYPE_AFTER_NEARLINE_POOL TargetReadPreferenceType = 1 + iota
    TARGET_READ_PREFERENCE_TYPE_AFTER_NON_EJECTABLE_TAPE TargetReadPreferenceType = 1 + iota
    TARGET_READ_PREFERENCE_TYPE_LAST_RESORT TargetReadPreferenceType = 1 + iota
    TARGET_READ_PREFERENCE_TYPE_NEVER TargetReadPreferenceType = 1 + iota
)

func (targetReadPreferenceType *TargetReadPreferenceType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *targetReadPreferenceType = UNDEFINED
        case "MINIMUM_LATENCY": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_MINIMUM_LATENCY
        case "AFTER_ONLINE_POOL": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_AFTER_ONLINE_POOL
        case "AFTER_NEARLINE_POOL": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_AFTER_NEARLINE_POOL
        case "AFTER_NON_EJECTABLE_TAPE": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_AFTER_NON_EJECTABLE_TAPE
        case "LAST_RESORT": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_LAST_RESORT
        case "NEVER": *targetReadPreferenceType = TARGET_READ_PREFERENCE_TYPE_NEVER
        default:
            *targetReadPreferenceType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TargetReadPreferenceType", str))
    }
    return nil
}

func (targetReadPreferenceType TargetReadPreferenceType) String() string {
    switch targetReadPreferenceType {
        case TARGET_READ_PREFERENCE_TYPE_MINIMUM_LATENCY: return "MINIMUM_LATENCY"
        case TARGET_READ_PREFERENCE_TYPE_AFTER_ONLINE_POOL: return "AFTER_ONLINE_POOL"
        case TARGET_READ_PREFERENCE_TYPE_AFTER_NEARLINE_POOL: return "AFTER_NEARLINE_POOL"
        case TARGET_READ_PREFERENCE_TYPE_AFTER_NON_EJECTABLE_TAPE: return "AFTER_NON_EJECTABLE_TAPE"
        case TARGET_READ_PREFERENCE_TYPE_LAST_RESORT: return "LAST_RESORT"
        case TARGET_READ_PREFERENCE_TYPE_NEVER: return "NEVER"
        default:
            log.Printf("Error: invalid TargetReadPreferenceType represented by '%d'", targetReadPreferenceType)
            return ""
    }
}

func (targetReadPreferenceType TargetReadPreferenceType) StringPtr() *string {
    if targetReadPreferenceType == UNDEFINED {
        return nil
    }
    result := targetReadPreferenceType.String()
    return &result
}

func newTargetReadPreferenceTypeFromContent(content []byte, aggErr *AggregateError) *TargetReadPreferenceType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(TargetReadPreferenceType)
    parseEnum(content, result, aggErr)
    return result
}
type TargetState Enum

const (
    TARGET_STATE_ONLINE TargetState = 1 + iota
    TARGET_STATE_OFFLINE TargetState = 1 + iota
)

func (targetState *TargetState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *targetState = UNDEFINED
        case "ONLINE": *targetState = TARGET_STATE_ONLINE
        case "OFFLINE": *targetState = TARGET_STATE_OFFLINE
        default:
            *targetState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into TargetState", str))
    }
    return nil
}

func (targetState TargetState) String() string {
    switch targetState {
        case TARGET_STATE_ONLINE: return "ONLINE"
        case TARGET_STATE_OFFLINE: return "OFFLINE"
        default:
            log.Printf("Error: invalid TargetState represented by '%d'", targetState)
            return ""
    }
}

func (targetState TargetState) StringPtr() *string {
    if targetState == UNDEFINED {
        return nil
    }
    result := targetState.String()
    return &result
}

func newTargetStateFromContent(content []byte, aggErr *AggregateError) *TargetState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(TargetState)
    parseEnum(content, result, aggErr)
    return result
}
type BulkObject struct {
    Bucket *string
    Id *string
    InCache *bool
    Latest bool
    Length int64
    Name *string
    Offset int64
    PhysicalPlacement *PhysicalPlacement
    VersionId string
}

func (bulkObject *BulkObject) parse(xmlNode *XmlNode, aggErr *AggregateError) {
    // Parse Attributes
    for _, attr := range xmlNode.Attrs {
        switch attr.Name.Local {
        case "Bucket":
            bulkObject.Bucket = parseNullableStringFromString(attr.Value)
        case "Id":
            bulkObject.Id = parseNullableStringFromString(attr.Value)
        case "InCache":
            bulkObject.InCache = parseNullableBoolFromString(attr.Value, aggErr)
        case "Latest":
            bulkObject.Latest = parseBoolFromString(attr.Value, aggErr)
        case "Length":
            bulkObject.Length = parseInt64FromString(attr.Value, aggErr)
        case "Name":
            bulkObject.Name = parseNullableStringFromString(attr.Value)
        case "Offset":
            bulkObject.Offset = parseInt64FromString(attr.Value, aggErr)
        case "VersionId":
            bulkObject.VersionId = attr.Value
        default:
            log.Printf("WARNING: unable to parse unknown attribute '%s' while parsing BulkObject.", attr.Name.Local)
        }
    }

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "PhysicalPlacement":
            var model PhysicalPlacement
            model.parse(&child, aggErr)
            bulkObject.PhysicalPlacement = &model
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BulkObject.", child.XMLName.Local)
        }
    }
}

type BulkObjectList struct {
    Objects []BulkObject
}

func (bulkObjectList *BulkObjectList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Object":
            var model BulkObject
            model.parse(&child, aggErr)
            bulkObjectList.Objects = append(bulkObjectList.Objects, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BulkObjectList.", child.XMLName.Local)
        }
    }
}

type BuildInformation struct {
    Branch *string
    Revision *string
    Version *string
}

func (buildInformation *BuildInformation) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Branch":
            buildInformation.Branch = parseNullableString(child.Content)
        case "Revision":
            buildInformation.Revision = parseNullableString(child.Content)
        case "Version":
            buildInformation.Version = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BuildInformation.", child.XMLName.Local)
        }
    }
}

type BlobStoreTaskInformation struct {
    DateScheduled string
    DateStarted *string
    Description *string
    DriveId *string
    Id int64
    Name *string
    PoolId *string
    Priority Priority
    State BlobStoreTaskState
    TapeId *string
    TargetId *string
    TargetType *string
}

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
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BlobStoreTaskInformation.", child.XMLName.Local)
        }
    }
}

type BlobStoreTaskState Enum

const (
    BLOB_STORE_TASK_STATE_NOT_READY BlobStoreTaskState = 1 + iota
    BLOB_STORE_TASK_STATE_READY BlobStoreTaskState = 1 + iota
    BLOB_STORE_TASK_STATE_PENDING_EXECUTION BlobStoreTaskState = 1 + iota
    BLOB_STORE_TASK_STATE_IN_PROGRESS BlobStoreTaskState = 1 + iota
    BLOB_STORE_TASK_STATE_COMPLETED BlobStoreTaskState = 1 + iota
)

func (blobStoreTaskState *BlobStoreTaskState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *blobStoreTaskState = UNDEFINED
        case "NOT_READY": *blobStoreTaskState = BLOB_STORE_TASK_STATE_NOT_READY
        case "READY": *blobStoreTaskState = BLOB_STORE_TASK_STATE_READY
        case "PENDING_EXECUTION": *blobStoreTaskState = BLOB_STORE_TASK_STATE_PENDING_EXECUTION
        case "IN_PROGRESS": *blobStoreTaskState = BLOB_STORE_TASK_STATE_IN_PROGRESS
        case "COMPLETED": *blobStoreTaskState = BLOB_STORE_TASK_STATE_COMPLETED
        default:
            *blobStoreTaskState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into BlobStoreTaskState", str))
    }
    return nil
}

func (blobStoreTaskState BlobStoreTaskState) String() string {
    switch blobStoreTaskState {
        case BLOB_STORE_TASK_STATE_NOT_READY: return "NOT_READY"
        case BLOB_STORE_TASK_STATE_READY: return "READY"
        case BLOB_STORE_TASK_STATE_PENDING_EXECUTION: return "PENDING_EXECUTION"
        case BLOB_STORE_TASK_STATE_IN_PROGRESS: return "IN_PROGRESS"
        case BLOB_STORE_TASK_STATE_COMPLETED: return "COMPLETED"
        default:
            log.Printf("Error: invalid BlobStoreTaskState represented by '%d'", blobStoreTaskState)
            return ""
    }
}

func (blobStoreTaskState BlobStoreTaskState) StringPtr() *string {
    if blobStoreTaskState == UNDEFINED {
        return nil
    }
    result := blobStoreTaskState.String()
    return &result
}

func newBlobStoreTaskStateFromContent(content []byte, aggErr *AggregateError) *BlobStoreTaskState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(BlobStoreTaskState)
    parseEnum(content, result, aggErr)
    return result
}
type BlobStoreTasksInformation struct {
    Tasks []BlobStoreTaskInformation
}

func (blobStoreTasksInformation *BlobStoreTasksInformation) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Tasks":
            var model BlobStoreTaskInformation
            model.parse(&child, aggErr)
            blobStoreTasksInformation.Tasks = append(blobStoreTasksInformation.Tasks, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BlobStoreTasksInformation.", child.XMLName.Local)
        }
    }
}

type CacheEntryInformation struct {
    Blob *Blob
    State CacheEntryState
}

func (cacheEntryInformation *CacheEntryInformation) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Blob":
            var model Blob
            model.parse(&child, aggErr)
            cacheEntryInformation.Blob = &model
        case "State":
            parseEnum(child.Content, &cacheEntryInformation.State, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CacheEntryInformation.", child.XMLName.Local)
        }
    }
}

type CacheEntryState Enum

const (
    CACHE_ENTRY_STATE_ALLOCATED CacheEntryState = 1 + iota
    CACHE_ENTRY_STATE_IN_CACHE CacheEntryState = 1 + iota
)

func (cacheEntryState *CacheEntryState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *cacheEntryState = UNDEFINED
        case "ALLOCATED": *cacheEntryState = CACHE_ENTRY_STATE_ALLOCATED
        case "IN_CACHE": *cacheEntryState = CACHE_ENTRY_STATE_IN_CACHE
        default:
            *cacheEntryState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into CacheEntryState", str))
    }
    return nil
}

func (cacheEntryState CacheEntryState) String() string {
    switch cacheEntryState {
        case CACHE_ENTRY_STATE_ALLOCATED: return "ALLOCATED"
        case CACHE_ENTRY_STATE_IN_CACHE: return "IN_CACHE"
        default:
            log.Printf("Error: invalid CacheEntryState represented by '%d'", cacheEntryState)
            return ""
    }
}

func (cacheEntryState CacheEntryState) StringPtr() *string {
    if cacheEntryState == UNDEFINED {
        return nil
    }
    result := cacheEntryState.String()
    return &result
}

func newCacheEntryStateFromContent(content []byte, aggErr *AggregateError) *CacheEntryState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(CacheEntryState)
    parseEnum(content, result, aggErr)
    return result
}
type CacheFilesystemInformation struct {
    AvailableCapacityInBytes int64
    CacheFilesystem CacheFilesystem
    Entries []CacheEntryInformation
    Summary *string
    TotalCapacityInBytes int64
    UnavailableCapacityInBytes int64
    UsedCapacityInBytes int64
}

func (cacheFilesystemInformation *CacheFilesystemInformation) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AvailableCapacityInBytes":
            cacheFilesystemInformation.AvailableCapacityInBytes = parseInt64(child.Content, aggErr)
        case "CacheFilesystem":
            cacheFilesystemInformation.CacheFilesystem.parse(&child, aggErr)
        case "Entries":
            var model CacheEntryInformation
            model.parse(&child, aggErr)
            cacheFilesystemInformation.Entries = append(cacheFilesystemInformation.Entries, model)
        case "Summary":
            cacheFilesystemInformation.Summary = parseNullableString(child.Content)
        case "TotalCapacityInBytes":
            cacheFilesystemInformation.TotalCapacityInBytes = parseInt64(child.Content, aggErr)
        case "UnavailableCapacityInBytes":
            cacheFilesystemInformation.UnavailableCapacityInBytes = parseInt64(child.Content, aggErr)
        case "UsedCapacityInBytes":
            cacheFilesystemInformation.UsedCapacityInBytes = parseInt64(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CacheFilesystemInformation.", child.XMLName.Local)
        }
    }
}

type CacheInformation struct {
    Filesystems []CacheFilesystemInformation
}

func (cacheInformation *CacheInformation) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Filesystems":
            var model CacheFilesystemInformation
            model.parse(&child, aggErr)
            cacheInformation.Filesystems = append(cacheInformation.Filesystems, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CacheInformation.", child.XMLName.Local)
        }
    }
}

type BucketDetails struct {
    CreationDate *string
    Name *string
}

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

type ListBucketResult struct {
    CommonPrefixes []string
    CreationDate *string
    Delimiter *string
    Marker *string
    MaxKeys int
    Name *string
    NextMarker *string
    Objects []Contents
    Prefix *string
    Truncated bool
    VersionedObjects []Contents
}

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

type ListAllMyBucketsResult struct {
    Buckets []BucketDetails
    Owner User
}

func (listAllMyBucketsResult *ListAllMyBucketsResult) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Buckets":
            listAllMyBucketsResult.Buckets = parseBucketDetailsSlice("Bucket", child.Children, aggErr)
        case "Owner":
            listAllMyBucketsResult.Owner.parse(&child, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing ListAllMyBucketsResult.", child.XMLName.Local)
        }
    }
}

type CompleteMultipartUploadResult struct {
    Bucket *string
    ETag *string
    Key *string
    Location *string
}

func (completeMultipartUploadResult *CompleteMultipartUploadResult) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Bucket":
            completeMultipartUploadResult.Bucket = parseNullableString(child.Content)
        case "ETag":
            completeMultipartUploadResult.ETag = parseNullableString(child.Content)
        case "Key":
            completeMultipartUploadResult.Key = parseNullableString(child.Content)
        case "Location":
            completeMultipartUploadResult.Location = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CompleteMultipartUploadResult.", child.XMLName.Local)
        }
    }
}

type DeleteObjectError struct {
    Code *string
    Key *string
    Message *string
    VersionId *string
}

func (deleteObjectError *DeleteObjectError) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Code":
            deleteObjectError.Code = parseNullableString(child.Content)
        case "Key":
            deleteObjectError.Key = parseNullableString(child.Content)
        case "Message":
            deleteObjectError.Message = parseNullableString(child.Content)
        case "VersionId":
            deleteObjectError.VersionId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DeleteObjectError.", child.XMLName.Local)
        }
    }
}

type DeleteResult struct {
    DeletedObjects []S3ObjectToDelete
    Errors []DeleteObjectError
}

func (deleteResult *DeleteResult) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Deleted":
            var model S3ObjectToDelete
            model.parse(&child, aggErr)
            deleteResult.DeletedObjects = append(deleteResult.DeletedObjects, model)
        case "Error":
            var model DeleteObjectError
            model.parse(&child, aggErr)
            deleteResult.Errors = append(deleteResult.Errors, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DeleteResult.", child.XMLName.Local)
        }
    }
}

type DetailedTapePartition struct {
    AutoCompactionEnabled bool
    DriveType *TapeDriveType
    DriveTypes []TapeDriveType
    ErrorMessage *string
    Id string
    ImportExportConfiguration ImportExportConfiguration
    LibraryId string
    MinimumReadReservedDrives int
    MinimumWriteReservedDrives int
    Name *string
    Quiesced Quiesced
    SerialNumber *string
    State TapePartitionState
    TapeTypes []string
}

func (detailedTapePartition *DetailedTapePartition) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AutoCompactionEnabled":
            detailedTapePartition.AutoCompactionEnabled = parseBool(child.Content, aggErr)
        case "DriveType":
            detailedTapePartition.DriveType = newTapeDriveTypeFromContent(child.Content, aggErr)
        case "DriveTypes":
            var model TapeDriveType
            parseEnum(child.Content, &model, aggErr)
            detailedTapePartition.DriveTypes = append(detailedTapePartition.DriveTypes, model)
        case "ErrorMessage":
            detailedTapePartition.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            detailedTapePartition.Id = parseString(child.Content)
        case "ImportExportConfiguration":
            parseEnum(child.Content, &detailedTapePartition.ImportExportConfiguration, aggErr)
        case "LibraryId":
            detailedTapePartition.LibraryId = parseString(child.Content)
        case "MinimumReadReservedDrives":
            detailedTapePartition.MinimumReadReservedDrives = parseInt(child.Content, aggErr)
        case "MinimumWriteReservedDrives":
            detailedTapePartition.MinimumWriteReservedDrives = parseInt(child.Content, aggErr)
        case "Name":
            detailedTapePartition.Name = parseNullableString(child.Content)
        case "Quiesced":
            parseEnum(child.Content, &detailedTapePartition.Quiesced, aggErr)
        case "SerialNumber":
            detailedTapePartition.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &detailedTapePartition.State, aggErr)
        case "TapeTypes":
            var str = parseString(child.Content)
            detailedTapePartition.TapeTypes = append(detailedTapePartition.TapeTypes, str)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DetailedTapePartition.", child.XMLName.Local)
        }
    }
}

type Error struct {
    Code *string
    HttpErrorCode int
    Message *string
    Resource *string
    ResourceId int64
}

func (error *Error) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Code":
            error.Code = parseNullableString(child.Content)
        case "HttpErrorCode":
            error.HttpErrorCode = parseInt(child.Content, aggErr)
        case "Message":
            error.Message = parseNullableString(child.Content)
        case "Resource":
            error.Resource = parseNullableString(child.Content)
        case "ResourceId":
            error.ResourceId = parseInt64(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Error.", child.XMLName.Local)
        }
    }
}

type InitiateMultipartUploadResult struct {
    Bucket *string
    Key *string
    UploadId *string
}

func (initiateMultipartUploadResult *InitiateMultipartUploadResult) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Bucket":
            initiateMultipartUploadResult.Bucket = parseNullableString(child.Content)
        case "Key":
            initiateMultipartUploadResult.Key = parseNullableString(child.Content)
        case "UploadId":
            initiateMultipartUploadResult.UploadId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing InitiateMultipartUploadResult.", child.XMLName.Local)
        }
    }
}

type Job struct {
    Aggregating bool
    BucketName *string
    CachedSizeInBytes int64
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    CompletedSizeInBytes int64
    EntirelyInCache *bool
    JobId string
    Naked bool
    Name *string
    Nodes []JobNode
    OriginalSizeInBytes int64
    Priority Priority
    RequestType JobRequestType
    StartDate string
    Status JobStatus
    UserId string
    UserName *string
}

func (job *Job) parse(xmlNode *XmlNode, aggErr *AggregateError) {
    // Parse Attributes
    for _, attr := range xmlNode.Attrs {
        switch attr.Name.Local {
        case "Aggregating":
            job.Aggregating = parseBoolFromString(attr.Value, aggErr)
        case "BucketName":
            job.BucketName = parseNullableStringFromString(attr.Value)
        case "CachedSizeInBytes":
            job.CachedSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "ChunkClientProcessingOrderGuarantee":
            parseEnumFromString(attr.Value, &job.ChunkClientProcessingOrderGuarantee, aggErr)
        case "CompletedSizeInBytes":
            job.CompletedSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "EntirelyInCache":
            job.EntirelyInCache = parseNullableBoolFromString(attr.Value, aggErr)
        case "JobId":
            job.JobId = attr.Value
        case "Naked":
            job.Naked = parseBoolFromString(attr.Value, aggErr)
        case "Name":
            job.Name = parseNullableStringFromString(attr.Value)
        case "OriginalSizeInBytes":
            job.OriginalSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "Priority":
            parseEnumFromString(attr.Value, &job.Priority, aggErr)
        case "RequestType":
            parseEnumFromString(attr.Value, &job.RequestType, aggErr)
        case "StartDate":
            job.StartDate = attr.Value
        case "Status":
            parseEnumFromString(attr.Value, &job.Status, aggErr)
        case "UserId":
            job.UserId = attr.Value
        case "UserName":
            job.UserName = parseNullableStringFromString(attr.Value)
        default:
            log.Printf("WARNING: unable to parse unknown attribute '%s' while parsing Job.", attr.Name.Local)
        }
    }

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Nodes":
            job.Nodes = parseJobNodeSlice("Node", child.Children, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Job.", child.XMLName.Local)
        }
    }
}


func parseJobSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []Job {
    var result []Job
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult Job
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing Job struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}

type Objects struct {
    ChunkId string
    ChunkNumber int
    NodeId *string
    Objects []BulkObject
}

func (objects *Objects) parse(xmlNode *XmlNode, aggErr *AggregateError) {
    // Parse Attributes
    for _, attr := range xmlNode.Attrs {
        switch attr.Name.Local {
        case "ChunkId":
            objects.ChunkId = attr.Value
        case "ChunkNumber":
            objects.ChunkNumber = parseIntFromString(attr.Value, aggErr)
        case "NodeId":
            objects.NodeId = parseNullableStringFromString(attr.Value)
        default:
            log.Printf("WARNING: unable to parse unknown attribute '%s' while parsing Objects.", attr.Name.Local)
        }
    }

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Object":
            var model BulkObject
            model.parse(&child, aggErr)
            objects.Objects = append(objects.Objects, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Objects.", child.XMLName.Local)
        }
    }
}

type JobStatus Enum

const (
    JOB_STATUS_IN_PROGRESS JobStatus = 1 + iota
    JOB_STATUS_COMPLETED JobStatus = 1 + iota
    JOB_STATUS_CANCELED JobStatus = 1 + iota
)

func (jobStatus *JobStatus) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *jobStatus = UNDEFINED
        case "IN_PROGRESS": *jobStatus = JOB_STATUS_IN_PROGRESS
        case "COMPLETED": *jobStatus = JOB_STATUS_COMPLETED
        case "CANCELED": *jobStatus = JOB_STATUS_CANCELED
        default:
            *jobStatus = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into JobStatus", str))
    }
    return nil
}

func (jobStatus JobStatus) String() string {
    switch jobStatus {
        case JOB_STATUS_IN_PROGRESS: return "IN_PROGRESS"
        case JOB_STATUS_COMPLETED: return "COMPLETED"
        case JOB_STATUS_CANCELED: return "CANCELED"
        default:
            log.Printf("Error: invalid JobStatus represented by '%d'", jobStatus)
            return ""
    }
}

func (jobStatus JobStatus) StringPtr() *string {
    if jobStatus == UNDEFINED {
        return nil
    }
    result := jobStatus.String()
    return &result
}

func newJobStatusFromContent(content []byte, aggErr *AggregateError) *JobStatus {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(JobStatus)
    parseEnum(content, result, aggErr)
    return result
}
type MasterObjectList struct {
    Aggregating bool
    BucketName *string
    CachedSizeInBytes int64
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    CompletedSizeInBytes int64
    EntirelyInCache *bool
    JobId string
    Naked bool
    Name *string
    Nodes []JobNode
    Objects []Objects
    OriginalSizeInBytes int64
    Priority Priority
    RequestType JobRequestType
    StartDate string
    Status JobStatus
    UserId string
    UserName *string
}

func (masterObjectList *MasterObjectList) parse(xmlNode *XmlNode, aggErr *AggregateError) {
    // Parse Attributes
    for _, attr := range xmlNode.Attrs {
        switch attr.Name.Local {
        case "Aggregating":
            masterObjectList.Aggregating = parseBoolFromString(attr.Value, aggErr)
        case "BucketName":
            masterObjectList.BucketName = parseNullableStringFromString(attr.Value)
        case "CachedSizeInBytes":
            masterObjectList.CachedSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "ChunkClientProcessingOrderGuarantee":
            parseEnumFromString(attr.Value, &masterObjectList.ChunkClientProcessingOrderGuarantee, aggErr)
        case "CompletedSizeInBytes":
            masterObjectList.CompletedSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "EntirelyInCache":
            masterObjectList.EntirelyInCache = parseNullableBoolFromString(attr.Value, aggErr)
        case "JobId":
            masterObjectList.JobId = attr.Value
        case "Naked":
            masterObjectList.Naked = parseBoolFromString(attr.Value, aggErr)
        case "Name":
            masterObjectList.Name = parseNullableStringFromString(attr.Value)
        case "OriginalSizeInBytes":
            masterObjectList.OriginalSizeInBytes = parseInt64FromString(attr.Value, aggErr)
        case "Priority":
            parseEnumFromString(attr.Value, &masterObjectList.Priority, aggErr)
        case "RequestType":
            parseEnumFromString(attr.Value, &masterObjectList.RequestType, aggErr)
        case "StartDate":
            masterObjectList.StartDate = attr.Value
        case "Status":
            parseEnumFromString(attr.Value, &masterObjectList.Status, aggErr)
        case "UserId":
            masterObjectList.UserId = attr.Value
        case "UserName":
            masterObjectList.UserName = parseNullableStringFromString(attr.Value)
        default:
            log.Printf("WARNING: unable to parse unknown attribute '%s' while parsing MasterObjectList.", attr.Name.Local)
        }
    }

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Nodes":
            masterObjectList.Nodes = parseJobNodeSlice("Node", child.Children, aggErr)
        case "Objects":
            var model Objects
            model.parse(&child, aggErr)
            masterObjectList.Objects = append(masterObjectList.Objects, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing MasterObjectList.", child.XMLName.Local)
        }
    }
}

type JobList struct {
    Jobs []Job
}

func (jobList *JobList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Job":
            var model Job
            model.parse(&child, aggErr)
            jobList.Jobs = append(jobList.Jobs, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing JobList.", child.XMLName.Local)
        }
    }
}

type ListPartsResult struct {
    Bucket *string
    Key *string
    MaxParts int
    NextPartNumberMarker int
    Owner User
    PartNumberMarker *int
    Parts []MultiPartUploadPart
    Truncated bool
    UploadId string
}

func (listPartsResult *ListPartsResult) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Bucket":
            listPartsResult.Bucket = parseNullableString(child.Content)
        case "Key":
            listPartsResult.Key = parseNullableString(child.Content)
        case "MaxParts":
            listPartsResult.MaxParts = parseInt(child.Content, aggErr)
        case "NextPartNumberMarker":
            listPartsResult.NextPartNumberMarker = parseInt(child.Content, aggErr)
        case "Owner":
            listPartsResult.Owner.parse(&child, aggErr)
        case "PartNumberMarker":
            listPartsResult.PartNumberMarker = parseNullableInt(child.Content, aggErr)
        case "Part":
            var model MultiPartUploadPart
            model.parse(&child, aggErr)
            listPartsResult.Parts = append(listPartsResult.Parts, model)
        case "IsTruncated":
            listPartsResult.Truncated = parseBool(child.Content, aggErr)
        case "UploadId":
            listPartsResult.UploadId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing ListPartsResult.", child.XMLName.Local)
        }
    }
}

type ListMultiPartUploadsResult struct {
    Bucket *string
    CommonPrefixes []string
    Delimiter *string
    KeyMarker *string
    MaxUploads int
    NextKeyMarker *string
    NextUploadIdMarker *string
    Prefix *string
    Truncated bool
    UploadIdMarker *string
    Uploads []MultiPartUpload
}

func (listMultiPartUploadsResult *ListMultiPartUploadsResult) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Bucket":
            listMultiPartUploadsResult.Bucket = parseNullableString(child.Content)
        case "CommonPrefixes":
            var prefixes []string
            prefixes = parseStringSlice("Prefix", child.Children, aggErr)
            listMultiPartUploadsResult.CommonPrefixes = append(listMultiPartUploadsResult.CommonPrefixes, prefixes...)
        case "Delimiter":
            listMultiPartUploadsResult.Delimiter = parseNullableString(child.Content)
        case "KeyMarker":
            listMultiPartUploadsResult.KeyMarker = parseNullableString(child.Content)
        case "MaxUploads":
            listMultiPartUploadsResult.MaxUploads = parseInt(child.Content, aggErr)
        case "NextKeyMarker":
            listMultiPartUploadsResult.NextKeyMarker = parseNullableString(child.Content)
        case "NextUploadIdMarker":
            listMultiPartUploadsResult.NextUploadIdMarker = parseNullableString(child.Content)
        case "Prefix":
            listMultiPartUploadsResult.Prefix = parseNullableString(child.Content)
        case "IsTruncated":
            listMultiPartUploadsResult.Truncated = parseBool(child.Content, aggErr)
        case "UploadIdMarker":
            listMultiPartUploadsResult.UploadIdMarker = parseNullableString(child.Content)
        case "Upload":
            var model MultiPartUpload
            model.parse(&child, aggErr)
            listMultiPartUploadsResult.Uploads = append(listMultiPartUploadsResult.Uploads, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing ListMultiPartUploadsResult.", child.XMLName.Local)
        }
    }
}

type MultiPartUpload struct {
    Initiated string
    Key *string
    Owner User
    UploadId string
}

func (multiPartUpload *MultiPartUpload) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Initiated":
            multiPartUpload.Initiated = parseString(child.Content)
        case "Key":
            multiPartUpload.Key = parseNullableString(child.Content)
        case "Owner":
            multiPartUpload.Owner.parse(&child, aggErr)
        case "UploadId":
            multiPartUpload.UploadId = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing MultiPartUpload.", child.XMLName.Local)
        }
    }
}

type MultiPartUploadPart struct {
    ETag *string
    LastModified string
    PartNumber int
}

func (multiPartUploadPart *MultiPartUploadPart) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "ETag":
            multiPartUploadPart.ETag = parseNullableString(child.Content)
        case "LastModified":
            multiPartUploadPart.LastModified = parseString(child.Content)
        case "PartNumber":
            multiPartUploadPart.PartNumber = parseInt(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing MultiPartUploadPart.", child.XMLName.Local)
        }
    }
}

type JobNode struct {
    EndPoint *string
    HttpPort *int
    HttpsPort *int
    Id string
}

func (jobNode *JobNode) parse(xmlNode *XmlNode, aggErr *AggregateError) {
    // Parse Attributes
    for _, attr := range xmlNode.Attrs {
        switch attr.Name.Local {
        case "EndPoint":
            jobNode.EndPoint = parseNullableStringFromString(attr.Value)
        case "HttpPort":
            jobNode.HttpPort = parseNullableIntFromString(attr.Value, aggErr)
        case "HttpsPort":
            jobNode.HttpsPort = parseNullableIntFromString(attr.Value, aggErr)
        case "Id":
            jobNode.Id = attr.Value
        default:
            log.Printf("WARNING: unable to parse unknown attribute '%s' while parsing JobNode.", attr.Name.Local)
        }
    }

}


func parseJobNodeSlice(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []JobNode {
    var result []JobNode
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult JobNode
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s' when parsing JobNode struct.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}

type Contents struct {
    ETag *string
    IsLatest *bool
    Key *string
    LastModified *string
    Owner User
    Size int64
    StorageClass *string
    VersionId *string
}

func (contents *Contents) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "ETag":
            contents.ETag = parseNullableString(child.Content)
        case "IsLatest":
            contents.IsLatest = parseNullableBool(child.Content, aggErr)
        case "Key":
            contents.Key = parseNullableString(child.Content)
        case "LastModified":
            contents.LastModified = parseNullableString(child.Content)
        case "Owner":
            contents.Owner.parse(&child, aggErr)
        case "Size":
            contents.Size = parseInt64(child.Content, aggErr)
        case "StorageClass":
            contents.StorageClass = parseNullableString(child.Content)
        case "VersionId":
            contents.VersionId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Contents.", child.XMLName.Local)
        }
    }
}

type S3ObjectToDelete struct {
    Key *string
    VersionId *string
}

func (s3ObjectToDelete *S3ObjectToDelete) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Key":
            s3ObjectToDelete.Key = parseNullableString(child.Content)
        case "VersionId":
            s3ObjectToDelete.VersionId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3ObjectToDelete.", child.XMLName.Local)
        }
    }
}

type User struct {
    DisplayName *string
    Id string
}

func (user *User) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DisplayName":
            user.DisplayName = parseNullableString(child.Content)
        case "ID":
            user.Id = parseString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing User.", child.XMLName.Local)
        }
    }
}

type DetailedS3Object struct {
    Blobs *BulkObjectList
    BlobsBeingPersisted *int
    BlobsDegraded *int
    BlobsInCache *int
    BlobsTotal *int
    BucketId string
    CreationDate *string
    ETag *string
    Id string
    Latest bool
    Name *string
    Owner *string
    Size int64
    Type S3ObjectType
}

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
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DetailedS3Object.", child.XMLName.Local)
        }
    }
}

type SystemInformation struct {
    ApiVersion *string
    BackendActivated bool
    BuildInformation BuildInformation
    InstanceId string
    Now int64
    SerialNumber *string
}

func (systemInformation *SystemInformation) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "ApiVersion":
            systemInformation.ApiVersion = parseNullableString(child.Content)
        case "BackendActivated":
            systemInformation.BackendActivated = parseBool(child.Content, aggErr)
        case "BuildInformation":
            systemInformation.BuildInformation.parse(&child, aggErr)
        case "InstanceId":
            systemInformation.InstanceId = parseString(child.Content)
        case "Now":
            systemInformation.Now = parseInt64(child.Content, aggErr)
        case "SerialNumber":
            systemInformation.SerialNumber = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SystemInformation.", child.XMLName.Local)
        }
    }
}

type HealthVerificationResult struct {
    DatabaseFilesystemFreeSpace DatabasePhysicalSpaceState
    MsRequiredToVerifyDataPlannerHealth int64
}

func (healthVerificationResult *HealthVerificationResult) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DatabaseFilesystemFreeSpace":
            parseEnum(child.Content, &healthVerificationResult.DatabaseFilesystemFreeSpace, aggErr)
        case "MsRequiredToVerifyDataPlannerHealth":
            healthVerificationResult.MsRequiredToVerifyDataPlannerHealth = parseInt64(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing HealthVerificationResult.", child.XMLName.Local)
        }
    }
}

type NamedDetailedTapePartition struct {
    AutoCompactionEnabled bool
    DriveType *TapeDriveType
    DriveTypes []TapeDriveType
    ErrorMessage *string
    Id string
    ImportExportConfiguration ImportExportConfiguration
    LibraryId string
    MinimumReadReservedDrives int
    MinimumWriteReservedDrives int
    Name *string
    Quiesced Quiesced
    SerialNumber *string
    State TapePartitionState
    TapeTypes []string
}

func (namedDetailedTapePartition *NamedDetailedTapePartition) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AutoCompactionEnabled":
            namedDetailedTapePartition.AutoCompactionEnabled = parseBool(child.Content, aggErr)
        case "DriveType":
            namedDetailedTapePartition.DriveType = newTapeDriveTypeFromContent(child.Content, aggErr)
        case "DriveTypes":
            var model TapeDriveType
            parseEnum(child.Content, &model, aggErr)
            namedDetailedTapePartition.DriveTypes = append(namedDetailedTapePartition.DriveTypes, model)
        case "ErrorMessage":
            namedDetailedTapePartition.ErrorMessage = parseNullableString(child.Content)
        case "Id":
            namedDetailedTapePartition.Id = parseString(child.Content)
        case "ImportExportConfiguration":
            parseEnum(child.Content, &namedDetailedTapePartition.ImportExportConfiguration, aggErr)
        case "LibraryId":
            namedDetailedTapePartition.LibraryId = parseString(child.Content)
        case "MinimumReadReservedDrives":
            namedDetailedTapePartition.MinimumReadReservedDrives = parseInt(child.Content, aggErr)
        case "MinimumWriteReservedDrives":
            namedDetailedTapePartition.MinimumWriteReservedDrives = parseInt(child.Content, aggErr)
        case "Name":
            namedDetailedTapePartition.Name = parseNullableString(child.Content)
        case "Quiesced":
            parseEnum(child.Content, &namedDetailedTapePartition.Quiesced, aggErr)
        case "SerialNumber":
            namedDetailedTapePartition.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &namedDetailedTapePartition.State, aggErr)
        case "TapeTypes":
            var str = parseString(child.Content)
            namedDetailedTapePartition.TapeTypes = append(namedDetailedTapePartition.TapeTypes, str)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing NamedDetailedTapePartition.", child.XMLName.Local)
        }
    }
}

type TapeFailure struct {
    Cause *string
    Tape Tape
}

func (tapeFailure *TapeFailure) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Cause":
            tapeFailure.Cause = parseNullableString(child.Content)
        case "Tape":
            tapeFailure.Tape.parse(&child, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeFailure.", child.XMLName.Local)
        }
    }
}

type TapeFailureList struct {
    Failures []TapeFailure
}

func (tapeFailureList *TapeFailureList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Failure":
            var model TapeFailure
            model.parse(&child, aggErr)
            tapeFailureList.Failures = append(tapeFailureList.Failures, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeFailureList.", child.XMLName.Local)
        }
    }
}

type RestOperationType Enum

const (
    REST_OPERATION_TYPE_ALLOCATE RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CANCEL_EJECT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CANCEL_FORMAT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CANCEL_IMPORT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CANCEL_ONLINE RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CANCEL_VERIFY RestOperationType = 1 + iota
    REST_OPERATION_TYPE_CLEAN RestOperationType = 1 + iota
    REST_OPERATION_TYPE_COMPACT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_DEALLOCATE RestOperationType = 1 + iota
    REST_OPERATION_TYPE_EJECT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_FORMAT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_GET_PHYSICAL_PLACEMENT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_IMPORT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_INSPECT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_ONLINE RestOperationType = 1 + iota
    REST_OPERATION_TYPE_PAIR_BACK RestOperationType = 1 + iota
    REST_OPERATION_TYPE_REGENERATE_SECRET_KEY RestOperationType = 1 + iota
    REST_OPERATION_TYPE_START_BULK_GET RestOperationType = 1 + iota
    REST_OPERATION_TYPE_START_BULK_PUT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_START_BULK_STAGE RestOperationType = 1 + iota
    REST_OPERATION_TYPE_START_BULK_VERIFY RestOperationType = 1 + iota
    REST_OPERATION_TYPE_VERIFY RestOperationType = 1 + iota
    REST_OPERATION_TYPE_VERIFY_SAFE_TO_START_BULK_PUT RestOperationType = 1 + iota
    REST_OPERATION_TYPE_VERIFY_PHYSICAL_PLACEMENT RestOperationType = 1 + iota
)

func (restOperationType *RestOperationType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *restOperationType = UNDEFINED
        case "ALLOCATE": *restOperationType = REST_OPERATION_TYPE_ALLOCATE
        case "CANCEL_EJECT": *restOperationType = REST_OPERATION_TYPE_CANCEL_EJECT
        case "CANCEL_FORMAT": *restOperationType = REST_OPERATION_TYPE_CANCEL_FORMAT
        case "CANCEL_IMPORT": *restOperationType = REST_OPERATION_TYPE_CANCEL_IMPORT
        case "CANCEL_ONLINE": *restOperationType = REST_OPERATION_TYPE_CANCEL_ONLINE
        case "CANCEL_VERIFY": *restOperationType = REST_OPERATION_TYPE_CANCEL_VERIFY
        case "CLEAN": *restOperationType = REST_OPERATION_TYPE_CLEAN
        case "COMPACT": *restOperationType = REST_OPERATION_TYPE_COMPACT
        case "DEALLOCATE": *restOperationType = REST_OPERATION_TYPE_DEALLOCATE
        case "EJECT": *restOperationType = REST_OPERATION_TYPE_EJECT
        case "FORMAT": *restOperationType = REST_OPERATION_TYPE_FORMAT
        case "GET_PHYSICAL_PLACEMENT": *restOperationType = REST_OPERATION_TYPE_GET_PHYSICAL_PLACEMENT
        case "IMPORT": *restOperationType = REST_OPERATION_TYPE_IMPORT
        case "INSPECT": *restOperationType = REST_OPERATION_TYPE_INSPECT
        case "ONLINE": *restOperationType = REST_OPERATION_TYPE_ONLINE
        case "PAIR_BACK": *restOperationType = REST_OPERATION_TYPE_PAIR_BACK
        case "REGENERATE_SECRET_KEY": *restOperationType = REST_OPERATION_TYPE_REGENERATE_SECRET_KEY
        case "START_BULK_GET": *restOperationType = REST_OPERATION_TYPE_START_BULK_GET
        case "START_BULK_PUT": *restOperationType = REST_OPERATION_TYPE_START_BULK_PUT
        case "START_BULK_STAGE": *restOperationType = REST_OPERATION_TYPE_START_BULK_STAGE
        case "START_BULK_VERIFY": *restOperationType = REST_OPERATION_TYPE_START_BULK_VERIFY
        case "VERIFY": *restOperationType = REST_OPERATION_TYPE_VERIFY
        case "VERIFY_SAFE_TO_START_BULK_PUT": *restOperationType = REST_OPERATION_TYPE_VERIFY_SAFE_TO_START_BULK_PUT
        case "VERIFY_PHYSICAL_PLACEMENT": *restOperationType = REST_OPERATION_TYPE_VERIFY_PHYSICAL_PLACEMENT
        default:
            *restOperationType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into RestOperationType", str))
    }
    return nil
}

func (restOperationType RestOperationType) String() string {
    switch restOperationType {
        case REST_OPERATION_TYPE_ALLOCATE: return "ALLOCATE"
        case REST_OPERATION_TYPE_CANCEL_EJECT: return "CANCEL_EJECT"
        case REST_OPERATION_TYPE_CANCEL_FORMAT: return "CANCEL_FORMAT"
        case REST_OPERATION_TYPE_CANCEL_IMPORT: return "CANCEL_IMPORT"
        case REST_OPERATION_TYPE_CANCEL_ONLINE: return "CANCEL_ONLINE"
        case REST_OPERATION_TYPE_CANCEL_VERIFY: return "CANCEL_VERIFY"
        case REST_OPERATION_TYPE_CLEAN: return "CLEAN"
        case REST_OPERATION_TYPE_COMPACT: return "COMPACT"
        case REST_OPERATION_TYPE_DEALLOCATE: return "DEALLOCATE"
        case REST_OPERATION_TYPE_EJECT: return "EJECT"
        case REST_OPERATION_TYPE_FORMAT: return "FORMAT"
        case REST_OPERATION_TYPE_GET_PHYSICAL_PLACEMENT: return "GET_PHYSICAL_PLACEMENT"
        case REST_OPERATION_TYPE_IMPORT: return "IMPORT"
        case REST_OPERATION_TYPE_INSPECT: return "INSPECT"
        case REST_OPERATION_TYPE_ONLINE: return "ONLINE"
        case REST_OPERATION_TYPE_PAIR_BACK: return "PAIR_BACK"
        case REST_OPERATION_TYPE_REGENERATE_SECRET_KEY: return "REGENERATE_SECRET_KEY"
        case REST_OPERATION_TYPE_START_BULK_GET: return "START_BULK_GET"
        case REST_OPERATION_TYPE_START_BULK_PUT: return "START_BULK_PUT"
        case REST_OPERATION_TYPE_START_BULK_STAGE: return "START_BULK_STAGE"
        case REST_OPERATION_TYPE_START_BULK_VERIFY: return "START_BULK_VERIFY"
        case REST_OPERATION_TYPE_VERIFY: return "VERIFY"
        case REST_OPERATION_TYPE_VERIFY_SAFE_TO_START_BULK_PUT: return "VERIFY_SAFE_TO_START_BULK_PUT"
        case REST_OPERATION_TYPE_VERIFY_PHYSICAL_PLACEMENT: return "VERIFY_PHYSICAL_PLACEMENT"
        default:
            log.Printf("Error: invalid RestOperationType represented by '%d'", restOperationType)
            return ""
    }
}

func (restOperationType RestOperationType) StringPtr() *string {
    if restOperationType == UNDEFINED {
        return nil
    }
    result := restOperationType.String()
    return &result
}

func newRestOperationTypeFromContent(content []byte, aggErr *AggregateError) *RestOperationType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(RestOperationType)
    parseEnum(content, result, aggErr)
    return result
}
type DatabasePhysicalSpaceState Enum

const (
    DATABASE_PHYSICAL_SPACE_STATE_CRITICAL DatabasePhysicalSpaceState = 1 + iota
    DATABASE_PHYSICAL_SPACE_STATE_LOW DatabasePhysicalSpaceState = 1 + iota
    DATABASE_PHYSICAL_SPACE_STATE_NEAR_LOW DatabasePhysicalSpaceState = 1 + iota
    DATABASE_PHYSICAL_SPACE_STATE_NORMAL DatabasePhysicalSpaceState = 1 + iota
)

func (databasePhysicalSpaceState *DatabasePhysicalSpaceState) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *databasePhysicalSpaceState = UNDEFINED
        case "CRITICAL": *databasePhysicalSpaceState = DATABASE_PHYSICAL_SPACE_STATE_CRITICAL
        case "LOW": *databasePhysicalSpaceState = DATABASE_PHYSICAL_SPACE_STATE_LOW
        case "NEAR_LOW": *databasePhysicalSpaceState = DATABASE_PHYSICAL_SPACE_STATE_NEAR_LOW
        case "NORMAL": *databasePhysicalSpaceState = DATABASE_PHYSICAL_SPACE_STATE_NORMAL
        default:
            *databasePhysicalSpaceState = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into DatabasePhysicalSpaceState", str))
    }
    return nil
}

func (databasePhysicalSpaceState DatabasePhysicalSpaceState) String() string {
    switch databasePhysicalSpaceState {
        case DATABASE_PHYSICAL_SPACE_STATE_CRITICAL: return "CRITICAL"
        case DATABASE_PHYSICAL_SPACE_STATE_LOW: return "LOW"
        case DATABASE_PHYSICAL_SPACE_STATE_NEAR_LOW: return "NEAR_LOW"
        case DATABASE_PHYSICAL_SPACE_STATE_NORMAL: return "NORMAL"
        default:
            log.Printf("Error: invalid DatabasePhysicalSpaceState represented by '%d'", databasePhysicalSpaceState)
            return ""
    }
}

func (databasePhysicalSpaceState DatabasePhysicalSpaceState) StringPtr() *string {
    if databasePhysicalSpaceState == UNDEFINED {
        return nil
    }
    result := databasePhysicalSpaceState.String()
    return &result
}

func newDatabasePhysicalSpaceStateFromContent(content []byte, aggErr *AggregateError) *DatabasePhysicalSpaceState {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(DatabasePhysicalSpaceState)
    parseEnum(content, result, aggErr)
    return result
}
type HttpResponseFormatType Enum

const (
    HTTP_RESPONSE_FORMAT_TYPE_DEFAULT HttpResponseFormatType = 1 + iota
    HTTP_RESPONSE_FORMAT_TYPE_JSON HttpResponseFormatType = 1 + iota
    HTTP_RESPONSE_FORMAT_TYPE_XML HttpResponseFormatType = 1 + iota
)

func (httpResponseFormatType *HttpResponseFormatType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *httpResponseFormatType = UNDEFINED
        case "DEFAULT": *httpResponseFormatType = HTTP_RESPONSE_FORMAT_TYPE_DEFAULT
        case "JSON": *httpResponseFormatType = HTTP_RESPONSE_FORMAT_TYPE_JSON
        case "XML": *httpResponseFormatType = HTTP_RESPONSE_FORMAT_TYPE_XML
        default:
            *httpResponseFormatType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into HttpResponseFormatType", str))
    }
    return nil
}

func (httpResponseFormatType HttpResponseFormatType) String() string {
    switch httpResponseFormatType {
        case HTTP_RESPONSE_FORMAT_TYPE_DEFAULT: return "DEFAULT"
        case HTTP_RESPONSE_FORMAT_TYPE_JSON: return "JSON"
        case HTTP_RESPONSE_FORMAT_TYPE_XML: return "XML"
        default:
            log.Printf("Error: invalid HttpResponseFormatType represented by '%d'", httpResponseFormatType)
            return ""
    }
}

func (httpResponseFormatType HttpResponseFormatType) StringPtr() *string {
    if httpResponseFormatType == UNDEFINED {
        return nil
    }
    result := httpResponseFormatType.String()
    return &result
}

func newHttpResponseFormatTypeFromContent(content []byte, aggErr *AggregateError) *HttpResponseFormatType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(HttpResponseFormatType)
    parseEnum(content, result, aggErr)
    return result
}
type RequestType Enum

const (
    REQUEST_TYPE_DELETE RequestType = 1 + iota
    REQUEST_TYPE_GET RequestType = 1 + iota
    REQUEST_TYPE_HEAD RequestType = 1 + iota
    REQUEST_TYPE_POST RequestType = 1 + iota
    REQUEST_TYPE_PUT RequestType = 1 + iota
)

func (requestType *RequestType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *requestType = UNDEFINED
        case "DELETE": *requestType = REQUEST_TYPE_DELETE
        case "GET": *requestType = REQUEST_TYPE_GET
        case "HEAD": *requestType = REQUEST_TYPE_HEAD
        case "POST": *requestType = REQUEST_TYPE_POST
        case "PUT": *requestType = REQUEST_TYPE_PUT
        default:
            *requestType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into RequestType", str))
    }
    return nil
}

func (requestType RequestType) String() string {
    switch requestType {
        case REQUEST_TYPE_DELETE: return "DELETE"
        case REQUEST_TYPE_GET: return "GET"
        case REQUEST_TYPE_HEAD: return "HEAD"
        case REQUEST_TYPE_POST: return "POST"
        case REQUEST_TYPE_PUT: return "PUT"
        default:
            log.Printf("Error: invalid RequestType represented by '%d'", requestType)
            return ""
    }
}

func (requestType RequestType) StringPtr() *string {
    if requestType == UNDEFINED {
        return nil
    }
    result := requestType.String()
    return &result
}

func newRequestTypeFromContent(content []byte, aggErr *AggregateError) *RequestType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(RequestType)
    parseEnum(content, result, aggErr)
    return result
}
type NamingConventionType Enum

const (
    NAMING_CONVENTION_TYPE_CONCAT_LOWERCASE NamingConventionType = 1 + iota
    NAMING_CONVENTION_TYPE_CONSTANT NamingConventionType = 1 + iota
    NAMING_CONVENTION_TYPE_UNDERSCORED NamingConventionType = 1 + iota
    NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_UPPERCASE NamingConventionType = 1 + iota
    NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_LOWERCASE NamingConventionType = 1 + iota
)

func (namingConventionType *NamingConventionType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *namingConventionType = UNDEFINED
        case "CONCAT_LOWERCASE": *namingConventionType = NAMING_CONVENTION_TYPE_CONCAT_LOWERCASE
        case "CONSTANT": *namingConventionType = NAMING_CONVENTION_TYPE_CONSTANT
        case "UNDERSCORED": *namingConventionType = NAMING_CONVENTION_TYPE_UNDERSCORED
        case "CAMEL_CASE_WITH_FIRST_LETTER_UPPERCASE": *namingConventionType = NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_UPPERCASE
        case "CAMEL_CASE_WITH_FIRST_LETTER_LOWERCASE": *namingConventionType = NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_LOWERCASE
        default:
            *namingConventionType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into NamingConventionType", str))
    }
    return nil
}

func (namingConventionType NamingConventionType) String() string {
    switch namingConventionType {
        case NAMING_CONVENTION_TYPE_CONCAT_LOWERCASE: return "CONCAT_LOWERCASE"
        case NAMING_CONVENTION_TYPE_CONSTANT: return "CONSTANT"
        case NAMING_CONVENTION_TYPE_UNDERSCORED: return "UNDERSCORED"
        case NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_UPPERCASE: return "CAMEL_CASE_WITH_FIRST_LETTER_UPPERCASE"
        case NAMING_CONVENTION_TYPE_CAMEL_CASE_WITH_FIRST_LETTER_LOWERCASE: return "CAMEL_CASE_WITH_FIRST_LETTER_LOWERCASE"
        default:
            log.Printf("Error: invalid NamingConventionType represented by '%d'", namingConventionType)
            return ""
    }
}

func (namingConventionType NamingConventionType) StringPtr() *string {
    if namingConventionType == UNDEFINED {
        return nil
    }
    result := namingConventionType.String()
    return &result
}

func newNamingConventionTypeFromContent(content []byte, aggErr *AggregateError) *NamingConventionType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(NamingConventionType)
    parseEnum(content, result, aggErr)
    return result
}
type ChecksumType Enum

const (
    CHECKSUM_TYPE_CRC_32 ChecksumType = 1 + iota
    CHECKSUM_TYPE_CRC_32C ChecksumType = 1 + iota
    CHECKSUM_TYPE_MD5 ChecksumType = 1 + iota
    CHECKSUM_TYPE_SHA_256 ChecksumType = 1 + iota
    CHECKSUM_TYPE_SHA_512 ChecksumType = 1 + iota
    NONE ChecksumType = 1 + iota
)

func (checksumType *ChecksumType) UnmarshalText(text []byte) error {
    var str string = string(bytes.ToUpper(text))
    switch str {
        case "": *checksumType = UNDEFINED
        case "CRC_32": *checksumType = CHECKSUM_TYPE_CRC_32
        case "CRC_32C": *checksumType = CHECKSUM_TYPE_CRC_32C
        case "MD5": *checksumType = CHECKSUM_TYPE_MD5
        case "SHA_256": *checksumType = CHECKSUM_TYPE_SHA_256
        case "SHA_512": *checksumType = CHECKSUM_TYPE_SHA_512
        default:
            *checksumType = UNDEFINED
            return errors.New(fmt.Sprintf("Cannot marshal '%s' into ChecksumType", str))
    }
    return nil
}

func (checksumType ChecksumType) String() string {
    switch checksumType {
        case CHECKSUM_TYPE_CRC_32: return "CRC_32"
        case CHECKSUM_TYPE_CRC_32C: return "CRC_32C"
        case CHECKSUM_TYPE_MD5: return "MD5"
        case CHECKSUM_TYPE_SHA_256: return "SHA_256"
        case CHECKSUM_TYPE_SHA_512: return "SHA_512"
        default:
            log.Printf("Error: invalid ChecksumType represented by '%d'", checksumType)
            return ""
    }
}

func (checksumType ChecksumType) StringPtr() *string {
    if checksumType == UNDEFINED {
        return nil
    }
    result := checksumType.String()
    return &result
}

func newChecksumTypeFromContent(content []byte, aggErr *AggregateError) *ChecksumType {
    if len(content) == 0 {
        // no value
        return nil
    }
    result := new(ChecksumType)
    parseEnum(content, result, aggErr)
    return result
}
type BucketAclList struct {
    BucketAcls []BucketAcl
}

func (bucketAclList *BucketAclList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketAcl":
            var model BucketAcl
            model.parse(&child, aggErr)
            bucketAclList.BucketAcls = append(bucketAclList.BucketAcls, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BucketAclList.", child.XMLName.Local)
        }
    }
}

type DataPolicyAclList struct {
    DataPolicyAcls []DataPolicyAcl
}

func (dataPolicyAclList *DataPolicyAclList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DataPolicyAcl":
            var model DataPolicyAcl
            model.parse(&child, aggErr)
            dataPolicyAclList.DataPolicyAcls = append(dataPolicyAclList.DataPolicyAcls, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DataPolicyAclList.", child.XMLName.Local)
        }
    }
}

type BucketList struct {
    Buckets []Bucket
}

func (bucketList *BucketList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Bucket":
            var model Bucket
            model.parse(&child, aggErr)
            bucketList.Buckets = append(bucketList.Buckets, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BucketList.", child.XMLName.Local)
        }
    }
}

type CacheFilesystemList struct {
    CacheFilesystems []CacheFilesystem
}

func (cacheFilesystemList *CacheFilesystemList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CacheFilesystem":
            var model CacheFilesystem
            model.parse(&child, aggErr)
            cacheFilesystemList.CacheFilesystems = append(cacheFilesystemList.CacheFilesystems, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CacheFilesystemList.", child.XMLName.Local)
        }
    }
}

type AzureDataReplicationRuleList struct {
    AzureDataReplicationRules []AzureDataReplicationRule
}

func (azureDataReplicationRuleList *AzureDataReplicationRuleList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AzureDataReplicationRule":
            var model AzureDataReplicationRule
            model.parse(&child, aggErr)
            azureDataReplicationRuleList.AzureDataReplicationRules = append(azureDataReplicationRuleList.AzureDataReplicationRules, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureDataReplicationRuleList.", child.XMLName.Local)
        }
    }
}

type DataPersistenceRuleList struct {
    DataPersistenceRules []DataPersistenceRule
}

func (dataPersistenceRuleList *DataPersistenceRuleList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DataPersistenceRule":
            var model DataPersistenceRule
            model.parse(&child, aggErr)
            dataPersistenceRuleList.DataPersistenceRules = append(dataPersistenceRuleList.DataPersistenceRules, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DataPersistenceRuleList.", child.XMLName.Local)
        }
    }
}

type DataPolicyList struct {
    DataPolicies []DataPolicy
}

func (dataPolicyList *DataPolicyList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DataPolicy":
            var model DataPolicy
            model.parse(&child, aggErr)
            dataPolicyList.DataPolicies = append(dataPolicyList.DataPolicies, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DataPolicyList.", child.XMLName.Local)
        }
    }
}

type Ds3DataReplicationRuleList struct {
    Ds3DataReplicationRules []Ds3DataReplicationRule
}

func (ds3DataReplicationRuleList *Ds3DataReplicationRuleList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Ds3DataReplicationRule":
            var model Ds3DataReplicationRule
            model.parse(&child, aggErr)
            ds3DataReplicationRuleList.Ds3DataReplicationRules = append(ds3DataReplicationRuleList.Ds3DataReplicationRules, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3DataReplicationRuleList.", child.XMLName.Local)
        }
    }
}

type S3DataReplicationRuleList struct {
    S3DataReplicationRules []S3DataReplicationRule
}

func (s3DataReplicationRuleList *S3DataReplicationRuleList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "S3DataReplicationRule":
            var model S3DataReplicationRule
            model.parse(&child, aggErr)
            s3DataReplicationRuleList.S3DataReplicationRules = append(s3DataReplicationRuleList.S3DataReplicationRules, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3DataReplicationRuleList.", child.XMLName.Local)
        }
    }
}

type DegradedBlobList struct {
    DegradedBlobs []DegradedBlob
}

func (degradedBlobList *DegradedBlobList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "DegradedBlob":
            var model DegradedBlob
            model.parse(&child, aggErr)
            degradedBlobList.DegradedBlobs = append(degradedBlobList.DegradedBlobs, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DegradedBlobList.", child.XMLName.Local)
        }
    }
}

type SuspectBlobAzureTargetList struct {
    SuspectBlobAzureTargets []SuspectBlobAzureTarget
}

func (suspectBlobAzureTargetList *SuspectBlobAzureTargetList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "SuspectBlobAzureTarget":
            var model SuspectBlobAzureTarget
            model.parse(&child, aggErr)
            suspectBlobAzureTargetList.SuspectBlobAzureTargets = append(suspectBlobAzureTargetList.SuspectBlobAzureTargets, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobAzureTargetList.", child.XMLName.Local)
        }
    }
}

type SuspectBlobDs3TargetList struct {
    SuspectBlobDs3Targets []SuspectBlobDs3Target
}

func (suspectBlobDs3TargetList *SuspectBlobDs3TargetList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "SuspectBlobDs3Target":
            var model SuspectBlobDs3Target
            model.parse(&child, aggErr)
            suspectBlobDs3TargetList.SuspectBlobDs3Targets = append(suspectBlobDs3TargetList.SuspectBlobDs3Targets, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobDs3TargetList.", child.XMLName.Local)
        }
    }
}

type SuspectBlobPoolList struct {
    SuspectBlobPools []SuspectBlobPool
}

func (suspectBlobPoolList *SuspectBlobPoolList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "SuspectBlobPool":
            var model SuspectBlobPool
            model.parse(&child, aggErr)
            suspectBlobPoolList.SuspectBlobPools = append(suspectBlobPoolList.SuspectBlobPools, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobPoolList.", child.XMLName.Local)
        }
    }
}

type SuspectBlobS3TargetList struct {
    SuspectBlobS3Targets []SuspectBlobS3Target
}

func (suspectBlobS3TargetList *SuspectBlobS3TargetList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "SuspectBlobS3Target":
            var model SuspectBlobS3Target
            model.parse(&child, aggErr)
            suspectBlobS3TargetList.SuspectBlobS3Targets = append(suspectBlobS3TargetList.SuspectBlobS3Targets, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobS3TargetList.", child.XMLName.Local)
        }
    }
}

type SuspectBlobTapeList struct {
    SuspectBlobTapes []SuspectBlobTape
}

func (suspectBlobTapeList *SuspectBlobTapeList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "SuspectBlobTape":
            var model SuspectBlobTape
            model.parse(&child, aggErr)
            suspectBlobTapeList.SuspectBlobTapes = append(suspectBlobTapeList.SuspectBlobTapes, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SuspectBlobTapeList.", child.XMLName.Local)
        }
    }
}

type S3ObjectList struct {
    S3Objects []S3Object
}

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

type GroupMemberList struct {
    GroupMembers []GroupMember
}

func (groupMemberList *GroupMemberList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "GroupMember":
            var model GroupMember
            model.parse(&child, aggErr)
            groupMemberList.GroupMembers = append(groupMemberList.GroupMembers, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing GroupMemberList.", child.XMLName.Local)
        }
    }
}

type GroupList struct {
    Groups []Group
}

func (groupList *GroupList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Group":
            var model Group
            model.parse(&child, aggErr)
            groupList.Groups = append(groupList.Groups, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing GroupList.", child.XMLName.Local)
        }
    }
}

type ActiveJobList struct {
    ActiveJobs []ActiveJob
}

func (activeJobList *ActiveJobList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Job":
            var model ActiveJob
            model.parse(&child, aggErr)
            activeJobList.ActiveJobs = append(activeJobList.ActiveJobs, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing ActiveJobList.", child.XMLName.Local)
        }
    }
}

type CanceledJobList struct {
    CanceledJobs []CanceledJob
}

func (canceledJobList *CanceledJobList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CanceledJob":
            var model CanceledJob
            model.parse(&child, aggErr)
            canceledJobList.CanceledJobs = append(canceledJobList.CanceledJobs, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CanceledJobList.", child.XMLName.Local)
        }
    }
}

type CompletedJobList struct {
    CompletedJobs []CompletedJob
}

func (completedJobList *CompletedJobList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CompletedJob":
            var model CompletedJob
            model.parse(&child, aggErr)
            completedJobList.CompletedJobs = append(completedJobList.CompletedJobs, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing CompletedJobList.", child.XMLName.Local)
        }
    }
}

type NodeList struct {
    Nodes []Node
}

func (nodeList *NodeList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Node":
            var model Node
            model.parse(&child, aggErr)
            nodeList.Nodes = append(nodeList.Nodes, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing NodeList.", child.XMLName.Local)
        }
    }
}

type AzureTargetFailureNotificationRegistrationList struct {
    AzureTargetFailureNotificationRegistrations []AzureTargetFailureNotificationRegistration
}

func (azureTargetFailureNotificationRegistrationList *AzureTargetFailureNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AzureTargetFailureNotificationRegistration":
            var model AzureTargetFailureNotificationRegistration
            model.parse(&child, aggErr)
            azureTargetFailureNotificationRegistrationList.AzureTargetFailureNotificationRegistrations = append(azureTargetFailureNotificationRegistrationList.AzureTargetFailureNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureTargetFailureNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type BucketChangesNotificationRegistrationList struct {
    BucketChangesNotificationRegistrations []BucketChangesNotificationRegistration
}

func (bucketChangesNotificationRegistrationList *BucketChangesNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketChangesNotificationRegistration":
            var model BucketChangesNotificationRegistration
            model.parse(&child, aggErr)
            bucketChangesNotificationRegistrationList.BucketChangesNotificationRegistrations = append(bucketChangesNotificationRegistrationList.BucketChangesNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BucketChangesNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type BucketHistoryEventList struct {
    BucketHistoryEvents []BucketHistoryEvent
}

func (bucketHistoryEventList *BucketHistoryEventList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "BucketHistoryEvent":
            var model BucketHistoryEvent
            model.parse(&child, aggErr)
            bucketHistoryEventList.BucketHistoryEvents = append(bucketHistoryEventList.BucketHistoryEvents, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing BucketHistoryEventList.", child.XMLName.Local)
        }
    }
}

type Ds3TargetFailureNotificationRegistrationList struct {
    Ds3TargetFailureNotificationRegistrations []Ds3TargetFailureNotificationRegistration
}

func (ds3TargetFailureNotificationRegistrationList *Ds3TargetFailureNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Ds3TargetFailureNotificationRegistration":
            var model Ds3TargetFailureNotificationRegistration
            model.parse(&child, aggErr)
            ds3TargetFailureNotificationRegistrationList.Ds3TargetFailureNotificationRegistrations = append(ds3TargetFailureNotificationRegistrationList.Ds3TargetFailureNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3TargetFailureNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type JobCompletedNotificationRegistrationList struct {
    JobCompletedNotificationRegistrations []JobCompletedNotificationRegistration
}

func (jobCompletedNotificationRegistrationList *JobCompletedNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "JobCompletedNotificationRegistration":
            var model JobCompletedNotificationRegistration
            model.parse(&child, aggErr)
            jobCompletedNotificationRegistrationList.JobCompletedNotificationRegistrations = append(jobCompletedNotificationRegistrationList.JobCompletedNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing JobCompletedNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type JobCreatedNotificationRegistrationList struct {
    JobCreatedNotificationRegistrations []JobCreatedNotificationRegistration
}

func (jobCreatedNotificationRegistrationList *JobCreatedNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "JobCreatedNotificationRegistration":
            var model JobCreatedNotificationRegistration
            model.parse(&child, aggErr)
            jobCreatedNotificationRegistrationList.JobCreatedNotificationRegistrations = append(jobCreatedNotificationRegistrationList.JobCreatedNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing JobCreatedNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type JobCreationFailedNotificationRegistrationList struct {
    JobCreationFailedNotificationRegistrations []JobCreationFailedNotificationRegistration
}

func (jobCreationFailedNotificationRegistrationList *JobCreationFailedNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "JobCreationFailedNotificationRegistration":
            var model JobCreationFailedNotificationRegistration
            model.parse(&child, aggErr)
            jobCreationFailedNotificationRegistrationList.JobCreationFailedNotificationRegistrations = append(jobCreationFailedNotificationRegistrationList.JobCreationFailedNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing JobCreationFailedNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type S3ObjectCachedNotificationRegistrationList struct {
    S3ObjectCachedNotificationRegistrations []S3ObjectCachedNotificationRegistration
}

func (s3ObjectCachedNotificationRegistrationList *S3ObjectCachedNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "S3ObjectCachedNotificationRegistration":
            var model S3ObjectCachedNotificationRegistration
            model.parse(&child, aggErr)
            s3ObjectCachedNotificationRegistrationList.S3ObjectCachedNotificationRegistrations = append(s3ObjectCachedNotificationRegistrationList.S3ObjectCachedNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3ObjectCachedNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type S3ObjectLostNotificationRegistrationList struct {
    S3ObjectLostNotificationRegistrations []S3ObjectLostNotificationRegistration
}

func (s3ObjectLostNotificationRegistrationList *S3ObjectLostNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "S3ObjectLostNotificationRegistration":
            var model S3ObjectLostNotificationRegistration
            model.parse(&child, aggErr)
            s3ObjectLostNotificationRegistrationList.S3ObjectLostNotificationRegistrations = append(s3ObjectLostNotificationRegistrationList.S3ObjectLostNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3ObjectLostNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type S3ObjectPersistedNotificationRegistrationList struct {
    S3ObjectPersistedNotificationRegistrations []S3ObjectPersistedNotificationRegistration
}

func (s3ObjectPersistedNotificationRegistrationList *S3ObjectPersistedNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "S3ObjectPersistedNotificationRegistration":
            var model S3ObjectPersistedNotificationRegistration
            model.parse(&child, aggErr)
            s3ObjectPersistedNotificationRegistrationList.S3ObjectPersistedNotificationRegistrations = append(s3ObjectPersistedNotificationRegistrationList.S3ObjectPersistedNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3ObjectPersistedNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type PoolFailureNotificationRegistrationList struct {
    PoolFailureNotificationRegistrations []PoolFailureNotificationRegistration
}

func (poolFailureNotificationRegistrationList *PoolFailureNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "PoolFailureNotificationRegistration":
            var model PoolFailureNotificationRegistration
            model.parse(&child, aggErr)
            poolFailureNotificationRegistrationList.PoolFailureNotificationRegistrations = append(poolFailureNotificationRegistrationList.PoolFailureNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing PoolFailureNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type S3TargetFailureNotificationRegistrationList struct {
    S3TargetFailureNotificationRegistrations []S3TargetFailureNotificationRegistration
}

func (s3TargetFailureNotificationRegistrationList *S3TargetFailureNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "S3TargetFailureNotificationRegistration":
            var model S3TargetFailureNotificationRegistration
            model.parse(&child, aggErr)
            s3TargetFailureNotificationRegistrationList.S3TargetFailureNotificationRegistrations = append(s3TargetFailureNotificationRegistrationList.S3TargetFailureNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3TargetFailureNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type StorageDomainFailureNotificationRegistrationList struct {
    StorageDomainFailureNotificationRegistrations []StorageDomainFailureNotificationRegistration
}

func (storageDomainFailureNotificationRegistrationList *StorageDomainFailureNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "StorageDomainFailureNotificationRegistration":
            var model StorageDomainFailureNotificationRegistration
            model.parse(&child, aggErr)
            storageDomainFailureNotificationRegistrationList.StorageDomainFailureNotificationRegistrations = append(storageDomainFailureNotificationRegistrationList.StorageDomainFailureNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing StorageDomainFailureNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type SystemFailureNotificationRegistrationList struct {
    SystemFailureNotificationRegistrations []SystemFailureNotificationRegistration
}

func (systemFailureNotificationRegistrationList *SystemFailureNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "SystemFailureNotificationRegistration":
            var model SystemFailureNotificationRegistration
            model.parse(&child, aggErr)
            systemFailureNotificationRegistrationList.SystemFailureNotificationRegistrations = append(systemFailureNotificationRegistrationList.SystemFailureNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SystemFailureNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type TapeFailureNotificationRegistrationList struct {
    TapeFailureNotificationRegistrations []TapeFailureNotificationRegistration
}

func (tapeFailureNotificationRegistrationList *TapeFailureNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "TapeFailureNotificationRegistration":
            var model TapeFailureNotificationRegistration
            model.parse(&child, aggErr)
            tapeFailureNotificationRegistrationList.TapeFailureNotificationRegistrations = append(tapeFailureNotificationRegistrationList.TapeFailureNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeFailureNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type TapePartitionFailureNotificationRegistrationList struct {
    TapePartitionFailureNotificationRegistrations []TapePartitionFailureNotificationRegistration
}

func (tapePartitionFailureNotificationRegistrationList *TapePartitionFailureNotificationRegistrationList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "TapePartitionFailureNotificationRegistration":
            var model TapePartitionFailureNotificationRegistration
            model.parse(&child, aggErr)
            tapePartitionFailureNotificationRegistrationList.TapePartitionFailureNotificationRegistrations = append(tapePartitionFailureNotificationRegistrationList.TapePartitionFailureNotificationRegistrations, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapePartitionFailureNotificationRegistrationList.", child.XMLName.Local)
        }
    }
}

type DetailedS3ObjectList struct {
    DetailedS3Objects []DetailedS3Object
}

func (detailedS3ObjectList *DetailedS3ObjectList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Object":
            var model DetailedS3Object
            model.parse(&child, aggErr)
            detailedS3ObjectList.DetailedS3Objects = append(detailedS3ObjectList.DetailedS3Objects, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DetailedS3ObjectList.", child.XMLName.Local)
        }
    }
}

type PoolFailureList struct {
    PoolFailures []PoolFailure
}

func (poolFailureList *PoolFailureList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "PoolFailure":
            var model PoolFailure
            model.parse(&child, aggErr)
            poolFailureList.PoolFailures = append(poolFailureList.PoolFailures, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing PoolFailureList.", child.XMLName.Local)
        }
    }
}

type PoolPartitionList struct {
    PoolPartitions []PoolPartition
}

func (poolPartitionList *PoolPartitionList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "PoolPartition":
            var model PoolPartition
            model.parse(&child, aggErr)
            poolPartitionList.PoolPartitions = append(poolPartitionList.PoolPartitions, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing PoolPartitionList.", child.XMLName.Local)
        }
    }
}

type PoolList struct {
    Pools []Pool
}

func (poolList *PoolList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Pool":
            var model Pool
            model.parse(&child, aggErr)
            poolList.Pools = append(poolList.Pools, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing PoolList.", child.XMLName.Local)
        }
    }
}

type StorageDomainFailureList struct {
    StorageDomainFailures []StorageDomainFailure
}

func (storageDomainFailureList *StorageDomainFailureList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "StorageDomainFailure":
            var model StorageDomainFailure
            model.parse(&child, aggErr)
            storageDomainFailureList.StorageDomainFailures = append(storageDomainFailureList.StorageDomainFailures, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing StorageDomainFailureList.", child.XMLName.Local)
        }
    }
}

type StorageDomainMemberList struct {
    StorageDomainMembers []StorageDomainMember
}

func (storageDomainMemberList *StorageDomainMemberList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "StorageDomainMember":
            var model StorageDomainMember
            model.parse(&child, aggErr)
            storageDomainMemberList.StorageDomainMembers = append(storageDomainMemberList.StorageDomainMembers, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing StorageDomainMemberList.", child.XMLName.Local)
        }
    }
}

type StorageDomainList struct {
    StorageDomains []StorageDomain
}

func (storageDomainList *StorageDomainList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "StorageDomain":
            var model StorageDomain
            model.parse(&child, aggErr)
            storageDomainList.StorageDomains = append(storageDomainList.StorageDomains, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing StorageDomainList.", child.XMLName.Local)
        }
    }
}

type FeatureKeyList struct {
    FeatureKeys []FeatureKey
}

func (featureKeyList *FeatureKeyList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "FeatureKey":
            var model FeatureKey
            model.parse(&child, aggErr)
            featureKeyList.FeatureKeys = append(featureKeyList.FeatureKeys, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing FeatureKeyList.", child.XMLName.Local)
        }
    }
}

type SystemFailureList struct {
    SystemFailures []SystemFailure
}

func (systemFailureList *SystemFailureList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "SystemFailure":
            var model SystemFailure
            model.parse(&child, aggErr)
            systemFailureList.SystemFailures = append(systemFailureList.SystemFailures, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SystemFailureList.", child.XMLName.Local)
        }
    }
}

type TapeDensityDirectiveList struct {
    TapeDensityDirectives []TapeDensityDirective
}

func (tapeDensityDirectiveList *TapeDensityDirectiveList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "TapeDensityDirective":
            var model TapeDensityDirective
            model.parse(&child, aggErr)
            tapeDensityDirectiveList.TapeDensityDirectives = append(tapeDensityDirectiveList.TapeDensityDirectives, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeDensityDirectiveList.", child.XMLName.Local)
        }
    }
}

type TapeDriveList struct {
    TapeDrives []TapeDrive
}

func (tapeDriveList *TapeDriveList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "TapeDrive":
            var model TapeDrive
            model.parse(&child, aggErr)
            tapeDriveList.TapeDrives = append(tapeDriveList.TapeDrives, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeDriveList.", child.XMLName.Local)
        }
    }
}

type DetailedTapeFailureList struct {
    DetailedTapeFailures []DetailedTapeFailure
}

func (detailedTapeFailureList *DetailedTapeFailureList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "TapeFailure":
            var model DetailedTapeFailure
            model.parse(&child, aggErr)
            detailedTapeFailureList.DetailedTapeFailures = append(detailedTapeFailureList.DetailedTapeFailures, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing DetailedTapeFailureList.", child.XMLName.Local)
        }
    }
}

type TapeLibraryList struct {
    TapeLibraries []TapeLibrary
}

func (tapeLibraryList *TapeLibraryList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "TapeLibrary":
            var model TapeLibrary
            model.parse(&child, aggErr)
            tapeLibraryList.TapeLibraries = append(tapeLibraryList.TapeLibraries, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeLibraryList.", child.XMLName.Local)
        }
    }
}

type TapePartitionFailureList struct {
    TapePartitionFailures []TapePartitionFailure
}

func (tapePartitionFailureList *TapePartitionFailureList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "TapePartitionFailure":
            var model TapePartitionFailure
            model.parse(&child, aggErr)
            tapePartitionFailureList.TapePartitionFailures = append(tapePartitionFailureList.TapePartitionFailures, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapePartitionFailureList.", child.XMLName.Local)
        }
    }
}

type TapePartitionList struct {
    TapePartitions []TapePartition
}

func (tapePartitionList *TapePartitionList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "TapePartition":
            var model TapePartition
            model.parse(&child, aggErr)
            tapePartitionList.TapePartitions = append(tapePartitionList.TapePartitions, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapePartitionList.", child.XMLName.Local)
        }
    }
}

type NamedDetailedTapePartitionList struct {
    NamedDetailedTapePartitions []NamedDetailedTapePartition
}

func (namedDetailedTapePartitionList *NamedDetailedTapePartitionList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "TapePartition":
            var model NamedDetailedTapePartition
            model.parse(&child, aggErr)
            namedDetailedTapePartitionList.NamedDetailedTapePartitions = append(namedDetailedTapePartitionList.NamedDetailedTapePartitions, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing NamedDetailedTapePartitionList.", child.XMLName.Local)
        }
    }
}

type TapeList struct {
    Tapes []Tape
}

func (tapeList *TapeList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Tape":
            var model Tape
            model.parse(&child, aggErr)
            tapeList.Tapes = append(tapeList.Tapes, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing TapeList.", child.XMLName.Local)
        }
    }
}

type AzureTargetBucketNameList struct {
    AzureTargetBucketNames []AzureTargetBucketName
}

func (azureTargetBucketNameList *AzureTargetBucketNameList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AzureTargetBucketName":
            var model AzureTargetBucketName
            model.parse(&child, aggErr)
            azureTargetBucketNameList.AzureTargetBucketNames = append(azureTargetBucketNameList.AzureTargetBucketNames, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureTargetBucketNameList.", child.XMLName.Local)
        }
    }
}

type AzureTargetFailureList struct {
    AzureTargetFailures []AzureTargetFailure
}

func (azureTargetFailureList *AzureTargetFailureList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AzureTargetFailure":
            var model AzureTargetFailure
            model.parse(&child, aggErr)
            azureTargetFailureList.AzureTargetFailures = append(azureTargetFailureList.AzureTargetFailures, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureTargetFailureList.", child.XMLName.Local)
        }
    }
}

type AzureTargetReadPreferenceList struct {
    AzureTargetReadPreferences []AzureTargetReadPreference
}

func (azureTargetReadPreferenceList *AzureTargetReadPreferenceList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AzureTargetReadPreference":
            var model AzureTargetReadPreference
            model.parse(&child, aggErr)
            azureTargetReadPreferenceList.AzureTargetReadPreferences = append(azureTargetReadPreferenceList.AzureTargetReadPreferences, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureTargetReadPreferenceList.", child.XMLName.Local)
        }
    }
}

type AzureTargetList struct {
    AzureTargets []AzureTarget
}

func (azureTargetList *AzureTargetList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "AzureTarget":
            var model AzureTarget
            model.parse(&child, aggErr)
            azureTargetList.AzureTargets = append(azureTargetList.AzureTargets, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing AzureTargetList.", child.XMLName.Local)
        }
    }
}

type Ds3TargetFailureList struct {
    Ds3TargetFailures []Ds3TargetFailure
}

func (ds3TargetFailureList *Ds3TargetFailureList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Ds3TargetFailure":
            var model Ds3TargetFailure
            model.parse(&child, aggErr)
            ds3TargetFailureList.Ds3TargetFailures = append(ds3TargetFailureList.Ds3TargetFailures, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3TargetFailureList.", child.XMLName.Local)
        }
    }
}

type Ds3TargetReadPreferenceList struct {
    Ds3TargetReadPreferences []Ds3TargetReadPreference
}

func (ds3TargetReadPreferenceList *Ds3TargetReadPreferenceList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Ds3TargetReadPreference":
            var model Ds3TargetReadPreference
            model.parse(&child, aggErr)
            ds3TargetReadPreferenceList.Ds3TargetReadPreferences = append(ds3TargetReadPreferenceList.Ds3TargetReadPreferences, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3TargetReadPreferenceList.", child.XMLName.Local)
        }
    }
}

type Ds3TargetList struct {
    Ds3Targets []Ds3Target
}

func (ds3TargetList *Ds3TargetList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "Ds3Target":
            var model Ds3Target
            model.parse(&child, aggErr)
            ds3TargetList.Ds3Targets = append(ds3TargetList.Ds3Targets, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing Ds3TargetList.", child.XMLName.Local)
        }
    }
}

type S3TargetBucketNameList struct {
    S3TargetBucketNames []S3TargetBucketName
}

func (s3TargetBucketNameList *S3TargetBucketNameList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "S3TargetBucketName":
            var model S3TargetBucketName
            model.parse(&child, aggErr)
            s3TargetBucketNameList.S3TargetBucketNames = append(s3TargetBucketNameList.S3TargetBucketNames, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3TargetBucketNameList.", child.XMLName.Local)
        }
    }
}

type S3TargetFailureList struct {
    S3TargetFailures []S3TargetFailure
}

func (s3TargetFailureList *S3TargetFailureList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "S3TargetFailure":
            var model S3TargetFailure
            model.parse(&child, aggErr)
            s3TargetFailureList.S3TargetFailures = append(s3TargetFailureList.S3TargetFailures, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3TargetFailureList.", child.XMLName.Local)
        }
    }
}

type S3TargetReadPreferenceList struct {
    S3TargetReadPreferences []S3TargetReadPreference
}

func (s3TargetReadPreferenceList *S3TargetReadPreferenceList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "S3TargetReadPreference":
            var model S3TargetReadPreference
            model.parse(&child, aggErr)
            s3TargetReadPreferenceList.S3TargetReadPreferences = append(s3TargetReadPreferenceList.S3TargetReadPreferences, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3TargetReadPreferenceList.", child.XMLName.Local)
        }
    }
}

type S3TargetList struct {
    S3Targets []S3Target
}

func (s3TargetList *S3TargetList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "S3Target":
            var model S3Target
            model.parse(&child, aggErr)
            s3TargetList.S3Targets = append(s3TargetList.S3Targets, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3TargetList.", child.XMLName.Local)
        }
    }
}

type SpectraUserList struct {
    SpectraUsers []SpectraUser
}

func (spectraUserList *SpectraUserList) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "User":
            var model SpectraUser
            model.parse(&child, aggErr)
            spectraUserList.SpectraUsers = append(spectraUserList.SpectraUsers, model)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing SpectraUserList.", child.XMLName.Local)
        }
    }
}

