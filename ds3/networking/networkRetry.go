package networking

import (
    "errors"
    "fmt"
    "log"
    "net/http"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

type networkRetryPolicy struct {
    maxRetries int
}

// Decorator for Network which handles network related retries
type NetworkRetryDecorator struct {
    network Network
    policy *networkRetryPolicy
}

func NewNetworkRetryDecorator(network Network, maxRetires int) (Network) {
    return &NetworkRetryDecorator{
        network: network,
        policy: &networkRetryPolicy{ maxRetries: maxRetires },
    }
}

func (networkRetryDecorator *NetworkRetryDecorator) Invoke(httpRequest *http.Request) (models.WebResponse, error) {
    // Handle as many Network related retries as we're allowed.
    var lastErr error
    for i := 0; i <= networkRetryDecorator.policy.maxRetries; i++ {
        ds3Response, err := networkRetryDecorator.network.Invoke(httpRequest)

        // If request was performed successfully then return response.
        if err == nil {
            return ds3Response, nil
        }

        // Log the network error, and try again if max retries has not been attempted
        log.Printf("Encountered network error '%s'.", err.Error())
        lastErr = err
    }

    // We had as many network related retries as we're allowed to use.
    return nil, errors.New(fmt.Sprintf(
        "Cannot send request. Retried the max number of %d times. Error: `%s`.",
        networkRetryDecorator.policy.maxRetries,
        lastErr.Error(),
    ))
}
