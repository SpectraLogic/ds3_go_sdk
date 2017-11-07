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

type PutBulkJobSpectraS3Request struct {
    BucketName                  string
    Aggregating                 *bool
    ImplicitJobIdResolution     *bool
    MaxUploadSize               *int64
    MinimizeSpanningAcrossMedia *bool
    Name                        *string
    Priority                    Priority
    VerifyAfterWrite            *bool
    Objects                     []Ds3PutObject
    IgnoreNamingConflicts       bool
    Force                       bool
}

func NewPutBulkJobSpectraS3Request(bucketName string, objects []Ds3PutObject) *PutBulkJobSpectraS3Request {
    return &PutBulkJobSpectraS3Request{
        BucketName:  bucketName,
        Objects:     objects,
    }
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithAggregating(aggregating bool) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.Aggregating = &aggregating
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithImplicitJobIdResolution(implicitJobIdResolution bool) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.ImplicitJobIdResolution = &implicitJobIdResolution
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithMaxUploadSize(maxUploadSize int64) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.MaxUploadSize = &maxUploadSize
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithMinimizeSpanningAcrossMedia(minimizeSpanningAcrossMedia bool) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.MinimizeSpanningAcrossMedia = &minimizeSpanningAcrossMedia
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithPriority(priority Priority) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.Priority = priority
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithVerifyAfterWrite(verifyAfterWrite bool) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.VerifyAfterWrite = &verifyAfterWrite
    return putBulkJobSpectraS3Request
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithName(name string) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.Name = &name
    return putBulkJobSpectraS3Request
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithForce() *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.Force = true
    return putBulkJobSpectraS3Request
}
func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithIgnoreNamingConflicts() *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.IgnoreNamingConflicts = true
    return putBulkJobSpectraS3Request
}
