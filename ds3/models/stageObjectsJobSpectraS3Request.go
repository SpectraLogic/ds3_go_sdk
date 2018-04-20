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

type StageObjectsJobSpectraS3Request struct {
    BucketName string
    Name *string
    Objects []Ds3GetObject
    Priority Priority
}

func NewStageObjectsJobSpectraS3Request(bucketName string, objectNames []string) *StageObjectsJobSpectraS3Request {

    return &StageObjectsJobSpectraS3Request{
        BucketName: bucketName,
        Objects: buildDs3GetObjectSliceFromNames(objectNames),
    }
}

func NewStageObjectsJobSpectraS3RequestWithPartialObjects(bucketName string, objects []Ds3GetObject) *StageObjectsJobSpectraS3Request {

    return &StageObjectsJobSpectraS3Request{
        BucketName: bucketName,
        Objects: objects,
    }
}

func (stageObjectsJobSpectraS3Request *StageObjectsJobSpectraS3Request) WithName(name string) *StageObjectsJobSpectraS3Request {
    stageObjectsJobSpectraS3Request.Name = &name
    return stageObjectsJobSpectraS3Request
}

func (stageObjectsJobSpectraS3Request *StageObjectsJobSpectraS3Request) WithPriority(priority Priority) *StageObjectsJobSpectraS3Request {
    stageObjectsJobSpectraS3Request.Priority = priority
    return stageObjectsJobSpectraS3Request
}

