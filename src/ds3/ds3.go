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
    clientPolicy *ClientPolicy
}

type ClientPolicy struct {
    maxRetries int // Maximum number of times to attempt sending a request amidst network issues
    maxRedirect int // Maximum number of times to attempt redirect retries
}

func NewClientBuilder(endpoint *url.URL, creds networking.Credentials) *ClientBuilder {
    return &ClientBuilder{
        &networking.ConnectionInfo{
            Endpoint: *endpoint,
            Creds: creds,
            Proxy: nil},
        &ClientPolicy{
            maxRetries: 5,
            maxRedirect: 5}}
}

func (clientBuilder *ClientBuilder) WithProxy(proxy *url.URL) *ClientBuilder {
    clientBuilder.connectionInfo.Proxy = proxy
    return clientBuilder
}

func (clientBuilder *ClientBuilder) WithRedirectRetryCount(count int) *ClientBuilder {
    clientBuilder.clientPolicy.maxRedirect = count
    return clientBuilder
}

func (clientBuilder *ClientBuilder) WithNetworkRetryCount(count int) *ClientBuilder {
    clientBuilder.clientPolicy.maxRetries = count
    return clientBuilder
}

func (clientBuilder *ClientBuilder) BuildClient() *Client {
    return &Client{networking.NewHttpNetwork(clientBuilder.connectionInfo)}
}

