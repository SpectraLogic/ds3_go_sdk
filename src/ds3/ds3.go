package ds3

import (
    "ds3/networking"
    "net/url"
)

type Client struct {
    netLayer networking.Network
}

type ClientBuilder struct {
    connectionInfo *networking.ConnectionInfo
}

func NewClientBuilder(endpoint *url.URL, creds networking.Credentials) *ClientBuilder {
    return &ClientBuilder{&networking.ConnectionInfo{
        Endpoint: *endpoint,
        Creds: creds,
        RedirectRetryCount: 5,
        Proxy: nil,
    }}
}

func (clientBuilder *ClientBuilder) WithProxy(proxy *url.URL) *ClientBuilder {
    clientBuilder.connectionInfo.Proxy = proxy
    return clientBuilder
}

func (clientBuilder *ClientBuilder) WithRedirectRetryCount(count int) *ClientBuilder {
    clientBuilder.connectionInfo.RedirectRetryCount = count
    return clientBuilder
}

func (clientBuilder *ClientBuilder) BuildClient() *Client {
    return &Client{networking.NewHttpNetwork(clientBuilder.connectionInfo)}
}

