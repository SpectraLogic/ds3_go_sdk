package net

import (
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
}

func NewHttpNetwork(connectionInfo *ConnectionInfo) Network {
    return &httpNetwork{connectionInfo}
}

func (net httpNetwork) Invoke(request models.Ds3Request) (*http.Response, error) {
    verb := request.Verb().String()
    path := buildUrl(net.connectionInfo, request).String()
    httpRequest, _ := http.NewRequest(verb, path, request.GetContentStream())
    httpResponse, _ := new(http.Client).Do(httpRequest)
    return httpResponse, nil
    //TODO: error handling
}

func buildUrl(connectionInfo *ConnectionInfo, request models.Ds3Request) *url.URL {
    url := connectionInfo.Endpoint
    url.Path = request.Path()
    url.RawQuery = request.QueryParams().Encode()
    return &url
}

