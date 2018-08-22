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
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/networking"
)

func (client *Client) CompleteBlob(request *models.CompleteBlobRequest) (*models.CompleteBlobResponse, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/" + request.BucketName + "/" + request.ObjectName).
        WithQueryParam("blob", request.Blob).
        WithQueryParam("job", request.Job).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCompleteBlobResponse(response)
}

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

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

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
        WithReadCloser(buildDeleteObjectsPayload(request.ObjectNames)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeleteObjectsResponse(response)
}

func (client *Client) InitiateMultiPartUpload(request *models.InitiateMultiPartUploadRequest) (*models.InitiateMultiPartUploadResponse, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/" + request.BucketName + "/" + request.ObjectName).
        WithQueryParam("uploads", "").
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewInitiateMultiPartUploadResponse(response)
}

func (client *Client) PutBucketAclForGroupSpectraS3(request *models.PutBucketAclForGroupSpectraS3Request) (*models.PutBucketAclForGroupSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/bucket_acl").
        WithQueryParam("bucket_id", request.BucketId).
        WithQueryParam("group_id", request.GroupId).
        WithQueryParam("permission", request.Permission.String()).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutBucketAclForGroupSpectraS3Response(response)
}

func (client *Client) PutBucketAclForUserSpectraS3(request *models.PutBucketAclForUserSpectraS3Request) (*models.PutBucketAclForUserSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/bucket_acl").
        WithQueryParam("bucket_id", request.BucketId).
        WithQueryParam("permission", request.Permission.String()).
        WithQueryParam("user_id", request.UserId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutBucketAclForUserSpectraS3Response(response)
}

func (client *Client) PutDataPolicyAclForGroupSpectraS3(request *models.PutDataPolicyAclForGroupSpectraS3Request) (*models.PutDataPolicyAclForGroupSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/data_policy_acl").
        WithQueryParam("data_policy_id", request.DataPolicyId).
        WithQueryParam("group_id", request.GroupId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDataPolicyAclForGroupSpectraS3Response(response)
}

func (client *Client) PutDataPolicyAclForUserSpectraS3(request *models.PutDataPolicyAclForUserSpectraS3Request) (*models.PutDataPolicyAclForUserSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/data_policy_acl").
        WithQueryParam("data_policy_id", request.DataPolicyId).
        WithQueryParam("user_id", request.UserId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDataPolicyAclForUserSpectraS3Response(response)
}

func (client *Client) PutGlobalBucketAclForGroupSpectraS3(request *models.PutGlobalBucketAclForGroupSpectraS3Request) (*models.PutGlobalBucketAclForGroupSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/bucket_acl").
        WithQueryParam("group_id", request.GroupId).
        WithQueryParam("permission", request.Permission.String()).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGlobalBucketAclForGroupSpectraS3Response(response)
}

func (client *Client) PutGlobalBucketAclForUserSpectraS3(request *models.PutGlobalBucketAclForUserSpectraS3Request) (*models.PutGlobalBucketAclForUserSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/bucket_acl").
        WithQueryParam("permission", request.Permission.String()).
        WithQueryParam("user_id", request.UserId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGlobalBucketAclForUserSpectraS3Response(response)
}

func (client *Client) PutGlobalDataPolicyAclForGroupSpectraS3(request *models.PutGlobalDataPolicyAclForGroupSpectraS3Request) (*models.PutGlobalDataPolicyAclForGroupSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/data_policy_acl").
        WithQueryParam("group_id", request.GroupId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGlobalDataPolicyAclForGroupSpectraS3Response(response)
}

func (client *Client) PutGlobalDataPolicyAclForUserSpectraS3(request *models.PutGlobalDataPolicyAclForUserSpectraS3Request) (*models.PutGlobalDataPolicyAclForUserSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/data_policy_acl").
        WithQueryParam("user_id", request.UserId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGlobalDataPolicyAclForUserSpectraS3Response(response)
}

func (client *Client) PutBucketSpectraS3(request *models.PutBucketSpectraS3Request) (*models.PutBucketSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/bucket").
        WithQueryParam("name", request.Name).
        WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
        WithOptionalQueryParam("id", request.Id).
        WithOptionalQueryParam("user_id", request.UserId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
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

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutAzureDataReplicationRuleSpectraS3Response(response)
}

func (client *Client) PutDataPersistenceRuleSpectraS3(request *models.PutDataPersistenceRuleSpectraS3Request) (*models.PutDataPersistenceRuleSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/data_persistence_rule").
        WithQueryParam("data_policy_id", request.DataPolicyId).
        WithQueryParam("isolation_level", request.IsolationLevel.String()).
        WithQueryParam("storage_domain_id", request.StorageDomainId).
        WithQueryParam("type", request.DataPersistenceRuleType.String()).
        WithOptionalQueryParam("minimum_days_to_retain", networking.IntPtrToStrPtr(request.MinimumDaysToRetain)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDataPersistenceRuleSpectraS3Response(response)
}

func (client *Client) PutDataPolicySpectraS3(request *models.PutDataPolicySpectraS3Request) (*models.PutDataPolicySpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/data_policy").
        WithQueryParam("name", request.Name).
        WithOptionalQueryParam("always_force_put_job_creation", networking.BoolPtrToStrPtr(request.AlwaysForcePutJobCreation)).
        WithOptionalQueryParam("always_minimize_spanning_across_media", networking.BoolPtrToStrPtr(request.AlwaysMinimizeSpanningAcrossMedia)).
        WithOptionalQueryParam("blobbing_enabled", networking.BoolPtrToStrPtr(request.BlobbingEnabled)).
        WithOptionalQueryParam("checksum_type", networking.InterfaceToStrPtr(request.ChecksumType)).
        WithOptionalQueryParam("default_blob_size", networking.Int64PtrToStrPtr(request.DefaultBlobSize)).
        WithOptionalQueryParam("default_get_job_priority", networking.InterfaceToStrPtr(request.DefaultGetJobPriority)).
        WithOptionalQueryParam("default_put_job_priority", networking.InterfaceToStrPtr(request.DefaultPutJobPriority)).
        WithOptionalQueryParam("default_verify_after_write", networking.BoolPtrToStrPtr(request.DefaultVerifyAfterWrite)).
        WithOptionalQueryParam("default_verify_job_priority", networking.InterfaceToStrPtr(request.DefaultVerifyJobPriority)).
        WithOptionalQueryParam("end_to_end_crc_required", networking.BoolPtrToStrPtr(request.EndToEndCrcRequired)).
        WithOptionalQueryParam("max_versions_to_keep", networking.IntPtrToStrPtr(request.MaxVersionsToKeep)).
        WithOptionalQueryParam("rebuild_priority", networking.InterfaceToStrPtr(request.RebuildPriority)).
        WithOptionalQueryParam("versioning", networking.InterfaceToStrPtr(request.Versioning)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDataPolicySpectraS3Response(response)
}

func (client *Client) PutDs3DataReplicationRuleSpectraS3(request *models.PutDs3DataReplicationRuleSpectraS3Request) (*models.PutDs3DataReplicationRuleSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/ds3_data_replication_rule").
        WithQueryParam("data_policy_id", request.DataPolicyId).
        WithQueryParam("target_id", request.TargetId).
        WithQueryParam("type", request.DataReplicationRuleType.String()).
        WithOptionalQueryParam("replicate_deletes", networking.BoolPtrToStrPtr(request.ReplicateDeletes)).
        WithOptionalQueryParam("target_data_policy", request.TargetDataPolicy).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDs3DataReplicationRuleSpectraS3Response(response)
}

func (client *Client) PutS3DataReplicationRuleSpectraS3(request *models.PutS3DataReplicationRuleSpectraS3Request) (*models.PutS3DataReplicationRuleSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/s3_data_replication_rule").
        WithQueryParam("data_policy_id", request.DataPolicyId).
        WithQueryParam("target_id", request.TargetId).
        WithQueryParam("type", request.DataReplicationRuleType.String()).
        WithOptionalQueryParam("initial_data_placement", networking.InterfaceToStrPtr(request.InitialDataPlacement)).
        WithOptionalQueryParam("max_blob_part_size_in_bytes", networking.Int64PtrToStrPtr(request.MaxBlobPartSizeInBytes)).
        WithOptionalQueryParam("replicate_deletes", networking.BoolPtrToStrPtr(request.ReplicateDeletes)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutS3DataReplicationRuleSpectraS3Response(response)
}

func (client *Client) PutGroupGroupMemberSpectraS3(request *models.PutGroupGroupMemberSpectraS3Request) (*models.PutGroupGroupMemberSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/group_member").
        WithQueryParam("group_id", request.GroupId).
        WithQueryParam("member_group_id", request.MemberGroupId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGroupGroupMemberSpectraS3Response(response)
}

func (client *Client) PutGroupSpectraS3(request *models.PutGroupSpectraS3Request) (*models.PutGroupSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/group").
        WithQueryParam("name", request.Name).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutGroupSpectraS3Response(response)
}

func (client *Client) PutUserGroupMemberSpectraS3(request *models.PutUserGroupMemberSpectraS3Request) (*models.PutUserGroupMemberSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/group_member").
        WithQueryParam("group_id", request.GroupId).
        WithQueryParam("member_user_id", request.MemberUserId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutUserGroupMemberSpectraS3Response(response)
}

func (client *Client) PutAzureTargetFailureNotificationRegistrationSpectraS3(request *models.PutAzureTargetFailureNotificationRegistrationSpectraS3Request) (*models.PutAzureTargetFailureNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/azure_target_failure_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutAzureTargetFailureNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutDs3TargetFailureNotificationRegistrationSpectraS3(request *models.PutDs3TargetFailureNotificationRegistrationSpectraS3Request) (*models.PutDs3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/ds3_target_failure_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDs3TargetFailureNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutJobCompletedNotificationRegistrationSpectraS3(request *models.PutJobCompletedNotificationRegistrationSpectraS3Request) (*models.PutJobCompletedNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/job_completed_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("job_id", request.JobId).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutJobCompletedNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutJobCreatedNotificationRegistrationSpectraS3(request *models.PutJobCreatedNotificationRegistrationSpectraS3Request) (*models.PutJobCreatedNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/job_created_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutJobCreatedNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutJobCreationFailedNotificationRegistrationSpectraS3(request *models.PutJobCreationFailedNotificationRegistrationSpectraS3Request) (*models.PutJobCreationFailedNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/job_creation_failed_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutJobCreationFailedNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutObjectCachedNotificationRegistrationSpectraS3(request *models.PutObjectCachedNotificationRegistrationSpectraS3Request) (*models.PutObjectCachedNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/object_cached_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("job_id", request.JobId).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutObjectCachedNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutObjectLostNotificationRegistrationSpectraS3(request *models.PutObjectLostNotificationRegistrationSpectraS3Request) (*models.PutObjectLostNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/object_lost_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutObjectLostNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutObjectPersistedNotificationRegistrationSpectraS3(request *models.PutObjectPersistedNotificationRegistrationSpectraS3Request) (*models.PutObjectPersistedNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/object_persisted_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("job_id", request.JobId).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutObjectPersistedNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutPoolFailureNotificationRegistrationSpectraS3(request *models.PutPoolFailureNotificationRegistrationSpectraS3Request) (*models.PutPoolFailureNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/pool_failure_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutPoolFailureNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutS3TargetFailureNotificationRegistrationSpectraS3(request *models.PutS3TargetFailureNotificationRegistrationSpectraS3Request) (*models.PutS3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/s3_target_failure_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutS3TargetFailureNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutStorageDomainFailureNotificationRegistrationSpectraS3(request *models.PutStorageDomainFailureNotificationRegistrationSpectraS3Request) (*models.PutStorageDomainFailureNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/storage_domain_failure_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutStorageDomainFailureNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutSystemFailureNotificationRegistrationSpectraS3(request *models.PutSystemFailureNotificationRegistrationSpectraS3Request) (*models.PutSystemFailureNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/system_failure_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutSystemFailureNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutTapeFailureNotificationRegistrationSpectraS3(request *models.PutTapeFailureNotificationRegistrationSpectraS3Request) (*models.PutTapeFailureNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/tape_failure_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutTapeFailureNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutTapePartitionFailureNotificationRegistrationSpectraS3(request *models.PutTapePartitionFailureNotificationRegistrationSpectraS3Request) (*models.PutTapePartitionFailureNotificationRegistrationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/tape_partition_failure_notification_registration").
        WithQueryParam("notification_end_point", request.NotificationEndPoint).
        WithOptionalQueryParam("format", networking.InterfaceToStrPtr(request.Format)).
        WithOptionalQueryParam("naming_convention", networking.InterfaceToStrPtr(request.NamingConvention)).
        WithOptionalQueryParam("notification_http_method", networking.InterfaceToStrPtr(request.NotificationHttpMethod)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutTapePartitionFailureNotificationRegistrationSpectraS3Response(response)
}

func (client *Client) PutPoolPartitionSpectraS3(request *models.PutPoolPartitionSpectraS3Request) (*models.PutPoolPartitionSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/pool_partition").
        WithQueryParam("name", request.Name).
        WithQueryParam("type", request.PoolType.String()).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutPoolPartitionSpectraS3Response(response)
}

func (client *Client) PutPoolStorageDomainMemberSpectraS3(request *models.PutPoolStorageDomainMemberSpectraS3Request) (*models.PutPoolStorageDomainMemberSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/storage_domain_member").
        WithQueryParam("pool_partition_id", request.PoolPartitionId).
        WithQueryParam("storage_domain_id", request.StorageDomainId).
        WithOptionalQueryParam("write_preference", networking.InterfaceToStrPtr(request.WritePreference)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutPoolStorageDomainMemberSpectraS3Response(response)
}

func (client *Client) PutStorageDomainSpectraS3(request *models.PutStorageDomainSpectraS3Request) (*models.PutStorageDomainSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/storage_domain").
        WithQueryParam("name", request.Name).
        WithOptionalQueryParam("auto_eject_media_full_threshold", networking.Int64PtrToStrPtr(request.AutoEjectMediaFullThreshold)).
        WithOptionalQueryParam("auto_eject_upon_cron", request.AutoEjectUponCron).
        WithOptionalQueryParam("auto_eject_upon_job_cancellation", networking.BoolPtrToStrPtr(request.AutoEjectUponJobCancellation)).
        WithOptionalQueryParam("auto_eject_upon_job_completion", networking.BoolPtrToStrPtr(request.AutoEjectUponJobCompletion)).
        WithOptionalQueryParam("auto_eject_upon_media_full", networking.BoolPtrToStrPtr(request.AutoEjectUponMediaFull)).
        WithOptionalQueryParam("ltfs_file_naming", networking.InterfaceToStrPtr(request.LtfsFileNaming)).
        WithOptionalQueryParam("max_tape_fragmentation_percent", networking.IntPtrToStrPtr(request.MaxTapeFragmentationPercent)).
        WithOptionalQueryParam("maximum_auto_verification_frequency_in_days", networking.IntPtrToStrPtr(request.MaximumAutoVerificationFrequencyInDays)).
        WithOptionalQueryParam("media_ejection_allowed", networking.BoolPtrToStrPtr(request.MediaEjectionAllowed)).
        WithOptionalQueryParam("secure_media_allocation", networking.BoolPtrToStrPtr(request.SecureMediaAllocation)).
        WithOptionalQueryParam("verify_prior_to_auto_eject", networking.InterfaceToStrPtr(request.VerifyPriorToAutoEject)).
        WithOptionalQueryParam("write_optimization", networking.InterfaceToStrPtr(request.WriteOptimization)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutStorageDomainSpectraS3Response(response)
}

func (client *Client) PutTapeStorageDomainMemberSpectraS3(request *models.PutTapeStorageDomainMemberSpectraS3Request) (*models.PutTapeStorageDomainMemberSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/storage_domain_member").
        WithQueryParam("storage_domain_id", request.StorageDomainId).
        WithQueryParam("tape_partition_id", request.TapePartitionId).
        WithQueryParam("tape_type", request.TapeType).
        WithOptionalQueryParam("auto_compaction_threshold", networking.IntPtrToStrPtr(request.AutoCompactionThreshold)).
        WithOptionalQueryParam("write_preference", networking.InterfaceToStrPtr(request.WritePreference)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutTapeStorageDomainMemberSpectraS3Response(response)
}

func (client *Client) PutTapeDensityDirectiveSpectraS3(request *models.PutTapeDensityDirectiveSpectraS3Request) (*models.PutTapeDensityDirectiveSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/tape_density_directive").
        WithQueryParam("density", request.Density.String()).
        WithQueryParam("partition_id", request.PartitionId).
        WithQueryParam("tape_type", request.TapeType).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutTapeDensityDirectiveSpectraS3Response(response)
}

func (client *Client) PutAzureTargetBucketNameSpectraS3(request *models.PutAzureTargetBucketNameSpectraS3Request) (*models.PutAzureTargetBucketNameSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/azure_target_bucket_name").
        WithQueryParam("bucket_id", request.BucketId).
        WithQueryParam("name", request.Name).
        WithQueryParam("target_id", request.TargetId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutAzureTargetBucketNameSpectraS3Response(response)
}

func (client *Client) PutAzureTargetReadPreferenceSpectraS3(request *models.PutAzureTargetReadPreferenceSpectraS3Request) (*models.PutAzureTargetReadPreferenceSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/azure_target_read_preference").
        WithQueryParam("bucket_id", request.BucketId).
        WithQueryParam("read_preference", request.ReadPreference.String()).
        WithQueryParam("target_id", request.TargetId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutAzureTargetReadPreferenceSpectraS3Response(response)
}

func (client *Client) RegisterAzureTargetSpectraS3(request *models.RegisterAzureTargetSpectraS3Request) (*models.RegisterAzureTargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/azure_target").
        WithQueryParam("account_key", request.AccountKey).
        WithQueryParam("account_name", request.AccountName).
        WithQueryParam("name", request.Name).
        WithOptionalQueryParam("auto_verify_frequency_in_days", networking.IntPtrToStrPtr(request.AutoVerifyFrequencyInDays)).
        WithOptionalQueryParam("cloud_bucket_prefix", request.CloudBucketPrefix).
        WithOptionalQueryParam("cloud_bucket_suffix", request.CloudBucketSuffix).
        WithOptionalQueryParam("default_read_preference", networking.InterfaceToStrPtr(request.DefaultReadPreference)).
        WithOptionalQueryParam("https", networking.BoolPtrToStrPtr(request.Https)).
        WithOptionalQueryParam("permit_going_out_of_sync", networking.BoolPtrToStrPtr(request.PermitGoingOutOfSync)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewRegisterAzureTargetSpectraS3Response(response)
}

func (client *Client) PutDs3TargetReadPreferenceSpectraS3(request *models.PutDs3TargetReadPreferenceSpectraS3Request) (*models.PutDs3TargetReadPreferenceSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/ds3_target_read_preference").
        WithQueryParam("bucket_id", request.BucketId).
        WithQueryParam("read_preference", request.ReadPreference.String()).
        WithQueryParam("target_id", request.TargetId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutDs3TargetReadPreferenceSpectraS3Response(response)
}

func (client *Client) RegisterDs3TargetSpectraS3(request *models.RegisterDs3TargetSpectraS3Request) (*models.RegisterDs3TargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/ds3_target").
        WithQueryParam("admin_auth_id", request.AdminAuthId).
        WithQueryParam("admin_secret_key", request.AdminSecretKey).
        WithQueryParam("data_path_end_point", request.DataPathEndPoint).
        WithQueryParam("name", request.Name).
        WithOptionalQueryParam("access_control_replication", networking.InterfaceToStrPtr(request.AccessControlReplication)).
        WithOptionalQueryParam("data_path_https", networking.BoolPtrToStrPtr(request.DataPathHttps)).
        WithOptionalQueryParam("data_path_port", networking.IntPtrToStrPtr(request.DataPathPort)).
        WithOptionalQueryParam("data_path_proxy", request.DataPathProxy).
        WithOptionalQueryParam("data_path_verify_certificate", networking.BoolPtrToStrPtr(request.DataPathVerifyCertificate)).
        WithOptionalQueryParam("default_read_preference", networking.InterfaceToStrPtr(request.DefaultReadPreference)).
        WithOptionalQueryParam("permit_going_out_of_sync", networking.BoolPtrToStrPtr(request.PermitGoingOutOfSync)).
        WithOptionalQueryParam("replicated_user_default_data_policy", request.ReplicatedUserDefaultDataPolicy).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewRegisterDs3TargetSpectraS3Response(response)
}

func (client *Client) PutS3TargetBucketNameSpectraS3(request *models.PutS3TargetBucketNameSpectraS3Request) (*models.PutS3TargetBucketNameSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/s3_target_bucket_name").
        WithQueryParam("bucket_id", request.BucketId).
        WithQueryParam("name", request.Name).
        WithQueryParam("target_id", request.TargetId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutS3TargetBucketNameSpectraS3Response(response)
}

func (client *Client) PutS3TargetReadPreferenceSpectraS3(request *models.PutS3TargetReadPreferenceSpectraS3Request) (*models.PutS3TargetReadPreferenceSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/s3_target_read_preference").
        WithQueryParam("bucket_id", request.BucketId).
        WithQueryParam("read_preference", request.ReadPreference.String()).
        WithQueryParam("target_id", request.TargetId).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutS3TargetReadPreferenceSpectraS3Response(response)
}

func (client *Client) RegisterS3TargetSpectraS3(request *models.RegisterS3TargetSpectraS3Request) (*models.RegisterS3TargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/s3_target").
        WithQueryParam("access_key", request.AccessKey).
        WithQueryParam("name", request.Name).
        WithQueryParam("secret_key", request.SecretKey).
        WithOptionalQueryParam("auto_verify_frequency_in_days", networking.IntPtrToStrPtr(request.AutoVerifyFrequencyInDays)).
        WithOptionalQueryParam("cloud_bucket_prefix", request.CloudBucketPrefix).
        WithOptionalQueryParam("cloud_bucket_suffix", request.CloudBucketSuffix).
        WithOptionalQueryParam("data_path_end_point", request.DataPathEndPoint).
        WithOptionalQueryParam("default_read_preference", networking.InterfaceToStrPtr(request.DefaultReadPreference)).
        WithOptionalQueryParam("https", networking.BoolPtrToStrPtr(request.Https)).
        WithOptionalQueryParam("naming_mode", networking.InterfaceToStrPtr(request.NamingMode)).
        WithOptionalQueryParam("offline_data_staging_window_in_tb", networking.IntPtrToStrPtr(request.OfflineDataStagingWindowInTb)).
        WithOptionalQueryParam("permit_going_out_of_sync", networking.BoolPtrToStrPtr(request.PermitGoingOutOfSync)).
        WithOptionalQueryParam("proxy_domain", request.ProxyDomain).
        WithOptionalQueryParam("proxy_host", request.ProxyHost).
        WithOptionalQueryParam("proxy_password", request.ProxyPassword).
        WithOptionalQueryParam("proxy_port", networking.IntPtrToStrPtr(request.ProxyPort)).
        WithOptionalQueryParam("proxy_username", request.ProxyUsername).
        WithOptionalQueryParam("region", networking.InterfaceToStrPtr(request.Region)).
        WithOptionalQueryParam("staged_data_expiration_in_days", networking.IntPtrToStrPtr(request.StagedDataExpirationInDays)).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewRegisterS3TargetSpectraS3Response(response)
}

func (client *Client) DelegateCreateUserSpectraS3(request *models.DelegateCreateUserSpectraS3Request) (*models.DelegateCreateUserSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_POST).
        WithPath("/_rest_/user").
        WithQueryParam("name", request.Name).
        WithOptionalQueryParam("default_data_policy_id", request.DefaultDataPolicyId).
        WithOptionalQueryParam("id", request.Id).
        WithOptionalQueryParam("max_buckets", networking.IntPtrToStrPtr(request.MaxBuckets)).
        WithOptionalQueryParam("secret_key", request.SecretKey).
        Build(client.connectionInfo)

    if err != nil {
        return nil, err
    }

    networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(httpRequest)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDelegateCreateUserSpectraS3Response(response)
}


