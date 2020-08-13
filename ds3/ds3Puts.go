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
    "strconv"
)

func (client *Client) PutBucket(request *models.PutBucketRequest) (*models.PutBucketResponse, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/" + request.BucketName).
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
    return models.NewPutBucketResponse(response)
}

func (client *Client) PutMultiPartUploadPart(request *models.PutMultiPartUploadPartRequest) (*models.PutMultiPartUploadPartResponse, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/" + request.BucketName + "/" + request.ObjectName).
        WithQueryParam("part_number", strconv.Itoa(request.PartNumber)).
        WithQueryParam("upload_id", request.UploadId).
        WithReader(request.Content).
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
    return models.NewPutMultiPartUploadPartResponse(response)
}

func (client *Client) PutObject(request *models.PutObjectRequest) (*models.PutObjectResponse, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/" + request.BucketName + "/" + request.ObjectName).
        WithOptionalQueryParam("job", request.Job).
        WithOptionalQueryParam("offset", networking.Int64PtrToStrPtr(request.Offset)).
        WithReader(request.Content).
        WithChecksum(request.Checksum).
        WithHeaders(request.Metadata).
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
    return models.NewPutObjectResponse(response)
}

func (client *Client) ModifyBucketSpectraS3(request *models.ModifyBucketSpectraS3Request) (*models.ModifyBucketSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/bucket/" + request.BucketName).
        WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
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
    return models.NewModifyBucketSpectraS3Response(response)
}

func (client *Client) ForceFullCacheReclaimSpectraS3(request *models.ForceFullCacheReclaimSpectraS3Request) (*models.ForceFullCacheReclaimSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/cache_filesystem").
        WithQueryParam("reclaim", "").
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
    return models.NewForceFullCacheReclaimSpectraS3Response(response)
}

func (client *Client) ModifyCacheFilesystemSpectraS3(request *models.ModifyCacheFilesystemSpectraS3Request) (*models.ModifyCacheFilesystemSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/cache_filesystem/" + request.CacheFilesystem).
        WithOptionalQueryParam("auto_reclaim_initiate_threshold", networking.Float64PtrToStrPtr(request.AutoReclaimInitiateThreshold)).
        WithOptionalQueryParam("auto_reclaim_terminate_threshold", networking.Float64PtrToStrPtr(request.AutoReclaimTerminateThreshold)).
        WithOptionalQueryParam("burst_threshold", networking.Float64PtrToStrPtr(request.BurstThreshold)).
        WithOptionalQueryParam("max_capacity_in_bytes", networking.Int64PtrToStrPtr(request.MaxCapacityInBytes)).
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
    return models.NewModifyCacheFilesystemSpectraS3Response(response)
}

func (client *Client) ModifyDataPathBackendSpectraS3(request *models.ModifyDataPathBackendSpectraS3Request) (*models.ModifyDataPathBackendSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/data_path_backend").
        WithOptionalQueryParam("activated", networking.BoolPtrToStrPtr(request.Activated)).
        WithOptionalQueryParam("allow_new_job_requests", networking.BoolPtrToStrPtr(request.AllowNewJobRequests)).
        WithOptionalQueryParam("auto_activate_timeout_in_mins", networking.IntPtrToStrPtr(request.AutoActivateTimeoutInMins)).
        WithOptionalQueryParam("auto_inspect", networking.InterfaceToStrPtr(request.AutoInspect)).
        WithOptionalQueryParam("cache_available_retry_after_in_seconds", networking.IntPtrToStrPtr(request.CacheAvailableRetryAfterInSeconds)).
        WithOptionalQueryParam("default_verify_data_after_import", networking.InterfaceToStrPtr(request.DefaultVerifyDataAfterImport)).
        WithOptionalQueryParam("default_verify_data_prior_to_import", networking.BoolPtrToStrPtr(request.DefaultVerifyDataPriorToImport)).
        WithOptionalQueryParam("iom_enabled", networking.BoolPtrToStrPtr(request.IomEnabled)).
        WithOptionalQueryParam("partially_verify_last_percent_of_tapes", networking.IntPtrToStrPtr(request.PartiallyVerifyLastPercentOfTapes)).
        WithOptionalQueryParam("unavailable_media_policy", networking.InterfaceToStrPtr(request.UnavailableMediaPolicy)).
        WithOptionalQueryParam("unavailable_pool_max_job_retry_in_mins", networking.IntPtrToStrPtr(request.UnavailablePoolMaxJobRetryInMins)).
        WithOptionalQueryParam("unavailable_tape_partition_max_job_retry_in_mins", networking.IntPtrToStrPtr(request.UnavailableTapePartitionMaxJobRetryInMins)).
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
    return models.NewModifyDataPathBackendSpectraS3Response(response)
}

func (client *Client) ModifyAzureDataReplicationRuleSpectraS3(request *models.ModifyAzureDataReplicationRuleSpectraS3Request) (*models.ModifyAzureDataReplicationRuleSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/azure_data_replication_rule/" + request.AzureDataReplicationRule).
        WithOptionalQueryParam("max_blob_part_size_in_bytes", networking.Int64PtrToStrPtr(request.MaxBlobPartSizeInBytes)).
        WithOptionalQueryParam("replicate_deletes", networking.BoolPtrToStrPtr(request.ReplicateDeletes)).
        WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataReplicationRuleType)).
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
    return models.NewModifyAzureDataReplicationRuleSpectraS3Response(response)
}

func (client *Client) ModifyDataPersistenceRuleSpectraS3(request *models.ModifyDataPersistenceRuleSpectraS3Request) (*models.ModifyDataPersistenceRuleSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/data_persistence_rule/" + request.DataPersistenceRuleId).
        WithOptionalQueryParam("isolation_level", networking.InterfaceToStrPtr(request.IsolationLevel)).
        WithOptionalQueryParam("minimum_days_to_retain", networking.IntPtrToStrPtr(request.MinimumDaysToRetain)).
        WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataPersistenceRuleType)).
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
    return models.NewModifyDataPersistenceRuleSpectraS3Response(response)
}

func (client *Client) ModifyDataPolicySpectraS3(request *models.ModifyDataPolicySpectraS3Request) (*models.ModifyDataPolicySpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/data_policy/" + request.DataPolicyId).
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
        WithOptionalQueryParam("name", request.Name).
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
    return models.NewModifyDataPolicySpectraS3Response(response)
}

