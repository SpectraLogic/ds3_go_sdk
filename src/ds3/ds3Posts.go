package ds3

import "ds3/models"

func (client *Client) InitiateMultipart(request *models.InitiateMultipartRequest) (*models.InitiateMultipartResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewInitiateMultipartResponse(response)
}

func (client *Client) CompleteMultipart(request *models.CompleteMultipartRequest) (*models.CompleteMultipartResponse, error) {
    // Invoke the HTTP request.
    response, requestErr := client.netLayer.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCompleteMultipartResponse(response)
}

