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

func (client *Client) PutBucket(request *models.PutBucketRequest) (*models.PutBucketResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutBucketResponse(response)
}
func (client *Client) PutMultiPartUploadPart(request *models.PutMultiPartUploadPartRequest) (*models.PutMultiPartUploadPartResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutMultiPartUploadPartResponse(response)
}
func (client *Client) PutObject(request *models.PutObjectRequest) (*models.PutObjectResponse, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutObjectResponse(response)
}
func (client *Client) ModifyBucketSpectraS3(request *models.ModifyBucketSpectraS3Request) (*models.ModifyBucketSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyBucketSpectraS3Response(response)
}
func (client *Client) ForceFullCacheReclaimSpectraS3(request *models.ForceFullCacheReclaimSpectraS3Request) (*models.ForceFullCacheReclaimSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewForceFullCacheReclaimSpectraS3Response(response)
}
func (client *Client) ModifyCacheFilesystemSpectraS3(request *models.ModifyCacheFilesystemSpectraS3Request) (*models.ModifyCacheFilesystemSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyCacheFilesystemSpectraS3Response(response)
}
func (client *Client) ModifyDataPathBackendSpectraS3(request *models.ModifyDataPathBackendSpectraS3Request) (*models.ModifyDataPathBackendSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyDataPathBackendSpectraS3Response(response)
}
func (client *Client) ModifyAzureDataReplicationRuleSpectraS3(request *models.ModifyAzureDataReplicationRuleSpectraS3Request) (*models.ModifyAzureDataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyAzureDataReplicationRuleSpectraS3Response(response)
}
func (client *Client) ModifyDataPersistenceRuleSpectraS3(request *models.ModifyDataPersistenceRuleSpectraS3Request) (*models.ModifyDataPersistenceRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyDataPersistenceRuleSpectraS3Response(response)
}
func (client *Client) ModifyDataPolicySpectraS3(request *models.ModifyDataPolicySpectraS3Request) (*models.ModifyDataPolicySpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyDataPolicySpectraS3Response(response)
}
func (client *Client) ModifyDs3DataReplicationRuleSpectraS3(request *models.ModifyDs3DataReplicationRuleSpectraS3Request) (*models.ModifyDs3DataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyDs3DataReplicationRuleSpectraS3Response(response)
}
func (client *Client) ModifyS3DataReplicationRuleSpectraS3(request *models.ModifyS3DataReplicationRuleSpectraS3Request) (*models.ModifyS3DataReplicationRuleSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyS3DataReplicationRuleSpectraS3Response(response)
}
func (client *Client) MarkSuspectBlobAzureTargetsAsDegradedSpectraS3(request *models.MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Request) (*models.MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewMarkSuspectBlobAzureTargetsAsDegradedSpectraS3Response(response)
}
func (client *Client) MarkSuspectBlobDs3TargetsAsDegradedSpectraS3(request *models.MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request) (*models.MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewMarkSuspectBlobDs3TargetsAsDegradedSpectraS3Response(response)
}
func (client *Client) MarkSuspectBlobPoolsAsDegradedSpectraS3(request *models.MarkSuspectBlobPoolsAsDegradedSpectraS3Request) (*models.MarkSuspectBlobPoolsAsDegradedSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewMarkSuspectBlobPoolsAsDegradedSpectraS3Response(response)
}
func (client *Client) MarkSuspectBlobS3TargetsAsDegradedSpectraS3(request *models.MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request) (*models.MarkSuspectBlobS3TargetsAsDegradedSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewMarkSuspectBlobS3TargetsAsDegradedSpectraS3Response(response)
}
func (client *Client) MarkSuspectBlobTapesAsDegradedSpectraS3(request *models.MarkSuspectBlobTapesAsDegradedSpectraS3Request) (*models.MarkSuspectBlobTapesAsDegradedSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewMarkSuspectBlobTapesAsDegradedSpectraS3Response(response)
}
func (client *Client) ModifyGroupSpectraS3(request *models.ModifyGroupSpectraS3Request) (*models.ModifyGroupSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyGroupSpectraS3Response(response)
}
func (client *Client) VerifyUserIsMemberOfGroupSpectraS3(request *models.VerifyUserIsMemberOfGroupSpectraS3Request) (*models.VerifyUserIsMemberOfGroupSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyUserIsMemberOfGroupSpectraS3Response(response)
}
func (client *Client) AllocateJobChunkSpectraS3(request *models.AllocateJobChunkSpectraS3Request) (*models.AllocateJobChunkSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewAllocateJobChunkSpectraS3Response(response)
}
func (client *Client) GetBulkJobSpectraS3(request *models.GetBulkJobSpectraS3Request) (*models.GetBulkJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetBulkJobSpectraS3Response(response)
}
func (client *Client) PutBulkJobSpectraS3(request *models.PutBulkJobSpectraS3Request) (*models.PutBulkJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPutBulkJobSpectraS3Response(response)
}
func (client *Client) VerifyBulkJobSpectraS3(request *models.VerifyBulkJobSpectraS3Request) (*models.VerifyBulkJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyBulkJobSpectraS3Response(response)
}
func (client *Client) ModifyActiveJobSpectraS3(request *models.ModifyActiveJobSpectraS3Request) (*models.ModifyActiveJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyActiveJobSpectraS3Response(response)
}
func (client *Client) ModifyJobSpectraS3(request *models.ModifyJobSpectraS3Request) (*models.ModifyJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyJobSpectraS3Response(response)
}
func (client *Client) ReplicatePutJobSpectraS3(request *models.ReplicatePutJobSpectraS3Request) (*models.ReplicatePutJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewReplicatePutJobSpectraS3Response(response)
}
func (client *Client) VerifySafeToCreatePutJobSpectraS3(request *models.VerifySafeToCreatePutJobSpectraS3Request) (*models.VerifySafeToCreatePutJobSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifySafeToCreatePutJobSpectraS3Response(response)
}
func (client *Client) ModifyNodeSpectraS3(request *models.ModifyNodeSpectraS3Request) (*models.ModifyNodeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyNodeSpectraS3Response(response)
}
func (client *Client) GetPhysicalPlacementForObjectsSpectraS3(request *models.GetPhysicalPlacementForObjectsSpectraS3Request) (*models.GetPhysicalPlacementForObjectsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetPhysicalPlacementForObjectsSpectraS3Response(response)
}
func (client *Client) GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3(request *models.GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) (*models.GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewGetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response(response)
}
func (client *Client) CancelImportOnAllPoolsSpectraS3(request *models.CancelImportOnAllPoolsSpectraS3Request) (*models.CancelImportOnAllPoolsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelImportOnAllPoolsSpectraS3Response(response)
}
func (client *Client) CancelImportPoolSpectraS3(request *models.CancelImportPoolSpectraS3Request) (*models.CancelImportPoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelImportPoolSpectraS3Response(response)
}
func (client *Client) CancelVerifyOnAllPoolsSpectraS3(request *models.CancelVerifyOnAllPoolsSpectraS3Request) (*models.CancelVerifyOnAllPoolsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelVerifyOnAllPoolsSpectraS3Response(response)
}
func (client *Client) CancelVerifyPoolSpectraS3(request *models.CancelVerifyPoolSpectraS3Request) (*models.CancelVerifyPoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelVerifyPoolSpectraS3Response(response)
}
func (client *Client) CompactAllPoolsSpectraS3(request *models.CompactAllPoolsSpectraS3Request) (*models.CompactAllPoolsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCompactAllPoolsSpectraS3Response(response)
}
func (client *Client) CompactPoolSpectraS3(request *models.CompactPoolSpectraS3Request) (*models.CompactPoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCompactPoolSpectraS3Response(response)
}
func (client *Client) DeallocatePoolSpectraS3(request *models.DeallocatePoolSpectraS3Request) (*models.DeallocatePoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewDeallocatePoolSpectraS3Response(response)
}
func (client *Client) ForcePoolEnvironmentRefreshSpectraS3(request *models.ForcePoolEnvironmentRefreshSpectraS3Request) (*models.ForcePoolEnvironmentRefreshSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewForcePoolEnvironmentRefreshSpectraS3Response(response)
}
func (client *Client) FormatAllForeignPoolsSpectraS3(request *models.FormatAllForeignPoolsSpectraS3Request) (*models.FormatAllForeignPoolsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewFormatAllForeignPoolsSpectraS3Response(response)
}
func (client *Client) FormatForeignPoolSpectraS3(request *models.FormatForeignPoolSpectraS3Request) (*models.FormatForeignPoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewFormatForeignPoolSpectraS3Response(response)
}
func (client *Client) ImportAllPoolsSpectraS3(request *models.ImportAllPoolsSpectraS3Request) (*models.ImportAllPoolsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewImportAllPoolsSpectraS3Response(response)
}
func (client *Client) ImportPoolSpectraS3(request *models.ImportPoolSpectraS3Request) (*models.ImportPoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewImportPoolSpectraS3Response(response)
}
func (client *Client) ModifyAllPoolsSpectraS3(request *models.ModifyAllPoolsSpectraS3Request) (*models.ModifyAllPoolsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyAllPoolsSpectraS3Response(response)
}
func (client *Client) ModifyPoolPartitionSpectraS3(request *models.ModifyPoolPartitionSpectraS3Request) (*models.ModifyPoolPartitionSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyPoolPartitionSpectraS3Response(response)
}
func (client *Client) ModifyPoolSpectraS3(request *models.ModifyPoolSpectraS3Request) (*models.ModifyPoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyPoolSpectraS3Response(response)
}
func (client *Client) VerifyAllPoolsSpectraS3(request *models.VerifyAllPoolsSpectraS3Request) (*models.VerifyAllPoolsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyAllPoolsSpectraS3Response(response)
}
func (client *Client) VerifyPoolSpectraS3(request *models.VerifyPoolSpectraS3Request) (*models.VerifyPoolSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyPoolSpectraS3Response(response)
}
func (client *Client) ConvertStorageDomainToDs3TargetSpectraS3(request *models.ConvertStorageDomainToDs3TargetSpectraS3Request) (*models.ConvertStorageDomainToDs3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewConvertStorageDomainToDs3TargetSpectraS3Response(response)
}
func (client *Client) ModifyStorageDomainMemberSpectraS3(request *models.ModifyStorageDomainMemberSpectraS3Request) (*models.ModifyStorageDomainMemberSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyStorageDomainMemberSpectraS3Response(response)
}
func (client *Client) ModifyStorageDomainSpectraS3(request *models.ModifyStorageDomainSpectraS3Request) (*models.ModifyStorageDomainSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyStorageDomainSpectraS3Response(response)
}
func (client *Client) ForceFeatureKeyValidationSpectraS3(request *models.ForceFeatureKeyValidationSpectraS3Request) (*models.ForceFeatureKeyValidationSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewForceFeatureKeyValidationSpectraS3Response(response)
}
func (client *Client) ResetInstanceIdentifierSpectraS3(request *models.ResetInstanceIdentifierSpectraS3Request) (*models.ResetInstanceIdentifierSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewResetInstanceIdentifierSpectraS3Response(response)
}
func (client *Client) CancelEjectOnAllTapesSpectraS3(request *models.CancelEjectOnAllTapesSpectraS3Request) (*models.CancelEjectOnAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelEjectOnAllTapesSpectraS3Response(response)
}
func (client *Client) CancelEjectTapeSpectraS3(request *models.CancelEjectTapeSpectraS3Request) (*models.CancelEjectTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelEjectTapeSpectraS3Response(response)
}
func (client *Client) CancelFormatOnAllTapesSpectraS3(request *models.CancelFormatOnAllTapesSpectraS3Request) (*models.CancelFormatOnAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelFormatOnAllTapesSpectraS3Response(response)
}
func (client *Client) CancelFormatTapeSpectraS3(request *models.CancelFormatTapeSpectraS3Request) (*models.CancelFormatTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelFormatTapeSpectraS3Response(response)
}
func (client *Client) CancelImportOnAllTapesSpectraS3(request *models.CancelImportOnAllTapesSpectraS3Request) (*models.CancelImportOnAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelImportOnAllTapesSpectraS3Response(response)
}
func (client *Client) CancelImportTapeSpectraS3(request *models.CancelImportTapeSpectraS3Request) (*models.CancelImportTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelImportTapeSpectraS3Response(response)
}
func (client *Client) CancelOnlineOnAllTapesSpectraS3(request *models.CancelOnlineOnAllTapesSpectraS3Request) (*models.CancelOnlineOnAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelOnlineOnAllTapesSpectraS3Response(response)
}
func (client *Client) CancelOnlineTapeSpectraS3(request *models.CancelOnlineTapeSpectraS3Request) (*models.CancelOnlineTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelOnlineTapeSpectraS3Response(response)
}
func (client *Client) CancelVerifyOnAllTapesSpectraS3(request *models.CancelVerifyOnAllTapesSpectraS3Request) (*models.CancelVerifyOnAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelVerifyOnAllTapesSpectraS3Response(response)
}
func (client *Client) CancelVerifyTapeSpectraS3(request *models.CancelVerifyTapeSpectraS3Request) (*models.CancelVerifyTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCancelVerifyTapeSpectraS3Response(response)
}
func (client *Client) CleanTapeDriveSpectraS3(request *models.CleanTapeDriveSpectraS3Request) (*models.CleanTapeDriveSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewCleanTapeDriveSpectraS3Response(response)
}
func (client *Client) EjectAllTapesSpectraS3(request *models.EjectAllTapesSpectraS3Request) (*models.EjectAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewEjectAllTapesSpectraS3Response(response)
}
func (client *Client) EjectStorageDomainBlobsSpectraS3(request *models.EjectStorageDomainBlobsSpectraS3Request) (*models.EjectStorageDomainBlobsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewEjectStorageDomainBlobsSpectraS3Response(response)
}
func (client *Client) EjectStorageDomainSpectraS3(request *models.EjectStorageDomainSpectraS3Request) (*models.EjectStorageDomainSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewEjectStorageDomainSpectraS3Response(response)
}
func (client *Client) EjectTapeSpectraS3(request *models.EjectTapeSpectraS3Request) (*models.EjectTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewEjectTapeSpectraS3Response(response)
}
func (client *Client) ForceTapeEnvironmentRefreshSpectraS3(request *models.ForceTapeEnvironmentRefreshSpectraS3Request) (*models.ForceTapeEnvironmentRefreshSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewForceTapeEnvironmentRefreshSpectraS3Response(response)
}
func (client *Client) FormatAllTapesSpectraS3(request *models.FormatAllTapesSpectraS3Request) (*models.FormatAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewFormatAllTapesSpectraS3Response(response)
}
func (client *Client) FormatTapeSpectraS3(request *models.FormatTapeSpectraS3Request) (*models.FormatTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewFormatTapeSpectraS3Response(response)
}
func (client *Client) ImportAllTapesSpectraS3(request *models.ImportAllTapesSpectraS3Request) (*models.ImportAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewImportAllTapesSpectraS3Response(response)
}
func (client *Client) ImportTapeSpectraS3(request *models.ImportTapeSpectraS3Request) (*models.ImportTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewImportTapeSpectraS3Response(response)
}
func (client *Client) InspectAllTapesSpectraS3(request *models.InspectAllTapesSpectraS3Request) (*models.InspectAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewInspectAllTapesSpectraS3Response(response)
}
func (client *Client) InspectTapeSpectraS3(request *models.InspectTapeSpectraS3Request) (*models.InspectTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewInspectTapeSpectraS3Response(response)
}
func (client *Client) ModifyAllTapePartitionsSpectraS3(request *models.ModifyAllTapePartitionsSpectraS3Request) (*models.ModifyAllTapePartitionsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyAllTapePartitionsSpectraS3Response(response)
}
func (client *Client) ModifyTapePartitionSpectraS3(request *models.ModifyTapePartitionSpectraS3Request) (*models.ModifyTapePartitionSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyTapePartitionSpectraS3Response(response)
}
func (client *Client) ModifyTapeSpectraS3(request *models.ModifyTapeSpectraS3Request) (*models.ModifyTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyTapeSpectraS3Response(response)
}
func (client *Client) OnlineAllTapesSpectraS3(request *models.OnlineAllTapesSpectraS3Request) (*models.OnlineAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewOnlineAllTapesSpectraS3Response(response)
}
func (client *Client) OnlineTapeSpectraS3(request *models.OnlineTapeSpectraS3Request) (*models.OnlineTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewOnlineTapeSpectraS3Response(response)
}
func (client *Client) RawImportAllTapesSpectraS3(request *models.RawImportAllTapesSpectraS3Request) (*models.RawImportAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewRawImportAllTapesSpectraS3Response(response)
}
func (client *Client) RawImportTapeSpectraS3(request *models.RawImportTapeSpectraS3Request) (*models.RawImportTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewRawImportTapeSpectraS3Response(response)
}
func (client *Client) VerifyAllTapesSpectraS3(request *models.VerifyAllTapesSpectraS3Request) (*models.VerifyAllTapesSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyAllTapesSpectraS3Response(response)
}
func (client *Client) VerifyTapeSpectraS3(request *models.VerifyTapeSpectraS3Request) (*models.VerifyTapeSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyTapeSpectraS3Response(response)
}
func (client *Client) ForceTargetEnvironmentRefreshSpectraS3(request *models.ForceTargetEnvironmentRefreshSpectraS3Request) (*models.ForceTargetEnvironmentRefreshSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewForceTargetEnvironmentRefreshSpectraS3Response(response)
}
func (client *Client) ImportAzureTargetSpectraS3(request *models.ImportAzureTargetSpectraS3Request) (*models.ImportAzureTargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewImportAzureTargetSpectraS3Response(response)
}
func (client *Client) ModifyAllAzureTargetsSpectraS3(request *models.ModifyAllAzureTargetsSpectraS3Request) (*models.ModifyAllAzureTargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyAllAzureTargetsSpectraS3Response(response)
}
func (client *Client) ModifyAzureTargetSpectraS3(request *models.ModifyAzureTargetSpectraS3Request) (*models.ModifyAzureTargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyAzureTargetSpectraS3Response(response)
}
func (client *Client) VerifyAzureTargetSpectraS3(request *models.VerifyAzureTargetSpectraS3Request) (*models.VerifyAzureTargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyAzureTargetSpectraS3Response(response)
}
func (client *Client) ModifyAllDs3TargetsSpectraS3(request *models.ModifyAllDs3TargetsSpectraS3Request) (*models.ModifyAllDs3TargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyAllDs3TargetsSpectraS3Response(response)
}
func (client *Client) ModifyDs3TargetSpectraS3(request *models.ModifyDs3TargetSpectraS3Request) (*models.ModifyDs3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyDs3TargetSpectraS3Response(response)
}
func (client *Client) PairBackRegisteredDs3TargetSpectraS3(request *models.PairBackRegisteredDs3TargetSpectraS3Request) (*models.PairBackRegisteredDs3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewPairBackRegisteredDs3TargetSpectraS3Response(response)
}
func (client *Client) VerifyDs3TargetSpectraS3(request *models.VerifyDs3TargetSpectraS3Request) (*models.VerifyDs3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyDs3TargetSpectraS3Response(response)
}
func (client *Client) ImportS3TargetSpectraS3(request *models.ImportS3TargetSpectraS3Request) (*models.ImportS3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewImportS3TargetSpectraS3Response(response)
}
func (client *Client) ModifyAllS3TargetsSpectraS3(request *models.ModifyAllS3TargetsSpectraS3Request) (*models.ModifyAllS3TargetsSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyAllS3TargetsSpectraS3Response(response)
}
func (client *Client) ModifyS3TargetSpectraS3(request *models.ModifyS3TargetSpectraS3Request) (*models.ModifyS3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyS3TargetSpectraS3Response(response)
}
func (client *Client) VerifyS3TargetSpectraS3(request *models.VerifyS3TargetSpectraS3Request) (*models.VerifyS3TargetSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewVerifyS3TargetSpectraS3Response(response)
}
func (client *Client) ModifyUserSpectraS3(request *models.ModifyUserSpectraS3Request) (*models.ModifyUserSpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewModifyUserSpectraS3Response(response)
}
func (client *Client) RegenerateUserSecretKeySpectraS3(request *models.RegenerateUserSecretKeySpectraS3Request) (*models.RegenerateUserSecretKeySpectraS3Response, error) {
    networkRetryDecorator := networking.NewNetworkRetryDecorator(&(client.netLayer), client.clientPolicy.maxRetries)

    // Invoke the HTTP request.
    response, requestErr := networkRetryDecorator.Invoke(request)
    if requestErr != nil {
        return nil, requestErr
    }

    // Create a response object based on the result.
    return models.NewRegenerateUserSecretKeySpectraS3Response(response)
}

