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


