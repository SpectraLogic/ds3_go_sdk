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

func (client *Client) AbortMultiPartUpload(request *models.AbortMultiPartUploadRequest) (*models.AbortMultiPartUploadResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewAbortMultiPartUploadResponse(response)
}
func (client *Client) DeleteBucket(request *models.DeleteBucketRequest) (*models.DeleteBucketResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteBucketResponse(response)
}
func (client *Client) DeleteObject(request *models.DeleteObjectRequest) (*models.DeleteObjectResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteObjectResponse(response)
}
func (client *Client) DeleteBucketAclSpectraS3(request *models.DeleteBucketAclSpectraS3Request) (*models.DeleteBucketAclSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteBucketAclSpectraS3Response(response)
}
func (client *Client) DeleteDataPolicyAclSpectraS3(request *models.DeleteDataPolicyAclSpectraS3Request) (*models.DeleteDataPolicyAclSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteDataPolicyAclSpectraS3Response(response)
}
func (client *Client) DeleteBucketSpectraS3(request *models.DeleteBucketSpectraS3Request) (*models.DeleteBucketSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteBucketSpectraS3Response(response)
}
func (client *Client) DeleteAzureDataReplicationRuleSpectraS3(request *models.DeleteAzureDataReplicationRuleSpectraS3Request) (*models.DeleteAzureDataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteAzureDataReplicationRuleSpectraS3Response(response)
}
func (client *Client) DeleteDataPersistenceRuleSpectraS3(request *models.DeleteDataPersistenceRuleSpectraS3Request) (*models.DeleteDataPersistenceRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteDataPersistenceRuleSpectraS3Response(response)
}
func (client *Client) DeleteDataPolicySpectraS3(request *models.DeleteDataPolicySpectraS3Request) (*models.DeleteDataPolicySpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteDataPolicySpectraS3Response(response)
}
func (client *Client) DeleteDs3DataReplicationRuleSpectraS3(request *models.DeleteDs3DataReplicationRuleSpectraS3Request) (*models.DeleteDs3DataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteDs3DataReplicationRuleSpectraS3Response(response)
}
func (client *Client) DeleteS3DataReplicationRuleSpectraS3(request *models.DeleteS3DataReplicationRuleSpectraS3Request) (*models.DeleteS3DataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteS3DataReplicationRuleSpectraS3Response(response)
}
func (client *Client) ClearSuspectBlobAzureTargetsSpectraS3(request *models.ClearSuspectBlobAzureTargetsSpectraS3Request) (*models.ClearSuspectBlobAzureTargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewClearSuspectBlobAzureTargetsSpectraS3Response(response)
}
func (client *Client) ClearSuspectBlobDs3TargetsSpectraS3(request *models.ClearSuspectBlobDs3TargetsSpectraS3Request) (*models.ClearSuspectBlobDs3TargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewClearSuspectBlobDs3TargetsSpectraS3Response(response)
}
func (client *Client) ClearSuspectBlobPoolsSpectraS3(request *models.ClearSuspectBlobPoolsSpectraS3Request) (*models.ClearSuspectBlobPoolsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewClearSuspectBlobPoolsSpectraS3Response(response)
}
func (client *Client) ClearSuspectBlobS3TargetsSpectraS3(request *models.ClearSuspectBlobS3TargetsSpectraS3Request) (*models.ClearSuspectBlobS3TargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewClearSuspectBlobS3TargetsSpectraS3Response(response)
}
func (client *Client) ClearSuspectBlobTapesSpectraS3(request *models.ClearSuspectBlobTapesSpectraS3Request) (*models.ClearSuspectBlobTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewClearSuspectBlobTapesSpectraS3Response(response)
}
func (client *Client) DeleteGroupMemberSpectraS3(request *models.DeleteGroupMemberSpectraS3Request) (*models.DeleteGroupMemberSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteGroupMemberSpectraS3Response(response)
}
func (client *Client) DeleteGroupSpectraS3(request *models.DeleteGroupSpectraS3Request) (*models.DeleteGroupSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteGroupSpectraS3Response(response)
}
func (client *Client) CancelActiveJobSpectraS3(request *models.CancelActiveJobSpectraS3Request) (*models.CancelActiveJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelActiveJobSpectraS3Response(response)
}
func (client *Client) CancelAllActiveJobsSpectraS3(request *models.CancelAllActiveJobsSpectraS3Request) (*models.CancelAllActiveJobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelAllActiveJobsSpectraS3Response(response)
}
func (client *Client) CancelAllJobsSpectraS3(request *models.CancelAllJobsSpectraS3Request) (*models.CancelAllJobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelAllJobsSpectraS3Response(response)
}
func (client *Client) CancelJobSpectraS3(request *models.CancelJobSpectraS3Request) (*models.CancelJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelJobSpectraS3Response(response)
}
func (client *Client) ClearAllCanceledJobsSpectraS3(request *models.ClearAllCanceledJobsSpectraS3Request) (*models.ClearAllCanceledJobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewClearAllCanceledJobsSpectraS3Response(response)
}
func (client *Client) ClearAllCompletedJobsSpectraS3(request *models.ClearAllCompletedJobsSpectraS3Request) (*models.ClearAllCompletedJobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewClearAllCompletedJobsSpectraS3Response(response)
}
func (client *Client) TruncateActiveJobSpectraS3(request *models.TruncateActiveJobSpectraS3Request) (*models.TruncateActiveJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewTruncateActiveJobSpectraS3Response(response)
}
func (client *Client) TruncateAllActiveJobsSpectraS3(request *models.TruncateAllActiveJobsSpectraS3Request) (*models.TruncateAllActiveJobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewTruncateAllActiveJobsSpectraS3Response(response)
}
func (client *Client) TruncateAllJobsSpectraS3(request *models.TruncateAllJobsSpectraS3Request) (*models.TruncateAllJobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewTruncateAllJobsSpectraS3Response(response)
}
func (client *Client) TruncateJobSpectraS3(request *models.TruncateJobSpectraS3Request) (*models.TruncateJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewTruncateJobSpectraS3Response(response)
}
func (client *Client) DeleteAzureTargetFailureNotificationRegistrationSpectraS3(request *models.DeleteAzureTargetFailureNotificationRegistrationSpectraS3Request) (*models.DeleteAzureTargetFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteAzureTargetFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteDs3TargetFailureNotificationRegistrationSpectraS3(request *models.DeleteDs3TargetFailureNotificationRegistrationSpectraS3Request) (*models.DeleteDs3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteDs3TargetFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteJobCompletedNotificationRegistrationSpectraS3(request *models.DeleteJobCompletedNotificationRegistrationSpectraS3Request) (*models.DeleteJobCompletedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteJobCompletedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteJobCreatedNotificationRegistrationSpectraS3(request *models.DeleteJobCreatedNotificationRegistrationSpectraS3Request) (*models.DeleteJobCreatedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteJobCreatedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteJobCreationFailedNotificationRegistrationSpectraS3(request *models.DeleteJobCreationFailedNotificationRegistrationSpectraS3Request) (*models.DeleteJobCreationFailedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteJobCreationFailedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteObjectCachedNotificationRegistrationSpectraS3(request *models.DeleteObjectCachedNotificationRegistrationSpectraS3Request) (*models.DeleteObjectCachedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteObjectCachedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteObjectLostNotificationRegistrationSpectraS3(request *models.DeleteObjectLostNotificationRegistrationSpectraS3Request) (*models.DeleteObjectLostNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteObjectLostNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteObjectPersistedNotificationRegistrationSpectraS3(request *models.DeleteObjectPersistedNotificationRegistrationSpectraS3Request) (*models.DeleteObjectPersistedNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteObjectPersistedNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeletePoolFailureNotificationRegistrationSpectraS3(request *models.DeletePoolFailureNotificationRegistrationSpectraS3Request) (*models.DeletePoolFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeletePoolFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteS3TargetFailureNotificationRegistrationSpectraS3(request *models.DeleteS3TargetFailureNotificationRegistrationSpectraS3Request) (*models.DeleteS3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteS3TargetFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteStorageDomainFailureNotificationRegistrationSpectraS3(request *models.DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request) (*models.DeleteStorageDomainFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteStorageDomainFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteSystemFailureNotificationRegistrationSpectraS3(request *models.DeleteSystemFailureNotificationRegistrationSpectraS3Request) (*models.DeleteSystemFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteSystemFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteTapeFailureNotificationRegistrationSpectraS3(request *models.DeleteTapeFailureNotificationRegistrationSpectraS3Request) (*models.DeleteTapeFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteTapeFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteTapePartitionFailureNotificationRegistrationSpectraS3(request *models.DeleteTapePartitionFailureNotificationRegistrationSpectraS3Request) (*models.DeleteTapePartitionFailureNotificationRegistrationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteTapePartitionFailureNotificationRegistrationSpectraS3Response(response)
}
func (client *Client) DeleteFolderRecursivelySpectraS3(request *models.DeleteFolderRecursivelySpectraS3Request) (*models.DeleteFolderRecursivelySpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteFolderRecursivelySpectraS3Response(response)
}
func (client *Client) DeletePermanentlyLostPoolSpectraS3(request *models.DeletePermanentlyLostPoolSpectraS3Request) (*models.DeletePermanentlyLostPoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeletePermanentlyLostPoolSpectraS3Response(response)
}
func (client *Client) DeletePoolFailureSpectraS3(request *models.DeletePoolFailureSpectraS3Request) (*models.DeletePoolFailureSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeletePoolFailureSpectraS3Response(response)
}
func (client *Client) DeletePoolPartitionSpectraS3(request *models.DeletePoolPartitionSpectraS3Request) (*models.DeletePoolPartitionSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeletePoolPartitionSpectraS3Response(response)
}
func (client *Client) DeleteStorageDomainFailureSpectraS3(request *models.DeleteStorageDomainFailureSpectraS3Request) (*models.DeleteStorageDomainFailureSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteStorageDomainFailureSpectraS3Response(response)
}
func (client *Client) DeleteStorageDomainMemberSpectraS3(request *models.DeleteStorageDomainMemberSpectraS3Request) (*models.DeleteStorageDomainMemberSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteStorageDomainMemberSpectraS3Response(response)
}
func (client *Client) DeleteStorageDomainSpectraS3(request *models.DeleteStorageDomainSpectraS3Request) (*models.DeleteStorageDomainSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteStorageDomainSpectraS3Response(response)
}
func (client *Client) DeletePermanentlyLostTapeSpectraS3(request *models.DeletePermanentlyLostTapeSpectraS3Request) (*models.DeletePermanentlyLostTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeletePermanentlyLostTapeSpectraS3Response(response)
}
func (client *Client) DeleteTapeDensityDirectiveSpectraS3(request *models.DeleteTapeDensityDirectiveSpectraS3Request) (*models.DeleteTapeDensityDirectiveSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteTapeDensityDirectiveSpectraS3Response(response)
}
func (client *Client) DeleteTapeDriveSpectraS3(request *models.DeleteTapeDriveSpectraS3Request) (*models.DeleteTapeDriveSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteTapeDriveSpectraS3Response(response)
}
func (client *Client) DeleteTapeFailureSpectraS3(request *models.DeleteTapeFailureSpectraS3Request) (*models.DeleteTapeFailureSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteTapeFailureSpectraS3Response(response)
}
func (client *Client) DeleteTapePartitionFailureSpectraS3(request *models.DeleteTapePartitionFailureSpectraS3Request) (*models.DeleteTapePartitionFailureSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteTapePartitionFailureSpectraS3Response(response)
}
func (client *Client) DeleteTapePartitionSpectraS3(request *models.DeleteTapePartitionSpectraS3Request) (*models.DeleteTapePartitionSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteTapePartitionSpectraS3Response(response)
}
func (client *Client) DeleteAzureTargetBucketNameSpectraS3(request *models.DeleteAzureTargetBucketNameSpectraS3Request) (*models.DeleteAzureTargetBucketNameSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteAzureTargetBucketNameSpectraS3Response(response)
}
func (client *Client) DeleteAzureTargetFailureSpectraS3(request *models.DeleteAzureTargetFailureSpectraS3Request) (*models.DeleteAzureTargetFailureSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteAzureTargetFailureSpectraS3Response(response)
}
func (client *Client) DeleteAzureTargetReadPreferenceSpectraS3(request *models.DeleteAzureTargetReadPreferenceSpectraS3Request) (*models.DeleteAzureTargetReadPreferenceSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteAzureTargetReadPreferenceSpectraS3Response(response)
}
func (client *Client) DeleteAzureTargetSpectraS3(request *models.DeleteAzureTargetSpectraS3Request) (*models.DeleteAzureTargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteAzureTargetSpectraS3Response(response)
}
func (client *Client) DeleteDs3TargetFailureSpectraS3(request *models.DeleteDs3TargetFailureSpectraS3Request) (*models.DeleteDs3TargetFailureSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteDs3TargetFailureSpectraS3Response(response)
}
func (client *Client) DeleteDs3TargetReadPreferenceSpectraS3(request *models.DeleteDs3TargetReadPreferenceSpectraS3Request) (*models.DeleteDs3TargetReadPreferenceSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteDs3TargetReadPreferenceSpectraS3Response(response)
}
func (client *Client) DeleteDs3TargetSpectraS3(request *models.DeleteDs3TargetSpectraS3Request) (*models.DeleteDs3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteDs3TargetSpectraS3Response(response)
}
func (client *Client) DeleteS3TargetBucketNameSpectraS3(request *models.DeleteS3TargetBucketNameSpectraS3Request) (*models.DeleteS3TargetBucketNameSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteS3TargetBucketNameSpectraS3Response(response)
}
func (client *Client) DeleteS3TargetFailureSpectraS3(request *models.DeleteS3TargetFailureSpectraS3Request) (*models.DeleteS3TargetFailureSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteS3TargetFailureSpectraS3Response(response)
}
func (client *Client) DeleteS3TargetReadPreferenceSpectraS3(request *models.DeleteS3TargetReadPreferenceSpectraS3Request) (*models.DeleteS3TargetReadPreferenceSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteS3TargetReadPreferenceSpectraS3Response(response)
}
func (client *Client) DeleteS3TargetSpectraS3(request *models.DeleteS3TargetSpectraS3Request) (*models.DeleteS3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteS3TargetSpectraS3Response(response)
}
func (client *Client) DelegateDeleteUserSpectraS3(request *models.DelegateDeleteUserSpectraS3Request) (*models.DelegateDeleteUserSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDelegateDeleteUserSpectraS3Response(response)
}

