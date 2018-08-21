// Copyright 2014-2018 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package networking

import (
    "net/url"
    "bytes"
    "sort"
)

// Wrapper for url.Values to be used for encoding query parameters
// without the '=' sign when there is no corresponding value, ex:
// key, value -> key=value
// key, '' -> key
type CustomUrlValues struct {
    values url.Values
}

func NewCustomUrlValues() CustomUrlValues {
    return CustomUrlValues{ values:url.Values{} }
}

// Copied and modified from url.Values to remove '=' when there are no values to add
func (customValues CustomUrlValues) Encode() string {
    if customValues.values == nil {
        return ""
    }
    var buf bytes.Buffer
    keys := make([]string, 0, len(customValues.values))
    for key := range customValues.values {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    for _, key := range keys {
        values := customValues.values[key]
        prefix := url.QueryEscape(key)
        for _, value := range values {
            if buf.Len() > 0 {
                buf.WriteByte('&')
            }
            buf.WriteString(prefix)

            // If there is a value, then append
            if value != "" {
                buf.WriteString("=" + url.QueryEscape(value))
            }
        }
    }
    return buf.String()
}

func (customValues CustomUrlValues) Add(key string, value string) {
    customValues.values.Add(key, value)
}

func (customValues CustomUrlValues) Size() int {
    return len(customValues.values)
}