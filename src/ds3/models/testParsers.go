package models

//TODO delete file
/*
import (
    "log"
)

func (jobList *JobList) parse(node *XmlNode, aggErr *AggregateError) {
    // No attributes to parse

    // Parse children nodes
    for _, child := range node.Children {
        switch child.XMLName.Local {
        case "Job":
            var job Job
            job.parse(&child, aggErr)
            jobList.Jobs = append(jobList.Jobs, job)
        }
    }
}

func (job *Job) parse(node *XmlNode, aggErr *AggregateError) {g
    // Parse attributes
    for _, attr := range node.Attrs {
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
            job.EntirelyInCache = parseBoolFromString(attr.Value, aggErr)
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
        }
    }

    // Parse children nodes
    for _, child := range node.Children {
        switch child.XMLName.Local {
        case "Nodes":
            job.Nodes = parseJobNodeSlice("Node", child.Children, aggErr)
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
            log.Printf("WARNING: Discovered unexpected xml tag '%s' when expected tag '%s'.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return result
}

func (jobNode *JobNode) parse(node *XmlNode, aggErr *AggregateError) {
    // Parse attributes
    for _, attr := range node.Attrs {
        switch attr.Name.Local {
        case "EndPoint":
            jobNode.EndPoint = parseNullableStringFromString(attr.Value)
        case "HttpPort":
            jobNode.HttpPort = parseNullableIntFromString(attr.Value, aggErr)
        case "HttpsPort":
            jobNode.HttpsPort = parseNullableIntFromString(attr.Value, aggErr)
        case "Id":
            jobNode.Id = attr.Value
        }
    }

    // No children nodes
}

func (tapeList *TapeList) parse(node *XmlNode, aggErr *AggregateError) {
    // No attributes to parse

    // Parse children nodes
    for _, child := range node.Children {
        switch child.XMLName.Local {
        case "Tape":
            var tape Tape
            tape.parse(&child, aggErr)
            tapeList.Tapes = append(tapeList.Tapes, tape)
        }
    }
}

func (tape *Tape) parse(node *XmlNode, aggErr *AggregateError) {
    // No attributes to parse for Tape object

    // Parse children nodes
    for _, child := range node.Children {
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
            parseNullableEnum(child.Content, tape.PreviousState, aggErr)
        case "SerialNumber":
            tape.SerialNumber = parseNullableString(child.Content)
        case "State":
            parseEnum(child.Content, &tape.State, aggErr)
        case "StorageDomainId":
            tape.StorageDomainId = parseNullableString(child.Content)
        case "TakeOwnershipPending":
            tape.TakeOwnershipPending = parseBool(child.Content, aggErr)
        case "TotalRawCapacity":
            tape.TotalRawCapacity = parseNullableInt64(child.Content, aggErr)
        case "Type":
            parseEnum(child.Content, &tape.Type, aggErr)
        case "VerifyPending":
            parseNullableEnum(child.Content, tape.VerifyPending, aggErr)
        case "WriteProtected":
            tape.WriteProtected = parseBool(child.Content, aggErr)
        }
    }
}

*/