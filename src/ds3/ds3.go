package ds3

import (
    "ds3/networking"
    "net/url"
)

type Client struct {
    netLayer networking.Network
}

type Builder struct {
    connectionInfo *networking.ConnectionInfo
}

func NewBuilder(endpoint *url.URL, creds networking.Credentials) *Builder {
    return &Builder{&networking.ConnectionInfo{
        Endpoint: *endpoint,
        Creds: creds,
        RedirectRetryCount: 5,
        Proxy: nil,
    }}
}

func (builder *Builder) WithProxy(proxy *url.URL) *Builder {
    builder.connectionInfo.Proxy = proxy
    return builder
}

func (builder *Builder) WithRedirectRetryCount(count int) *Builder {
    builder.connectionInfo.RedirectRetryCount = count
    return builder
}

func (builder *Builder) Build() *Client {
    return &Client{networking.NewHttpNetwork(builder.connectionInfo)}
}

