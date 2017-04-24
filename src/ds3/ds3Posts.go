package ds3

import (
    "ds3/models"
    "ds3/networking"
)

func (client *Client) InitiateMultipart(request *models.InitiateMultipartRequest) (*models.InitiateMultipartResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewInitiateMultipartResponse(response)
}

func (client *Client) CompleteMultipart(request *models.CompleteMultipartRequest) (*models.CompleteMultipartResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCompleteMultipartResponse(response)
}

