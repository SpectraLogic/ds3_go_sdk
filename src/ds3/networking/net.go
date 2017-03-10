package networking

import (
    "io"
    "net/http"
    "net/url"
)

type Network interface {
    Invoke(ds3Request Ds3Request) (WebResponse, error)
}

type ConnectionInfo struct {
    Endpoint url.URL
    Creds Credentials
    Proxy *url.URL
}

type Credentials struct {
    AccessId string
    Key string
}

type httpNetwork struct {
    connectionInfo *ConnectionInfo
    transport *http.Transport
}

func NewHttpNetwork(connectionInfo *ConnectionInfo) Network {
    return &httpNetwork{
        connectionInfo,
        &http.Transport{ Proxy: http.ProxyURL(connectionInfo.Proxy) },
    }
}

func (httpNetwork *httpNetwork) Invoke(ds3Request Ds3Request) (WebResponse, error) {
    // Open up the content stream.
    stream := ds3Request.GetContentStream()
    if stream != nil {
        defer stream.Close()
    }

    // Build the request.
    httpRequest, makeReqErr := buildHttpRequest(httpNetwork.connectionInfo, ds3Request, stream)
    if makeReqErr != nil {
        return nil, makeReqErr
    }

    // Perform the request.
    httpResponse, reqErr := httpNetwork.transport.RoundTrip(httpRequest)
    if reqErr != nil {
        return nil, RoundTripError(reqErr.Error())
    }

    //If it was a redirect then return redirect error
    if httpResponse.StatusCode == http.StatusTemporaryRedirect {
        return nil, TemporaryRedirectError(httpResponse.StatusCode)
    }

    return &wrappedHttpResponse{httpResponse}, nil
}

func buildHttpRequest(conn *ConnectionInfo, ds3Request Ds3Request, stream SizedReadCloser) (*http.Request, error) {
    var reader io.Reader
    if stream != nil {
        reader = &proxiedReader{stream}
    }

    // Build the basic request with the verb, url, and payload (if any).
    verb, verbErr := ds3Request.Verb().String()
    if verbErr != nil {
        return nil, verbErr
    }

    httpRequest, err := http.NewRequest(
        verb,
        buildUrl(conn, ds3Request).String(),
        reader,
    )
    if err != nil {
        return nil, err
    }

    // Set the content length if we have a payload.
    if stream != nil {
        var sizeErr error
        httpRequest.ContentLength, sizeErr = stream.Size()
        if sizeErr != nil {
            return nil, sizeErr
        }
    }

    // Build the authentication
    now := getCurrentTime()

    authHeaderVal := buildAuthHeaderValue(conn.Creds, signatureFields{
        Verb: verb,
        ContentHash: ds3Request.GetChecksum().ContentHash,
        ContentType: ds3Request.Header().Get(ContentTypeKey),
        Path: ds3Request.Path(),
        Date: now,
    })

    // Set the http request headers such as authorization and date.
    setHttpRequestHeaders(httpRequest, ds3Request, headerFields{
        AuthHeaderVal: authHeaderVal,
        DateHeaderVal: now,
        ChecksumType:  ds3Request.GetChecksum().Type,
        ContentHash:   ds3Request.GetChecksum().ContentHash,
    })

    return httpRequest, nil
}

// http.NewRequest takes an io.Reader, but the client methods do a run time
// type check for io.Closer and call close. Since we want to be able to do
// multiple reads of the same stream in the case of 307 redirects, we use this
// structure to prevent the http client from calling Close.
type proxiedReader struct {
    innerReader io.Reader
}

func (proxiedReader *proxiedReader) Read(p []byte) (n int, err error) {
    return proxiedReader.innerReader.Read(p)
}

func buildUrl(conn *ConnectionInfo, ds3Request Ds3Request) *url.URL {
    httpUrl := conn.Endpoint
    httpUrl.Path = ds3Request.Path()
    httpUrl.RawQuery = ds3Request.QueryParams().Encode()
    return &httpUrl
}

