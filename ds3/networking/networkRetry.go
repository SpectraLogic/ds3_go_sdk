package networking

import (
	"errors"
	"fmt"
	"github.com/SpectraLogic/ds3_go_sdk/ds3/models"
	"log"
	"net/http"
)

type networkRetryPolicy struct {
	maxRetries int
}

// Decorator for Network which handles network related retries
type NetworkRetryDecorator struct {
	network Network
	policy  *networkRetryPolicy
}

func NewNetworkRetryDecorator(network Network, maxRetires int) Network {
	return &NetworkRetryDecorator{
		network: network,
		policy:  &networkRetryPolicy{maxRetries: maxRetires},
	}
}

func (networkRetryDecorator *NetworkRetryDecorator) Invoke(httpRequest *http.Request) (models.WebResponse, error) {
	// Handle as many Network related retries as we're allowed.
	var lastErr error
	for i := 0; i <= networkRetryDecorator.policy.maxRetries; i++ {
		// Bail early if the request's context has been cancelled or its deadline exceeded
		// — otherwise we'd keep slamming attempts at an already-dead context.
		if err := httpRequest.Context().Err(); err != nil {
			return nil, err
		}

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
		"Cannot send request. Retried the network connection the maximum number of %d times. Error: `%s`.",
		networkRetryDecorator.policy.maxRetries,
		lastErr.Error(),
	))
}
