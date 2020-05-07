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

import (
    "fmt"
    "strings"
)

type AbortMultiPartUploadRequest struct {
    BucketName string
    ObjectName string
    UploadId string
}

func NewAbortMultiPartUploadRequest(bucketName string, objectName string, uploadId string) *AbortMultiPartUploadRequest {
    return &AbortMultiPartUploadRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        UploadId: uploadId,
    }
}

type CompleteBlobRequest struct {
    BucketName string
    ObjectName string
    Blob string
    Checksum Checksum
    Job string
    Metadata map[string]string
    Size *int64
}

func NewCompleteBlobRequest(bucketName string, objectName string, blob string, job string) *CompleteBlobRequest {
    return &CompleteBlobRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        Blob: blob,
        Job: job,
        Checksum: NewNoneChecksum(),
        Metadata: make(map[string]string),
    }
}

func (completeBlobRequest *CompleteBlobRequest) WithSize(size int64) *CompleteBlobRequest {
    completeBlobRequest.Size = &size
    return completeBlobRequest
}


func (completeBlobRequest *CompleteBlobRequest) WithChecksum(contentHash string, checksumType ChecksumType) *CompleteBlobRequest {
    completeBlobRequest.Checksum.ContentHash = contentHash
    completeBlobRequest.Checksum.Type = checksumType
    return completeBlobRequest
}

func (completeBlobRequest *CompleteBlobRequest) WithMetaData(key string, values ...string) *CompleteBlobRequest {
    if strings.HasPrefix(strings.ToLower(key), AMZ_META_HEADER) {
        completeBlobRequest.Metadata[key] = strings.Join(values, ",")
    } else {
        completeBlobRequest.Metadata[strings.ToLower(AMZ_META_HEADER + key)] = strings.Join(values, ",")
    }
    return completeBlobRequest
}
type CompleteMultiPartUploadRequest struct {
    BucketName string
    ObjectName string
    Parts []Part
    UploadId string
}

func NewCompleteMultiPartUploadRequest(bucketName string, objectName string, parts []Part, uploadId string) *CompleteMultiPartUploadRequest {
    return &CompleteMultiPartUploadRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        UploadId: uploadId,
        Parts: parts,
    }
}

type PutBucketRequest struct {
    BucketName string
}

func NewPutBucketRequest(bucketName string) *PutBucketRequest {
    return &PutBucketRequest{
        BucketName: bucketName,
    }
}

type PutMultiPartUploadPartRequest struct {
    BucketName string
    ObjectName string
    Content ReaderWithSizeDecorator
    PartNumber int
    UploadId string
}

func NewPutMultiPartUploadPartRequest(bucketName string, objectName string, content ReaderWithSizeDecorator, partNumber int, uploadId string) *PutMultiPartUploadPartRequest {
    return &PutMultiPartUploadPartRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        PartNumber: partNumber,
        UploadId: uploadId,
        Content: content,
    }
}

type PutObjectRequest struct {
    BucketName string
    ObjectName string
    Checksum Checksum
    Content ReaderWithSizeDecorator
    Job *string
    Metadata map[string]string
    Offset *int64
}

func NewPutObjectRequest(bucketName string, objectName string, content ReaderWithSizeDecorator) *PutObjectRequest {
    return &PutObjectRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        Content: content,
        Checksum: NewNoneChecksum(),
        Metadata: make(map[string]string),
    }
}

func (putObjectRequest *PutObjectRequest) WithJob(job string) *PutObjectRequest {
    putObjectRequest.Job = &job
    return putObjectRequest
}

func (putObjectRequest *PutObjectRequest) WithOffset(offset int64) *PutObjectRequest {
    putObjectRequest.Offset = &offset
    return putObjectRequest
}


func (putObjectRequest *PutObjectRequest) WithChecksum(contentHash string, checksumType ChecksumType) *PutObjectRequest {
    putObjectRequest.Checksum.ContentHash = contentHash
    putObjectRequest.Checksum.Type = checksumType
    return putObjectRequest
}

func (putObjectRequest *PutObjectRequest) WithMetaData(key string, values ...string) *PutObjectRequest {
    if strings.HasPrefix(strings.ToLower(key), AMZ_META_HEADER) {
        putObjectRequest.Metadata[key] = strings.Join(values, ",")
    } else {
        putObjectRequest.Metadata[strings.ToLower(AMZ_META_HEADER + key)] = strings.Join(values, ",")
    }
    return putObjectRequest
}
type DeleteBucketRequest struct {
    BucketName string
}

func NewDeleteBucketRequest(bucketName string) *DeleteBucketRequest {
    return &DeleteBucketRequest{
        BucketName: bucketName,
    }
}

type DeleteObjectRequest struct {
    BucketName string
    ObjectName string
    VersionId *string
}

func NewDeleteObjectRequest(bucketName string, objectName string) *DeleteObjectRequest {
    return &DeleteObjectRequest{
        BucketName: bucketName,
        ObjectName: objectName,
    }
}

func (deleteObjectRequest *DeleteObjectRequest) WithVersionId(versionId string) *DeleteObjectRequest {
    deleteObjectRequest.VersionId = &versionId
    return deleteObjectRequest
}

type DeleteObjectsRequest struct {
    BucketName string
    ObjectNames []string
}

func NewDeleteObjectsRequest(bucketName string, objectNames []string) *DeleteObjectsRequest {
    return &DeleteObjectsRequest{
        BucketName: bucketName,
        ObjectNames: objectNames,
    }
}

type GetBucketRequest struct {
    BucketName string
    Delimiter *string
    Marker *string
    MaxKeys *int
    Prefix *string
    Versions bool
}

func NewGetBucketRequest(bucketName string) *GetBucketRequest {
    return &GetBucketRequest{
        BucketName: bucketName,
    }
}

