package ds3

import (
    "ds3/net"
    "ds3/models"
)

type Client struct {
    netLayer net.Network
}

//TODO: improve error handling
func (client *Client) GetService(request *models.GetServiceRequest) (*models.GetServiceResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    getServiceResponse, responseError := models.NewGetServiceResponse(response)
    if responseError != nil {
        return nil, responseError
    }

    // Return the result.
    return getServiceResponse, nil
}