func (client *Client) ModifyDs3DataReplicationRuleSpectraS3(request *models.ModifyDs3DataReplicationRuleSpectraS3Request) (*models.ModifyDs3DataReplicationRuleSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/ds3_data_replication_rule/" + request.Ds3DataReplicationRule).
        WithOptionalQueryParam("replicate_deletes", networking.BoolPtrToStrPtr(request.ReplicateDeletes)).
        WithOptionalQueryParam("target_data_policy", request.TargetDataPolicy).
        WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataReplicationRuleType)).
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
    return models.NewModifyDs3DataReplicationRuleSpectraS3Response(response)
}

func (client *Client) ModifyS3DataReplicationRuleSpectraS3(request *models.ModifyS3DataReplicationRuleSpectraS3Request) (*models.ModifyS3DataReplicationRuleSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/s3_data_replication_rule/" + request.S3DataReplicationRule).
        WithOptionalQueryParam("initial_data_placement", networking.InterfaceToStrPtr(request.InitialDataPlacement)).
        WithOptionalQueryParam("max_blob_part_size_in_bytes", networking.Int64PtrToStrPtr(request.MaxBlobPartSizeInBytes)).
        WithOptionalQueryParam("replicate_deletes", networking.BoolPtrToStrPtr(request.ReplicateDeletes)).
        WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataReplicationRuleType)).
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
    return models.NewModifyS3DataReplicationRuleSpectraS3Response(response)
}

func (client *Client) MarkSuspectBlobAzureTargetsAsDegradedSpectraS3(request *models.MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Request) (*models.MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/suspect_blob_azure_target").
        WithOptionalVoidQueryParam("force", request.Force).
        WithReadCloser(buildIdListPayload(request.Ids)).
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
    return models.NewMarkSuspectBlobAzureTargetsAsDegradedSpectraS3Response(response)
}

func (client *Client) MarkSuspectBlobDs3TargetsAsDegradedSpectraS3(request *models.MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request) (*models.MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/suspect_blob_ds3_target").
        WithOptionalVoidQueryParam("force", request.Force).
        WithReadCloser(buildIdListPayload(request.Ids)).
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
    return models.NewMarkSuspectBlobDs3TargetsAsDegradedSpectraS3Response(response)
}

func (client *Client) MarkSuspectBlobPoolsAsDegradedSpectraS3(request *models.MarkSuspectBlobPoolsAsDegradedSpectraS3Request) (*models.MarkSuspectBlobPoolsAsDegradedSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/suspect_blob_pool").
        WithOptionalVoidQueryParam("force", request.Force).
        WithReadCloser(buildIdListPayload(request.Ids)).
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
    return models.NewMarkSuspectBlobPoolsAsDegradedSpectraS3Response(response)
}

func (client *Client) MarkSuspectBlobS3TargetsAsDegradedSpectraS3(request *models.MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request) (*models.MarkSuspectBlobS3TargetsAsDegradedSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/suspect_blob_s3_target").
        WithOptionalVoidQueryParam("force", request.Force).
        WithReadCloser(buildIdListPayload(request.Ids)).
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
    return models.NewMarkSuspectBlobS3TargetsAsDegradedSpectraS3Response(response)
}

func (client *Client) MarkSuspectBlobTapesAsDegradedSpectraS3(request *models.MarkSuspectBlobTapesAsDegradedSpectraS3Request) (*models.MarkSuspectBlobTapesAsDegradedSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/suspect_blob_tape").
        WithOptionalVoidQueryParam("force", request.Force).
        WithReadCloser(buildIdListPayload(request.Ids)).
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
    return models.NewMarkSuspectBlobTapesAsDegradedSpectraS3Response(response)
}

func (client *Client) ModifyGroupSpectraS3(request *models.ModifyGroupSpectraS3Request) (*models.ModifyGroupSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/group/" + request.Group).
        WithOptionalQueryParam("name", request.Name).
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
    return models.NewModifyGroupSpectraS3Response(response)
}

func (client *Client) VerifyUserIsMemberOfGroupSpectraS3(request *models.VerifyUserIsMemberOfGroupSpectraS3Request) (*models.VerifyUserIsMemberOfGroupSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/group/" + request.Group).
        WithOptionalQueryParam("user_id", request.UserId).
        WithQueryParam("operation", "verify").
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
    return models.NewVerifyUserIsMemberOfGroupSpectraS3Response(response)
}

func (client *Client) AllocateJobChunkSpectraS3(request *models.AllocateJobChunkSpectraS3Request) (*models.AllocateJobChunkSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/job_chunk/" + request.JobChunkId).
        WithQueryParam("operation", "allocate").
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
    return models.NewAllocateJobChunkSpectraS3Response(response)
}

func (client *Client) CloseAggregatingJobSpectraS3(request *models.CloseAggregatingJobSpectraS3Request) (*models.CloseAggregatingJobSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/job/" + request.JobId).
        WithQueryParam("close_aggregating_job", "").
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
    return models.NewCloseAggregatingJobSpectraS3Response(response)
}

func (client *Client) GetBulkJobSpectraS3(request *models.GetBulkJobSpectraS3Request) (*models.GetBulkJobSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/bucket/" + request.BucketName).
        WithOptionalQueryParam("aggregating", networking.BoolPtrToStrPtr(request.Aggregating)).
        WithOptionalQueryParam("chunk_client_processing_order_guarantee", networking.InterfaceToStrPtr(request.ChunkClientProcessingOrderGuarantee)).
        WithOptionalQueryParam("implicit_job_id_resolution", networking.BoolPtrToStrPtr(request.ImplicitJobIdResolution)).
        WithOptionalQueryParam("name", request.Name).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithQueryParam("operation", "start_bulk_get").
        WithReadCloser(buildDs3GetObjectListStream(request.Objects)).
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
    return models.NewGetBulkJobSpectraS3Response(response)
}

