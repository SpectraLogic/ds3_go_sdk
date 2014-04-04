package ds3

import "ds3/models"

func (client *Client) DeleteBucket(request *models.DeleteBucketRequest) (*models.DeleteBucketResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteBucketResponse(response)
}

func (client *Client) DeleteObject(request *models.DeleteObjectRequest) (*models.DeleteObjectResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteObjectResponse(response)
}

