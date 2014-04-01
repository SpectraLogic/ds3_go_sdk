package ds3

import (
    "ds3/net"
    "ds3/models"
)

type Client struct {
    netLayer net.Network
}

func (client *Client) GetService(request *models.GetServiceRequest) (*models.GetServiceResponse, error) {
    //TODO: error handling
    response, _ := client.netLayer.Invoke(request)
    getServiceResponse, _ := models.NewServiceResponse(response)
    return getServiceResponse, nil
}

