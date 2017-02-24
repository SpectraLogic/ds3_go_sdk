package networking

import (
    "io"
    "fmt"
    "errors"
    "net/http"
    "net/url"
)

type Network interface {
    Invoke(ds3Request Ds3Request) (Ds3Response, error)
}

type ConnectionInfo struct {
    Endpoint url.URL
    Creds Credentials

    Proxy *url.URL
    RedirectRetryCount int
}

type Credentials struct {
    AccessId string
    Key string
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

func (httpNetwork *httpNetwork) Invoke(ds3Request Ds3Request) (Ds3Response, error) {
    // Open up the content stream.
    stream := ds3Request.GetContentStream()
    if stream != nil {
        defer stream.Close()
    }

    // Handle as many 307's as we're allowed.
    for i := 0; i < httpNetwork.connectionInfo.RedirectRetryCount; i++ {
        // Seek to the beginning of the request stream.
        if stream != nil {
            if _, seekErr := stream.Seek(0, 0); seekErr != nil {
                return nil, seekErr
            }
        }

        // Build the request.
        httpRequest, makeReqErr := buildHttpRequest(httpNetwork.connectionInfo, ds3Request, stream)
        if makeReqErr != nil {
            return nil, makeReqErr
        }

        // Perform the request.
        httpResponse, reqErr := httpNetwork.transport.RoundTrip(httpRequest)
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
        httpNetwork.connectionInfo.RedirectRetryCount,
    ))
}

func buildHttpRequest(conn *ConnectionInfo, ds3Request Ds3Request, stream SizedReadCloser) (*http.Request, error) {
    var reader io.Reader
    if stream != nil {
        reader = &proxiedReader{stream}
    }

    // Build the basic request with the verb, url, and payload (if any).
    verb, verbErr := ds3Request.Verb().String()
    if verbErr != nil {
        return nil, verbErr
    }

    httpRequest, err := http.NewRequest(
        verb,
        buildUrl(conn, ds3Request).String(),
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
    headersErr := setRequestHeaders(httpRequest, conn.Creds, ds3Request)
    if headersErr != nil {
        return nil, headersErr
    }

    return httpRequest, nil
}

// http.NewRequest takes an io.Reader, but the client methods do a run time
// type check for io.Closer and call close. Since we want to be able to do
// multiple reads of the same stream in the case of 307 redirects, we use this
// structure to prevent the http client from calling Close.
type proxiedReader struct {
    innerReader io.Reader
}

func (proxiedReader *proxiedReader) Read(p []byte) (n int, err error) {
    return proxiedReader.innerReader.Read(p)
}

func buildUrl(conn *ConnectionInfo, ds3Request Ds3Request) *url.URL {
    httpUrl := conn.Endpoint
    httpUrl.Path = ds3Request.Path()
    httpUrl.RawQuery = ds3Request.QueryParams().Encode()
    return &httpUrl
}

