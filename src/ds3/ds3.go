package ds3

import (
    "ds3/net"
    "ds3/models"
)

type Client struct {
    netLayer net.Network
}

func (client *Client) GetService(request *models.GetServiceRequest) *models.GetServiceResponse {
    return NewServiceResponse(client.netLayer.Invoke(request))
}

