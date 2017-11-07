package networking

import (
    "net/http"
    "net/url"
)

type ConnectionInfo struct {
    Endpoint    *url.URL
    Credentials *Credentials
    Proxy       *url.URL
}

type Credentials struct {
    AccessId string
    Key string
}

type Network interface {
    Invoke(httpRequest *http.Request) (WebResponse, error)
}

// Performs http requests
type SendNetwork struct {
    transport *http.Transport
}

func NewSendNetwork(connectionInfo *ConnectionInfo) Network {
    return &SendNetwork{
        transport: &http.Transport{Proxy: http.ProxyURL(connectionInfo.Proxy)},
    }
}

func (sendNetwork *SendNetwork) Invoke(httpRequest *http.Request) (WebResponse, error) {
    // Perform the request.
    httpResponse, reqErr := sendNetwork.transport.RoundTrip(httpRequest)
    if reqErr != nil {
        return nil, reqErr
    }

    return &wrappedHttpResponse{httpResponse}, nil
}
