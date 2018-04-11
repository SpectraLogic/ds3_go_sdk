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
    "strings"
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

func TestBuildingAuthorizationDigestWithMetadata(t *testing.T) {
	httpRequestBuilder := NewHttpRequestBuilder()

	endpointUrl, err := url.Parse("https://google.com")
	ds3Testing.AssertNilError(t, err)

	const gracie = "Gracie"
	const eskimo = "Eskimo"
	const shasta = "Shasta"
	const samoyed = "Samoyed"

	connectionInfo := ConnectionInfo{
		Endpoint:    endpointUrl,
		Credentials: &Credentials{AccessId: gracie, Key: eskimo},
		Proxy:       nil}

	httpRequestBuilder.
		WithHeader(AmazonMetadataPrefix + shasta, samoyed).
		WithHeader(AmazonMetadataPrefix + gracie, eskimo).
		Build(&connectionInfo)

	amazonHeaders := httpRequestBuilder.signatureFields.CanonicalizedAmzHeaders

	ds3Testing.AssertBool(t, "expected amazonHeader to have something in it", true, len(amazonHeaders) > 0)

	expected := AmazonMetadataPrefix + strings.ToLower(gracie) + ":" + eskimo + "\n" + AmazonMetadataPrefix + strings.ToLower(shasta) + ":" + samoyed + "\n"
	ds3Testing.AssertBool(t, "amazonHeader string isn't what we expected", true, strings.Compare(amazonHeaders, expected) == 0)
}

