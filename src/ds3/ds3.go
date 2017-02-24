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
    connectionPolicy *networking.ConnectionPolicy
}

func NewClientBuilder(endpoint *url.URL, creds networking.Credentials) *ClientBuilder {
    return &ClientBuilder{
        &networking.ConnectionInfo{
            Endpoint: *endpoint,
            Creds: creds,
            Proxy: nil},
        &networking.ConnectionPolicy{RedirectRetryCount: 5}}
}

func (clientBuilder *ClientBuilder) WithProxy(proxy *url.URL) *ClientBuilder {
    clientBuilder.connectionInfo.Proxy = proxy
    return clientBuilder
}

func (clientBuilder *ClientBuilder) WithRedirectRetryCount(count int) *ClientBuilder {
    clientBuilder.connectionPolicy.RedirectRetryCount = count
    return clientBuilder
}

func (clientBuilder *ClientBuilder) BuildClient() *Client {
    return &Client{networking.NewHttpNetwork(clientBuilder.connectionInfo, clientBuilder.connectionPolicy)}
}

