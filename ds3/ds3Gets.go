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
	"context"
	"github.com/SpectraLogic/ds3_go_sdk/ds3/models"
	"github.com/SpectraLogic/ds3_go_sdk/ds3/networking"
)

func (client *Client) GetObject(ctx context.Context, request *models.GetObjectRequest) (*models.GetObjectResponse, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/"+request.BucketName+"/"+request.ObjectName).
		WithOptionalVoidQueryParam("cached_only", request.CachedOnly).
		WithOptionalQueryParam("job", request.Job).
		WithOptionalQueryParam("offset", networking.Int64PtrToStrPtr(request.Offset)).
		WithOptionalQueryParam("version_id", request.VersionId).
		WithChecksum(request.Checksum).
		WithHeaders(request.Metadata).
		Build(ctx, client.connectionInfo)

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
	return models.NewGetObjectResponse(response)
}

func (client *Client) GetBucket(ctx context.Context, request *models.GetBucketRequest) (*models.GetBucketResponse, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/"+request.BucketName).
		WithOptionalQueryParam("delimiter", request.Delimiter).
		WithOptionalQueryParam("marker", request.Marker).
		WithOptionalQueryParam("max_keys", networking.IntPtrToStrPtr(request.MaxKeys)).
		WithOptionalQueryParam("prefix", request.Prefix).
		WithOptionalVoidQueryParam("versions", request.Versions).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBucketResponse(response, client.Logger)
}

func (client *Client) GetService(ctx context.Context, request *models.GetServiceRequest) (*models.GetServiceResponse, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetServiceResponse(response, client.Logger)
}