func (client *Client) PutBulkJobSpectraS3(request *models.PutBulkJobSpectraS3Request) (*models.PutBulkJobSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/bucket/" + request.BucketName).
        WithOptionalQueryParam("aggregating", networking.BoolPtrToStrPtr(request.Aggregating)).
        WithOptionalVoidQueryParam("force", request.Force).
        WithOptionalVoidQueryParam("ignore_naming_conflicts", request.IgnoreNamingConflicts).
        WithOptionalQueryParam("implicit_job_id_resolution", networking.BoolPtrToStrPtr(request.ImplicitJobIdResolution)).
        WithOptionalQueryParam("max_upload_size", networking.Int64PtrToStrPtr(request.MaxUploadSize)).
        WithOptionalQueryParam("minimize_spanning_across_media", networking.BoolPtrToStrPtr(request.MinimizeSpanningAcrossMedia)).
        WithOptionalQueryParam("name", request.Name).
        WithOptionalVoidQueryParam("pre_allocate_job_space", request.PreAllocateJobSpace).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithOptionalQueryParam("verify_after_write", networking.BoolPtrToStrPtr(request.VerifyAfterWrite)).
        WithQueryParam("operation", "start_bulk_put").
        WithReadCloser(buildDs3PutObjectListStream(request.Objects)).
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
    return models.NewPutBulkJobSpectraS3Response(response)
}

func (client *Client) VerifyBulkJobSpectraS3(request *models.VerifyBulkJobSpectraS3Request) (*models.VerifyBulkJobSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/bucket/" + request.BucketName).
        WithOptionalQueryParam("aggregating", networking.BoolPtrToStrPtr(request.Aggregating)).
        WithOptionalQueryParam("name", request.Name).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithQueryParam("operation", "start_bulk_verify").
        WithReadCloser(buildDs3GetObjectListStream(request.Objects)).
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
    return models.NewVerifyBulkJobSpectraS3Response(response)
}

func (client *Client) ModifyActiveJobSpectraS3(request *models.ModifyActiveJobSpectraS3Request) (*models.ModifyActiveJobSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/active_job/" + request.ActiveJobId).
        WithOptionalQueryParam("created_at", request.CreatedAt).
        WithOptionalQueryParam("name", request.Name).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
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
    return models.NewModifyActiveJobSpectraS3Response(response)
}

func (client *Client) ModifyJobSpectraS3(request *models.ModifyJobSpectraS3Request) (*models.ModifyJobSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/job/" + request.JobId).
        WithOptionalQueryParam("created_at", request.CreatedAt).
        WithOptionalQueryParam("name", request.Name).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
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
    return models.NewModifyJobSpectraS3Response(response)
}

func (client *Client) ReplicatePutJobSpectraS3(request *models.ReplicatePutJobSpectraS3Request) (*models.ReplicatePutJobSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/bucket/" + request.BucketName).
        WithQueryParam("replicate", "").
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithQueryParam("operation", "start_bulk_put").
        WithReadCloser(buildStreamFromString(request.RequestPayload)).
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
    return models.NewReplicatePutJobSpectraS3Response(response)
}

func (client *Client) StageObjectsJobSpectraS3(request *models.StageObjectsJobSpectraS3Request) (*models.StageObjectsJobSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/bucket/" + request.BucketName).
        WithOptionalQueryParam("name", request.Name).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithQueryParam("operation", "start_bulk_stage").
        WithReadCloser(buildDs3GetObjectListStream(request.Objects)).
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
    return models.NewStageObjectsJobSpectraS3Response(response)
}

func (client *Client) VerifySafeToCreatePutJobSpectraS3(request *models.VerifySafeToCreatePutJobSpectraS3Request) (*models.VerifySafeToCreatePutJobSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/bucket/" + request.BucketName).
        WithQueryParam("operation", "verify_safe_to_start_bulk_put").
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
    return models.NewVerifySafeToCreatePutJobSpectraS3Response(response)
}

func (client *Client) ModifyNodeSpectraS3(request *models.ModifyNodeSpectraS3Request) (*models.ModifyNodeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/node/" + request.Node).
        WithOptionalQueryParam("dns_name", request.DnsName).
        WithOptionalQueryParam("name", request.Name).
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
    return models.NewModifyNodeSpectraS3Response(response)
}

func (client *Client) GetPhysicalPlacementForObjectsSpectraS3(request *models.GetPhysicalPlacementForObjectsSpectraS3Request) (*models.GetPhysicalPlacementForObjectsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/bucket/" + request.BucketName).
        WithOptionalQueryParam("storage_domain", request.StorageDomain).
        WithQueryParam("operation", "get_physical_placement").
        WithReadCloser(buildDs3GetObjectListStream(request.Objects)).
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
    return models.NewGetPhysicalPlacementForObjectsSpectraS3Response(response)
}

func (client *Client) GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3(request *models.GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) (*models.GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/bucket/" + request.BucketName).
        WithQueryParam("full_details", "").
        WithOptionalQueryParam("storage_domain", request.StorageDomain).
        WithQueryParam("operation", "get_physical_placement").
        WithReadCloser(buildDs3GetObjectListStream(request.Objects)).
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
    return models.NewGetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response(response)
}

func (client *Client) UndeleteObjectSpectraS3(request *models.UndeleteObjectSpectraS3Request) (*models.UndeleteObjectSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/object").
        WithQueryParam("bucket_id", request.BucketId).
        WithQueryParam("name", request.Name).
        WithOptionalQueryParam("version_id", request.VersionId).
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
    return models.NewUndeleteObjectSpectraS3Response(response)
}

func (client *Client) CancelImportOnAllPoolsSpectraS3(request *models.CancelImportOnAllPoolsSpectraS3Request) (*models.CancelImportOnAllPoolsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool").
        WithQueryParam("operation", "cancel_import").
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
    return models.NewCancelImportOnAllPoolsSpectraS3Response(response)
}

func (client *Client) CancelImportPoolSpectraS3(request *models.CancelImportPoolSpectraS3Request) (*models.CancelImportPoolSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool/" + request.Pool).
        WithQueryParam("operation", "cancel_import").
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
    return models.NewCancelImportPoolSpectraS3Response(response)
}

func (client *Client) CancelVerifyOnAllPoolsSpectraS3(request *models.CancelVerifyOnAllPoolsSpectraS3Request) (*models.CancelVerifyOnAllPoolsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool").
        WithQueryParam("operation", "cancel_verify").
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
    return models.NewCancelVerifyOnAllPoolsSpectraS3Response(response)
}

