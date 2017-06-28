package models

import (
    "strconv"
)

//TODO test parser- hand implemented

func (tape *Tape) parse(node *XmlNode) error {
    // No attributes to parse for Tape object

    // Parse children nodes
    var err error
    for _, child := range node.Children {
        switch child.XMLName.Local {
        case "AssignedToStorageDomain":
            if tape.AssignedToStorageDomain, err = parseBool(child.Content); err != nil {
                return err
            }
        case "AvailableRawCapacity":
            var err error
            if tape.AvailableRawCapacity, err = parseNullableInt(child.Content); err != nil {
                return err
            }
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
            if tape.FullOfData, err = parseBool(child.Content); err != nil {
                return err
            }
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
            if err := parseNullableEnum(child.Content, tape.PreviousState); err != nil {
                return err
            }
        case "SerialNumber":
            tape.SerialNumber = parseNullableString(child.Content)
        case "State":
            if err := parseEnum(child.Content, &tape.State); err != nil {
                return err
            }
        case "StorageDomainId":
            tape.StorageDomainId = parseNullableString(child.Content)
        case "TakeOwnershipPending":
            if tape.TakeOwnershipPending, err = parseBool(child.Content); err != nil {
                return err
            }
        case "TotalRawCapacity":
            if tape.TotalRawCapacity, err = parseNullableInt(child.Content); err != nil {
                return err
            }
        case "Type":
            if err := parseEnum(child.Content, &tape.Type); err != nil {
                return err
            }
        case "VerifyPending":
            if err := parseNullableEnum(child.Content, tape.VerifyPending); err != nil {
                return err
            }
        case "WriteProtected":
            if tape.WriteProtected, err = parseBool(child.Content); err != nil {
                return err
            }
        }
    }

    return nil
}

// Interface defined for spectra defined enums
// Used for generic parsing of enums
// within: parseEnum and parseNullableEnum
type Ds3Enum interface {
    UnmarshalText(text []byte) error
}

func parseEnum(content []byte, param Ds3Enum) error {
    return param.UnmarshalText(content)
}

func parseNullableEnum(content []byte, param Ds3Enum) error {
    if len(content) == 0 {
        return nil
    }
    return parseEnum(content, param)
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
func parseInt(content []byte) (int64, error) {
    return strconv.ParseInt(string(content), 10, 64)
}

func parseNullableInt(content []byte) (*int64, error) {
    if len(content) == 0 {
        return nil, nil
    }
    result, err := parseInt(content)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Parses a boolean value and expects a value
func parseBool(content []byte) (bool, error) {
    return strconv.ParseBool(string(content))
}

func parseNullableBool(content []byte) (*bool, error) {
    if len(content) == 0 {
        return nil, nil
    }
    result, err := parseBool(content)
    if err != nil {
        return nil, err
    }
    return &result, nil
}