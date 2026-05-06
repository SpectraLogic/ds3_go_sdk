package networking

import (
    "crypto/tls"
    "net/http"
    "net/url"
    "time"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

type ConnectionInfo struct {
    Endpoint                *url.URL
    Credentials             *Credentials
    Proxy                   *url.URL
    IgnoreServerCertificate bool
    MaxIdleConnsPerHost     int
    IdleConnTimeout         time.Duration
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

// Default idle connection timeout that's applied to the transport.
// Beats BlackPearl's 60s Tomcat keepAliveTimeout so the client always
// closes idle conns first and never races a server-side FIN.
const DEFAULT_IDLE_CONN_TIMEOUT = 30 * time.Second

func NewSendNetwork(connectionInfo *ConnectionInfo) Network {
    transport := http.DefaultTransport.(*http.Transport).Clone()
    transport.IdleConnTimeout = DEFAULT_IDLE_CONN_TIMEOUT

    if connectionInfo.Proxy != nil {
        transport.Proxy = http.ProxyURL(connectionInfo.Proxy)
    }
    if connectionInfo.IgnoreServerCertificate {
        transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    }
    if connectionInfo.MaxIdleConnsPerHost > 0 {
        transport.MaxIdleConnsPerHost = connectionInfo.MaxIdleConnsPerHost
    }
    if connectionInfo.IdleConnTimeout > 0 {
        transport.IdleConnTimeout = connectionInfo.IdleConnTimeout
    }

    // Pin HTTP/1.1: BlackPearl's Tomcat h2 connector trips a 64 KB
    // stream-window stall on streaming PUTs (RMS-10959). The C SDK
    // (libds3v5) does the same.
    transport.ForceAttemptHTTP2 = false
    transport.TLSNextProto = map[string]func(string, *tls.Conn) http.RoundTripper{}

    return &SendNetwork{transport: transport}
}

func (sendNetwork *SendNetwork) Invoke(httpRequest *http.Request) (models.WebResponse, error) {
    // Perform the request.
    httpResponse, reqErr := sendNetwork.transport.RoundTrip(httpRequest)
    if reqErr != nil {
        return nil, reqErr
    }

    return models.NewWrappedHttpResponse(httpResponse), nil
}
