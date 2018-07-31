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
    "testing"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_utils/ds3Testing"
)

func TestBuildAuthHeaderValue(t *testing.T) {
    expected := "AWS access:1JbBTxv5KRFKbvju7w27c2J6bKk="

    fields := signatureFields{
        Verb:"PUT",
        Date:"Tue, 31 Jul 2018 17:23:50 +0000",
        Path:"abc/1+2+3/e f g/aÀˠϠૐ༺ᎀ",
    }

    credentials := Credentials{
        AccessId:"access",
        Key:"key",
    }

    result := fields.BuildAuthHeaderValue(&credentials)

    ds3Testing.AssertString(t, "authorization header", expected, result)
}