func (getBucketRequest *GetBucketRequest) WithDelimiter(delimiter string) *GetBucketRequest {
    getBucketRequest.Delimiter = &delimiter
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithMarker(marker string) *GetBucketRequest {
    getBucketRequest.Marker = &marker
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithMaxKeys(maxKeys int) *GetBucketRequest {
    getBucketRequest.MaxKeys = &maxKeys
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithPrefix(prefix string) *GetBucketRequest {
    getBucketRequest.Prefix = &prefix
    return getBucketRequest
}

func (getBucketRequest *GetBucketRequest) WithVersions() *GetBucketRequest {
    getBucketRequest.Versions = true
    return getBucketRequest
}

type GetServiceRequest struct {
}

func NewGetServiceRequest() *GetServiceRequest {
    return &GetServiceRequest{
    }
}

type Range struct {
    Start int64
    End int64
}

type GetObjectRequest struct {
    BucketName string
    ObjectName string
    Checksum Checksum
    Job *string
    Metadata map[string]string
    Offset *int64
    VersionId *string
}

func NewGetObjectRequest(bucketName string, objectName string) *GetObjectRequest {
    return &GetObjectRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        Checksum: NewNoneChecksum(),
        Metadata: make(map[string]string),
    }
}

func (getObjectRequest *GetObjectRequest) WithJob(job string) *GetObjectRequest {
    getObjectRequest.Job = &job
    return getObjectRequest
}

func (getObjectRequest *GetObjectRequest) WithOffset(offset int64) *GetObjectRequest {
    getObjectRequest.Offset = &offset
    return getObjectRequest
}

func (getObjectRequest *GetObjectRequest) WithVersionId(versionId string) *GetObjectRequest {
    getObjectRequest.VersionId = &versionId
    return getObjectRequest
}


func (getObjectRequest *GetObjectRequest) WithChecksum(contentHash string, checksumType ChecksumType) *GetObjectRequest {
    getObjectRequest.Checksum.ContentHash = contentHash
    getObjectRequest.Checksum.Type = checksumType
    return getObjectRequest
}

func (getObjectRequest *GetObjectRequest) WithRanges(ranges ...Range) *GetObjectRequest {
    var rangeElements []string
    for _, cur := range ranges {
        rangeElements = append(rangeElements, fmt.Sprintf("%d-%d", cur.Start, cur.End))
    }
    getObjectRequest.Metadata["Range"] = fmt.Sprintf("bytes=%s", strings.Join(rangeElements[:], ","))
    return getObjectRequest
}
type HeadBucketRequest struct {
    BucketName string
}

func NewHeadBucketRequest(bucketName string) *HeadBucketRequest {
    return &HeadBucketRequest{
        BucketName: bucketName,
    }
}

type HeadObjectRequest struct {
    BucketName string
    ObjectName string
    VersionId *string
}

func NewHeadObjectRequest(bucketName string, objectName string) *HeadObjectRequest {
    return &HeadObjectRequest{
        BucketName: bucketName,
        ObjectName: objectName,
    }
}

func (headObjectRequest *HeadObjectRequest) WithVersionId(versionId string) *HeadObjectRequest {
    headObjectRequest.VersionId = &versionId
    return headObjectRequest
}

type InitiateMultiPartUploadRequest struct {
    BucketName string
    ObjectName string
}

func NewInitiateMultiPartUploadRequest(bucketName string, objectName string) *InitiateMultiPartUploadRequest {
    return &InitiateMultiPartUploadRequest{
        BucketName: bucketName,
        ObjectName: objectName,
    }
}

type ListMultiPartUploadPartsRequest struct {
    BucketName string
    ObjectName string
    MaxParts *int
    PartNumberMarker *int
    UploadId string
}

func NewListMultiPartUploadPartsRequest(bucketName string, objectName string, uploadId string) *ListMultiPartUploadPartsRequest {
    return &ListMultiPartUploadPartsRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        UploadId: uploadId,
    }
}

func (listMultiPartUploadPartsRequest *ListMultiPartUploadPartsRequest) WithMaxParts(maxParts int) *ListMultiPartUploadPartsRequest {
    listMultiPartUploadPartsRequest.MaxParts = &maxParts
    return listMultiPartUploadPartsRequest
}

func (listMultiPartUploadPartsRequest *ListMultiPartUploadPartsRequest) WithPartNumberMarker(partNumberMarker int) *ListMultiPartUploadPartsRequest {
    listMultiPartUploadPartsRequest.PartNumberMarker = &partNumberMarker
    return listMultiPartUploadPartsRequest
}

type ListMultiPartUploadsRequest struct {
    BucketName string
    Delimiter *string
    KeyMarker *string
    MaxUploads *int
    Prefix *string
    UploadIdMarker *string
}

func NewListMultiPartUploadsRequest(bucketName string) *ListMultiPartUploadsRequest {
    return &ListMultiPartUploadsRequest{
        BucketName: bucketName,
    }
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithDelimiter(delimiter string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.Delimiter = &delimiter
    return listMultiPartUploadsRequest
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithKeyMarker(keyMarker string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.KeyMarker = &keyMarker
    return listMultiPartUploadsRequest
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithMaxUploads(maxUploads int) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.MaxUploads = &maxUploads
    return listMultiPartUploadsRequest
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithPrefix(prefix string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.Prefix = &prefix
    return listMultiPartUploadsRequest
}

func (listMultiPartUploadsRequest *ListMultiPartUploadsRequest) WithUploadIdMarker(uploadIdMarker string) *ListMultiPartUploadsRequest {
    listMultiPartUploadsRequest.UploadIdMarker = &uploadIdMarker
    return listMultiPartUploadsRequest
}

type PutBucketAclForGroupSpectraS3Request struct {
    BucketId string
    GroupId string
    Permission BucketAclPermission
}

func NewPutBucketAclForGroupSpectraS3Request(bucketId string, groupId string, permission BucketAclPermission) *PutBucketAclForGroupSpectraS3Request {
    return &PutBucketAclForGroupSpectraS3Request{
        BucketId: bucketId,
        GroupId: groupId,
        Permission: permission,
    }
}

type PutBucketAclForUserSpectraS3Request struct {
    BucketId string
    Permission BucketAclPermission
    UserId string
}

func NewPutBucketAclForUserSpectraS3Request(bucketId string, permission BucketAclPermission, userId string) *PutBucketAclForUserSpectraS3Request {
    return &PutBucketAclForUserSpectraS3Request{
        BucketId: bucketId,
        Permission: permission,
        UserId: userId,
    }
}

type PutDataPolicyAclForGroupSpectraS3Request struct {
    DataPolicyId string
    GroupId string
}

func NewPutDataPolicyAclForGroupSpectraS3Request(dataPolicyId string, groupId string) *PutDataPolicyAclForGroupSpectraS3Request {
    return &PutDataPolicyAclForGroupSpectraS3Request{
        DataPolicyId: dataPolicyId,
        GroupId: groupId,
    }
}

type PutDataPolicyAclForUserSpectraS3Request struct {
    DataPolicyId string
    UserId string
}

func NewPutDataPolicyAclForUserSpectraS3Request(dataPolicyId string, userId string) *PutDataPolicyAclForUserSpectraS3Request {
    return &PutDataPolicyAclForUserSpectraS3Request{
        DataPolicyId: dataPolicyId,
        UserId: userId,
    }
}

type PutGlobalBucketAclForGroupSpectraS3Request struct {
    GroupId string
    Permission BucketAclPermission
}

func NewPutGlobalBucketAclForGroupSpectraS3Request(groupId string, permission BucketAclPermission) *PutGlobalBucketAclForGroupSpectraS3Request {
    return &PutGlobalBucketAclForGroupSpectraS3Request{
        GroupId: groupId,
        Permission: permission,
    }
}

type PutGlobalBucketAclForUserSpectraS3Request struct {
    Permission BucketAclPermission
    UserId string
}

func NewPutGlobalBucketAclForUserSpectraS3Request(permission BucketAclPermission, userId string) *PutGlobalBucketAclForUserSpectraS3Request {
    return &PutGlobalBucketAclForUserSpectraS3Request{
        Permission: permission,
        UserId: userId,
    }
}

type PutGlobalDataPolicyAclForGroupSpectraS3Request struct {
    GroupId string
}

func NewPutGlobalDataPolicyAclForGroupSpectraS3Request(groupId string) *PutGlobalDataPolicyAclForGroupSpectraS3Request {
    return &PutGlobalDataPolicyAclForGroupSpectraS3Request{
        GroupId: groupId,
    }
}

type PutGlobalDataPolicyAclForUserSpectraS3Request struct {
    UserId string
}

func NewPutGlobalDataPolicyAclForUserSpectraS3Request(userId string) *PutGlobalDataPolicyAclForUserSpectraS3Request {
    return &PutGlobalDataPolicyAclForUserSpectraS3Request{
        UserId: userId,
    }
}

type DeleteBucketAclSpectraS3Request struct {
    BucketAcl string
}

func NewDeleteBucketAclSpectraS3Request(bucketAcl string) *DeleteBucketAclSpectraS3Request {
    return &DeleteBucketAclSpectraS3Request{
        BucketAcl: bucketAcl,
    }
}

type DeleteDataPolicyAclSpectraS3Request struct {
    DataPolicyAcl string
}

func NewDeleteDataPolicyAclSpectraS3Request(dataPolicyAcl string) *DeleteDataPolicyAclSpectraS3Request {
    return &DeleteDataPolicyAclSpectraS3Request{
        DataPolicyAcl: dataPolicyAcl,
    }
}

type GetBucketAclSpectraS3Request struct {
    BucketAcl string
}

func NewGetBucketAclSpectraS3Request(bucketAcl string) *GetBucketAclSpectraS3Request {
    return &GetBucketAclSpectraS3Request{
        BucketAcl: bucketAcl,
    }
}

type GetBucketAclsSpectraS3Request struct {
    BucketId *string
    GroupId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    Permission BucketAclPermission
    UserId *string
}

func NewGetBucketAclsSpectraS3Request() *GetBucketAclsSpectraS3Request {
    return &GetBucketAclsSpectraS3Request{
    }
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithBucketId(bucketId string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.BucketId = &bucketId
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithGroupId(groupId string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.GroupId = &groupId
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithLastPage() *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.LastPage = true
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPageLength(pageLength int) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.PageLength = &pageLength
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPageOffset(pageOffset int) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.PageOffset = &pageOffset
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithPermission(permission BucketAclPermission) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.Permission = permission
    return getBucketAclsSpectraS3Request
}

func (getBucketAclsSpectraS3Request *GetBucketAclsSpectraS3Request) WithUserId(userId string) *GetBucketAclsSpectraS3Request {
    getBucketAclsSpectraS3Request.UserId = &userId
    return getBucketAclsSpectraS3Request
}

type GetDataPolicyAclSpectraS3Request struct {
    DataPolicyAcl string
}

func NewGetDataPolicyAclSpectraS3Request(dataPolicyAcl string) *GetDataPolicyAclSpectraS3Request {
    return &GetDataPolicyAclSpectraS3Request{
        DataPolicyAcl: dataPolicyAcl,
    }
}

type GetDataPolicyAclsSpectraS3Request struct {
    DataPolicyId *string
    GroupId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetDataPolicyAclsSpectraS3Request() *GetDataPolicyAclsSpectraS3Request {
    return &GetDataPolicyAclsSpectraS3Request{
    }
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithGroupId(groupId string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.GroupId = &groupId
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithLastPage() *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.LastPage = true
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithPageLength(pageLength int) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.PageLength = &pageLength
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithPageOffset(pageOffset int) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.PageOffset = &pageOffset
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDataPolicyAclsSpectraS3Request
}

func (getDataPolicyAclsSpectraS3Request *GetDataPolicyAclsSpectraS3Request) WithUserId(userId string) *GetDataPolicyAclsSpectraS3Request {
    getDataPolicyAclsSpectraS3Request.UserId = &userId
    return getDataPolicyAclsSpectraS3Request
}

type PutBucketSpectraS3Request struct {
    DataPolicyId *string
    Id *string
    Name string
    UserId *string
}

func NewPutBucketSpectraS3Request(name string) *PutBucketSpectraS3Request {
    return &PutBucketSpectraS3Request{
        Name: name,
    }
}

func (putBucketSpectraS3Request *PutBucketSpectraS3Request) WithDataPolicyId(dataPolicyId string) *PutBucketSpectraS3Request {
    putBucketSpectraS3Request.DataPolicyId = &dataPolicyId
    return putBucketSpectraS3Request
}

func (putBucketSpectraS3Request *PutBucketSpectraS3Request) WithId(id string) *PutBucketSpectraS3Request {
    putBucketSpectraS3Request.Id = &id
    return putBucketSpectraS3Request
}

func (putBucketSpectraS3Request *PutBucketSpectraS3Request) WithUserId(userId string) *PutBucketSpectraS3Request {
    putBucketSpectraS3Request.UserId = &userId
    return putBucketSpectraS3Request
}

type DeleteBucketSpectraS3Request struct {
    BucketName string
    Force bool
}

func NewDeleteBucketSpectraS3Request(bucketName string) *DeleteBucketSpectraS3Request {
    return &DeleteBucketSpectraS3Request{
        BucketName: bucketName,
    }
}

func (deleteBucketSpectraS3Request *DeleteBucketSpectraS3Request) WithForce() *DeleteBucketSpectraS3Request {
    deleteBucketSpectraS3Request.Force = true
    return deleteBucketSpectraS3Request
}

type GetBucketSpectraS3Request struct {
    BucketName string
}

func NewGetBucketSpectraS3Request(bucketName string) *GetBucketSpectraS3Request {
    return &GetBucketSpectraS3Request{
        BucketName: bucketName,
    }
}

type GetBucketsSpectraS3Request struct {
    DataPolicyId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetBucketsSpectraS3Request() *GetBucketsSpectraS3Request {
    return &GetBucketsSpectraS3Request{
    }
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.DataPolicyId = &dataPolicyId
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithLastPage() *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.LastPage = true
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithName(name string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.Name = &name
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithPageLength(pageLength int) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.PageLength = &pageLength
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithPageOffset(pageOffset int) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.PageOffset = &pageOffset
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getBucketsSpectraS3Request
}

func (getBucketsSpectraS3Request *GetBucketsSpectraS3Request) WithUserId(userId string) *GetBucketsSpectraS3Request {
    getBucketsSpectraS3Request.UserId = &userId
    return getBucketsSpectraS3Request
}

type ModifyBucketSpectraS3Request struct {
    BucketName string
    DataPolicyId *string
    UserId *string
}

func NewModifyBucketSpectraS3Request(bucketName string) *ModifyBucketSpectraS3Request {
    return &ModifyBucketSpectraS3Request{
        BucketName: bucketName,
    }
}

func (modifyBucketSpectraS3Request *ModifyBucketSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ModifyBucketSpectraS3Request {
    modifyBucketSpectraS3Request.DataPolicyId = &dataPolicyId
    return modifyBucketSpectraS3Request
}

func (modifyBucketSpectraS3Request *ModifyBucketSpectraS3Request) WithUserId(userId string) *ModifyBucketSpectraS3Request {
    modifyBucketSpectraS3Request.UserId = &userId
    return modifyBucketSpectraS3Request
}

type ForceFullCacheReclaimSpectraS3Request struct {
}

func NewForceFullCacheReclaimSpectraS3Request() *ForceFullCacheReclaimSpectraS3Request {
    return &ForceFullCacheReclaimSpectraS3Request{
    }
}

type GetCacheFilesystemSpectraS3Request struct {
    CacheFilesystem string
}

func NewGetCacheFilesystemSpectraS3Request(cacheFilesystem string) *GetCacheFilesystemSpectraS3Request {
    return &GetCacheFilesystemSpectraS3Request{
        CacheFilesystem: cacheFilesystem,
    }
}

type GetCacheFilesystemsSpectraS3Request struct {
    LastPage bool
    NodeId *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetCacheFilesystemsSpectraS3Request() *GetCacheFilesystemsSpectraS3Request {
    return &GetCacheFilesystemsSpectraS3Request{
    }
}

func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) WithLastPage() *GetCacheFilesystemsSpectraS3Request {
    getCacheFilesystemsSpectraS3Request.LastPage = true
    return getCacheFilesystemsSpectraS3Request
}

func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) WithNodeId(nodeId string) *GetCacheFilesystemsSpectraS3Request {
    getCacheFilesystemsSpectraS3Request.NodeId = &nodeId
    return getCacheFilesystemsSpectraS3Request
}

func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) WithPageLength(pageLength int) *GetCacheFilesystemsSpectraS3Request {
    getCacheFilesystemsSpectraS3Request.PageLength = &pageLength
    return getCacheFilesystemsSpectraS3Request
}

func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) WithPageOffset(pageOffset int) *GetCacheFilesystemsSpectraS3Request {
    getCacheFilesystemsSpectraS3Request.PageOffset = &pageOffset
    return getCacheFilesystemsSpectraS3Request
}

func (getCacheFilesystemsSpectraS3Request *GetCacheFilesystemsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetCacheFilesystemsSpectraS3Request {
    getCacheFilesystemsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getCacheFilesystemsSpectraS3Request
}

type GetCacheStateSpectraS3Request struct {
}

func NewGetCacheStateSpectraS3Request() *GetCacheStateSpectraS3Request {
    return &GetCacheStateSpectraS3Request{
    }
}

type ModifyCacheFilesystemSpectraS3Request struct {
    AutoReclaimInitiateThreshold *float64
    AutoReclaimTerminateThreshold *float64
    BurstThreshold *float64
    CacheFilesystem string
    MaxCapacityInBytes *int64
}

func NewModifyCacheFilesystemSpectraS3Request(cacheFilesystem string) *ModifyCacheFilesystemSpectraS3Request {
    return &ModifyCacheFilesystemSpectraS3Request{
        CacheFilesystem: cacheFilesystem,
    }
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithAutoReclaimInitiateThreshold(autoReclaimInitiateThreshold float64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.AutoReclaimInitiateThreshold = &autoReclaimInitiateThreshold
    return modifyCacheFilesystemSpectraS3Request
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithAutoReclaimTerminateThreshold(autoReclaimTerminateThreshold float64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.AutoReclaimTerminateThreshold = &autoReclaimTerminateThreshold
    return modifyCacheFilesystemSpectraS3Request
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithBurstThreshold(burstThreshold float64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.BurstThreshold = &burstThreshold
    return modifyCacheFilesystemSpectraS3Request
}

func (modifyCacheFilesystemSpectraS3Request *ModifyCacheFilesystemSpectraS3Request) WithMaxCapacityInBytes(maxCapacityInBytes int64) *ModifyCacheFilesystemSpectraS3Request {
    modifyCacheFilesystemSpectraS3Request.MaxCapacityInBytes = &maxCapacityInBytes
    return modifyCacheFilesystemSpectraS3Request
}

type GetBucketCapacitySummarySpectraS3Request struct {
    BucketId string
    PoolHealth PoolHealth
    PoolState PoolState
    PoolType PoolType
    StorageDomainId string
    TapeState TapeState
    TapeType *string
}

func NewGetBucketCapacitySummarySpectraS3Request(bucketId string, storageDomainId string) *GetBucketCapacitySummarySpectraS3Request {
    return &GetBucketCapacitySummarySpectraS3Request{
        BucketId: bucketId,
        StorageDomainId: storageDomainId,
    }
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithPoolHealth(poolHealth PoolHealth) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.PoolHealth = poolHealth
    return getBucketCapacitySummarySpectraS3Request
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithPoolState(poolState PoolState) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.PoolState = poolState
    return getBucketCapacitySummarySpectraS3Request
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithPoolType(poolType PoolType) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.PoolType = poolType
    return getBucketCapacitySummarySpectraS3Request
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithTapeState(tapeState TapeState) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.TapeState = tapeState
    return getBucketCapacitySummarySpectraS3Request
}

func (getBucketCapacitySummarySpectraS3Request *GetBucketCapacitySummarySpectraS3Request) WithTapeType(tapeType string) *GetBucketCapacitySummarySpectraS3Request {
    getBucketCapacitySummarySpectraS3Request.TapeType = &tapeType
    return getBucketCapacitySummarySpectraS3Request
}

type GetStorageDomainCapacitySummarySpectraS3Request struct {
    PoolHealth PoolHealth
    PoolState PoolState
    PoolType PoolType
    StorageDomainId string
    TapeState TapeState
    TapeType *string
}

func NewGetStorageDomainCapacitySummarySpectraS3Request(storageDomainId string) *GetStorageDomainCapacitySummarySpectraS3Request {
    return &GetStorageDomainCapacitySummarySpectraS3Request{
        StorageDomainId: storageDomainId,
    }
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithPoolHealth(poolHealth PoolHealth) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.PoolHealth = poolHealth
    return getStorageDomainCapacitySummarySpectraS3Request
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithPoolState(poolState PoolState) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.PoolState = poolState
    return getStorageDomainCapacitySummarySpectraS3Request
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithPoolType(poolType PoolType) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.PoolType = poolType
    return getStorageDomainCapacitySummarySpectraS3Request
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithTapeState(tapeState TapeState) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.TapeState = tapeState
    return getStorageDomainCapacitySummarySpectraS3Request
}

func (getStorageDomainCapacitySummarySpectraS3Request *GetStorageDomainCapacitySummarySpectraS3Request) WithTapeType(tapeType string) *GetStorageDomainCapacitySummarySpectraS3Request {
    getStorageDomainCapacitySummarySpectraS3Request.TapeType = &tapeType
    return getStorageDomainCapacitySummarySpectraS3Request
}

type GetSystemCapacitySummarySpectraS3Request struct {
    PoolHealth PoolHealth
    PoolState PoolState
    PoolType PoolType
    TapeState TapeState
    TapeType *string
}

func NewGetSystemCapacitySummarySpectraS3Request() *GetSystemCapacitySummarySpectraS3Request {
    return &GetSystemCapacitySummarySpectraS3Request{
    }
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithPoolHealth(poolHealth PoolHealth) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.PoolHealth = poolHealth
    return getSystemCapacitySummarySpectraS3Request
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithPoolState(poolState PoolState) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.PoolState = poolState
    return getSystemCapacitySummarySpectraS3Request
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithPoolType(poolType PoolType) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.PoolType = poolType
    return getSystemCapacitySummarySpectraS3Request
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithTapeState(tapeState TapeState) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.TapeState = tapeState
    return getSystemCapacitySummarySpectraS3Request
}

func (getSystemCapacitySummarySpectraS3Request *GetSystemCapacitySummarySpectraS3Request) WithTapeType(tapeType string) *GetSystemCapacitySummarySpectraS3Request {
    getSystemCapacitySummarySpectraS3Request.TapeType = &tapeType
    return getSystemCapacitySummarySpectraS3Request
}

type GetDataPathBackendSpectraS3Request struct {
}

func NewGetDataPathBackendSpectraS3Request() *GetDataPathBackendSpectraS3Request {
    return &GetDataPathBackendSpectraS3Request{
    }
}

type GetDataPlannerBlobStoreTasksSpectraS3Request struct {
    FullDetails bool
}

func NewGetDataPlannerBlobStoreTasksSpectraS3Request() *GetDataPlannerBlobStoreTasksSpectraS3Request {
    return &GetDataPlannerBlobStoreTasksSpectraS3Request{
    }
}

func (getDataPlannerBlobStoreTasksSpectraS3Request *GetDataPlannerBlobStoreTasksSpectraS3Request) WithFullDetails() *GetDataPlannerBlobStoreTasksSpectraS3Request {
    getDataPlannerBlobStoreTasksSpectraS3Request.FullDetails = true
    return getDataPlannerBlobStoreTasksSpectraS3Request
}

type ModifyDataPathBackendSpectraS3Request struct {
    Activated *bool
    AllowNewJobRequests *bool
    AutoActivateTimeoutInMins *int
    AutoInspect AutoInspectMode
    CacheAvailableRetryAfterInSeconds *int
    DefaultVerifyDataAfterImport Priority
    DefaultVerifyDataPriorToImport *bool
    IomEnabled *bool
    PartiallyVerifyLastPercentOfTapes *int
    UnavailableMediaPolicy UnavailableMediaUsagePolicy
    UnavailablePoolMaxJobRetryInMins *int
    UnavailableTapePartitionMaxJobRetryInMins *int
}

func NewModifyDataPathBackendSpectraS3Request() *ModifyDataPathBackendSpectraS3Request {
    return &ModifyDataPathBackendSpectraS3Request{
    }
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithActivated(activated bool) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.Activated = &activated
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithAllowNewJobRequests(allowNewJobRequests bool) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.AllowNewJobRequests = &allowNewJobRequests
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithAutoActivateTimeoutInMins(autoActivateTimeoutInMins int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.AutoActivateTimeoutInMins = &autoActivateTimeoutInMins
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithAutoInspect(autoInspect AutoInspectMode) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.AutoInspect = autoInspect
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithCacheAvailableRetryAfterInSeconds(cacheAvailableRetryAfterInSeconds int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.CacheAvailableRetryAfterInSeconds = &cacheAvailableRetryAfterInSeconds
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithDefaultVerifyDataAfterImport(defaultVerifyDataAfterImport Priority) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.DefaultVerifyDataAfterImport = defaultVerifyDataAfterImport
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithDefaultVerifyDataPriorToImport(defaultVerifyDataPriorToImport bool) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.DefaultVerifyDataPriorToImport = &defaultVerifyDataPriorToImport
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithIomEnabled(iomEnabled bool) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.IomEnabled = &iomEnabled
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithPartiallyVerifyLastPercentOfTapes(partiallyVerifyLastPercentOfTapes int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.PartiallyVerifyLastPercentOfTapes = &partiallyVerifyLastPercentOfTapes
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithUnavailableMediaPolicy(unavailableMediaPolicy UnavailableMediaUsagePolicy) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.UnavailableMediaPolicy = unavailableMediaPolicy
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithUnavailablePoolMaxJobRetryInMins(unavailablePoolMaxJobRetryInMins int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.UnavailablePoolMaxJobRetryInMins = &unavailablePoolMaxJobRetryInMins
    return modifyDataPathBackendSpectraS3Request
}

func (modifyDataPathBackendSpectraS3Request *ModifyDataPathBackendSpectraS3Request) WithUnavailableTapePartitionMaxJobRetryInMins(unavailableTapePartitionMaxJobRetryInMins int) *ModifyDataPathBackendSpectraS3Request {
    modifyDataPathBackendSpectraS3Request.UnavailableTapePartitionMaxJobRetryInMins = &unavailableTapePartitionMaxJobRetryInMins
    return modifyDataPathBackendSpectraS3Request
}

type PutAzureDataReplicationRuleSpectraS3Request struct {
    DataPolicyId string
    DataReplicationRuleType DataReplicationRuleType
    MaxBlobPartSizeInBytes *int64
    ReplicateDeletes *bool
    TargetId string
}

func NewPutAzureDataReplicationRuleSpectraS3Request(dataPolicyId string, dataReplicationRuleType DataReplicationRuleType, targetId string) *PutAzureDataReplicationRuleSpectraS3Request {
    return &PutAzureDataReplicationRuleSpectraS3Request{
        DataPolicyId: dataPolicyId,
        TargetId: targetId,
        DataReplicationRuleType: dataReplicationRuleType,
    }
}

func (putAzureDataReplicationRuleSpectraS3Request *PutAzureDataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *PutAzureDataReplicationRuleSpectraS3Request {
    putAzureDataReplicationRuleSpectraS3Request.MaxBlobPartSizeInBytes = &maxBlobPartSizeInBytes
    return putAzureDataReplicationRuleSpectraS3Request
}

func (putAzureDataReplicationRuleSpectraS3Request *PutAzureDataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *PutAzureDataReplicationRuleSpectraS3Request {
    putAzureDataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return putAzureDataReplicationRuleSpectraS3Request
}

type PutDataPersistenceRuleSpectraS3Request struct {
    DataPersistenceRuleType DataPersistenceRuleType
    DataPolicyId string
    IsolationLevel DataIsolationLevel
    MinimumDaysToRetain *int
    StorageDomainId string
}

func NewPutDataPersistenceRuleSpectraS3Request(dataPersistenceRuleType DataPersistenceRuleType, dataPolicyId string, isolationLevel DataIsolationLevel, storageDomainId string) *PutDataPersistenceRuleSpectraS3Request {
    return &PutDataPersistenceRuleSpectraS3Request{
        DataPolicyId: dataPolicyId,
        IsolationLevel: isolationLevel,
        StorageDomainId: storageDomainId,
        DataPersistenceRuleType: dataPersistenceRuleType,
    }
}

func (putDataPersistenceRuleSpectraS3Request *PutDataPersistenceRuleSpectraS3Request) WithMinimumDaysToRetain(minimumDaysToRetain int) *PutDataPersistenceRuleSpectraS3Request {
    putDataPersistenceRuleSpectraS3Request.MinimumDaysToRetain = &minimumDaysToRetain
    return putDataPersistenceRuleSpectraS3Request
}

type PutDataPolicySpectraS3Request struct {
    AlwaysForcePutJobCreation *bool
    AlwaysMinimizeSpanningAcrossMedia *bool
    BlobbingEnabled *bool
    ChecksumType ChecksumType
    DefaultBlobSize *int64
    DefaultGetJobPriority Priority
    DefaultPutJobPriority Priority
    DefaultVerifyAfterWrite *bool
    DefaultVerifyJobPriority Priority
    EndToEndCrcRequired *bool
    MaxVersionsToKeep *int
    Name string
    RebuildPriority Priority
    Versioning VersioningLevel
}

func NewPutDataPolicySpectraS3Request(name string) *PutDataPolicySpectraS3Request {
    return &PutDataPolicySpectraS3Request{
        Name: name,
    }
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithAlwaysForcePutJobCreation(alwaysForcePutJobCreation bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.AlwaysForcePutJobCreation = &alwaysForcePutJobCreation
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithAlwaysMinimizeSpanningAcrossMedia(alwaysMinimizeSpanningAcrossMedia bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.AlwaysMinimizeSpanningAcrossMedia = &alwaysMinimizeSpanningAcrossMedia
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithBlobbingEnabled(blobbingEnabled bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.BlobbingEnabled = &blobbingEnabled
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithChecksumType(checksumType ChecksumType) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.ChecksumType = checksumType
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultBlobSize(defaultBlobSize int64) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.DefaultBlobSize = &defaultBlobSize
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultGetJobPriority(defaultGetJobPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.DefaultGetJobPriority = defaultGetJobPriority
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultPutJobPriority(defaultPutJobPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.DefaultPutJobPriority = defaultPutJobPriority
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultVerifyAfterWrite(defaultVerifyAfterWrite bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.DefaultVerifyAfterWrite = &defaultVerifyAfterWrite
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithDefaultVerifyJobPriority(defaultVerifyJobPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.DefaultVerifyJobPriority = defaultVerifyJobPriority
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithEndToEndCrcRequired(endToEndCrcRequired bool) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.EndToEndCrcRequired = &endToEndCrcRequired
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithMaxVersionsToKeep(maxVersionsToKeep int) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.MaxVersionsToKeep = &maxVersionsToKeep
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithRebuildPriority(rebuildPriority Priority) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.RebuildPriority = rebuildPriority
    return putDataPolicySpectraS3Request
}

func (putDataPolicySpectraS3Request *PutDataPolicySpectraS3Request) WithVersioning(versioning VersioningLevel) *PutDataPolicySpectraS3Request {
    putDataPolicySpectraS3Request.Versioning = versioning
    return putDataPolicySpectraS3Request
}

type PutDs3DataReplicationRuleSpectraS3Request struct {
    DataPolicyId string
    DataReplicationRuleType DataReplicationRuleType
    ReplicateDeletes *bool
    TargetDataPolicy *string
    TargetId string
}

func NewPutDs3DataReplicationRuleSpectraS3Request(dataPolicyId string, dataReplicationRuleType DataReplicationRuleType, targetId string) *PutDs3DataReplicationRuleSpectraS3Request {
    return &PutDs3DataReplicationRuleSpectraS3Request{
        DataPolicyId: dataPolicyId,
        TargetId: targetId,
        DataReplicationRuleType: dataReplicationRuleType,
    }
}

func (putDs3DataReplicationRuleSpectraS3Request *PutDs3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *PutDs3DataReplicationRuleSpectraS3Request {
    putDs3DataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return putDs3DataReplicationRuleSpectraS3Request
}

func (putDs3DataReplicationRuleSpectraS3Request *PutDs3DataReplicationRuleSpectraS3Request) WithTargetDataPolicy(targetDataPolicy string) *PutDs3DataReplicationRuleSpectraS3Request {
    putDs3DataReplicationRuleSpectraS3Request.TargetDataPolicy = &targetDataPolicy
    return putDs3DataReplicationRuleSpectraS3Request
}

type PutS3DataReplicationRuleSpectraS3Request struct {
    DataPolicyId string
    DataReplicationRuleType DataReplicationRuleType
    InitialDataPlacement S3InitialDataPlacementPolicy
    MaxBlobPartSizeInBytes *int64
    ReplicateDeletes *bool
    TargetId string
}

func NewPutS3DataReplicationRuleSpectraS3Request(dataPolicyId string, dataReplicationRuleType DataReplicationRuleType, targetId string) *PutS3DataReplicationRuleSpectraS3Request {
    return &PutS3DataReplicationRuleSpectraS3Request{
        DataPolicyId: dataPolicyId,
        TargetId: targetId,
        DataReplicationRuleType: dataReplicationRuleType,
    }
}

func (putS3DataReplicationRuleSpectraS3Request *PutS3DataReplicationRuleSpectraS3Request) WithInitialDataPlacement(initialDataPlacement S3InitialDataPlacementPolicy) *PutS3DataReplicationRuleSpectraS3Request {
    putS3DataReplicationRuleSpectraS3Request.InitialDataPlacement = initialDataPlacement
    return putS3DataReplicationRuleSpectraS3Request
}

func (putS3DataReplicationRuleSpectraS3Request *PutS3DataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *PutS3DataReplicationRuleSpectraS3Request {
    putS3DataReplicationRuleSpectraS3Request.MaxBlobPartSizeInBytes = &maxBlobPartSizeInBytes
    return putS3DataReplicationRuleSpectraS3Request
}

func (putS3DataReplicationRuleSpectraS3Request *PutS3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *PutS3DataReplicationRuleSpectraS3Request {
    putS3DataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return putS3DataReplicationRuleSpectraS3Request
}

type DeleteAzureDataReplicationRuleSpectraS3Request struct {
    AzureDataReplicationRule string
}

func NewDeleteAzureDataReplicationRuleSpectraS3Request(azureDataReplicationRule string) *DeleteAzureDataReplicationRuleSpectraS3Request {
    return &DeleteAzureDataReplicationRuleSpectraS3Request{
        AzureDataReplicationRule: azureDataReplicationRule,
    }
}

type DeleteDataPersistenceRuleSpectraS3Request struct {
    DataPersistenceRuleId string
}

func NewDeleteDataPersistenceRuleSpectraS3Request(dataPersistenceRuleId string) *DeleteDataPersistenceRuleSpectraS3Request {
    return &DeleteDataPersistenceRuleSpectraS3Request{
        DataPersistenceRuleId: dataPersistenceRuleId,
    }
}

type DeleteDataPolicySpectraS3Request struct {
    DataPolicyId string
}

func NewDeleteDataPolicySpectraS3Request(dataPolicyId string) *DeleteDataPolicySpectraS3Request {
    return &DeleteDataPolicySpectraS3Request{
        DataPolicyId: dataPolicyId,
    }
}

type DeleteDs3DataReplicationRuleSpectraS3Request struct {
    Ds3DataReplicationRule string
}

func NewDeleteDs3DataReplicationRuleSpectraS3Request(ds3DataReplicationRule string) *DeleteDs3DataReplicationRuleSpectraS3Request {
    return &DeleteDs3DataReplicationRuleSpectraS3Request{
        Ds3DataReplicationRule: ds3DataReplicationRule,
    }
}

type DeleteS3DataReplicationRuleSpectraS3Request struct {
    S3DataReplicationRule string
}

func NewDeleteS3DataReplicationRuleSpectraS3Request(s3DataReplicationRule string) *DeleteS3DataReplicationRuleSpectraS3Request {
    return &DeleteS3DataReplicationRuleSpectraS3Request{
        S3DataReplicationRule: s3DataReplicationRule,
    }
}

type GetAzureDataReplicationRuleSpectraS3Request struct {
    AzureDataReplicationRule string
}

func NewGetAzureDataReplicationRuleSpectraS3Request(azureDataReplicationRule string) *GetAzureDataReplicationRuleSpectraS3Request {
    return &GetAzureDataReplicationRuleSpectraS3Request{
        AzureDataReplicationRule: azureDataReplicationRule,
    }
}

type GetAzureDataReplicationRulesSpectraS3Request struct {
    DataPolicyId *string
    DataReplicationRuleType DataReplicationRuleType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    ReplicateDeletes *bool
    State DataPlacementRuleState
    TargetId *string
}

func NewGetAzureDataReplicationRulesSpectraS3Request() *GetAzureDataReplicationRulesSpectraS3Request {
    return &GetAzureDataReplicationRulesSpectraS3Request{
    }
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithLastPage() *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.LastPage = true
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.State = state
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getAzureDataReplicationRulesSpectraS3Request
}

func (getAzureDataReplicationRulesSpectraS3Request *GetAzureDataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetAzureDataReplicationRulesSpectraS3Request {
    getAzureDataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getAzureDataReplicationRulesSpectraS3Request
}

type GetDataPersistenceRuleSpectraS3Request struct {
    DataPersistenceRuleId string
}

func NewGetDataPersistenceRuleSpectraS3Request(dataPersistenceRuleId string) *GetDataPersistenceRuleSpectraS3Request {
    return &GetDataPersistenceRuleSpectraS3Request{
        DataPersistenceRuleId: dataPersistenceRuleId,
    }
}

type GetDataPersistenceRulesSpectraS3Request struct {
    DataPersistenceRuleType DataPersistenceRuleType
    DataPolicyId *string
    IsolationLevel DataIsolationLevel
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    State DataPlacementRuleState
    StorageDomainId *string
}

func NewGetDataPersistenceRulesSpectraS3Request() *GetDataPersistenceRulesSpectraS3Request {
    return &GetDataPersistenceRulesSpectraS3Request{
    }
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithIsolationLevel(isolationLevel DataIsolationLevel) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.IsolationLevel = isolationLevel
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithLastPage() *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.LastPage = true
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithPageLength(pageLength int) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.PageLength = &pageLength
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.PageOffset = &pageOffset
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.State = state
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.StorageDomainId = &storageDomainId
    return getDataPersistenceRulesSpectraS3Request
}

func (getDataPersistenceRulesSpectraS3Request *GetDataPersistenceRulesSpectraS3Request) WithDataPersistenceRuleType(dataPersistenceRuleType DataPersistenceRuleType) *GetDataPersistenceRulesSpectraS3Request {
    getDataPersistenceRulesSpectraS3Request.DataPersistenceRuleType = dataPersistenceRuleType
    return getDataPersistenceRulesSpectraS3Request
}

type GetDataPoliciesSpectraS3Request struct {
    AlwaysForcePutJobCreation *bool
    AlwaysMinimizeSpanningAcrossMedia *bool
    ChecksumType ChecksumType
    EndToEndCrcRequired *bool
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetDataPoliciesSpectraS3Request() *GetDataPoliciesSpectraS3Request {
    return &GetDataPoliciesSpectraS3Request{
    }
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithAlwaysForcePutJobCreation(alwaysForcePutJobCreation bool) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.AlwaysForcePutJobCreation = &alwaysForcePutJobCreation
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithAlwaysMinimizeSpanningAcrossMedia(alwaysMinimizeSpanningAcrossMedia bool) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.AlwaysMinimizeSpanningAcrossMedia = &alwaysMinimizeSpanningAcrossMedia
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithChecksumType(checksumType ChecksumType) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.ChecksumType = checksumType
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithEndToEndCrcRequired(endToEndCrcRequired bool) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.EndToEndCrcRequired = &endToEndCrcRequired
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithLastPage() *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.LastPage = true
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithName(name string) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.Name = &name
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithPageLength(pageLength int) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.PageLength = &pageLength
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithPageOffset(pageOffset int) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.PageOffset = &pageOffset
    return getDataPoliciesSpectraS3Request
}

func (getDataPoliciesSpectraS3Request *GetDataPoliciesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDataPoliciesSpectraS3Request {
    getDataPoliciesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDataPoliciesSpectraS3Request
}

type GetDataPolicySpectraS3Request struct {
    DataPolicyId string
}

func NewGetDataPolicySpectraS3Request(dataPolicyId string) *GetDataPolicySpectraS3Request {
    return &GetDataPolicySpectraS3Request{
        DataPolicyId: dataPolicyId,
    }
}

type GetDs3DataReplicationRuleSpectraS3Request struct {
    Ds3DataReplicationRule string
}

func NewGetDs3DataReplicationRuleSpectraS3Request(ds3DataReplicationRule string) *GetDs3DataReplicationRuleSpectraS3Request {
    return &GetDs3DataReplicationRuleSpectraS3Request{
        Ds3DataReplicationRule: ds3DataReplicationRule,
    }
}

type GetDs3DataReplicationRulesSpectraS3Request struct {
    DataPolicyId *string
    DataReplicationRuleType DataReplicationRuleType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    ReplicateDeletes *bool
    State DataPlacementRuleState
    TargetId *string
}

func NewGetDs3DataReplicationRulesSpectraS3Request() *GetDs3DataReplicationRulesSpectraS3Request {
    return &GetDs3DataReplicationRulesSpectraS3Request{
    }
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithLastPage() *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.LastPage = true
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.State = state
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getDs3DataReplicationRulesSpectraS3Request
}

func (getDs3DataReplicationRulesSpectraS3Request *GetDs3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDs3DataReplicationRulesSpectraS3Request {
    getDs3DataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getDs3DataReplicationRulesSpectraS3Request
}

type GetS3DataReplicationRuleSpectraS3Request struct {
    S3DataReplicationRule string
}

func NewGetS3DataReplicationRuleSpectraS3Request(s3DataReplicationRule string) *GetS3DataReplicationRuleSpectraS3Request {
    return &GetS3DataReplicationRuleSpectraS3Request{
        S3DataReplicationRule: s3DataReplicationRule,
    }
}

type GetS3DataReplicationRulesSpectraS3Request struct {
    DataPolicyId *string
    DataReplicationRuleType DataReplicationRuleType
    InitialDataPlacement S3InitialDataPlacementPolicy
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    ReplicateDeletes *bool
    State DataPlacementRuleState
    TargetId *string
}

func NewGetS3DataReplicationRulesSpectraS3Request() *GetS3DataReplicationRulesSpectraS3Request {
    return &GetS3DataReplicationRulesSpectraS3Request{
    }
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithInitialDataPlacement(initialDataPlacement S3InitialDataPlacementPolicy) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.InitialDataPlacement = initialDataPlacement
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithLastPage() *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.LastPage = true
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.State = state
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getS3DataReplicationRulesSpectraS3Request
}

func (getS3DataReplicationRulesSpectraS3Request *GetS3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetS3DataReplicationRulesSpectraS3Request {
    getS3DataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getS3DataReplicationRulesSpectraS3Request
}

type ModifyAzureDataReplicationRuleSpectraS3Request struct {
    AzureDataReplicationRule string
    DataReplicationRuleType DataReplicationRuleType
    MaxBlobPartSizeInBytes *int64
    ReplicateDeletes *bool
}

func NewModifyAzureDataReplicationRuleSpectraS3Request(azureDataReplicationRule string) *ModifyAzureDataReplicationRuleSpectraS3Request {
    return &ModifyAzureDataReplicationRuleSpectraS3Request{
        AzureDataReplicationRule: azureDataReplicationRule,
    }
}

func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *ModifyAzureDataReplicationRuleSpectraS3Request {
    modifyAzureDataReplicationRuleSpectraS3Request.MaxBlobPartSizeInBytes = &maxBlobPartSizeInBytes
    return modifyAzureDataReplicationRuleSpectraS3Request
}

func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *ModifyAzureDataReplicationRuleSpectraS3Request {
    modifyAzureDataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return modifyAzureDataReplicationRuleSpectraS3Request
}

func (modifyAzureDataReplicationRuleSpectraS3Request *ModifyAzureDataReplicationRuleSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *ModifyAzureDataReplicationRuleSpectraS3Request {
    modifyAzureDataReplicationRuleSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return modifyAzureDataReplicationRuleSpectraS3Request
}

type ModifyDataPersistenceRuleSpectraS3Request struct {
    DataPersistenceRuleId string
    DataPersistenceRuleType DataPersistenceRuleType
    IsolationLevel DataIsolationLevel
    MinimumDaysToRetain *int
}

func NewModifyDataPersistenceRuleSpectraS3Request(dataPersistenceRuleId string) *ModifyDataPersistenceRuleSpectraS3Request {
    return &ModifyDataPersistenceRuleSpectraS3Request{
        DataPersistenceRuleId: dataPersistenceRuleId,
    }
}

func (modifyDataPersistenceRuleSpectraS3Request *ModifyDataPersistenceRuleSpectraS3Request) WithIsolationLevel(isolationLevel DataIsolationLevel) *ModifyDataPersistenceRuleSpectraS3Request {
    modifyDataPersistenceRuleSpectraS3Request.IsolationLevel = isolationLevel
    return modifyDataPersistenceRuleSpectraS3Request
}

func (modifyDataPersistenceRuleSpectraS3Request *ModifyDataPersistenceRuleSpectraS3Request) WithMinimumDaysToRetain(minimumDaysToRetain int) *ModifyDataPersistenceRuleSpectraS3Request {
    modifyDataPersistenceRuleSpectraS3Request.MinimumDaysToRetain = &minimumDaysToRetain
    return modifyDataPersistenceRuleSpectraS3Request
}

func (modifyDataPersistenceRuleSpectraS3Request *ModifyDataPersistenceRuleSpectraS3Request) WithDataPersistenceRuleType(dataPersistenceRuleType DataPersistenceRuleType) *ModifyDataPersistenceRuleSpectraS3Request {
    modifyDataPersistenceRuleSpectraS3Request.DataPersistenceRuleType = dataPersistenceRuleType
    return modifyDataPersistenceRuleSpectraS3Request
}

type ModifyDataPolicySpectraS3Request struct {
    AlwaysForcePutJobCreation *bool
    AlwaysMinimizeSpanningAcrossMedia *bool
    BlobbingEnabled *bool
    ChecksumType ChecksumType
    DataPolicyId string
    DefaultBlobSize *int64
    DefaultGetJobPriority Priority
    DefaultPutJobPriority Priority
    DefaultVerifyAfterWrite *bool
    DefaultVerifyJobPriority Priority
    EndToEndCrcRequired *bool
    MaxVersionsToKeep *int
    Name *string
    RebuildPriority Priority
    Versioning VersioningLevel
}

func NewModifyDataPolicySpectraS3Request(dataPolicyId string) *ModifyDataPolicySpectraS3Request {
    return &ModifyDataPolicySpectraS3Request{
        DataPolicyId: dataPolicyId,
    }
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithAlwaysForcePutJobCreation(alwaysForcePutJobCreation bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.AlwaysForcePutJobCreation = &alwaysForcePutJobCreation
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithAlwaysMinimizeSpanningAcrossMedia(alwaysMinimizeSpanningAcrossMedia bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.AlwaysMinimizeSpanningAcrossMedia = &alwaysMinimizeSpanningAcrossMedia
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithBlobbingEnabled(blobbingEnabled bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.BlobbingEnabled = &blobbingEnabled
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithChecksumType(checksumType ChecksumType) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.ChecksumType = checksumType
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultBlobSize(defaultBlobSize int64) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.DefaultBlobSize = &defaultBlobSize
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultGetJobPriority(defaultGetJobPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.DefaultGetJobPriority = defaultGetJobPriority
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultPutJobPriority(defaultPutJobPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.DefaultPutJobPriority = defaultPutJobPriority
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultVerifyAfterWrite(defaultVerifyAfterWrite bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.DefaultVerifyAfterWrite = &defaultVerifyAfterWrite
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithDefaultVerifyJobPriority(defaultVerifyJobPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.DefaultVerifyJobPriority = defaultVerifyJobPriority
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithEndToEndCrcRequired(endToEndCrcRequired bool) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.EndToEndCrcRequired = &endToEndCrcRequired
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithMaxVersionsToKeep(maxVersionsToKeep int) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.MaxVersionsToKeep = &maxVersionsToKeep
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithName(name string) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.Name = &name
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithRebuildPriority(rebuildPriority Priority) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.RebuildPriority = rebuildPriority
    return modifyDataPolicySpectraS3Request
}

func (modifyDataPolicySpectraS3Request *ModifyDataPolicySpectraS3Request) WithVersioning(versioning VersioningLevel) *ModifyDataPolicySpectraS3Request {
    modifyDataPolicySpectraS3Request.Versioning = versioning
    return modifyDataPolicySpectraS3Request
}

type ModifyDs3DataReplicationRuleSpectraS3Request struct {
    DataReplicationRuleType DataReplicationRuleType
    Ds3DataReplicationRule string
    ReplicateDeletes *bool
    TargetDataPolicy *string
}

func NewModifyDs3DataReplicationRuleSpectraS3Request(ds3DataReplicationRule string) *ModifyDs3DataReplicationRuleSpectraS3Request {
    return &ModifyDs3DataReplicationRuleSpectraS3Request{
        Ds3DataReplicationRule: ds3DataReplicationRule,
    }
}

func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *ModifyDs3DataReplicationRuleSpectraS3Request {
    modifyDs3DataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return modifyDs3DataReplicationRuleSpectraS3Request
}

func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) WithTargetDataPolicy(targetDataPolicy string) *ModifyDs3DataReplicationRuleSpectraS3Request {
    modifyDs3DataReplicationRuleSpectraS3Request.TargetDataPolicy = &targetDataPolicy
    return modifyDs3DataReplicationRuleSpectraS3Request
}

func (modifyDs3DataReplicationRuleSpectraS3Request *ModifyDs3DataReplicationRuleSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *ModifyDs3DataReplicationRuleSpectraS3Request {
    modifyDs3DataReplicationRuleSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return modifyDs3DataReplicationRuleSpectraS3Request
}

type ModifyS3DataReplicationRuleSpectraS3Request struct {
    DataReplicationRuleType DataReplicationRuleType
    InitialDataPlacement S3InitialDataPlacementPolicy
    MaxBlobPartSizeInBytes *int64
    ReplicateDeletes *bool
    S3DataReplicationRule string
}

func NewModifyS3DataReplicationRuleSpectraS3Request(s3DataReplicationRule string) *ModifyS3DataReplicationRuleSpectraS3Request {
    return &ModifyS3DataReplicationRuleSpectraS3Request{
        S3DataReplicationRule: s3DataReplicationRule,
    }
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithInitialDataPlacement(initialDataPlacement S3InitialDataPlacementPolicy) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.InitialDataPlacement = initialDataPlacement
    return modifyS3DataReplicationRuleSpectraS3Request
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithMaxBlobPartSizeInBytes(maxBlobPartSizeInBytes int64) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.MaxBlobPartSizeInBytes = &maxBlobPartSizeInBytes
    return modifyS3DataReplicationRuleSpectraS3Request
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithReplicateDeletes(replicateDeletes bool) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.ReplicateDeletes = &replicateDeletes
    return modifyS3DataReplicationRuleSpectraS3Request
}

func (modifyS3DataReplicationRuleSpectraS3Request *ModifyS3DataReplicationRuleSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *ModifyS3DataReplicationRuleSpectraS3Request {
    modifyS3DataReplicationRuleSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return modifyS3DataReplicationRuleSpectraS3Request
}

type ClearSuspectBlobAzureTargetsSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewClearSuspectBlobAzureTargetsSpectraS3Request(ids []string) *ClearSuspectBlobAzureTargetsSpectraS3Request {
    return &ClearSuspectBlobAzureTargetsSpectraS3Request{
        Ids: ids,
    }
}

func (clearSuspectBlobAzureTargetsSpectraS3Request *ClearSuspectBlobAzureTargetsSpectraS3Request) WithForce() *ClearSuspectBlobAzureTargetsSpectraS3Request {
    clearSuspectBlobAzureTargetsSpectraS3Request.Force = true
    return clearSuspectBlobAzureTargetsSpectraS3Request
}

type ClearSuspectBlobDs3TargetsSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewClearSuspectBlobDs3TargetsSpectraS3Request(ids []string) *ClearSuspectBlobDs3TargetsSpectraS3Request {
    return &ClearSuspectBlobDs3TargetsSpectraS3Request{
        Ids: ids,
    }
}

func (clearSuspectBlobDs3TargetsSpectraS3Request *ClearSuspectBlobDs3TargetsSpectraS3Request) WithForce() *ClearSuspectBlobDs3TargetsSpectraS3Request {
    clearSuspectBlobDs3TargetsSpectraS3Request.Force = true
    return clearSuspectBlobDs3TargetsSpectraS3Request
}

type ClearSuspectBlobPoolsSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewClearSuspectBlobPoolsSpectraS3Request(ids []string) *ClearSuspectBlobPoolsSpectraS3Request {
    return &ClearSuspectBlobPoolsSpectraS3Request{
        Ids: ids,
    }
}

func (clearSuspectBlobPoolsSpectraS3Request *ClearSuspectBlobPoolsSpectraS3Request) WithForce() *ClearSuspectBlobPoolsSpectraS3Request {
    clearSuspectBlobPoolsSpectraS3Request.Force = true
    return clearSuspectBlobPoolsSpectraS3Request
}

type ClearSuspectBlobS3TargetsSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewClearSuspectBlobS3TargetsSpectraS3Request(ids []string) *ClearSuspectBlobS3TargetsSpectraS3Request {
    return &ClearSuspectBlobS3TargetsSpectraS3Request{
        Ids: ids,
    }
}

func (clearSuspectBlobS3TargetsSpectraS3Request *ClearSuspectBlobS3TargetsSpectraS3Request) WithForce() *ClearSuspectBlobS3TargetsSpectraS3Request {
    clearSuspectBlobS3TargetsSpectraS3Request.Force = true
    return clearSuspectBlobS3TargetsSpectraS3Request
}

type ClearSuspectBlobTapesSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewClearSuspectBlobTapesSpectraS3Request(ids []string) *ClearSuspectBlobTapesSpectraS3Request {
    return &ClearSuspectBlobTapesSpectraS3Request{
        Ids: ids,
    }
}

func (clearSuspectBlobTapesSpectraS3Request *ClearSuspectBlobTapesSpectraS3Request) WithForce() *ClearSuspectBlobTapesSpectraS3Request {
    clearSuspectBlobTapesSpectraS3Request.Force = true
    return clearSuspectBlobTapesSpectraS3Request
}

type GetDegradedAzureDataReplicationRulesSpectraS3Request struct {
    DataPolicyId *string
    DataReplicationRuleType DataReplicationRuleType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    State DataPlacementRuleState
    TargetId *string
}

func NewGetDegradedAzureDataReplicationRulesSpectraS3Request() *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    return &GetDegradedAzureDataReplicationRulesSpectraS3Request{
    }
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithLastPage() *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.LastPage = true
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.State = state
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

func (getDegradedAzureDataReplicationRulesSpectraS3Request *GetDegradedAzureDataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDegradedAzureDataReplicationRulesSpectraS3Request {
    getDegradedAzureDataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getDegradedAzureDataReplicationRulesSpectraS3Request
}

type GetDegradedBlobsSpectraS3Request struct {
    BlobId *string
    BucketId *string
    Ds3ReplicationRuleId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PersistenceRuleId *string
}

func NewGetDegradedBlobsSpectraS3Request() *GetDegradedBlobsSpectraS3Request {
    return &GetDegradedBlobsSpectraS3Request{
    }
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithBlobId(blobId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.BlobId = &blobId
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithBucketId(bucketId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.BucketId = &bucketId
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithDs3ReplicationRuleId(ds3ReplicationRuleId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.Ds3ReplicationRuleId = &ds3ReplicationRuleId
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithLastPage() *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.LastPage = true
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPageLength(pageLength int) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.PageLength = &pageLength
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.PageOffset = &pageOffset
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedBlobsSpectraS3Request
}

func (getDegradedBlobsSpectraS3Request *GetDegradedBlobsSpectraS3Request) WithPersistenceRuleId(persistenceRuleId string) *GetDegradedBlobsSpectraS3Request {
    getDegradedBlobsSpectraS3Request.PersistenceRuleId = &persistenceRuleId
    return getDegradedBlobsSpectraS3Request
}

type GetDegradedBucketsSpectraS3Request struct {
    DataPolicyId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetDegradedBucketsSpectraS3Request() *GetDegradedBucketsSpectraS3Request {
    return &GetDegradedBucketsSpectraS3Request{
    }
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithLastPage() *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.LastPage = true
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithName(name string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.Name = &name
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithPageLength(pageLength int) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.PageLength = &pageLength
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.PageOffset = &pageOffset
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedBucketsSpectraS3Request
}

func (getDegradedBucketsSpectraS3Request *GetDegradedBucketsSpectraS3Request) WithUserId(userId string) *GetDegradedBucketsSpectraS3Request {
    getDegradedBucketsSpectraS3Request.UserId = &userId
    return getDegradedBucketsSpectraS3Request
}

type GetDegradedDataPersistenceRulesSpectraS3Request struct {
    DataPersistenceRuleType DataPersistenceRuleType
    DataPolicyId *string
    IsolationLevel DataIsolationLevel
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    State DataPlacementRuleState
    StorageDomainId *string
}

func NewGetDegradedDataPersistenceRulesSpectraS3Request() *GetDegradedDataPersistenceRulesSpectraS3Request {
    return &GetDegradedDataPersistenceRulesSpectraS3Request{
    }
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithIsolationLevel(isolationLevel DataIsolationLevel) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.IsolationLevel = isolationLevel
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithLastPage() *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.LastPage = true
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.PageLength = &pageLength
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.PageOffset = &pageOffset
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.State = state
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.StorageDomainId = &storageDomainId
    return getDegradedDataPersistenceRulesSpectraS3Request
}

func (getDegradedDataPersistenceRulesSpectraS3Request *GetDegradedDataPersistenceRulesSpectraS3Request) WithDataPersistenceRuleType(dataPersistenceRuleType DataPersistenceRuleType) *GetDegradedDataPersistenceRulesSpectraS3Request {
    getDegradedDataPersistenceRulesSpectraS3Request.DataPersistenceRuleType = dataPersistenceRuleType
    return getDegradedDataPersistenceRulesSpectraS3Request
}

type GetDegradedDs3DataReplicationRulesSpectraS3Request struct {
    DataPolicyId *string
    DataReplicationRuleType DataReplicationRuleType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    State DataPlacementRuleState
    TargetId *string
}

func NewGetDegradedDs3DataReplicationRulesSpectraS3Request() *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    return &GetDegradedDs3DataReplicationRulesSpectraS3Request{
    }
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithLastPage() *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.LastPage = true
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.State = state
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

func (getDegradedDs3DataReplicationRulesSpectraS3Request *GetDegradedDs3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDegradedDs3DataReplicationRulesSpectraS3Request {
    getDegradedDs3DataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getDegradedDs3DataReplicationRulesSpectraS3Request
}

type GetDegradedS3DataReplicationRulesSpectraS3Request struct {
    DataPolicyId *string
    DataReplicationRuleType DataReplicationRuleType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    State DataPlacementRuleState
    TargetId *string
}

func NewGetDegradedS3DataReplicationRulesSpectraS3Request() *GetDegradedS3DataReplicationRulesSpectraS3Request {
    return &GetDegradedS3DataReplicationRulesSpectraS3Request{
    }
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.DataPolicyId = &dataPolicyId
    return getDegradedS3DataReplicationRulesSpectraS3Request
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithLastPage() *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.LastPage = true
    return getDegradedS3DataReplicationRulesSpectraS3Request
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithPageLength(pageLength int) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.PageLength = &pageLength
    return getDegradedS3DataReplicationRulesSpectraS3Request
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithPageOffset(pageOffset int) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.PageOffset = &pageOffset
    return getDegradedS3DataReplicationRulesSpectraS3Request
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDegradedS3DataReplicationRulesSpectraS3Request
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithState(state DataPlacementRuleState) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.State = state
    return getDegradedS3DataReplicationRulesSpectraS3Request
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithTargetId(targetId string) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.TargetId = &targetId
    return getDegradedS3DataReplicationRulesSpectraS3Request
}

func (getDegradedS3DataReplicationRulesSpectraS3Request *GetDegradedS3DataReplicationRulesSpectraS3Request) WithDataReplicationRuleType(dataReplicationRuleType DataReplicationRuleType) *GetDegradedS3DataReplicationRulesSpectraS3Request {
    getDegradedS3DataReplicationRulesSpectraS3Request.DataReplicationRuleType = dataReplicationRuleType
    return getDegradedS3DataReplicationRulesSpectraS3Request
}

type GetSuspectBlobAzureTargetsSpectraS3Request struct {
    BlobId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetId *string
}

func NewGetSuspectBlobAzureTargetsSpectraS3Request() *GetSuspectBlobAzureTargetsSpectraS3Request {
    return &GetSuspectBlobAzureTargetsSpectraS3Request{
    }
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.BlobId = &blobId
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithLastPage() *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.LastPage = true
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.PageLength = &pageLength
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.PageOffset = &pageOffset
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectBlobAzureTargetsSpectraS3Request
}

func (getSuspectBlobAzureTargetsSpectraS3Request *GetSuspectBlobAzureTargetsSpectraS3Request) WithTargetId(targetId string) *GetSuspectBlobAzureTargetsSpectraS3Request {
    getSuspectBlobAzureTargetsSpectraS3Request.TargetId = &targetId
    return getSuspectBlobAzureTargetsSpectraS3Request
}

type GetSuspectBlobDs3TargetsSpectraS3Request struct {
    BlobId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetId *string
}

func NewGetSuspectBlobDs3TargetsSpectraS3Request() *GetSuspectBlobDs3TargetsSpectraS3Request {
    return &GetSuspectBlobDs3TargetsSpectraS3Request{
    }
}

func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.BlobId = &blobId
    return getSuspectBlobDs3TargetsSpectraS3Request
}

func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithLastPage() *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.LastPage = true
    return getSuspectBlobDs3TargetsSpectraS3Request
}

func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.PageLength = &pageLength
    return getSuspectBlobDs3TargetsSpectraS3Request
}

func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.PageOffset = &pageOffset
    return getSuspectBlobDs3TargetsSpectraS3Request
}

func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectBlobDs3TargetsSpectraS3Request
}

func (getSuspectBlobDs3TargetsSpectraS3Request *GetSuspectBlobDs3TargetsSpectraS3Request) WithTargetId(targetId string) *GetSuspectBlobDs3TargetsSpectraS3Request {
    getSuspectBlobDs3TargetsSpectraS3Request.TargetId = &targetId
    return getSuspectBlobDs3TargetsSpectraS3Request
}

type GetSuspectBlobPoolsSpectraS3Request struct {
    BlobId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PoolId *string
}

func NewGetSuspectBlobPoolsSpectraS3Request() *GetSuspectBlobPoolsSpectraS3Request {
    return &GetSuspectBlobPoolsSpectraS3Request{
    }
}

func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.BlobId = &blobId
    return getSuspectBlobPoolsSpectraS3Request
}

func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithLastPage() *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.LastPage = true
    return getSuspectBlobPoolsSpectraS3Request
}

func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.PageLength = &pageLength
    return getSuspectBlobPoolsSpectraS3Request
}

func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.PageOffset = &pageOffset
    return getSuspectBlobPoolsSpectraS3Request
}

func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectBlobPoolsSpectraS3Request
}

func (getSuspectBlobPoolsSpectraS3Request *GetSuspectBlobPoolsSpectraS3Request) WithPoolId(poolId string) *GetSuspectBlobPoolsSpectraS3Request {
    getSuspectBlobPoolsSpectraS3Request.PoolId = &poolId
    return getSuspectBlobPoolsSpectraS3Request
}

type GetSuspectBlobS3TargetsSpectraS3Request struct {
    BlobId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetId *string
}

func NewGetSuspectBlobS3TargetsSpectraS3Request() *GetSuspectBlobS3TargetsSpectraS3Request {
    return &GetSuspectBlobS3TargetsSpectraS3Request{
    }
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.BlobId = &blobId
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithLastPage() *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.LastPage = true
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.PageLength = &pageLength
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.PageOffset = &pageOffset
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectBlobS3TargetsSpectraS3Request
}

func (getSuspectBlobS3TargetsSpectraS3Request *GetSuspectBlobS3TargetsSpectraS3Request) WithTargetId(targetId string) *GetSuspectBlobS3TargetsSpectraS3Request {
    getSuspectBlobS3TargetsSpectraS3Request.TargetId = &targetId
    return getSuspectBlobS3TargetsSpectraS3Request
}

type GetSuspectBlobTapesSpectraS3Request struct {
    BlobId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TapeId *string
}

func NewGetSuspectBlobTapesSpectraS3Request() *GetSuspectBlobTapesSpectraS3Request {
    return &GetSuspectBlobTapesSpectraS3Request{
    }
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithBlobId(blobId string) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.BlobId = &blobId
    return getSuspectBlobTapesSpectraS3Request
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithLastPage() *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.LastPage = true
    return getSuspectBlobTapesSpectraS3Request
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.PageLength = &pageLength
    return getSuspectBlobTapesSpectraS3Request
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.PageOffset = &pageOffset
    return getSuspectBlobTapesSpectraS3Request
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectBlobTapesSpectraS3Request
}

func (getSuspectBlobTapesSpectraS3Request *GetSuspectBlobTapesSpectraS3Request) WithTapeId(tapeId string) *GetSuspectBlobTapesSpectraS3Request {
    getSuspectBlobTapesSpectraS3Request.TapeId = &tapeId
    return getSuspectBlobTapesSpectraS3Request
}

type GetSuspectBucketsSpectraS3Request struct {
    DataPolicyId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetSuspectBucketsSpectraS3Request() *GetSuspectBucketsSpectraS3Request {
    return &GetSuspectBucketsSpectraS3Request{
    }
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.DataPolicyId = &dataPolicyId
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithLastPage() *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.LastPage = true
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithName(name string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.Name = &name
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.PageLength = &pageLength
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.PageOffset = &pageOffset
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectBucketsSpectraS3Request
}

func (getSuspectBucketsSpectraS3Request *GetSuspectBucketsSpectraS3Request) WithUserId(userId string) *GetSuspectBucketsSpectraS3Request {
    getSuspectBucketsSpectraS3Request.UserId = &userId
    return getSuspectBucketsSpectraS3Request
}

type GetSuspectObjectsSpectraS3Request struct {
    BucketId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetSuspectObjectsSpectraS3Request() *GetSuspectObjectsSpectraS3Request {
    return &GetSuspectObjectsSpectraS3Request{
    }
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithBucketId(bucketId string) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.BucketId = &bucketId
    return getSuspectObjectsSpectraS3Request
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithLastPage() *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.LastPage = true
    return getSuspectObjectsSpectraS3Request
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithPageLength(pageLength int) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.PageLength = &pageLength
    return getSuspectObjectsSpectraS3Request
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithPageOffset(pageOffset int) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.PageOffset = &pageOffset
    return getSuspectObjectsSpectraS3Request
}

func (getSuspectObjectsSpectraS3Request *GetSuspectObjectsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSuspectObjectsSpectraS3Request {
    getSuspectObjectsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSuspectObjectsSpectraS3Request
}

type GetSuspectObjectsWithFullDetailsSpectraS3Request struct {
    BucketId *string
    StorageDomain *string
}

func NewGetSuspectObjectsWithFullDetailsSpectraS3Request() *GetSuspectObjectsWithFullDetailsSpectraS3Request {
    return &GetSuspectObjectsWithFullDetailsSpectraS3Request{
    }
}

func (getSuspectObjectsWithFullDetailsSpectraS3Request *GetSuspectObjectsWithFullDetailsSpectraS3Request) WithBucketId(bucketId string) *GetSuspectObjectsWithFullDetailsSpectraS3Request {
    getSuspectObjectsWithFullDetailsSpectraS3Request.BucketId = &bucketId
    return getSuspectObjectsWithFullDetailsSpectraS3Request
}

func (getSuspectObjectsWithFullDetailsSpectraS3Request *GetSuspectObjectsWithFullDetailsSpectraS3Request) WithStorageDomain(storageDomain string) *GetSuspectObjectsWithFullDetailsSpectraS3Request {
    getSuspectObjectsWithFullDetailsSpectraS3Request.StorageDomain = &storageDomain
    return getSuspectObjectsWithFullDetailsSpectraS3Request
}

type MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewMarkSuspectBlobAzureTargetsAsDegradedSpectraS3Request(ids []string) *MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Request {
    return &MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Request{
        Ids: ids,
    }
}

func (markSuspectBlobAzureTargetsAsDegradedSpectraS3Request *MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Request) WithForce() *MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Request {
    markSuspectBlobAzureTargetsAsDegradedSpectraS3Request.Force = true
    return markSuspectBlobAzureTargetsAsDegradedSpectraS3Request
}

type MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewMarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request(ids []string) *MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request {
    return &MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request{
        Ids: ids,
    }
}

func (markSuspectBlobDs3TargetsAsDegradedSpectraS3Request *MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request) WithForce() *MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Request {
    markSuspectBlobDs3TargetsAsDegradedSpectraS3Request.Force = true
    return markSuspectBlobDs3TargetsAsDegradedSpectraS3Request
}

type MarkSuspectBlobPoolsAsDegradedSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewMarkSuspectBlobPoolsAsDegradedSpectraS3Request(ids []string) *MarkSuspectBlobPoolsAsDegradedSpectraS3Request {
    return &MarkSuspectBlobPoolsAsDegradedSpectraS3Request{
        Ids: ids,
    }
}

func (markSuspectBlobPoolsAsDegradedSpectraS3Request *MarkSuspectBlobPoolsAsDegradedSpectraS3Request) WithForce() *MarkSuspectBlobPoolsAsDegradedSpectraS3Request {
    markSuspectBlobPoolsAsDegradedSpectraS3Request.Force = true
    return markSuspectBlobPoolsAsDegradedSpectraS3Request
}

type MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewMarkSuspectBlobS3TargetsAsDegradedSpectraS3Request(ids []string) *MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request {
    return &MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request{
        Ids: ids,
    }
}

func (markSuspectBlobS3TargetsAsDegradedSpectraS3Request *MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request) WithForce() *MarkSuspectBlobS3TargetsAsDegradedSpectraS3Request {
    markSuspectBlobS3TargetsAsDegradedSpectraS3Request.Force = true
    return markSuspectBlobS3TargetsAsDegradedSpectraS3Request
}

type MarkSuspectBlobTapesAsDegradedSpectraS3Request struct {
    Force bool
    Ids []string
}

func NewMarkSuspectBlobTapesAsDegradedSpectraS3Request(ids []string) *MarkSuspectBlobTapesAsDegradedSpectraS3Request {
    return &MarkSuspectBlobTapesAsDegradedSpectraS3Request{
        Ids: ids,
    }
}

func (markSuspectBlobTapesAsDegradedSpectraS3Request *MarkSuspectBlobTapesAsDegradedSpectraS3Request) WithForce() *MarkSuspectBlobTapesAsDegradedSpectraS3Request {
    markSuspectBlobTapesAsDegradedSpectraS3Request.Force = true
    return markSuspectBlobTapesAsDegradedSpectraS3Request
}

type PutGroupGroupMemberSpectraS3Request struct {
    GroupId string
    MemberGroupId string
}

func NewPutGroupGroupMemberSpectraS3Request(groupId string, memberGroupId string) *PutGroupGroupMemberSpectraS3Request {
    return &PutGroupGroupMemberSpectraS3Request{
        GroupId: groupId,
        MemberGroupId: memberGroupId,
    }
}

type PutGroupSpectraS3Request struct {
    Name string
}

func NewPutGroupSpectraS3Request(name string) *PutGroupSpectraS3Request {
    return &PutGroupSpectraS3Request{
        Name: name,
    }
}

type PutUserGroupMemberSpectraS3Request struct {
    GroupId string
    MemberUserId string
}

func NewPutUserGroupMemberSpectraS3Request(groupId string, memberUserId string) *PutUserGroupMemberSpectraS3Request {
    return &PutUserGroupMemberSpectraS3Request{
        GroupId: groupId,
        MemberUserId: memberUserId,
    }
}

type DeleteGroupMemberSpectraS3Request struct {
    GroupMember string
}

func NewDeleteGroupMemberSpectraS3Request(groupMember string) *DeleteGroupMemberSpectraS3Request {
    return &DeleteGroupMemberSpectraS3Request{
        GroupMember: groupMember,
    }
}

type DeleteGroupSpectraS3Request struct {
    Group string
}

func NewDeleteGroupSpectraS3Request(group string) *DeleteGroupSpectraS3Request {
    return &DeleteGroupSpectraS3Request{
        Group: group,
    }
}

type GetGroupMemberSpectraS3Request struct {
    GroupMember string
}

func NewGetGroupMemberSpectraS3Request(groupMember string) *GetGroupMemberSpectraS3Request {
    return &GetGroupMemberSpectraS3Request{
        GroupMember: groupMember,
    }
}

type GetGroupMembersSpectraS3Request struct {
    GroupId *string
    LastPage bool
    MemberGroupId *string
    MemberUserId *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetGroupMembersSpectraS3Request() *GetGroupMembersSpectraS3Request {
    return &GetGroupMembersSpectraS3Request{
    }
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithGroupId(groupId string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.GroupId = &groupId
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithLastPage() *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.LastPage = true
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithMemberGroupId(memberGroupId string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.MemberGroupId = &memberGroupId
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithMemberUserId(memberUserId string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.MemberUserId = &memberUserId
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithPageLength(pageLength int) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.PageLength = &pageLength
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithPageOffset(pageOffset int) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.PageOffset = &pageOffset
    return getGroupMembersSpectraS3Request
}

func (getGroupMembersSpectraS3Request *GetGroupMembersSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetGroupMembersSpectraS3Request {
    getGroupMembersSpectraS3Request.PageStartMarker = &pageStartMarker
    return getGroupMembersSpectraS3Request
}

type GetGroupSpectraS3Request struct {
    Group string
}

func NewGetGroupSpectraS3Request(group string) *GetGroupSpectraS3Request {
    return &GetGroupSpectraS3Request{
        Group: group,
    }
}

type GetGroupsSpectraS3Request struct {
    BuiltIn *bool
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetGroupsSpectraS3Request() *GetGroupsSpectraS3Request {
    return &GetGroupsSpectraS3Request{
    }
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithBuiltIn(builtIn bool) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.BuiltIn = &builtIn
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithLastPage() *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.LastPage = true
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithName(name string) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.Name = &name
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithPageLength(pageLength int) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.PageLength = &pageLength
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithPageOffset(pageOffset int) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.PageOffset = &pageOffset
    return getGroupsSpectraS3Request
}

func (getGroupsSpectraS3Request *GetGroupsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetGroupsSpectraS3Request {
    getGroupsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getGroupsSpectraS3Request
}

type ModifyGroupSpectraS3Request struct {
    Group string
    Name *string
}

func NewModifyGroupSpectraS3Request(group string) *ModifyGroupSpectraS3Request {
    return &ModifyGroupSpectraS3Request{
        Group: group,
    }
}

func (modifyGroupSpectraS3Request *ModifyGroupSpectraS3Request) WithName(name string) *ModifyGroupSpectraS3Request {
    modifyGroupSpectraS3Request.Name = &name
    return modifyGroupSpectraS3Request
}

type VerifyUserIsMemberOfGroupSpectraS3Request struct {
    Group string
    UserId *string
}

func NewVerifyUserIsMemberOfGroupSpectraS3Request(group string) *VerifyUserIsMemberOfGroupSpectraS3Request {
    return &VerifyUserIsMemberOfGroupSpectraS3Request{
        Group: group,
    }
}

func (verifyUserIsMemberOfGroupSpectraS3Request *VerifyUserIsMemberOfGroupSpectraS3Request) WithUserId(userId string) *VerifyUserIsMemberOfGroupSpectraS3Request {
    verifyUserIsMemberOfGroupSpectraS3Request.UserId = &userId
    return verifyUserIsMemberOfGroupSpectraS3Request
}

type AllocateJobChunkSpectraS3Request struct {
    JobChunkId string
}

func NewAllocateJobChunkSpectraS3Request(jobChunkId string) *AllocateJobChunkSpectraS3Request {
    return &AllocateJobChunkSpectraS3Request{
        JobChunkId: jobChunkId,
    }
}

type CancelActiveJobSpectraS3Request struct {
    ActiveJobId string
}

func NewCancelActiveJobSpectraS3Request(activeJobId string) *CancelActiveJobSpectraS3Request {
    return &CancelActiveJobSpectraS3Request{
        ActiveJobId: activeJobId,
    }
}

type CancelAllActiveJobsSpectraS3Request struct {
    BucketId *string
    RequestType JobRequestType
}

func NewCancelAllActiveJobsSpectraS3Request() *CancelAllActiveJobsSpectraS3Request {
    return &CancelAllActiveJobsSpectraS3Request{
    }
}

func (cancelAllActiveJobsSpectraS3Request *CancelAllActiveJobsSpectraS3Request) WithBucketId(bucketId string) *CancelAllActiveJobsSpectraS3Request {
    cancelAllActiveJobsSpectraS3Request.BucketId = &bucketId
    return cancelAllActiveJobsSpectraS3Request
}

func (cancelAllActiveJobsSpectraS3Request *CancelAllActiveJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *CancelAllActiveJobsSpectraS3Request {
    cancelAllActiveJobsSpectraS3Request.RequestType = requestType
    return cancelAllActiveJobsSpectraS3Request
}

type CancelAllJobsSpectraS3Request struct {
    BucketId *string
    RequestType JobRequestType
}

func NewCancelAllJobsSpectraS3Request() *CancelAllJobsSpectraS3Request {
    return &CancelAllJobsSpectraS3Request{
    }
}

func (cancelAllJobsSpectraS3Request *CancelAllJobsSpectraS3Request) WithBucketId(bucketId string) *CancelAllJobsSpectraS3Request {
    cancelAllJobsSpectraS3Request.BucketId = &bucketId
    return cancelAllJobsSpectraS3Request
}

func (cancelAllJobsSpectraS3Request *CancelAllJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *CancelAllJobsSpectraS3Request {
    cancelAllJobsSpectraS3Request.RequestType = requestType
    return cancelAllJobsSpectraS3Request
}

type CancelJobSpectraS3Request struct {
    JobId string
}

func NewCancelJobSpectraS3Request(jobId string) *CancelJobSpectraS3Request {
    return &CancelJobSpectraS3Request{
        JobId: jobId,
    }
}

type ClearAllCanceledJobsSpectraS3Request struct {
}

func NewClearAllCanceledJobsSpectraS3Request() *ClearAllCanceledJobsSpectraS3Request {
    return &ClearAllCanceledJobsSpectraS3Request{
    }
}

type ClearAllCompletedJobsSpectraS3Request struct {
}

func NewClearAllCompletedJobsSpectraS3Request() *ClearAllCompletedJobsSpectraS3Request {
    return &ClearAllCompletedJobsSpectraS3Request{
    }
}

type CloseAggregatingJobSpectraS3Request struct {
    JobId string
}

func NewCloseAggregatingJobSpectraS3Request(jobId string) *CloseAggregatingJobSpectraS3Request {
    return &CloseAggregatingJobSpectraS3Request{
        JobId: jobId,
    }
}

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

type PutBulkJobSpectraS3Request struct {
    BucketName string
    Aggregating *bool
    Force bool
    IgnoreNamingConflicts bool
    ImplicitJobIdResolution *bool
    MaxUploadSize *int64
    MinimizeSpanningAcrossMedia *bool
    Name *string
    Objects []Ds3PutObject
    PreAllocateJobSpace bool
    Priority Priority
    VerifyAfterWrite *bool
}

func NewPutBulkJobSpectraS3Request(bucketName string, objects []Ds3PutObject) *PutBulkJobSpectraS3Request {
    return &PutBulkJobSpectraS3Request{
        BucketName: bucketName,
        Objects: objects,
    }
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithAggregating(aggregating bool) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.Aggregating = &aggregating
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

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithName(name string) *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.Name = &name
    return putBulkJobSpectraS3Request
}

func (putBulkJobSpectraS3Request *PutBulkJobSpectraS3Request) WithPreAllocateJobSpace() *PutBulkJobSpectraS3Request {
    putBulkJobSpectraS3Request.PreAllocateJobSpace = true
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

type GetActiveJobSpectraS3Request struct {
    ActiveJobId string
}

func NewGetActiveJobSpectraS3Request(activeJobId string) *GetActiveJobSpectraS3Request {
    return &GetActiveJobSpectraS3Request{
        ActiveJobId: activeJobId,
    }
}

type GetActiveJobsSpectraS3Request struct {
    Aggregating *bool
    BucketId *string
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    Priority Priority
    Rechunked *string
    RequestType JobRequestType
    Truncated *bool
    UserId *string
}

func NewGetActiveJobsSpectraS3Request() *GetActiveJobsSpectraS3Request {
    return &GetActiveJobsSpectraS3Request{
    }
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithAggregating(aggregating bool) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.Aggregating = &aggregating
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithBucketId(bucketId string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.BucketId = &bucketId
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.ChunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithLastPage() *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.LastPage = true
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithName(name string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.Name = &name
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPageLength(pageLength int) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.PageLength = &pageLength
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPageOffset(pageOffset int) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.PageOffset = &pageOffset
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithPriority(priority Priority) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.Priority = priority
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithRechunked(rechunked string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.Rechunked = &rechunked
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.RequestType = requestType
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithTruncated(truncated bool) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.Truncated = &truncated
    return getActiveJobsSpectraS3Request
}

func (getActiveJobsSpectraS3Request *GetActiveJobsSpectraS3Request) WithUserId(userId string) *GetActiveJobsSpectraS3Request {
    getActiveJobsSpectraS3Request.UserId = &userId
    return getActiveJobsSpectraS3Request
}

type GetCanceledJobSpectraS3Request struct {
    CanceledJob string
}

func NewGetCanceledJobSpectraS3Request(canceledJob string) *GetCanceledJobSpectraS3Request {
    return &GetCanceledJobSpectraS3Request{
        CanceledJob: canceledJob,
    }
}

type GetCanceledJobsSpectraS3Request struct {
    BucketId *string
    CanceledDueToTimeout *bool
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    Priority Priority
    Rechunked *string
    RequestType JobRequestType
    Truncated *bool
    UserId *string
}

func NewGetCanceledJobsSpectraS3Request() *GetCanceledJobsSpectraS3Request {
    return &GetCanceledJobsSpectraS3Request{
    }
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithBucketId(bucketId string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.BucketId = &bucketId
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithCanceledDueToTimeout(canceledDueToTimeout bool) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.CanceledDueToTimeout = &canceledDueToTimeout
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.ChunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithLastPage() *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.LastPage = true
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithName(name string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.Name = &name
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPageLength(pageLength int) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.PageLength = &pageLength
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPageOffset(pageOffset int) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.PageOffset = &pageOffset
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithPriority(priority Priority) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.Priority = priority
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithRechunked(rechunked string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.Rechunked = &rechunked
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.RequestType = requestType
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithTruncated(truncated bool) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.Truncated = &truncated
    return getCanceledJobsSpectraS3Request
}

func (getCanceledJobsSpectraS3Request *GetCanceledJobsSpectraS3Request) WithUserId(userId string) *GetCanceledJobsSpectraS3Request {
    getCanceledJobsSpectraS3Request.UserId = &userId
    return getCanceledJobsSpectraS3Request
}

type GetCompletedJobSpectraS3Request struct {
    CompletedJob string
}

func NewGetCompletedJobSpectraS3Request(completedJob string) *GetCompletedJobSpectraS3Request {
    return &GetCompletedJobSpectraS3Request{
        CompletedJob: completedJob,
    }
}

type GetCompletedJobsSpectraS3Request struct {
    BucketId *string
    ChunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    Priority Priority
    Rechunked *string
    RequestType JobRequestType
    Truncated *bool
    UserId *string
}

func NewGetCompletedJobsSpectraS3Request() *GetCompletedJobsSpectraS3Request {
    return &GetCompletedJobsSpectraS3Request{
    }
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithBucketId(bucketId string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.BucketId = &bucketId
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithChunkClientProcessingOrderGuarantee(chunkClientProcessingOrderGuarantee JobChunkClientProcessingOrderGuarantee) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.ChunkClientProcessingOrderGuarantee = chunkClientProcessingOrderGuarantee
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithLastPage() *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.LastPage = true
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithName(name string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.Name = &name
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPageLength(pageLength int) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.PageLength = &pageLength
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPageOffset(pageOffset int) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.PageOffset = &pageOffset
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithPriority(priority Priority) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.Priority = priority
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithRechunked(rechunked string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.Rechunked = &rechunked
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.RequestType = requestType
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithTruncated(truncated bool) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.Truncated = &truncated
    return getCompletedJobsSpectraS3Request
}

func (getCompletedJobsSpectraS3Request *GetCompletedJobsSpectraS3Request) WithUserId(userId string) *GetCompletedJobsSpectraS3Request {
    getCompletedJobsSpectraS3Request.UserId = &userId
    return getCompletedJobsSpectraS3Request
}

type GetJobChunkDaoSpectraS3Request struct {
    JobChunkDao string
}

func NewGetJobChunkDaoSpectraS3Request(jobChunkDao string) *GetJobChunkDaoSpectraS3Request {
    return &GetJobChunkDaoSpectraS3Request{
        JobChunkDao: jobChunkDao,
    }
}

type GetJobChunkSpectraS3Request struct {
    JobChunkId string
}

func NewGetJobChunkSpectraS3Request(jobChunkId string) *GetJobChunkSpectraS3Request {
    return &GetJobChunkSpectraS3Request{
        JobChunkId: jobChunkId,
    }
}

type GetJobChunksReadyForClientProcessingSpectraS3Request struct {
    Job string
    JobChunk *string
    PreferredNumberOfChunks *int
}

func NewGetJobChunksReadyForClientProcessingSpectraS3Request(job string) *GetJobChunksReadyForClientProcessingSpectraS3Request {
    return &GetJobChunksReadyForClientProcessingSpectraS3Request{
        Job: job,
    }
}

func (getJobChunksReadyForClientProcessingSpectraS3Request *GetJobChunksReadyForClientProcessingSpectraS3Request) WithJobChunk(jobChunk string) *GetJobChunksReadyForClientProcessingSpectraS3Request {
    getJobChunksReadyForClientProcessingSpectraS3Request.JobChunk = &jobChunk
    return getJobChunksReadyForClientProcessingSpectraS3Request
}

func (getJobChunksReadyForClientProcessingSpectraS3Request *GetJobChunksReadyForClientProcessingSpectraS3Request) WithPreferredNumberOfChunks(preferredNumberOfChunks int) *GetJobChunksReadyForClientProcessingSpectraS3Request {
    getJobChunksReadyForClientProcessingSpectraS3Request.PreferredNumberOfChunks = &preferredNumberOfChunks
    return getJobChunksReadyForClientProcessingSpectraS3Request
}

type GetJobSpectraS3Request struct {
    JobId string
}

func NewGetJobSpectraS3Request(jobId string) *GetJobSpectraS3Request {
    return &GetJobSpectraS3Request{
        JobId: jobId,
    }
}

type GetJobToReplicateSpectraS3Request struct {
    JobId string
}

func NewGetJobToReplicateSpectraS3Request(jobId string) *GetJobToReplicateSpectraS3Request {
    return &GetJobToReplicateSpectraS3Request{
        JobId: jobId,
    }
}

type GetJobsSpectraS3Request struct {
    BucketId *string
    FullDetails bool
}

func NewGetJobsSpectraS3Request() *GetJobsSpectraS3Request {
    return &GetJobsSpectraS3Request{
    }
}

func (getJobsSpectraS3Request *GetJobsSpectraS3Request) WithBucketId(bucketId string) *GetJobsSpectraS3Request {
    getJobsSpectraS3Request.BucketId = &bucketId
    return getJobsSpectraS3Request
}

func (getJobsSpectraS3Request *GetJobsSpectraS3Request) WithFullDetails() *GetJobsSpectraS3Request {
    getJobsSpectraS3Request.FullDetails = true
    return getJobsSpectraS3Request
}

type ModifyActiveJobSpectraS3Request struct {
    ActiveJobId string
    CreatedAt *string
    Name *string
    Priority Priority
}

func NewModifyActiveJobSpectraS3Request(activeJobId string) *ModifyActiveJobSpectraS3Request {
    return &ModifyActiveJobSpectraS3Request{
        ActiveJobId: activeJobId,
    }
}

func (modifyActiveJobSpectraS3Request *ModifyActiveJobSpectraS3Request) WithCreatedAt(createdAt string) *ModifyActiveJobSpectraS3Request {
    modifyActiveJobSpectraS3Request.CreatedAt = &createdAt
    return modifyActiveJobSpectraS3Request
}

func (modifyActiveJobSpectraS3Request *ModifyActiveJobSpectraS3Request) WithName(name string) *ModifyActiveJobSpectraS3Request {
    modifyActiveJobSpectraS3Request.Name = &name
    return modifyActiveJobSpectraS3Request
}

func (modifyActiveJobSpectraS3Request *ModifyActiveJobSpectraS3Request) WithPriority(priority Priority) *ModifyActiveJobSpectraS3Request {
    modifyActiveJobSpectraS3Request.Priority = priority
    return modifyActiveJobSpectraS3Request
}

type ModifyJobSpectraS3Request struct {
    CreatedAt *string
    JobId string
    Name *string
    Priority Priority
}

func NewModifyJobSpectraS3Request(jobId string) *ModifyJobSpectraS3Request {
    return &ModifyJobSpectraS3Request{
        JobId: jobId,
    }
}

func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) WithCreatedAt(createdAt string) *ModifyJobSpectraS3Request {
    modifyJobSpectraS3Request.CreatedAt = &createdAt
    return modifyJobSpectraS3Request
}

func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) WithName(name string) *ModifyJobSpectraS3Request {
    modifyJobSpectraS3Request.Name = &name
    return modifyJobSpectraS3Request
}

func (modifyJobSpectraS3Request *ModifyJobSpectraS3Request) WithPriority(priority Priority) *ModifyJobSpectraS3Request {
    modifyJobSpectraS3Request.Priority = priority
    return modifyJobSpectraS3Request
}

type ReplicatePutJobSpectraS3Request struct {
    BucketName string
    Priority Priority
    RequestPayload string
}

func NewReplicatePutJobSpectraS3Request(bucketName string, requestPayload string) *ReplicatePutJobSpectraS3Request {
    return &ReplicatePutJobSpectraS3Request{
        BucketName: bucketName,
        RequestPayload: requestPayload,
    }
}

func (replicatePutJobSpectraS3Request *ReplicatePutJobSpectraS3Request) WithPriority(priority Priority) *ReplicatePutJobSpectraS3Request {
    replicatePutJobSpectraS3Request.Priority = priority
    return replicatePutJobSpectraS3Request
}

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

type TruncateActiveJobSpectraS3Request struct {
    ActiveJobId string
}

func NewTruncateActiveJobSpectraS3Request(activeJobId string) *TruncateActiveJobSpectraS3Request {
    return &TruncateActiveJobSpectraS3Request{
        ActiveJobId: activeJobId,
    }
}

type TruncateAllActiveJobsSpectraS3Request struct {
    BucketId *string
    RequestType JobRequestType
}

func NewTruncateAllActiveJobsSpectraS3Request() *TruncateAllActiveJobsSpectraS3Request {
    return &TruncateAllActiveJobsSpectraS3Request{
    }
}

func (truncateAllActiveJobsSpectraS3Request *TruncateAllActiveJobsSpectraS3Request) WithBucketId(bucketId string) *TruncateAllActiveJobsSpectraS3Request {
    truncateAllActiveJobsSpectraS3Request.BucketId = &bucketId
    return truncateAllActiveJobsSpectraS3Request
}

func (truncateAllActiveJobsSpectraS3Request *TruncateAllActiveJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *TruncateAllActiveJobsSpectraS3Request {
    truncateAllActiveJobsSpectraS3Request.RequestType = requestType
    return truncateAllActiveJobsSpectraS3Request
}

type TruncateAllJobsSpectraS3Request struct {
    BucketId *string
    RequestType JobRequestType
}

func NewTruncateAllJobsSpectraS3Request() *TruncateAllJobsSpectraS3Request {
    return &TruncateAllJobsSpectraS3Request{
    }
}

func (truncateAllJobsSpectraS3Request *TruncateAllJobsSpectraS3Request) WithBucketId(bucketId string) *TruncateAllJobsSpectraS3Request {
    truncateAllJobsSpectraS3Request.BucketId = &bucketId
    return truncateAllJobsSpectraS3Request
}

func (truncateAllJobsSpectraS3Request *TruncateAllJobsSpectraS3Request) WithRequestType(requestType JobRequestType) *TruncateAllJobsSpectraS3Request {
    truncateAllJobsSpectraS3Request.RequestType = requestType
    return truncateAllJobsSpectraS3Request
}

type TruncateJobSpectraS3Request struct {
    JobId string
}

func NewTruncateJobSpectraS3Request(jobId string) *TruncateJobSpectraS3Request {
    return &TruncateJobSpectraS3Request{
        JobId: jobId,
    }
}

type VerifySafeToCreatePutJobSpectraS3Request struct {
    BucketName string
}

func NewVerifySafeToCreatePutJobSpectraS3Request(bucketName string) *VerifySafeToCreatePutJobSpectraS3Request {
    return &VerifySafeToCreatePutJobSpectraS3Request{
        BucketName: bucketName,
    }
}

type GetNodeSpectraS3Request struct {
    Node string
}

func NewGetNodeSpectraS3Request(node string) *GetNodeSpectraS3Request {
    return &GetNodeSpectraS3Request{
        Node: node,
    }
}

type GetNodesSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetNodesSpectraS3Request() *GetNodesSpectraS3Request {
    return &GetNodesSpectraS3Request{
    }
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithLastPage() *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.LastPage = true
    return getNodesSpectraS3Request
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithPageLength(pageLength int) *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.PageLength = &pageLength
    return getNodesSpectraS3Request
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithPageOffset(pageOffset int) *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.PageOffset = &pageOffset
    return getNodesSpectraS3Request
}

func (getNodesSpectraS3Request *GetNodesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetNodesSpectraS3Request {
    getNodesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getNodesSpectraS3Request
}

type ModifyNodeSpectraS3Request struct {
    DnsName *string
    Name *string
    Node string
}

func NewModifyNodeSpectraS3Request(node string) *ModifyNodeSpectraS3Request {
    return &ModifyNodeSpectraS3Request{
        Node: node,
    }
}

func (modifyNodeSpectraS3Request *ModifyNodeSpectraS3Request) WithDnsName(dnsName string) *ModifyNodeSpectraS3Request {
    modifyNodeSpectraS3Request.DnsName = &dnsName
    return modifyNodeSpectraS3Request
}

func (modifyNodeSpectraS3Request *ModifyNodeSpectraS3Request) WithName(name string) *ModifyNodeSpectraS3Request {
    modifyNodeSpectraS3Request.Name = &name
    return modifyNodeSpectraS3Request
}

type PutAzureTargetFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutAzureTargetFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    return &PutAzureTargetFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.Format = format
    return putAzureTargetFailureNotificationRegistrationSpectraS3Request
}

func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putAzureTargetFailureNotificationRegistrationSpectraS3Request
}

func (putAzureTargetFailureNotificationRegistrationSpectraS3Request *PutAzureTargetFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutAzureTargetFailureNotificationRegistrationSpectraS3Request {
    putAzureTargetFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putAzureTargetFailureNotificationRegistrationSpectraS3Request
}

type PutBucketChangesNotificationRegistrationSpectraS3Request struct {
    BucketId *string
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutBucketChangesNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutBucketChangesNotificationRegistrationSpectraS3Request {
    return &PutBucketChangesNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putBucketChangesNotificationRegistrationSpectraS3Request *PutBucketChangesNotificationRegistrationSpectraS3Request) WithBucketId(bucketId string) *PutBucketChangesNotificationRegistrationSpectraS3Request {
    putBucketChangesNotificationRegistrationSpectraS3Request.BucketId = &bucketId
    return putBucketChangesNotificationRegistrationSpectraS3Request
}

func (putBucketChangesNotificationRegistrationSpectraS3Request *PutBucketChangesNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutBucketChangesNotificationRegistrationSpectraS3Request {
    putBucketChangesNotificationRegistrationSpectraS3Request.Format = format
    return putBucketChangesNotificationRegistrationSpectraS3Request
}

func (putBucketChangesNotificationRegistrationSpectraS3Request *PutBucketChangesNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutBucketChangesNotificationRegistrationSpectraS3Request {
    putBucketChangesNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putBucketChangesNotificationRegistrationSpectraS3Request
}

func (putBucketChangesNotificationRegistrationSpectraS3Request *PutBucketChangesNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutBucketChangesNotificationRegistrationSpectraS3Request {
    putBucketChangesNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putBucketChangesNotificationRegistrationSpectraS3Request
}

type PutDs3TargetFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutDs3TargetFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutDs3TargetFailureNotificationRegistrationSpectraS3Request {
    return &PutDs3TargetFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putDs3TargetFailureNotificationRegistrationSpectraS3Request *PutDs3TargetFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutDs3TargetFailureNotificationRegistrationSpectraS3Request {
    putDs3TargetFailureNotificationRegistrationSpectraS3Request.Format = format
    return putDs3TargetFailureNotificationRegistrationSpectraS3Request
}

func (putDs3TargetFailureNotificationRegistrationSpectraS3Request *PutDs3TargetFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutDs3TargetFailureNotificationRegistrationSpectraS3Request {
    putDs3TargetFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putDs3TargetFailureNotificationRegistrationSpectraS3Request
}

func (putDs3TargetFailureNotificationRegistrationSpectraS3Request *PutDs3TargetFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutDs3TargetFailureNotificationRegistrationSpectraS3Request {
    putDs3TargetFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putDs3TargetFailureNotificationRegistrationSpectraS3Request
}

type PutJobCompletedNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    JobId *string
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutJobCompletedNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutJobCompletedNotificationRegistrationSpectraS3Request {
    return &PutJobCompletedNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putJobCompletedNotificationRegistrationSpectraS3Request *PutJobCompletedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutJobCompletedNotificationRegistrationSpectraS3Request {
    putJobCompletedNotificationRegistrationSpectraS3Request.Format = format
    return putJobCompletedNotificationRegistrationSpectraS3Request
}

func (putJobCompletedNotificationRegistrationSpectraS3Request *PutJobCompletedNotificationRegistrationSpectraS3Request) WithJobId(jobId string) *PutJobCompletedNotificationRegistrationSpectraS3Request {
    putJobCompletedNotificationRegistrationSpectraS3Request.JobId = &jobId
    return putJobCompletedNotificationRegistrationSpectraS3Request
}

func (putJobCompletedNotificationRegistrationSpectraS3Request *PutJobCompletedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutJobCompletedNotificationRegistrationSpectraS3Request {
    putJobCompletedNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putJobCompletedNotificationRegistrationSpectraS3Request
}

func (putJobCompletedNotificationRegistrationSpectraS3Request *PutJobCompletedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutJobCompletedNotificationRegistrationSpectraS3Request {
    putJobCompletedNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putJobCompletedNotificationRegistrationSpectraS3Request
}

type PutJobCreatedNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutJobCreatedNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutJobCreatedNotificationRegistrationSpectraS3Request {
    return &PutJobCreatedNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putJobCreatedNotificationRegistrationSpectraS3Request *PutJobCreatedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutJobCreatedNotificationRegistrationSpectraS3Request {
    putJobCreatedNotificationRegistrationSpectraS3Request.Format = format
    return putJobCreatedNotificationRegistrationSpectraS3Request
}

func (putJobCreatedNotificationRegistrationSpectraS3Request *PutJobCreatedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutJobCreatedNotificationRegistrationSpectraS3Request {
    putJobCreatedNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putJobCreatedNotificationRegistrationSpectraS3Request
}

func (putJobCreatedNotificationRegistrationSpectraS3Request *PutJobCreatedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutJobCreatedNotificationRegistrationSpectraS3Request {
    putJobCreatedNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putJobCreatedNotificationRegistrationSpectraS3Request
}

type PutJobCreationFailedNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutJobCreationFailedNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    return &PutJobCreationFailedNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    putJobCreationFailedNotificationRegistrationSpectraS3Request.Format = format
    return putJobCreationFailedNotificationRegistrationSpectraS3Request
}

func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    putJobCreationFailedNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putJobCreationFailedNotificationRegistrationSpectraS3Request
}

func (putJobCreationFailedNotificationRegistrationSpectraS3Request *PutJobCreationFailedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutJobCreationFailedNotificationRegistrationSpectraS3Request {
    putJobCreationFailedNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putJobCreationFailedNotificationRegistrationSpectraS3Request
}

type PutObjectCachedNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    JobId *string
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutObjectCachedNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutObjectCachedNotificationRegistrationSpectraS3Request {
    return &PutObjectCachedNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putObjectCachedNotificationRegistrationSpectraS3Request *PutObjectCachedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutObjectCachedNotificationRegistrationSpectraS3Request {
    putObjectCachedNotificationRegistrationSpectraS3Request.Format = format
    return putObjectCachedNotificationRegistrationSpectraS3Request
}

func (putObjectCachedNotificationRegistrationSpectraS3Request *PutObjectCachedNotificationRegistrationSpectraS3Request) WithJobId(jobId string) *PutObjectCachedNotificationRegistrationSpectraS3Request {
    putObjectCachedNotificationRegistrationSpectraS3Request.JobId = &jobId
    return putObjectCachedNotificationRegistrationSpectraS3Request
}

func (putObjectCachedNotificationRegistrationSpectraS3Request *PutObjectCachedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutObjectCachedNotificationRegistrationSpectraS3Request {
    putObjectCachedNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putObjectCachedNotificationRegistrationSpectraS3Request
}

func (putObjectCachedNotificationRegistrationSpectraS3Request *PutObjectCachedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutObjectCachedNotificationRegistrationSpectraS3Request {
    putObjectCachedNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putObjectCachedNotificationRegistrationSpectraS3Request
}

type PutObjectLostNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutObjectLostNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutObjectLostNotificationRegistrationSpectraS3Request {
    return &PutObjectLostNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutObjectLostNotificationRegistrationSpectraS3Request {
    putObjectLostNotificationRegistrationSpectraS3Request.Format = format
    return putObjectLostNotificationRegistrationSpectraS3Request
}

func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutObjectLostNotificationRegistrationSpectraS3Request {
    putObjectLostNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putObjectLostNotificationRegistrationSpectraS3Request
}

func (putObjectLostNotificationRegistrationSpectraS3Request *PutObjectLostNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutObjectLostNotificationRegistrationSpectraS3Request {
    putObjectLostNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putObjectLostNotificationRegistrationSpectraS3Request
}

type PutObjectPersistedNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    JobId *string
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutObjectPersistedNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    return &PutObjectPersistedNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.Format = format
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithJobId(jobId string) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.JobId = &jobId
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}

func (putObjectPersistedNotificationRegistrationSpectraS3Request *PutObjectPersistedNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutObjectPersistedNotificationRegistrationSpectraS3Request {
    putObjectPersistedNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putObjectPersistedNotificationRegistrationSpectraS3Request
}

type PutPoolFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutPoolFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    return &PutPoolFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    putPoolFailureNotificationRegistrationSpectraS3Request.Format = format
    return putPoolFailureNotificationRegistrationSpectraS3Request
}

func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    putPoolFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putPoolFailureNotificationRegistrationSpectraS3Request
}

func (putPoolFailureNotificationRegistrationSpectraS3Request *PutPoolFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutPoolFailureNotificationRegistrationSpectraS3Request {
    putPoolFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putPoolFailureNotificationRegistrationSpectraS3Request
}

type PutS3TargetFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutS3TargetFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutS3TargetFailureNotificationRegistrationSpectraS3Request {
    return &PutS3TargetFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putS3TargetFailureNotificationRegistrationSpectraS3Request *PutS3TargetFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutS3TargetFailureNotificationRegistrationSpectraS3Request {
    putS3TargetFailureNotificationRegistrationSpectraS3Request.Format = format
    return putS3TargetFailureNotificationRegistrationSpectraS3Request
}

func (putS3TargetFailureNotificationRegistrationSpectraS3Request *PutS3TargetFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutS3TargetFailureNotificationRegistrationSpectraS3Request {
    putS3TargetFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putS3TargetFailureNotificationRegistrationSpectraS3Request
}

func (putS3TargetFailureNotificationRegistrationSpectraS3Request *PutS3TargetFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutS3TargetFailureNotificationRegistrationSpectraS3Request {
    putS3TargetFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putS3TargetFailureNotificationRegistrationSpectraS3Request
}

type PutStorageDomainFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutStorageDomainFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutStorageDomainFailureNotificationRegistrationSpectraS3Request {
    return &PutStorageDomainFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putStorageDomainFailureNotificationRegistrationSpectraS3Request *PutStorageDomainFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutStorageDomainFailureNotificationRegistrationSpectraS3Request {
    putStorageDomainFailureNotificationRegistrationSpectraS3Request.Format = format
    return putStorageDomainFailureNotificationRegistrationSpectraS3Request
}

func (putStorageDomainFailureNotificationRegistrationSpectraS3Request *PutStorageDomainFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutStorageDomainFailureNotificationRegistrationSpectraS3Request {
    putStorageDomainFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putStorageDomainFailureNotificationRegistrationSpectraS3Request
}

func (putStorageDomainFailureNotificationRegistrationSpectraS3Request *PutStorageDomainFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutStorageDomainFailureNotificationRegistrationSpectraS3Request {
    putStorageDomainFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putStorageDomainFailureNotificationRegistrationSpectraS3Request
}

type PutSystemFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutSystemFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutSystemFailureNotificationRegistrationSpectraS3Request {
    return &PutSystemFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putSystemFailureNotificationRegistrationSpectraS3Request *PutSystemFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutSystemFailureNotificationRegistrationSpectraS3Request {
    putSystemFailureNotificationRegistrationSpectraS3Request.Format = format
    return putSystemFailureNotificationRegistrationSpectraS3Request
}

func (putSystemFailureNotificationRegistrationSpectraS3Request *PutSystemFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutSystemFailureNotificationRegistrationSpectraS3Request {
    putSystemFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putSystemFailureNotificationRegistrationSpectraS3Request
}

func (putSystemFailureNotificationRegistrationSpectraS3Request *PutSystemFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutSystemFailureNotificationRegistrationSpectraS3Request {
    putSystemFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putSystemFailureNotificationRegistrationSpectraS3Request
}

type PutTapeFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutTapeFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    return &PutTapeFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    putTapeFailureNotificationRegistrationSpectraS3Request.Format = format
    return putTapeFailureNotificationRegistrationSpectraS3Request
}

func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    putTapeFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putTapeFailureNotificationRegistrationSpectraS3Request
}

func (putTapeFailureNotificationRegistrationSpectraS3Request *PutTapeFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutTapeFailureNotificationRegistrationSpectraS3Request {
    putTapeFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putTapeFailureNotificationRegistrationSpectraS3Request
}

type PutTapePartitionFailureNotificationRegistrationSpectraS3Request struct {
    Format HttpResponseFormatType
    NamingConvention NamingConventionType
    NotificationEndPoint string
    NotificationHttpMethod RequestType
}

func NewPutTapePartitionFailureNotificationRegistrationSpectraS3Request(notificationEndPoint string) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    return &PutTapePartitionFailureNotificationRegistrationSpectraS3Request{
        NotificationEndPoint: notificationEndPoint,
    }
}

func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) WithFormat(format HttpResponseFormatType) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.Format = format
    return putTapePartitionFailureNotificationRegistrationSpectraS3Request
}

func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) WithNamingConvention(namingConvention NamingConventionType) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.NamingConvention = namingConvention
    return putTapePartitionFailureNotificationRegistrationSpectraS3Request
}

func (putTapePartitionFailureNotificationRegistrationSpectraS3Request *PutTapePartitionFailureNotificationRegistrationSpectraS3Request) WithNotificationHttpMethod(notificationHttpMethod RequestType) *PutTapePartitionFailureNotificationRegistrationSpectraS3Request {
    putTapePartitionFailureNotificationRegistrationSpectraS3Request.NotificationHttpMethod = notificationHttpMethod
    return putTapePartitionFailureNotificationRegistrationSpectraS3Request
}

type DeleteAzureTargetFailureNotificationRegistrationSpectraS3Request struct {
    AzureTargetFailureNotificationRegistration string
}

func NewDeleteAzureTargetFailureNotificationRegistrationSpectraS3Request(azureTargetFailureNotificationRegistration string) *DeleteAzureTargetFailureNotificationRegistrationSpectraS3Request {
    return &DeleteAzureTargetFailureNotificationRegistrationSpectraS3Request{
        AzureTargetFailureNotificationRegistration: azureTargetFailureNotificationRegistration,
    }
}

type DeleteBucketChangesNotificationRegistrationSpectraS3Request struct {
    BucketChangesNotificationRegistration string
}

func NewDeleteBucketChangesNotificationRegistrationSpectraS3Request(bucketChangesNotificationRegistration string) *DeleteBucketChangesNotificationRegistrationSpectraS3Request {
    return &DeleteBucketChangesNotificationRegistrationSpectraS3Request{
        BucketChangesNotificationRegistration: bucketChangesNotificationRegistration,
    }
}

type DeleteDs3TargetFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteDs3TargetFailureNotificationRegistrationSpectraS3Request(notificationId string) *DeleteDs3TargetFailureNotificationRegistrationSpectraS3Request {
    return &DeleteDs3TargetFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeleteJobCompletedNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteJobCompletedNotificationRegistrationSpectraS3Request(notificationId string) *DeleteJobCompletedNotificationRegistrationSpectraS3Request {
    return &DeleteJobCompletedNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeleteJobCreatedNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteJobCreatedNotificationRegistrationSpectraS3Request(notificationId string) *DeleteJobCreatedNotificationRegistrationSpectraS3Request {
    return &DeleteJobCreatedNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeleteJobCreationFailedNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteJobCreationFailedNotificationRegistrationSpectraS3Request(notificationId string) *DeleteJobCreationFailedNotificationRegistrationSpectraS3Request {
    return &DeleteJobCreationFailedNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeleteObjectCachedNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteObjectCachedNotificationRegistrationSpectraS3Request(notificationId string) *DeleteObjectCachedNotificationRegistrationSpectraS3Request {
    return &DeleteObjectCachedNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeleteObjectLostNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteObjectLostNotificationRegistrationSpectraS3Request(notificationId string) *DeleteObjectLostNotificationRegistrationSpectraS3Request {
    return &DeleteObjectLostNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeleteObjectPersistedNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteObjectPersistedNotificationRegistrationSpectraS3Request(notificationId string) *DeleteObjectPersistedNotificationRegistrationSpectraS3Request {
    return &DeleteObjectPersistedNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeletePoolFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeletePoolFailureNotificationRegistrationSpectraS3Request(notificationId string) *DeletePoolFailureNotificationRegistrationSpectraS3Request {
    return &DeletePoolFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeleteS3TargetFailureNotificationRegistrationSpectraS3Request struct {
    S3TargetFailureNotificationRegistration string
}

func NewDeleteS3TargetFailureNotificationRegistrationSpectraS3Request(s3TargetFailureNotificationRegistration string) *DeleteS3TargetFailureNotificationRegistrationSpectraS3Request {
    return &DeleteS3TargetFailureNotificationRegistrationSpectraS3Request{
        S3TargetFailureNotificationRegistration: s3TargetFailureNotificationRegistration,
    }
}

type DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteStorageDomainFailureNotificationRegistrationSpectraS3Request(notificationId string) *DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request {
    return &DeleteStorageDomainFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeleteSystemFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteSystemFailureNotificationRegistrationSpectraS3Request(notificationId string) *DeleteSystemFailureNotificationRegistrationSpectraS3Request {
    return &DeleteSystemFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeleteTapeFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteTapeFailureNotificationRegistrationSpectraS3Request(notificationId string) *DeleteTapeFailureNotificationRegistrationSpectraS3Request {
    return &DeleteTapeFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type DeleteTapePartitionFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewDeleteTapePartitionFailureNotificationRegistrationSpectraS3Request(notificationId string) *DeleteTapePartitionFailureNotificationRegistrationSpectraS3Request {
    return &DeleteTapePartitionFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetAzureTargetFailureNotificationRegistrationSpectraS3Request struct {
    AzureTargetFailureNotificationRegistration string
}

func NewGetAzureTargetFailureNotificationRegistrationSpectraS3Request(azureTargetFailureNotificationRegistration string) *GetAzureTargetFailureNotificationRegistrationSpectraS3Request {
    return &GetAzureTargetFailureNotificationRegistrationSpectraS3Request{
        AzureTargetFailureNotificationRegistration: azureTargetFailureNotificationRegistration,
    }
}

type GetAzureTargetFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetAzureTargetFailureNotificationRegistrationsSpectraS3Request() *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    return &GetAzureTargetFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Request *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetAzureTargetFailureNotificationRegistrationsSpectraS3Request {
    getAzureTargetFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getAzureTargetFailureNotificationRegistrationsSpectraS3Request
}

type GetBucketChangesNotificationRegistrationSpectraS3Request struct {
    BucketChangesNotificationRegistration string
}

func NewGetBucketChangesNotificationRegistrationSpectraS3Request(bucketChangesNotificationRegistration string) *GetBucketChangesNotificationRegistrationSpectraS3Request {
    return &GetBucketChangesNotificationRegistrationSpectraS3Request{
        BucketChangesNotificationRegistration: bucketChangesNotificationRegistration,
    }
}

type GetBucketChangesNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetBucketChangesNotificationRegistrationsSpectraS3Request() *GetBucketChangesNotificationRegistrationsSpectraS3Request {
    return &GetBucketChangesNotificationRegistrationsSpectraS3Request{
    }
}

func (getBucketChangesNotificationRegistrationsSpectraS3Request *GetBucketChangesNotificationRegistrationsSpectraS3Request) WithLastPage() *GetBucketChangesNotificationRegistrationsSpectraS3Request {
    getBucketChangesNotificationRegistrationsSpectraS3Request.LastPage = true
    return getBucketChangesNotificationRegistrationsSpectraS3Request
}

func (getBucketChangesNotificationRegistrationsSpectraS3Request *GetBucketChangesNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetBucketChangesNotificationRegistrationsSpectraS3Request {
    getBucketChangesNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getBucketChangesNotificationRegistrationsSpectraS3Request
}

func (getBucketChangesNotificationRegistrationsSpectraS3Request *GetBucketChangesNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetBucketChangesNotificationRegistrationsSpectraS3Request {
    getBucketChangesNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getBucketChangesNotificationRegistrationsSpectraS3Request
}

func (getBucketChangesNotificationRegistrationsSpectraS3Request *GetBucketChangesNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetBucketChangesNotificationRegistrationsSpectraS3Request {
    getBucketChangesNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getBucketChangesNotificationRegistrationsSpectraS3Request
}

func (getBucketChangesNotificationRegistrationsSpectraS3Request *GetBucketChangesNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetBucketChangesNotificationRegistrationsSpectraS3Request {
    getBucketChangesNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getBucketChangesNotificationRegistrationsSpectraS3Request
}

type GetBucketHistorySpectraS3Request struct {
    BucketId *string
    LastPage bool
    MinSequenceNumber *int64
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetBucketHistorySpectraS3Request() *GetBucketHistorySpectraS3Request {
    return &GetBucketHistorySpectraS3Request{
    }
}

func (getBucketHistorySpectraS3Request *GetBucketHistorySpectraS3Request) WithBucketId(bucketId string) *GetBucketHistorySpectraS3Request {
    getBucketHistorySpectraS3Request.BucketId = &bucketId
    return getBucketHistorySpectraS3Request
}

func (getBucketHistorySpectraS3Request *GetBucketHistorySpectraS3Request) WithLastPage() *GetBucketHistorySpectraS3Request {
    getBucketHistorySpectraS3Request.LastPage = true
    return getBucketHistorySpectraS3Request
}

func (getBucketHistorySpectraS3Request *GetBucketHistorySpectraS3Request) WithMinSequenceNumber(minSequenceNumber int64) *GetBucketHistorySpectraS3Request {
    getBucketHistorySpectraS3Request.MinSequenceNumber = &minSequenceNumber
    return getBucketHistorySpectraS3Request
}

func (getBucketHistorySpectraS3Request *GetBucketHistorySpectraS3Request) WithPageLength(pageLength int) *GetBucketHistorySpectraS3Request {
    getBucketHistorySpectraS3Request.PageLength = &pageLength
    return getBucketHistorySpectraS3Request
}

func (getBucketHistorySpectraS3Request *GetBucketHistorySpectraS3Request) WithPageOffset(pageOffset int) *GetBucketHistorySpectraS3Request {
    getBucketHistorySpectraS3Request.PageOffset = &pageOffset
    return getBucketHistorySpectraS3Request
}

func (getBucketHistorySpectraS3Request *GetBucketHistorySpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetBucketHistorySpectraS3Request {
    getBucketHistorySpectraS3Request.PageStartMarker = &pageStartMarker
    return getBucketHistorySpectraS3Request
}

type GetDs3TargetFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetDs3TargetFailureNotificationRegistrationSpectraS3Request(notificationId string) *GetDs3TargetFailureNotificationRegistrationSpectraS3Request {
    return &GetDs3TargetFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetDs3TargetFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetDs3TargetFailureNotificationRegistrationsSpectraS3Request() *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    return &GetDs3TargetFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Request *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetDs3TargetFailureNotificationRegistrationsSpectraS3Request {
    getDs3TargetFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getDs3TargetFailureNotificationRegistrationsSpectraS3Request
}

type GetJobCompletedNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetJobCompletedNotificationRegistrationSpectraS3Request(notificationId string) *GetJobCompletedNotificationRegistrationSpectraS3Request {
    return &GetJobCompletedNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetJobCompletedNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetJobCompletedNotificationRegistrationsSpectraS3Request() *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    return &GetJobCompletedNotificationRegistrationsSpectraS3Request{
    }
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.LastPage = true
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

func (getJobCompletedNotificationRegistrationsSpectraS3Request *GetJobCompletedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetJobCompletedNotificationRegistrationsSpectraS3Request {
    getJobCompletedNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getJobCompletedNotificationRegistrationsSpectraS3Request
}

type GetJobCreatedNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetJobCreatedNotificationRegistrationSpectraS3Request(notificationId string) *GetJobCreatedNotificationRegistrationSpectraS3Request {
    return &GetJobCreatedNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetJobCreatedNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetJobCreatedNotificationRegistrationsSpectraS3Request() *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    return &GetJobCreatedNotificationRegistrationsSpectraS3Request{
    }
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.LastPage = true
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

func (getJobCreatedNotificationRegistrationsSpectraS3Request *GetJobCreatedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetJobCreatedNotificationRegistrationsSpectraS3Request {
    getJobCreatedNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getJobCreatedNotificationRegistrationsSpectraS3Request
}

type GetJobCreationFailedNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetJobCreationFailedNotificationRegistrationSpectraS3Request(notificationId string) *GetJobCreationFailedNotificationRegistrationSpectraS3Request {
    return &GetJobCreationFailedNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetJobCreationFailedNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetJobCreationFailedNotificationRegistrationsSpectraS3Request() *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    return &GetJobCreationFailedNotificationRegistrationsSpectraS3Request{
    }
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.LastPage = true
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Request *GetJobCreationFailedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetJobCreationFailedNotificationRegistrationsSpectraS3Request {
    getJobCreationFailedNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getJobCreationFailedNotificationRegistrationsSpectraS3Request
}

type GetObjectCachedNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetObjectCachedNotificationRegistrationSpectraS3Request(notificationId string) *GetObjectCachedNotificationRegistrationSpectraS3Request {
    return &GetObjectCachedNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetObjectCachedNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetObjectCachedNotificationRegistrationsSpectraS3Request() *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    return &GetObjectCachedNotificationRegistrationsSpectraS3Request{
    }
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.LastPage = true
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

func (getObjectCachedNotificationRegistrationsSpectraS3Request *GetObjectCachedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetObjectCachedNotificationRegistrationsSpectraS3Request {
    getObjectCachedNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getObjectCachedNotificationRegistrationsSpectraS3Request
}

type GetObjectLostNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetObjectLostNotificationRegistrationSpectraS3Request(notificationId string) *GetObjectLostNotificationRegistrationSpectraS3Request {
    return &GetObjectLostNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetObjectLostNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetObjectLostNotificationRegistrationsSpectraS3Request() *GetObjectLostNotificationRegistrationsSpectraS3Request {
    return &GetObjectLostNotificationRegistrationsSpectraS3Request{
    }
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithLastPage() *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.LastPage = true
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

func (getObjectLostNotificationRegistrationsSpectraS3Request *GetObjectLostNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetObjectLostNotificationRegistrationsSpectraS3Request {
    getObjectLostNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getObjectLostNotificationRegistrationsSpectraS3Request
}

type GetObjectPersistedNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetObjectPersistedNotificationRegistrationSpectraS3Request(notificationId string) *GetObjectPersistedNotificationRegistrationSpectraS3Request {
    return &GetObjectPersistedNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetObjectPersistedNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetObjectPersistedNotificationRegistrationsSpectraS3Request() *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    return &GetObjectPersistedNotificationRegistrationsSpectraS3Request{
    }
}

func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) WithLastPage() *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    getObjectPersistedNotificationRegistrationsSpectraS3Request.LastPage = true
    return getObjectPersistedNotificationRegistrationsSpectraS3Request
}

func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    getObjectPersistedNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getObjectPersistedNotificationRegistrationsSpectraS3Request
}

func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    getObjectPersistedNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getObjectPersistedNotificationRegistrationsSpectraS3Request
}

func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    getObjectPersistedNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getObjectPersistedNotificationRegistrationsSpectraS3Request
}

func (getObjectPersistedNotificationRegistrationsSpectraS3Request *GetObjectPersistedNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetObjectPersistedNotificationRegistrationsSpectraS3Request {
    getObjectPersistedNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getObjectPersistedNotificationRegistrationsSpectraS3Request
}

type GetPoolFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetPoolFailureNotificationRegistrationSpectraS3Request(notificationId string) *GetPoolFailureNotificationRegistrationSpectraS3Request {
    return &GetPoolFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetPoolFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetPoolFailureNotificationRegistrationsSpectraS3Request() *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    return &GetPoolFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

func (getPoolFailureNotificationRegistrationsSpectraS3Request *GetPoolFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetPoolFailureNotificationRegistrationsSpectraS3Request {
    getPoolFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getPoolFailureNotificationRegistrationsSpectraS3Request
}

type GetS3TargetFailureNotificationRegistrationSpectraS3Request struct {
    S3TargetFailureNotificationRegistration string
}

func NewGetS3TargetFailureNotificationRegistrationSpectraS3Request(s3TargetFailureNotificationRegistration string) *GetS3TargetFailureNotificationRegistrationSpectraS3Request {
    return &GetS3TargetFailureNotificationRegistrationSpectraS3Request{
        S3TargetFailureNotificationRegistration: s3TargetFailureNotificationRegistration,
    }
}

type GetS3TargetFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetS3TargetFailureNotificationRegistrationsSpectraS3Request() *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    return &GetS3TargetFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Request *GetS3TargetFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetS3TargetFailureNotificationRegistrationsSpectraS3Request {
    getS3TargetFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getS3TargetFailureNotificationRegistrationsSpectraS3Request
}

type GetStorageDomainFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetStorageDomainFailureNotificationRegistrationSpectraS3Request(notificationId string) *GetStorageDomainFailureNotificationRegistrationSpectraS3Request {
    return &GetStorageDomainFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetStorageDomainFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetStorageDomainFailureNotificationRegistrationsSpectraS3Request() *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    return &GetStorageDomainFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Request *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetStorageDomainFailureNotificationRegistrationsSpectraS3Request {
    getStorageDomainFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getStorageDomainFailureNotificationRegistrationsSpectraS3Request
}

type GetSystemFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetSystemFailureNotificationRegistrationSpectraS3Request(notificationId string) *GetSystemFailureNotificationRegistrationSpectraS3Request {
    return &GetSystemFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetSystemFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetSystemFailureNotificationRegistrationsSpectraS3Request() *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    return &GetSystemFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

type GetTapeFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetTapeFailureNotificationRegistrationSpectraS3Request(notificationId string) *GetTapeFailureNotificationRegistrationSpectraS3Request {
    return &GetTapeFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetTapeFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetTapeFailureNotificationRegistrationsSpectraS3Request() *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    return &GetTapeFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    getTapeFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getTapeFailureNotificationRegistrationsSpectraS3Request
}

func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    getTapeFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getTapeFailureNotificationRegistrationsSpectraS3Request
}

func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    getTapeFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getTapeFailureNotificationRegistrationsSpectraS3Request
}

func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    getTapeFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapeFailureNotificationRegistrationsSpectraS3Request
}

func (getTapeFailureNotificationRegistrationsSpectraS3Request *GetTapeFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetTapeFailureNotificationRegistrationsSpectraS3Request {
    getTapeFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getTapeFailureNotificationRegistrationsSpectraS3Request
}

type GetTapePartitionFailureNotificationRegistrationSpectraS3Request struct {
    NotificationId string
}

func NewGetTapePartitionFailureNotificationRegistrationSpectraS3Request(notificationId string) *GetTapePartitionFailureNotificationRegistrationSpectraS3Request {
    return &GetTapePartitionFailureNotificationRegistrationSpectraS3Request{
        NotificationId: notificationId,
    }
}

type GetTapePartitionFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetTapePartitionFailureNotificationRegistrationsSpectraS3Request() *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    return &GetTapePartitionFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Request *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetTapePartitionFailureNotificationRegistrationsSpectraS3Request {
    getTapePartitionFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getTapePartitionFailureNotificationRegistrationsSpectraS3Request
}

type DeleteFolderRecursivelySpectraS3Request struct {
    BucketId string
    Folder string
}

func NewDeleteFolderRecursivelySpectraS3Request(bucketId string, folder string) *DeleteFolderRecursivelySpectraS3Request {
    return &DeleteFolderRecursivelySpectraS3Request{
        Folder: folder,
        BucketId: bucketId,
    }
}

type GetBlobPersistenceSpectraS3Request struct {
    RequestPayload string
}

func NewGetBlobPersistenceSpectraS3Request(requestPayload string) *GetBlobPersistenceSpectraS3Request {
    return &GetBlobPersistenceSpectraS3Request{
        RequestPayload: requestPayload,
    }
}

type GetObjectDetailsSpectraS3Request struct {
    ObjectName string
    BucketId string
}

func NewGetObjectDetailsSpectraS3Request(objectName string, bucketId string) *GetObjectDetailsSpectraS3Request {
    return &GetObjectDetailsSpectraS3Request{
        ObjectName: objectName,
        BucketId: bucketId,
    }
}

type GetObjectsDetailsSpectraS3Request struct {
    BucketId *string
    EndDate *int64
    LastPage bool
    Latest *bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    S3ObjectType S3ObjectType
    StartDate *int64
}

func NewGetObjectsDetailsSpectraS3Request() *GetObjectsDetailsSpectraS3Request {
    return &GetObjectsDetailsSpectraS3Request{
    }
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithBucketId(bucketId string) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.BucketId = &bucketId
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithEndDate(endDate int64) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.EndDate = &endDate
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithLastPage() *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.LastPage = true
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithLatest(latest bool) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.Latest = &latest
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithName(name string) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.Name = &name
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithPageLength(pageLength int) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.PageLength = &pageLength
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.PageOffset = &pageOffset
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithStartDate(startDate int64) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.StartDate = &startDate
    return getObjectsDetailsSpectraS3Request
}

func (getObjectsDetailsSpectraS3Request *GetObjectsDetailsSpectraS3Request) WithS3ObjectType(s3ObjectType S3ObjectType) *GetObjectsDetailsSpectraS3Request {
    getObjectsDetailsSpectraS3Request.S3ObjectType = s3ObjectType
    return getObjectsDetailsSpectraS3Request
}

type GetObjectsWithFullDetailsSpectraS3Request struct {
    BucketId *string
    EndDate *int64
    IncludePhysicalPlacement bool
    LastPage bool
    Latest *bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    S3ObjectType S3ObjectType
    StartDate *int64
}

func NewGetObjectsWithFullDetailsSpectraS3Request() *GetObjectsWithFullDetailsSpectraS3Request {
    return &GetObjectsWithFullDetailsSpectraS3Request{
    }
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithBucketId(bucketId string) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.BucketId = &bucketId
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithEndDate(endDate int64) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.EndDate = &endDate
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithIncludePhysicalPlacement() *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.IncludePhysicalPlacement = true
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithLastPage() *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.LastPage = true
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithLatest(latest bool) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.Latest = &latest
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithName(name string) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.Name = &name
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithPageLength(pageLength int) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.PageLength = &pageLength
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithPageOffset(pageOffset int) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.PageOffset = &pageOffset
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithStartDate(startDate int64) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.StartDate = &startDate
    return getObjectsWithFullDetailsSpectraS3Request
}

func (getObjectsWithFullDetailsSpectraS3Request *GetObjectsWithFullDetailsSpectraS3Request) WithS3ObjectType(s3ObjectType S3ObjectType) *GetObjectsWithFullDetailsSpectraS3Request {
    getObjectsWithFullDetailsSpectraS3Request.S3ObjectType = s3ObjectType
    return getObjectsWithFullDetailsSpectraS3Request
}

type GetPhysicalPlacementForObjectsSpectraS3Request struct {
    BucketName string
    Objects []Ds3GetObject
    StorageDomain *string
}

func NewGetPhysicalPlacementForObjectsSpectraS3Request(bucketName string, objectNames []string) *GetPhysicalPlacementForObjectsSpectraS3Request {

    return &GetPhysicalPlacementForObjectsSpectraS3Request{
        BucketName: bucketName,
        Objects: buildDs3GetObjectSliceFromNames(objectNames),
    }
}

func NewGetPhysicalPlacementForObjectsSpectraS3RequestWithPartialObjects(bucketName string, objects []Ds3GetObject) *GetPhysicalPlacementForObjectsSpectraS3Request {

    return &GetPhysicalPlacementForObjectsSpectraS3Request{
        BucketName: bucketName,
        Objects: objects,
    }
}

func (getPhysicalPlacementForObjectsSpectraS3Request *GetPhysicalPlacementForObjectsSpectraS3Request) WithStorageDomain(storageDomain string) *GetPhysicalPlacementForObjectsSpectraS3Request {
    getPhysicalPlacementForObjectsSpectraS3Request.StorageDomain = &storageDomain
    return getPhysicalPlacementForObjectsSpectraS3Request
}

type GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request struct {
    BucketName string
    Objects []Ds3GetObject
    StorageDomain *string
}

func NewGetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request(bucketName string, objectNames []string) *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request {

    return &GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request{
        BucketName: bucketName,
        Objects: buildDs3GetObjectSliceFromNames(objectNames),
    }
}

func NewGetPhysicalPlacementForObjectsWithFullDetailsSpectraS3RequestWithPartialObjects(bucketName string, objects []Ds3GetObject) *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request {

    return &GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request{
        BucketName: bucketName,
        Objects: objects,
    }
}

func (getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) WithStorageDomain(storageDomain string) *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request {
    getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.StorageDomain = &storageDomain
    return getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request
}

type UndeleteObjectSpectraS3Request struct {
    BucketId string
    Name string
    VersionId *string
}

func NewUndeleteObjectSpectraS3Request(bucketId string, name string) *UndeleteObjectSpectraS3Request {
    return &UndeleteObjectSpectraS3Request{
        BucketId: bucketId,
        Name: name,
    }
}

func (undeleteObjectSpectraS3Request *UndeleteObjectSpectraS3Request) WithVersionId(versionId string) *UndeleteObjectSpectraS3Request {
    undeleteObjectSpectraS3Request.VersionId = &versionId
    return undeleteObjectSpectraS3Request
}

type VerifyPhysicalPlacementForObjectsSpectraS3Request struct {
    BucketName string
    Objects []Ds3GetObject
    StorageDomain *string
}

func NewVerifyPhysicalPlacementForObjectsSpectraS3Request(bucketName string, objectNames []string) *VerifyPhysicalPlacementForObjectsSpectraS3Request {

    return &VerifyPhysicalPlacementForObjectsSpectraS3Request{
        BucketName: bucketName,
        Objects: buildDs3GetObjectSliceFromNames(objectNames),
    }
}

func NewVerifyPhysicalPlacementForObjectsSpectraS3RequestWithPartialObjects(bucketName string, objects []Ds3GetObject) *VerifyPhysicalPlacementForObjectsSpectraS3Request {

    return &VerifyPhysicalPlacementForObjectsSpectraS3Request{
        BucketName: bucketName,
        Objects: objects,
    }
}

func (verifyPhysicalPlacementForObjectsSpectraS3Request *VerifyPhysicalPlacementForObjectsSpectraS3Request) WithStorageDomain(storageDomain string) *VerifyPhysicalPlacementForObjectsSpectraS3Request {
    verifyPhysicalPlacementForObjectsSpectraS3Request.StorageDomain = &storageDomain
    return verifyPhysicalPlacementForObjectsSpectraS3Request
}

type VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request struct {
    BucketName string
    Objects []Ds3GetObject
    StorageDomain *string
}

func NewVerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request(bucketName string, objectNames []string) *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request {

    return &VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request{
        BucketName: bucketName,
        Objects: buildDs3GetObjectSliceFromNames(objectNames),
    }
}

func NewVerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3RequestWithPartialObjects(bucketName string, objects []Ds3GetObject) *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request {

    return &VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request{
        BucketName: bucketName,
        Objects: objects,
    }
}

func (verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request) WithStorageDomain(storageDomain string) *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request {
    verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request.StorageDomain = &storageDomain
    return verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Request
}

type CancelImportOnAllPoolsSpectraS3Request struct {
}

func NewCancelImportOnAllPoolsSpectraS3Request() *CancelImportOnAllPoolsSpectraS3Request {
    return &CancelImportOnAllPoolsSpectraS3Request{
    }
}

type CancelImportPoolSpectraS3Request struct {
    Pool string
}

func NewCancelImportPoolSpectraS3Request(pool string) *CancelImportPoolSpectraS3Request {
    return &CancelImportPoolSpectraS3Request{
        Pool: pool,
    }
}

type CancelVerifyOnAllPoolsSpectraS3Request struct {
}

func NewCancelVerifyOnAllPoolsSpectraS3Request() *CancelVerifyOnAllPoolsSpectraS3Request {
    return &CancelVerifyOnAllPoolsSpectraS3Request{
    }
}

type CancelVerifyPoolSpectraS3Request struct {
    Pool string
}

func NewCancelVerifyPoolSpectraS3Request(pool string) *CancelVerifyPoolSpectraS3Request {
    return &CancelVerifyPoolSpectraS3Request{
        Pool: pool,
    }
}

type CompactAllPoolsSpectraS3Request struct {
    Priority Priority
}

func NewCompactAllPoolsSpectraS3Request() *CompactAllPoolsSpectraS3Request {
    return &CompactAllPoolsSpectraS3Request{
    }
}

func (compactAllPoolsSpectraS3Request *CompactAllPoolsSpectraS3Request) WithPriority(priority Priority) *CompactAllPoolsSpectraS3Request {
    compactAllPoolsSpectraS3Request.Priority = priority
    return compactAllPoolsSpectraS3Request
}

type CompactPoolSpectraS3Request struct {
    Pool string
    Priority Priority
}

func NewCompactPoolSpectraS3Request(pool string) *CompactPoolSpectraS3Request {
    return &CompactPoolSpectraS3Request{
        Pool: pool,
    }
}

func (compactPoolSpectraS3Request *CompactPoolSpectraS3Request) WithPriority(priority Priority) *CompactPoolSpectraS3Request {
    compactPoolSpectraS3Request.Priority = priority
    return compactPoolSpectraS3Request
}

type PutPoolPartitionSpectraS3Request struct {
    Name string
    PoolType PoolType
}

func NewPutPoolPartitionSpectraS3Request(name string, poolType PoolType) *PutPoolPartitionSpectraS3Request {
    return &PutPoolPartitionSpectraS3Request{
        Name: name,
        PoolType: poolType,
    }
}

type DeallocatePoolSpectraS3Request struct {
    Pool string
}

func NewDeallocatePoolSpectraS3Request(pool string) *DeallocatePoolSpectraS3Request {
    return &DeallocatePoolSpectraS3Request{
        Pool: pool,
    }
}

type DeletePermanentlyLostPoolSpectraS3Request struct {
    Pool string
}

func NewDeletePermanentlyLostPoolSpectraS3Request(pool string) *DeletePermanentlyLostPoolSpectraS3Request {
    return &DeletePermanentlyLostPoolSpectraS3Request{
        Pool: pool,
    }
}

type DeletePoolFailureSpectraS3Request struct {
    PoolFailure string
}

func NewDeletePoolFailureSpectraS3Request(poolFailure string) *DeletePoolFailureSpectraS3Request {
    return &DeletePoolFailureSpectraS3Request{
        PoolFailure: poolFailure,
    }
}

type DeletePoolPartitionSpectraS3Request struct {
    PoolPartition string
}

func NewDeletePoolPartitionSpectraS3Request(poolPartition string) *DeletePoolPartitionSpectraS3Request {
    return &DeletePoolPartitionSpectraS3Request{
        PoolPartition: poolPartition,
    }
}

type ForcePoolEnvironmentRefreshSpectraS3Request struct {
}

func NewForcePoolEnvironmentRefreshSpectraS3Request() *ForcePoolEnvironmentRefreshSpectraS3Request {
    return &ForcePoolEnvironmentRefreshSpectraS3Request{
    }
}

type FormatAllForeignPoolsSpectraS3Request struct {
}

func NewFormatAllForeignPoolsSpectraS3Request() *FormatAllForeignPoolsSpectraS3Request {
    return &FormatAllForeignPoolsSpectraS3Request{
    }
}

type FormatForeignPoolSpectraS3Request struct {
    Pool string
}

func NewFormatForeignPoolSpectraS3Request(pool string) *FormatForeignPoolSpectraS3Request {
    return &FormatForeignPoolSpectraS3Request{
        Pool: pool,
    }
}

type GetBlobsOnPoolSpectraS3Request struct {
    Pool string
}

func NewGetBlobsOnPoolSpectraS3Request(pool string) *GetBlobsOnPoolSpectraS3Request {
    return &GetBlobsOnPoolSpectraS3Request{
        Pool: pool,
    }
}

type GetPoolFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PoolFailureType PoolFailureType
    PoolId *string
}

func NewGetPoolFailuresSpectraS3Request() *GetPoolFailuresSpectraS3Request {
    return &GetPoolFailuresSpectraS3Request{
    }
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithLastPage() *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.LastPage = true
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPageLength(pageLength int) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.PageLength = &pageLength
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.PageOffset = &pageOffset
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPoolId(poolId string) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.PoolId = &poolId
    return getPoolFailuresSpectraS3Request
}

func (getPoolFailuresSpectraS3Request *GetPoolFailuresSpectraS3Request) WithPoolFailureType(poolFailureType PoolFailureType) *GetPoolFailuresSpectraS3Request {
    getPoolFailuresSpectraS3Request.PoolFailureType = poolFailureType
    return getPoolFailuresSpectraS3Request
}

type GetPoolPartitionSpectraS3Request struct {
    PoolPartition string
}

func NewGetPoolPartitionSpectraS3Request(poolPartition string) *GetPoolPartitionSpectraS3Request {
    return &GetPoolPartitionSpectraS3Request{
        PoolPartition: poolPartition,
    }
}

type GetPoolPartitionsSpectraS3Request struct {
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PoolType PoolType
}

func NewGetPoolPartitionsSpectraS3Request() *GetPoolPartitionsSpectraS3Request {
    return &GetPoolPartitionsSpectraS3Request{
    }
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithLastPage() *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.LastPage = true
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithName(name string) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.Name = &name
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPageLength(pageLength int) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.PageLength = &pageLength
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.PageOffset = &pageOffset
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getPoolPartitionsSpectraS3Request
}

func (getPoolPartitionsSpectraS3Request *GetPoolPartitionsSpectraS3Request) WithPoolType(poolType PoolType) *GetPoolPartitionsSpectraS3Request {
    getPoolPartitionsSpectraS3Request.PoolType = poolType
    return getPoolPartitionsSpectraS3Request
}

type GetPoolSpectraS3Request struct {
    Pool string
}

func NewGetPoolSpectraS3Request(pool string) *GetPoolSpectraS3Request {
    return &GetPoolSpectraS3Request{
        Pool: pool,
    }
}

type GetPoolsSpectraS3Request struct {
    AssignedToStorageDomain *bool
    BucketId *string
    Guid *string
    Health PoolHealth
    LastPage bool
    LastVerified *string
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PartitionId *string
    PoolType PoolType
    PoweredOn *bool
    State PoolState
    StorageDomainMemberId *string
}

func NewGetPoolsSpectraS3Request() *GetPoolsSpectraS3Request {
    return &GetPoolsSpectraS3Request{
    }
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithAssignedToStorageDomain(assignedToStorageDomain bool) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.AssignedToStorageDomain = &assignedToStorageDomain
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithBucketId(bucketId string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.BucketId = &bucketId
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithGuid(guid string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.Guid = &guid
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithHealth(health PoolHealth) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.Health = health
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithLastPage() *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.LastPage = true
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithLastVerified(lastVerified string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.LastVerified = &lastVerified
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithName(name string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.Name = &name
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPageLength(pageLength int) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PageLength = &pageLength
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPageOffset(pageOffset int) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PageOffset = &pageOffset
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPartitionId(partitionId string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PartitionId = &partitionId
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPoweredOn(poweredOn bool) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PoweredOn = &poweredOn
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithState(state PoolState) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.State = state
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithStorageDomainMemberId(storageDomainMemberId string) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.StorageDomainMemberId = &storageDomainMemberId
    return getPoolsSpectraS3Request
}

func (getPoolsSpectraS3Request *GetPoolsSpectraS3Request) WithPoolType(poolType PoolType) *GetPoolsSpectraS3Request {
    getPoolsSpectraS3Request.PoolType = poolType
    return getPoolsSpectraS3Request
}

type ImportAllPoolsSpectraS3Request struct {
    DataPolicyId *string
    Priority Priority
    StorageDomainId *string
    UserId *string
    VerifyDataAfterImport Priority
    VerifyDataPriorToImport *bool
}

func NewImportAllPoolsSpectraS3Request() *ImportAllPoolsSpectraS3Request {
    return &ImportAllPoolsSpectraS3Request{
    }
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.DataPolicyId = &dataPolicyId
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithPriority(priority Priority) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.Priority = priority
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.StorageDomainId = &storageDomainId
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithUserId(userId string) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.UserId = &userId
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.VerifyDataAfterImport = verifyDataAfterImport
    return importAllPoolsSpectraS3Request
}

func (importAllPoolsSpectraS3Request *ImportAllPoolsSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportAllPoolsSpectraS3Request {
    importAllPoolsSpectraS3Request.VerifyDataPriorToImport = &verifyDataPriorToImport
    return importAllPoolsSpectraS3Request
}

type ImportPoolSpectraS3Request struct {
    DataPolicyId *string
    Pool string
    Priority Priority
    StorageDomainId *string
    UserId *string
    VerifyDataAfterImport Priority
    VerifyDataPriorToImport *bool
}

func NewImportPoolSpectraS3Request(pool string) *ImportPoolSpectraS3Request {
    return &ImportPoolSpectraS3Request{
        Pool: pool,
    }
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.DataPolicyId = &dataPolicyId
    return importPoolSpectraS3Request
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithPriority(priority Priority) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.Priority = priority
    return importPoolSpectraS3Request
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.StorageDomainId = &storageDomainId
    return importPoolSpectraS3Request
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithUserId(userId string) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.UserId = &userId
    return importPoolSpectraS3Request
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.VerifyDataAfterImport = verifyDataAfterImport
    return importPoolSpectraS3Request
}

func (importPoolSpectraS3Request *ImportPoolSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportPoolSpectraS3Request {
    importPoolSpectraS3Request.VerifyDataPriorToImport = &verifyDataPriorToImport
    return importPoolSpectraS3Request
}

type ModifyAllPoolsSpectraS3Request struct {
    Quiesced Quiesced
}

func NewModifyAllPoolsSpectraS3Request(quiesced Quiesced) *ModifyAllPoolsSpectraS3Request {
    return &ModifyAllPoolsSpectraS3Request{
        Quiesced: quiesced,
    }
}

type ModifyPoolPartitionSpectraS3Request struct {
    Name *string
    PoolPartition string
}

func NewModifyPoolPartitionSpectraS3Request(poolPartition string) *ModifyPoolPartitionSpectraS3Request {
    return &ModifyPoolPartitionSpectraS3Request{
        PoolPartition: poolPartition,
    }
}

func (modifyPoolPartitionSpectraS3Request *ModifyPoolPartitionSpectraS3Request) WithName(name string) *ModifyPoolPartitionSpectraS3Request {
    modifyPoolPartitionSpectraS3Request.Name = &name
    return modifyPoolPartitionSpectraS3Request
}

type ModifyPoolSpectraS3Request struct {
    PartitionId *string
    Pool string
    Quiesced Quiesced
}

func NewModifyPoolSpectraS3Request(pool string) *ModifyPoolSpectraS3Request {
    return &ModifyPoolSpectraS3Request{
        Pool: pool,
    }
}

func (modifyPoolSpectraS3Request *ModifyPoolSpectraS3Request) WithPartitionId(partitionId string) *ModifyPoolSpectraS3Request {
    modifyPoolSpectraS3Request.PartitionId = &partitionId
    return modifyPoolSpectraS3Request
}

func (modifyPoolSpectraS3Request *ModifyPoolSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyPoolSpectraS3Request {
    modifyPoolSpectraS3Request.Quiesced = quiesced
    return modifyPoolSpectraS3Request
}

type VerifyAllPoolsSpectraS3Request struct {
    Priority Priority
}

func NewVerifyAllPoolsSpectraS3Request() *VerifyAllPoolsSpectraS3Request {
    return &VerifyAllPoolsSpectraS3Request{
    }
}

func (verifyAllPoolsSpectraS3Request *VerifyAllPoolsSpectraS3Request) WithPriority(priority Priority) *VerifyAllPoolsSpectraS3Request {
    verifyAllPoolsSpectraS3Request.Priority = priority
    return verifyAllPoolsSpectraS3Request
}

type VerifyPoolSpectraS3Request struct {
    Pool string
    Priority Priority
}

func NewVerifyPoolSpectraS3Request(pool string) *VerifyPoolSpectraS3Request {
    return &VerifyPoolSpectraS3Request{
        Pool: pool,
    }
}

func (verifyPoolSpectraS3Request *VerifyPoolSpectraS3Request) WithPriority(priority Priority) *VerifyPoolSpectraS3Request {
    verifyPoolSpectraS3Request.Priority = priority
    return verifyPoolSpectraS3Request
}

type ConvertStorageDomainToDs3TargetSpectraS3Request struct {
    ConvertToDs3Target string
    StorageDomain string
}

func NewConvertStorageDomainToDs3TargetSpectraS3Request(convertToDs3Target string, storageDomain string) *ConvertStorageDomainToDs3TargetSpectraS3Request {
    return &ConvertStorageDomainToDs3TargetSpectraS3Request{
        StorageDomain: storageDomain,
        ConvertToDs3Target: convertToDs3Target,
    }
}

type PutPoolStorageDomainMemberSpectraS3Request struct {
    PoolPartitionId string
    StorageDomainId string
    WritePreference WritePreferenceLevel
}

func NewPutPoolStorageDomainMemberSpectraS3Request(poolPartitionId string, storageDomainId string) *PutPoolStorageDomainMemberSpectraS3Request {
    return &PutPoolStorageDomainMemberSpectraS3Request{
        PoolPartitionId: poolPartitionId,
        StorageDomainId: storageDomainId,
    }
}

func (putPoolStorageDomainMemberSpectraS3Request *PutPoolStorageDomainMemberSpectraS3Request) WithWritePreference(writePreference WritePreferenceLevel) *PutPoolStorageDomainMemberSpectraS3Request {
    putPoolStorageDomainMemberSpectraS3Request.WritePreference = writePreference
    return putPoolStorageDomainMemberSpectraS3Request
}

type PutStorageDomainSpectraS3Request struct {
    AutoEjectMediaFullThreshold *int64
    AutoEjectUponCron *string
    AutoEjectUponJobCancellation *bool
    AutoEjectUponJobCompletion *bool
    AutoEjectUponMediaFull *bool
    LtfsFileNaming LtfsFileNamingMode
    MaximumAutoVerificationFrequencyInDays *int
    MaxTapeFragmentationPercent *int
    MediaEjectionAllowed *bool
    Name string
    SecureMediaAllocation *bool
    VerifyPriorToAutoEject Priority
    WriteOptimization WriteOptimization
}

func NewPutStorageDomainSpectraS3Request(name string) *PutStorageDomainSpectraS3Request {
    return &PutStorageDomainSpectraS3Request{
        Name: name,
    }
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectMediaFullThreshold(autoEjectMediaFullThreshold int64) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.AutoEjectMediaFullThreshold = &autoEjectMediaFullThreshold
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponCron(autoEjectUponCron string) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.AutoEjectUponCron = &autoEjectUponCron
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponJobCancellation(autoEjectUponJobCancellation bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.AutoEjectUponJobCancellation = &autoEjectUponJobCancellation
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponJobCompletion(autoEjectUponJobCompletion bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.AutoEjectUponJobCompletion = &autoEjectUponJobCompletion
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithAutoEjectUponMediaFull(autoEjectUponMediaFull bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.AutoEjectUponMediaFull = &autoEjectUponMediaFull
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithLtfsFileNaming(ltfsFileNaming LtfsFileNamingMode) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.LtfsFileNaming = ltfsFileNaming
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithMaxTapeFragmentationPercent(maxTapeFragmentationPercent int) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.MaxTapeFragmentationPercent = &maxTapeFragmentationPercent
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithMaximumAutoVerificationFrequencyInDays(maximumAutoVerificationFrequencyInDays int) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.MaximumAutoVerificationFrequencyInDays = &maximumAutoVerificationFrequencyInDays
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithMediaEjectionAllowed(mediaEjectionAllowed bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.MediaEjectionAllowed = &mediaEjectionAllowed
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithSecureMediaAllocation(secureMediaAllocation bool) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.SecureMediaAllocation = &secureMediaAllocation
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithVerifyPriorToAutoEject(verifyPriorToAutoEject Priority) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.VerifyPriorToAutoEject = verifyPriorToAutoEject
    return putStorageDomainSpectraS3Request
}

func (putStorageDomainSpectraS3Request *PutStorageDomainSpectraS3Request) WithWriteOptimization(writeOptimization WriteOptimization) *PutStorageDomainSpectraS3Request {
    putStorageDomainSpectraS3Request.WriteOptimization = writeOptimization
    return putStorageDomainSpectraS3Request
}

type PutTapeStorageDomainMemberSpectraS3Request struct {
    AutoCompactionThreshold *int
    StorageDomainId string
    TapePartitionId string
    TapeType string
    WritePreference WritePreferenceLevel
}

func NewPutTapeStorageDomainMemberSpectraS3Request(storageDomainId string, tapePartitionId string, tapeType string) *PutTapeStorageDomainMemberSpectraS3Request {
    return &PutTapeStorageDomainMemberSpectraS3Request{
        StorageDomainId: storageDomainId,
        TapePartitionId: tapePartitionId,
        TapeType: tapeType,
    }
}

func (putTapeStorageDomainMemberSpectraS3Request *PutTapeStorageDomainMemberSpectraS3Request) WithAutoCompactionThreshold(autoCompactionThreshold int) *PutTapeStorageDomainMemberSpectraS3Request {
    putTapeStorageDomainMemberSpectraS3Request.AutoCompactionThreshold = &autoCompactionThreshold
    return putTapeStorageDomainMemberSpectraS3Request
}

func (putTapeStorageDomainMemberSpectraS3Request *PutTapeStorageDomainMemberSpectraS3Request) WithWritePreference(writePreference WritePreferenceLevel) *PutTapeStorageDomainMemberSpectraS3Request {
    putTapeStorageDomainMemberSpectraS3Request.WritePreference = writePreference
    return putTapeStorageDomainMemberSpectraS3Request
}

type DeleteStorageDomainFailureSpectraS3Request struct {
    StorageDomainFailure string
}

func NewDeleteStorageDomainFailureSpectraS3Request(storageDomainFailure string) *DeleteStorageDomainFailureSpectraS3Request {
    return &DeleteStorageDomainFailureSpectraS3Request{
        StorageDomainFailure: storageDomainFailure,
    }
}

type DeleteStorageDomainMemberSpectraS3Request struct {
    StorageDomainMember string
}

func NewDeleteStorageDomainMemberSpectraS3Request(storageDomainMember string) *DeleteStorageDomainMemberSpectraS3Request {
    return &DeleteStorageDomainMemberSpectraS3Request{
        StorageDomainMember: storageDomainMember,
    }
}

type DeleteStorageDomainSpectraS3Request struct {
    StorageDomain string
}

func NewDeleteStorageDomainSpectraS3Request(storageDomain string) *DeleteStorageDomainSpectraS3Request {
    return &DeleteStorageDomainSpectraS3Request{
        StorageDomain: storageDomain,
    }
}

type GetStorageDomainFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    StorageDomainFailureType StorageDomainFailureType
    StorageDomainId *string
}

func NewGetStorageDomainFailuresSpectraS3Request() *GetStorageDomainFailuresSpectraS3Request {
    return &GetStorageDomainFailuresSpectraS3Request{
    }
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithLastPage() *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.LastPage = true
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.PageLength = &pageLength
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.PageOffset = &pageOffset
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.StorageDomainId = &storageDomainId
    return getStorageDomainFailuresSpectraS3Request
}

func (getStorageDomainFailuresSpectraS3Request *GetStorageDomainFailuresSpectraS3Request) WithStorageDomainFailureType(storageDomainFailureType StorageDomainFailureType) *GetStorageDomainFailuresSpectraS3Request {
    getStorageDomainFailuresSpectraS3Request.StorageDomainFailureType = storageDomainFailureType
    return getStorageDomainFailuresSpectraS3Request
}

type GetStorageDomainMemberSpectraS3Request struct {
    StorageDomainMember string
}

func NewGetStorageDomainMemberSpectraS3Request(storageDomainMember string) *GetStorageDomainMemberSpectraS3Request {
    return &GetStorageDomainMemberSpectraS3Request{
        StorageDomainMember: storageDomainMember,
    }
}

type GetStorageDomainMembersSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PoolPartitionId *string
    State StorageDomainMemberState
    StorageDomainId *string
    TapePartitionId *string
    TapeType *string
    WritePreference WritePreferenceLevel
}

func NewGetStorageDomainMembersSpectraS3Request() *GetStorageDomainMembersSpectraS3Request {
    return &GetStorageDomainMembersSpectraS3Request{
    }
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithLastPage() *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.LastPage = true
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.PageLength = &pageLength
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.PageOffset = &pageOffset
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.PageStartMarker = &pageStartMarker
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithPoolPartitionId(poolPartitionId string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.PoolPartitionId = &poolPartitionId
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithState(state StorageDomainMemberState) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.State = state
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithStorageDomainId(storageDomainId string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.StorageDomainId = &storageDomainId
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithTapePartitionId(tapePartitionId string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.TapePartitionId = &tapePartitionId
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithTapeType(tapeType string) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.TapeType = &tapeType
    return getStorageDomainMembersSpectraS3Request
}

func (getStorageDomainMembersSpectraS3Request *GetStorageDomainMembersSpectraS3Request) WithWritePreference(writePreference WritePreferenceLevel) *GetStorageDomainMembersSpectraS3Request {
    getStorageDomainMembersSpectraS3Request.WritePreference = writePreference
    return getStorageDomainMembersSpectraS3Request
}

type GetStorageDomainSpectraS3Request struct {
    StorageDomain string
}

func NewGetStorageDomainSpectraS3Request(storageDomain string) *GetStorageDomainSpectraS3Request {
    return &GetStorageDomainSpectraS3Request{
        StorageDomain: storageDomain,
    }
}

type GetStorageDomainsSpectraS3Request struct {
    AutoEjectUponCron *string
    AutoEjectUponJobCancellation *bool
    AutoEjectUponJobCompletion *bool
    AutoEjectUponMediaFull *bool
    LastPage bool
    MediaEjectionAllowed *bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    SecureMediaAllocation *bool
    WriteOptimization WriteOptimization
}

func NewGetStorageDomainsSpectraS3Request() *GetStorageDomainsSpectraS3Request {
    return &GetStorageDomainsSpectraS3Request{
    }
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponCron(autoEjectUponCron string) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.AutoEjectUponCron = &autoEjectUponCron
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponJobCancellation(autoEjectUponJobCancellation bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.AutoEjectUponJobCancellation = &autoEjectUponJobCancellation
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponJobCompletion(autoEjectUponJobCompletion bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.AutoEjectUponJobCompletion = &autoEjectUponJobCompletion
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithAutoEjectUponMediaFull(autoEjectUponMediaFull bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.AutoEjectUponMediaFull = &autoEjectUponMediaFull
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithLastPage() *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.LastPage = true
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithMediaEjectionAllowed(mediaEjectionAllowed bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.MediaEjectionAllowed = &mediaEjectionAllowed
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithName(name string) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.Name = &name
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithPageLength(pageLength int) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.PageLength = &pageLength
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithPageOffset(pageOffset int) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.PageOffset = &pageOffset
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithSecureMediaAllocation(secureMediaAllocation bool) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.SecureMediaAllocation = &secureMediaAllocation
    return getStorageDomainsSpectraS3Request
}

func (getStorageDomainsSpectraS3Request *GetStorageDomainsSpectraS3Request) WithWriteOptimization(writeOptimization WriteOptimization) *GetStorageDomainsSpectraS3Request {
    getStorageDomainsSpectraS3Request.WriteOptimization = writeOptimization
    return getStorageDomainsSpectraS3Request
}

type ModifyStorageDomainMemberSpectraS3Request struct {
    AutoCompactionThreshold *int
    State StorageDomainMemberState
    StorageDomainMember string
    WritePreference WritePreferenceLevel
}

func NewModifyStorageDomainMemberSpectraS3Request(storageDomainMember string) *ModifyStorageDomainMemberSpectraS3Request {
    return &ModifyStorageDomainMemberSpectraS3Request{
        StorageDomainMember: storageDomainMember,
    }
}

func (modifyStorageDomainMemberSpectraS3Request *ModifyStorageDomainMemberSpectraS3Request) WithAutoCompactionThreshold(autoCompactionThreshold int) *ModifyStorageDomainMemberSpectraS3Request {
    modifyStorageDomainMemberSpectraS3Request.AutoCompactionThreshold = &autoCompactionThreshold
    return modifyStorageDomainMemberSpectraS3Request
}

func (modifyStorageDomainMemberSpectraS3Request *ModifyStorageDomainMemberSpectraS3Request) WithState(state StorageDomainMemberState) *ModifyStorageDomainMemberSpectraS3Request {
    modifyStorageDomainMemberSpectraS3Request.State = state
    return modifyStorageDomainMemberSpectraS3Request
}

func (modifyStorageDomainMemberSpectraS3Request *ModifyStorageDomainMemberSpectraS3Request) WithWritePreference(writePreference WritePreferenceLevel) *ModifyStorageDomainMemberSpectraS3Request {
    modifyStorageDomainMemberSpectraS3Request.WritePreference = writePreference
    return modifyStorageDomainMemberSpectraS3Request
}

type ModifyStorageDomainSpectraS3Request struct {
    AutoEjectMediaFullThreshold *int64
    AutoEjectUponCron *string
    AutoEjectUponJobCancellation *bool
    AutoEjectUponJobCompletion *bool
    AutoEjectUponMediaFull *bool
    LtfsFileNaming LtfsFileNamingMode
    MaximumAutoVerificationFrequencyInDays *int
    MaxTapeFragmentationPercent *int
    MediaEjectionAllowed *bool
    Name *string
    SecureMediaAllocation *bool
    StorageDomain string
    VerifyPriorToAutoEject Priority
    WriteOptimization WriteOptimization
}

func NewModifyStorageDomainSpectraS3Request(storageDomain string) *ModifyStorageDomainSpectraS3Request {
    return &ModifyStorageDomainSpectraS3Request{
        StorageDomain: storageDomain,
    }
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectMediaFullThreshold(autoEjectMediaFullThreshold int64) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.AutoEjectMediaFullThreshold = &autoEjectMediaFullThreshold
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponCron(autoEjectUponCron string) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.AutoEjectUponCron = &autoEjectUponCron
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponJobCancellation(autoEjectUponJobCancellation bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.AutoEjectUponJobCancellation = &autoEjectUponJobCancellation
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponJobCompletion(autoEjectUponJobCompletion bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.AutoEjectUponJobCompletion = &autoEjectUponJobCompletion
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithAutoEjectUponMediaFull(autoEjectUponMediaFull bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.AutoEjectUponMediaFull = &autoEjectUponMediaFull
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithLtfsFileNaming(ltfsFileNaming LtfsFileNamingMode) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.LtfsFileNaming = ltfsFileNaming
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithMaxTapeFragmentationPercent(maxTapeFragmentationPercent int) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.MaxTapeFragmentationPercent = &maxTapeFragmentationPercent
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithMaximumAutoVerificationFrequencyInDays(maximumAutoVerificationFrequencyInDays int) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.MaximumAutoVerificationFrequencyInDays = &maximumAutoVerificationFrequencyInDays
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithMediaEjectionAllowed(mediaEjectionAllowed bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.MediaEjectionAllowed = &mediaEjectionAllowed
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithName(name string) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.Name = &name
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithSecureMediaAllocation(secureMediaAllocation bool) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.SecureMediaAllocation = &secureMediaAllocation
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithVerifyPriorToAutoEject(verifyPriorToAutoEject Priority) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.VerifyPriorToAutoEject = verifyPriorToAutoEject
    return modifyStorageDomainSpectraS3Request
}

func (modifyStorageDomainSpectraS3Request *ModifyStorageDomainSpectraS3Request) WithWriteOptimization(writeOptimization WriteOptimization) *ModifyStorageDomainSpectraS3Request {
    modifyStorageDomainSpectraS3Request.WriteOptimization = writeOptimization
    return modifyStorageDomainSpectraS3Request
}

type ForceFeatureKeyValidationSpectraS3Request struct {
}

func NewForceFeatureKeyValidationSpectraS3Request() *ForceFeatureKeyValidationSpectraS3Request {
    return &ForceFeatureKeyValidationSpectraS3Request{
    }
}

type GetFeatureKeysSpectraS3Request struct {
    ErrorMessage *string
    ExpirationDate *string
    Key FeatureKeyType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetFeatureKeysSpectraS3Request() *GetFeatureKeysSpectraS3Request {
    return &GetFeatureKeysSpectraS3Request{
    }
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithErrorMessage(errorMessage string) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.ErrorMessage = &errorMessage
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithExpirationDate(expirationDate string) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.ExpirationDate = &expirationDate
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithKey(key FeatureKeyType) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.Key = key
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithLastPage() *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.LastPage = true
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithPageLength(pageLength int) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.PageLength = &pageLength
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithPageOffset(pageOffset int) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.PageOffset = &pageOffset
    return getFeatureKeysSpectraS3Request
}

func (getFeatureKeysSpectraS3Request *GetFeatureKeysSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetFeatureKeysSpectraS3Request {
    getFeatureKeysSpectraS3Request.PageStartMarker = &pageStartMarker
    return getFeatureKeysSpectraS3Request
}

type GetSystemFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    SystemFailureType SystemFailureType
}

func NewGetSystemFailuresSpectraS3Request() *GetSystemFailuresSpectraS3Request {
    return &GetSystemFailuresSpectraS3Request{
    }
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithLastPage() *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.LastPage = true
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithPageLength(pageLength int) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.PageLength = &pageLength
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.PageOffset = &pageOffset
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSystemFailuresSpectraS3Request
}

func (getSystemFailuresSpectraS3Request *GetSystemFailuresSpectraS3Request) WithSystemFailureType(systemFailureType SystemFailureType) *GetSystemFailuresSpectraS3Request {
    getSystemFailuresSpectraS3Request.SystemFailureType = systemFailureType
    return getSystemFailuresSpectraS3Request
}

type GetSystemInformationSpectraS3Request struct {
}

func NewGetSystemInformationSpectraS3Request() *GetSystemInformationSpectraS3Request {
    return &GetSystemInformationSpectraS3Request{
    }
}

type ResetInstanceIdentifierSpectraS3Request struct {
}

func NewResetInstanceIdentifierSpectraS3Request() *ResetInstanceIdentifierSpectraS3Request {
    return &ResetInstanceIdentifierSpectraS3Request{
    }
}

type VerifySystemHealthSpectraS3Request struct {
}

func NewVerifySystemHealthSpectraS3Request() *VerifySystemHealthSpectraS3Request {
    return &VerifySystemHealthSpectraS3Request{
    }
}

type CancelEjectOnAllTapesSpectraS3Request struct {
}

func NewCancelEjectOnAllTapesSpectraS3Request() *CancelEjectOnAllTapesSpectraS3Request {
    return &CancelEjectOnAllTapesSpectraS3Request{
    }
}

type CancelEjectTapeSpectraS3Request struct {
    TapeId string
}

func NewCancelEjectTapeSpectraS3Request(tapeId string) *CancelEjectTapeSpectraS3Request {
    return &CancelEjectTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

type CancelFormatOnAllTapesSpectraS3Request struct {
}

func NewCancelFormatOnAllTapesSpectraS3Request() *CancelFormatOnAllTapesSpectraS3Request {
    return &CancelFormatOnAllTapesSpectraS3Request{
    }
}

type CancelFormatTapeSpectraS3Request struct {
    TapeId string
}

func NewCancelFormatTapeSpectraS3Request(tapeId string) *CancelFormatTapeSpectraS3Request {
    return &CancelFormatTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

type CancelImportOnAllTapesSpectraS3Request struct {
}

func NewCancelImportOnAllTapesSpectraS3Request() *CancelImportOnAllTapesSpectraS3Request {
    return &CancelImportOnAllTapesSpectraS3Request{
    }
}

type CancelImportTapeSpectraS3Request struct {
    TapeId string
}

func NewCancelImportTapeSpectraS3Request(tapeId string) *CancelImportTapeSpectraS3Request {
    return &CancelImportTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

type CancelOnlineOnAllTapesSpectraS3Request struct {
}

func NewCancelOnlineOnAllTapesSpectraS3Request() *CancelOnlineOnAllTapesSpectraS3Request {
    return &CancelOnlineOnAllTapesSpectraS3Request{
    }
}

type CancelOnlineTapeSpectraS3Request struct {
    TapeId string
}

func NewCancelOnlineTapeSpectraS3Request(tapeId string) *CancelOnlineTapeSpectraS3Request {
    return &CancelOnlineTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

type CancelVerifyOnAllTapesSpectraS3Request struct {
}

func NewCancelVerifyOnAllTapesSpectraS3Request() *CancelVerifyOnAllTapesSpectraS3Request {
    return &CancelVerifyOnAllTapesSpectraS3Request{
    }
}

type CancelVerifyTapeSpectraS3Request struct {
    TapeId string
}

func NewCancelVerifyTapeSpectraS3Request(tapeId string) *CancelVerifyTapeSpectraS3Request {
    return &CancelVerifyTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

type CleanTapeDriveSpectraS3Request struct {
    TapeDriveId string
}

func NewCleanTapeDriveSpectraS3Request(tapeDriveId string) *CleanTapeDriveSpectraS3Request {
    return &CleanTapeDriveSpectraS3Request{
        TapeDriveId: tapeDriveId,
    }
}

type PutTapeDensityDirectiveSpectraS3Request struct {
    Density TapeDriveType
    PartitionId string
    TapeType string
}

func NewPutTapeDensityDirectiveSpectraS3Request(density TapeDriveType, partitionId string, tapeType string) *PutTapeDensityDirectiveSpectraS3Request {
    return &PutTapeDensityDirectiveSpectraS3Request{
        Density: density,
        PartitionId: partitionId,
        TapeType: tapeType,
    }
}

type DeletePermanentlyLostTapeSpectraS3Request struct {
    TapeId string
}

func NewDeletePermanentlyLostTapeSpectraS3Request(tapeId string) *DeletePermanentlyLostTapeSpectraS3Request {
    return &DeletePermanentlyLostTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

type DeleteTapeDensityDirectiveSpectraS3Request struct {
    TapeDensityDirective string
}

func NewDeleteTapeDensityDirectiveSpectraS3Request(tapeDensityDirective string) *DeleteTapeDensityDirectiveSpectraS3Request {
    return &DeleteTapeDensityDirectiveSpectraS3Request{
        TapeDensityDirective: tapeDensityDirective,
    }
}

type DeleteTapeDriveSpectraS3Request struct {
    TapeDriveId string
}

func NewDeleteTapeDriveSpectraS3Request(tapeDriveId string) *DeleteTapeDriveSpectraS3Request {
    return &DeleteTapeDriveSpectraS3Request{
        TapeDriveId: tapeDriveId,
    }
}

type DeleteTapeFailureSpectraS3Request struct {
    TapeFailure string
}

func NewDeleteTapeFailureSpectraS3Request(tapeFailure string) *DeleteTapeFailureSpectraS3Request {
    return &DeleteTapeFailureSpectraS3Request{
        TapeFailure: tapeFailure,
    }
}

type DeleteTapePartitionFailureSpectraS3Request struct {
    TapePartitionFailure string
}

func NewDeleteTapePartitionFailureSpectraS3Request(tapePartitionFailure string) *DeleteTapePartitionFailureSpectraS3Request {
    return &DeleteTapePartitionFailureSpectraS3Request{
        TapePartitionFailure: tapePartitionFailure,
    }
}

type DeleteTapePartitionSpectraS3Request struct {
    TapePartition string
}

func NewDeleteTapePartitionSpectraS3Request(tapePartition string) *DeleteTapePartitionSpectraS3Request {
    return &DeleteTapePartitionSpectraS3Request{
        TapePartition: tapePartition,
    }
}

type EjectAllTapesSpectraS3Request struct {
    EjectLabel *string
    EjectLocation *string
}

func NewEjectAllTapesSpectraS3Request() *EjectAllTapesSpectraS3Request {
    return &EjectAllTapesSpectraS3Request{
    }
}

func (ejectAllTapesSpectraS3Request *EjectAllTapesSpectraS3Request) WithEjectLabel(ejectLabel string) *EjectAllTapesSpectraS3Request {
    ejectAllTapesSpectraS3Request.EjectLabel = &ejectLabel
    return ejectAllTapesSpectraS3Request
}

func (ejectAllTapesSpectraS3Request *EjectAllTapesSpectraS3Request) WithEjectLocation(ejectLocation string) *EjectAllTapesSpectraS3Request {
    ejectAllTapesSpectraS3Request.EjectLocation = &ejectLocation
    return ejectAllTapesSpectraS3Request
}

type EjectStorageDomainBlobsSpectraS3Request struct {
    BucketId string
    EjectLabel *string
    EjectLocation *string
    Objects []Ds3GetObject
    StorageDomain string
}

func NewEjectStorageDomainBlobsSpectraS3Request(bucketId string, storageDomain string, objectNames []string) *EjectStorageDomainBlobsSpectraS3Request {

    return &EjectStorageDomainBlobsSpectraS3Request{
        BucketId: bucketId,
        StorageDomain: storageDomain,
        Objects: buildDs3GetObjectSliceFromNames(objectNames),
    }
}

func NewEjectStorageDomainBlobsSpectraS3RequestWithPartialObjects(bucketId string, storageDomain string, objects []Ds3GetObject) *EjectStorageDomainBlobsSpectraS3Request {

    return &EjectStorageDomainBlobsSpectraS3Request{
        BucketId: bucketId,
        StorageDomain: storageDomain,
        Objects: objects,
    }
}

func (ejectStorageDomainBlobsSpectraS3Request *EjectStorageDomainBlobsSpectraS3Request) WithEjectLabel(ejectLabel string) *EjectStorageDomainBlobsSpectraS3Request {
    ejectStorageDomainBlobsSpectraS3Request.EjectLabel = &ejectLabel
    return ejectStorageDomainBlobsSpectraS3Request
}

func (ejectStorageDomainBlobsSpectraS3Request *EjectStorageDomainBlobsSpectraS3Request) WithEjectLocation(ejectLocation string) *EjectStorageDomainBlobsSpectraS3Request {
    ejectStorageDomainBlobsSpectraS3Request.EjectLocation = &ejectLocation
    return ejectStorageDomainBlobsSpectraS3Request
}

type EjectStorageDomainSpectraS3Request struct {
    BucketId *string
    EjectLabel *string
    EjectLocation *string
    StorageDomain string
}

func NewEjectStorageDomainSpectraS3Request(storageDomain string) *EjectStorageDomainSpectraS3Request {
    return &EjectStorageDomainSpectraS3Request{
        StorageDomain: storageDomain,
    }
}

func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) WithBucketId(bucketId string) *EjectStorageDomainSpectraS3Request {
    ejectStorageDomainSpectraS3Request.BucketId = &bucketId
    return ejectStorageDomainSpectraS3Request
}

func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) WithEjectLabel(ejectLabel string) *EjectStorageDomainSpectraS3Request {
    ejectStorageDomainSpectraS3Request.EjectLabel = &ejectLabel
    return ejectStorageDomainSpectraS3Request
}

func (ejectStorageDomainSpectraS3Request *EjectStorageDomainSpectraS3Request) WithEjectLocation(ejectLocation string) *EjectStorageDomainSpectraS3Request {
    ejectStorageDomainSpectraS3Request.EjectLocation = &ejectLocation
    return ejectStorageDomainSpectraS3Request
}

type EjectTapeSpectraS3Request struct {
    EjectLabel *string
    EjectLocation *string
    TapeId string
}

func NewEjectTapeSpectraS3Request(tapeId string) *EjectTapeSpectraS3Request {
    return &EjectTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (ejectTapeSpectraS3Request *EjectTapeSpectraS3Request) WithEjectLabel(ejectLabel string) *EjectTapeSpectraS3Request {
    ejectTapeSpectraS3Request.EjectLabel = &ejectLabel
    return ejectTapeSpectraS3Request
}

func (ejectTapeSpectraS3Request *EjectTapeSpectraS3Request) WithEjectLocation(ejectLocation string) *EjectTapeSpectraS3Request {
    ejectTapeSpectraS3Request.EjectLocation = &ejectLocation
    return ejectTapeSpectraS3Request
}

type ForceTapeEnvironmentRefreshSpectraS3Request struct {
}

func NewForceTapeEnvironmentRefreshSpectraS3Request() *ForceTapeEnvironmentRefreshSpectraS3Request {
    return &ForceTapeEnvironmentRefreshSpectraS3Request{
    }
}

type FormatAllTapesSpectraS3Request struct {
    Force bool
}

func NewFormatAllTapesSpectraS3Request() *FormatAllTapesSpectraS3Request {
    return &FormatAllTapesSpectraS3Request{
    }
}

func (formatAllTapesSpectraS3Request *FormatAllTapesSpectraS3Request) WithForce() *FormatAllTapesSpectraS3Request {
    formatAllTapesSpectraS3Request.Force = true
    return formatAllTapesSpectraS3Request
}

type FormatTapeSpectraS3Request struct {
    Force bool
    TapeId string
}

func NewFormatTapeSpectraS3Request(tapeId string) *FormatTapeSpectraS3Request {
    return &FormatTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (formatTapeSpectraS3Request *FormatTapeSpectraS3Request) WithForce() *FormatTapeSpectraS3Request {
    formatTapeSpectraS3Request.Force = true
    return formatTapeSpectraS3Request
}

type GetBlobsOnTapeSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TapeId string
}

func NewGetBlobsOnTapeSpectraS3Request(tapeId string) *GetBlobsOnTapeSpectraS3Request {
    return &GetBlobsOnTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (getBlobsOnTapeSpectraS3Request *GetBlobsOnTapeSpectraS3Request) WithLastPage() *GetBlobsOnTapeSpectraS3Request {
    getBlobsOnTapeSpectraS3Request.LastPage = true
    return getBlobsOnTapeSpectraS3Request
}

func (getBlobsOnTapeSpectraS3Request *GetBlobsOnTapeSpectraS3Request) WithPageLength(pageLength int) *GetBlobsOnTapeSpectraS3Request {
    getBlobsOnTapeSpectraS3Request.PageLength = &pageLength
    return getBlobsOnTapeSpectraS3Request
}

func (getBlobsOnTapeSpectraS3Request *GetBlobsOnTapeSpectraS3Request) WithPageOffset(pageOffset int) *GetBlobsOnTapeSpectraS3Request {
    getBlobsOnTapeSpectraS3Request.PageOffset = &pageOffset
    return getBlobsOnTapeSpectraS3Request
}

func (getBlobsOnTapeSpectraS3Request *GetBlobsOnTapeSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetBlobsOnTapeSpectraS3Request {
    getBlobsOnTapeSpectraS3Request.PageStartMarker = &pageStartMarker
    return getBlobsOnTapeSpectraS3Request
}

type GetTapeDensityDirectiveSpectraS3Request struct {
    TapeDensityDirective string
}

func NewGetTapeDensityDirectiveSpectraS3Request(tapeDensityDirective string) *GetTapeDensityDirectiveSpectraS3Request {
    return &GetTapeDensityDirectiveSpectraS3Request{
        TapeDensityDirective: tapeDensityDirective,
    }
}

type GetTapeDensityDirectivesSpectraS3Request struct {
    Density TapeDriveType
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PartitionId *string
    TapeType *string
}

func NewGetTapeDensityDirectivesSpectraS3Request() *GetTapeDensityDirectivesSpectraS3Request {
    return &GetTapeDensityDirectivesSpectraS3Request{
    }
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithDensity(density TapeDriveType) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.Density = density
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithLastPage() *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.LastPage = true
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPageLength(pageLength int) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.PageLength = &pageLength
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.PageOffset = &pageOffset
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithPartitionId(partitionId string) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.PartitionId = &partitionId
    return getTapeDensityDirectivesSpectraS3Request
}

func (getTapeDensityDirectivesSpectraS3Request *GetTapeDensityDirectivesSpectraS3Request) WithTapeType(tapeType string) *GetTapeDensityDirectivesSpectraS3Request {
    getTapeDensityDirectivesSpectraS3Request.TapeType = &tapeType
    return getTapeDensityDirectivesSpectraS3Request
}

type GetTapeDriveSpectraS3Request struct {
    TapeDriveId string
}

func NewGetTapeDriveSpectraS3Request(tapeDriveId string) *GetTapeDriveSpectraS3Request {
    return &GetTapeDriveSpectraS3Request{
        TapeDriveId: tapeDriveId,
    }
}

type GetTapeDrivesSpectraS3Request struct {
    LastPage bool
    MinimumTaskPriority Priority
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PartitionId *string
    ReservedTaskType ReservedTaskType
    SerialNumber *string
    State TapeDriveState
    TapeDriveType TapeDriveType
}

func NewGetTapeDrivesSpectraS3Request() *GetTapeDrivesSpectraS3Request {
    return &GetTapeDrivesSpectraS3Request{
    }
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithLastPage() *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.LastPage = true
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithMinimumTaskPriority(minimumTaskPriority Priority) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.MinimumTaskPriority = minimumTaskPriority
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPageLength(pageLength int) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.PageLength = &pageLength
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.PageOffset = &pageOffset
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithPartitionId(partitionId string) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.PartitionId = &partitionId
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithReservedTaskType(reservedTaskType ReservedTaskType) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.ReservedTaskType = reservedTaskType
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithSerialNumber(serialNumber string) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.SerialNumber = &serialNumber
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithState(state TapeDriveState) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.State = state
    return getTapeDrivesSpectraS3Request
}

func (getTapeDrivesSpectraS3Request *GetTapeDrivesSpectraS3Request) WithTapeDriveType(tapeDriveType TapeDriveType) *GetTapeDrivesSpectraS3Request {
    getTapeDrivesSpectraS3Request.TapeDriveType = tapeDriveType
    return getTapeDrivesSpectraS3Request
}

type GetTapeFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TapeDriveId *string
    TapeFailureType TapeFailureType
    TapeId *string
}

func NewGetTapeFailuresSpectraS3Request() *GetTapeFailuresSpectraS3Request {
    return &GetTapeFailuresSpectraS3Request{
    }
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithLastPage() *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.LastPage = true
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithPageLength(pageLength int) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.PageLength = &pageLength
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.PageOffset = &pageOffset
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithTapeDriveId(tapeDriveId string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.TapeDriveId = &tapeDriveId
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithTapeId(tapeId string) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.TapeId = &tapeId
    return getTapeFailuresSpectraS3Request
}

func (getTapeFailuresSpectraS3Request *GetTapeFailuresSpectraS3Request) WithTapeFailureType(tapeFailureType TapeFailureType) *GetTapeFailuresSpectraS3Request {
    getTapeFailuresSpectraS3Request.TapeFailureType = tapeFailureType
    return getTapeFailuresSpectraS3Request
}

type GetTapeLibrariesSpectraS3Request struct {
    LastPage bool
    ManagementUrl *string
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    SerialNumber *string
}

func NewGetTapeLibrariesSpectraS3Request() *GetTapeLibrariesSpectraS3Request {
    return &GetTapeLibrariesSpectraS3Request{
    }
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithLastPage() *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.LastPage = true
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithManagementUrl(managementUrl string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.ManagementUrl = &managementUrl
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithName(name string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.Name = &name
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithPageLength(pageLength int) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.PageLength = &pageLength
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.PageOffset = &pageOffset
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapeLibrariesSpectraS3Request
}

func (getTapeLibrariesSpectraS3Request *GetTapeLibrariesSpectraS3Request) WithSerialNumber(serialNumber string) *GetTapeLibrariesSpectraS3Request {
    getTapeLibrariesSpectraS3Request.SerialNumber = &serialNumber
    return getTapeLibrariesSpectraS3Request
}

type GetTapeLibrarySpectraS3Request struct {
    TapeLibraryId string
}

func NewGetTapeLibrarySpectraS3Request(tapeLibraryId string) *GetTapeLibrarySpectraS3Request {
    return &GetTapeLibrarySpectraS3Request{
        TapeLibraryId: tapeLibraryId,
    }
}

type GetTapePartitionFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PartitionId *string
    TapePartitionFailureType TapePartitionFailureType
}

func NewGetTapePartitionFailuresSpectraS3Request() *GetTapePartitionFailuresSpectraS3Request {
    return &GetTapePartitionFailuresSpectraS3Request{
    }
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithLastPage() *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.LastPage = true
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.PageLength = &pageLength
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.PageOffset = &pageOffset
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithPartitionId(partitionId string) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.PartitionId = &partitionId
    return getTapePartitionFailuresSpectraS3Request
}

func (getTapePartitionFailuresSpectraS3Request *GetTapePartitionFailuresSpectraS3Request) WithTapePartitionFailureType(tapePartitionFailureType TapePartitionFailureType) *GetTapePartitionFailuresSpectraS3Request {
    getTapePartitionFailuresSpectraS3Request.TapePartitionFailureType = tapePartitionFailureType
    return getTapePartitionFailuresSpectraS3Request
}

type GetTapePartitionSpectraS3Request struct {
    TapePartition string
}

func NewGetTapePartitionSpectraS3Request(tapePartition string) *GetTapePartitionSpectraS3Request {
    return &GetTapePartitionSpectraS3Request{
        TapePartition: tapePartition,
    }
}

type GetTapePartitionWithFullDetailsSpectraS3Request struct {
    TapePartition string
}

func NewGetTapePartitionWithFullDetailsSpectraS3Request(tapePartition string) *GetTapePartitionWithFullDetailsSpectraS3Request {
    return &GetTapePartitionWithFullDetailsSpectraS3Request{
        TapePartition: tapePartition,
    }
}

type GetTapePartitionsSpectraS3Request struct {
    ImportExportConfiguration ImportExportConfiguration
    LastPage bool
    LibraryId *string
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    Quiesced Quiesced
    SerialNumber *string
    State TapePartitionState
}

func NewGetTapePartitionsSpectraS3Request() *GetTapePartitionsSpectraS3Request {
    return &GetTapePartitionsSpectraS3Request{
    }
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithImportExportConfiguration(importExportConfiguration ImportExportConfiguration) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.ImportExportConfiguration = importExportConfiguration
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithLastPage() *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.LastPage = true
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithLibraryId(libraryId string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.LibraryId = &libraryId
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithName(name string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.Name = &name
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.PageLength = &pageLength
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.PageOffset = &pageOffset
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.Quiesced = quiesced
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithSerialNumber(serialNumber string) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.SerialNumber = &serialNumber
    return getTapePartitionsSpectraS3Request
}

func (getTapePartitionsSpectraS3Request *GetTapePartitionsSpectraS3Request) WithState(state TapePartitionState) *GetTapePartitionsSpectraS3Request {
    getTapePartitionsSpectraS3Request.State = state
    return getTapePartitionsSpectraS3Request
}

type GetTapePartitionsWithFullDetailsSpectraS3Request struct {
    ImportExportConfiguration ImportExportConfiguration
    LastPage bool
    LibraryId *string
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    Quiesced Quiesced
    SerialNumber *string
    State TapePartitionState
}

func NewGetTapePartitionsWithFullDetailsSpectraS3Request() *GetTapePartitionsWithFullDetailsSpectraS3Request {
    return &GetTapePartitionsWithFullDetailsSpectraS3Request{
    }
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithImportExportConfiguration(importExportConfiguration ImportExportConfiguration) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.ImportExportConfiguration = importExportConfiguration
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithLastPage() *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.LastPage = true
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithLibraryId(libraryId string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.LibraryId = &libraryId
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithName(name string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.Name = &name
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithPageLength(pageLength int) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.PageLength = &pageLength
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithPageOffset(pageOffset int) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.PageOffset = &pageOffset
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.Quiesced = quiesced
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithSerialNumber(serialNumber string) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.SerialNumber = &serialNumber
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

func (getTapePartitionsWithFullDetailsSpectraS3Request *GetTapePartitionsWithFullDetailsSpectraS3Request) WithState(state TapePartitionState) *GetTapePartitionsWithFullDetailsSpectraS3Request {
    getTapePartitionsWithFullDetailsSpectraS3Request.State = state
    return getTapePartitionsWithFullDetailsSpectraS3Request
}

type GetTapeSpectraS3Request struct {
    TapeId string
}

func NewGetTapeSpectraS3Request(tapeId string) *GetTapeSpectraS3Request {
    return &GetTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

type GetTapesSpectraS3Request struct {
    AssignedToStorageDomain *bool
    AvailableRawCapacity *int64
    BarCode *string
    BucketId *string
    EjectLabel *string
    EjectLocation *string
    FullOfData *bool
    LastPage bool
    LastVerified *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PartiallyVerifiedEndOfTape *string
    PartitionId *string
    PreviousState TapeState
    SerialNumber *string
    SortBy *string
    State TapeState
    StorageDomainMemberId *string
    String *string
    VerifyPending Priority
    WriteProtected *bool
}

func NewGetTapesSpectraS3Request() *GetTapesSpectraS3Request {
    return &GetTapesSpectraS3Request{
    }
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithAssignedToStorageDomain(assignedToStorageDomain bool) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.AssignedToStorageDomain = &assignedToStorageDomain
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithAvailableRawCapacity(availableRawCapacity int64) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.AvailableRawCapacity = &availableRawCapacity
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithBarCode(barCode string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.BarCode = &barCode
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithBucketId(bucketId string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.BucketId = &bucketId
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithEjectLabel(ejectLabel string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.EjectLabel = &ejectLabel
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithEjectLocation(ejectLocation string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.EjectLocation = &ejectLocation
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithFullOfData(fullOfData bool) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.FullOfData = &fullOfData
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithLastPage() *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.LastPage = true
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithLastVerified(lastVerified string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.LastVerified = &lastVerified
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPageLength(pageLength int) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PageLength = &pageLength
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPageOffset(pageOffset int) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PageOffset = &pageOffset
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPartiallyVerifiedEndOfTape(partiallyVerifiedEndOfTape string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PartiallyVerifiedEndOfTape = &partiallyVerifiedEndOfTape
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPartitionId(partitionId string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PartitionId = &partitionId
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithPreviousState(previousState TapeState) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.PreviousState = previousState
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithSerialNumber(serialNumber string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.SerialNumber = &serialNumber
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithSortBy(sortBy string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.SortBy = &sortBy
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithState(state TapeState) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.State = state
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithStorageDomainMemberId(storageDomainMemberId string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.StorageDomainMemberId = &storageDomainMemberId
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithString(string string) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.String = &string
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithVerifyPending(verifyPending Priority) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.VerifyPending = verifyPending
    return getTapesSpectraS3Request
}

func (getTapesSpectraS3Request *GetTapesSpectraS3Request) WithWriteProtected(writeProtected bool) *GetTapesSpectraS3Request {
    getTapesSpectraS3Request.WriteProtected = &writeProtected
    return getTapesSpectraS3Request
}

type ImportAllTapesSpectraS3Request struct {
    DataPolicyId *string
    Priority Priority
    StorageDomainId *string
    UserId *string
    VerifyDataAfterImport Priority
    VerifyDataPriorToImport *bool
}

func NewImportAllTapesSpectraS3Request() *ImportAllTapesSpectraS3Request {
    return &ImportAllTapesSpectraS3Request{
    }
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.DataPolicyId = &dataPolicyId
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithPriority(priority Priority) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.Priority = priority
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.StorageDomainId = &storageDomainId
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithUserId(userId string) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.UserId = &userId
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.VerifyDataAfterImport = verifyDataAfterImport
    return importAllTapesSpectraS3Request
}

func (importAllTapesSpectraS3Request *ImportAllTapesSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportAllTapesSpectraS3Request {
    importAllTapesSpectraS3Request.VerifyDataPriorToImport = &verifyDataPriorToImport
    return importAllTapesSpectraS3Request
}

type ImportTapeSpectraS3Request struct {
    DataPolicyId *string
    Priority Priority
    StorageDomainId *string
    TapeId string
    UserId *string
    VerifyDataAfterImport Priority
    VerifyDataPriorToImport *bool
}

func NewImportTapeSpectraS3Request(tapeId string) *ImportTapeSpectraS3Request {
    return &ImportTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.DataPolicyId = &dataPolicyId
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithPriority(priority Priority) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.Priority = priority
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithStorageDomainId(storageDomainId string) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.StorageDomainId = &storageDomainId
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithUserId(userId string) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.UserId = &userId
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithVerifyDataAfterImport(verifyDataAfterImport Priority) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.VerifyDataAfterImport = verifyDataAfterImport
    return importTapeSpectraS3Request
}

func (importTapeSpectraS3Request *ImportTapeSpectraS3Request) WithVerifyDataPriorToImport(verifyDataPriorToImport bool) *ImportTapeSpectraS3Request {
    importTapeSpectraS3Request.VerifyDataPriorToImport = &verifyDataPriorToImport
    return importTapeSpectraS3Request
}

type InspectAllTapesSpectraS3Request struct {
    TaskPriority Priority
}

func NewInspectAllTapesSpectraS3Request() *InspectAllTapesSpectraS3Request {
    return &InspectAllTapesSpectraS3Request{
    }
}

func (inspectAllTapesSpectraS3Request *InspectAllTapesSpectraS3Request) WithTaskPriority(taskPriority Priority) *InspectAllTapesSpectraS3Request {
    inspectAllTapesSpectraS3Request.TaskPriority = taskPriority
    return inspectAllTapesSpectraS3Request
}

type InspectTapeSpectraS3Request struct {
    TapeId string
    TaskPriority Priority
}

func NewInspectTapeSpectraS3Request(tapeId string) *InspectTapeSpectraS3Request {
    return &InspectTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (inspectTapeSpectraS3Request *InspectTapeSpectraS3Request) WithTaskPriority(taskPriority Priority) *InspectTapeSpectraS3Request {
    inspectTapeSpectraS3Request.TaskPriority = taskPriority
    return inspectTapeSpectraS3Request
}

type ModifyAllTapePartitionsSpectraS3Request struct {
    Quiesced Quiesced
}

func NewModifyAllTapePartitionsSpectraS3Request(quiesced Quiesced) *ModifyAllTapePartitionsSpectraS3Request {
    return &ModifyAllTapePartitionsSpectraS3Request{
        Quiesced: quiesced,
    }
}

type ModifyTapeDriveSpectraS3Request struct {
    MinimumTaskPriority Priority
    Quiesced Quiesced
    ReservedTaskType ReservedTaskType
    TapeDriveId string
}

func NewModifyTapeDriveSpectraS3Request(tapeDriveId string) *ModifyTapeDriveSpectraS3Request {
    return &ModifyTapeDriveSpectraS3Request{
        TapeDriveId: tapeDriveId,
    }
}

func (modifyTapeDriveSpectraS3Request *ModifyTapeDriveSpectraS3Request) WithMinimumTaskPriority(minimumTaskPriority Priority) *ModifyTapeDriveSpectraS3Request {
    modifyTapeDriveSpectraS3Request.MinimumTaskPriority = minimumTaskPriority
    return modifyTapeDriveSpectraS3Request
}

func (modifyTapeDriveSpectraS3Request *ModifyTapeDriveSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyTapeDriveSpectraS3Request {
    modifyTapeDriveSpectraS3Request.Quiesced = quiesced
    return modifyTapeDriveSpectraS3Request
}

func (modifyTapeDriveSpectraS3Request *ModifyTapeDriveSpectraS3Request) WithReservedTaskType(reservedTaskType ReservedTaskType) *ModifyTapeDriveSpectraS3Request {
    modifyTapeDriveSpectraS3Request.ReservedTaskType = reservedTaskType
    return modifyTapeDriveSpectraS3Request
}

type ModifyTapePartitionSpectraS3Request struct {
    AutoCompactionEnabled *bool
    MinimumReadReservedDrives *int
    MinimumWriteReservedDrives *int
    Quiesced Quiesced
    SerialNumber *string
    TapePartition string
}

func NewModifyTapePartitionSpectraS3Request(tapePartition string) *ModifyTapePartitionSpectraS3Request {
    return &ModifyTapePartitionSpectraS3Request{
        TapePartition: tapePartition,
    }
}

func (modifyTapePartitionSpectraS3Request *ModifyTapePartitionSpectraS3Request) WithAutoCompactionEnabled(autoCompactionEnabled bool) *ModifyTapePartitionSpectraS3Request {
    modifyTapePartitionSpectraS3Request.AutoCompactionEnabled = &autoCompactionEnabled
    return modifyTapePartitionSpectraS3Request
}

func (modifyTapePartitionSpectraS3Request *ModifyTapePartitionSpectraS3Request) WithMinimumReadReservedDrives(minimumReadReservedDrives int) *ModifyTapePartitionSpectraS3Request {
    modifyTapePartitionSpectraS3Request.MinimumReadReservedDrives = &minimumReadReservedDrives
    return modifyTapePartitionSpectraS3Request
}

func (modifyTapePartitionSpectraS3Request *ModifyTapePartitionSpectraS3Request) WithMinimumWriteReservedDrives(minimumWriteReservedDrives int) *ModifyTapePartitionSpectraS3Request {
    modifyTapePartitionSpectraS3Request.MinimumWriteReservedDrives = &minimumWriteReservedDrives
    return modifyTapePartitionSpectraS3Request
}

func (modifyTapePartitionSpectraS3Request *ModifyTapePartitionSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyTapePartitionSpectraS3Request {
    modifyTapePartitionSpectraS3Request.Quiesced = quiesced
    return modifyTapePartitionSpectraS3Request
}

func (modifyTapePartitionSpectraS3Request *ModifyTapePartitionSpectraS3Request) WithSerialNumber(serialNumber string) *ModifyTapePartitionSpectraS3Request {
    modifyTapePartitionSpectraS3Request.SerialNumber = &serialNumber
    return modifyTapePartitionSpectraS3Request
}

type ModifyTapeSpectraS3Request struct {
    EjectLabel *string
    EjectLocation *string
    State TapeState
    TapeId string
}

func NewModifyTapeSpectraS3Request(tapeId string) *ModifyTapeSpectraS3Request {
    return &ModifyTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) WithEjectLabel(ejectLabel string) *ModifyTapeSpectraS3Request {
    modifyTapeSpectraS3Request.EjectLabel = &ejectLabel
    return modifyTapeSpectraS3Request
}

func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) WithEjectLocation(ejectLocation string) *ModifyTapeSpectraS3Request {
    modifyTapeSpectraS3Request.EjectLocation = &ejectLocation
    return modifyTapeSpectraS3Request
}

func (modifyTapeSpectraS3Request *ModifyTapeSpectraS3Request) WithState(state TapeState) *ModifyTapeSpectraS3Request {
    modifyTapeSpectraS3Request.State = state
    return modifyTapeSpectraS3Request
}

type OnlineAllTapesSpectraS3Request struct {
}

func NewOnlineAllTapesSpectraS3Request() *OnlineAllTapesSpectraS3Request {
    return &OnlineAllTapesSpectraS3Request{
    }
}

type OnlineTapeSpectraS3Request struct {
    TapeId string
}

func NewOnlineTapeSpectraS3Request(tapeId string) *OnlineTapeSpectraS3Request {
    return &OnlineTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

type RawImportAllTapesSpectraS3Request struct {
    BucketId string
    StorageDomainId *string
    TaskPriority Priority
}

func NewRawImportAllTapesSpectraS3Request(bucketId string) *RawImportAllTapesSpectraS3Request {
    return &RawImportAllTapesSpectraS3Request{
        BucketId: bucketId,
    }
}

func (rawImportAllTapesSpectraS3Request *RawImportAllTapesSpectraS3Request) WithStorageDomainId(storageDomainId string) *RawImportAllTapesSpectraS3Request {
    rawImportAllTapesSpectraS3Request.StorageDomainId = &storageDomainId
    return rawImportAllTapesSpectraS3Request
}

func (rawImportAllTapesSpectraS3Request *RawImportAllTapesSpectraS3Request) WithTaskPriority(taskPriority Priority) *RawImportAllTapesSpectraS3Request {
    rawImportAllTapesSpectraS3Request.TaskPriority = taskPriority
    return rawImportAllTapesSpectraS3Request
}

type RawImportTapeSpectraS3Request struct {
    BucketId string
    StorageDomainId *string
    TapeId string
    TaskPriority Priority
}

func NewRawImportTapeSpectraS3Request(bucketId string, tapeId string) *RawImportTapeSpectraS3Request {
    return &RawImportTapeSpectraS3Request{
        TapeId: tapeId,
        BucketId: bucketId,
    }
}

func (rawImportTapeSpectraS3Request *RawImportTapeSpectraS3Request) WithStorageDomainId(storageDomainId string) *RawImportTapeSpectraS3Request {
    rawImportTapeSpectraS3Request.StorageDomainId = &storageDomainId
    return rawImportTapeSpectraS3Request
}

func (rawImportTapeSpectraS3Request *RawImportTapeSpectraS3Request) WithTaskPriority(taskPriority Priority) *RawImportTapeSpectraS3Request {
    rawImportTapeSpectraS3Request.TaskPriority = taskPriority
    return rawImportTapeSpectraS3Request
}

type VerifyAllTapesSpectraS3Request struct {
    TaskPriority Priority
}

func NewVerifyAllTapesSpectraS3Request() *VerifyAllTapesSpectraS3Request {
    return &VerifyAllTapesSpectraS3Request{
    }
}

func (verifyAllTapesSpectraS3Request *VerifyAllTapesSpectraS3Request) WithTaskPriority(taskPriority Priority) *VerifyAllTapesSpectraS3Request {
    verifyAllTapesSpectraS3Request.TaskPriority = taskPriority
    return verifyAllTapesSpectraS3Request
}

type VerifyTapeSpectraS3Request struct {
    TapeId string
    TaskPriority Priority
}

func NewVerifyTapeSpectraS3Request(tapeId string) *VerifyTapeSpectraS3Request {
    return &VerifyTapeSpectraS3Request{
        TapeId: tapeId,
    }
}

func (verifyTapeSpectraS3Request *VerifyTapeSpectraS3Request) WithTaskPriority(taskPriority Priority) *VerifyTapeSpectraS3Request {
    verifyTapeSpectraS3Request.TaskPriority = taskPriority
    return verifyTapeSpectraS3Request
}

type ForceTargetEnvironmentRefreshSpectraS3Request struct {
}

func NewForceTargetEnvironmentRefreshSpectraS3Request() *ForceTargetEnvironmentRefreshSpectraS3Request {
    return &ForceTargetEnvironmentRefreshSpectraS3Request{
    }
}

type PutAzureTargetBucketNameSpectraS3Request struct {
    BucketId string
    Name string
    TargetId string
}

func NewPutAzureTargetBucketNameSpectraS3Request(bucketId string, name string, targetId string) *PutAzureTargetBucketNameSpectraS3Request {
    return &PutAzureTargetBucketNameSpectraS3Request{
        BucketId: bucketId,
        Name: name,
        TargetId: targetId,
    }
}

type PutAzureTargetReadPreferenceSpectraS3Request struct {
    BucketId string
    ReadPreference TargetReadPreferenceType
    TargetId string
}

func NewPutAzureTargetReadPreferenceSpectraS3Request(bucketId string, readPreference TargetReadPreferenceType, targetId string) *PutAzureTargetReadPreferenceSpectraS3Request {
    return &PutAzureTargetReadPreferenceSpectraS3Request{
        BucketId: bucketId,
        ReadPreference: readPreference,
        TargetId: targetId,
    }
}

type DeleteAzureTargetBucketNameSpectraS3Request struct {
    AzureTargetBucketName string
}

func NewDeleteAzureTargetBucketNameSpectraS3Request(azureTargetBucketName string) *DeleteAzureTargetBucketNameSpectraS3Request {
    return &DeleteAzureTargetBucketNameSpectraS3Request{
        AzureTargetBucketName: azureTargetBucketName,
    }
}

type DeleteAzureTargetFailureSpectraS3Request struct {
    AzureTargetFailure string
}

func NewDeleteAzureTargetFailureSpectraS3Request(azureTargetFailure string) *DeleteAzureTargetFailureSpectraS3Request {
    return &DeleteAzureTargetFailureSpectraS3Request{
        AzureTargetFailure: azureTargetFailure,
    }
}

type DeleteAzureTargetReadPreferenceSpectraS3Request struct {
    AzureTargetReadPreference string
}

func NewDeleteAzureTargetReadPreferenceSpectraS3Request(azureTargetReadPreference string) *DeleteAzureTargetReadPreferenceSpectraS3Request {
    return &DeleteAzureTargetReadPreferenceSpectraS3Request{
        AzureTargetReadPreference: azureTargetReadPreference,
    }
}

type DeleteAzureTargetSpectraS3Request struct {
    AzureTarget string
}

func NewDeleteAzureTargetSpectraS3Request(azureTarget string) *DeleteAzureTargetSpectraS3Request {
    return &DeleteAzureTargetSpectraS3Request{
        AzureTarget: azureTarget,
    }
}

type GetAzureTargetBucketNamesSpectraS3Request struct {
    BucketId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetId *string
}

func NewGetAzureTargetBucketNamesSpectraS3Request() *GetAzureTargetBucketNamesSpectraS3Request {
    return &GetAzureTargetBucketNamesSpectraS3Request{
    }
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithBucketId(bucketId string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.BucketId = &bucketId
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithLastPage() *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.LastPage = true
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithName(name string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.Name = &name
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.PageLength = &pageLength
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.PageOffset = &pageOffset
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureTargetBucketNamesSpectraS3Request
}

func (getAzureTargetBucketNamesSpectraS3Request *GetAzureTargetBucketNamesSpectraS3Request) WithTargetId(targetId string) *GetAzureTargetBucketNamesSpectraS3Request {
    getAzureTargetBucketNamesSpectraS3Request.TargetId = &targetId
    return getAzureTargetBucketNamesSpectraS3Request
}

type GetAzureTargetFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetFailureType TargetFailureType
    TargetId *string
}

func NewGetAzureTargetFailuresSpectraS3Request() *GetAzureTargetFailuresSpectraS3Request {
    return &GetAzureTargetFailuresSpectraS3Request{
    }
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithLastPage() *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.LastPage = true
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.PageLength = &pageLength
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.PageOffset = &pageOffset
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithTargetId(targetId string) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.TargetId = &targetId
    return getAzureTargetFailuresSpectraS3Request
}

func (getAzureTargetFailuresSpectraS3Request *GetAzureTargetFailuresSpectraS3Request) WithTargetFailureType(targetFailureType TargetFailureType) *GetAzureTargetFailuresSpectraS3Request {
    getAzureTargetFailuresSpectraS3Request.TargetFailureType = targetFailureType
    return getAzureTargetFailuresSpectraS3Request
}

type GetAzureTargetReadPreferenceSpectraS3Request struct {
    AzureTargetReadPreference string
}

func NewGetAzureTargetReadPreferenceSpectraS3Request(azureTargetReadPreference string) *GetAzureTargetReadPreferenceSpectraS3Request {
    return &GetAzureTargetReadPreferenceSpectraS3Request{
        AzureTargetReadPreference: azureTargetReadPreference,
    }
}

type GetAzureTargetReadPreferencesSpectraS3Request struct {
    BucketId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    ReadPreference TargetReadPreferenceType
    TargetId *string
}

func NewGetAzureTargetReadPreferencesSpectraS3Request() *GetAzureTargetReadPreferencesSpectraS3Request {
    return &GetAzureTargetReadPreferencesSpectraS3Request{
    }
}

func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithBucketId(bucketId string) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.BucketId = &bucketId
    return getAzureTargetReadPreferencesSpectraS3Request
}

func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithLastPage() *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.LastPage = true
    return getAzureTargetReadPreferencesSpectraS3Request
}

func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.PageLength = &pageLength
    return getAzureTargetReadPreferencesSpectraS3Request
}

func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.PageOffset = &pageOffset
    return getAzureTargetReadPreferencesSpectraS3Request
}

func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureTargetReadPreferencesSpectraS3Request
}

func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithReadPreference(readPreference TargetReadPreferenceType) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.ReadPreference = readPreference
    return getAzureTargetReadPreferencesSpectraS3Request
}

func (getAzureTargetReadPreferencesSpectraS3Request *GetAzureTargetReadPreferencesSpectraS3Request) WithTargetId(targetId string) *GetAzureTargetReadPreferencesSpectraS3Request {
    getAzureTargetReadPreferencesSpectraS3Request.TargetId = &targetId
    return getAzureTargetReadPreferencesSpectraS3Request
}

type GetAzureTargetSpectraS3Request struct {
    AzureTarget string
}

func NewGetAzureTargetSpectraS3Request(azureTarget string) *GetAzureTargetSpectraS3Request {
    return &GetAzureTargetSpectraS3Request{
        AzureTarget: azureTarget,
    }
}

type GetAzureTargetsSpectraS3Request struct {
    AccountName *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PermitGoingOutOfSync *bool
    Quiesced Quiesced
    State TargetState
}

func NewGetAzureTargetsSpectraS3Request() *GetAzureTargetsSpectraS3Request {
    return &GetAzureTargetsSpectraS3Request{
    }
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithAccountName(accountName string) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.AccountName = &accountName
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithHttps(https bool) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.Https = &https
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithLastPage() *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.LastPage = true
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithName(name string) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.Name = &name
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPageLength(pageLength int) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.PageLength = &pageLength
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.PageOffset = &pageOffset
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.Quiesced = quiesced
    return getAzureTargetsSpectraS3Request
}

func (getAzureTargetsSpectraS3Request *GetAzureTargetsSpectraS3Request) WithState(state TargetState) *GetAzureTargetsSpectraS3Request {
    getAzureTargetsSpectraS3Request.State = state
    return getAzureTargetsSpectraS3Request
}

type GetBlobsOnAzureTargetSpectraS3Request struct {
    AzureTarget string
}

func NewGetBlobsOnAzureTargetSpectraS3Request(azureTarget string) *GetBlobsOnAzureTargetSpectraS3Request {
    return &GetBlobsOnAzureTargetSpectraS3Request{
        AzureTarget: azureTarget,
    }
}

type ImportAzureTargetSpectraS3Request struct {
    AzureTarget string
    CloudBucketName string
    DataPolicyId *string
    Priority Priority
    UserId *string
}

func NewImportAzureTargetSpectraS3Request(azureTarget string, cloudBucketName string) *ImportAzureTargetSpectraS3Request {
    return &ImportAzureTargetSpectraS3Request{
        AzureTarget: azureTarget,
        CloudBucketName: cloudBucketName,
    }
}

func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.DataPolicyId = &dataPolicyId
    return importAzureTargetSpectraS3Request
}

func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithPriority(priority Priority) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.Priority = priority
    return importAzureTargetSpectraS3Request
}

func (importAzureTargetSpectraS3Request *ImportAzureTargetSpectraS3Request) WithUserId(userId string) *ImportAzureTargetSpectraS3Request {
    importAzureTargetSpectraS3Request.UserId = &userId
    return importAzureTargetSpectraS3Request
}

type ModifyAllAzureTargetsSpectraS3Request struct {
    Quiesced Quiesced
}

func NewModifyAllAzureTargetsSpectraS3Request(quiesced Quiesced) *ModifyAllAzureTargetsSpectraS3Request {
    return &ModifyAllAzureTargetsSpectraS3Request{
        Quiesced: quiesced,
    }
}

type ModifyAzureTargetSpectraS3Request struct {
    AccountKey *string
    AccountName *string
    AutoVerifyFrequencyInDays *int
    AzureTarget string
    CloudBucketPrefix *string
    CloudBucketSuffix *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    Name *string
    PermitGoingOutOfSync *bool
    Quiesced Quiesced
}

func NewModifyAzureTargetSpectraS3Request(azureTarget string) *ModifyAzureTargetSpectraS3Request {
    return &ModifyAzureTargetSpectraS3Request{
        AzureTarget: azureTarget,
    }
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithAccountKey(accountKey string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.AccountKey = &accountKey
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithAccountName(accountName string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.AccountName = &accountName
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays int) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.AutoVerifyFrequencyInDays = &autoVerifyFrequencyInDays
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.CloudBucketPrefix = &cloudBucketPrefix
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.CloudBucketSuffix = &cloudBucketSuffix
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithHttps(https bool) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.Https = &https
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithName(name string) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.Name = &name
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return modifyAzureTargetSpectraS3Request
}

func (modifyAzureTargetSpectraS3Request *ModifyAzureTargetSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyAzureTargetSpectraS3Request {
    modifyAzureTargetSpectraS3Request.Quiesced = quiesced
    return modifyAzureTargetSpectraS3Request
}

type RegisterAzureTargetSpectraS3Request struct {
    AccountKey string
    AccountName string
    AutoVerifyFrequencyInDays *int
    CloudBucketPrefix *string
    CloudBucketSuffix *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    Name string
    PermitGoingOutOfSync *bool
}

func NewRegisterAzureTargetSpectraS3Request(accountKey string, accountName string, name string) *RegisterAzureTargetSpectraS3Request {
    return &RegisterAzureTargetSpectraS3Request{
        AccountKey: accountKey,
        AccountName: accountName,
        Name: name,
    }
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays int) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.AutoVerifyFrequencyInDays = &autoVerifyFrequencyInDays
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix string) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.CloudBucketPrefix = &cloudBucketPrefix
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix string) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.CloudBucketSuffix = &cloudBucketSuffix
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithHttps(https bool) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.Https = &https
    return registerAzureTargetSpectraS3Request
}

func (registerAzureTargetSpectraS3Request *RegisterAzureTargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *RegisterAzureTargetSpectraS3Request {
    registerAzureTargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return registerAzureTargetSpectraS3Request
}

type VerifyAzureTargetSpectraS3Request struct {
    AzureTarget string
    FullDetails bool
}

func NewVerifyAzureTargetSpectraS3Request(azureTarget string) *VerifyAzureTargetSpectraS3Request {
    return &VerifyAzureTargetSpectraS3Request{
        AzureTarget: azureTarget,
    }
}

func (verifyAzureTargetSpectraS3Request *VerifyAzureTargetSpectraS3Request) WithFullDetails() *VerifyAzureTargetSpectraS3Request {
    verifyAzureTargetSpectraS3Request.FullDetails = true
    return verifyAzureTargetSpectraS3Request
}

type PutDs3TargetReadPreferenceSpectraS3Request struct {
    BucketId string
    ReadPreference TargetReadPreferenceType
    TargetId string
}

func NewPutDs3TargetReadPreferenceSpectraS3Request(bucketId string, readPreference TargetReadPreferenceType, targetId string) *PutDs3TargetReadPreferenceSpectraS3Request {
    return &PutDs3TargetReadPreferenceSpectraS3Request{
        BucketId: bucketId,
        ReadPreference: readPreference,
        TargetId: targetId,
    }
}

type DeleteDs3TargetFailureSpectraS3Request struct {
    Ds3TargetFailure string
}

func NewDeleteDs3TargetFailureSpectraS3Request(ds3TargetFailure string) *DeleteDs3TargetFailureSpectraS3Request {
    return &DeleteDs3TargetFailureSpectraS3Request{
        Ds3TargetFailure: ds3TargetFailure,
    }
}

type DeleteDs3TargetReadPreferenceSpectraS3Request struct {
    Ds3TargetReadPreference string
}

func NewDeleteDs3TargetReadPreferenceSpectraS3Request(ds3TargetReadPreference string) *DeleteDs3TargetReadPreferenceSpectraS3Request {
    return &DeleteDs3TargetReadPreferenceSpectraS3Request{
        Ds3TargetReadPreference: ds3TargetReadPreference,
    }
}

type DeleteDs3TargetSpectraS3Request struct {
    Ds3Target string
}

func NewDeleteDs3TargetSpectraS3Request(ds3Target string) *DeleteDs3TargetSpectraS3Request {
    return &DeleteDs3TargetSpectraS3Request{
        Ds3Target: ds3Target,
    }
}

type GetBlobsOnDs3TargetSpectraS3Request struct {
    Ds3Target string
}

func NewGetBlobsOnDs3TargetSpectraS3Request(ds3Target string) *GetBlobsOnDs3TargetSpectraS3Request {
    return &GetBlobsOnDs3TargetSpectraS3Request{
        Ds3Target: ds3Target,
    }
}

type GetDs3TargetDataPoliciesSpectraS3Request struct {
    Ds3TargetDataPolicies string
}

func NewGetDs3TargetDataPoliciesSpectraS3Request(ds3TargetDataPolicies string) *GetDs3TargetDataPoliciesSpectraS3Request {
    return &GetDs3TargetDataPoliciesSpectraS3Request{
        Ds3TargetDataPolicies: ds3TargetDataPolicies,
    }
}

type GetDs3TargetFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetFailureType TargetFailureType
    TargetId *string
}

func NewGetDs3TargetFailuresSpectraS3Request() *GetDs3TargetFailuresSpectraS3Request {
    return &GetDs3TargetFailuresSpectraS3Request{
    }
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithLastPage() *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.LastPage = true
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.PageLength = &pageLength
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.PageOffset = &pageOffset
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithTargetId(targetId string) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.TargetId = &targetId
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithTargetFailureType(targetFailureType TargetFailureType) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.TargetFailureType = targetFailureType
    return getDs3TargetFailuresSpectraS3Request
}

type GetDs3TargetReadPreferenceSpectraS3Request struct {
    Ds3TargetReadPreference string
}

func NewGetDs3TargetReadPreferenceSpectraS3Request(ds3TargetReadPreference string) *GetDs3TargetReadPreferenceSpectraS3Request {
    return &GetDs3TargetReadPreferenceSpectraS3Request{
        Ds3TargetReadPreference: ds3TargetReadPreference,
    }
}

type GetDs3TargetReadPreferencesSpectraS3Request struct {
    BucketId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    ReadPreference TargetReadPreferenceType
    TargetId *string
}

func NewGetDs3TargetReadPreferencesSpectraS3Request() *GetDs3TargetReadPreferencesSpectraS3Request {
    return &GetDs3TargetReadPreferencesSpectraS3Request{
    }
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithBucketId(bucketId string) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.BucketId = &bucketId
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithLastPage() *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.LastPage = true
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.PageLength = &pageLength
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.PageOffset = &pageOffset
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithReadPreference(readPreference TargetReadPreferenceType) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.ReadPreference = readPreference
    return getDs3TargetReadPreferencesSpectraS3Request
}

func (getDs3TargetReadPreferencesSpectraS3Request *GetDs3TargetReadPreferencesSpectraS3Request) WithTargetId(targetId string) *GetDs3TargetReadPreferencesSpectraS3Request {
    getDs3TargetReadPreferencesSpectraS3Request.TargetId = &targetId
    return getDs3TargetReadPreferencesSpectraS3Request
}

type GetDs3TargetSpectraS3Request struct {
    Ds3Target string
}

func NewGetDs3TargetSpectraS3Request(ds3Target string) *GetDs3TargetSpectraS3Request {
    return &GetDs3TargetSpectraS3Request{
        Ds3Target: ds3Target,
    }
}

type GetDs3TargetsSpectraS3Request struct {
    AdminAuthId *string
    DataPathEndPoint *string
    DataPathHttps *bool
    DataPathPort *int
    DataPathProxy *string
    DataPathVerifyCertificate *bool
    DefaultReadPreference TargetReadPreferenceType
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PermitGoingOutOfSync *bool
    Quiesced Quiesced
    State TargetState
}

func NewGetDs3TargetsSpectraS3Request() *GetDs3TargetsSpectraS3Request {
    return &GetDs3TargetsSpectraS3Request{
    }
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithAdminAuthId(adminAuthId string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.AdminAuthId = &adminAuthId
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DataPathHttps = &dataPathHttps
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathPort(dataPathPort int) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DataPathPort = &dataPathPort
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathProxy(dataPathProxy string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DataPathProxy = &dataPathProxy
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DataPathVerifyCertificate = &dataPathVerifyCertificate
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithLastPage() *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.LastPage = true
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithName(name string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.Name = &name
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.PageLength = &pageLength
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.PageOffset = &pageOffset
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.Quiesced = quiesced
    return getDs3TargetsSpectraS3Request
}

func (getDs3TargetsSpectraS3Request *GetDs3TargetsSpectraS3Request) WithState(state TargetState) *GetDs3TargetsSpectraS3Request {
    getDs3TargetsSpectraS3Request.State = state
    return getDs3TargetsSpectraS3Request
}

type ModifyAllDs3TargetsSpectraS3Request struct {
    Quiesced Quiesced
}

func NewModifyAllDs3TargetsSpectraS3Request(quiesced Quiesced) *ModifyAllDs3TargetsSpectraS3Request {
    return &ModifyAllDs3TargetsSpectraS3Request{
        Quiesced: quiesced,
    }
}

type ModifyDs3TargetSpectraS3Request struct {
    AccessControlReplication Ds3TargetAccessControlReplication
    AdminAuthId *string
    AdminSecretKey *string
    DataPathEndPoint *string
    DataPathHttps *bool
    DataPathPort *int
    DataPathProxy *string
    DataPathVerifyCertificate *bool
    DefaultReadPreference TargetReadPreferenceType
    Ds3Target string
    Name *string
    PermitGoingOutOfSync *bool
    Quiesced Quiesced
    ReplicatedUserDefaultDataPolicy *string
}

func NewModifyDs3TargetSpectraS3Request(ds3Target string) *ModifyDs3TargetSpectraS3Request {
    return &ModifyDs3TargetSpectraS3Request{
        Ds3Target: ds3Target,
    }
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithAccessControlReplication(accessControlReplication Ds3TargetAccessControlReplication) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.AccessControlReplication = accessControlReplication
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithAdminAuthId(adminAuthId string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.AdminAuthId = &adminAuthId
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithAdminSecretKey(adminSecretKey string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.AdminSecretKey = &adminSecretKey
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DataPathHttps = &dataPathHttps
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathPort(dataPathPort int) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DataPathPort = &dataPathPort
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathProxy(dataPathProxy string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DataPathProxy = &dataPathProxy
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DataPathVerifyCertificate = &dataPathVerifyCertificate
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithName(name string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.Name = &name
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.Quiesced = quiesced
    return modifyDs3TargetSpectraS3Request
}

func (modifyDs3TargetSpectraS3Request *ModifyDs3TargetSpectraS3Request) WithReplicatedUserDefaultDataPolicy(replicatedUserDefaultDataPolicy string) *ModifyDs3TargetSpectraS3Request {
    modifyDs3TargetSpectraS3Request.ReplicatedUserDefaultDataPolicy = &replicatedUserDefaultDataPolicy
    return modifyDs3TargetSpectraS3Request
}

type PairBackRegisteredDs3TargetSpectraS3Request struct {
    AccessControlReplication Ds3TargetAccessControlReplication
    AdminAuthId *string
    AdminSecretKey *string
    DataPathEndPoint *string
    DataPathHttps *bool
    DataPathPort *int
    DataPathProxy *string
    DataPathVerifyCertificate *bool
    DefaultReadPreference TargetReadPreferenceType
    Ds3Target string
    Name *string
    PermitGoingOutOfSync *bool
    ReplicatedUserDefaultDataPolicy *string
}

func NewPairBackRegisteredDs3TargetSpectraS3Request(ds3Target string) *PairBackRegisteredDs3TargetSpectraS3Request {
    return &PairBackRegisteredDs3TargetSpectraS3Request{
        Ds3Target: ds3Target,
    }
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithAccessControlReplication(accessControlReplication Ds3TargetAccessControlReplication) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.AccessControlReplication = accessControlReplication
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithAdminAuthId(adminAuthId string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.AdminAuthId = &adminAuthId
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithAdminSecretKey(adminSecretKey string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.AdminSecretKey = &adminSecretKey
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DataPathHttps = &dataPathHttps
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathPort(dataPathPort int) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DataPathPort = &dataPathPort
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathProxy(dataPathProxy string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DataPathProxy = &dataPathProxy
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DataPathVerifyCertificate = &dataPathVerifyCertificate
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithName(name string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.Name = &name
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return pairBackRegisteredDs3TargetSpectraS3Request
}

func (pairBackRegisteredDs3TargetSpectraS3Request *PairBackRegisteredDs3TargetSpectraS3Request) WithReplicatedUserDefaultDataPolicy(replicatedUserDefaultDataPolicy string) *PairBackRegisteredDs3TargetSpectraS3Request {
    pairBackRegisteredDs3TargetSpectraS3Request.ReplicatedUserDefaultDataPolicy = &replicatedUserDefaultDataPolicy
    return pairBackRegisteredDs3TargetSpectraS3Request
}

type RegisterDs3TargetSpectraS3Request struct {
    AccessControlReplication Ds3TargetAccessControlReplication
    AdminAuthId string
    AdminSecretKey string
    DataPathEndPoint string
    DataPathHttps *bool
    DataPathPort *int
    DataPathProxy *string
    DataPathVerifyCertificate *bool
    DefaultReadPreference TargetReadPreferenceType
    Name string
    PermitGoingOutOfSync *bool
    ReplicatedUserDefaultDataPolicy *string
}

func NewRegisterDs3TargetSpectraS3Request(adminAuthId string, adminSecretKey string, dataPathEndPoint string, name string) *RegisterDs3TargetSpectraS3Request {
    return &RegisterDs3TargetSpectraS3Request{
        AdminAuthId: adminAuthId,
        AdminSecretKey: adminSecretKey,
        DataPathEndPoint: dataPathEndPoint,
        Name: name,
    }
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithAccessControlReplication(accessControlReplication Ds3TargetAccessControlReplication) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.AccessControlReplication = accessControlReplication
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathHttps(dataPathHttps bool) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.DataPathHttps = &dataPathHttps
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathPort(dataPathPort int) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.DataPathPort = &dataPathPort
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathProxy(dataPathProxy string) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.DataPathProxy = &dataPathProxy
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDataPathVerifyCertificate(dataPathVerifyCertificate bool) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.DataPathVerifyCertificate = &dataPathVerifyCertificate
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return registerDs3TargetSpectraS3Request
}

func (registerDs3TargetSpectraS3Request *RegisterDs3TargetSpectraS3Request) WithReplicatedUserDefaultDataPolicy(replicatedUserDefaultDataPolicy string) *RegisterDs3TargetSpectraS3Request {
    registerDs3TargetSpectraS3Request.ReplicatedUserDefaultDataPolicy = &replicatedUserDefaultDataPolicy
    return registerDs3TargetSpectraS3Request
}

type VerifyDs3TargetSpectraS3Request struct {
    Ds3Target string
    FullDetails bool
}

func NewVerifyDs3TargetSpectraS3Request(ds3Target string) *VerifyDs3TargetSpectraS3Request {
    return &VerifyDs3TargetSpectraS3Request{
        Ds3Target: ds3Target,
    }
}

func (verifyDs3TargetSpectraS3Request *VerifyDs3TargetSpectraS3Request) WithFullDetails() *VerifyDs3TargetSpectraS3Request {
    verifyDs3TargetSpectraS3Request.FullDetails = true
    return verifyDs3TargetSpectraS3Request
}

type PutS3TargetBucketNameSpectraS3Request struct {
    BucketId string
    Name string
    TargetId string
}

func NewPutS3TargetBucketNameSpectraS3Request(bucketId string, name string, targetId string) *PutS3TargetBucketNameSpectraS3Request {
    return &PutS3TargetBucketNameSpectraS3Request{
        BucketId: bucketId,
        Name: name,
        TargetId: targetId,
    }
}

type PutS3TargetReadPreferenceSpectraS3Request struct {
    BucketId string
    ReadPreference TargetReadPreferenceType
    TargetId string
}

func NewPutS3TargetReadPreferenceSpectraS3Request(bucketId string, readPreference TargetReadPreferenceType, targetId string) *PutS3TargetReadPreferenceSpectraS3Request {
    return &PutS3TargetReadPreferenceSpectraS3Request{
        BucketId: bucketId,
        ReadPreference: readPreference,
        TargetId: targetId,
    }
}

type DeleteS3TargetBucketNameSpectraS3Request struct {
    S3TargetBucketName string
}

func NewDeleteS3TargetBucketNameSpectraS3Request(s3TargetBucketName string) *DeleteS3TargetBucketNameSpectraS3Request {
    return &DeleteS3TargetBucketNameSpectraS3Request{
        S3TargetBucketName: s3TargetBucketName,
    }
}

type DeleteS3TargetFailureSpectraS3Request struct {
    S3TargetFailure string
}

func NewDeleteS3TargetFailureSpectraS3Request(s3TargetFailure string) *DeleteS3TargetFailureSpectraS3Request {
    return &DeleteS3TargetFailureSpectraS3Request{
        S3TargetFailure: s3TargetFailure,
    }
}

type DeleteS3TargetReadPreferenceSpectraS3Request struct {
    S3TargetReadPreference string
}

func NewDeleteS3TargetReadPreferenceSpectraS3Request(s3TargetReadPreference string) *DeleteS3TargetReadPreferenceSpectraS3Request {
    return &DeleteS3TargetReadPreferenceSpectraS3Request{
        S3TargetReadPreference: s3TargetReadPreference,
    }
}

type DeleteS3TargetSpectraS3Request struct {
    S3Target string
}

func NewDeleteS3TargetSpectraS3Request(s3Target string) *DeleteS3TargetSpectraS3Request {
    return &DeleteS3TargetSpectraS3Request{
        S3Target: s3Target,
    }
}

type GetBlobsOnS3TargetSpectraS3Request struct {
    S3Target string
}

func NewGetBlobsOnS3TargetSpectraS3Request(s3Target string) *GetBlobsOnS3TargetSpectraS3Request {
    return &GetBlobsOnS3TargetSpectraS3Request{
        S3Target: s3Target,
    }
}

type GetS3TargetBucketNamesSpectraS3Request struct {
    BucketId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetId *string
}

func NewGetS3TargetBucketNamesSpectraS3Request() *GetS3TargetBucketNamesSpectraS3Request {
    return &GetS3TargetBucketNamesSpectraS3Request{
    }
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithBucketId(bucketId string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.BucketId = &bucketId
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithLastPage() *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.LastPage = true
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithName(name string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.Name = &name
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.PageLength = &pageLength
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.PageOffset = &pageOffset
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3TargetBucketNamesSpectraS3Request
}

func (getS3TargetBucketNamesSpectraS3Request *GetS3TargetBucketNamesSpectraS3Request) WithTargetId(targetId string) *GetS3TargetBucketNamesSpectraS3Request {
    getS3TargetBucketNamesSpectraS3Request.TargetId = &targetId
    return getS3TargetBucketNamesSpectraS3Request
}

type GetS3TargetFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetFailureType TargetFailureType
    TargetId *string
}

func NewGetS3TargetFailuresSpectraS3Request() *GetS3TargetFailuresSpectraS3Request {
    return &GetS3TargetFailuresSpectraS3Request{
    }
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithLastPage() *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.LastPage = true
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.PageLength = &pageLength
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.PageOffset = &pageOffset
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithTargetId(targetId string) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.TargetId = &targetId
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithTargetFailureType(targetFailureType TargetFailureType) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.TargetFailureType = targetFailureType
    return getS3TargetFailuresSpectraS3Request
}

type GetS3TargetReadPreferenceSpectraS3Request struct {
    S3TargetReadPreference string
}

func NewGetS3TargetReadPreferenceSpectraS3Request(s3TargetReadPreference string) *GetS3TargetReadPreferenceSpectraS3Request {
    return &GetS3TargetReadPreferenceSpectraS3Request{
        S3TargetReadPreference: s3TargetReadPreference,
    }
}

type GetS3TargetReadPreferencesSpectraS3Request struct {
    BucketId *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    ReadPreference TargetReadPreferenceType
    TargetId *string
}

func NewGetS3TargetReadPreferencesSpectraS3Request() *GetS3TargetReadPreferencesSpectraS3Request {
    return &GetS3TargetReadPreferencesSpectraS3Request{
    }
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithBucketId(bucketId string) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.BucketId = &bucketId
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithLastPage() *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.LastPage = true
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.PageLength = &pageLength
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.PageOffset = &pageOffset
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithReadPreference(readPreference TargetReadPreferenceType) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.ReadPreference = readPreference
    return getS3TargetReadPreferencesSpectraS3Request
}

func (getS3TargetReadPreferencesSpectraS3Request *GetS3TargetReadPreferencesSpectraS3Request) WithTargetId(targetId string) *GetS3TargetReadPreferencesSpectraS3Request {
    getS3TargetReadPreferencesSpectraS3Request.TargetId = &targetId
    return getS3TargetReadPreferencesSpectraS3Request
}

type GetS3TargetSpectraS3Request struct {
    S3Target string
}

func NewGetS3TargetSpectraS3Request(s3Target string) *GetS3TargetSpectraS3Request {
    return &GetS3TargetSpectraS3Request{
        S3Target: s3Target,
    }
}

type GetS3TargetsSpectraS3Request struct {
    AccessKey *string
    DataPathEndPoint *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    LastPage bool
    Name *string
    NamingMode CloudNamingMode
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    PermitGoingOutOfSync *bool
    Quiesced Quiesced
    Region S3Region
    State TargetState
}

func NewGetS3TargetsSpectraS3Request() *GetS3TargetsSpectraS3Request {
    return &GetS3TargetsSpectraS3Request{
    }
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithAccessKey(accessKey string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.AccessKey = &accessKey
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithHttps(https bool) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.Https = &https
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithLastPage() *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.LastPage = true
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithName(name string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.Name = &name
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithNamingMode(namingMode CloudNamingMode) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.NamingMode = namingMode
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.PageLength = &pageLength
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.PageOffset = &pageOffset
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithQuiesced(quiesced Quiesced) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.Quiesced = quiesced
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithRegion(region S3Region) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.Region = region
    return getS3TargetsSpectraS3Request
}

func (getS3TargetsSpectraS3Request *GetS3TargetsSpectraS3Request) WithState(state TargetState) *GetS3TargetsSpectraS3Request {
    getS3TargetsSpectraS3Request.State = state
    return getS3TargetsSpectraS3Request
}

type ImportS3TargetSpectraS3Request struct {
    CloudBucketName string
    DataPolicyId *string
    Priority Priority
    S3Target string
    UserId *string
}

func NewImportS3TargetSpectraS3Request(cloudBucketName string, s3Target string) *ImportS3TargetSpectraS3Request {
    return &ImportS3TargetSpectraS3Request{
        S3Target: s3Target,
        CloudBucketName: cloudBucketName,
    }
}

func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithDataPolicyId(dataPolicyId string) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.DataPolicyId = &dataPolicyId
    return importS3TargetSpectraS3Request
}

func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithPriority(priority Priority) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.Priority = priority
    return importS3TargetSpectraS3Request
}

func (importS3TargetSpectraS3Request *ImportS3TargetSpectraS3Request) WithUserId(userId string) *ImportS3TargetSpectraS3Request {
    importS3TargetSpectraS3Request.UserId = &userId
    return importS3TargetSpectraS3Request
}

type ModifyAllS3TargetsSpectraS3Request struct {
    Quiesced Quiesced
}

func NewModifyAllS3TargetsSpectraS3Request(quiesced Quiesced) *ModifyAllS3TargetsSpectraS3Request {
    return &ModifyAllS3TargetsSpectraS3Request{
        Quiesced: quiesced,
    }
}

type ModifyS3TargetSpectraS3Request struct {
    AccessKey *string
    AutoVerifyFrequencyInDays *int
    CloudBucketPrefix *string
    CloudBucketSuffix *string
    DataPathEndPoint *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    Name *string
    NamingMode CloudNamingMode
    OfflineDataStagingWindowInTb *int
    PermitGoingOutOfSync *bool
    ProxyDomain *string
    ProxyHost *string
    ProxyPassword *string
    ProxyPort *int
    ProxyUsername *string
    Quiesced Quiesced
    Region S3Region
    S3Target string
    SecretKey *string
    StagedDataExpirationInDays *int
}

func NewModifyS3TargetSpectraS3Request(s3Target string) *ModifyS3TargetSpectraS3Request {
    return &ModifyS3TargetSpectraS3Request{
        S3Target: s3Target,
    }
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithAccessKey(accessKey string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.AccessKey = &accessKey
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.AutoVerifyFrequencyInDays = &autoVerifyFrequencyInDays
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.CloudBucketPrefix = &cloudBucketPrefix
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.CloudBucketSuffix = &cloudBucketSuffix
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithHttps(https bool) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.Https = &https
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithName(name string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.Name = &name
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithNamingMode(namingMode CloudNamingMode) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.NamingMode = namingMode
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithOfflineDataStagingWindowInTb(offlineDataStagingWindowInTb int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.OfflineDataStagingWindowInTb = &offlineDataStagingWindowInTb
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyDomain(proxyDomain string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.ProxyDomain = &proxyDomain
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyHost(proxyHost string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.ProxyHost = &proxyHost
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyPassword(proxyPassword string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.ProxyPassword = &proxyPassword
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyPort(proxyPort int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.ProxyPort = &proxyPort
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithProxyUsername(proxyUsername string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.ProxyUsername = &proxyUsername
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithQuiesced(quiesced Quiesced) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.Quiesced = quiesced
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithRegion(region S3Region) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.Region = region
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithSecretKey(secretKey string) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.SecretKey = &secretKey
    return modifyS3TargetSpectraS3Request
}

func (modifyS3TargetSpectraS3Request *ModifyS3TargetSpectraS3Request) WithStagedDataExpirationInDays(stagedDataExpirationInDays int) *ModifyS3TargetSpectraS3Request {
    modifyS3TargetSpectraS3Request.StagedDataExpirationInDays = &stagedDataExpirationInDays
    return modifyS3TargetSpectraS3Request
}

type RegisterS3TargetSpectraS3Request struct {
    AccessKey string
    AutoVerifyFrequencyInDays *int
    CloudBucketPrefix *string
    CloudBucketSuffix *string
    DataPathEndPoint *string
    DefaultReadPreference TargetReadPreferenceType
    Https *bool
    Name string
    NamingMode CloudNamingMode
    OfflineDataStagingWindowInTb *int
    PermitGoingOutOfSync *bool
    ProxyDomain *string
    ProxyHost *string
    ProxyPassword *string
    ProxyPort *int
    ProxyUsername *string
    Region S3Region
    SecretKey string
    StagedDataExpirationInDays *int
}

func NewRegisterS3TargetSpectraS3Request(accessKey string, name string, secretKey string) *RegisterS3TargetSpectraS3Request {
    return &RegisterS3TargetSpectraS3Request{
        AccessKey: accessKey,
        Name: name,
        SecretKey: secretKey,
    }
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithAutoVerifyFrequencyInDays(autoVerifyFrequencyInDays int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.AutoVerifyFrequencyInDays = &autoVerifyFrequencyInDays
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithCloudBucketPrefix(cloudBucketPrefix string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.CloudBucketPrefix = &cloudBucketPrefix
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithCloudBucketSuffix(cloudBucketSuffix string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.CloudBucketSuffix = &cloudBucketSuffix
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithDataPathEndPoint(dataPathEndPoint string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.DataPathEndPoint = &dataPathEndPoint
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithDefaultReadPreference(defaultReadPreference TargetReadPreferenceType) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.DefaultReadPreference = defaultReadPreference
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithHttps(https bool) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.Https = &https
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithNamingMode(namingMode CloudNamingMode) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.NamingMode = namingMode
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithOfflineDataStagingWindowInTb(offlineDataStagingWindowInTb int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.OfflineDataStagingWindowInTb = &offlineDataStagingWindowInTb
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithPermitGoingOutOfSync(permitGoingOutOfSync bool) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.PermitGoingOutOfSync = &permitGoingOutOfSync
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyDomain(proxyDomain string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.ProxyDomain = &proxyDomain
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyHost(proxyHost string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.ProxyHost = &proxyHost
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyPassword(proxyPassword string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.ProxyPassword = &proxyPassword
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyPort(proxyPort int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.ProxyPort = &proxyPort
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithProxyUsername(proxyUsername string) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.ProxyUsername = &proxyUsername
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithRegion(region S3Region) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.Region = region
    return registerS3TargetSpectraS3Request
}

func (registerS3TargetSpectraS3Request *RegisterS3TargetSpectraS3Request) WithStagedDataExpirationInDays(stagedDataExpirationInDays int) *RegisterS3TargetSpectraS3Request {
    registerS3TargetSpectraS3Request.StagedDataExpirationInDays = &stagedDataExpirationInDays
    return registerS3TargetSpectraS3Request
}

type VerifyS3TargetSpectraS3Request struct {
    FullDetails bool
    S3Target string
}

func NewVerifyS3TargetSpectraS3Request(s3Target string) *VerifyS3TargetSpectraS3Request {
    return &VerifyS3TargetSpectraS3Request{
        S3Target: s3Target,
    }
}

func (verifyS3TargetSpectraS3Request *VerifyS3TargetSpectraS3Request) WithFullDetails() *VerifyS3TargetSpectraS3Request {
    verifyS3TargetSpectraS3Request.FullDetails = true
    return verifyS3TargetSpectraS3Request
}

type DelegateCreateUserSpectraS3Request struct {
    DefaultDataPolicyId *string
    Id *string
    MaxBuckets *int
    Name string
    SecretKey *string
}

func NewDelegateCreateUserSpectraS3Request(name string) *DelegateCreateUserSpectraS3Request {
    return &DelegateCreateUserSpectraS3Request{
        Name: name,
    }
}

func (delegateCreateUserSpectraS3Request *DelegateCreateUserSpectraS3Request) WithDefaultDataPolicyId(defaultDataPolicyId string) *DelegateCreateUserSpectraS3Request {
    delegateCreateUserSpectraS3Request.DefaultDataPolicyId = &defaultDataPolicyId
    return delegateCreateUserSpectraS3Request
}

func (delegateCreateUserSpectraS3Request *DelegateCreateUserSpectraS3Request) WithId(id string) *DelegateCreateUserSpectraS3Request {
    delegateCreateUserSpectraS3Request.Id = &id
    return delegateCreateUserSpectraS3Request
}

func (delegateCreateUserSpectraS3Request *DelegateCreateUserSpectraS3Request) WithMaxBuckets(maxBuckets int) *DelegateCreateUserSpectraS3Request {
    delegateCreateUserSpectraS3Request.MaxBuckets = &maxBuckets
    return delegateCreateUserSpectraS3Request
}

func (delegateCreateUserSpectraS3Request *DelegateCreateUserSpectraS3Request) WithSecretKey(secretKey string) *DelegateCreateUserSpectraS3Request {
    delegateCreateUserSpectraS3Request.SecretKey = &secretKey
    return delegateCreateUserSpectraS3Request
}

type DelegateDeleteUserSpectraS3Request struct {
    UserId string
}

func NewDelegateDeleteUserSpectraS3Request(userId string) *DelegateDeleteUserSpectraS3Request {
    return &DelegateDeleteUserSpectraS3Request{
        UserId: userId,
    }
}

type GetUserSpectraS3Request struct {
    UserId string
}

func NewGetUserSpectraS3Request(userId string) *GetUserSpectraS3Request {
    return &GetUserSpectraS3Request{
        UserId: userId,
    }
}

type GetUsersSpectraS3Request struct {
    AuthId *string
    DefaultDataPolicyId *string
    LastPage bool
    Name *string
    PageLength *int
    PageOffset *int
    PageStartMarker *string
}

func NewGetUsersSpectraS3Request() *GetUsersSpectraS3Request {
    return &GetUsersSpectraS3Request{
    }
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithAuthId(authId string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.AuthId = &authId
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithDefaultDataPolicyId(defaultDataPolicyId string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.DefaultDataPolicyId = &defaultDataPolicyId
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithLastPage() *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.LastPage = true
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithName(name string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.Name = &name
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithPageLength(pageLength int) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.PageLength = &pageLength
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithPageOffset(pageOffset int) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.PageOffset = &pageOffset
    return getUsersSpectraS3Request
}

func (getUsersSpectraS3Request *GetUsersSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetUsersSpectraS3Request {
    getUsersSpectraS3Request.PageStartMarker = &pageStartMarker
    return getUsersSpectraS3Request
}

type ModifyUserSpectraS3Request struct {
    DefaultDataPolicyId *string
    MaxBuckets *int
    Name *string
    SecretKey *string
    UserId string
}

func NewModifyUserSpectraS3Request(userId string) *ModifyUserSpectraS3Request {
    return &ModifyUserSpectraS3Request{
        UserId: userId,
    }
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) WithDefaultDataPolicyId(defaultDataPolicyId string) *ModifyUserSpectraS3Request {
    modifyUserSpectraS3Request.DefaultDataPolicyId = &defaultDataPolicyId
    return modifyUserSpectraS3Request
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) WithMaxBuckets(maxBuckets int) *ModifyUserSpectraS3Request {
    modifyUserSpectraS3Request.MaxBuckets = &maxBuckets
    return modifyUserSpectraS3Request
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) WithName(name string) *ModifyUserSpectraS3Request {
    modifyUserSpectraS3Request.Name = &name
    return modifyUserSpectraS3Request
}

func (modifyUserSpectraS3Request *ModifyUserSpectraS3Request) WithSecretKey(secretKey string) *ModifyUserSpectraS3Request {
    modifyUserSpectraS3Request.SecretKey = &secretKey
    return modifyUserSpectraS3Request
}

type RegenerateUserSecretKeySpectraS3Request struct {
    UserId string
}

func NewRegenerateUserSecretKeySpectraS3Request(userId string) *RegenerateUserSecretKeySpectraS3Request {
    return &RegenerateUserSecretKeySpectraS3Request{
        UserId: userId,
    }
}

