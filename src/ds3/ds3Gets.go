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
func (client *Client) ListMultiPartUploadParts(request *models.ListMultiPartUploadPartsRequest) (*models.ListMultiPartUploadPartsResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewListMultiPartUploadPartsResponse(response)
}
func (client *Client) ListMultiPartUploads(request *models.ListMultiPartUploadsRequest) (*models.ListMultiPartUploadsResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewListMultiPartUploadsResponse(response)
}
func (client *Client) GetBucketAclSpectraS3(request *models.GetBucketAclSpectraS3Request) (*models.GetBucketAclSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBucketAclSpectraS3Response(response)
}
func (client *Client) GetBucketAclsSpectraS3(request *models.GetBucketAclsSpectraS3Request) (*models.GetBucketAclsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBucketAclsSpectraS3Response(response)
}
func (client *Client) GetDataPolicyAclSpectraS3(request *models.GetDataPolicyAclSpectraS3Request) (*models.GetDataPolicyAclSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDataPolicyAclSpectraS3Response(response)
}
func (client *Client) GetDataPolicyAclsSpectraS3(request *models.GetDataPolicyAclsSpectraS3Request) (*models.GetDataPolicyAclsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDataPolicyAclsSpectraS3Response(response)
}
func (client *Client) GetBucketSpectraS3(request *models.GetBucketSpectraS3Request) (*models.GetBucketSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBucketSpectraS3Response(response)
}
func (client *Client) GetBucketsSpectraS3(request *models.GetBucketsSpectraS3Request) (*models.GetBucketsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBucketsSpectraS3Response(response)
}
func (client *Client) GetCacheFilesystemSpectraS3(request *models.GetCacheFilesystemSpectraS3Request) (*models.GetCacheFilesystemSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetCacheFilesystemSpectraS3Response(response)
}
func (client *Client) GetCacheFilesystemsSpectraS3(request *models.GetCacheFilesystemsSpectraS3Request) (*models.GetCacheFilesystemsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetCacheFilesystemsSpectraS3Response(response)
}
func (client *Client) GetCacheStateSpectraS3(request *models.GetCacheStateSpectraS3Request) (*models.GetCacheStateSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetCacheStateSpectraS3Response(response)
}
func (client *Client) GetBucketCapacitySummarySpectraS3(request *models.GetBucketCapacitySummarySpectraS3Request) (*models.GetBucketCapacitySummarySpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBucketCapacitySummarySpectraS3Response(response)
}
func (client *Client) GetStorageDomainCapacitySummarySpectraS3(request *models.GetStorageDomainCapacitySummarySpectraS3Request) (*models.GetStorageDomainCapacitySummarySpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetStorageDomainCapacitySummarySpectraS3Response(response)
}
func (client *Client) GetSystemCapacitySummarySpectraS3(request *models.GetSystemCapacitySummarySpectraS3Request) (*models.GetSystemCapacitySummarySpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSystemCapacitySummarySpectraS3Response(response)
}
func (client *Client) GetDataPathBackendSpectraS3(request *models.GetDataPathBackendSpectraS3Request) (*models.GetDataPathBackendSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDataPathBackendSpectraS3Response(response)
}
func (client *Client) GetDataPlannerBlobStoreTasksSpectraS3(request *models.GetDataPlannerBlobStoreTasksSpectraS3Request) (*models.GetDataPlannerBlobStoreTasksSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDataPlannerBlobStoreTasksSpectraS3Response(response)
}
func (client *Client) GetAzureDataReplicationRuleSpectraS3(request *models.GetAzureDataReplicationRuleSpectraS3Request) (*models.GetAzureDataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetAzureDataReplicationRuleSpectraS3Response(response)
}
func (client *Client) GetAzureDataReplicationRulesSpectraS3(request *models.GetAzureDataReplicationRulesSpectraS3Request) (*models.GetAzureDataReplicationRulesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetAzureDataReplicationRulesSpectraS3Response(response)
}
func (client *Client) GetDataPersistenceRuleSpectraS3(request *models.GetDataPersistenceRuleSpectraS3Request) (*models.GetDataPersistenceRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDataPersistenceRuleSpectraS3Response(response)
}
func (client *Client) GetDataPersistenceRulesSpectraS3(request *models.GetDataPersistenceRulesSpectraS3Request) (*models.GetDataPersistenceRulesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDataPersistenceRulesSpectraS3Response(response)
}
func (client *Client) GetDataPoliciesSpectraS3(request *models.GetDataPoliciesSpectraS3Request) (*models.GetDataPoliciesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDataPoliciesSpectraS3Response(response)
}
func (client *Client) GetDataPolicySpectraS3(request *models.GetDataPolicySpectraS3Request) (*models.GetDataPolicySpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDataPolicySpectraS3Response(response)
}
func (client *Client) GetDs3DataReplicationRuleSpectraS3(request *models.GetDs3DataReplicationRuleSpectraS3Request) (*models.GetDs3DataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDs3DataReplicationRuleSpectraS3Response(response)
}
func (client *Client) GetDs3DataReplicationRulesSpectraS3(request *models.GetDs3DataReplicationRulesSpectraS3Request) (*models.GetDs3DataReplicationRulesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDs3DataReplicationRulesSpectraS3Response(response)
}
func (client *Client) GetS3DataReplicationRuleSpectraS3(request *models.GetS3DataReplicationRuleSpectraS3Request) (*models.GetS3DataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetS3DataReplicationRuleSpectraS3Response(response)
}
func (client *Client) GetS3DataReplicationRulesSpectraS3(request *models.GetS3DataReplicationRulesSpectraS3Request) (*models.GetS3DataReplicationRulesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetS3DataReplicationRulesSpectraS3Response(response)
}
func (client *Client) GetDegradedAzureDataReplicationRulesSpectraS3(request *models.GetDegradedAzureDataReplicationRulesSpectraS3Request) (*models.GetDegradedAzureDataReplicationRulesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDegradedAzureDataReplicationRulesSpectraS3Response(response)
}
func (client *Client) GetDegradedBlobsSpectraS3(request *models.GetDegradedBlobsSpectraS3Request) (*models.GetDegradedBlobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDegradedBlobsSpectraS3Response(response)
}
func (client *Client) GetDegradedBucketsSpectraS3(request *models.GetDegradedBucketsSpectraS3Request) (*models.GetDegradedBucketsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDegradedBucketsSpectraS3Response(response)
}
func (client *Client) GetDegradedDataPersistenceRulesSpectraS3(request *models.GetDegradedDataPersistenceRulesSpectraS3Request) (*models.GetDegradedDataPersistenceRulesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDegradedDataPersistenceRulesSpectraS3Response(response)
}
func (client *Client) GetDegradedDs3DataReplicationRulesSpectraS3(request *models.GetDegradedDs3DataReplicationRulesSpectraS3Request) (*models.GetDegradedDs3DataReplicationRulesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDegradedDs3DataReplicationRulesSpectraS3Response(response)
}
func (client *Client) GetDegradedS3DataReplicationRulesSpectraS3(request *models.GetDegradedS3DataReplicationRulesSpectraS3Request) (*models.GetDegradedS3DataReplicationRulesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDegradedS3DataReplicationRulesSpectraS3Response(response)
}
func (client *Client) GetSuspectBlobAzureTargetsSpectraS3(request *models.GetSuspectBlobAzureTargetsSpectraS3Request) (*models.GetSuspectBlobAzureTargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSuspectBlobAzureTargetsSpectraS3Response(response)
}
func (client *Client) GetSuspectBlobDs3TargetsSpectraS3(request *models.GetSuspectBlobDs3TargetsSpectraS3Request) (*models.GetSuspectBlobDs3TargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSuspectBlobDs3TargetsSpectraS3Response(response)
}
func (client *Client) GetSuspectBlobPoolsSpectraS3(request *models.GetSuspectBlobPoolsSpectraS3Request) (*models.GetSuspectBlobPoolsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSuspectBlobPoolsSpectraS3Response(response)
}
func (client *Client) GetSuspectBlobS3TargetsSpectraS3(request *models.GetSuspectBlobS3TargetsSpectraS3Request) (*models.GetSuspectBlobS3TargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSuspectBlobS3TargetsSpectraS3Response(response)
}
func (client *Client) GetSuspectBlobTapesSpectraS3(request *models.GetSuspectBlobTapesSpectraS3Request) (*models.GetSuspectBlobTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSuspectBlobTapesSpectraS3Response(response)
}
func (client *Client) GetSuspectBucketsSpectraS3(request *models.GetSuspectBucketsSpectraS3Request) (*models.GetSuspectBucketsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSuspectBucketsSpectraS3Response(response)
}
func (client *Client) GetSuspectObjectsSpectraS3(request *models.GetSuspectObjectsSpectraS3Request) (*models.GetSuspectObjectsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSuspectObjectsSpectraS3Response(response)
}
func (client *Client) GetSuspectObjectsWithFullDetailsSpectraS3(request *models.GetSuspectObjectsWithFullDetailsSpectraS3Request) (*models.GetSuspectObjectsWithFullDetailsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSuspectObjectsWithFullDetailsSpectraS3Response(response)
}
func (client *Client) GetGroupMemberSpectraS3(request *models.GetGroupMemberSpectraS3Request) (*models.GetGroupMemberSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetGroupMemberSpectraS3Response(response)
}
func (client *Client) GetGroupMembersSpectraS3(request *models.GetGroupMembersSpectraS3Request) (*models.GetGroupMembersSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetGroupMembersSpectraS3Response(response)
}
func (client *Client) GetGroupSpectraS3(request *models.GetGroupSpectraS3Request) (*models.GetGroupSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetGroupSpectraS3Response(response)
}
func (client *Client) GetGroupsSpectraS3(request *models.GetGroupsSpectraS3Request) (*models.GetGroupsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetGroupsSpectraS3Response(response)
}
func (client *Client) GetActiveJobSpectraS3(request *models.GetActiveJobSpectraS3Request) (*models.GetActiveJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetActiveJobSpectraS3Response(response)
}
func (client *Client) GetActiveJobsSpectraS3(request *models.GetActiveJobsSpectraS3Request) (*models.GetActiveJobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetActiveJobsSpectraS3Response(response)
}
func (client *Client) GetCanceledJobSpectraS3(request *models.GetCanceledJobSpectraS3Request) (*models.GetCanceledJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetCanceledJobSpectraS3Response(response)
}
func (client *Client) GetCanceledJobsSpectraS3(request *models.GetCanceledJobsSpectraS3Request) (*models.GetCanceledJobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetCanceledJobsSpectraS3Response(response)
}
func (client *Client) GetCompletedJobSpectraS3(request *models.GetCompletedJobSpectraS3Request) (*models.GetCompletedJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetCompletedJobSpectraS3Response(response)
}
func (client *Client) GetCompletedJobsSpectraS3(request *models.GetCompletedJobsSpectraS3Request) (*models.GetCompletedJobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetCompletedJobsSpectraS3Response(response)
}
func (client *Client) GetJobChunkDaoSpectraS3(request *models.GetJobChunkDaoSpectraS3Request) (*models.GetJobChunkDaoSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobChunkDaoSpectraS3Response(response)
}
func (client *Client) GetJobChunkSpectraS3(request *models.GetJobChunkSpectraS3Request) (*models.GetJobChunkSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobChunkSpectraS3Response(response)
}
func (client *Client) GetJobChunksReadyForClientProcessingSpectraS3(request *models.GetJobChunksReadyForClientProcessingSpectraS3Request) (*models.GetJobChunksReadyForClientProcessingSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobChunksReadyForClientProcessingSpectraS3Response(response)
}
func (client *Client) GetJobSpectraS3(request *models.GetJobSpectraS3Request) (*models.GetJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobSpectraS3Response(response)
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
func (client *Client) GetJobsSpectraS3(request *models.GetJobsSpectraS3Request) (*models.GetJobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobsSpectraS3Response(response)
}
func (client *Client) GetNodeSpectraS3(request *models.GetNodeSpectraS3Request) (*models.GetNodeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetNodeSpectraS3Response(response)
}
func (client *Client) GetNodesSpectraS3(request *models.GetNodesSpectraS3Request) (*models.GetNodesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetNodesSpectraS3Response(response)
}
func (client *Client) GetAzureTargetFailureNotificationRegistrationSpectraS3(request *models.GetAzureTargetFailureNotificationRegistrationSpectraS3Request) (*models.GetAzureTargetFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetAzureTargetFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetAzureTargetFailureNotificationRegistrationsSpectraS3(request *models.GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) (*models.GetAzureTargetFailureNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetAzureTargetFailureNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetDs3TargetFailureNotificationRegistrationSpectraS3(request *models.GetDs3TargetFailureNotificationRegistrationSpectraS3Request) (*models.GetDs3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDs3TargetFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetDs3TargetFailureNotificationRegistrationsSpectraS3(request *models.GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) (*models.GetDs3TargetFailureNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDs3TargetFailureNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetJobCompletedNotificationRegistrationSpectraS3(request *models.GetJobCompletedNotificationRegistrationSpectraS3Request) (*models.GetJobCompletedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobCompletedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetJobCompletedNotificationRegistrationsSpectraS3(request *models.GetJobCompletedNotificationRegistrationsSpectraS3Request) (*models.GetJobCompletedNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobCompletedNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetJobCreatedNotificationRegistrationSpectraS3(request *models.GetJobCreatedNotificationRegistrationSpectraS3Request) (*models.GetJobCreatedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobCreatedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetJobCreatedNotificationRegistrationsSpectraS3(request *models.GetJobCreatedNotificationRegistrationsSpectraS3Request) (*models.GetJobCreatedNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobCreatedNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetJobCreationFailedNotificationRegistrationSpectraS3(request *models.GetJobCreationFailedNotificationRegistrationSpectraS3Request) (*models.GetJobCreationFailedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobCreationFailedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetJobCreationFailedNotificationRegistrationsSpectraS3(request *models.GetJobCreationFailedNotificationRegistrationsSpectraS3Request) (*models.GetJobCreationFailedNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetJobCreationFailedNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetObjectCachedNotificationRegistrationSpectraS3(request *models.GetObjectCachedNotificationRegistrationSpectraS3Request) (*models.GetObjectCachedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectCachedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetObjectCachedNotificationRegistrationsSpectraS3(request *models.GetObjectCachedNotificationRegistrationsSpectraS3Request) (*models.GetObjectCachedNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectCachedNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetObjectLostNotificationRegistrationSpectraS3(request *models.GetObjectLostNotificationRegistrationSpectraS3Request) (*models.GetObjectLostNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectLostNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetObjectLostNotificationRegistrationsSpectraS3(request *models.GetObjectLostNotificationRegistrationsSpectraS3Request) (*models.GetObjectLostNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectLostNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetObjectPersistedNotificationRegistrationSpectraS3(request *models.GetObjectPersistedNotificationRegistrationSpectraS3Request) (*models.GetObjectPersistedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectPersistedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetObjectPersistedNotificationRegistrationsSpectraS3(request *models.GetObjectPersistedNotificationRegistrationsSpectraS3Request) (*models.GetObjectPersistedNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectPersistedNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetPoolFailureNotificationRegistrationSpectraS3(request *models.GetPoolFailureNotificationRegistrationSpectraS3Request) (*models.GetPoolFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetPoolFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetPoolFailureNotificationRegistrationsSpectraS3(request *models.GetPoolFailureNotificationRegistrationsSpectraS3Request) (*models.GetPoolFailureNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetPoolFailureNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetS3TargetFailureNotificationRegistrationSpectraS3(request *models.GetS3TargetFailureNotificationRegistrationSpectraS3Request) (*models.GetS3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetS3TargetFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetS3TargetFailureNotificationRegistrationsSpectraS3(request *models.GetS3TargetFailureNotificationRegistrationsSpectraS3Request) (*models.GetS3TargetFailureNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetS3TargetFailureNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetStorageDomainFailureNotificationRegistrationSpectraS3(request *models.GetStorageDomainFailureNotificationRegistrationSpectraS3Request) (*models.GetStorageDomainFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetStorageDomainFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetStorageDomainFailureNotificationRegistrationsSpectraS3(request *models.GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) (*models.GetStorageDomainFailureNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetStorageDomainFailureNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetSystemFailureNotificationRegistrationSpectraS3(request *models.GetSystemFailureNotificationRegistrationSpectraS3Request) (*models.GetSystemFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSystemFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetSystemFailureNotificationRegistrationsSpectraS3(request *models.GetSystemFailureNotificationRegistrationsSpectraS3Request) (*models.GetSystemFailureNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSystemFailureNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetTapeFailureNotificationRegistrationSpectraS3(request *models.GetTapeFailureNotificationRegistrationSpectraS3Request) (*models.GetTapeFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapeFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetTapeFailureNotificationRegistrationsSpectraS3(request *models.GetTapeFailureNotificationRegistrationsSpectraS3Request) (*models.GetTapeFailureNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapeFailureNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetTapePartitionFailureNotificationRegistrationSpectraS3(request *models.GetTapePartitionFailureNotificationRegistrationSpectraS3Request) (*models.GetTapePartitionFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapePartitionFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) GetTapePartitionFailureNotificationRegistrationsSpectraS3(request *models.GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) (*models.GetTapePartitionFailureNotificationRegistrationsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapePartitionFailureNotificationRegistrationsSpectraS3Response(response)
}
func (client *Client) GetBlobPersistenceSpectraS3(request *models.GetBlobPersistenceSpectraS3Request) (*models.GetBlobPersistenceSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBlobPersistenceSpectraS3Response(response)
}
func (client *Client) GetObjectDetailsSpectraS3(request *models.GetObjectDetailsSpectraS3Request) (*models.GetObjectDetailsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectDetailsSpectraS3Response(response)
}
func (client *Client) GetObjectsDetailsSpectraS3(request *models.GetObjectsDetailsSpectraS3Request) (*models.GetObjectsDetailsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectsDetailsSpectraS3Response(response)
}
func (client *Client) GetObjectsWithFullDetailsSpectraS3(request *models.GetObjectsWithFullDetailsSpectraS3Request) (*models.GetObjectsWithFullDetailsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetObjectsWithFullDetailsSpectraS3Response(response)
}
func (client *Client) VerifyPhysicalPlacementForObjectsSpectraS3(request *models.VerifyPhysicalPlacementForObjectsSpectraS3Request) (*models.VerifyPhysicalPlacementForObjectsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyPhysicalPlacementForObjectsSpectraS3Response(response)
}
func (client *Client) VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3(request *models.VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) (*models.VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response(response)
}
func (client *Client) GetBlobsOnPoolSpectraS3(request *models.GetBlobsOnPoolSpectraS3Request) (*models.GetBlobsOnPoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBlobsOnPoolSpectraS3Response(response)
}
func (client *Client) GetPoolFailuresSpectraS3(request *models.GetPoolFailuresSpectraS3Request) (*models.GetPoolFailuresSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetPoolFailuresSpectraS3Response(response)
}
func (client *Client) GetPoolPartitionSpectraS3(request *models.GetPoolPartitionSpectraS3Request) (*models.GetPoolPartitionSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetPoolPartitionSpectraS3Response(response)
}
func (client *Client) GetPoolPartitionsSpectraS3(request *models.GetPoolPartitionsSpectraS3Request) (*models.GetPoolPartitionsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetPoolPartitionsSpectraS3Response(response)
}
func (client *Client) GetPoolSpectraS3(request *models.GetPoolSpectraS3Request) (*models.GetPoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetPoolSpectraS3Response(response)
}
func (client *Client) GetPoolsSpectraS3(request *models.GetPoolsSpectraS3Request) (*models.GetPoolsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetPoolsSpectraS3Response(response)
}
func (client *Client) GetStorageDomainFailuresSpectraS3(request *models.GetStorageDomainFailuresSpectraS3Request) (*models.GetStorageDomainFailuresSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetStorageDomainFailuresSpectraS3Response(response)
}
func (client *Client) GetStorageDomainMemberSpectraS3(request *models.GetStorageDomainMemberSpectraS3Request) (*models.GetStorageDomainMemberSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetStorageDomainMemberSpectraS3Response(response)
}
func (client *Client) GetStorageDomainMembersSpectraS3(request *models.GetStorageDomainMembersSpectraS3Request) (*models.GetStorageDomainMembersSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetStorageDomainMembersSpectraS3Response(response)
}
func (client *Client) GetStorageDomainSpectraS3(request *models.GetStorageDomainSpectraS3Request) (*models.GetStorageDomainSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetStorageDomainSpectraS3Response(response)
}
func (client *Client) GetStorageDomainsSpectraS3(request *models.GetStorageDomainsSpectraS3Request) (*models.GetStorageDomainsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetStorageDomainsSpectraS3Response(response)
}
func (client *Client) GetFeatureKeysSpectraS3(request *models.GetFeatureKeysSpectraS3Request) (*models.GetFeatureKeysSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetFeatureKeysSpectraS3Response(response)
}
func (client *Client) GetSystemFailuresSpectraS3(request *models.GetSystemFailuresSpectraS3Request) (*models.GetSystemFailuresSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSystemFailuresSpectraS3Response(response)
}
func (client *Client) GetSystemInformationSpectraS3(request *models.GetSystemInformationSpectraS3Request) (*models.GetSystemInformationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetSystemInformationSpectraS3Response(response)
}
func (client *Client) VerifySystemHealthSpectraS3(request *models.VerifySystemHealthSpectraS3Request) (*models.VerifySystemHealthSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifySystemHealthSpectraS3Response(response)
}
func (client *Client) GetBlobsOnTapeSpectraS3(request *models.GetBlobsOnTapeSpectraS3Request) (*models.GetBlobsOnTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBlobsOnTapeSpectraS3Response(response)
}
func (client *Client) GetTapeDensityDirectiveSpectraS3(request *models.GetTapeDensityDirectiveSpectraS3Request) (*models.GetTapeDensityDirectiveSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapeDensityDirectiveSpectraS3Response(response)
}
func (client *Client) GetTapeDensityDirectivesSpectraS3(request *models.GetTapeDensityDirectivesSpectraS3Request) (*models.GetTapeDensityDirectivesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapeDensityDirectivesSpectraS3Response(response)
}
func (client *Client) GetTapeDriveSpectraS3(request *models.GetTapeDriveSpectraS3Request) (*models.GetTapeDriveSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapeDriveSpectraS3Response(response)
}
func (client *Client) GetTapeDrivesSpectraS3(request *models.GetTapeDrivesSpectraS3Request) (*models.GetTapeDrivesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapeDrivesSpectraS3Response(response)
}
func (client *Client) GetTapeFailuresSpectraS3(request *models.GetTapeFailuresSpectraS3Request) (*models.GetTapeFailuresSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapeFailuresSpectraS3Response(response)
}
func (client *Client) GetTapeLibrariesSpectraS3(request *models.GetTapeLibrariesSpectraS3Request) (*models.GetTapeLibrariesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapeLibrariesSpectraS3Response(response)
}
func (client *Client) GetTapeLibrarySpectraS3(request *models.GetTapeLibrarySpectraS3Request) (*models.GetTapeLibrarySpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapeLibrarySpectraS3Response(response)
}
func (client *Client) GetTapePartitionFailuresSpectraS3(request *models.GetTapePartitionFailuresSpectraS3Request) (*models.GetTapePartitionFailuresSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapePartitionFailuresSpectraS3Response(response)
}
func (client *Client) GetTapePartitionSpectraS3(request *models.GetTapePartitionSpectraS3Request) (*models.GetTapePartitionSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapePartitionSpectraS3Response(response)
}
func (client *Client) GetTapePartitionWithFullDetailsSpectraS3(request *models.GetTapePartitionWithFullDetailsSpectraS3Request) (*models.GetTapePartitionWithFullDetailsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapePartitionWithFullDetailsSpectraS3Response(response)
}
func (client *Client) GetTapePartitionsSpectraS3(request *models.GetTapePartitionsSpectraS3Request) (*models.GetTapePartitionsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapePartitionsSpectraS3Response(response)
}
func (client *Client) GetTapePartitionsWithFullDetailsSpectraS3(request *models.GetTapePartitionsWithFullDetailsSpectraS3Request) (*models.GetTapePartitionsWithFullDetailsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapePartitionsWithFullDetailsSpectraS3Response(response)
}
func (client *Client) GetTapeSpectraS3(request *models.GetTapeSpectraS3Request) (*models.GetTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapeSpectraS3Response(response)
}
func (client *Client) GetTapesSpectraS3(request *models.GetTapesSpectraS3Request) (*models.GetTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetTapesSpectraS3Response(response)
}
func (client *Client) GetAzureTargetBucketNamesSpectraS3(request *models.GetAzureTargetBucketNamesSpectraS3Request) (*models.GetAzureTargetBucketNamesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetAzureTargetBucketNamesSpectraS3Response(response)
}
func (client *Client) GetAzureTargetFailuresSpectraS3(request *models.GetAzureTargetFailuresSpectraS3Request) (*models.GetAzureTargetFailuresSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetAzureTargetFailuresSpectraS3Response(response)
}
func (client *Client) GetAzureTargetReadPreferenceSpectraS3(request *models.GetAzureTargetReadPreferenceSpectraS3Request) (*models.GetAzureTargetReadPreferenceSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetAzureTargetReadPreferenceSpectraS3Response(response)
}
func (client *Client) GetAzureTargetReadPreferencesSpectraS3(request *models.GetAzureTargetReadPreferencesSpectraS3Request) (*models.GetAzureTargetReadPreferencesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetAzureTargetReadPreferencesSpectraS3Response(response)
}
func (client *Client) GetAzureTargetSpectraS3(request *models.GetAzureTargetSpectraS3Request) (*models.GetAzureTargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetAzureTargetSpectraS3Response(response)
}
func (client *Client) GetAzureTargetsSpectraS3(request *models.GetAzureTargetsSpectraS3Request) (*models.GetAzureTargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetAzureTargetsSpectraS3Response(response)
}
func (client *Client) GetBlobsOnAzureTargetSpectraS3(request *models.GetBlobsOnAzureTargetSpectraS3Request) (*models.GetBlobsOnAzureTargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBlobsOnAzureTargetSpectraS3Response(response)
}
func (client *Client) GetBlobsOnDs3TargetSpectraS3(request *models.GetBlobsOnDs3TargetSpectraS3Request) (*models.GetBlobsOnDs3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBlobsOnDs3TargetSpectraS3Response(response)
}
func (client *Client) GetDs3TargetDataPoliciesSpectraS3(request *models.GetDs3TargetDataPoliciesSpectraS3Request) (*models.GetDs3TargetDataPoliciesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDs3TargetDataPoliciesSpectraS3Response(response)
}
func (client *Client) GetDs3TargetFailuresSpectraS3(request *models.GetDs3TargetFailuresSpectraS3Request) (*models.GetDs3TargetFailuresSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDs3TargetFailuresSpectraS3Response(response)
}
func (client *Client) GetDs3TargetReadPreferenceSpectraS3(request *models.GetDs3TargetReadPreferenceSpectraS3Request) (*models.GetDs3TargetReadPreferenceSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDs3TargetReadPreferenceSpectraS3Response(response)
}
func (client *Client) GetDs3TargetReadPreferencesSpectraS3(request *models.GetDs3TargetReadPreferencesSpectraS3Request) (*models.GetDs3TargetReadPreferencesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDs3TargetReadPreferencesSpectraS3Response(response)
}
func (client *Client) GetDs3TargetSpectraS3(request *models.GetDs3TargetSpectraS3Request) (*models.GetDs3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDs3TargetSpectraS3Response(response)
}
func (client *Client) GetDs3TargetsSpectraS3(request *models.GetDs3TargetsSpectraS3Request) (*models.GetDs3TargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetDs3TargetsSpectraS3Response(response)
}
func (client *Client) GetBlobsOnS3TargetSpectraS3(request *models.GetBlobsOnS3TargetSpectraS3Request) (*models.GetBlobsOnS3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBlobsOnS3TargetSpectraS3Response(response)
}
func (client *Client) GetS3TargetBucketNamesSpectraS3(request *models.GetS3TargetBucketNamesSpectraS3Request) (*models.GetS3TargetBucketNamesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetS3TargetBucketNamesSpectraS3Response(response)
}
func (client *Client) GetS3TargetFailuresSpectraS3(request *models.GetS3TargetFailuresSpectraS3Request) (*models.GetS3TargetFailuresSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetS3TargetFailuresSpectraS3Response(response)
}
func (client *Client) GetS3TargetReadPreferenceSpectraS3(request *models.GetS3TargetReadPreferenceSpectraS3Request) (*models.GetS3TargetReadPreferenceSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetS3TargetReadPreferenceSpectraS3Response(response)
}
func (client *Client) GetS3TargetReadPreferencesSpectraS3(request *models.GetS3TargetReadPreferencesSpectraS3Request) (*models.GetS3TargetReadPreferencesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetS3TargetReadPreferencesSpectraS3Response(response)
}
func (client *Client) GetS3TargetSpectraS3(request *models.GetS3TargetSpectraS3Request) (*models.GetS3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetS3TargetSpectraS3Response(response)
}
func (client *Client) GetS3TargetsSpectraS3(request *models.GetS3TargetsSpectraS3Request) (*models.GetS3TargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetS3TargetsSpectraS3Response(response)
}
func (client *Client) GetUserSpectraS3(request *models.GetUserSpectraS3Request) (*models.GetUserSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetUserSpectraS3Response(response)
}
func (client *Client) GetUsersSpectraS3(request *models.GetUsersSpectraS3Request) (*models.GetUsersSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)
    httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(&networkRetryDecorator, client.clientPolicy.maxRedirect)

    // Invoke the HTTP request.
    response, requestErr := httpRedirectDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetUsersSpectraS3Response(response)
}
