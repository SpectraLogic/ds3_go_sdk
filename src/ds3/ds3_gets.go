package ds3

import "ds3/models"

func (client *Client) GetService(request *models.GetServiceRequest) (*models.GetServiceResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetServiceResponse(response)
}

func (client *Client) GetBucket(request *models.GetBucketRequest) (*models.GetBucketResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBucketResponse(response)
}

func (client *Client) GetObject(request *models.GetObjectRequest) (*models.GetObjectResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectResponse(response)
}

func (client *Client) ListParts(request *models.ListPartsRequest) (*models.ListPartsResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewListPartsResponse(response)
}

