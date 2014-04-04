package ds3

import "ds3/models"

func (client *Client) PutBucket(request *models.PutBucketRequest) (*models.PutBucketResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutBucketResponse(response)
}

func (client *Client) PutObject(request *models.PutObjectRequest) (*models.PutObjectResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutObjectResponse(response)
}

func (client *Client) BulkGet(request *models.BulkGetRequest) (*models.BulkGetResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewBulkGetResponse(response)
}

func (client *Client) BulkPut(request *models.BulkPutRequest) (*models.BulkPutResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewBulkPutResponse(response)
}

