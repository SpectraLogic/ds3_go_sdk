package networking

import (
    "crypto/tls"
    "net/http"
    "net/url"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

type ConnectionInfo struct {
    Endpoint                *url.URL
    Credentials             *Credentials
    Proxy                   *url.URL
    IgnoreServerCertificate bool
}

type Credentials struct {
    AccessId string
    Key string
}

type Network interface {
    Invoke(httpRequest *http.Request) (models.WebResponse, error)
}

// Performs http requests
type SendNetwork struct {
    transport *http.Transport
}

func NewSendNetwork(connectionInfo *ConnectionInfo) Network {
    sendNetwork := &SendNetwork{
        transport: &http.Transport{
            Proxy: http.ProxyURL(connectionInfo.Proxy),
        },
    }
    if connectionInfo.IgnoreServerCertificate {
        sendNetwork.transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    }
    return sendNetwork
}

func (sendNetwork *SendNetwork) Invoke(httpRequest *http.Request) (models.WebResponse, error) {
    // Perform the request.
    httpResponse, reqErr := sendNetwork.transport.RoundTrip(httpRequest)
    if reqErr != nil {
        return nil, reqErr
    }

    return models.NewWrappedHttpResponse(httpResponse), nil
}