func (client *Client) CancelVerifyPoolSpectraS3(request *models.CancelVerifyPoolSpectraS3Request) (*models.CancelVerifyPoolSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool/" + request.Pool).
        WithQueryParam("operation", "cancel_verify").
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
    return models.NewCancelVerifyPoolSpectraS3Response(response)
}

func (client *Client) CompactAllPoolsSpectraS3(request *models.CompactAllPoolsSpectraS3Request) (*models.CompactAllPoolsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool").
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithQueryParam("operation", "compact").
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
    return models.NewCompactAllPoolsSpectraS3Response(response)
}

func (client *Client) CompactPoolSpectraS3(request *models.CompactPoolSpectraS3Request) (*models.CompactPoolSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool/" + request.Pool).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithQueryParam("operation", "compact").
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
    return models.NewCompactPoolSpectraS3Response(response)
}

func (client *Client) DeallocatePoolSpectraS3(request *models.DeallocatePoolSpectraS3Request) (*models.DeallocatePoolSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool/" + request.Pool).
        WithQueryParam("operation", "deallocate").
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
    return models.NewDeallocatePoolSpectraS3Response(response)
}

func (client *Client) ForcePoolEnvironmentRefreshSpectraS3(request *models.ForcePoolEnvironmentRefreshSpectraS3Request) (*models.ForcePoolEnvironmentRefreshSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool_environment").
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
    return models.NewForcePoolEnvironmentRefreshSpectraS3Response(response)
}

func (client *Client) FormatAllForeignPoolsSpectraS3(request *models.FormatAllForeignPoolsSpectraS3Request) (*models.FormatAllForeignPoolsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool").
        WithQueryParam("operation", "format").
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
    return models.NewFormatAllForeignPoolsSpectraS3Response(response)
}

func (client *Client) FormatForeignPoolSpectraS3(request *models.FormatForeignPoolSpectraS3Request) (*models.FormatForeignPoolSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool/" + request.Pool).
        WithQueryParam("operation", "format").
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
    return models.NewFormatForeignPoolSpectraS3Response(response)
}

func (client *Client) ImportAllPoolsSpectraS3(request *models.ImportAllPoolsSpectraS3Request) (*models.ImportAllPoolsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool").
        WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithOptionalQueryParam("storage_domain_id", request.StorageDomainId).
        WithOptionalQueryParam("user_id", request.UserId).
        WithOptionalQueryParam("verify_data_after_import", networking.InterfaceToStrPtr(request.VerifyDataAfterImport)).
        WithOptionalQueryParam("verify_data_prior_to_import", networking.BoolPtrToStrPtr(request.VerifyDataPriorToImport)).
        WithQueryParam("operation", "import").
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
    return models.NewImportAllPoolsSpectraS3Response(response)
}

func (client *Client) ImportPoolSpectraS3(request *models.ImportPoolSpectraS3Request) (*models.ImportPoolSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool/" + request.Pool).
        WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithOptionalQueryParam("storage_domain_id", request.StorageDomainId).
        WithOptionalQueryParam("user_id", request.UserId).
        WithOptionalQueryParam("verify_data_after_import", networking.InterfaceToStrPtr(request.VerifyDataAfterImport)).
        WithOptionalQueryParam("verify_data_prior_to_import", networking.BoolPtrToStrPtr(request.VerifyDataPriorToImport)).
        WithQueryParam("operation", "import").
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
    return models.NewImportPoolSpectraS3Response(response)
}

func (client *Client) ModifyAllPoolsSpectraS3(request *models.ModifyAllPoolsSpectraS3Request) (*models.ModifyAllPoolsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool").
        WithQueryParam("quiesced", request.Quiesced.String()).
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
    return models.NewModifyAllPoolsSpectraS3Response(response)
}

func (client *Client) ModifyPoolPartitionSpectraS3(request *models.ModifyPoolPartitionSpectraS3Request) (*models.ModifyPoolPartitionSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool_partition/" + request.PoolPartition).
        WithOptionalQueryParam("name", request.Name).
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
    return models.NewModifyPoolPartitionSpectraS3Response(response)
}

func (client *Client) ModifyPoolSpectraS3(request *models.ModifyPoolSpectraS3Request) (*models.ModifyPoolSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool/" + request.Pool).
        WithOptionalQueryParam("partition_id", request.PartitionId).
        WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
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
    return models.NewModifyPoolSpectraS3Response(response)
}

func (client *Client) VerifyAllPoolsSpectraS3(request *models.VerifyAllPoolsSpectraS3Request) (*models.VerifyAllPoolsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool").
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithQueryParam("operation", "verify").
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
    return models.NewVerifyAllPoolsSpectraS3Response(response)
}

func (client *Client) VerifyPoolSpectraS3(request *models.VerifyPoolSpectraS3Request) (*models.VerifyPoolSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/pool/" + request.Pool).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithQueryParam("operation", "verify").
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
    return models.NewVerifyPoolSpectraS3Response(response)
}

func (client *Client) ConvertStorageDomainToDs3TargetSpectraS3(request *models.ConvertStorageDomainToDs3TargetSpectraS3Request) (*models.ConvertStorageDomainToDs3TargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/storage_domain/" + request.StorageDomain).
        WithQueryParam("convert_to_ds3_target", request.ConvertToDs3Target).
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
    return models.NewConvertStorageDomainToDs3TargetSpectraS3Response(response)
}

func (client *Client) ModifyStorageDomainMemberSpectraS3(request *models.ModifyStorageDomainMemberSpectraS3Request) (*models.ModifyStorageDomainMemberSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/storage_domain_member/" + request.StorageDomainMember).
        WithOptionalQueryParam("auto_compaction_threshold", networking.IntPtrToStrPtr(request.AutoCompactionThreshold)).
        WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
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
    return models.NewModifyStorageDomainMemberSpectraS3Response(response)
}

