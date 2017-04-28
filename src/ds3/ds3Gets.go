package ds3

import (
    "ds3/models"
    networking "ds3/networking"
)

func (client *Client) GetService(request *models.GetServiceRequest) (*models.GetServiceResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetServiceResponse(response)
}

func (client *Client) GetBucket(request *models.GetBucketRequest) (*models.GetBucketResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBucketResponse(response)
}

func (client *Client) GetObject(request *models.GetObjectRequest) (*models.GetObjectResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectResponse(response)
}

func (client *Client) ListParts(request *models.ListPartsRequest) (*models.ListPartsResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewListPartsResponse(response)
}

func (client *Client) ListMultipart(request *models.ListMultipartRequest) (*models.ListMultipartResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewListMultipartResponse(response)
}

func (client *Client) GetJobToReplicateSpectraS3(request *models.GetJobToReplicateSpectraS3Request) (*models.GetJobToReplicateSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobToReplicateSpectraS3Response(response)
}

