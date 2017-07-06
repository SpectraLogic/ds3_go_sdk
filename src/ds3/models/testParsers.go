package models

import (
    "strconv"
    "log"
    "fmt"
    "reflect"
)

// TODO move to separate file

type AggregateError struct {
    Errors []error
}

func (aggregateError *AggregateError) Error() string {
    msg := fmt.Sprintf("Multiple errors occured: %d\n", len(aggregateError.Errors))

    for i, err := range aggregateError.Errors {
        msg += fmt.Sprintf("%d) %s\n", i, err.Error())
    }

    return msg
}

// Returns the aggregate error if at least one error exists,
// else returns nil
func (aggregateError *AggregateError) GetErrors() error {
    if len (aggregateError.Errors) == 0 {
        return nil
    }
    return aggregateError
}

func (aggregateError *AggregateError) Append(err error) {
    if err != nil {
        aggregateError.Errors = append(aggregateError.Errors, err)
    }
}

//TODO test parsers- hand implemented

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

func (job *Job) parse(node *XmlNode, aggErr *AggregateError) {
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
            var jobNode JobNode
            var list modelSlice = parseDs3ObjectList("Node", child.Children, &jobNode, aggErr)
            for _, n := range list {
                if v, ok := n.(*JobNode); ok {
                    job.Nodes = append(job.Nodes, *v)
                }
            }
        }
    }
}

type modelSlice []modelParser

// TODO Convert to generic function using reflection that works for any modelParser type
func parseDs3ObjectList(tagName string, xmlNodes []XmlNode, model modelParser, aggErr *AggregateError) modelSlice {
    var list modelSlice

    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            v := reflect.ValueOf(model)
            curResult := v.Interface().(modelParser)
            curResult.parse(&curXmlNode, aggErr)
            list = append(list, curResult)
        } else {
            // TODO possibly change from logging to appending to aggregate error
            log.Printf("WARNING: Discovered unexpected tag '%s' when expected tag '%s'.\n", curXmlNode.XMLName.Local, tagName)
        }
    }
    return list
}

// TODO Convert to generic function using reflection that works for any modelParser type
func parseJobNodeList(tagName string, xmlNodes []XmlNode, aggErr *AggregateError) []JobNode {
    var result []JobNode
    for _, curXmlNode := range xmlNodes {
        if curXmlNode.XMLName.Local == tagName {
            var curResult JobNode
            curResult.parse(&curXmlNode, aggErr)
            result = append(result, curResult)
        } else {
            // TODO possibly change from logging to appending to aggregate error
            log.Printf("WARNING: Discovered unexpected tag '%s' when expected tag '%s'.\n", curXmlNode.XMLName.Local, tagName)
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

// Interface defined for spectra defined enums
// Used for generic parsing of enums
// within: parseEnum and parseNullableEnum
type Ds3Enum interface {
    UnmarshalText(text []byte) error
}

func parseEnum(content []byte, param Ds3Enum, aggErr *AggregateError) {
    err := param.UnmarshalText(content)
    if err != nil {
        aggErr.Append(err)
    }
}

func parseNullableEnum(content []byte, param Ds3Enum, aggErr *AggregateError) {
    if len(content) > 0 {
        parseEnum(content, param, aggErr)
    }

}

// Parses a string value
func parseString(content []byte) string {
    return string(content)
}

// Parses a string, where no bytes is converted to nil
func parseNullableString(content []byte) *string {
    if len(content) == 0 {
        return nil
    }
    result := parseString(content)
    return &result
}

// Parses an int64 value and expects a value to exist
func parseInt(content []byte, aggErr *AggregateError) int {
    result, err :=  strconv.Atoi(string(content))
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

func parseNullableInt(content []byte, aggErr *AggregateError) *int {
    if len(content) == 0 {
        return nil
    }
    result := parseInt(content, aggErr)
    return &result
}

// Parses an int64 value and expects a value to exist
func parseInt64(content []byte, aggErr *AggregateError) int64 {
    result, err := strconv.ParseInt(string(content), 10, 64)
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

func parseNullableInt64(content []byte, aggErr *AggregateError) *int64 {
    if len(content) == 0 {
        return nil
    }
    result := parseInt64(content, aggErr)
    return &result
}

// Parses a boolean value and expects a value
func parseBool(content []byte, aggErr *AggregateError) bool {
    result, err := strconv.ParseBool(string(content))
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

func parseNullableBool(content []byte, aggErr *AggregateError) *bool {
    if len(content) == 0 {
        return nil
    }
    result := parseBool(content, aggErr)
    return &result
}

// TODO ----------------------------- parsing attributes from strings

func parseEnumFromString(content string, param Ds3Enum, aggErr *AggregateError) {
    err := param.UnmarshalText([]byte(content))
    if err != nil {
        aggErr.Append(err)
    }
}

func parseNullableEnumFromString(content string, param Ds3Enum, aggErr *AggregateError) {
    if len(content) > 0 {
        parseEnumFromString(content, param, aggErr)
    }
}

// Parses a string, where no bytes is converted to nil
func parseNullableStringFromString(content string) *string {
    if len(content) == 0 {
        return nil
    }
    result := content
    return &result
}

// Parses an int value and expects a value to exist
func parseIntFromString(content string, aggErr *AggregateError) int {
    result, err :=  strconv.Atoi(content)
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

func parseNullableIntFromString(content string, aggErr *AggregateError) *int {
    if len(content) == 0 {
        return nil
    }
    result := parseIntFromString(content, aggErr)
    return &result
}

// Parses an int64 value and expects a value to exist
func parseInt64FromString(content string, aggErr *AggregateError) int64 {
    result, err := strconv.ParseInt(content, 10, 64)
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

func parseNullableInt64FromString(content string, aggErr *AggregateError) *int64 {
    if len(content) == 0 {
        return nil
    }
    result := parseInt64FromString(content, aggErr)
    return &result
}

// Parses a boolean value and expects a value
func parseBoolFromString(content string, aggErr *AggregateError) bool {
    result, err := strconv.ParseBool(content)
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

func parseNullableBoolFromString(content string, aggErr *AggregateError) *bool {
    if len(content) == 0 {
        return nil
    }
    result := parseBoolFromString(content, aggErr)
    return &result
}