func (client *Client) ModifyStorageDomainSpectraS3(request *models.ModifyStorageDomainSpectraS3Request) (*models.ModifyStorageDomainSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/storage_domain/" + request.StorageDomain).
        WithOptionalQueryParam("auto_eject_media_full_threshold", networking.Int64PtrToStrPtr(request.AutoEjectMediaFullThreshold)).
        WithOptionalQueryParam("auto_eject_upon_cron", request.AutoEjectUponCron).
        WithOptionalQueryParam("auto_eject_upon_job_cancellation", networking.BoolPtrToStrPtr(request.AutoEjectUponJobCancellation)).
        WithOptionalQueryParam("auto_eject_upon_job_completion", networking.BoolPtrToStrPtr(request.AutoEjectUponJobCompletion)).
        WithOptionalQueryParam("auto_eject_upon_media_full", networking.BoolPtrToStrPtr(request.AutoEjectUponMediaFull)).
        WithOptionalQueryParam("ltfs_file_naming", networking.InterfaceToStrPtr(request.LtfsFileNaming)).
        WithOptionalQueryParam("max_tape_fragmentation_percent", networking.IntPtrToStrPtr(request.MaxTapeFragmentationPercent)).
        WithOptionalQueryParam("maximum_auto_verification_frequency_in_days", networking.IntPtrToStrPtr(request.MaximumAutoVerificationFrequencyInDays)).
        WithOptionalQueryParam("media_ejection_allowed", networking.BoolPtrToStrPtr(request.MediaEjectionAllowed)).
        WithOptionalQueryParam("name", request.Name).
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
    return models.NewModifyStorageDomainSpectraS3Response(response)
}

func (client *Client) ForceFeatureKeyValidationSpectraS3(request *models.ForceFeatureKeyValidationSpectraS3Request) (*models.ForceFeatureKeyValidationSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/feature_key").
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
    return models.NewForceFeatureKeyValidationSpectraS3Response(response)
}

func (client *Client) ResetInstanceIdentifierSpectraS3(request *models.ResetInstanceIdentifierSpectraS3Request) (*models.ResetInstanceIdentifierSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/instance_identifier").
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
    return models.NewResetInstanceIdentifierSpectraS3Response(response)
}

func (client *Client) CancelEjectOnAllTapesSpectraS3(request *models.CancelEjectOnAllTapesSpectraS3Request) (*models.CancelEjectOnAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithQueryParam("operation", "cancel_eject").
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
    return models.NewCancelEjectOnAllTapesSpectraS3Response(response)
}

func (client *Client) CancelEjectTapeSpectraS3(request *models.CancelEjectTapeSpectraS3Request) (*models.CancelEjectTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithQueryParam("operation", "cancel_eject").
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
    return models.NewCancelEjectTapeSpectraS3Response(response)
}

func (client *Client) CancelFormatOnAllTapesSpectraS3(request *models.CancelFormatOnAllTapesSpectraS3Request) (*models.CancelFormatOnAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithQueryParam("operation", "cancel_format").
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
    return models.NewCancelFormatOnAllTapesSpectraS3Response(response)
}

func (client *Client) CancelFormatTapeSpectraS3(request *models.CancelFormatTapeSpectraS3Request) (*models.CancelFormatTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithQueryParam("operation", "cancel_format").
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
    return models.NewCancelFormatTapeSpectraS3Response(response)
}

func (client *Client) CancelImportOnAllTapesSpectraS3(request *models.CancelImportOnAllTapesSpectraS3Request) (*models.CancelImportOnAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithQueryParam("operation", "cancel_import").
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
    return models.NewCancelImportOnAllTapesSpectraS3Response(response)
}

func (client *Client) CancelImportTapeSpectraS3(request *models.CancelImportTapeSpectraS3Request) (*models.CancelImportTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithQueryParam("operation", "cancel_import").
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
    return models.NewCancelImportTapeSpectraS3Response(response)
}

func (client *Client) CancelOnlineOnAllTapesSpectraS3(request *models.CancelOnlineOnAllTapesSpectraS3Request) (*models.CancelOnlineOnAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithQueryParam("operation", "cancel_online").
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
    return models.NewCancelOnlineOnAllTapesSpectraS3Response(response)
}

func (client *Client) CancelOnlineTapeSpectraS3(request *models.CancelOnlineTapeSpectraS3Request) (*models.CancelOnlineTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithQueryParam("operation", "cancel_online").
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
    return models.NewCancelOnlineTapeSpectraS3Response(response)
}

func (client *Client) CancelVerifyOnAllTapesSpectraS3(request *models.CancelVerifyOnAllTapesSpectraS3Request) (*models.CancelVerifyOnAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithQueryParam("operation", "cancel_verify").
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
    return models.NewCancelVerifyOnAllTapesSpectraS3Response(response)
}

func (client *Client) CancelVerifyTapeSpectraS3(request *models.CancelVerifyTapeSpectraS3Request) (*models.CancelVerifyTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithQueryParam("operation", "cancel_verify").
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
    return models.NewCancelVerifyTapeSpectraS3Response(response)
}

func (client *Client) CleanTapeDriveSpectraS3(request *models.CleanTapeDriveSpectraS3Request) (*models.CleanTapeDriveSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape_drive/" + request.TapeDriveId).
        WithQueryParam("operation", "clean").
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
    return models.NewCleanTapeDriveSpectraS3Response(response)
}

func (client *Client) EjectAllTapesSpectraS3(request *models.EjectAllTapesSpectraS3Request) (*models.EjectAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithOptionalQueryParam("eject_label", request.EjectLabel).
        WithOptionalQueryParam("eject_location", request.EjectLocation).
        WithQueryParam("operation", "eject").
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
    return models.NewEjectAllTapesSpectraS3Response(response)
}

func (client *Client) EjectStorageDomainBlobsSpectraS3(request *models.EjectStorageDomainBlobsSpectraS3Request) (*models.EjectStorageDomainBlobsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithQueryParam("blobs", "").
        WithQueryParam("bucket_id", request.BucketId).
        WithQueryParam("storage_domain", request.StorageDomain).
        WithOptionalQueryParam("eject_label", request.EjectLabel).
        WithOptionalQueryParam("eject_location", request.EjectLocation).
        WithQueryParam("operation", "eject").
        WithReadCloser(buildDs3GetObjectListStream(request.Objects)).
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
    return models.NewEjectStorageDomainBlobsSpectraS3Response(response)
}

