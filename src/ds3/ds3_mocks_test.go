package ds3

import (
    "io"
    "testing"
    "net/url"
    "ds3/net"
)

func mockedClient(t *testing.T) mockedClientWithTest {
    return &mockedNet{t: t}
}

type mockedClientWithTest interface {
    Expecting(verb net.HttpVerb, path string, queryParams *url.Values) mockedClientWithExpectation
}

type mockedClientWithExpectation interface {
    Returning(statusCode int, response string) *Client
}

type mockedNet struct {
    // Test manager.
    t *testing.T

    // Expected data to validate.
    verb net.HttpVerb
    path string
    queryParams *url.Values

    // Response data to return.
    statusCode int
    response string
}

func (self *mockedNet) Expecting(verb net.HttpVerb, path string, queryParams *url.Values) mockedClientWithExpectation {
    self.verb = verb
    self.path = path
    self.queryParams = queryParams
    return self
}

func (self *mockedNet) Returning(statusCode int, response string) *Client {
    self.statusCode = statusCode
    self.response = response
    return &Client{self}
}

func (self *mockedNet) Invoke(request net.Request) (net.Response, error) {
    // Verify the verb.
    verb := request.Verb()
    if verb != self.verb {
        self.t.Errorf("Expected verb '%s' but got '%s'.", self.verb.String(), verb.String())
    }

    // Verify the path.
    path := request.Path()
    if path != self.path {
        self.t.Errorf("Expected path '%s' but got '%s'.", self.path, path)
    }

    // Verify the query parameters.
    actualQueryParams := request.QueryParams().Encode()
    expectedQueryParams := self.queryParams.Encode()
    if actualQueryParams != expectedQueryParams {
        self.t.Errorf(
            "Expected query params '%s' but got '%s'.",
            actualQueryParams,
            expectedQueryParams,
        )
    }

    // Return the net.Response object, which we just made mockedNet also implement.
    return self, nil
}

func (self *mockedNet) StatusCode() int {
    return self.statusCode
}

func (self *mockedNet) Body() io.ReadCloser {
    return net.BuildSizedReadCloser([]byte(self.response))
}

