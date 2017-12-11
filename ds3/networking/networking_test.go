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

package networking

import (
    "testing"
    "net/url"
    "spectra/ds3_go_sdk/ds3_utils/ds3Testing"
)

func TestEncodeQueryParams(t *testing.T) {
    expected := "key%20space%20=value%20space%20&key%2Bplus=value%2Bplus&key%3Bsemicolon=value%3Bsemicolon"

    queryParams := &url.Values{}
    queryParams.Set("key space ", "value space ")
    queryParams.Set("key;semicolon", "value;semicolon")
    queryParams.Set("key+plus", "value+plus")

    result := encodeQueryParams(queryParams)

    ds3Testing.AssertString(t, "Encoded Query Params", expected, result)
}