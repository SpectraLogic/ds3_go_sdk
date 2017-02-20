package networking

import (
    "io"
    "fmt"
    "errors"
    "net/http"
    "net/url"
)

type Network interface {
    Invoke(request Request) (Response, error)
}

type ConnectionInfo struct {
    Endpoint url.URL
    Creds Credentials

    Proxy *url.URL
    RedirectRetryCount int
}

type Credentials struct {
    AccessId, Key string
}

type httpNetwork struct {
    connectionInfo *ConnectionInfo
    transport *http.Transport
}

func NewHttpNetwork(connectionInfo *ConnectionInfo) Network {
    return &httpNetwork{
        connectionInfo,
        &http.Transport{ Proxy: http.ProxyURL(connectionInfo.Proxy) },
    }
}

func (net httpNetwork) Invoke(request Request) (Response, error) {
    // Open up the content stream.
    stream := request.GetContentStream()
    if stream != nil {
        defer stream.Close()
    }

    // Handle as many 307's as we're allowed.
    for i := 0; i < net.connectionInfo.RedirectRetryCount; i++ {
        // Seek to the beginning of the request stream.
        if stream != nil {
            if _, seekErr := stream.Seek(0, 0); seekErr != nil {
                return nil, seekErr
            }
        }

        // Build the request.
        httpRequest, makeReqErr := buildHttpRequest(net.connectionInfo, request, stream)
        if makeReqErr != nil {
            return nil, makeReqErr
        }

        // Perform the request.
        httpResponse, reqErr := net.transport.RoundTrip(httpRequest)
        if reqErr != nil {
            return nil, reqErr
        }

        // If it wasn't a redirect then return.
        if httpResponse.StatusCode != http.StatusTemporaryRedirect {
            return &wrappedHttpResponse{httpResponse}, nil
        }
    }

    // We had as many 307 redirects as we were allowed to use.
    return nil, errors.New(fmt.Sprintf(
        "The server is busy. Retried the max number of %d times.",
        net.connectionInfo.RedirectRetryCount,
    ))
}

func buildHttpRequest(conn *ConnectionInfo, request Request, stream SizedReadCloser) (*http.Request, error) {
    var reader io.Reader
    if stream != nil {
        reader = proxiedReader{stream}
    }

    // Build the basic request with the verb, url, and payload (if any).
    httpRequest, err := http.NewRequest(
        request.Verb().String(),
        buildUrl(conn, request).String(),
        reader,
    )
    if err != nil {
        return nil, err
    }

    // Set the content length if we have a payload.
    if stream != nil {
        httpRequest.ContentLength = stream.Size()
    }

    // Set the request headers such as authorization and date.
    setRequestHeaders(httpRequest, conn.Creds, request)

    return httpRequest, nil
}

// http.NewRequest takes an io.Reader, but the client methods do a run time
// type check for io.Closer and call close. Since we want to be able to do
// multiple reads of the same stream in the case of 307 redirects, we use this
// structure to prevent the http client from calling Close.
type proxiedReader struct {
    innerReader io.Reader
}

func (self proxiedReader) Read(p []byte) (n int, err error) {
    return self.innerReader.Read(p)
}

func buildUrl(conn *ConnectionInfo, request Request) *url.URL {
    url := conn.Endpoint
    url.Path = request.Path()
    url.RawQuery = request.QueryParams().Encode()
    return &url
}

