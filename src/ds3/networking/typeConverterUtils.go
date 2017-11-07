package networking

import "strconv"

// Contains utils for converting pointers to primitives into string pointers.
// Used to handle conversions of request handler optional parameters into
// query parameters within clients.

// Converts an *int64 into a *string
func Int64PtrToStrPtr(int *int64) *string {
    if int == nil {
        return nil
    }
    str := strconv.FormatInt(*int, 10)
    return &str
}

// Converts an *int into a *string
func IntPtrToStrPtr(int *int) *string {
    if int == nil {
        return nil
    }
    str := strconv.Itoa(*int)
    return &str
}

// Converts a *bool into a *string
func BoolPtrToStrPtr(bool *bool) *string {
    if bool == nil {
        return nil
    }
    str := strconv.FormatBool(*bool)
    return &str
}

// Converts a *float64 into a *string
func Float64PtrToStrPtr(float *float64) *string {
    if float == nil {
        return nil
    }
    str := strconv.FormatFloat(*float, 'f', -1, 64)
    return &str
}

// Interface for a to string method
type ToStringInterface interface {
    String() string
}

// Converts an item that meets the ToStringInterface into a *string
func InterfaceToStrPtr(item ToStringInterface) *string {
    if item == nil {
        return nil
    }
    str := item.String()
    return &str
}