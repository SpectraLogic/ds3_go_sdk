package ds3

import (
    "io"
    "io/ioutil"
    "testing"
    "net/url"
    "net/http"
    "reflect"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/networking"
)

func mockedClient(t *testing.T) mockedClientWithTest {
    return &mockedNet{t: t}
}

type mockedClientWithTest interface {
    Expecting(verb string, path string, queryParams *url.Values, requestHeaders *http.Header, request *string) mockedClientWithExpectation
}

type mockedClientWithExpectation interface {
    Returning(statusCode int, response string, headers *http.Header) *Client
}

type mockedNet struct {
    // Test manager.
    t *testing.T

    // Expected data to validate.
    verb string
    path string
    queryParams *url.Values
    requestHeaders *http.Header
    request *string

    // Response data to return.
    statusCode int
    response string
    headers *http.Header
}

func (mockedNet *mockedNet) Expecting(verb string, path string, queryParams *url.Values, requestHeaders *http.Header, request *string) mockedClientWithExpectation {
    mockedNet.verb = verb
    mockedNet.path = path
    mockedNet.queryParams = queryParams
    mockedNet.requestHeaders = requestHeaders
    mockedNet.request = request
    return mockedNet
}

func (mockedNet *mockedNet) Returning(statusCode int, response string, headers *http.Header) *Client {
    mockedNet.statusCode = statusCode
    mockedNet.response = response
    mockedNet.headers = headers
    var url url.URL
    var endpointUrl, _ = url.Parse("/")
    return &Client{
        SendNetwork:  mockedNet,
        clientPolicy: &ClientPolicy{5, 5},
        connectionInfo: &networking.ConnectionInfo{
            Endpoint: endpointUrl,
            Credentials: &networking.Credentials{AccessId: "AccessId", Key: "Key"},
        },
    }
}

func (mockedNet *mockedNet) Invoke(httpRequest *http.Request) (models.WebResponse, error) {
    // Verify the verb.
    if httpRequest.Method != mockedNet.verb {
        mockedNet.t.Errorf("Expected verb '%s' but got '%s'.", mockedNet.verb, httpRequest.Method)
    }

    // Verify the path.
    if httpRequest.URL.Path != mockedNet.path {
        mockedNet.t.Errorf("Expected path '%s' but got '%s'.", mockedNet.path, httpRequest.URL.Path)
    }

    // Verify the query parameters.
    actualQueryParams := httpRequest.URL.RawQuery
    expectedQueryParams := mockedNet.queryParams.Encode()
    if actualQueryParams != expectedQueryParams {
        mockedNet.t.Errorf(
            "Expected query params '%s' but got '%s'.",
            expectedQueryParams,
            actualQueryParams,
        )
    }

    // Verify request headers
    actualRequestHeaders := httpRequest.Header
    actualRequestHeaders.Del("Date") //remove date
    actualRequestHeaders.Del("Authorization") //remove authorization
    expectedRequestHeaders := *mockedNet.requestHeaders
    if len(actualRequestHeaders) != len(expectedRequestHeaders) {
        mockedNet.t.Errorf(
            "Expected request headers '%s' but got '%s'.",
            expectedRequestHeaders,
            actualRequestHeaders,
        )
    } else {
        for key, expectedValues := range expectedRequestHeaders {
            actualValues := actualRequestHeaders[key]
            if !reflect.DeepEqual(actualValues, expectedValues) {
                mockedNet.t.Errorf(
                    "For request header with key '%s', expected values '%v' but got '%v'.",
                    key,
                    expectedValues,
                    actualValues,
                )
            }
        }
    }

    // Verify the request contents.
    if mockedNet.request != nil {
        contentStream := httpRequest.Body
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
    return BuildByteReaderWithSizeDecorator([]byte(mockedNet.response))
}

func (mockedNet *mockedNet) Header() *http.Header {
    if mockedNet.headers == nil {
        return &http.Header{}
    } else {
        return mockedNet.headers
    }
}

