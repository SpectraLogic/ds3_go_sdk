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
    "github.com/SpectraLogic/ds3_go_sdk/ds3_utils/ds3Testing"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
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
		WithHeader(models.AMZ_META_HEADER + shasta, samoyed).
		WithHeader(models.AMZ_META_HEADER + gracie, eskimo).
		Build(&connectionInfo)

	amazonHeaders := httpRequestBuilder.signatureFields.CanonicalizedAmzHeaders

	ds3Testing.AssertBool(t, "expected amazonHeader to have something in it", true, len(amazonHeaders) > 0)

	expected := models.AMZ_META_HEADER + strings.ToLower(gracie) + ":" + eskimo + "\n" + models.AMZ_META_HEADER + strings.ToLower(shasta) + ":" + samoyed + "\n"
	ds3Testing.AssertBool(t, "amazonHeader string isn't what we expected", true, strings.Compare(amazonHeaders, expected) == 0)
}

func TestBuildingAuthorizationDigestWithSignatureQueryParams(t *testing.T) {
	httpRequestBuilder := NewHttpRequestBuilder()

	endpointUrl, err := url.Parse("https://google.com")
	ds3Testing.AssertNilError(t, err)

	connectionInfo := ConnectionInfo{
		Endpoint:    endpointUrl,
		Credentials: &Credentials{AccessId: "access", Key: "key"},
		Proxy:       nil}

	httpRequestBuilder.
		WithHeader("versioning", "one").
		WithHeader("delete", "").
		WithHeader("uploads", "two").
		WithHeader("uploads", "three").
		WithHeader("doesnotexist", "four").
		WithHeader("VersionId", "version").
		Build(&connectionInfo)

	subResources := httpRequestBuilder.signatureFields.CanonicalizedSubResources

	ds3Testing.AssertBool(t, "expected sub resources to have something in it", true, len(subResources) > 0)

	expected := "?Delete&Uploads=two,three&Versionid=version&Versioning=one"
	ds3Testing.AssertString(t, "sub resources", expected, subResources)
}