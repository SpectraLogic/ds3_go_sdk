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

type VerifyBulkJobSpectraS3Request struct {
    BucketName string
    Aggregating *bool
    Name *string
    Objects []Ds3GetObject
    Priority Priority
}

func NewVerifyBulkJobSpectraS3Request(bucketName string, objectNames []string) *VerifyBulkJobSpectraS3Request {

    return &VerifyBulkJobSpectraS3Request{
        BucketName: bucketName,
        Objects: buildDs3GetObjectSliceFromNames(objectNames),
    }
}

func NewVerifyBulkJobSpectraS3RequestWithPartialObjects(bucketName string, objects []Ds3GetObject) *VerifyBulkJobSpectraS3Request {

    return &VerifyBulkJobSpectraS3Request{
        BucketName: bucketName,
        Objects: objects,
    }
}

func (verifyBulkJobSpectraS3Request *VerifyBulkJobSpectraS3Request) WithAggregating(aggregating bool) *VerifyBulkJobSpectraS3Request {
    verifyBulkJobSpectraS3Request.Aggregating = &aggregating
    return verifyBulkJobSpectraS3Request
}

func (verifyBulkJobSpectraS3Request *VerifyBulkJobSpectraS3Request) WithName(name string) *VerifyBulkJobSpectraS3Request {
    verifyBulkJobSpectraS3Request.Name = &name
    return verifyBulkJobSpectraS3Request
}

func (verifyBulkJobSpectraS3Request *VerifyBulkJobSpectraS3Request) WithPriority(priority Priority) *VerifyBulkJobSpectraS3Request {
    verifyBulkJobSpectraS3Request.Priority = priority
    return verifyBulkJobSpectraS3Request
}

