package ds3

import "testing"

// Asserts if a specified bool pointer has the expected value. If not, Fatal.
func assertNonNilBoolPtr(t *testing.T, xmlTag string, expected bool, actual *bool) {
    if actual == nil {
        t.Fatalf("Expected %s to be '%t' but was 'nil'.", xmlTag, expected)
    }
    assertBool(t, xmlTag, expected, *actual)
}

// Asserts if a specified bool pointer is nil. If not, Fatal.
func assertBoolPtrIsNil(t *testing.T, xmlTag bool, actual *bool) {
    if actual != nil {
        t.Fatalf("Expected %s to be 'nil' but was '%t'.", xmlTag, *actual)
    }
}

// Asserts if a specified bool has the expected value. If not, Fatal.
func assertBool(t *testing.T, xmlTag string, expected bool, actual bool) {
    if expected != actual {
        t.Fatalf("Expected %s to be '%t' but was '%t'.", xmlTag, expected, actual)
    }
}

// Asserts if a specified int64 pointer has the expected value. If not, Fatal.
func assertNonNilIntPtr(t *testing.T, xmlTag string, expected int64, actual *int64) {
    if actual == nil {
        t.Fatalf("Expected %s to be '%d' but was 'nil'.", xmlTag, expected)
    }
    assertInt(t, xmlTag, expected, *actual)
}

// Asserts if a specified int64 is nil. If not, Fatal.
func assertIntPtrIsNil(t *testing.T, xmlTag int64, actual *int64) {
    if actual != nil && *actual != 0 { //TODO remove default value check once nil ptr parsing is fixed
        t.Fatalf("Expected %s to be 'nil' but was '%d'.", xmlTag, *actual)
    }
}

// Asserts if a specified int64 has the expected value. If not, Fatal.
func assertInt(t *testing.T, xmlTag string, expected int64, actual int64) {
    if expected != actual {
        t.Fatalf("Expected %s to be '%d' but was '%d'.", xmlTag, expected, actual)
    }
}

// Asserts if a specified string pointer has the expected value. If not, Fatal.
func assertNonNilStringPtr(t *testing.T, xmlTag string, expected string, actual *string) {
    if actual == nil {
        t.Fatalf("Expected %s to be '%s' but was 'nil'.", xmlTag, expected)
    }
    assertString(t, xmlTag, expected, *actual)
}

// Asserts if a specified string pointer is nil. If not, Fatal.
func assertStringPtrIsNil(t *testing.T, xmlTag string, actual *string) {
    if actual != nil && *actual != "" { //TODO remove empty string comparison once nil ptr parsing is fixed
        t.Fatalf("Expected %s to be 'nil' but was '%s'.", xmlTag, *actual)
    }
}

// Asserts if a specified string has the expected value. If not, Fatal.
func assertString(t *testing.T, xmlTag string, expected string, actual string) {
    if expected != actual {
        t.Fatalf("Expected %s to be '%s' but was '%s'.", xmlTag, expected, actual)
    }
}