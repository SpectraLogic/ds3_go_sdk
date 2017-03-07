package networking

import (
    "errors"
    "fmt"
)

type httpRedirectPolicy struct {
    maxRedirect int
}

type HttpTempRedirectDecorator struct {
    network *Network
    policy *httpRedirectPolicy
}


// Decorator for Network which handles 307 redirect retries
func NewHttpTempRedirectDecorator(network *Network, maxRedirect int) (Network) {
    return &HttpTempRedirectDecorator{
        network: network,
        policy: &httpRedirectPolicy{maxRedirect: maxRedirect},
    }
}

func (tempRedirectDecorator *HttpTempRedirectDecorator) Invoke(request Ds3Request) (WebResponse, error) {
    // Handle as many 307's as we're allowed.
    for i := 0; i <= tempRedirectDecorator.policy.maxRedirect; i++ {
        ds3Response, err := (*tempRedirectDecorator.network).Invoke(request)

        // If request was performed successfully then return response.
        if err == nil {
            return ds3Response, nil
        }

        // If there was a non-redirect error then return.
        if _, ok := err.(TemporaryRedirectError); ok == false {
            return nil, err
        }
    }

    // We had as many 307 redirects as we were allowed to use.
    return nil, errors.New(fmt.Sprintf(
        "The server is busy. Retried the max number of %d times.",
        tempRedirectDecorator.policy.maxRedirect,
    ))
}