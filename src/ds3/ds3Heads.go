// Copyright 2014-2017 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// This code is auto-generated, do not modify

package ds3

import (
    "ds3/models"
    "ds3/networking"
)

func (client *Client) HeadBucket(request *models.HeadBucketRequest) (*models.HeadBucketResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewHeadBucketResponse(response)
}
func (client *Client) HeadObject(request *models.HeadObjectRequest) (*models.HeadObjectResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewHeadObjectResponse(response)
}