func (client *Client) GetBucketAclSpectraS3(ctx context.Context, request *models.GetBucketAclSpectraS3Request) (*models.GetBucketAclSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/bucket_acl/"+request.BucketAcl).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBucketAclSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBucketAclsSpectraS3(ctx context.Context, request *models.GetBucketAclsSpectraS3Request) (*models.GetBucketAclsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/bucket_acl").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("group_id", request.GroupId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("permission", networking.InterfaceToStrPtr(request.Permission)).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBucketAclsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDataPolicyAclSpectraS3(ctx context.Context, request *models.GetDataPolicyAclSpectraS3Request) (*models.GetDataPolicyAclSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/data_policy_acl/"+request.DataPolicyAcl).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDataPolicyAclSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDataPolicyAclsSpectraS3(ctx context.Context, request *models.GetDataPolicyAclsSpectraS3Request) (*models.GetDataPolicyAclsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/data_policy_acl").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalQueryParam("group_id", request.GroupId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDataPolicyAclsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBucketSpectraS3(ctx context.Context, request *models.GetBucketSpectraS3Request) (*models.GetBucketSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/bucket/"+request.BucketName).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBucketSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBucketsSpectraS3(ctx context.Context, request *models.GetBucketsSpectraS3Request) (*models.GetBucketsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/bucket").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBucketsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetCacheFilesystemSpectraS3(ctx context.Context, request *models.GetCacheFilesystemSpectraS3Request) (*models.GetCacheFilesystemSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/cache_filesystem/"+request.CacheFilesystem).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetCacheFilesystemSpectraS3Response(response, client.Logger)
}

func (client *Client) GetCacheFilesystemsSpectraS3(ctx context.Context, request *models.GetCacheFilesystemsSpectraS3Request) (*models.GetCacheFilesystemsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/cache_filesystem").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("node_id", request.NodeId).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetCacheFilesystemsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetCacheStateSpectraS3(ctx context.Context, request *models.GetCacheStateSpectraS3Request) (*models.GetCacheStateSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/cache_state").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetCacheStateSpectraS3Response(response, client.Logger)
}

func (client *Client) GetCacheThrottleRuleSpectraS3(ctx context.Context, request *models.GetCacheThrottleRuleSpectraS3Request) (*models.GetCacheThrottleRuleSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/cache_throttle_rule/"+request.CacheThrottleRule).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetCacheThrottleRuleSpectraS3Response(response, client.Logger)
}

func (client *Client) GetCacheThrottleRulesSpectraS3(ctx context.Context, request *models.GetCacheThrottleRulesSpectraS3Request) (*models.GetCacheThrottleRulesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/cache_throttle_rule").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("burst_threshold", networking.Float64PtrToStrPtr(request.BurstThreshold)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("max_cache_percent", networking.Float64PtrToStrPtr(request.MaxCachePercent)).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
		WithOptionalQueryParam("request_type", networking.InterfaceToStrPtr(request.RequestType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetCacheThrottleRulesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBucketCapacitySummarySpectraS3(ctx context.Context, request *models.GetBucketCapacitySummarySpectraS3Request) (*models.GetBucketCapacitySummarySpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/capacity_summary").
		WithQueryParam("bucket_id", request.BucketId).
		WithQueryParam("storage_domain_id", request.StorageDomainId).
		WithOptionalQueryParam("pool_health", networking.InterfaceToStrPtr(request.PoolHealth)).
		WithOptionalQueryParam("pool_state", networking.InterfaceToStrPtr(request.PoolState)).
		WithOptionalQueryParam("pool_type", networking.InterfaceToStrPtr(request.PoolType)).
		WithOptionalQueryParam("tape_state", networking.InterfaceToStrPtr(request.TapeState)).
		WithOptionalQueryParam("tape_type", request.TapeType).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBucketCapacitySummarySpectraS3Response(response, client.Logger)
}

func (client *Client) GetStorageDomainCapacitySummarySpectraS3(ctx context.Context, request *models.GetStorageDomainCapacitySummarySpectraS3Request) (*models.GetStorageDomainCapacitySummarySpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/capacity_summary").
		WithQueryParam("storage_domain_id", request.StorageDomainId).
		WithOptionalQueryParam("pool_health", networking.InterfaceToStrPtr(request.PoolHealth)).
		WithOptionalQueryParam("pool_state", networking.InterfaceToStrPtr(request.PoolState)).
		WithOptionalQueryParam("pool_type", networking.InterfaceToStrPtr(request.PoolType)).
		WithOptionalQueryParam("tape_state", networking.InterfaceToStrPtr(request.TapeState)).
		WithOptionalQueryParam("tape_type", request.TapeType).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetStorageDomainCapacitySummarySpectraS3Response(response, client.Logger)
}

func (client *Client) GetSystemCapacitySummarySpectraS3(ctx context.Context, request *models.GetSystemCapacitySummarySpectraS3Request) (*models.GetSystemCapacitySummarySpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/capacity_summary").
		WithOptionalQueryParam("pool_health", networking.InterfaceToStrPtr(request.PoolHealth)).
		WithOptionalQueryParam("pool_state", networking.InterfaceToStrPtr(request.PoolState)).
		WithOptionalQueryParam("pool_type", networking.InterfaceToStrPtr(request.PoolType)).
		WithOptionalQueryParam("tape_state", networking.InterfaceToStrPtr(request.TapeState)).
		WithOptionalQueryParam("tape_type", request.TapeType).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSystemCapacitySummarySpectraS3Response(response, client.Logger)
}

func (client *Client) GetDataPathBackendSpectraS3(ctx context.Context, request *models.GetDataPathBackendSpectraS3Request) (*models.GetDataPathBackendSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/data_path_backend").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDataPathBackendSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDataPlannerBlobStoreTasksSpectraS3(ctx context.Context, request *models.GetDataPlannerBlobStoreTasksSpectraS3Request) (*models.GetDataPlannerBlobStoreTasksSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/blob_store_task").
		WithOptionalVoidQueryParam("full_details", request.FullDetails).
		WithOptionalQueryParam("job", request.Job).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDataPlannerBlobStoreTasksSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAzureDataReplicationRuleSpectraS3(ctx context.Context, request *models.GetAzureDataReplicationRuleSpectraS3Request) (*models.GetAzureDataReplicationRuleSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_data_replication_rule/"+request.AzureDataReplicationRule).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAzureDataReplicationRuleSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAzureDataReplicationRulesSpectraS3(ctx context.Context, request *models.GetAzureDataReplicationRulesSpectraS3Request) (*models.GetAzureDataReplicationRulesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_data_replication_rule").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("replicate_deletes", networking.BoolPtrToStrPtr(request.ReplicateDeletes)).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("target_id", request.TargetId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataReplicationRuleType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAzureDataReplicationRulesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDataPersistenceRuleSpectraS3(ctx context.Context, request *models.GetDataPersistenceRuleSpectraS3Request) (*models.GetDataPersistenceRuleSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/data_persistence_rule/"+request.DataPersistenceRuleId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDataPersistenceRuleSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDataPersistenceRulesSpectraS3(ctx context.Context, request *models.GetDataPersistenceRulesSpectraS3Request) (*models.GetDataPersistenceRulesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/data_persistence_rule").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalQueryParam("isolation_level", networking.InterfaceToStrPtr(request.IsolationLevel)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("storage_domain_id", request.StorageDomainId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataPersistenceRuleType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDataPersistenceRulesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDataPoliciesSpectraS3(ctx context.Context, request *models.GetDataPoliciesSpectraS3Request) (*models.GetDataPoliciesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/data_policy").
		WithOptionalQueryParam("always_force_put_job_creation", networking.BoolPtrToStrPtr(request.AlwaysForcePutJobCreation)).
		WithOptionalQueryParam("always_minimize_spanning_across_media", networking.BoolPtrToStrPtr(request.AlwaysMinimizeSpanningAcrossMedia)).
		WithOptionalQueryParam("checksum_type", networking.InterfaceToStrPtr(request.ChecksumType)).
		WithOptionalQueryParam("end_to_end_crc_required", networking.BoolPtrToStrPtr(request.EndToEndCrcRequired)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDataPoliciesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDataPolicySpectraS3(ctx context.Context, request *models.GetDataPolicySpectraS3Request) (*models.GetDataPolicySpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/data_policy/"+request.DataPolicyId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDataPolicySpectraS3Response(response, client.Logger)
}

func (client *Client) GetDs3DataReplicationRuleSpectraS3(ctx context.Context, request *models.GetDs3DataReplicationRuleSpectraS3Request) (*models.GetDs3DataReplicationRuleSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_data_replication_rule/"+request.Ds3DataReplicationRule).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDs3DataReplicationRuleSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDs3DataReplicationRulesSpectraS3(ctx context.Context, request *models.GetDs3DataReplicationRulesSpectraS3Request) (*models.GetDs3DataReplicationRulesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_data_replication_rule").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("replicate_deletes", networking.BoolPtrToStrPtr(request.ReplicateDeletes)).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("target_id", request.TargetId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataReplicationRuleType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDs3DataReplicationRulesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetS3DataReplicationRuleSpectraS3(ctx context.Context, request *models.GetS3DataReplicationRuleSpectraS3Request) (*models.GetS3DataReplicationRuleSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_data_replication_rule/"+request.S3DataReplicationRule).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetS3DataReplicationRuleSpectraS3Response(response, client.Logger)
}

func (client *Client) GetS3DataReplicationRulesSpectraS3(ctx context.Context, request *models.GetS3DataReplicationRulesSpectraS3Request) (*models.GetS3DataReplicationRulesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_data_replication_rule").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalQueryParam("initial_data_placement", networking.InterfaceToStrPtr(request.InitialDataPlacement)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("replicate_deletes", networking.BoolPtrToStrPtr(request.ReplicateDeletes)).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("target_id", request.TargetId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataReplicationRuleType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetS3DataReplicationRulesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDegradedAzureDataReplicationRulesSpectraS3(ctx context.Context, request *models.GetDegradedAzureDataReplicationRulesSpectraS3Request) (*models.GetDegradedAzureDataReplicationRulesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/degraded_azure_data_replication_rule").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("target_id", request.TargetId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataReplicationRuleType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDegradedAzureDataReplicationRulesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDegradedBlobsSpectraS3(ctx context.Context, request *models.GetDegradedBlobsSpectraS3Request) (*models.GetDegradedBlobsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/degraded_blob").
		WithOptionalQueryParam("blob_id", request.BlobId).
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("ds3_replication_rule_id", request.Ds3ReplicationRuleId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("persistence_rule_id", request.PersistenceRuleId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDegradedBlobsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDegradedBucketsSpectraS3(ctx context.Context, request *models.GetDegradedBucketsSpectraS3Request) (*models.GetDegradedBucketsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/degraded_bucket").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDegradedBucketsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDegradedDataPersistenceRulesSpectraS3(ctx context.Context, request *models.GetDegradedDataPersistenceRulesSpectraS3Request) (*models.GetDegradedDataPersistenceRulesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/degraded_data_persistence_rule").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalQueryParam("isolation_level", networking.InterfaceToStrPtr(request.IsolationLevel)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("storage_domain_id", request.StorageDomainId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataPersistenceRuleType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDegradedDataPersistenceRulesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDegradedDs3DataReplicationRulesSpectraS3(ctx context.Context, request *models.GetDegradedDs3DataReplicationRulesSpectraS3Request) (*models.GetDegradedDs3DataReplicationRulesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/degraded_ds3_data_replication_rule").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("target_id", request.TargetId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataReplicationRuleType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDegradedDs3DataReplicationRulesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDegradedS3DataReplicationRulesSpectraS3(ctx context.Context, request *models.GetDegradedS3DataReplicationRulesSpectraS3Request) (*models.GetDegradedS3DataReplicationRulesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/degraded_s3_data_replication_rule").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("target_id", request.TargetId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.DataReplicationRuleType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDegradedS3DataReplicationRulesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSuspectBlobAzureTargetsSpectraS3(ctx context.Context, request *models.GetSuspectBlobAzureTargetsSpectraS3Request) (*models.GetSuspectBlobAzureTargetsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/suspect_blob_azure_target").
		WithOptionalQueryParam("blob_id", request.BlobId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("target_id", request.TargetId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSuspectBlobAzureTargetsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSuspectBlobDs3TargetsSpectraS3(ctx context.Context, request *models.GetSuspectBlobDs3TargetsSpectraS3Request) (*models.GetSuspectBlobDs3TargetsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/suspect_blob_ds3_target").
		WithOptionalQueryParam("blob_id", request.BlobId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("target_id", request.TargetId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSuspectBlobDs3TargetsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSuspectBlobPoolsSpectraS3(ctx context.Context, request *models.GetSuspectBlobPoolsSpectraS3Request) (*models.GetSuspectBlobPoolsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/suspect_blob_pool").
		WithOptionalQueryParam("blob_id", request.BlobId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("pool_id", request.PoolId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSuspectBlobPoolsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSuspectBlobS3TargetsSpectraS3(ctx context.Context, request *models.GetSuspectBlobS3TargetsSpectraS3Request) (*models.GetSuspectBlobS3TargetsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/suspect_blob_s3_target").
		WithOptionalQueryParam("blob_id", request.BlobId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("target_id", request.TargetId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSuspectBlobS3TargetsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSuspectBlobTapesSpectraS3(ctx context.Context, request *models.GetSuspectBlobTapesSpectraS3Request) (*models.GetSuspectBlobTapesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/suspect_blob_tape").
		WithOptionalQueryParam("blob_id", request.BlobId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("tape_id", request.TapeId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSuspectBlobTapesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSuspectBucketsSpectraS3(ctx context.Context, request *models.GetSuspectBucketsSpectraS3Request) (*models.GetSuspectBucketsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/suspect_bucket").
		WithOptionalQueryParam("data_policy_id", request.DataPolicyId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSuspectBucketsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSuspectObjectsSpectraS3(ctx context.Context, request *models.GetSuspectObjectsSpectraS3Request) (*models.GetSuspectObjectsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/suspect_object").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSuspectObjectsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSuspectObjectsWithFullDetailsSpectraS3(ctx context.Context, request *models.GetSuspectObjectsWithFullDetailsSpectraS3Request) (*models.GetSuspectObjectsWithFullDetailsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/suspect_object").
		WithQueryParam("full_details", "").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("storage_domain", request.StorageDomain).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSuspectObjectsWithFullDetailsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetGroupMemberSpectraS3(ctx context.Context, request *models.GetGroupMemberSpectraS3Request) (*models.GetGroupMemberSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/group_member/"+request.GroupMember).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetGroupMemberSpectraS3Response(response, client.Logger)
}

func (client *Client) GetGroupMembersSpectraS3(ctx context.Context, request *models.GetGroupMembersSpectraS3Request) (*models.GetGroupMembersSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/group_member").
		WithOptionalQueryParam("group_id", request.GroupId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("member_group_id", request.MemberGroupId).
		WithOptionalQueryParam("member_user_id", request.MemberUserId).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetGroupMembersSpectraS3Response(response, client.Logger)
}

func (client *Client) GetGroupSpectraS3(ctx context.Context, request *models.GetGroupSpectraS3Request) (*models.GetGroupSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/group/"+request.Group).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetGroupSpectraS3Response(response, client.Logger)
}

func (client *Client) GetGroupsSpectraS3(ctx context.Context, request *models.GetGroupsSpectraS3Request) (*models.GetGroupsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/group").
		WithOptionalQueryParam("built_in", networking.BoolPtrToStrPtr(request.BuiltIn)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetGroupsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetActiveJobSpectraS3(ctx context.Context, request *models.GetActiveJobSpectraS3Request) (*models.GetActiveJobSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/active_job/"+request.ActiveJobId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetActiveJobSpectraS3Response(response, client.Logger)
}

func (client *Client) GetActiveJobsSpectraS3(ctx context.Context, request *models.GetActiveJobsSpectraS3Request) (*models.GetActiveJobsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/active_job").
		WithOptionalQueryParam("aggregating", networking.BoolPtrToStrPtr(request.Aggregating)).
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("chunk_client_processing_order_guarantee", networking.InterfaceToStrPtr(request.ChunkClientProcessingOrderGuarantee)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
		WithOptionalQueryParam("rechunked", request.Rechunked).
		WithOptionalQueryParam("request_type", networking.InterfaceToStrPtr(request.RequestType)).
		WithOptionalQueryParam("truncated", networking.BoolPtrToStrPtr(request.Truncated)).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetActiveJobsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetCanceledJobSpectraS3(ctx context.Context, request *models.GetCanceledJobSpectraS3Request) (*models.GetCanceledJobSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/canceled_job/"+request.CanceledJob).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetCanceledJobSpectraS3Response(response, client.Logger)
}

func (client *Client) GetCanceledJobsSpectraS3(ctx context.Context, request *models.GetCanceledJobsSpectraS3Request) (*models.GetCanceledJobsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/canceled_job").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("canceled_due_to_timeout", networking.BoolPtrToStrPtr(request.CanceledDueToTimeout)).
		WithOptionalQueryParam("chunk_client_processing_order_guarantee", networking.InterfaceToStrPtr(request.ChunkClientProcessingOrderGuarantee)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
		WithOptionalQueryParam("rechunked", request.Rechunked).
		WithOptionalQueryParam("request_type", networking.InterfaceToStrPtr(request.RequestType)).
		WithOptionalQueryParam("truncated", networking.BoolPtrToStrPtr(request.Truncated)).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetCanceledJobsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetCompletedJobSpectraS3(ctx context.Context, request *models.GetCompletedJobSpectraS3Request) (*models.GetCompletedJobSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/completed_job/"+request.CompletedJob).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetCompletedJobSpectraS3Response(response, client.Logger)
}

func (client *Client) GetCompletedJobsSpectraS3(ctx context.Context, request *models.GetCompletedJobsSpectraS3Request) (*models.GetCompletedJobsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/completed_job").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("chunk_client_processing_order_guarantee", networking.InterfaceToStrPtr(request.ChunkClientProcessingOrderGuarantee)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("priority", networking.InterfaceToStrPtr(request.Priority)).
		WithOptionalQueryParam("rechunked", request.Rechunked).
		WithOptionalQueryParam("request_type", networking.InterfaceToStrPtr(request.RequestType)).
		WithOptionalQueryParam("truncated", networking.BoolPtrToStrPtr(request.Truncated)).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetCompletedJobsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobChunkDaoSpectraS3(ctx context.Context, request *models.GetJobChunkDaoSpectraS3Request) (*models.GetJobChunkDaoSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_chunk_dao/"+request.JobChunkDao).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobChunkDaoSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobChunkSpectraS3(ctx context.Context, request *models.GetJobChunkSpectraS3Request) (*models.GetJobChunkSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_chunk/"+request.JobChunkId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobChunkSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobChunksReadyForClientProcessingSpectraS3(ctx context.Context, request *models.GetJobChunksReadyForClientProcessingSpectraS3Request) (*models.GetJobChunksReadyForClientProcessingSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_chunk").
		WithQueryParam("job", request.Job).
		WithOptionalQueryParam("job_chunk", request.JobChunk).
		WithOptionalQueryParam("preferred_number_of_chunks", networking.IntPtrToStrPtr(request.PreferredNumberOfChunks)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobChunksReadyForClientProcessingSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobCreationFailuresSpectraS3(ctx context.Context, request *models.GetJobCreationFailuresSpectraS3Request) (*models.GetJobCreationFailuresSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_creation_failed").
		WithOptionalQueryParam("date", request.Date).
		WithOptionalQueryParam("error_message", request.ErrorMessage).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_name", request.UserName).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobCreationFailuresSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobEntriesSpectraS3(ctx context.Context, request *models.GetJobEntriesSpectraS3Request) (*models.GetJobEntriesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_chunk_dao").
		WithQueryParam("job_id", request.JobId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobEntriesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobSpectraS3(ctx context.Context, request *models.GetJobSpectraS3Request) (*models.GetJobSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job/"+request.JobId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobSummarySpectraS3(ctx context.Context, request *models.GetJobSummarySpectraS3Request) (*models.GetJobSummarySpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job/"+request.JobId).
		WithQueryParam("summary", "").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobSummarySpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobToReplicateSpectraS3(ctx context.Context, request *models.GetJobToReplicateSpectraS3Request) (*models.GetJobToReplicateSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job/"+request.JobId).
		WithQueryParam("replicate", "").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobToReplicateSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobsSpectraS3(ctx context.Context, request *models.GetJobsSpectraS3Request) (*models.GetJobsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalVoidQueryParam("full_details", request.FullDetails).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetNodeSpectraS3(ctx context.Context, request *models.GetNodeSpectraS3Request) (*models.GetNodeSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/node/"+request.Node).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetNodeSpectraS3Response(response, client.Logger)
}

func (client *Client) GetNodesSpectraS3(ctx context.Context, request *models.GetNodesSpectraS3Request) (*models.GetNodesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/node").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetNodesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAzureTargetFailureNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetAzureTargetFailureNotificationRegistrationSpectraS3Request) (*models.GetAzureTargetFailureNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_target_failure_notification_registration/"+request.AzureTargetFailureNotificationRegistration).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAzureTargetFailureNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAzureTargetFailureNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) (*models.GetAzureTargetFailureNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_target_failure_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAzureTargetFailureNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBucketChangesNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetBucketChangesNotificationRegistrationSpectraS3Request) (*models.GetBucketChangesNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/bucket_changes_notification_registration/"+request.BucketChangesNotificationRegistration).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBucketChangesNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBucketChangesNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetBucketChangesNotificationRegistrationsSpectraS3Request) (*models.GetBucketChangesNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/bucket_changes_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBucketChangesNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBucketHistorySpectraS3(ctx context.Context, request *models.GetBucketHistorySpectraS3Request) (*models.GetBucketHistorySpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/bucket_history").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("min_sequence_number", networking.Int64PtrToStrPtr(request.MinSequenceNumber)).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBucketHistorySpectraS3Response(response, client.Logger)
}

func (client *Client) GetDs3TargetFailureNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetDs3TargetFailureNotificationRegistrationSpectraS3Request) (*models.GetDs3TargetFailureNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_target_failure_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDs3TargetFailureNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDs3TargetFailureNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) (*models.GetDs3TargetFailureNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_target_failure_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDs3TargetFailureNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobCompletedNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetJobCompletedNotificationRegistrationSpectraS3Request) (*models.GetJobCompletedNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_completed_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobCompletedNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobCompletedNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetJobCompletedNotificationRegistrationsSpectraS3Request) (*models.GetJobCompletedNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_completed_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobCompletedNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobCreatedNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetJobCreatedNotificationRegistrationSpectraS3Request) (*models.GetJobCreatedNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_created_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobCreatedNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobCreatedNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetJobCreatedNotificationRegistrationsSpectraS3Request) (*models.GetJobCreatedNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_created_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobCreatedNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobCreationFailedNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetJobCreationFailedNotificationRegistrationSpectraS3Request) (*models.GetJobCreationFailedNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_creation_failed_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobCreationFailedNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetJobCreationFailedNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetJobCreationFailedNotificationRegistrationsSpectraS3Request) (*models.GetJobCreationFailedNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/job_creation_failed_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetJobCreationFailedNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetObjectCachedNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetObjectCachedNotificationRegistrationSpectraS3Request) (*models.GetObjectCachedNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/object_cached_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetObjectCachedNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetObjectCachedNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetObjectCachedNotificationRegistrationsSpectraS3Request) (*models.GetObjectCachedNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/object_cached_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetObjectCachedNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetObjectLostNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetObjectLostNotificationRegistrationSpectraS3Request) (*models.GetObjectLostNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/object_lost_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetObjectLostNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetObjectLostNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetObjectLostNotificationRegistrationsSpectraS3Request) (*models.GetObjectLostNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/object_lost_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetObjectLostNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetObjectPersistedNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetObjectPersistedNotificationRegistrationSpectraS3Request) (*models.GetObjectPersistedNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/object_persisted_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetObjectPersistedNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetObjectPersistedNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetObjectPersistedNotificationRegistrationsSpectraS3Request) (*models.GetObjectPersistedNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/object_persisted_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetObjectPersistedNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetPoolFailureNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetPoolFailureNotificationRegistrationSpectraS3Request) (*models.GetPoolFailureNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/pool_failure_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetPoolFailureNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetPoolFailureNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetPoolFailureNotificationRegistrationsSpectraS3Request) (*models.GetPoolFailureNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/pool_failure_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetPoolFailureNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetS3TargetFailureNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetS3TargetFailureNotificationRegistrationSpectraS3Request) (*models.GetS3TargetFailureNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_target_failure_notification_registration/"+request.S3TargetFailureNotificationRegistration).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetS3TargetFailureNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetS3TargetFailureNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetS3TargetFailureNotificationRegistrationsSpectraS3Request) (*models.GetS3TargetFailureNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_target_failure_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetS3TargetFailureNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetStorageDomainFailureNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetStorageDomainFailureNotificationRegistrationSpectraS3Request) (*models.GetStorageDomainFailureNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/storage_domain_failure_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetStorageDomainFailureNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetStorageDomainFailureNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) (*models.GetStorageDomainFailureNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/storage_domain_failure_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetStorageDomainFailureNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSystemFailureNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetSystemFailureNotificationRegistrationSpectraS3Request) (*models.GetSystemFailureNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/system_failure_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSystemFailureNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSystemFailureNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetSystemFailureNotificationRegistrationsSpectraS3Request) (*models.GetSystemFailureNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/system_failure_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSystemFailureNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapeFailureNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetTapeFailureNotificationRegistrationSpectraS3Request) (*models.GetTapeFailureNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_failure_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapeFailureNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapeFailureNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetTapeFailureNotificationRegistrationsSpectraS3Request) (*models.GetTapeFailureNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_failure_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapeFailureNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapePartitionFailureNotificationRegistrationSpectraS3(ctx context.Context, request *models.GetTapePartitionFailureNotificationRegistrationSpectraS3Request) (*models.GetTapePartitionFailureNotificationRegistrationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_partition_failure_notification_registration/"+request.NotificationId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapePartitionFailureNotificationRegistrationSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapePartitionFailureNotificationRegistrationsSpectraS3(ctx context.Context, request *models.GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) (*models.GetTapePartitionFailureNotificationRegistrationsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_partition_failure_notification_registration").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("user_id", request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapePartitionFailureNotificationRegistrationsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBlobPersistenceSpectraS3(ctx context.Context, request *models.GetBlobPersistenceSpectraS3Request) (*models.GetBlobPersistenceSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/blob_persistence").
		WithReadCloser(buildStreamFromString(request.RequestPayload)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBlobPersistenceSpectraS3Response(response, client.Logger)
}

func (client *Client) GetObjectDetailsSpectraS3(ctx context.Context, request *models.GetObjectDetailsSpectraS3Request) (*models.GetObjectDetailsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/object/"+request.ObjectName).
		WithQueryParam("bucket_id", request.BucketId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetObjectDetailsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetObjectsDetailsSpectraS3(ctx context.Context, request *models.GetObjectsDetailsSpectraS3Request) (*models.GetObjectsDetailsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/object").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("end_date", networking.Int64PtrToStrPtr(request.EndDate)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("latest", networking.BoolPtrToStrPtr(request.Latest)).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("start_date", networking.Int64PtrToStrPtr(request.StartDate)).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.S3ObjectType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetObjectsDetailsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetObjectsWithFullDetailsSpectraS3(ctx context.Context, request *models.GetObjectsWithFullDetailsSpectraS3Request) (*models.GetObjectsWithFullDetailsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/object").
		WithQueryParam("full_details", "").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("end_date", networking.Int64PtrToStrPtr(request.EndDate)).
		WithOptionalVoidQueryParam("include_physical_placement", request.IncludePhysicalPlacement).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("latest", networking.BoolPtrToStrPtr(request.Latest)).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("start_date", networking.Int64PtrToStrPtr(request.StartDate)).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.S3ObjectType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetObjectsWithFullDetailsSpectraS3Response(response, client.Logger)
}

func (client *Client) VerifyPhysicalPlacementForObjectsSpectraS3(ctx context.Context, request *models.VerifyPhysicalPlacementForObjectsSpectraS3Request) (*models.VerifyPhysicalPlacementForObjectsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/bucket/"+request.BucketName).
		WithOptionalQueryParam("storage_domain", request.StorageDomain).
		WithQueryParam("operation", "verify_physical_placement").
		WithReadCloser(buildDs3GetObjectListStream(request.Objects)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewVerifyPhysicalPlacementForObjectsSpectraS3Response(response, client.Logger)
}

func (client *Client) VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3(ctx context.Context, request *models.VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) (*models.VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/bucket/"+request.BucketName).
		WithQueryParam("full_details", "").
		WithOptionalQueryParam("storage_domain", request.StorageDomain).
		WithQueryParam("operation", "verify_physical_placement").
		WithReadCloser(buildDs3GetObjectListStream(request.Objects)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewVerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBlobsOnPoolSpectraS3(ctx context.Context, request *models.GetBlobsOnPoolSpectraS3Request) (*models.GetBlobsOnPoolSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/pool/"+request.Pool).
		WithQueryParam("operation", "get_physical_placement").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBlobsOnPoolSpectraS3Response(response, client.Logger)
}

func (client *Client) GetPoolFailuresSpectraS3(ctx context.Context, request *models.GetPoolFailuresSpectraS3Request) (*models.GetPoolFailuresSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/pool_failure").
		WithOptionalQueryParam("error_message", request.ErrorMessage).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("pool_id", request.PoolId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.PoolFailureType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetPoolFailuresSpectraS3Response(response, client.Logger)
}

func (client *Client) GetPoolPartitionSpectraS3(ctx context.Context, request *models.GetPoolPartitionSpectraS3Request) (*models.GetPoolPartitionSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/pool_partition/"+request.PoolPartition).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetPoolPartitionSpectraS3Response(response, client.Logger)
}

func (client *Client) GetPoolPartitionsSpectraS3(ctx context.Context, request *models.GetPoolPartitionsSpectraS3Request) (*models.GetPoolPartitionsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/pool_partition").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.PoolType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetPoolPartitionsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetPoolSpectraS3(ctx context.Context, request *models.GetPoolSpectraS3Request) (*models.GetPoolSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/pool/"+request.Pool).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetPoolSpectraS3Response(response, client.Logger)
}

func (client *Client) GetPoolsSpectraS3(ctx context.Context, request *models.GetPoolsSpectraS3Request) (*models.GetPoolsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/pool").
		WithOptionalQueryParam("assigned_to_storage_domain", networking.BoolPtrToStrPtr(request.AssignedToStorageDomain)).
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("guid", request.Guid).
		WithOptionalQueryParam("health", networking.InterfaceToStrPtr(request.Health)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("last_verified", request.LastVerified).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("partition_id", request.PartitionId).
		WithOptionalQueryParam("powered_on", networking.BoolPtrToStrPtr(request.PoweredOn)).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("storage_domain_member_id", request.StorageDomainMemberId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.PoolType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetPoolsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetStorageDomainFailuresSpectraS3(ctx context.Context, request *models.GetStorageDomainFailuresSpectraS3Request) (*models.GetStorageDomainFailuresSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/storage_domain_failure").
		WithOptionalQueryParam("error_message", request.ErrorMessage).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("storage_domain_id", request.StorageDomainId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.StorageDomainFailureType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetStorageDomainFailuresSpectraS3Response(response, client.Logger)
}

func (client *Client) GetStorageDomainMemberSpectraS3(ctx context.Context, request *models.GetStorageDomainMemberSpectraS3Request) (*models.GetStorageDomainMemberSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/storage_domain_member/"+request.StorageDomainMember).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetStorageDomainMemberSpectraS3Response(response, client.Logger)
}

func (client *Client) GetStorageDomainMembersSpectraS3(ctx context.Context, request *models.GetStorageDomainMembersSpectraS3Request) (*models.GetStorageDomainMembersSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/storage_domain_member").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("pool_partition_id", request.PoolPartitionId).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("storage_domain_id", request.StorageDomainId).
		WithOptionalQueryParam("tape_partition_id", request.TapePartitionId).
		WithOptionalQueryParam("tape_type", request.TapeType).
		WithOptionalQueryParam("write_preference", networking.InterfaceToStrPtr(request.WritePreference)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetStorageDomainMembersSpectraS3Response(response, client.Logger)
}

func (client *Client) GetStorageDomainSpectraS3(ctx context.Context, request *models.GetStorageDomainSpectraS3Request) (*models.GetStorageDomainSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/storage_domain/"+request.StorageDomain).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetStorageDomainSpectraS3Response(response, client.Logger)
}

func (client *Client) GetStorageDomainsSpectraS3(ctx context.Context, request *models.GetStorageDomainsSpectraS3Request) (*models.GetStorageDomainsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/storage_domain").
		WithOptionalQueryParam("auto_eject_upon_cron", request.AutoEjectUponCron).
		WithOptionalQueryParam("auto_eject_upon_job_cancellation", networking.BoolPtrToStrPtr(request.AutoEjectUponJobCancellation)).
		WithOptionalQueryParam("auto_eject_upon_job_completion", networking.BoolPtrToStrPtr(request.AutoEjectUponJobCompletion)).
		WithOptionalQueryParam("auto_eject_upon_media_full", networking.BoolPtrToStrPtr(request.AutoEjectUponMediaFull)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("media_ejection_allowed", networking.BoolPtrToStrPtr(request.MediaEjectionAllowed)).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("secure_media_allocation", networking.BoolPtrToStrPtr(request.SecureMediaAllocation)).
		WithOptionalQueryParam("write_optimization", networking.InterfaceToStrPtr(request.WriteOptimization)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetStorageDomainsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAbmConfigSpectraS3(ctx context.Context, request *models.GetAbmConfigSpectraS3Request) (*models.GetAbmConfigSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/abm_config").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAbmConfigSpectraS3Response(response, client.Logger)
}

func (client *Client) GetFeatureKeysSpectraS3(ctx context.Context, request *models.GetFeatureKeysSpectraS3Request) (*models.GetFeatureKeysSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/feature_key").
		WithOptionalQueryParam("error_message", request.ErrorMessage).
		WithOptionalQueryParam("expiration_date", request.ExpirationDate).
		WithOptionalQueryParam("key", networking.InterfaceToStrPtr(request.Key)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetFeatureKeysSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSystemFailuresSpectraS3(ctx context.Context, request *models.GetSystemFailuresSpectraS3Request) (*models.GetSystemFailuresSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/system_failure").
		WithOptionalQueryParam("error_message", request.ErrorMessage).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.SystemFailureType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSystemFailuresSpectraS3Response(response, client.Logger)
}

func (client *Client) GetSystemInformationSpectraS3(ctx context.Context, request *models.GetSystemInformationSpectraS3Request) (*models.GetSystemInformationSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/system_information").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetSystemInformationSpectraS3Response(response, client.Logger)
}

func (client *Client) VerifySystemHealthSpectraS3(ctx context.Context, request *models.VerifySystemHealthSpectraS3Request) (*models.VerifySystemHealthSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/system_health").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewVerifySystemHealthSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBlobsOnTapeSpectraS3(ctx context.Context, request *models.GetBlobsOnTapeSpectraS3Request) (*models.GetBlobsOnTapeSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape/"+request.TapeId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithQueryParam("operation", "get_physical_placement").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBlobsOnTapeSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapeDensityDirectiveSpectraS3(ctx context.Context, request *models.GetTapeDensityDirectiveSpectraS3Request) (*models.GetTapeDensityDirectiveSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_density_directive/"+request.TapeDensityDirective).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapeDensityDirectiveSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapeDensityDirectivesSpectraS3(ctx context.Context, request *models.GetTapeDensityDirectivesSpectraS3Request) (*models.GetTapeDensityDirectivesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_density_directive").
		WithOptionalQueryParam("density", networking.InterfaceToStrPtr(request.Density)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("partition_id", request.PartitionId).
		WithOptionalQueryParam("tape_type", request.TapeType).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapeDensityDirectivesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapeDriveSpectraS3(ctx context.Context, request *models.GetTapeDriveSpectraS3Request) (*models.GetTapeDriveSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_drive/"+request.TapeDriveId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapeDriveSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapeDrivesSpectraS3(ctx context.Context, request *models.GetTapeDrivesSpectraS3Request) (*models.GetTapeDrivesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_drive").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("minimum_task_priority", networking.InterfaceToStrPtr(request.MinimumTaskPriority)).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("partition_id", request.PartitionId).
		WithOptionalQueryParam("reserved_task_type", networking.InterfaceToStrPtr(request.ReservedTaskType)).
		WithOptionalQueryParam("serial_number", request.SerialNumber).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.TapeDriveType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapeDrivesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapeFailuresSpectraS3(ctx context.Context, request *models.GetTapeFailuresSpectraS3Request) (*models.GetTapeFailuresSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_failure").
		WithOptionalQueryParam("error_message", request.ErrorMessage).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("tape_drive_id", request.TapeDriveId).
		WithOptionalQueryParam("tape_id", request.TapeId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.TapeFailureType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapeFailuresSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapeLibrariesSpectraS3(ctx context.Context, request *models.GetTapeLibrariesSpectraS3Request) (*models.GetTapeLibrariesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_library").
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("management_url", request.ManagementUrl).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("serial_number", request.SerialNumber).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapeLibrariesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapeLibrarySpectraS3(ctx context.Context, request *models.GetTapeLibrarySpectraS3Request) (*models.GetTapeLibrarySpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_library/"+request.TapeLibraryId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapeLibrarySpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapePartitionFailuresSpectraS3(ctx context.Context, request *models.GetTapePartitionFailuresSpectraS3Request) (*models.GetTapePartitionFailuresSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_partition_failure").
		WithOptionalQueryParam("error_message", request.ErrorMessage).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("partition_id", request.PartitionId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.TapePartitionFailureType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapePartitionFailuresSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapePartitionSpectraS3(ctx context.Context, request *models.GetTapePartitionSpectraS3Request) (*models.GetTapePartitionSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_partition/"+request.TapePartition).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapePartitionSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapePartitionWithFullDetailsSpectraS3(ctx context.Context, request *models.GetTapePartitionWithFullDetailsSpectraS3Request) (*models.GetTapePartitionWithFullDetailsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_partition/"+request.TapePartition).
		WithQueryParam("full_details", "").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapePartitionWithFullDetailsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapePartitionsSpectraS3(ctx context.Context, request *models.GetTapePartitionsSpectraS3Request) (*models.GetTapePartitionsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_partition").
		WithOptionalQueryParam("import_export_configuration", networking.InterfaceToStrPtr(request.ImportExportConfiguration)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("library_id", request.LibraryId).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
		WithOptionalQueryParam("serial_number", request.SerialNumber).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapePartitionsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapePartitionsWithFullDetailsSpectraS3(ctx context.Context, request *models.GetTapePartitionsWithFullDetailsSpectraS3Request) (*models.GetTapePartitionsWithFullDetailsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape_partition").
		WithQueryParam("full_details", "").
		WithOptionalQueryParam("import_export_configuration", networking.InterfaceToStrPtr(request.ImportExportConfiguration)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("library_id", request.LibraryId).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
		WithOptionalQueryParam("serial_number", request.SerialNumber).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapePartitionsWithFullDetailsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapeSpectraS3(ctx context.Context, request *models.GetTapeSpectraS3Request) (*models.GetTapeSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape/"+request.TapeId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapeSpectraS3Response(response, client.Logger)
}

func (client *Client) GetTapesSpectraS3(ctx context.Context, request *models.GetTapesSpectraS3Request) (*models.GetTapesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/tape").
		WithOptionalQueryParam("assigned_to_storage_domain", networking.BoolPtrToStrPtr(request.AssignedToStorageDomain)).
		WithOptionalQueryParam("available_raw_capacity", networking.Int64PtrToStrPtr(request.AvailableRawCapacity)).
		WithOptionalQueryParam("bar_code", request.BarCode).
		WithOptionalQueryParam("bucket", request.Bucket).
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalQueryParam("eject_label", request.EjectLabel).
		WithOptionalQueryParam("eject_location", request.EjectLocation).
		WithOptionalQueryParam("full_of_data", networking.BoolPtrToStrPtr(request.FullOfData)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("last_verified", request.LastVerified).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("partially_verified_end_of_tape", request.PartiallyVerifiedEndOfTape).
		WithOptionalQueryParam("partition", request.Partition).
		WithOptionalQueryParam("partition_id", request.PartitionId).
		WithOptionalQueryParam("previous_state", networking.InterfaceToStrPtr(request.PreviousState)).
		WithOptionalQueryParam("serial_number", request.SerialNumber).
		WithOptionalQueryParam("sort_by", request.SortBy).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		WithOptionalQueryParam("storage_domain", request.StorageDomain).
		WithOptionalQueryParam("storage_domain_member_id", request.StorageDomainMemberId).
		WithOptionalQueryParam("type", request.String).
		WithOptionalQueryParam("verify_pending", networking.InterfaceToStrPtr(request.VerifyPending)).
		WithOptionalQueryParam("write_protected", networking.BoolPtrToStrPtr(request.WriteProtected)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetTapesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAzureTargetBucketNamesSpectraS3(ctx context.Context, request *models.GetAzureTargetBucketNamesSpectraS3Request) (*models.GetAzureTargetBucketNamesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_target_bucket_name").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("target_id", request.TargetId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAzureTargetBucketNamesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAzureTargetFailuresSpectraS3(ctx context.Context, request *models.GetAzureTargetFailuresSpectraS3Request) (*models.GetAzureTargetFailuresSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_target_failure").
		WithOptionalQueryParam("error_message", request.ErrorMessage).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("target_id", request.TargetId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.TargetFailureType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAzureTargetFailuresSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAzureTargetReadPreferenceSpectraS3(ctx context.Context, request *models.GetAzureTargetReadPreferenceSpectraS3Request) (*models.GetAzureTargetReadPreferenceSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_target_read_preference/"+request.AzureTargetReadPreference).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAzureTargetReadPreferenceSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAzureTargetReadPreferencesSpectraS3(ctx context.Context, request *models.GetAzureTargetReadPreferencesSpectraS3Request) (*models.GetAzureTargetReadPreferencesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_target_read_preference").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("read_preference", networking.InterfaceToStrPtr(request.ReadPreference)).
		WithOptionalQueryParam("target_id", request.TargetId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAzureTargetReadPreferencesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAzureTargetSpectraS3(ctx context.Context, request *models.GetAzureTargetSpectraS3Request) (*models.GetAzureTargetSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_target/"+request.AzureTarget).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAzureTargetSpectraS3Response(response, client.Logger)
}

func (client *Client) GetAzureTargetsSpectraS3(ctx context.Context, request *models.GetAzureTargetsSpectraS3Request) (*models.GetAzureTargetsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_target").
		WithOptionalQueryParam("account_name", request.AccountName).
		WithOptionalQueryParam("default_read_preference", networking.InterfaceToStrPtr(request.DefaultReadPreference)).
		WithOptionalQueryParam("https", networking.BoolPtrToStrPtr(request.Https)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("permit_going_out_of_sync", networking.BoolPtrToStrPtr(request.PermitGoingOutOfSync)).
		WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetAzureTargetsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBlobsOnAzureTargetSpectraS3(ctx context.Context, request *models.GetBlobsOnAzureTargetSpectraS3Request) (*models.GetBlobsOnAzureTargetSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/azure_target/"+request.AzureTarget).
		WithQueryParam("operation", "get_physical_placement").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBlobsOnAzureTargetSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBlobsOnDs3TargetSpectraS3(ctx context.Context, request *models.GetBlobsOnDs3TargetSpectraS3Request) (*models.GetBlobsOnDs3TargetSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_target/"+request.Ds3Target).
		WithQueryParam("operation", "get_physical_placement").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBlobsOnDs3TargetSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDs3TargetDataPoliciesSpectraS3(ctx context.Context, request *models.GetDs3TargetDataPoliciesSpectraS3Request) (*models.GetDs3TargetDataPoliciesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_target_data_policies/"+request.Ds3TargetDataPolicies).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDs3TargetDataPoliciesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDs3TargetFailuresSpectraS3(ctx context.Context, request *models.GetDs3TargetFailuresSpectraS3Request) (*models.GetDs3TargetFailuresSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_target_failure").
		WithOptionalQueryParam("error_message", request.ErrorMessage).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("target_id", request.TargetId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.TargetFailureType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDs3TargetFailuresSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDs3TargetReadPreferenceSpectraS3(ctx context.Context, request *models.GetDs3TargetReadPreferenceSpectraS3Request) (*models.GetDs3TargetReadPreferenceSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_target_read_preference/"+request.Ds3TargetReadPreference).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDs3TargetReadPreferenceSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDs3TargetReadPreferencesSpectraS3(ctx context.Context, request *models.GetDs3TargetReadPreferencesSpectraS3Request) (*models.GetDs3TargetReadPreferencesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_target_read_preference").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("read_preference", networking.InterfaceToStrPtr(request.ReadPreference)).
		WithOptionalQueryParam("target_id", request.TargetId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDs3TargetReadPreferencesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDs3TargetSpectraS3(ctx context.Context, request *models.GetDs3TargetSpectraS3Request) (*models.GetDs3TargetSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_target/"+request.Ds3Target).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDs3TargetSpectraS3Response(response, client.Logger)
}

func (client *Client) GetDs3TargetsSpectraS3(ctx context.Context, request *models.GetDs3TargetsSpectraS3Request) (*models.GetDs3TargetsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/ds3_target").
		WithOptionalQueryParam("admin_auth_id", request.AdminAuthId).
		WithOptionalQueryParam("data_path_end_point", request.DataPathEndPoint).
		WithOptionalQueryParam("data_path_https", networking.BoolPtrToStrPtr(request.DataPathHttps)).
		WithOptionalQueryParam("data_path_port", networking.IntPtrToStrPtr(request.DataPathPort)).
		WithOptionalQueryParam("data_path_proxy", request.DataPathProxy).
		WithOptionalQueryParam("data_path_verify_certificate", networking.BoolPtrToStrPtr(request.DataPathVerifyCertificate)).
		WithOptionalQueryParam("default_read_preference", networking.InterfaceToStrPtr(request.DefaultReadPreference)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("permit_going_out_of_sync", networking.BoolPtrToStrPtr(request.PermitGoingOutOfSync)).
		WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetDs3TargetsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetBlobsOnS3TargetSpectraS3(ctx context.Context, request *models.GetBlobsOnS3TargetSpectraS3Request) (*models.GetBlobsOnS3TargetSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_target/"+request.S3Target).
		WithQueryParam("operation", "get_physical_placement").
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetBlobsOnS3TargetSpectraS3Response(response, client.Logger)
}

func (client *Client) GetS3TargetBucketNamesSpectraS3(ctx context.Context, request *models.GetS3TargetBucketNamesSpectraS3Request) (*models.GetS3TargetBucketNamesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_target_bucket_name").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("target_id", request.TargetId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetS3TargetBucketNamesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetS3TargetFailuresSpectraS3(ctx context.Context, request *models.GetS3TargetFailuresSpectraS3Request) (*models.GetS3TargetFailuresSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_target_failure").
		WithOptionalQueryParam("error_message", request.ErrorMessage).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("target_id", request.TargetId).
		WithOptionalQueryParam("type", networking.InterfaceToStrPtr(request.TargetFailureType)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetS3TargetFailuresSpectraS3Response(response, client.Logger)
}

func (client *Client) GetS3TargetReadPreferenceSpectraS3(ctx context.Context, request *models.GetS3TargetReadPreferenceSpectraS3Request) (*models.GetS3TargetReadPreferenceSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_target_read_preference/"+request.S3TargetReadPreference).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetS3TargetReadPreferenceSpectraS3Response(response, client.Logger)
}

func (client *Client) GetS3TargetReadPreferencesSpectraS3(ctx context.Context, request *models.GetS3TargetReadPreferencesSpectraS3Request) (*models.GetS3TargetReadPreferencesSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_target_read_preference").
		WithOptionalQueryParam("bucket_id", request.BucketId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("read_preference", networking.InterfaceToStrPtr(request.ReadPreference)).
		WithOptionalQueryParam("target_id", request.TargetId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetS3TargetReadPreferencesSpectraS3Response(response, client.Logger)
}

func (client *Client) GetS3TargetSpectraS3(ctx context.Context, request *models.GetS3TargetSpectraS3Request) (*models.GetS3TargetSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_target/"+request.S3Target).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetS3TargetSpectraS3Response(response, client.Logger)
}

func (client *Client) GetS3TargetsSpectraS3(ctx context.Context, request *models.GetS3TargetsSpectraS3Request) (*models.GetS3TargetsSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/s3_target").
		WithOptionalQueryParam("access_key", request.AccessKey).
		WithOptionalQueryParam("data_path_end_point", request.DataPathEndPoint).
		WithOptionalQueryParam("default_read_preference", networking.InterfaceToStrPtr(request.DefaultReadPreference)).
		WithOptionalQueryParam("https", networking.BoolPtrToStrPtr(request.Https)).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("naming_mode", networking.InterfaceToStrPtr(request.NamingMode)).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		WithOptionalQueryParam("permit_going_out_of_sync", networking.BoolPtrToStrPtr(request.PermitGoingOutOfSync)).
		WithOptionalQueryParam("quiesced", networking.InterfaceToStrPtr(request.Quiesced)).
		WithOptionalQueryParam("region", networking.InterfaceToStrPtr(request.Region)).
		WithOptionalQueryParam("state", networking.InterfaceToStrPtr(request.State)).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetS3TargetsSpectraS3Response(response, client.Logger)
}

func (client *Client) GetUserSpectraS3(ctx context.Context, request *models.GetUserSpectraS3Request) (*models.GetUserSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/user/"+request.UserId).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetUserSpectraS3Response(response, client.Logger)
}

func (client *Client) GetUsersSpectraS3(ctx context.Context, request *models.GetUsersSpectraS3Request) (*models.GetUsersSpectraS3Response, error) {
	// Build the http request
	httpRequest, err := networking.NewHttpRequestBuilder().
		WithHttpVerb(HTTP_VERB_GET).
		WithPath("/_rest_/user").
		WithOptionalQueryParam("auth_id", request.AuthId).
		WithOptionalQueryParam("default_data_policy_id", request.DefaultDataPolicyId).
		WithOptionalVoidQueryParam("last_page", request.LastPage).
		WithOptionalQueryParam("name", request.Name).
		WithOptionalQueryParam("page_length", networking.IntPtrToStrPtr(request.PageLength)).
		WithOptionalQueryParam("page_offset", networking.IntPtrToStrPtr(request.PageOffset)).
		WithOptionalQueryParam("page_start_marker", request.PageStartMarker).
		Build(ctx, client.connectionInfo)

	if err != nil {
		return nil, err
	}

	networkRetryDecorator := networking.NewNetworkRetryDecorator(client.sendNetwork, client.clientPolicy.maxRetries)
	httpRedirectDecorator := networking.NewHttpTempRedirectDecorator(networkRetryDecorator, client.clientPolicy.maxRedirect)

	// Invoke the HTTP request.
	response, requestErr := httpRedirectDecorator.Invoke(httpRequest)
	if requestErr != nil {
		return nil, requestErr
	}

	// Create a response object based on the result.
	return models.NewGetUsersSpectraS3Response(response, client.Logger)
}
