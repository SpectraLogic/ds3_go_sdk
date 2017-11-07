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

func (client *Client) CompleteMultiPartUpload(request *models.CompleteMultiPartUploadRequest) (*models.CompleteMultiPartUploadResponse, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/" + request.BucketName + "/" + request.ObjectName).
        WithQueryParam("upload_id", request.UploadId).
        WithReadCloser(buildPartsListStream(request.Parts)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCompleteMultiPartUploadResponse(response)
}

func (client *Client) DeleteObjects(request *models.DeleteObjectsRequest) (*models.DeleteObjectsResponse, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/" + request.BucketName).
        WithQueryParam("delete", "").
        WithOptionalVoidQueryParam("roll_back", request.RollBack).
        WithReadCloser(buildDeleteObjectsPayload(request.ObjectNames)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteObjectsResponse(response)
}

func (client *Client) InitiateMultiPartUpload(request *models.InitiateMultiPartUploadRequest) (*models.InitiateMultiPartUploadResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewInitiateMultiPartUploadResponse(response)
}
func (client *Client) PutBucketAclForGroupSpectraS3(request *models.PutBucketAclForGroupSpectraS3Request) (*models.PutBucketAclForGroupSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutBucketAclForGroupSpectraS3Response(response)
}
func (client *Client) PutBucketAclForUserSpectraS3(request *models.PutBucketAclForUserSpectraS3Request) (*models.PutBucketAclForUserSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutBucketAclForUserSpectraS3Response(response)
}
func (client *Client) PutDataPolicyAclForGroupSpectraS3(request *models.PutDataPolicyAclForGroupSpectraS3Request) (*models.PutDataPolicyAclForGroupSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDataPolicyAclForGroupSpectraS3Response(response)
}
func (client *Client) PutDataPolicyAclForUserSpectraS3(request *models.PutDataPolicyAclForUserSpectraS3Request) (*models.PutDataPolicyAclForUserSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDataPolicyAclForUserSpectraS3Response(response)
}
func (client *Client) PutGlobalBucketAclForGroupSpectraS3(request *models.PutGlobalBucketAclForGroupSpectraS3Request) (*models.PutGlobalBucketAclForGroupSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGlobalBucketAclForGroupSpectraS3Response(response)
}
func (client *Client) PutGlobalBucketAclForUserSpectraS3(request *models.PutGlobalBucketAclForUserSpectraS3Request) (*models.PutGlobalBucketAclForUserSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGlobalBucketAclForUserSpectraS3Response(response)
}
func (client *Client) PutGlobalDataPolicyAclForGroupSpectraS3(request *models.PutGlobalDataPolicyAclForGroupSpectraS3Request) (*models.PutGlobalDataPolicyAclForGroupSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGlobalDataPolicyAclForGroupSpectraS3Response(response)
}
func (client *Client) PutGlobalDataPolicyAclForUserSpectraS3(request *models.PutGlobalDataPolicyAclForUserSpectraS3Request) (*models.PutGlobalDataPolicyAclForUserSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGlobalDataPolicyAclForUserSpectraS3Response(response)
}
func (client *Client) PutBucketSpectraS3(request *models.PutBucketSpectraS3Request) (*models.PutBucketSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutBucketSpectraS3Response(response)
}
func (client *Client) PutAzureDataReplicationRuleSpectraS3(request *models.PutAzureDataReplicationRuleSpectraS3Request) (*models.PutAzureDataReplicationRuleSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/azure_data_replication_rule").
        WithQueryParam("data_policy_id", request.DataPolicyId).
        WithQueryParam("target_id", request.TargetId).
        WithQueryParam("type", request.DataReplicationRuleType.String()).
        WithOptionalQueryParam("max_blob_part_size_in_bytes", networking.Int64PtrToStrPtr(request.MaxBlobPartSizeInBytes)).
        WithOptionalQueryParam("replicate_deletes", networking.BoolPtrToStrPtr(request.ReplicateDeletes)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutAzureDataReplicationRuleSpectraS3Response(response)
}
func (client *Client) PutDataPersistenceRuleSpectraS3(request *models.PutDataPersistenceRuleSpectraS3Request) (*models.PutDataPersistenceRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDataPersistenceRuleSpectraS3Response(response)
}
func (client *Client) PutDataPolicySpectraS3(request *models.PutDataPolicySpectraS3Request) (*models.PutDataPolicySpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDataPolicySpectraS3Response(response)
}
func (client *Client) PutDs3DataReplicationRuleSpectraS3(request *models.PutDs3DataReplicationRuleSpectraS3Request) (*models.PutDs3DataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDs3DataReplicationRuleSpectraS3Response(response)
}
func (client *Client) PutS3DataReplicationRuleSpectraS3(request *models.PutS3DataReplicationRuleSpectraS3Request) (*models.PutS3DataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutS3DataReplicationRuleSpectraS3Response(response)
}
func (client *Client) PutGroupGroupMemberSpectraS3(request *models.PutGroupGroupMemberSpectraS3Request) (*models.PutGroupGroupMemberSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGroupGroupMemberSpectraS3Response(response)
}
func (client *Client) PutGroupSpectraS3(request *models.PutGroupSpectraS3Request) (*models.PutGroupSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGroupSpectraS3Response(response)
}
func (client *Client) PutUserGroupMemberSpectraS3(request *models.PutUserGroupMemberSpectraS3Request) (*models.PutUserGroupMemberSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutUserGroupMemberSpectraS3Response(response)
}
func (client *Client) PutAzureTargetFailureNotificationRegistrationSpectraS3(request *models.PutAzureTargetFailureNotificationRegistrationSpectraS3Request) (*models.PutAzureTargetFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutAzureTargetFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutDs3TargetFailureNotificationRegistrationSpectraS3(request *models.PutDs3TargetFailureNotificationRegistrationSpectraS3Request) (*models.PutDs3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDs3TargetFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutJobCompletedNotificationRegistrationSpectraS3(request *models.PutJobCompletedNotificationRegistrationSpectraS3Request) (*models.PutJobCompletedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutJobCompletedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutJobCreatedNotificationRegistrationSpectraS3(request *models.PutJobCreatedNotificationRegistrationSpectraS3Request) (*models.PutJobCreatedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutJobCreatedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutJobCreationFailedNotificationRegistrationSpectraS3(request *models.PutJobCreationFailedNotificationRegistrationSpectraS3Request) (*models.PutJobCreationFailedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutJobCreationFailedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutObjectCachedNotificationRegistrationSpectraS3(request *models.PutObjectCachedNotificationRegistrationSpectraS3Request) (*models.PutObjectCachedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutObjectCachedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutObjectLostNotificationRegistrationSpectraS3(request *models.PutObjectLostNotificationRegistrationSpectraS3Request) (*models.PutObjectLostNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutObjectLostNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutObjectPersistedNotificationRegistrationSpectraS3(request *models.PutObjectPersistedNotificationRegistrationSpectraS3Request) (*models.PutObjectPersistedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutObjectPersistedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutPoolFailureNotificationRegistrationSpectraS3(request *models.PutPoolFailureNotificationRegistrationSpectraS3Request) (*models.PutPoolFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutPoolFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutS3TargetFailureNotificationRegistrationSpectraS3(request *models.PutS3TargetFailureNotificationRegistrationSpectraS3Request) (*models.PutS3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutS3TargetFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutStorageDomainFailureNotificationRegistrationSpectraS3(request *models.PutStorageDomainFailureNotificationRegistrationSpectraS3Request) (*models.PutStorageDomainFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutStorageDomainFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutSystemFailureNotificationRegistrationSpectraS3(request *models.PutSystemFailureNotificationRegistrationSpectraS3Request) (*models.PutSystemFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutSystemFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutTapeFailureNotificationRegistrationSpectraS3(request *models.PutTapeFailureNotificationRegistrationSpectraS3Request) (*models.PutTapeFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutTapeFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutTapePartitionFailureNotificationRegistrationSpectraS3(request *models.PutTapePartitionFailureNotificationRegistrationSpectraS3Request) (*models.PutTapePartitionFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutTapePartitionFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) PutPoolPartitionSpectraS3(request *models.PutPoolPartitionSpectraS3Request) (*models.PutPoolPartitionSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutPoolPartitionSpectraS3Response(response)
}
func (client *Client) PutPoolStorageDomainMemberSpectraS3(request *models.PutPoolStorageDomainMemberSpectraS3Request) (*models.PutPoolStorageDomainMemberSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutPoolStorageDomainMemberSpectraS3Response(response)
}
func (client *Client) PutStorageDomainSpectraS3(request *models.PutStorageDomainSpectraS3Request) (*models.PutStorageDomainSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutStorageDomainSpectraS3Response(response)
}
func (client *Client) PutTapeStorageDomainMemberSpectraS3(request *models.PutTapeStorageDomainMemberSpectraS3Request) (*models.PutTapeStorageDomainMemberSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutTapeStorageDomainMemberSpectraS3Response(response)
}
func (client *Client) PutTapeDensityDirectiveSpectraS3(request *models.PutTapeDensityDirectiveSpectraS3Request) (*models.PutTapeDensityDirectiveSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutTapeDensityDirectiveSpectraS3Response(response)
}
func (client *Client) PutAzureTargetBucketNameSpectraS3(request *models.PutAzureTargetBucketNameSpectraS3Request) (*models.PutAzureTargetBucketNameSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutAzureTargetBucketNameSpectraS3Response(response)
}
func (client *Client) PutAzureTargetReadPreferenceSpectraS3(request *models.PutAzureTargetReadPreferenceSpectraS3Request) (*models.PutAzureTargetReadPreferenceSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutAzureTargetReadPreferenceSpectraS3Response(response)
}
func (client *Client) RegisterAzureTargetSpectraS3(request *models.RegisterAzureTargetSpectraS3Request) (*models.RegisterAzureTargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewRegisterAzureTargetSpectraS3Response(response)
}
func (client *Client) PutDs3TargetReadPreferenceSpectraS3(request *models.PutDs3TargetReadPreferenceSpectraS3Request) (*models.PutDs3TargetReadPreferenceSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDs3TargetReadPreferenceSpectraS3Response(response)
}
func (client *Client) RegisterDs3TargetSpectraS3(request *models.RegisterDs3TargetSpectraS3Request) (*models.RegisterDs3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewRegisterDs3TargetSpectraS3Response(response)
}
func (client *Client) PutS3TargetBucketNameSpectraS3(request *models.PutS3TargetBucketNameSpectraS3Request) (*models.PutS3TargetBucketNameSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutS3TargetBucketNameSpectraS3Response(response)
}
func (client *Client) PutS3TargetReadPreferenceSpectraS3(request *models.PutS3TargetReadPreferenceSpectraS3Request) (*models.PutS3TargetReadPreferenceSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutS3TargetReadPreferenceSpectraS3Response(response)
}
func (client *Client) RegisterS3TargetSpectraS3(request *models.RegisterS3TargetSpectraS3Request) (*models.RegisterS3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewRegisterS3TargetSpectraS3Response(response)
}
func (client *Client) DelegateCreateUserSpectraS3(request *models.DelegateCreateUserSpectraS3Request) (*models.DelegateCreateUserSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.sendNetwork), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDelegateCreateUserSpectraS3Response(response)
}