func (client *Client) EjectStorageDomainSpectraS3(request *models.EjectStorageDomainSpectraS3Request) (*models.EjectStorageDomainSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithQueryParam("storage_domain", request.StorageDomain).
        WithOptionalQueryParam("bucket_id", request.BucketId).
        WithOptionalQueryParam("eject_label", request.EjectLabel).
        WithOptionalQueryParam("eject_location", request.EjectLocation).
        WithQueryParam("operation", "eject").
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
    return models.NewEjectStorageDomainSpectraS3Response(response)
}

func (client *Client) EjectTapeSpectraS3(request *models.EjectTapeSpectraS3Request) (*models.EjectTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithOptionalQueryParam("eject_label", request.EjectLabel).
        WithOptionalQueryParam("eject_location", request.EjectLocation).
        WithQueryParam("operation", "eject").
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
    return models.NewEjectTapeSpectraS3Response(response)
}

func (client *Client) ForceTapeEnvironmentRefreshSpectraS3(request *models.ForceTapeEnvironmentRefreshSpectraS3Request) (*models.ForceTapeEnvironmentRefreshSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape_environment").
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
    return models.NewForceTapeEnvironmentRefreshSpectraS3Response(response)
}

func (client *Client) FormatAllTapesSpectraS3(request *models.FormatAllTapesSpectraS3Request) (*models.FormatAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithOptionalVoidQueryParam("force", request.Force).
        WithQueryParam("operation", "format").
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
    return models.NewFormatAllTapesSpectraS3Response(response)
}

func (client *Client) FormatTapeSpectraS3(request *models.FormatTapeSpectraS3Request) (*models.FormatTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithOptionalVoidQueryParam("force", request.Force).
        WithQueryParam("operation", "format").
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
    return models.NewFormatTapeSpectraS3Response(response)
}

func (client *Client) ImportAllTapesSpectraS3(request *models.ImportAllTapesSpectraS3Request) (*models.ImportAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithOptionalQueryParam("storage_domain_id", request.StorageDomainId).
        WithOptionalQueryParam("user_id", request.UserId).
        WithOptionalQueryParam("verify_data_after_import", networking.InterfaceToStrPtr(request.VerifyDataAfterImport)).
        WithOptionalQueryParam("verify_data_prior_to_import", networking.BoolPtrToStrPtr(request.VerifyDataPriorToImport)).
        WithQueryParam("operation", "import").
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
    return models.NewImportAllTapesSpectraS3Response(response)
}

func (client *Client) ImportTapeSpectraS3(request *models.ImportTapeSpectraS3Request) (*models.ImportTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithOptionalQueryParam("storage_domain_id", request.StorageDomainId).
        WithOptionalQueryParam("user_id", request.UserId).
        WithOptionalQueryParam("verify_data_after_import", networking.InterfaceToStrPtr(request.VerifyDataAfterImport)).
        WithOptionalQueryParam("verify_data_prior_to_import", networking.BoolPtrToStrPtr(request.VerifyDataPriorToImport)).
        WithQueryParam("operation", "import").
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
    return models.NewImportTapeSpectraS3Response(response)
}

func (client *Client) InspectAllTapesSpectraS3(request *models.InspectAllTapesSpectraS3Request) (*models.InspectAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithOptionalQueryParam("task_priority", networking.InterfaceToStrPtr(request.TaskPriority)).
        WithQueryParam("operation", "inspect").
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
    return models.NewInspectAllTapesSpectraS3Response(response)
}

func (client *Client) InspectTapeSpectraS3(request *models.InspectTapeSpectraS3Request) (*models.InspectTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithOptionalQueryParam("task_priority", networking.InterfaceToStrPtr(request.TaskPriority)).
        WithQueryParam("operation", "inspect").
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
    return models.NewInspectTapeSpectraS3Response(response)
}

func (client *Client) ModifyAllTapePartitionsSpectraS3(request *models.ModifyAllTapePartitionsSpectraS3Request) (*models.ModifyAllTapePartitionsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape_partition").
        WithQueryParam("quiesced", request.Quiesced.String()).
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
    return models.NewModifyAllTapePartitionsSpectraS3Response(response)
}

func (client *Client) ModifyTapeDriveSpectraS3(request *models.ModifyTapeDriveSpectraS3Request) (*models.ModifyTapeDriveSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape_drive/" + request.TapeDriveId).
        WithOptionalQueryParam("minimum_task_priority", networking.InterfaceToStrPtr(request.MinimumTaskPriority)).
        WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
        WithOptionalQueryParam("reserved_task_type", networking.InterfaceToStrPtr(request.ReservedTaskType)).
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
    return models.NewModifyTapeDriveSpectraS3Response(response)
}

func (client *Client) ModifyTapePartitionSpectraS3(request *models.ModifyTapePartitionSpectraS3Request) (*models.ModifyTapePartitionSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape_partition/" + request.TapePartition).
        WithOptionalQueryParam("auto_compaction_enabled", networking.BoolPtrToStrPtr(request.AutoCompactionEnabled)).
        WithOptionalQueryParam("minimum_read_reserved_drives", networking.IntPtrToStrPtr(request.MinimumReadReservedDrives)).
        WithOptionalQueryParam("minimum_write_reserved_drives", networking.IntPtrToStrPtr(request.MinimumWriteReservedDrives)).
        WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
        WithOptionalQueryParam("serial_number", request.SerialNumber).
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
    return models.NewModifyTapePartitionSpectraS3Response(response)
}

func (client *Client) ModifyTapeSpectraS3(request *models.ModifyTapeSpectraS3Request) (*models.ModifyTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithOptionalQueryParam("eject_label", request.EjectLabel).
        WithOptionalQueryParam("eject_location", request.EjectLocation).
        WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
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
    return models.NewModifyTapeSpectraS3Response(response)
}

func (client *Client) OnlineAllTapesSpectraS3(request *models.OnlineAllTapesSpectraS3Request) (*models.OnlineAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithQueryParam("operation", "online").
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
    return models.NewOnlineAllTapesSpectraS3Response(response)
}

func (client *Client) OnlineTapeSpectraS3(request *models.OnlineTapeSpectraS3Request) (*models.OnlineTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithQueryParam("operation", "online").
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
    return models.NewOnlineTapeSpectraS3Response(response)
}

