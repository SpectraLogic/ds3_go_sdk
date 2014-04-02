package net

import (
    "fmt"
    "errors"
    "net/http"
    "net/url"
    "ds3/models"
)

type Network interface {
    Invoke(request models.Ds3Request) (*http.Response, error)
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

//TODO: improve error handling
func (net httpNetwork) Invoke(request models.Ds3Request) (*http.Response, error) {
    // Build the request.
    httpRequest, makeReqErr := buildHttpRequest(net.connectionInfo, request)
    if makeReqErr != nil {
        return nil, makeReqErr
    }

    // Handle as many 307's as we're allowed.
    for i := 0; i < net.connectionInfo.RedirectRetryCount; i++ {
        // Perform the request.
        httpResponse, reqErr := net.transport.RoundTrip(httpRequest)
        if reqErr != nil {
            return nil, reqErr
        }

        // If it wasn't a redirect then return.
        if httpResponse.StatusCode != http.StatusTemporaryRedirect {
            return httpResponse, nil
        }
    }

    // We had as many 307 redirects as we were allowed to use.
    return nil, errors.New(fmt.Sprintf(
        "The server is busy. Retried the max number of %d times.",
        net.connectionInfo.RedirectRetryCount,
    ))
}

func buildHttpRequest(conn *ConnectionInfo, request models.Ds3Request) (*http.Request, error) {
    httpRequest, err := http.NewRequest(
        request.Verb().String(),
        buildUrl(conn, request).String(),
        request.GetContentStream(),
    )
    if err != nil {
        return nil, err
    }
    setRequestHeaders(httpRequest, conn.Creds, request)
    return httpRequest, nil
}

func buildUrl(conn *ConnectionInfo, request models.Ds3Request) *url.URL {
    url := conn.Endpoint
    url.Path = request.Path()
    url.RawQuery = request.QueryParams().Encode()
    return &url
}

