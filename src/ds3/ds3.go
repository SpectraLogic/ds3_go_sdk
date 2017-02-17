package ds3

import (
    "ds3/network"
    "net/url"
)

type Client struct {
    netLayer net.Network
}

type Builder struct {
    connectionInfo *net.ConnectionInfo
}

func NewBuilder(endpoint *url.URL, creds net.Credentials) *Builder {
    return &Builder{&net.ConnectionInfo{
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
    return &Client{net.NewHttpNetwork(builder.connectionInfo)}
}

