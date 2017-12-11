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

package ds3Testing

import (
    "testing"
    "spectra/ds3_go_sdk/ds3/models"
    "reflect"
)

// Asserts if a specified bool pointer has the expected value. If not, Fatal.
func AssertNonNilBoolPtr(t *testing.T, label string, expected bool, actual *bool) {
    if actual == nil {
        t.Fatalf("Expected %s to be '%t' but was 'nil'.", label, expected)
    }
    AssertBool(t, label, expected, *actual)
}

// Asserts if a specified bool pointer is nil. If not, Fatal.
func AssertBoolPtrIsNil(t *testing.T, label string, actual *bool) {
    if actual != nil {
        t.Fatalf("Expected %s to be 'nil' but was '%t'.", label, *actual)
    }
}

// Asserts if a specified bool has the expected value. If not, Fatal.
func AssertBool(t *testing.T, label string, expected bool, actual bool) {
    if expected != actual {
        t.Fatalf("Expected %s to be '%t' but was '%t'.", label, expected, actual)
    }
}

// Asserts if a specified int64 pointer has the expected value. If not, Fatal.
func AssertNonNilIntPtr(t *testing.T, label string, expected int, actual *int) {
    if actual == nil {
        t.Fatalf("Expected %s to be '%d' but was 'nil'.", label, expected)
    }
    AssertInt(t, label, expected, *actual)
}

// Asserts if a specified int64 is nil. If not, Fatal.
func AssertIntPtrIsNil(t *testing.T, label string, actual *int) {
    if actual != nil {
        t.Fatalf("Expected %s to be 'nil' but was '%d'.", label, *actual)
    }
}

// Asserts if a specified int64 has the expected value. If not, Fatal.
func AssertInt(t *testing.T, label string, expected int, actual int) {
    if expected != actual {
        t.Fatalf("Expected %s to be '%d' but was '%d'.", label, expected, actual)
    }
}

// Asserts if a specified int64 pointer has the expected value. If not, Fatal.
func AssertNonNilInt64Ptr(t *testing.T, label string, expected int64, actual *int64) {
    if actual == nil {
        t.Fatalf("Expected %s to be '%d' but was 'nil'.", label, expected)
    }
    AssertInt64(t, label, expected, *actual)
}

// Asserts if a specified int64 is nil. If not, Fatal.
func AssertInt64PtrIsNil(t *testing.T, label int64, actual *int64) {
    if actual != nil {
        t.Fatalf("Expected %s to be 'nil' but was '%d'.", label, *actual)
    }
}

// Asserts if a specified int64 has the expected value. If not, Fatal.
func AssertInt64(t *testing.T, label string, expected int64, actual int64) {
    if expected != actual {
        t.Fatalf("Expected %s to be '%d' but was '%d'.", label, expected, actual)
    }
}

// Asserts if a specified string pointer has the expected value. If not, Fatal.
func AssertNonNilStringPtr(t *testing.T, label string, expected string, actual *string) {
    if actual == nil {
        t.Fatalf("Expected %s to be '%s' but was 'nil'.", label, expected)
    }
    AssertString(t, label, expected, *actual)
}

// Asserts if a specified string pointer is nil. If not, Fatal.
func AssertStringPtrIsNil(t *testing.T, label string, actual *string) {
    if actual != nil {
        t.Fatalf("Expected %s to be 'nil' but was '%s'.", label, *actual)
    }
}

// Asserts if a specified string has the expected value. If not, Fatal.
func AssertString(t *testing.T, label string, expected string, actual string) {
    if expected != actual {
        t.Fatalf("Expected %s to be '%s' but was '%s'.", label, expected, actual)
    }
}

// Asserts that an error is nil, else Fatal.
func AssertNilError(t *testing.T, err error) {
    if err != nil {
        t.Fatalf("Unexpected error '%s'.", err.Error())
    }
}

func AssertBadStatusCodeError(t *testing.T, expectedCode int, err error) {
    if err == nil {
        t.Fatal("Expected a bad status code error, but the error is nil.")
        return
    }
    if statusCodeErr, ok := err.(*models.BadStatusCodeError); ok {
        if statusCodeErr.ActualStatusCode != expectedCode {
            t.Fatalf("Expected a BadStatusCode of '%d' but got '%d'.", expectedCode, statusCodeErr.ActualStatusCode)
        }
        return
    }
    t.Fatalf("Expected '*models.BadStatusCodeError' but instead got '%v'.", reflect.TypeOf(err))
}