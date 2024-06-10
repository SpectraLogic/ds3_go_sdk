package ds3

import (
    "github.com/SpectraLogic/ds3_go_sdk/ds3/networking"
    "github.com/SpectraLogic/ds3_go_sdk/sdk_log"
    "net/url"
)

const (
    HTTP_VERB_GET    = "GET"
    HTTP_VERB_PUT    = "PUT"
    HTTP_VERB_POST   = "POST"
    HTTP_VERB_DELETE = "DELETE"
    HTTP_VERB_HEAD   = "HEAD"
    HTTP_VERB_PATCH  = "PATCH"
)

type Client struct {
    SendNetwork  networking.Network
    clientPolicy *ClientPolicy
    connectionInfo *networking.ConnectionInfo

    // Logger where all messages will be logged to
    sdk_log.Logger
}

type ClientBuilder struct {
    connectionInfo *networking.ConnectionInfo
    clientPolicy   *ClientPolicy
    logger         sdk_log.Logger
}

type ClientPolicy struct {
    maxRetries int // Maximum number of times to attempt sending a request amidst network issues
    maxRedirect int // Maximum number of times to attempt redirect retries
}

const DEFAULT_MAX_RETRIES = 5
const DEFAULT_MAX_REDIRECTS = 5

func NewClientBuilder(endpoint *url.URL, creds *networking.Credentials) *ClientBuilder {
    return &ClientBuilder{
        connectionInfo: &networking.ConnectionInfo{
            Endpoint:    endpoint,
            Credentials: creds,
            Proxy:       nil},
        clientPolicy: &ClientPolicy{
            maxRetries: DEFAULT_MAX_RETRIES,
            maxRedirect: DEFAULT_MAX_REDIRECTS},
    }
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

func (clientBuilder *ClientBuilder) WithIgnoreServerCertificate(ignoreServerCert bool) *ClientBuilder {
    clientBuilder.connectionInfo.IgnoreServerCertificate = ignoreServerCert
    return clientBuilder
}

func (clientBuilder *ClientBuilder) WithLogger(logger sdk_log.Logger) *ClientBuilder {
    clientBuilder.logger = logger
    return clientBuilder
}

func (clientBuilder *ClientBuilder) BuildClient() *Client {
    if clientBuilder.logger == nil {
        clientBuilder.logger = sdk_log.NewSimpleLogger()
    }

    return &Client{
        SendNetwork:    networking.NewSendNetwork(clientBuilder.connectionInfo),
        clientPolicy:   clientBuilder.clientPolicy,
        connectionInfo: clientBuilder.connectionInfo,
        Logger:         clientBuilder.logger,
    }
}
