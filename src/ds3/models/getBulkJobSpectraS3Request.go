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

package models

type GetBulkJobSpectraS3Request struct {
    BucketName string
    Aggregating *bool
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    ImplicitJobIdResolution *bool
    Name *string
    Objects []Ds3GetObject
    Priority Priority
}

func NewGetBulkJobSpectraS3Request(bucketName string, objectNames []string) *GetBulkJobSpectraS3Request {

    return &GetBulkJobSpectraS3Request{
        BucketName: bucketName,
        Objects: buildDs3GetObjectSliceFromNames(objectNames),
    }
}

func NewGetBulkJobSpectraS3RequestWithPartialObjects(bucketName string, objects []Ds3GetObject) *GetBulkJobSpectraS3Request {

    return &GetBulkJobSpectraS3Request{
        BucketName: bucketName,
        Objects: objects,
    }
}

func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) WithAggregating(aggregating bool) *GetBulkJobSpectraS3Request {
    getBulkJobSpectraS3Request.Aggregating = &aggregating
    return getBulkJobSpectraS3Request
}

func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetBulkJobSpectraS3Request {
    getBulkJobSpectraS3Request.ChunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    return getBulkJobSpectraS3Request
}

func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) WithImplicitJobIdResolution(implicitJobIdResolution bool) *GetBulkJobSpectraS3Request {
    getBulkJobSpectraS3Request.ImplicitJobIdResolution = &implicitJobIdResolution
    return getBulkJobSpectraS3Request
}

func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) WithName(name string) *GetBulkJobSpectraS3Request {
    getBulkJobSpectraS3Request.Name = &name
    return getBulkJobSpectraS3Request
}

func (getBulkJobSpectraS3Request *GetBulkJobSpectraS3Request) WithPriority(priority Priority) *GetBulkJobSpectraS3Request {
    getBulkJobSpectraS3Request.Priority = priority
    return getBulkJobSpectraS3Request
}

