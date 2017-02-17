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
    Expecting(verb net.HttpVerb, path string, queryParams *url.Values, request *string) mockedClientWithExpectation
}

type mockedClientWithExpectation interface {
    Returning(statusCode int, response string, headers *http.Header) *Client
}

type mockedNet struct {
    // Test manager.
    t *testing.T

    // Expected data to validate.
    verb net.HttpVerb
    path string
    queryParams *url.Values
    request *string

    // Response data to return.
    statusCode int
    response string
    headers *http.Header
}

func (self *mockedNet) Expecting(verb net.HttpVerb, path string, queryParams *url.Values, request *string) mockedClientWithExpectation {
    self.verb = verb
    self.path = path
    self.queryParams = queryParams
    self.request = request
    return self
}

func (self *mockedNet) Returning(statusCode int, response string, headers *http.Header) *Client {
    self.statusCode = statusCode
    self.response = response
    self.headers = headers
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
            expectedQueryParams,
            actualQueryParams,
        )
    }

    // Verify the request contents.
    if self.request != nil {
        contentStream := request.GetContentStream()
        if contentStream == nil {
            self.t.Error("Expected a non-nil content stream, but got nil.")
        } else {
            defer contentStream.Close()
            bs, readErr := ioutil.ReadAll(contentStream)
            if readErr != nil {
                self.t.Errorf("Unexpected error '%s' while reading content stream.", readErr.Error())
            } else if string(bs) != *self.request {
                self.t.Errorf("Expected '%s' but got '%s'.", *self.request, string(bs))
            }
        }
    }

    // Return the net.Response interface, which we just made mockedNet also implement.
    return self, nil
}

func (self *mockedNet) StatusCode() int {
    return self.statusCode
}

func (self *mockedNet) Body() io.ReadCloser {
    return net.BuildSizedReadCloser([]byte(self.response))
}

func (self *mockedNet) Header() *http.Header {
    if self.headers == nil {
        return &http.Header{}
    } else {
        return self.headers
    }
}

