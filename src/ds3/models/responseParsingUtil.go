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

package models

import "strconv"

// Contains utils used by model parsers to parse response payloads.

// Interface defined for spectra defined enums.
// Used for generic parsing of enums.
type Ds3Enum interface {
    UnmarshalText(text []byte) error
}

// Parses a byte slice into an enum value. Assumes that a valid enum value is
// contained within the byte slice.
// Used in response payload parsing to convert xml node content into expected types.
func parseEnum(content []byte, param Ds3Enum, aggErr *AggregateError) {
    err := param.UnmarshalText(content)
    if err != nil {
        aggErr.Append(err)
    }
}

// Parses a byte slice into an enum value. An empty slice is converted into a nil value.
// Used in response payload parsing to convert xml node content into expected types.
func parseNullableEnum(content []byte, param Ds3Enum, aggErr *AggregateError) {
    if len(content) > 0 {
        parseEnum(content, param, aggErr)
    }
}

// Converts a byte slice into a string value.
// Used in response payload parsing to convert xml node content into expected types.
func parseString(content []byte) string {
    return string(content)
}

// Converts a byte slice into a string value. An empty slice is converted into a nil value.
// Used in response payload parsing to convert xml node content into expected types.
func parseNullableString(content []byte) *string {
    if len(content) == 0 {
        return nil
    }
    result := parseString(content)
    return &result
}

// Converts a byte slice into an int value. Assumes that a valid int value is
// contained within the byte slice.
// Used in response payload parsing to convert xml node content into expected types.
func parseInt(content []byte, aggErr *AggregateError) int {
    result, err :=  strconv.Atoi(string(content))
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

// Converts a byte slice into an int value. An empty slice is converted into a nil value.
// Used in response payload parsing to convert xml node content into expected types.
func parseNullableInt(content []byte, aggErr *AggregateError) *int {
    if len(content) == 0 {
        return nil
    }
    result := parseInt(content, aggErr)
    return &result
}

// Converts a byte slice into an int64 value. Assumes that a valid int64 value is
// contained within the byte slice.
// Used in response payload parsing to convert xml node content into expected types.
func parseInt64(content []byte, aggErr *AggregateError) int64 {
    result, err := strconv.ParseInt(string(content), 10, 64)
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

// Converts a byte slice into an int64 value. An empty slice is converted into a nil value.
// Used in response payload parsing to convert xml node content into expected types.
func parseNullableInt64(content []byte, aggErr *AggregateError) *int64 {
    if len(content) == 0 {
        return nil
    }
    result := parseInt64(content, aggErr)
    return &result
}

// Converts a byte slice into a boolean value. Assumes that a valid boolean value is
// contained within the byte slice.
// Used in response payload parsing to convert xml node content into expected types.
func parseBool(content []byte, aggErr *AggregateError) bool {
    result, err := strconv.ParseBool(string(content))
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

// Converts a byte slice into a boolean value. An empty slice is converted into a nil value.
// Used in response payload parsing to convert xml node content into expected types.
func parseNullableBool(content []byte, aggErr *AggregateError) *bool {
    if len(content) == 0 {
        return nil
    }
    result := parseBool(content, aggErr)
    return &result
}

// Parses an enum value from a string and expects a valid value to exist.
// Used for parsing response payload attributes into expected types.
func parseEnumFromString(content string, param Ds3Enum, aggErr *AggregateError) {
    err := param.UnmarshalText([]byte(content))
    if err != nil {
        aggErr.Append(err)
    }
}

// Parse nullable enum from string, where empty content represents a nil enum result.
// Used for parsing response payload attributes into expected types.
func parseNullableEnumFromString(content string, param Ds3Enum, aggErr *AggregateError) {
    if len(content) > 0 {
        parseEnumFromString(content, param, aggErr)
    }
}

// Parses a nullable string, where empty content represents a nil string result.
// Used for parsing response payload attributes into expected types.
func parseNullableStringFromString(content string) *string {
    if len(content) == 0 {
        return nil
    }
    result := content
    return &result
}

// Parses an int value from a string and expects a value to exist.
// Used for parsing response payload attributes into expected types.
func parseIntFromString(content string, aggErr *AggregateError) int {
    result, err :=  strconv.Atoi(content)
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

// Parses a nullable int from string, where empty content represents a nil int result.
// Used for parsing response payload attributes into expected types.
func parseNullableIntFromString(content string, aggErr *AggregateError) *int {
    if len(content) == 0 {
        return nil
    }
    result := parseIntFromString(content, aggErr)
    return &result
}

// Parses an int64 value from a string and expects a value to exist.
// Used for parsing response payload attributes into expected types.
func parseInt64FromString(content string, aggErr *AggregateError) int64 {
    result, err := strconv.ParseInt(content, 10, 64)
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

// Parses a nullable int64 from string, where an empty string represents a nil int64 result.
// Used for parsing response payload attributes into expected types.
func parseNullableInt64FromString(content string, aggErr *AggregateError) *int64 {
    if len(content) == 0 {
        return nil
    }
    result := parseInt64FromString(content, aggErr)
    return &result
}

// Parses an float64 value from a string and expects a value to exist.
// Used for parsing response payload attributes into expected types.
func parseFloat64FromString(content string, aggErr *AggregateError) float64 {
    result, err := strconv.ParseFloat(content, 64)
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

// Parses a nullable float64 from string, where an empty string represents a nil float64 result.
// Used for parsing response payload attributes into expected types.
func parseNullableFloat64FromString(content string, aggErr *AggregateError) *float64 {
    if len(content) == 0 {
        return nil
    }
    result := parseFloat64FromString(content, aggErr)
    return &result
}

// Parses a boolean value from a string and expects a valid boolean value to exist.
// Used for parsing response payload attributes into expected types.
func parseBoolFromString(content string, aggErr *AggregateError) bool {
    result, err := strconv.ParseBool(content)
    if err != nil {
        aggErr.Append(err)
    }
    return result
}

// Parses a nullable boolean value from a string, where an empty string represents a nil boolean result.
// Used for parsing response payload attributes into expected types.
func parseNullableBoolFromString(content string, aggErr *AggregateError) *bool {
    if len(content) == 0 {
        return nil
    }
    result := parseBoolFromString(content, aggErr)
    return &result
}