func (client *Client) RawImportAllTapesSpectraS3(request *models.RawImportAllTapesSpectraS3Request) (*models.RawImportAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithQueryParam("bucket_id", request.BucketId).
        WithOptionalQueryParam("storage_domain_id", request.StorageDomainId).
        WithOptionalQueryParam("task_priority", networking.InterfaceToStrPtr(request.TaskPriority)).
        WithQueryParam("operation", "import").
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
    return models.NewRawImportAllTapesSpectraS3Response(response)
}

func (client *Client) RawImportTapeSpectraS3(request *models.RawImportTapeSpectraS3Request) (*models.RawImportTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithQueryParam("bucket_id", request.BucketId).
        WithOptionalQueryParam("storage_domain_id", request.StorageDomainId).
        WithOptionalQueryParam("task_priority", networking.InterfaceToStrPtr(request.TaskPriority)).
        WithQueryParam("operation", "import").
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
    return models.NewRawImportTapeSpectraS3Response(response)
}

func (client *Client) VerifyAllTapesSpectraS3(request *models.VerifyAllTapesSpectraS3Request) (*models.VerifyAllTapesSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape").
        WithOptionalQueryParam("task_priority", networking.InterfaceToStrPtr(request.TaskPriority)).
        WithQueryParam("operation", "verify").
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
    return models.NewVerifyAllTapesSpectraS3Response(response)
}

func (client *Client) VerifyTapeSpectraS3(request *models.VerifyTapeSpectraS3Request) (*models.VerifyTapeSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/tape/" + request.TapeId).
        WithOptionalQueryParam("task_priority", networking.InterfaceToStrPtr(request.TaskPriority)).
        WithQueryParam("operation", "verify").
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
    return models.NewVerifyTapeSpectraS3Response(response)
}

func (client *Client) ForceTargetEnvironmentRefreshSpectraS3(request *models.ForceTargetEnvironmentRefreshSpectraS3Request) (*models.ForceTargetEnvironmentRefreshSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/target_environment").
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
    return models.NewForceTargetEnvironmentRefreshSpectraS3Response(response)
}

func (client *Client) ImportAzureTargetSpectraS3(request *models.ImportAzureTargetSpectraS3Request) (*models.ImportAzureTargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/azure_target/" + request.AzureTarget).
        WithQueryParam("cloud_bucket_name", request.CloudBucketName).
        WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithOptionalQueryParam("user_id", request.UserId).
        WithQueryParam("operation", "import").
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
    return models.NewImportAzureTargetSpectraS3Response(response)
}

func (client *Client) ModifyAllAzureTargetsSpectraS3(request *models.ModifyAllAzureTargetsSpectraS3Request) (*models.ModifyAllAzureTargetsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/azure_target").
        WithQueryParam("quiesced", request.Quiesced.String()).
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
    return models.NewModifyAllAzureTargetsSpectraS3Response(response)
}

func (client *Client) ModifyAzureTargetSpectraS3(request *models.ModifyAzureTargetSpectraS3Request) (*models.ModifyAzureTargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/azure_target/" + request.AzureTarget).
        WithOptionalQueryParam("account_key", request.AccountKey).
        WithOptionalQueryParam("account_name", request.AccountName).
        WithOptionalQueryParam("auto_verify_frequency_in_days", networking.IntPtrToStrPtr(request.AutoVerifyFrequencyInDays)).
        WithOptionalQueryParam("cloud_bucket_prefix", request.CloudBucketPrefix).
        WithOptionalQueryParam("cloud_bucket_suffix", request.CloudBucketSuffix).
        WithOptionalQueryParam("default_read_preference", networking.InterfaceToStrPtr(request.DefaultReadPreference)).
        WithOptionalQueryParam("https", networking.BoolPtrToStrPtr(request.Https)).
        WithOptionalQueryParam("name", request.Name).
        WithOptionalQueryParam("permit_going_out_of_sync", networking.BoolPtrToStrPtr(request.PermitGoingOutOfSync)).
        WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
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
    return models.NewModifyAzureTargetSpectraS3Response(response)
}

func (client *Client) VerifyAzureTargetSpectraS3(request *models.VerifyAzureTargetSpectraS3Request) (*models.VerifyAzureTargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/azure_target/" + request.AzureTarget).
        WithOptionalVoidQueryParam("full_details", request.FullDetails).
        WithQueryParam("operation", "verify").
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
    return models.NewVerifyAzureTargetSpectraS3Response(response)
}

func (client *Client) ModifyAllDs3TargetsSpectraS3(request *models.ModifyAllDs3TargetsSpectraS3Request) (*models.ModifyAllDs3TargetsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/ds3_target").
        WithQueryParam("quiesced", request.Quiesced.String()).
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
    return models.NewModifyAllDs3TargetsSpectraS3Response(response)
}

func (client *Client) ModifyDs3TargetSpectraS3(request *models.ModifyDs3TargetSpectraS3Request) (*models.ModifyDs3TargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/ds3_target/" + request.Ds3Target).
        WithOptionalQueryParam("access_control_replication", networking.InterfaceToStrPtr(request.AccessControlReplication)).
        WithOptionalQueryParam("admin_auth_id", request.AdminAuthId).
        WithOptionalQueryParam("admin_secret_key", request.AdminSecretKey).
        WithOptionalQueryParam("data_path_end_point", request.DataPathEndPoint).
        WithOptionalQueryParam("data_path_https", networking.BoolPtrToStrPtr(request.DataPathHttps)).
        WithOptionalQueryParam("data_path_port", networking.IntPtrToStrPtr(request.DataPathPort)).
        WithOptionalQueryParam("data_path_proxy", request.DataPathProxy).
        WithOptionalQueryParam("data_path_verify_certificate", networking.BoolPtrToStrPtr(request.DataPathVerifyCertificate)).
        WithOptionalQueryParam("default_read_preference", networking.InterfaceToStrPtr(request.DefaultReadPreference)).
        WithOptionalQueryParam("name", request.Name).
        WithOptionalQueryParam("permit_going_out_of_sync", networking.BoolPtrToStrPtr(request.PermitGoingOutOfSync)).
        WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
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
    return models.NewModifyDs3TargetSpectraS3Response(response)
}

