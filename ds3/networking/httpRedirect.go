package networking

import (
    "errors"
    "fmt"
    "net/http"
    "spectra/ds3_go_sdk/ds3/models"
)

type httpRedirectPolicy struct {
    maxRedirect int
}

// Decorator for Network which handles 307 temporary redirect retries
type HttpTempRedirectDecorator struct {
    network Network
    policy *httpRedirectPolicy
}

func NewHttpTempRedirectDecorator(network Network, maxRedirect int) (Network) {
    return &HttpTempRedirectDecorator{
        network: network,
        policy: &httpRedirectPolicy{maxRedirect: maxRedirect},
    }
}

func (tempRedirectDecorator *HttpTempRedirectDecorator) Invoke(httpRequest *http.Request) (models.WebResponse, error) {
    // Handle as many 307's as we're allowed.
    for i := 0; i <= tempRedirectDecorator.policy.maxRedirect; i++ {
        ds3Response, err := tempRedirectDecorator.network.Invoke(httpRequest)

        // If request was performed successfully then return response.
        if err == nil {
            return ds3Response, nil
        }

        // If there was a non-redirect error then return.
        if statusErr, ok := err.(models.BadStatusCodeError); ok == false || statusErr.ActualStatusCode != http.StatusTemporaryRedirect {
            return nil, err
        }
    }

    // We had as many 307 redirects as we were allowed to use.
    return nil, errors.New(fmt.Sprintf(
        "The server is busy. Retried the max number of %d times.",
        tempRedirectDecorator.policy.maxRedirect,
    ))
}