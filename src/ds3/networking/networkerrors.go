package networking

import "fmt"

// Error denoting a request has returned a temporary redirect status code
type TemporaryRedirectError int

func (temporaryRedirectError TemporaryRedirectError) Error() string {
    return fmt.Sprintf("Temporary Redirect, received status code %d", temporaryRedirectError)
}

// Denotes an error occurred during attempt to perform request
type RoundTripError string

func (roundTripError RoundTripError) Error() string {
    return string(roundTripError)
}