func (client *Client) PairBackRegisteredDs3TargetSpectraS3(request *models.PairBackRegisteredDs3TargetSpectraS3Request) (*models.PairBackRegisteredDs3TargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/ds3_target/" + request.Ds3Target).
        WithOptionalQueryParam("access_control_replication", networking.InterfaceToStrPtr(request.AccessControlReplication)).
        WithOptionalQueryParam("admin_auth_id", request.AdminAuthId).
        WithOptionalQueryParam("admin_secret_key", request.AdminSecretKey).
        WithOptionalQueryParam("data_path_end_point", request.DataPathEndPoint).
        WithOptionalQueryParam("data_path_https", networking.BoolPtrToStrPtr(request.DataPathHttps)).
        WithOptionalQueryParam("data_path_port", networking.IntPtrToStrPtr(request.DataPathPort)).
        WithOptionalQueryParam("data_path_proxy", request.DataPathProxy).
        WithOptionalQueryParam("data_path_verify_certificate", networking.BoolPtrToStrPtr(request.DataPathVerifyCertificate)).
        WithOptionalQueryParam("default_read_preference", networking.InterfaceToStrPtr(request.DefaultReadPreference)).
        WithOptionalQueryParam("name", request.Name).
        WithOptionalQueryParam("permit_going_out_of_sync", networking.BoolPtrToStrPtr(request.PermitGoingOutOfSync)).
        WithOptionalQueryParam("replicated_user_default_data_policy", request.ReplicatedUserDefaultDataPolicy).
        WithQueryParam("operation", "pair_back").
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
    return models.NewPairBackRegisteredDs3TargetSpectraS3Response(response)
}

func (client *Client) VerifyDs3TargetSpectraS3(request *models.VerifyDs3TargetSpectraS3Request) (*models.VerifyDs3TargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/ds3_target/" + request.Ds3Target).
        WithOptionalVoidQueryParam("full_details", request.FullDetails).
        WithQueryParam("operation", "verify").
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
    return models.NewVerifyDs3TargetSpectraS3Response(response)
}

func (client *Client) ImportS3TargetSpectraS3(request *models.ImportS3TargetSpectraS3Request) (*models.ImportS3TargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/s3_target/" + request.S3Target).
        WithQueryParam("cloud_bucket_name", request.CloudBucketName).
        WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
        WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
        WithOptionalQueryParam("user_id", request.UserId).
        WithQueryParam("operation", "import").
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
    return models.NewImportS3TargetSpectraS3Response(response)
}

func (client *Client) ModifyAllS3TargetsSpectraS3(request *models.ModifyAllS3TargetsSpectraS3Request) (*models.ModifyAllS3TargetsSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/s3_target").
        WithQueryParam("quiesced", request.Quiesced.String()).
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
    return models.NewModifyAllS3TargetsSpectraS3Response(response)
}

func (client *Client) ModifyS3TargetSpectraS3(request *models.ModifyS3TargetSpectraS3Request) (*models.ModifyS3TargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/s3_target/" + request.S3Target).
        WithOptionalQueryParam("access_key", request.AccessKey).
        WithOptionalQueryParam("auto_verify_frequency_in_days", networking.IntPtrToStrPtr(request.AutoVerifyFrequencyInDays)).
        WithOptionalQueryParam("cloud_bucket_prefix", request.CloudBucketPrefix).
        WithOptionalQueryParam("cloud_bucket_suffix", request.CloudBucketSuffix).
        WithOptionalQueryParam("data_path_end_point", request.DataPathEndPoint).
        WithOptionalQueryParam("default_read_preference", networking.InterfaceToStrPtr(request.DefaultReadPreference)).
        WithOptionalQueryParam("https", networking.BoolPtrToStrPtr(request.Https)).
        WithOptionalQueryParam("name", request.Name).
        WithOptionalQueryParam("naming_mode", networking.InterfaceToStrPtr(request.NamingMode)).
        WithOptionalQueryParam("offline_data_staging_window_in_tb", networking.IntPtrToStrPtr(request.OfflineDataStagingWindowInTb)).
        WithOptionalQueryParam("permit_going_out_of_sync", networking.BoolPtrToStrPtr(request.PermitGoingOutOfSync)).
        WithOptionalQueryParam("proxy_domain", request.ProxyDomain).
        WithOptionalQueryParam("proxy_host", request.ProxyHost).
        WithOptionalQueryParam("proxy_password", request.ProxyPassword).
        WithOptionalQueryParam("proxy_port", networking.IntPtrToStrPtr(request.ProxyPort)).
        WithOptionalQueryParam("proxy_username", request.ProxyUsername).
        WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
        WithOptionalQueryParam("region", networking.InterfaceToStrPtr(request.Region)).
        WithOptionalQueryParam("restricted_access", networking.BoolPtrToStrPtr(request.RestrictedAccess)).
        WithOptionalQueryParam("secret_key", request.SecretKey).
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
    return models.NewModifyS3TargetSpectraS3Response(response)
}

func (client *Client) VerifyS3TargetSpectraS3(request *models.VerifyS3TargetSpectraS3Request) (*models.VerifyS3TargetSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/s3_target/" + request.S3Target).
        WithOptionalVoidQueryParam("full_details", request.FullDetails).
        WithQueryParam("operation", "verify").
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
    return models.NewVerifyS3TargetSpectraS3Response(response)
}

func (client *Client) ModifyUserSpectraS3(request *models.ModifyUserSpectraS3Request) (*models.ModifyUserSpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/user/" + request.UserId).
        WithOptionalQueryParam("default_data_policy_id", request.DefaultDataPolicyId).
        WithOptionalQueryParam("max_buckets", networking.IntPtrToStrPtr(request.MaxBuckets)).
        WithOptionalQueryParam("name", request.Name).
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
    return models.NewModifyUserSpectraS3Response(response)
}

func (client *Client) RegenerateUserSecretKeySpectraS3(request *models.RegenerateUserSecretKeySpectraS3Request) (*models.RegenerateUserSecretKeySpectraS3Response, error) {
    // Build the http request
    httpRequest, err := networking.NewHttpRequestBuilder().
        WithHttpVerb(HTTP_VERB_PUT).
        WithPath("/_rest_/user/" + request.UserId).
        WithQueryParam("operation", "regenerate_secret_key").
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
    return models.NewRegenerateUserSecretKeySpectraS3Response(response)
}


