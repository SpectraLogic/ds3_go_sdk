package ds3

import (
    "io"
    "io/ioutil"
    "testing"
    "net/url"
    "net/http"
    "ds3/networking"
)

func mockedClient(t *testing.T) mockedClientWithTest {
    return &mockedNet{t: t}
}

type mockedClientWithTest interface {
    Expecting(verb networking.HttpVerb, path string, queryParams *url.Values, request *string) mockedClientWithExpectation
}

type mockedClientWithExpectation interface {
    Returning(statusCode int, response string, headers *http.Header) *Client
}

type mockedNet struct {
    // Test manager.
    t *testing.T

    // Expected data to validate.
    verb networking.HttpVerb
    path string
    queryParams *url.Values
    request *string

    // Response data to return.
    statusCode int
    response string
    headers *http.Header
}

func (mockedNet *mockedNet) Expecting(verb networking.HttpVerb, path string, queryParams *url.Values, request *string) mockedClientWithExpectation {
    mockedNet.verb = verb
    mockedNet.path = path
    mockedNet.queryParams = queryParams
    mockedNet.request = request
    return mockedNet
}

func (mockedNet *mockedNet) Returning(statusCode int, response string, headers *http.Header) *Client {
    mockedNet.statusCode = statusCode
    mockedNet.response = response
    mockedNet.headers = headers
    return &Client{mockedNet, &ClientPolicy{5, 5}}
}

func (mockedNet *mockedNet) Invoke(ds3Request networking.Ds3Request) (networking.WebResponse, error) {
    // Verify the verb.
    verb := ds3Request.Verb()
    if verb != mockedNet.verb {
        actualVerb, err := verb.String()
        if err != nil {
            mockedNet.t.Error(err)
        } else {
            mockedVerb, _ := mockedNet.verb.String()
            mockedNet.t.Errorf("Expected verb '%s' but got '%s'.", mockedVerb, actualVerb)
        }
    }

    // Verify the path.
    path := ds3Request.Path()
    if path != mockedNet.path {
        mockedNet.t.Errorf("Expected path '%s' but got '%s'.", mockedNet.path, path)
    }

    // Verify the query parameters.
    actualQueryParams := ds3Request.QueryParams().Encode()
    expectedQueryParams := mockedNet.queryParams.Encode()
    if actualQueryParams != expectedQueryParams {
        mockedNet.t.Errorf(
            "Expected query params '%s' but got '%s'.",
            expectedQueryParams,
            actualQueryParams,
        )
    }

    // Verify the request contents.
    if mockedNet.request != nil {
        contentStream := ds3Request.GetContentStream()
        if contentStream == nil {
            mockedNet.t.Error("Expected a non-nil content stream, but got nil.")
        } else {
            defer contentStream.Close()
            bs, readErr := ioutil.ReadAll(contentStream)
            if readErr != nil {
                mockedNet.t.Errorf("Unexpected error '%s' while reading content stream.", readErr.Error())
            } else if string(bs) != *mockedNet.request {
                mockedNet.t.Errorf("Expected '%s' but got '%s'.", *mockedNet.request, string(bs))
            }
        }
    }

    // Return the net.WebResponse interface, which we just made mockedNet also implement.
    return mockedNet, nil
}

func (mockedNet *mockedNet) StatusCode() int {
    return mockedNet.statusCode
}

func (mockedNet *mockedNet) Body() io.ReadCloser {
    return networking.BuildByteReaderWithSizeDecorator([]byte(mockedNet.response))
}

func (mockedNet *mockedNet) Header() *http.Header {
    if mockedNet.headers == nil {
        return &http.Header{}
    } else {
        return mockedNet.headers
    }
}

