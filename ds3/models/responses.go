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
    "net/http"
    "io"
)

type AbortMultiPartUploadResponse struct {
    
    Headers *http.Header
}



func NewAbortMultiPartUploadResponse(webResponse WebResponse) (*AbortMultiPartUploadResponse, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &AbortMultiPartUploadResponse{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CompleteBlobjectResponse struct {
    
    Headers *http.Header
}



func NewCompleteBlobjectResponse(webResponse WebResponse) (*CompleteBlobjectResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &CompleteBlobjectResponse{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CompleteMultiPartUploadResponse struct {
    CompleteMultipartUploadResult CompleteMultipartUploadResult
    Headers *http.Header
}

func (completeMultiPartUploadResponse *CompleteMultiPartUploadResponse) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &completeMultiPartUploadResponse.CompleteMultipartUploadResult)
}

func NewCompleteMultiPartUploadResponse(webResponse WebResponse) (*CompleteMultiPartUploadResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CompleteMultiPartUploadResponse
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutBucketResponse struct {
    
    Headers *http.Header
}



func NewPutBucketResponse(webResponse WebResponse) (*PutBucketResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &PutBucketResponse{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutMultiPartUploadPartResponse struct {
    
    Headers *http.Header
}



func NewPutMultiPartUploadPartResponse(webResponse WebResponse) (*PutMultiPartUploadPartResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &PutMultiPartUploadPartResponse{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutObjectResponse struct {
    
    Headers *http.Header
}



func NewPutObjectResponse(webResponse WebResponse) (*PutObjectResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &PutObjectResponse{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteBucketResponse struct {
    
    Headers *http.Header
}



func NewDeleteBucketResponse(webResponse WebResponse) (*DeleteBucketResponse, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteBucketResponse{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteObjectResponse struct {
    
    Headers *http.Header
}



func NewDeleteObjectResponse(webResponse WebResponse) (*DeleteObjectResponse, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteObjectResponse{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteObjectsResponse struct {
    DeleteResult DeleteResult
    Headers *http.Header
}

func (deleteObjectsResponse *DeleteObjectsResponse) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &deleteObjectsResponse.DeleteResult)
}

func NewDeleteObjectsResponse(webResponse WebResponse) (*DeleteObjectsResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body DeleteObjectsResponse
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBucketResponse struct {
    ListBucketResult ListBucketResult
    Headers *http.Header
}

func (getBucketResponse *GetBucketResponse) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBucketResponse.ListBucketResult)
}

func NewGetBucketResponse(webResponse WebResponse) (*GetBucketResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBucketResponse
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetServiceResponse struct {
    ListAllMyBucketsResult ListAllMyBucketsResult
    Headers *http.Header
}

func (getServiceResponse *GetServiceResponse) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getServiceResponse.ListAllMyBucketsResult)
}

func NewGetServiceResponse(webResponse WebResponse) (*GetServiceResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetServiceResponse
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetObjectResponse struct {
    Content io.ReadCloser
    Headers *http.Header
}



func NewGetObjectResponse(webResponse WebResponse) (*GetObjectResponse, error) {
    expectedStatusCodes := []int { 200, 206 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &GetObjectResponse{ Content: webResponse.Body(), Headers: webResponse.Header() }, nil
    case 206:
        return &GetObjectResponse{ Content: webResponse.Body(), Headers: webResponse.Header() }, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type HeadBucketResponse struct {
    
    Headers *http.Header
}



func NewHeadBucketResponse(webResponse WebResponse) (*HeadBucketResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &HeadBucketResponse{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type HeadObjectResponse struct {
    BlobChecksumType ChecksumType
    BlobChecksums map[int64]string
    Headers *http.Header
}



func NewHeadObjectResponse(webResponse WebResponse) (*HeadObjectResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        checksumType, err := getBlobChecksumType(webResponse.Header())
        if err != nil {
            return nil, err
        }
        checksumMap, err := getBlobChecksumMap(webResponse.Header())
        if err != nil {
            return nil, err
        }
        return &HeadObjectResponse{BlobChecksumType: checksumType, BlobChecksums: checksumMap, Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type InitiateMultiPartUploadResponse struct {
    InitiateMultipartUploadResult InitiateMultipartUploadResult
    Headers *http.Header
}

func (initiateMultiPartUploadResponse *InitiateMultiPartUploadResponse) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &initiateMultiPartUploadResponse.InitiateMultipartUploadResult)
}

func NewInitiateMultiPartUploadResponse(webResponse WebResponse) (*InitiateMultiPartUploadResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body InitiateMultiPartUploadResponse
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ListMultiPartUploadPartsResponse struct {
    ListPartsResult ListPartsResult
    Headers *http.Header
}

func (listMultiPartUploadPartsResponse *ListMultiPartUploadPartsResponse) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &listMultiPartUploadPartsResponse.ListPartsResult)
}

func NewListMultiPartUploadPartsResponse(webResponse WebResponse) (*ListMultiPartUploadPartsResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ListMultiPartUploadPartsResponse
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ListMultiPartUploadsResponse struct {
    ListMultiPartUploadsResult ListMultiPartUploadsResult
    Headers *http.Header
}

func (listMultiPartUploadsResponse *ListMultiPartUploadsResponse) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &listMultiPartUploadsResponse.ListMultiPartUploadsResult)
}

func NewListMultiPartUploadsResponse(webResponse WebResponse) (*ListMultiPartUploadsResponse, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ListMultiPartUploadsResponse
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutBucketAclForGroupSpectraS3Response struct {
    BucketAcl BucketAcl
    Headers *http.Header
}

func (putBucketAclForGroupSpectraS3Response *PutBucketAclForGroupSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putBucketAclForGroupSpectraS3Response.BucketAcl)
}

func NewPutBucketAclForGroupSpectraS3Response(webResponse WebResponse) (*PutBucketAclForGroupSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutBucketAclForGroupSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutBucketAclForUserSpectraS3Response struct {
    BucketAcl BucketAcl
    Headers *http.Header
}

func (putBucketAclForUserSpectraS3Response *PutBucketAclForUserSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putBucketAclForUserSpectraS3Response.BucketAcl)
}

func NewPutBucketAclForUserSpectraS3Response(webResponse WebResponse) (*PutBucketAclForUserSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutBucketAclForUserSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutDataPolicyAclForGroupSpectraS3Response struct {
    DataPolicyAcl DataPolicyAcl
    Headers *http.Header
}

func (putDataPolicyAclForGroupSpectraS3Response *PutDataPolicyAclForGroupSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putDataPolicyAclForGroupSpectraS3Response.DataPolicyAcl)
}

func NewPutDataPolicyAclForGroupSpectraS3Response(webResponse WebResponse) (*PutDataPolicyAclForGroupSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutDataPolicyAclForGroupSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutDataPolicyAclForUserSpectraS3Response struct {
    DataPolicyAcl DataPolicyAcl
    Headers *http.Header
}

func (putDataPolicyAclForUserSpectraS3Response *PutDataPolicyAclForUserSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putDataPolicyAclForUserSpectraS3Response.DataPolicyAcl)
}

func NewPutDataPolicyAclForUserSpectraS3Response(webResponse WebResponse) (*PutDataPolicyAclForUserSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutDataPolicyAclForUserSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutGlobalBucketAclForGroupSpectraS3Response struct {
    BucketAcl BucketAcl
    Headers *http.Header
}

func (putGlobalBucketAclForGroupSpectraS3Response *PutGlobalBucketAclForGroupSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putGlobalBucketAclForGroupSpectraS3Response.BucketAcl)
}

func NewPutGlobalBucketAclForGroupSpectraS3Response(webResponse WebResponse) (*PutGlobalBucketAclForGroupSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutGlobalBucketAclForGroupSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutGlobalBucketAclForUserSpectraS3Response struct {
    BucketAcl BucketAcl
    Headers *http.Header
}

func (putGlobalBucketAclForUserSpectraS3Response *PutGlobalBucketAclForUserSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putGlobalBucketAclForUserSpectraS3Response.BucketAcl)
}

func NewPutGlobalBucketAclForUserSpectraS3Response(webResponse WebResponse) (*PutGlobalBucketAclForUserSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutGlobalBucketAclForUserSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutGlobalDataPolicyAclForGroupSpectraS3Response struct {
    DataPolicyAcl DataPolicyAcl
    Headers *http.Header
}

func (putGlobalDataPolicyAclForGroupSpectraS3Response *PutGlobalDataPolicyAclForGroupSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putGlobalDataPolicyAclForGroupSpectraS3Response.DataPolicyAcl)
}

func NewPutGlobalDataPolicyAclForGroupSpectraS3Response(webResponse WebResponse) (*PutGlobalDataPolicyAclForGroupSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutGlobalDataPolicyAclForGroupSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutGlobalDataPolicyAclForUserSpectraS3Response struct {
    DataPolicyAcl DataPolicyAcl
    Headers *http.Header
}

func (putGlobalDataPolicyAclForUserSpectraS3Response *PutGlobalDataPolicyAclForUserSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putGlobalDataPolicyAclForUserSpectraS3Response.DataPolicyAcl)
}

func NewPutGlobalDataPolicyAclForUserSpectraS3Response(webResponse WebResponse) (*PutGlobalDataPolicyAclForUserSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutGlobalDataPolicyAclForUserSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteBucketAclSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteBucketAclSpectraS3Response(webResponse WebResponse) (*DeleteBucketAclSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteBucketAclSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteDataPolicyAclSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteDataPolicyAclSpectraS3Response(webResponse WebResponse) (*DeleteDataPolicyAclSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteDataPolicyAclSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBucketAclSpectraS3Response struct {
    BucketAcl BucketAcl
    Headers *http.Header
}

func (getBucketAclSpectraS3Response *GetBucketAclSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBucketAclSpectraS3Response.BucketAcl)
}

func NewGetBucketAclSpectraS3Response(webResponse WebResponse) (*GetBucketAclSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBucketAclSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBucketAclsSpectraS3Response struct {
    BucketAclList BucketAclList
    Headers *http.Header
}

func (getBucketAclsSpectraS3Response *GetBucketAclsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBucketAclsSpectraS3Response.BucketAclList)
}

func NewGetBucketAclsSpectraS3Response(webResponse WebResponse) (*GetBucketAclsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBucketAclsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDataPolicyAclSpectraS3Response struct {
    DataPolicyAcl DataPolicyAcl
    Headers *http.Header
}

func (getDataPolicyAclSpectraS3Response *GetDataPolicyAclSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDataPolicyAclSpectraS3Response.DataPolicyAcl)
}

func NewGetDataPolicyAclSpectraS3Response(webResponse WebResponse) (*GetDataPolicyAclSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDataPolicyAclSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDataPolicyAclsSpectraS3Response struct {
    DataPolicyAclList DataPolicyAclList
    Headers *http.Header
}

func (getDataPolicyAclsSpectraS3Response *GetDataPolicyAclsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDataPolicyAclsSpectraS3Response.DataPolicyAclList)
}

func NewGetDataPolicyAclsSpectraS3Response(webResponse WebResponse) (*GetDataPolicyAclsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDataPolicyAclsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutBucketSpectraS3Response struct {
    Bucket Bucket
    Headers *http.Header
}

func (putBucketSpectraS3Response *PutBucketSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putBucketSpectraS3Response.Bucket)
}

func NewPutBucketSpectraS3Response(webResponse WebResponse) (*PutBucketSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutBucketSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteBucketSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteBucketSpectraS3Response(webResponse WebResponse) (*DeleteBucketSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteBucketSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBucketSpectraS3Response struct {
    Bucket Bucket
    Headers *http.Header
}

func (getBucketSpectraS3Response *GetBucketSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBucketSpectraS3Response.Bucket)
}

func NewGetBucketSpectraS3Response(webResponse WebResponse) (*GetBucketSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBucketSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBucketsSpectraS3Response struct {
    BucketList BucketList
    Headers *http.Header
}

func (getBucketsSpectraS3Response *GetBucketsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBucketsSpectraS3Response.BucketList)
}

func NewGetBucketsSpectraS3Response(webResponse WebResponse) (*GetBucketsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBucketsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyBucketSpectraS3Response struct {
    Bucket Bucket
    Headers *http.Header
}

func (modifyBucketSpectraS3Response *ModifyBucketSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyBucketSpectraS3Response.Bucket)
}

func NewModifyBucketSpectraS3Response(webResponse WebResponse) (*ModifyBucketSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyBucketSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ForceFullCacheReclaimSpectraS3Response struct {
    
    Headers *http.Header
}



func NewForceFullCacheReclaimSpectraS3Response(webResponse WebResponse) (*ForceFullCacheReclaimSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ForceFullCacheReclaimSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetCacheFilesystemSpectraS3Response struct {
    CacheFilesystem CacheFilesystem
    Headers *http.Header
}

func (getCacheFilesystemSpectraS3Response *GetCacheFilesystemSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getCacheFilesystemSpectraS3Response.CacheFilesystem)
}

func NewGetCacheFilesystemSpectraS3Response(webResponse WebResponse) (*GetCacheFilesystemSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetCacheFilesystemSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetCacheFilesystemsSpectraS3Response struct {
    CacheFilesystemList CacheFilesystemList
    Headers *http.Header
}

func (getCacheFilesystemsSpectraS3Response *GetCacheFilesystemsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getCacheFilesystemsSpectraS3Response.CacheFilesystemList)
}

func NewGetCacheFilesystemsSpectraS3Response(webResponse WebResponse) (*GetCacheFilesystemsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetCacheFilesystemsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetCacheStateSpectraS3Response struct {
    CacheInformation CacheInformation
    Headers *http.Header
}

func (getCacheStateSpectraS3Response *GetCacheStateSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getCacheStateSpectraS3Response.CacheInformation)
}

func NewGetCacheStateSpectraS3Response(webResponse WebResponse) (*GetCacheStateSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetCacheStateSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyCacheFilesystemSpectraS3Response struct {
    CacheFilesystem CacheFilesystem
    Headers *http.Header
}

func (modifyCacheFilesystemSpectraS3Response *ModifyCacheFilesystemSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyCacheFilesystemSpectraS3Response.CacheFilesystem)
}

func NewModifyCacheFilesystemSpectraS3Response(webResponse WebResponse) (*ModifyCacheFilesystemSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyCacheFilesystemSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBucketCapacitySummarySpectraS3Response struct {
    CapacitySummaryContainer CapacitySummaryContainer
    Headers *http.Header
}

func (getBucketCapacitySummarySpectraS3Response *GetBucketCapacitySummarySpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBucketCapacitySummarySpectraS3Response.CapacitySummaryContainer)
}

func NewGetBucketCapacitySummarySpectraS3Response(webResponse WebResponse) (*GetBucketCapacitySummarySpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBucketCapacitySummarySpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetStorageDomainCapacitySummarySpectraS3Response struct {
    CapacitySummaryContainer CapacitySummaryContainer
    Headers *http.Header
}

func (getStorageDomainCapacitySummarySpectraS3Response *GetStorageDomainCapacitySummarySpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getStorageDomainCapacitySummarySpectraS3Response.CapacitySummaryContainer)
}

func NewGetStorageDomainCapacitySummarySpectraS3Response(webResponse WebResponse) (*GetStorageDomainCapacitySummarySpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetStorageDomainCapacitySummarySpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSystemCapacitySummarySpectraS3Response struct {
    CapacitySummaryContainer CapacitySummaryContainer
    Headers *http.Header
}

func (getSystemCapacitySummarySpectraS3Response *GetSystemCapacitySummarySpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSystemCapacitySummarySpectraS3Response.CapacitySummaryContainer)
}

func NewGetSystemCapacitySummarySpectraS3Response(webResponse WebResponse) (*GetSystemCapacitySummarySpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSystemCapacitySummarySpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDataPathBackendSpectraS3Response struct {
    DataPathBackend DataPathBackend
    Headers *http.Header
}

func (getDataPathBackendSpectraS3Response *GetDataPathBackendSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDataPathBackendSpectraS3Response.DataPathBackend)
}

func NewGetDataPathBackendSpectraS3Response(webResponse WebResponse) (*GetDataPathBackendSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDataPathBackendSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDataPlannerBlobStoreTasksSpectraS3Response struct {
    BlobStoreTasksInformation BlobStoreTasksInformation
    Headers *http.Header
}

func (getDataPlannerBlobStoreTasksSpectraS3Response *GetDataPlannerBlobStoreTasksSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDataPlannerBlobStoreTasksSpectraS3Response.BlobStoreTasksInformation)
}

func NewGetDataPlannerBlobStoreTasksSpectraS3Response(webResponse WebResponse) (*GetDataPlannerBlobStoreTasksSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDataPlannerBlobStoreTasksSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyDataPathBackendSpectraS3Response struct {
    DataPathBackend DataPathBackend
    Headers *http.Header
}

func (modifyDataPathBackendSpectraS3Response *ModifyDataPathBackendSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyDataPathBackendSpectraS3Response.DataPathBackend)
}

func NewModifyDataPathBackendSpectraS3Response(webResponse WebResponse) (*ModifyDataPathBackendSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyDataPathBackendSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutAzureDataReplicationRuleSpectraS3Response struct {
    AzureDataReplicationRule AzureDataReplicationRule
    Headers *http.Header
}

func (putAzureDataReplicationRuleSpectraS3Response *PutAzureDataReplicationRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putAzureDataReplicationRuleSpectraS3Response.AzureDataReplicationRule)
}

func NewPutAzureDataReplicationRuleSpectraS3Response(webResponse WebResponse) (*PutAzureDataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutAzureDataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutDataPersistenceRuleSpectraS3Response struct {
    DataPersistenceRule DataPersistenceRule
    Headers *http.Header
}

func (putDataPersistenceRuleSpectraS3Response *PutDataPersistenceRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putDataPersistenceRuleSpectraS3Response.DataPersistenceRule)
}

func NewPutDataPersistenceRuleSpectraS3Response(webResponse WebResponse) (*PutDataPersistenceRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutDataPersistenceRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutDataPolicySpectraS3Response struct {
    DataPolicy DataPolicy
    Headers *http.Header
}

func (putDataPolicySpectraS3Response *PutDataPolicySpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putDataPolicySpectraS3Response.DataPolicy)
}

func NewPutDataPolicySpectraS3Response(webResponse WebResponse) (*PutDataPolicySpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutDataPolicySpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutDs3DataReplicationRuleSpectraS3Response struct {
    Ds3DataReplicationRule Ds3DataReplicationRule
    Headers *http.Header
}

func (putDs3DataReplicationRuleSpectraS3Response *PutDs3DataReplicationRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putDs3DataReplicationRuleSpectraS3Response.Ds3DataReplicationRule)
}

func NewPutDs3DataReplicationRuleSpectraS3Response(webResponse WebResponse) (*PutDs3DataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutDs3DataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutS3DataReplicationRuleSpectraS3Response struct {
    S3DataReplicationRule S3DataReplicationRule
    Headers *http.Header
}

func (putS3DataReplicationRuleSpectraS3Response *PutS3DataReplicationRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putS3DataReplicationRuleSpectraS3Response.S3DataReplicationRule)
}

func NewPutS3DataReplicationRuleSpectraS3Response(webResponse WebResponse) (*PutS3DataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutS3DataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteAzureDataReplicationRuleSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteAzureDataReplicationRuleSpectraS3Response(webResponse WebResponse) (*DeleteAzureDataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteAzureDataReplicationRuleSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteDataPersistenceRuleSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteDataPersistenceRuleSpectraS3Response(webResponse WebResponse) (*DeleteDataPersistenceRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteDataPersistenceRuleSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteDataPolicySpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteDataPolicySpectraS3Response(webResponse WebResponse) (*DeleteDataPolicySpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteDataPolicySpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteDs3DataReplicationRuleSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteDs3DataReplicationRuleSpectraS3Response(webResponse WebResponse) (*DeleteDs3DataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteDs3DataReplicationRuleSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteS3DataReplicationRuleSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteS3DataReplicationRuleSpectraS3Response(webResponse WebResponse) (*DeleteS3DataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteS3DataReplicationRuleSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetAzureDataReplicationRuleSpectraS3Response struct {
    AzureDataReplicationRule AzureDataReplicationRule
    Headers *http.Header
}

func (getAzureDataReplicationRuleSpectraS3Response *GetAzureDataReplicationRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getAzureDataReplicationRuleSpectraS3Response.AzureDataReplicationRule)
}

func NewGetAzureDataReplicationRuleSpectraS3Response(webResponse WebResponse) (*GetAzureDataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetAzureDataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetAzureDataReplicationRulesSpectraS3Response struct {
    AzureDataReplicationRuleList AzureDataReplicationRuleList
    Headers *http.Header
}

func (getAzureDataReplicationRulesSpectraS3Response *GetAzureDataReplicationRulesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getAzureDataReplicationRulesSpectraS3Response.AzureDataReplicationRuleList)
}

func NewGetAzureDataReplicationRulesSpectraS3Response(webResponse WebResponse) (*GetAzureDataReplicationRulesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetAzureDataReplicationRulesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDataPersistenceRuleSpectraS3Response struct {
    DataPersistenceRule DataPersistenceRule
    Headers *http.Header
}

func (getDataPersistenceRuleSpectraS3Response *GetDataPersistenceRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDataPersistenceRuleSpectraS3Response.DataPersistenceRule)
}

func NewGetDataPersistenceRuleSpectraS3Response(webResponse WebResponse) (*GetDataPersistenceRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDataPersistenceRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDataPersistenceRulesSpectraS3Response struct {
    DataPersistenceRuleList DataPersistenceRuleList
    Headers *http.Header
}

func (getDataPersistenceRulesSpectraS3Response *GetDataPersistenceRulesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDataPersistenceRulesSpectraS3Response.DataPersistenceRuleList)
}

func NewGetDataPersistenceRulesSpectraS3Response(webResponse WebResponse) (*GetDataPersistenceRulesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDataPersistenceRulesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDataPoliciesSpectraS3Response struct {
    DataPolicyList DataPolicyList
    Headers *http.Header
}

func (getDataPoliciesSpectraS3Response *GetDataPoliciesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDataPoliciesSpectraS3Response.DataPolicyList)
}

func NewGetDataPoliciesSpectraS3Response(webResponse WebResponse) (*GetDataPoliciesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDataPoliciesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDataPolicySpectraS3Response struct {
    DataPolicy DataPolicy
    Headers *http.Header
}

func (getDataPolicySpectraS3Response *GetDataPolicySpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDataPolicySpectraS3Response.DataPolicy)
}

func NewGetDataPolicySpectraS3Response(webResponse WebResponse) (*GetDataPolicySpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDataPolicySpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDs3DataReplicationRuleSpectraS3Response struct {
    Ds3DataReplicationRule Ds3DataReplicationRule
    Headers *http.Header
}

func (getDs3DataReplicationRuleSpectraS3Response *GetDs3DataReplicationRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDs3DataReplicationRuleSpectraS3Response.Ds3DataReplicationRule)
}

func NewGetDs3DataReplicationRuleSpectraS3Response(webResponse WebResponse) (*GetDs3DataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDs3DataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDs3DataReplicationRulesSpectraS3Response struct {
    Ds3DataReplicationRuleList Ds3DataReplicationRuleList
    Headers *http.Header
}

func (getDs3DataReplicationRulesSpectraS3Response *GetDs3DataReplicationRulesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDs3DataReplicationRulesSpectraS3Response.Ds3DataReplicationRuleList)
}

func NewGetDs3DataReplicationRulesSpectraS3Response(webResponse WebResponse) (*GetDs3DataReplicationRulesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDs3DataReplicationRulesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetS3DataReplicationRuleSpectraS3Response struct {
    S3DataReplicationRule S3DataReplicationRule
    Headers *http.Header
}

func (getS3DataReplicationRuleSpectraS3Response *GetS3DataReplicationRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getS3DataReplicationRuleSpectraS3Response.S3DataReplicationRule)
}

func NewGetS3DataReplicationRuleSpectraS3Response(webResponse WebResponse) (*GetS3DataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetS3DataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetS3DataReplicationRulesSpectraS3Response struct {
    S3DataReplicationRuleList S3DataReplicationRuleList
    Headers *http.Header
}

func (getS3DataReplicationRulesSpectraS3Response *GetS3DataReplicationRulesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getS3DataReplicationRulesSpectraS3Response.S3DataReplicationRuleList)
}

func NewGetS3DataReplicationRulesSpectraS3Response(webResponse WebResponse) (*GetS3DataReplicationRulesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetS3DataReplicationRulesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyAzureDataReplicationRuleSpectraS3Response struct {
    AzureDataReplicationRule AzureDataReplicationRule
    Headers *http.Header
}

func (modifyAzureDataReplicationRuleSpectraS3Response *ModifyAzureDataReplicationRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyAzureDataReplicationRuleSpectraS3Response.AzureDataReplicationRule)
}

func NewModifyAzureDataReplicationRuleSpectraS3Response(webResponse WebResponse) (*ModifyAzureDataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyAzureDataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyDataPersistenceRuleSpectraS3Response struct {
    DataPersistenceRule DataPersistenceRule
    Headers *http.Header
}

func (modifyDataPersistenceRuleSpectraS3Response *ModifyDataPersistenceRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyDataPersistenceRuleSpectraS3Response.DataPersistenceRule)
}

func NewModifyDataPersistenceRuleSpectraS3Response(webResponse WebResponse) (*ModifyDataPersistenceRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyDataPersistenceRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyDataPolicySpectraS3Response struct {
    DataPolicy DataPolicy
    Headers *http.Header
}

func (modifyDataPolicySpectraS3Response *ModifyDataPolicySpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyDataPolicySpectraS3Response.DataPolicy)
}

func NewModifyDataPolicySpectraS3Response(webResponse WebResponse) (*ModifyDataPolicySpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyDataPolicySpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyDs3DataReplicationRuleSpectraS3Response struct {
    Ds3DataReplicationRule Ds3DataReplicationRule
    Headers *http.Header
}

func (modifyDs3DataReplicationRuleSpectraS3Response *ModifyDs3DataReplicationRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyDs3DataReplicationRuleSpectraS3Response.Ds3DataReplicationRule)
}

func NewModifyDs3DataReplicationRuleSpectraS3Response(webResponse WebResponse) (*ModifyDs3DataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyDs3DataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyS3DataReplicationRuleSpectraS3Response struct {
    S3DataReplicationRule S3DataReplicationRule
    Headers *http.Header
}

func (modifyS3DataReplicationRuleSpectraS3Response *ModifyS3DataReplicationRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyS3DataReplicationRuleSpectraS3Response.S3DataReplicationRule)
}

func NewModifyS3DataReplicationRuleSpectraS3Response(webResponse WebResponse) (*ModifyS3DataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyS3DataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ClearSuspectBlobAzureTargetsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewClearSuspectBlobAzureTargetsSpectraS3Response(webResponse WebResponse) (*ClearSuspectBlobAzureTargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ClearSuspectBlobAzureTargetsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ClearSuspectBlobDs3TargetsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewClearSuspectBlobDs3TargetsSpectraS3Response(webResponse WebResponse) (*ClearSuspectBlobDs3TargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ClearSuspectBlobDs3TargetsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ClearSuspectBlobPoolsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewClearSuspectBlobPoolsSpectraS3Response(webResponse WebResponse) (*ClearSuspectBlobPoolsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ClearSuspectBlobPoolsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ClearSuspectBlobS3TargetsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewClearSuspectBlobS3TargetsSpectraS3Response(webResponse WebResponse) (*ClearSuspectBlobS3TargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ClearSuspectBlobS3TargetsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ClearSuspectBlobTapesSpectraS3Response struct {
    
    Headers *http.Header
}



func NewClearSuspectBlobTapesSpectraS3Response(webResponse WebResponse) (*ClearSuspectBlobTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ClearSuspectBlobTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDegradedAzureDataReplicationRulesSpectraS3Response struct {
    AzureDataReplicationRuleList AzureDataReplicationRuleList
    Headers *http.Header
}

func (getDegradedAzureDataReplicationRulesSpectraS3Response *GetDegradedAzureDataReplicationRulesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDegradedAzureDataReplicationRulesSpectraS3Response.AzureDataReplicationRuleList)
}

func NewGetDegradedAzureDataReplicationRulesSpectraS3Response(webResponse WebResponse) (*GetDegradedAzureDataReplicationRulesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDegradedAzureDataReplicationRulesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDegradedBlobsSpectraS3Response struct {
    DegradedBlobList DegradedBlobList
    Headers *http.Header
}

func (getDegradedBlobsSpectraS3Response *GetDegradedBlobsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDegradedBlobsSpectraS3Response.DegradedBlobList)
}

func NewGetDegradedBlobsSpectraS3Response(webResponse WebResponse) (*GetDegradedBlobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDegradedBlobsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDegradedBucketsSpectraS3Response struct {
    BucketList BucketList
    Headers *http.Header
}

func (getDegradedBucketsSpectraS3Response *GetDegradedBucketsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDegradedBucketsSpectraS3Response.BucketList)
}

func NewGetDegradedBucketsSpectraS3Response(webResponse WebResponse) (*GetDegradedBucketsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDegradedBucketsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDegradedDataPersistenceRulesSpectraS3Response struct {
    DataPersistenceRuleList DataPersistenceRuleList
    Headers *http.Header
}

func (getDegradedDataPersistenceRulesSpectraS3Response *GetDegradedDataPersistenceRulesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDegradedDataPersistenceRulesSpectraS3Response.DataPersistenceRuleList)
}

func NewGetDegradedDataPersistenceRulesSpectraS3Response(webResponse WebResponse) (*GetDegradedDataPersistenceRulesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDegradedDataPersistenceRulesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDegradedDs3DataReplicationRulesSpectraS3Response struct {
    Ds3DataReplicationRuleList Ds3DataReplicationRuleList
    Headers *http.Header
}

func (getDegradedDs3DataReplicationRulesSpectraS3Response *GetDegradedDs3DataReplicationRulesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDegradedDs3DataReplicationRulesSpectraS3Response.Ds3DataReplicationRuleList)
}

func NewGetDegradedDs3DataReplicationRulesSpectraS3Response(webResponse WebResponse) (*GetDegradedDs3DataReplicationRulesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDegradedDs3DataReplicationRulesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDegradedS3DataReplicationRulesSpectraS3Response struct {
    S3DataReplicationRuleList S3DataReplicationRuleList
    Headers *http.Header
}

func (getDegradedS3DataReplicationRulesSpectraS3Response *GetDegradedS3DataReplicationRulesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDegradedS3DataReplicationRulesSpectraS3Response.S3DataReplicationRuleList)
}

func NewGetDegradedS3DataReplicationRulesSpectraS3Response(webResponse WebResponse) (*GetDegradedS3DataReplicationRulesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDegradedS3DataReplicationRulesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSuspectBlobAzureTargetsSpectraS3Response struct {
    SuspectBlobAzureTargetList SuspectBlobAzureTargetList
    Headers *http.Header
}

func (getSuspectBlobAzureTargetsSpectraS3Response *GetSuspectBlobAzureTargetsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSuspectBlobAzureTargetsSpectraS3Response.SuspectBlobAzureTargetList)
}

func NewGetSuspectBlobAzureTargetsSpectraS3Response(webResponse WebResponse) (*GetSuspectBlobAzureTargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSuspectBlobAzureTargetsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSuspectBlobDs3TargetsSpectraS3Response struct {
    SuspectBlobDs3TargetList SuspectBlobDs3TargetList
    Headers *http.Header
}

func (getSuspectBlobDs3TargetsSpectraS3Response *GetSuspectBlobDs3TargetsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSuspectBlobDs3TargetsSpectraS3Response.SuspectBlobDs3TargetList)
}

func NewGetSuspectBlobDs3TargetsSpectraS3Response(webResponse WebResponse) (*GetSuspectBlobDs3TargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSuspectBlobDs3TargetsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSuspectBlobPoolsSpectraS3Response struct {
    SuspectBlobPoolList SuspectBlobPoolList
    Headers *http.Header
}

func (getSuspectBlobPoolsSpectraS3Response *GetSuspectBlobPoolsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSuspectBlobPoolsSpectraS3Response.SuspectBlobPoolList)
}

func NewGetSuspectBlobPoolsSpectraS3Response(webResponse WebResponse) (*GetSuspectBlobPoolsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSuspectBlobPoolsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSuspectBlobS3TargetsSpectraS3Response struct {
    SuspectBlobS3TargetList SuspectBlobS3TargetList
    Headers *http.Header
}

func (getSuspectBlobS3TargetsSpectraS3Response *GetSuspectBlobS3TargetsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSuspectBlobS3TargetsSpectraS3Response.SuspectBlobS3TargetList)
}

func NewGetSuspectBlobS3TargetsSpectraS3Response(webResponse WebResponse) (*GetSuspectBlobS3TargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSuspectBlobS3TargetsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSuspectBlobTapesSpectraS3Response struct {
    SuspectBlobTapeList SuspectBlobTapeList
    Headers *http.Header
}

func (getSuspectBlobTapesSpectraS3Response *GetSuspectBlobTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSuspectBlobTapesSpectraS3Response.SuspectBlobTapeList)
}

func NewGetSuspectBlobTapesSpectraS3Response(webResponse WebResponse) (*GetSuspectBlobTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSuspectBlobTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSuspectBucketsSpectraS3Response struct {
    BucketList BucketList
    Headers *http.Header
}

func (getSuspectBucketsSpectraS3Response *GetSuspectBucketsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSuspectBucketsSpectraS3Response.BucketList)
}

func NewGetSuspectBucketsSpectraS3Response(webResponse WebResponse) (*GetSuspectBucketsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSuspectBucketsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSuspectObjectsSpectraS3Response struct {
    S3ObjectList S3ObjectList
    Headers *http.Header
}

func (getSuspectObjectsSpectraS3Response *GetSuspectObjectsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSuspectObjectsSpectraS3Response.S3ObjectList)
}

func NewGetSuspectObjectsSpectraS3Response(webResponse WebResponse) (*GetSuspectObjectsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSuspectObjectsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSuspectObjectsWithFullDetailsSpectraS3Response struct {
    BulkObjectList BulkObjectList
    Headers *http.Header
}

func (getSuspectObjectsWithFullDetailsSpectraS3Response *GetSuspectObjectsWithFullDetailsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSuspectObjectsWithFullDetailsSpectraS3Response.BulkObjectList)
}

func NewGetSuspectObjectsWithFullDetailsSpectraS3Response(webResponse WebResponse) (*GetSuspectObjectsWithFullDetailsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSuspectObjectsWithFullDetailsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Response struct {
    
    Headers *http.Header
}



func NewMarkSuspectBlobAzureTargetsAsDegradedSpectraS3Response(webResponse WebResponse) (*MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &MarkSuspectBlobAzureTargetsAsDegradedSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Response struct {
    
    Headers *http.Header
}



func NewMarkSuspectBlobDs3TargetsAsDegradedSpectraS3Response(webResponse WebResponse) (*MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &MarkSuspectBlobDs3TargetsAsDegradedSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type MarkSuspectBlobPoolsAsDegradedSpectraS3Response struct {
    
    Headers *http.Header
}



func NewMarkSuspectBlobPoolsAsDegradedSpectraS3Response(webResponse WebResponse) (*MarkSuspectBlobPoolsAsDegradedSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &MarkSuspectBlobPoolsAsDegradedSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type MarkSuspectBlobS3TargetsAsDegradedSpectraS3Response struct {
    
    Headers *http.Header
}



func NewMarkSuspectBlobS3TargetsAsDegradedSpectraS3Response(webResponse WebResponse) (*MarkSuspectBlobS3TargetsAsDegradedSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &MarkSuspectBlobS3TargetsAsDegradedSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type MarkSuspectBlobTapesAsDegradedSpectraS3Response struct {
    
    Headers *http.Header
}



func NewMarkSuspectBlobTapesAsDegradedSpectraS3Response(webResponse WebResponse) (*MarkSuspectBlobTapesAsDegradedSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &MarkSuspectBlobTapesAsDegradedSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutGroupGroupMemberSpectraS3Response struct {
    GroupMember GroupMember
    Headers *http.Header
}

func (putGroupGroupMemberSpectraS3Response *PutGroupGroupMemberSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putGroupGroupMemberSpectraS3Response.GroupMember)
}

func NewPutGroupGroupMemberSpectraS3Response(webResponse WebResponse) (*PutGroupGroupMemberSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutGroupGroupMemberSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutGroupSpectraS3Response struct {
    Group Group
    Headers *http.Header
}

func (putGroupSpectraS3Response *PutGroupSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putGroupSpectraS3Response.Group)
}

func NewPutGroupSpectraS3Response(webResponse WebResponse) (*PutGroupSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutGroupSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutUserGroupMemberSpectraS3Response struct {
    GroupMember GroupMember
    Headers *http.Header
}

func (putUserGroupMemberSpectraS3Response *PutUserGroupMemberSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putUserGroupMemberSpectraS3Response.GroupMember)
}

func NewPutUserGroupMemberSpectraS3Response(webResponse WebResponse) (*PutUserGroupMemberSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutUserGroupMemberSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteGroupMemberSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteGroupMemberSpectraS3Response(webResponse WebResponse) (*DeleteGroupMemberSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteGroupMemberSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteGroupSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteGroupSpectraS3Response(webResponse WebResponse) (*DeleteGroupSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteGroupSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetGroupMemberSpectraS3Response struct {
    GroupMember GroupMember
    Headers *http.Header
}

func (getGroupMemberSpectraS3Response *GetGroupMemberSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getGroupMemberSpectraS3Response.GroupMember)
}

func NewGetGroupMemberSpectraS3Response(webResponse WebResponse) (*GetGroupMemberSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetGroupMemberSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetGroupMembersSpectraS3Response struct {
    GroupMemberList GroupMemberList
    Headers *http.Header
}

func (getGroupMembersSpectraS3Response *GetGroupMembersSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getGroupMembersSpectraS3Response.GroupMemberList)
}

func NewGetGroupMembersSpectraS3Response(webResponse WebResponse) (*GetGroupMembersSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetGroupMembersSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetGroupSpectraS3Response struct {
    Group Group
    Headers *http.Header
}

func (getGroupSpectraS3Response *GetGroupSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getGroupSpectraS3Response.Group)
}

func NewGetGroupSpectraS3Response(webResponse WebResponse) (*GetGroupSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetGroupSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetGroupsSpectraS3Response struct {
    GroupList GroupList
    Headers *http.Header
}

func (getGroupsSpectraS3Response *GetGroupsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getGroupsSpectraS3Response.GroupList)
}

func NewGetGroupsSpectraS3Response(webResponse WebResponse) (*GetGroupsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetGroupsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyGroupSpectraS3Response struct {
    Group Group
    Headers *http.Header
}

func (modifyGroupSpectraS3Response *ModifyGroupSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyGroupSpectraS3Response.Group)
}

func NewModifyGroupSpectraS3Response(webResponse WebResponse) (*ModifyGroupSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyGroupSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyUserIsMemberOfGroupSpectraS3Response struct {
    Group *Group
    Headers *http.Header
}

func (verifyUserIsMemberOfGroupSpectraS3Response *VerifyUserIsMemberOfGroupSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, verifyUserIsMemberOfGroupSpectraS3Response.Group)
}

func NewVerifyUserIsMemberOfGroupSpectraS3Response(webResponse WebResponse) (*VerifyUserIsMemberOfGroupSpectraS3Response, error) {
    expectedStatusCodes := []int { 200, 204 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body VerifyUserIsMemberOfGroupSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    case 204:
        return &VerifyUserIsMemberOfGroupSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type AllocateJobChunkSpectraS3Response struct {
    Objects Objects
    Headers *http.Header
}

func (allocateJobChunkSpectraS3Response *AllocateJobChunkSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &allocateJobChunkSpectraS3Response.Objects)
}

func NewAllocateJobChunkSpectraS3Response(webResponse WebResponse) (*AllocateJobChunkSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body AllocateJobChunkSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelActiveJobSpectraS3Response struct {
    
    Headers *http.Header
}



func NewCancelActiveJobSpectraS3Response(webResponse WebResponse) (*CancelActiveJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelActiveJobSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelAllActiveJobsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewCancelAllActiveJobsSpectraS3Response(webResponse WebResponse) (*CancelAllActiveJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelAllActiveJobsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelAllJobsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewCancelAllJobsSpectraS3Response(webResponse WebResponse) (*CancelAllJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelAllJobsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelJobSpectraS3Response struct {
    
    Headers *http.Header
}



func NewCancelJobSpectraS3Response(webResponse WebResponse) (*CancelJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelJobSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ClearAllCanceledJobsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewClearAllCanceledJobsSpectraS3Response(webResponse WebResponse) (*ClearAllCanceledJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ClearAllCanceledJobsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ClearAllCompletedJobsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewClearAllCompletedJobsSpectraS3Response(webResponse WebResponse) (*ClearAllCompletedJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ClearAllCompletedJobsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CloseAggregatingJobSpectraS3Response struct {
    MasterObjectList MasterObjectList
    Headers *http.Header
}

func (closeAggregatingJobSpectraS3Response *CloseAggregatingJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &closeAggregatingJobSpectraS3Response.MasterObjectList)
}

func NewCloseAggregatingJobSpectraS3Response(webResponse WebResponse) (*CloseAggregatingJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CloseAggregatingJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBulkJobSpectraS3Response struct {
    MasterObjectList MasterObjectList
    Headers *http.Header
}

func (getBulkJobSpectraS3Response *GetBulkJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBulkJobSpectraS3Response.MasterObjectList)
}

func NewGetBulkJobSpectraS3Response(webResponse WebResponse) (*GetBulkJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBulkJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutBulkJobSpectraS3Response struct {
    MasterObjectList MasterObjectList
    Headers *http.Header
}

func (putBulkJobSpectraS3Response *PutBulkJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putBulkJobSpectraS3Response.MasterObjectList)
}

func NewPutBulkJobSpectraS3Response(webResponse WebResponse) (*PutBulkJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body PutBulkJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyBulkJobSpectraS3Response struct {
    MasterObjectList MasterObjectList
    Headers *http.Header
}

func (verifyBulkJobSpectraS3Response *VerifyBulkJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &verifyBulkJobSpectraS3Response.MasterObjectList)
}

func NewVerifyBulkJobSpectraS3Response(webResponse WebResponse) (*VerifyBulkJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body VerifyBulkJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetActiveJobSpectraS3Response struct {
    ActiveJob ActiveJob
    Headers *http.Header
}

func (getActiveJobSpectraS3Response *GetActiveJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getActiveJobSpectraS3Response.ActiveJob)
}

func NewGetActiveJobSpectraS3Response(webResponse WebResponse) (*GetActiveJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetActiveJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetActiveJobsSpectraS3Response struct {
    ActiveJobList ActiveJobList
    Headers *http.Header
}

func (getActiveJobsSpectraS3Response *GetActiveJobsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getActiveJobsSpectraS3Response.ActiveJobList)
}

func NewGetActiveJobsSpectraS3Response(webResponse WebResponse) (*GetActiveJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetActiveJobsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetCanceledJobSpectraS3Response struct {
    CanceledJob CanceledJob
    Headers *http.Header
}

func (getCanceledJobSpectraS3Response *GetCanceledJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getCanceledJobSpectraS3Response.CanceledJob)
}

func NewGetCanceledJobSpectraS3Response(webResponse WebResponse) (*GetCanceledJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetCanceledJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetCanceledJobsSpectraS3Response struct {
    CanceledJobList CanceledJobList
    Headers *http.Header
}

func (getCanceledJobsSpectraS3Response *GetCanceledJobsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getCanceledJobsSpectraS3Response.CanceledJobList)
}

func NewGetCanceledJobsSpectraS3Response(webResponse WebResponse) (*GetCanceledJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetCanceledJobsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetCompletedJobSpectraS3Response struct {
    CompletedJob CompletedJob
    Headers *http.Header
}

func (getCompletedJobSpectraS3Response *GetCompletedJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getCompletedJobSpectraS3Response.CompletedJob)
}

func NewGetCompletedJobSpectraS3Response(webResponse WebResponse) (*GetCompletedJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetCompletedJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetCompletedJobsSpectraS3Response struct {
    CompletedJobList CompletedJobList
    Headers *http.Header
}

func (getCompletedJobsSpectraS3Response *GetCompletedJobsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getCompletedJobsSpectraS3Response.CompletedJobList)
}

func NewGetCompletedJobsSpectraS3Response(webResponse WebResponse) (*GetCompletedJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetCompletedJobsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobChunkDaoSpectraS3Response struct {
    JobChunk JobChunk
    Headers *http.Header
}

func (getJobChunkDaoSpectraS3Response *GetJobChunkDaoSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobChunkDaoSpectraS3Response.JobChunk)
}

func NewGetJobChunkDaoSpectraS3Response(webResponse WebResponse) (*GetJobChunkDaoSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobChunkDaoSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobChunkSpectraS3Response struct {
    Objects Objects
    Headers *http.Header
}

func (getJobChunkSpectraS3Response *GetJobChunkSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobChunkSpectraS3Response.Objects)
}

func NewGetJobChunkSpectraS3Response(webResponse WebResponse) (*GetJobChunkSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobChunkSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobChunksReadyForClientProcessingSpectraS3Response struct {
    MasterObjectList MasterObjectList
    Headers *http.Header
}

func (getJobChunksReadyForClientProcessingSpectraS3Response *GetJobChunksReadyForClientProcessingSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobChunksReadyForClientProcessingSpectraS3Response.MasterObjectList)
}

func NewGetJobChunksReadyForClientProcessingSpectraS3Response(webResponse WebResponse) (*GetJobChunksReadyForClientProcessingSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobChunksReadyForClientProcessingSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobSpectraS3Response struct {
    MasterObjectList MasterObjectList
    Headers *http.Header
}

func (getJobSpectraS3Response *GetJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobSpectraS3Response.MasterObjectList)
}

func NewGetJobSpectraS3Response(webResponse WebResponse) (*GetJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobToReplicateSpectraS3Response struct {
    Content string
    Headers *http.Header
}



func NewGetJobToReplicateSpectraS3Response(webResponse WebResponse) (*GetJobToReplicateSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        content, err := getResponseBodyAsString(webResponse)
        if err != nil {
            return nil, err
        }
        return &GetJobToReplicateSpectraS3Response{Content: content, Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobsSpectraS3Response struct {
    JobList JobList
    Headers *http.Header
}

func (getJobsSpectraS3Response *GetJobsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobsSpectraS3Response.JobList)
}

func NewGetJobsSpectraS3Response(webResponse WebResponse) (*GetJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyActiveJobSpectraS3Response struct {
    MasterObjectList MasterObjectList
    Headers *http.Header
}

func (modifyActiveJobSpectraS3Response *ModifyActiveJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyActiveJobSpectraS3Response.MasterObjectList)
}

func NewModifyActiveJobSpectraS3Response(webResponse WebResponse) (*ModifyActiveJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyActiveJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyJobSpectraS3Response struct {
    MasterObjectList MasterObjectList
    Headers *http.Header
}

func (modifyJobSpectraS3Response *ModifyJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyJobSpectraS3Response.MasterObjectList)
}

func NewModifyJobSpectraS3Response(webResponse WebResponse) (*ModifyJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ReplicatePutJobSpectraS3Response struct {
    MasterObjectList *MasterObjectList
    Headers *http.Header
}

func (replicatePutJobSpectraS3Response *ReplicatePutJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, replicatePutJobSpectraS3Response.MasterObjectList)
}

func NewReplicatePutJobSpectraS3Response(webResponse WebResponse) (*ReplicatePutJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200, 204 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ReplicatePutJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    case 204:
        return &ReplicatePutJobSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type StageObjectsJobSpectraS3Response struct {
    MasterObjectList MasterObjectList
    Headers *http.Header
}

func (stageObjectsJobSpectraS3Response *StageObjectsJobSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &stageObjectsJobSpectraS3Response.MasterObjectList)
}

func NewStageObjectsJobSpectraS3Response(webResponse WebResponse) (*StageObjectsJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body StageObjectsJobSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type TruncateActiveJobSpectraS3Response struct {
    
    Headers *http.Header
}



func NewTruncateActiveJobSpectraS3Response(webResponse WebResponse) (*TruncateActiveJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &TruncateActiveJobSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type TruncateAllActiveJobsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewTruncateAllActiveJobsSpectraS3Response(webResponse WebResponse) (*TruncateAllActiveJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &TruncateAllActiveJobsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type TruncateAllJobsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewTruncateAllJobsSpectraS3Response(webResponse WebResponse) (*TruncateAllJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &TruncateAllJobsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type TruncateJobSpectraS3Response struct {
    
    Headers *http.Header
}



func NewTruncateJobSpectraS3Response(webResponse WebResponse) (*TruncateJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &TruncateJobSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifySafeToCreatePutJobSpectraS3Response struct {
    
    Headers *http.Header
}



func NewVerifySafeToCreatePutJobSpectraS3Response(webResponse WebResponse) (*VerifySafeToCreatePutJobSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        return &VerifySafeToCreatePutJobSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetNodeSpectraS3Response struct {
    Node Node
    Headers *http.Header
}

func (getNodeSpectraS3Response *GetNodeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getNodeSpectraS3Response.Node)
}

func NewGetNodeSpectraS3Response(webResponse WebResponse) (*GetNodeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetNodeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetNodesSpectraS3Response struct {
    NodeList NodeList
    Headers *http.Header
}

func (getNodesSpectraS3Response *GetNodesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getNodesSpectraS3Response.NodeList)
}

func NewGetNodesSpectraS3Response(webResponse WebResponse) (*GetNodesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetNodesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyNodeSpectraS3Response struct {
    Node Node
    Headers *http.Header
}

func (modifyNodeSpectraS3Response *ModifyNodeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyNodeSpectraS3Response.Node)
}

func NewModifyNodeSpectraS3Response(webResponse WebResponse) (*ModifyNodeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyNodeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutAzureTargetFailureNotificationRegistrationSpectraS3Response struct {
    AzureTargetFailureNotificationRegistration AzureTargetFailureNotificationRegistration
    Headers *http.Header
}

func (putAzureTargetFailureNotificationRegistrationSpectraS3Response *PutAzureTargetFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putAzureTargetFailureNotificationRegistrationSpectraS3Response.AzureTargetFailureNotificationRegistration)
}

func NewPutAzureTargetFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutAzureTargetFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutAzureTargetFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutDs3TargetFailureNotificationRegistrationSpectraS3Response struct {
    Ds3TargetFailureNotificationRegistration Ds3TargetFailureNotificationRegistration
    Headers *http.Header
}

func (putDs3TargetFailureNotificationRegistrationSpectraS3Response *PutDs3TargetFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putDs3TargetFailureNotificationRegistrationSpectraS3Response.Ds3TargetFailureNotificationRegistration)
}

func NewPutDs3TargetFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutDs3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutDs3TargetFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutJobCompletedNotificationRegistrationSpectraS3Response struct {
    JobCompletedNotificationRegistration JobCompletedNotificationRegistration
    Headers *http.Header
}

func (putJobCompletedNotificationRegistrationSpectraS3Response *PutJobCompletedNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putJobCompletedNotificationRegistrationSpectraS3Response.JobCompletedNotificationRegistration)
}

func NewPutJobCompletedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutJobCompletedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutJobCompletedNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutJobCreatedNotificationRegistrationSpectraS3Response struct {
    JobCreatedNotificationRegistration JobCreatedNotificationRegistration
    Headers *http.Header
}

func (putJobCreatedNotificationRegistrationSpectraS3Response *PutJobCreatedNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putJobCreatedNotificationRegistrationSpectraS3Response.JobCreatedNotificationRegistration)
}

func NewPutJobCreatedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutJobCreatedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutJobCreatedNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutJobCreationFailedNotificationRegistrationSpectraS3Response struct {
    JobCreationFailedNotificationRegistration JobCreationFailedNotificationRegistration
    Headers *http.Header
}

func (putJobCreationFailedNotificationRegistrationSpectraS3Response *PutJobCreationFailedNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putJobCreationFailedNotificationRegistrationSpectraS3Response.JobCreationFailedNotificationRegistration)
}

func NewPutJobCreationFailedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutJobCreationFailedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutJobCreationFailedNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutObjectCachedNotificationRegistrationSpectraS3Response struct {
    S3ObjectCachedNotificationRegistration S3ObjectCachedNotificationRegistration
    Headers *http.Header
}

func (putObjectCachedNotificationRegistrationSpectraS3Response *PutObjectCachedNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putObjectCachedNotificationRegistrationSpectraS3Response.S3ObjectCachedNotificationRegistration)
}

func NewPutObjectCachedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutObjectCachedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutObjectCachedNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutObjectLostNotificationRegistrationSpectraS3Response struct {
    S3ObjectLostNotificationRegistration S3ObjectLostNotificationRegistration
    Headers *http.Header
}

func (putObjectLostNotificationRegistrationSpectraS3Response *PutObjectLostNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putObjectLostNotificationRegistrationSpectraS3Response.S3ObjectLostNotificationRegistration)
}

func NewPutObjectLostNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutObjectLostNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutObjectLostNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutObjectPersistedNotificationRegistrationSpectraS3Response struct {
    S3ObjectPersistedNotificationRegistration S3ObjectPersistedNotificationRegistration
    Headers *http.Header
}

func (putObjectPersistedNotificationRegistrationSpectraS3Response *PutObjectPersistedNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putObjectPersistedNotificationRegistrationSpectraS3Response.S3ObjectPersistedNotificationRegistration)
}

func NewPutObjectPersistedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutObjectPersistedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutObjectPersistedNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutPoolFailureNotificationRegistrationSpectraS3Response struct {
    PoolFailureNotificationRegistration PoolFailureNotificationRegistration
    Headers *http.Header
}

func (putPoolFailureNotificationRegistrationSpectraS3Response *PutPoolFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putPoolFailureNotificationRegistrationSpectraS3Response.PoolFailureNotificationRegistration)
}

func NewPutPoolFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutPoolFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutPoolFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutS3TargetFailureNotificationRegistrationSpectraS3Response struct {
    S3TargetFailureNotificationRegistration S3TargetFailureNotificationRegistration
    Headers *http.Header
}

func (putS3TargetFailureNotificationRegistrationSpectraS3Response *PutS3TargetFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putS3TargetFailureNotificationRegistrationSpectraS3Response.S3TargetFailureNotificationRegistration)
}

func NewPutS3TargetFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutS3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutS3TargetFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutStorageDomainFailureNotificationRegistrationSpectraS3Response struct {
    StorageDomainFailureNotificationRegistration StorageDomainFailureNotificationRegistration
    Headers *http.Header
}

func (putStorageDomainFailureNotificationRegistrationSpectraS3Response *PutStorageDomainFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putStorageDomainFailureNotificationRegistrationSpectraS3Response.StorageDomainFailureNotificationRegistration)
}

func NewPutStorageDomainFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutStorageDomainFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutStorageDomainFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutSystemFailureNotificationRegistrationSpectraS3Response struct {
    SystemFailureNotificationRegistration SystemFailureNotificationRegistration
    Headers *http.Header
}

func (putSystemFailureNotificationRegistrationSpectraS3Response *PutSystemFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putSystemFailureNotificationRegistrationSpectraS3Response.SystemFailureNotificationRegistration)
}

func NewPutSystemFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutSystemFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutSystemFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutTapeFailureNotificationRegistrationSpectraS3Response struct {
    TapeFailureNotificationRegistration TapeFailureNotificationRegistration
    Headers *http.Header
}

func (putTapeFailureNotificationRegistrationSpectraS3Response *PutTapeFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putTapeFailureNotificationRegistrationSpectraS3Response.TapeFailureNotificationRegistration)
}

func NewPutTapeFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutTapeFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutTapeFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutTapePartitionFailureNotificationRegistrationSpectraS3Response struct {
    TapePartitionFailureNotificationRegistration TapePartitionFailureNotificationRegistration
    Headers *http.Header
}

func (putTapePartitionFailureNotificationRegistrationSpectraS3Response *PutTapePartitionFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putTapePartitionFailureNotificationRegistrationSpectraS3Response.TapePartitionFailureNotificationRegistration)
}

func NewPutTapePartitionFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*PutTapePartitionFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutTapePartitionFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteAzureTargetFailureNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteAzureTargetFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteAzureTargetFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteAzureTargetFailureNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteDs3TargetFailureNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteDs3TargetFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteDs3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteDs3TargetFailureNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteJobCompletedNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteJobCompletedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteJobCompletedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteJobCompletedNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteJobCreatedNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteJobCreatedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteJobCreatedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteJobCreatedNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteJobCreationFailedNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteJobCreationFailedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteJobCreationFailedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteJobCreationFailedNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteObjectCachedNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteObjectCachedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteObjectCachedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteObjectCachedNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteObjectLostNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteObjectLostNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteObjectLostNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteObjectLostNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteObjectPersistedNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteObjectPersistedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteObjectPersistedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteObjectPersistedNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeletePoolFailureNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeletePoolFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeletePoolFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeletePoolFailureNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteS3TargetFailureNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteS3TargetFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteS3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteS3TargetFailureNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteStorageDomainFailureNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteStorageDomainFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteStorageDomainFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteStorageDomainFailureNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteSystemFailureNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteSystemFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteSystemFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteSystemFailureNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteTapeFailureNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteTapeFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteTapeFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteTapeFailureNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteTapePartitionFailureNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteTapePartitionFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteTapePartitionFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteTapePartitionFailureNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetAzureTargetFailureNotificationRegistrationSpectraS3Response struct {
    AzureTargetFailureNotificationRegistration AzureTargetFailureNotificationRegistration
    Headers *http.Header
}

func (getAzureTargetFailureNotificationRegistrationSpectraS3Response *GetAzureTargetFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getAzureTargetFailureNotificationRegistrationSpectraS3Response.AzureTargetFailureNotificationRegistration)
}

func NewGetAzureTargetFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetAzureTargetFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetAzureTargetFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetAzureTargetFailureNotificationRegistrationsSpectraS3Response struct {
    AzureTargetFailureNotificationRegistrationList AzureTargetFailureNotificationRegistrationList
    Headers *http.Header
}

func (getAzureTargetFailureNotificationRegistrationsSpectraS3Response *GetAzureTargetFailureNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getAzureTargetFailureNotificationRegistrationsSpectraS3Response.AzureTargetFailureNotificationRegistrationList)
}

func NewGetAzureTargetFailureNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetAzureTargetFailureNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetAzureTargetFailureNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDs3TargetFailureNotificationRegistrationSpectraS3Response struct {
    Ds3TargetFailureNotificationRegistration Ds3TargetFailureNotificationRegistration
    Headers *http.Header
}

func (getDs3TargetFailureNotificationRegistrationSpectraS3Response *GetDs3TargetFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDs3TargetFailureNotificationRegistrationSpectraS3Response.Ds3TargetFailureNotificationRegistration)
}

func NewGetDs3TargetFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetDs3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDs3TargetFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDs3TargetFailureNotificationRegistrationsSpectraS3Response struct {
    Ds3TargetFailureNotificationRegistrationList Ds3TargetFailureNotificationRegistrationList
    Headers *http.Header
}

func (getDs3TargetFailureNotificationRegistrationsSpectraS3Response *GetDs3TargetFailureNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDs3TargetFailureNotificationRegistrationsSpectraS3Response.Ds3TargetFailureNotificationRegistrationList)
}

func NewGetDs3TargetFailureNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetDs3TargetFailureNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDs3TargetFailureNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobCompletedNotificationRegistrationSpectraS3Response struct {
    JobCompletedNotificationRegistration JobCompletedNotificationRegistration
    Headers *http.Header
}

func (getJobCompletedNotificationRegistrationSpectraS3Response *GetJobCompletedNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobCompletedNotificationRegistrationSpectraS3Response.JobCompletedNotificationRegistration)
}

func NewGetJobCompletedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetJobCompletedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobCompletedNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobCompletedNotificationRegistrationsSpectraS3Response struct {
    JobCompletedNotificationRegistrationList JobCompletedNotificationRegistrationList
    Headers *http.Header
}

func (getJobCompletedNotificationRegistrationsSpectraS3Response *GetJobCompletedNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobCompletedNotificationRegistrationsSpectraS3Response.JobCompletedNotificationRegistrationList)
}

func NewGetJobCompletedNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetJobCompletedNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobCompletedNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobCreatedNotificationRegistrationSpectraS3Response struct {
    JobCreatedNotificationRegistration JobCreatedNotificationRegistration
    Headers *http.Header
}

func (getJobCreatedNotificationRegistrationSpectraS3Response *GetJobCreatedNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobCreatedNotificationRegistrationSpectraS3Response.JobCreatedNotificationRegistration)
}

func NewGetJobCreatedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetJobCreatedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobCreatedNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobCreatedNotificationRegistrationsSpectraS3Response struct {
    JobCreatedNotificationRegistrationList JobCreatedNotificationRegistrationList
    Headers *http.Header
}

func (getJobCreatedNotificationRegistrationsSpectraS3Response *GetJobCreatedNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobCreatedNotificationRegistrationsSpectraS3Response.JobCreatedNotificationRegistrationList)
}

func NewGetJobCreatedNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetJobCreatedNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobCreatedNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobCreationFailedNotificationRegistrationSpectraS3Response struct {
    JobCreationFailedNotificationRegistration JobCreationFailedNotificationRegistration
    Headers *http.Header
}

func (getJobCreationFailedNotificationRegistrationSpectraS3Response *GetJobCreationFailedNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobCreationFailedNotificationRegistrationSpectraS3Response.JobCreationFailedNotificationRegistration)
}

func NewGetJobCreationFailedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetJobCreationFailedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobCreationFailedNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetJobCreationFailedNotificationRegistrationsSpectraS3Response struct {
    JobCreationFailedNotificationRegistrationList JobCreationFailedNotificationRegistrationList
    Headers *http.Header
}

func (getJobCreationFailedNotificationRegistrationsSpectraS3Response *GetJobCreationFailedNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getJobCreationFailedNotificationRegistrationsSpectraS3Response.JobCreationFailedNotificationRegistrationList)
}

func NewGetJobCreationFailedNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetJobCreationFailedNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetJobCreationFailedNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetObjectCachedNotificationRegistrationSpectraS3Response struct {
    S3ObjectCachedNotificationRegistration S3ObjectCachedNotificationRegistration
    Headers *http.Header
}

func (getObjectCachedNotificationRegistrationSpectraS3Response *GetObjectCachedNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getObjectCachedNotificationRegistrationSpectraS3Response.S3ObjectCachedNotificationRegistration)
}

func NewGetObjectCachedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetObjectCachedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetObjectCachedNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetObjectCachedNotificationRegistrationsSpectraS3Response struct {
    S3ObjectCachedNotificationRegistrationList S3ObjectCachedNotificationRegistrationList
    Headers *http.Header
}

func (getObjectCachedNotificationRegistrationsSpectraS3Response *GetObjectCachedNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getObjectCachedNotificationRegistrationsSpectraS3Response.S3ObjectCachedNotificationRegistrationList)
}

func NewGetObjectCachedNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetObjectCachedNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetObjectCachedNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetObjectLostNotificationRegistrationSpectraS3Response struct {
    S3ObjectLostNotificationRegistration S3ObjectLostNotificationRegistration
    Headers *http.Header
}

func (getObjectLostNotificationRegistrationSpectraS3Response *GetObjectLostNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getObjectLostNotificationRegistrationSpectraS3Response.S3ObjectLostNotificationRegistration)
}

func NewGetObjectLostNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetObjectLostNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetObjectLostNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetObjectLostNotificationRegistrationsSpectraS3Response struct {
    S3ObjectLostNotificationRegistrationList S3ObjectLostNotificationRegistrationList
    Headers *http.Header
}

func (getObjectLostNotificationRegistrationsSpectraS3Response *GetObjectLostNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getObjectLostNotificationRegistrationsSpectraS3Response.S3ObjectLostNotificationRegistrationList)
}

func NewGetObjectLostNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetObjectLostNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetObjectLostNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetObjectPersistedNotificationRegistrationSpectraS3Response struct {
    S3ObjectPersistedNotificationRegistration S3ObjectPersistedNotificationRegistration
    Headers *http.Header
}

func (getObjectPersistedNotificationRegistrationSpectraS3Response *GetObjectPersistedNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getObjectPersistedNotificationRegistrationSpectraS3Response.S3ObjectPersistedNotificationRegistration)
}

func NewGetObjectPersistedNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetObjectPersistedNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetObjectPersistedNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetObjectPersistedNotificationRegistrationsSpectraS3Response struct {
    S3ObjectPersistedNotificationRegistrationList S3ObjectPersistedNotificationRegistrationList
    Headers *http.Header
}

func (getObjectPersistedNotificationRegistrationsSpectraS3Response *GetObjectPersistedNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getObjectPersistedNotificationRegistrationsSpectraS3Response.S3ObjectPersistedNotificationRegistrationList)
}

func NewGetObjectPersistedNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetObjectPersistedNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetObjectPersistedNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetPoolFailureNotificationRegistrationSpectraS3Response struct {
    PoolFailureNotificationRegistration PoolFailureNotificationRegistration
    Headers *http.Header
}

func (getPoolFailureNotificationRegistrationSpectraS3Response *GetPoolFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getPoolFailureNotificationRegistrationSpectraS3Response.PoolFailureNotificationRegistration)
}

func NewGetPoolFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetPoolFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetPoolFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetPoolFailureNotificationRegistrationsSpectraS3Response struct {
    PoolFailureNotificationRegistrationList PoolFailureNotificationRegistrationList
    Headers *http.Header
}

func (getPoolFailureNotificationRegistrationsSpectraS3Response *GetPoolFailureNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getPoolFailureNotificationRegistrationsSpectraS3Response.PoolFailureNotificationRegistrationList)
}

func NewGetPoolFailureNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetPoolFailureNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetPoolFailureNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetS3TargetFailureNotificationRegistrationSpectraS3Response struct {
    S3TargetFailureNotificationRegistration S3TargetFailureNotificationRegistration
    Headers *http.Header
}

func (getS3TargetFailureNotificationRegistrationSpectraS3Response *GetS3TargetFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getS3TargetFailureNotificationRegistrationSpectraS3Response.S3TargetFailureNotificationRegistration)
}

func NewGetS3TargetFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetS3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetS3TargetFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetS3TargetFailureNotificationRegistrationsSpectraS3Response struct {
    S3TargetFailureNotificationRegistrationList S3TargetFailureNotificationRegistrationList
    Headers *http.Header
}

func (getS3TargetFailureNotificationRegistrationsSpectraS3Response *GetS3TargetFailureNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getS3TargetFailureNotificationRegistrationsSpectraS3Response.S3TargetFailureNotificationRegistrationList)
}

func NewGetS3TargetFailureNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetS3TargetFailureNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetS3TargetFailureNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetStorageDomainFailureNotificationRegistrationSpectraS3Response struct {
    StorageDomainFailureNotificationRegistration StorageDomainFailureNotificationRegistration
    Headers *http.Header
}

func (getStorageDomainFailureNotificationRegistrationSpectraS3Response *GetStorageDomainFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getStorageDomainFailureNotificationRegistrationSpectraS3Response.StorageDomainFailureNotificationRegistration)
}

func NewGetStorageDomainFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetStorageDomainFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetStorageDomainFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetStorageDomainFailureNotificationRegistrationsSpectraS3Response struct {
    StorageDomainFailureNotificationRegistrationList StorageDomainFailureNotificationRegistrationList
    Headers *http.Header
}

func (getStorageDomainFailureNotificationRegistrationsSpectraS3Response *GetStorageDomainFailureNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getStorageDomainFailureNotificationRegistrationsSpectraS3Response.StorageDomainFailureNotificationRegistrationList)
}

func NewGetStorageDomainFailureNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetStorageDomainFailureNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetStorageDomainFailureNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSystemFailureNotificationRegistrationSpectraS3Response struct {
    SystemFailureNotificationRegistration SystemFailureNotificationRegistration
    Headers *http.Header
}

func (getSystemFailureNotificationRegistrationSpectraS3Response *GetSystemFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSystemFailureNotificationRegistrationSpectraS3Response.SystemFailureNotificationRegistration)
}

func NewGetSystemFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetSystemFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSystemFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSystemFailureNotificationRegistrationsSpectraS3Response struct {
    SystemFailureNotificationRegistrationList SystemFailureNotificationRegistrationList
    Headers *http.Header
}

func (getSystemFailureNotificationRegistrationsSpectraS3Response *GetSystemFailureNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSystemFailureNotificationRegistrationsSpectraS3Response.SystemFailureNotificationRegistrationList)
}

func NewGetSystemFailureNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetSystemFailureNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSystemFailureNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapeFailureNotificationRegistrationSpectraS3Response struct {
    TapeFailureNotificationRegistration TapeFailureNotificationRegistration
    Headers *http.Header
}

func (getTapeFailureNotificationRegistrationSpectraS3Response *GetTapeFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapeFailureNotificationRegistrationSpectraS3Response.TapeFailureNotificationRegistration)
}

func NewGetTapeFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetTapeFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapeFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapeFailureNotificationRegistrationsSpectraS3Response struct {
    TapeFailureNotificationRegistrationList TapeFailureNotificationRegistrationList
    Headers *http.Header
}

func (getTapeFailureNotificationRegistrationsSpectraS3Response *GetTapeFailureNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapeFailureNotificationRegistrationsSpectraS3Response.TapeFailureNotificationRegistrationList)
}

func NewGetTapeFailureNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetTapeFailureNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapeFailureNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapePartitionFailureNotificationRegistrationSpectraS3Response struct {
    TapePartitionFailureNotificationRegistration TapePartitionFailureNotificationRegistration
    Headers *http.Header
}

func (getTapePartitionFailureNotificationRegistrationSpectraS3Response *GetTapePartitionFailureNotificationRegistrationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapePartitionFailureNotificationRegistrationSpectraS3Response.TapePartitionFailureNotificationRegistration)
}

func NewGetTapePartitionFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*GetTapePartitionFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapePartitionFailureNotificationRegistrationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapePartitionFailureNotificationRegistrationsSpectraS3Response struct {
    TapePartitionFailureNotificationRegistrationList TapePartitionFailureNotificationRegistrationList
    Headers *http.Header
}

func (getTapePartitionFailureNotificationRegistrationsSpectraS3Response *GetTapePartitionFailureNotificationRegistrationsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapePartitionFailureNotificationRegistrationsSpectraS3Response.TapePartitionFailureNotificationRegistrationList)
}

func NewGetTapePartitionFailureNotificationRegistrationsSpectraS3Response(webResponse WebResponse) (*GetTapePartitionFailureNotificationRegistrationsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapePartitionFailureNotificationRegistrationsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteFolderRecursivelySpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteFolderRecursivelySpectraS3Response(webResponse WebResponse) (*DeleteFolderRecursivelySpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteFolderRecursivelySpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBlobPersistenceSpectraS3Response struct {
    Content string
    Headers *http.Header
}



func NewGetBlobPersistenceSpectraS3Response(webResponse WebResponse) (*GetBlobPersistenceSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        content, err := getResponseBodyAsString(webResponse)
        if err != nil {
            return nil, err
        }
        return &GetBlobPersistenceSpectraS3Response{Content: content, Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetObjectDetailsSpectraS3Response struct {
    S3Object S3Object
    Headers *http.Header
}

func (getObjectDetailsSpectraS3Response *GetObjectDetailsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getObjectDetailsSpectraS3Response.S3Object)
}

func NewGetObjectDetailsSpectraS3Response(webResponse WebResponse) (*GetObjectDetailsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetObjectDetailsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetObjectsDetailsSpectraS3Response struct {
    S3ObjectList S3ObjectList
    Headers *http.Header
}

func (getObjectsDetailsSpectraS3Response *GetObjectsDetailsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getObjectsDetailsSpectraS3Response.S3ObjectList)
}

func NewGetObjectsDetailsSpectraS3Response(webResponse WebResponse) (*GetObjectsDetailsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetObjectsDetailsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetObjectsWithFullDetailsSpectraS3Response struct {
    DetailedS3ObjectList DetailedS3ObjectList
    Headers *http.Header
}

func (getObjectsWithFullDetailsSpectraS3Response *GetObjectsWithFullDetailsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getObjectsWithFullDetailsSpectraS3Response.DetailedS3ObjectList)
}

func NewGetObjectsWithFullDetailsSpectraS3Response(webResponse WebResponse) (*GetObjectsWithFullDetailsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetObjectsWithFullDetailsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetPhysicalPlacementForObjectsSpectraS3Response struct {
    PhysicalPlacement PhysicalPlacement
    Headers *http.Header
}

func (getPhysicalPlacementForObjectsSpectraS3Response *GetPhysicalPlacementForObjectsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getPhysicalPlacementForObjectsSpectraS3Response.PhysicalPlacement)
}

func NewGetPhysicalPlacementForObjectsSpectraS3Response(webResponse WebResponse) (*GetPhysicalPlacementForObjectsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetPhysicalPlacementForObjectsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response struct {
    BulkObjectList BulkObjectList
    Headers *http.Header
}

func (getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response *GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response.BulkObjectList)
}

func NewGetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response(webResponse WebResponse) (*GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type UndeleteObjectSpectraS3Response struct {
    S3Object S3Object
    Headers *http.Header
}

func (undeleteObjectSpectraS3Response *UndeleteObjectSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &undeleteObjectSpectraS3Response.S3Object)
}

func NewUndeleteObjectSpectraS3Response(webResponse WebResponse) (*UndeleteObjectSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body UndeleteObjectSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyPhysicalPlacementForObjectsSpectraS3Response struct {
    PhysicalPlacement PhysicalPlacement
    Headers *http.Header
}

func (verifyPhysicalPlacementForObjectsSpectraS3Response *VerifyPhysicalPlacementForObjectsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &verifyPhysicalPlacementForObjectsSpectraS3Response.PhysicalPlacement)
}

func NewVerifyPhysicalPlacementForObjectsSpectraS3Response(webResponse WebResponse) (*VerifyPhysicalPlacementForObjectsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body VerifyPhysicalPlacementForObjectsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response struct {
    BulkObjectList BulkObjectList
    Headers *http.Header
}

func (verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response *VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &verifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response.BulkObjectList)
}

func NewVerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response(webResponse WebResponse) (*VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body VerifyPhysicalPlacementForObjectsWithFullDetailsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelImportOnAllPoolsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewCancelImportOnAllPoolsSpectraS3Response(webResponse WebResponse) (*CancelImportOnAllPoolsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelImportOnAllPoolsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelImportPoolSpectraS3Response struct {
    Pool Pool
    Headers *http.Header
}

func (cancelImportPoolSpectraS3Response *CancelImportPoolSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &cancelImportPoolSpectraS3Response.Pool)
}

func NewCancelImportPoolSpectraS3Response(webResponse WebResponse) (*CancelImportPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CancelImportPoolSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelVerifyOnAllPoolsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewCancelVerifyOnAllPoolsSpectraS3Response(webResponse WebResponse) (*CancelVerifyOnAllPoolsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelVerifyOnAllPoolsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelVerifyPoolSpectraS3Response struct {
    Pool Pool
    Headers *http.Header
}

func (cancelVerifyPoolSpectraS3Response *CancelVerifyPoolSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &cancelVerifyPoolSpectraS3Response.Pool)
}

func NewCancelVerifyPoolSpectraS3Response(webResponse WebResponse) (*CancelVerifyPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CancelVerifyPoolSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CompactAllPoolsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewCompactAllPoolsSpectraS3Response(webResponse WebResponse) (*CompactAllPoolsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CompactAllPoolsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CompactPoolSpectraS3Response struct {
    Pool Pool
    Headers *http.Header
}

func (compactPoolSpectraS3Response *CompactPoolSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &compactPoolSpectraS3Response.Pool)
}

func NewCompactPoolSpectraS3Response(webResponse WebResponse) (*CompactPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CompactPoolSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutPoolPartitionSpectraS3Response struct {
    PoolPartition PoolPartition
    Headers *http.Header
}

func (putPoolPartitionSpectraS3Response *PutPoolPartitionSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putPoolPartitionSpectraS3Response.PoolPartition)
}

func NewPutPoolPartitionSpectraS3Response(webResponse WebResponse) (*PutPoolPartitionSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutPoolPartitionSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeallocatePoolSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeallocatePoolSpectraS3Response(webResponse WebResponse) (*DeallocatePoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeallocatePoolSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeletePermanentlyLostPoolSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeletePermanentlyLostPoolSpectraS3Response(webResponse WebResponse) (*DeletePermanentlyLostPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeletePermanentlyLostPoolSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeletePoolFailureSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeletePoolFailureSpectraS3Response(webResponse WebResponse) (*DeletePoolFailureSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeletePoolFailureSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeletePoolPartitionSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeletePoolPartitionSpectraS3Response(webResponse WebResponse) (*DeletePoolPartitionSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeletePoolPartitionSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ForcePoolEnvironmentRefreshSpectraS3Response struct {
    
    Headers *http.Header
}



func NewForcePoolEnvironmentRefreshSpectraS3Response(webResponse WebResponse) (*ForcePoolEnvironmentRefreshSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ForcePoolEnvironmentRefreshSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type FormatAllForeignPoolsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewFormatAllForeignPoolsSpectraS3Response(webResponse WebResponse) (*FormatAllForeignPoolsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &FormatAllForeignPoolsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type FormatForeignPoolSpectraS3Response struct {
    Pool Pool
    Headers *http.Header
}

func (formatForeignPoolSpectraS3Response *FormatForeignPoolSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &formatForeignPoolSpectraS3Response.Pool)
}

func NewFormatForeignPoolSpectraS3Response(webResponse WebResponse) (*FormatForeignPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body FormatForeignPoolSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBlobsOnPoolSpectraS3Response struct {
    BulkObjectList BulkObjectList
    Headers *http.Header
}

func (getBlobsOnPoolSpectraS3Response *GetBlobsOnPoolSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBlobsOnPoolSpectraS3Response.BulkObjectList)
}

func NewGetBlobsOnPoolSpectraS3Response(webResponse WebResponse) (*GetBlobsOnPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBlobsOnPoolSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetPoolFailuresSpectraS3Response struct {
    PoolFailureList PoolFailureList
    Headers *http.Header
}

func (getPoolFailuresSpectraS3Response *GetPoolFailuresSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getPoolFailuresSpectraS3Response.PoolFailureList)
}

func NewGetPoolFailuresSpectraS3Response(webResponse WebResponse) (*GetPoolFailuresSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetPoolFailuresSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetPoolPartitionSpectraS3Response struct {
    PoolPartition PoolPartition
    Headers *http.Header
}

func (getPoolPartitionSpectraS3Response *GetPoolPartitionSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getPoolPartitionSpectraS3Response.PoolPartition)
}

func NewGetPoolPartitionSpectraS3Response(webResponse WebResponse) (*GetPoolPartitionSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetPoolPartitionSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetPoolPartitionsSpectraS3Response struct {
    PoolPartitionList PoolPartitionList
    Headers *http.Header
}

func (getPoolPartitionsSpectraS3Response *GetPoolPartitionsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getPoolPartitionsSpectraS3Response.PoolPartitionList)
}

func NewGetPoolPartitionsSpectraS3Response(webResponse WebResponse) (*GetPoolPartitionsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetPoolPartitionsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetPoolSpectraS3Response struct {
    Pool Pool
    Headers *http.Header
}

func (getPoolSpectraS3Response *GetPoolSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getPoolSpectraS3Response.Pool)
}

func NewGetPoolSpectraS3Response(webResponse WebResponse) (*GetPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetPoolSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetPoolsSpectraS3Response struct {
    PoolList PoolList
    Headers *http.Header
}

func (getPoolsSpectraS3Response *GetPoolsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getPoolsSpectraS3Response.PoolList)
}

func NewGetPoolsSpectraS3Response(webResponse WebResponse) (*GetPoolsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetPoolsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ImportAllPoolsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewImportAllPoolsSpectraS3Response(webResponse WebResponse) (*ImportAllPoolsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ImportAllPoolsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ImportPoolSpectraS3Response struct {
    Pool Pool
    Headers *http.Header
}

func (importPoolSpectraS3Response *ImportPoolSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &importPoolSpectraS3Response.Pool)
}

func NewImportPoolSpectraS3Response(webResponse WebResponse) (*ImportPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ImportPoolSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyAllPoolsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewModifyAllPoolsSpectraS3Response(webResponse WebResponse) (*ModifyAllPoolsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ModifyAllPoolsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyPoolPartitionSpectraS3Response struct {
    PoolPartition PoolPartition
    Headers *http.Header
}

func (modifyPoolPartitionSpectraS3Response *ModifyPoolPartitionSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyPoolPartitionSpectraS3Response.PoolPartition)
}

func NewModifyPoolPartitionSpectraS3Response(webResponse WebResponse) (*ModifyPoolPartitionSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyPoolPartitionSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyPoolSpectraS3Response struct {
    Pool Pool
    Headers *http.Header
}

func (modifyPoolSpectraS3Response *ModifyPoolSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyPoolSpectraS3Response.Pool)
}

func NewModifyPoolSpectraS3Response(webResponse WebResponse) (*ModifyPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyPoolSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyAllPoolsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewVerifyAllPoolsSpectraS3Response(webResponse WebResponse) (*VerifyAllPoolsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &VerifyAllPoolsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyPoolSpectraS3Response struct {
    Pool Pool
    Headers *http.Header
}

func (verifyPoolSpectraS3Response *VerifyPoolSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &verifyPoolSpectraS3Response.Pool)
}

func NewVerifyPoolSpectraS3Response(webResponse WebResponse) (*VerifyPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body VerifyPoolSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ConvertStorageDomainToDs3TargetSpectraS3Response struct {
    
    Headers *http.Header
}



func NewConvertStorageDomainToDs3TargetSpectraS3Response(webResponse WebResponse) (*ConvertStorageDomainToDs3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ConvertStorageDomainToDs3TargetSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutPoolStorageDomainMemberSpectraS3Response struct {
    StorageDomainMember StorageDomainMember
    Headers *http.Header
}

func (putPoolStorageDomainMemberSpectraS3Response *PutPoolStorageDomainMemberSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putPoolStorageDomainMemberSpectraS3Response.StorageDomainMember)
}

func NewPutPoolStorageDomainMemberSpectraS3Response(webResponse WebResponse) (*PutPoolStorageDomainMemberSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutPoolStorageDomainMemberSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutStorageDomainSpectraS3Response struct {
    StorageDomain StorageDomain
    Headers *http.Header
}

func (putStorageDomainSpectraS3Response *PutStorageDomainSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putStorageDomainSpectraS3Response.StorageDomain)
}

func NewPutStorageDomainSpectraS3Response(webResponse WebResponse) (*PutStorageDomainSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutStorageDomainSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutTapeStorageDomainMemberSpectraS3Response struct {
    StorageDomainMember StorageDomainMember
    Headers *http.Header
}

func (putTapeStorageDomainMemberSpectraS3Response *PutTapeStorageDomainMemberSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putTapeStorageDomainMemberSpectraS3Response.StorageDomainMember)
}

func NewPutTapeStorageDomainMemberSpectraS3Response(webResponse WebResponse) (*PutTapeStorageDomainMemberSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutTapeStorageDomainMemberSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteStorageDomainFailureSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteStorageDomainFailureSpectraS3Response(webResponse WebResponse) (*DeleteStorageDomainFailureSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteStorageDomainFailureSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteStorageDomainMemberSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteStorageDomainMemberSpectraS3Response(webResponse WebResponse) (*DeleteStorageDomainMemberSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteStorageDomainMemberSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteStorageDomainSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteStorageDomainSpectraS3Response(webResponse WebResponse) (*DeleteStorageDomainSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteStorageDomainSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetStorageDomainFailuresSpectraS3Response struct {
    StorageDomainFailureList StorageDomainFailureList
    Headers *http.Header
}

func (getStorageDomainFailuresSpectraS3Response *GetStorageDomainFailuresSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getStorageDomainFailuresSpectraS3Response.StorageDomainFailureList)
}

func NewGetStorageDomainFailuresSpectraS3Response(webResponse WebResponse) (*GetStorageDomainFailuresSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetStorageDomainFailuresSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetStorageDomainMemberSpectraS3Response struct {
    StorageDomainMember StorageDomainMember
    Headers *http.Header
}

func (getStorageDomainMemberSpectraS3Response *GetStorageDomainMemberSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getStorageDomainMemberSpectraS3Response.StorageDomainMember)
}

func NewGetStorageDomainMemberSpectraS3Response(webResponse WebResponse) (*GetStorageDomainMemberSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetStorageDomainMemberSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetStorageDomainMembersSpectraS3Response struct {
    StorageDomainMemberList StorageDomainMemberList
    Headers *http.Header
}

func (getStorageDomainMembersSpectraS3Response *GetStorageDomainMembersSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getStorageDomainMembersSpectraS3Response.StorageDomainMemberList)
}

func NewGetStorageDomainMembersSpectraS3Response(webResponse WebResponse) (*GetStorageDomainMembersSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetStorageDomainMembersSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetStorageDomainSpectraS3Response struct {
    StorageDomain StorageDomain
    Headers *http.Header
}

func (getStorageDomainSpectraS3Response *GetStorageDomainSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getStorageDomainSpectraS3Response.StorageDomain)
}

func NewGetStorageDomainSpectraS3Response(webResponse WebResponse) (*GetStorageDomainSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetStorageDomainSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetStorageDomainsSpectraS3Response struct {
    StorageDomainList StorageDomainList
    Headers *http.Header
}

func (getStorageDomainsSpectraS3Response *GetStorageDomainsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getStorageDomainsSpectraS3Response.StorageDomainList)
}

func NewGetStorageDomainsSpectraS3Response(webResponse WebResponse) (*GetStorageDomainsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetStorageDomainsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyStorageDomainMemberSpectraS3Response struct {
    StorageDomainMember StorageDomainMember
    Headers *http.Header
}

func (modifyStorageDomainMemberSpectraS3Response *ModifyStorageDomainMemberSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyStorageDomainMemberSpectraS3Response.StorageDomainMember)
}

func NewModifyStorageDomainMemberSpectraS3Response(webResponse WebResponse) (*ModifyStorageDomainMemberSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyStorageDomainMemberSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyStorageDomainSpectraS3Response struct {
    StorageDomain StorageDomain
    Headers *http.Header
}

func (modifyStorageDomainSpectraS3Response *ModifyStorageDomainSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyStorageDomainSpectraS3Response.StorageDomain)
}

func NewModifyStorageDomainSpectraS3Response(webResponse WebResponse) (*ModifyStorageDomainSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyStorageDomainSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ForceFeatureKeyValidationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewForceFeatureKeyValidationSpectraS3Response(webResponse WebResponse) (*ForceFeatureKeyValidationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ForceFeatureKeyValidationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetFeatureKeysSpectraS3Response struct {
    FeatureKeyList FeatureKeyList
    Headers *http.Header
}

func (getFeatureKeysSpectraS3Response *GetFeatureKeysSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getFeatureKeysSpectraS3Response.FeatureKeyList)
}

func NewGetFeatureKeysSpectraS3Response(webResponse WebResponse) (*GetFeatureKeysSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetFeatureKeysSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSystemFailuresSpectraS3Response struct {
    SystemFailureList SystemFailureList
    Headers *http.Header
}

func (getSystemFailuresSpectraS3Response *GetSystemFailuresSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSystemFailuresSpectraS3Response.SystemFailureList)
}

func NewGetSystemFailuresSpectraS3Response(webResponse WebResponse) (*GetSystemFailuresSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSystemFailuresSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetSystemInformationSpectraS3Response struct {
    SystemInformation SystemInformation
    Headers *http.Header
}

func (getSystemInformationSpectraS3Response *GetSystemInformationSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getSystemInformationSpectraS3Response.SystemInformation)
}

func NewGetSystemInformationSpectraS3Response(webResponse WebResponse) (*GetSystemInformationSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetSystemInformationSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ResetInstanceIdentifierSpectraS3Response struct {
    DataPathBackend DataPathBackend
    Headers *http.Header
}

func (resetInstanceIdentifierSpectraS3Response *ResetInstanceIdentifierSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &resetInstanceIdentifierSpectraS3Response.DataPathBackend)
}

func NewResetInstanceIdentifierSpectraS3Response(webResponse WebResponse) (*ResetInstanceIdentifierSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ResetInstanceIdentifierSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifySystemHealthSpectraS3Response struct {
    HealthVerificationResult HealthVerificationResult
    Headers *http.Header
}

func (verifySystemHealthSpectraS3Response *VerifySystemHealthSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &verifySystemHealthSpectraS3Response.HealthVerificationResult)
}

func NewVerifySystemHealthSpectraS3Response(webResponse WebResponse) (*VerifySystemHealthSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body VerifySystemHealthSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelEjectOnAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (cancelEjectOnAllTapesSpectraS3Response *CancelEjectOnAllTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, cancelEjectOnAllTapesSpectraS3Response.TapeFailureList)
}

func NewCancelEjectOnAllTapesSpectraS3Response(webResponse WebResponse) (*CancelEjectOnAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelEjectOnAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body CancelEjectOnAllTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelEjectTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (cancelEjectTapeSpectraS3Response *CancelEjectTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &cancelEjectTapeSpectraS3Response.Tape)
}

func NewCancelEjectTapeSpectraS3Response(webResponse WebResponse) (*CancelEjectTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CancelEjectTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelFormatOnAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (cancelFormatOnAllTapesSpectraS3Response *CancelFormatOnAllTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, cancelFormatOnAllTapesSpectraS3Response.TapeFailureList)
}

func NewCancelFormatOnAllTapesSpectraS3Response(webResponse WebResponse) (*CancelFormatOnAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelFormatOnAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body CancelFormatOnAllTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelFormatTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (cancelFormatTapeSpectraS3Response *CancelFormatTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &cancelFormatTapeSpectraS3Response.Tape)
}

func NewCancelFormatTapeSpectraS3Response(webResponse WebResponse) (*CancelFormatTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CancelFormatTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelImportOnAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (cancelImportOnAllTapesSpectraS3Response *CancelImportOnAllTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, cancelImportOnAllTapesSpectraS3Response.TapeFailureList)
}

func NewCancelImportOnAllTapesSpectraS3Response(webResponse WebResponse) (*CancelImportOnAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelImportOnAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body CancelImportOnAllTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelImportTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (cancelImportTapeSpectraS3Response *CancelImportTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &cancelImportTapeSpectraS3Response.Tape)
}

func NewCancelImportTapeSpectraS3Response(webResponse WebResponse) (*CancelImportTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CancelImportTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelOnlineOnAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (cancelOnlineOnAllTapesSpectraS3Response *CancelOnlineOnAllTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, cancelOnlineOnAllTapesSpectraS3Response.TapeFailureList)
}

func NewCancelOnlineOnAllTapesSpectraS3Response(webResponse WebResponse) (*CancelOnlineOnAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelOnlineOnAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body CancelOnlineOnAllTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelOnlineTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (cancelOnlineTapeSpectraS3Response *CancelOnlineTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &cancelOnlineTapeSpectraS3Response.Tape)
}

func NewCancelOnlineTapeSpectraS3Response(webResponse WebResponse) (*CancelOnlineTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CancelOnlineTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelVerifyOnAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (cancelVerifyOnAllTapesSpectraS3Response *CancelVerifyOnAllTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, cancelVerifyOnAllTapesSpectraS3Response.TapeFailureList)
}

func NewCancelVerifyOnAllTapesSpectraS3Response(webResponse WebResponse) (*CancelVerifyOnAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &CancelVerifyOnAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body CancelVerifyOnAllTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CancelVerifyTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (cancelVerifyTapeSpectraS3Response *CancelVerifyTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &cancelVerifyTapeSpectraS3Response.Tape)
}

func NewCancelVerifyTapeSpectraS3Response(webResponse WebResponse) (*CancelVerifyTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CancelVerifyTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type CleanTapeDriveSpectraS3Response struct {
    TapeDrive TapeDrive
    Headers *http.Header
}

func (cleanTapeDriveSpectraS3Response *CleanTapeDriveSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &cleanTapeDriveSpectraS3Response.TapeDrive)
}

func NewCleanTapeDriveSpectraS3Response(webResponse WebResponse) (*CleanTapeDriveSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body CleanTapeDriveSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutTapeDensityDirectiveSpectraS3Response struct {
    TapeDensityDirective TapeDensityDirective
    Headers *http.Header
}

func (putTapeDensityDirectiveSpectraS3Response *PutTapeDensityDirectiveSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putTapeDensityDirectiveSpectraS3Response.TapeDensityDirective)
}

func NewPutTapeDensityDirectiveSpectraS3Response(webResponse WebResponse) (*PutTapeDensityDirectiveSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutTapeDensityDirectiveSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeletePermanentlyLostTapeSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeletePermanentlyLostTapeSpectraS3Response(webResponse WebResponse) (*DeletePermanentlyLostTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeletePermanentlyLostTapeSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteTapeDensityDirectiveSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteTapeDensityDirectiveSpectraS3Response(webResponse WebResponse) (*DeleteTapeDensityDirectiveSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteTapeDensityDirectiveSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteTapeDriveSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteTapeDriveSpectraS3Response(webResponse WebResponse) (*DeleteTapeDriveSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteTapeDriveSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteTapeFailureSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteTapeFailureSpectraS3Response(webResponse WebResponse) (*DeleteTapeFailureSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteTapeFailureSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteTapePartitionFailureSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteTapePartitionFailureSpectraS3Response(webResponse WebResponse) (*DeleteTapePartitionFailureSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteTapePartitionFailureSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteTapePartitionSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteTapePartitionSpectraS3Response(webResponse WebResponse) (*DeleteTapePartitionSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteTapePartitionSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type EjectAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (ejectAllTapesSpectraS3Response *EjectAllTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, ejectAllTapesSpectraS3Response.TapeFailureList)
}

func NewEjectAllTapesSpectraS3Response(webResponse WebResponse) (*EjectAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &EjectAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body EjectAllTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type EjectStorageDomainBlobsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewEjectStorageDomainBlobsSpectraS3Response(webResponse WebResponse) (*EjectStorageDomainBlobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &EjectStorageDomainBlobsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type EjectStorageDomainSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (ejectStorageDomainSpectraS3Response *EjectStorageDomainSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, ejectStorageDomainSpectraS3Response.TapeFailureList)
}

func NewEjectStorageDomainSpectraS3Response(webResponse WebResponse) (*EjectStorageDomainSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &EjectStorageDomainSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body EjectStorageDomainSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type EjectTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (ejectTapeSpectraS3Response *EjectTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &ejectTapeSpectraS3Response.Tape)
}

func NewEjectTapeSpectraS3Response(webResponse WebResponse) (*EjectTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body EjectTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ForceTapeEnvironmentRefreshSpectraS3Response struct {
    
    Headers *http.Header
}



func NewForceTapeEnvironmentRefreshSpectraS3Response(webResponse WebResponse) (*ForceTapeEnvironmentRefreshSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ForceTapeEnvironmentRefreshSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type FormatAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (formatAllTapesSpectraS3Response *FormatAllTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, formatAllTapesSpectraS3Response.TapeFailureList)
}

func NewFormatAllTapesSpectraS3Response(webResponse WebResponse) (*FormatAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &FormatAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body FormatAllTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type FormatTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (formatTapeSpectraS3Response *FormatTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &formatTapeSpectraS3Response.Tape)
}

func NewFormatTapeSpectraS3Response(webResponse WebResponse) (*FormatTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body FormatTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBlobsOnTapeSpectraS3Response struct {
    BulkObjectList BulkObjectList
    Headers *http.Header
}

func (getBlobsOnTapeSpectraS3Response *GetBlobsOnTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBlobsOnTapeSpectraS3Response.BulkObjectList)
}

func NewGetBlobsOnTapeSpectraS3Response(webResponse WebResponse) (*GetBlobsOnTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBlobsOnTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapeDensityDirectiveSpectraS3Response struct {
    TapeDensityDirective TapeDensityDirective
    Headers *http.Header
}

func (getTapeDensityDirectiveSpectraS3Response *GetTapeDensityDirectiveSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapeDensityDirectiveSpectraS3Response.TapeDensityDirective)
}

func NewGetTapeDensityDirectiveSpectraS3Response(webResponse WebResponse) (*GetTapeDensityDirectiveSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapeDensityDirectiveSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapeDensityDirectivesSpectraS3Response struct {
    TapeDensityDirectiveList TapeDensityDirectiveList
    Headers *http.Header
}

func (getTapeDensityDirectivesSpectraS3Response *GetTapeDensityDirectivesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapeDensityDirectivesSpectraS3Response.TapeDensityDirectiveList)
}

func NewGetTapeDensityDirectivesSpectraS3Response(webResponse WebResponse) (*GetTapeDensityDirectivesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapeDensityDirectivesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapeDriveSpectraS3Response struct {
    TapeDrive TapeDrive
    Headers *http.Header
}

func (getTapeDriveSpectraS3Response *GetTapeDriveSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapeDriveSpectraS3Response.TapeDrive)
}

func NewGetTapeDriveSpectraS3Response(webResponse WebResponse) (*GetTapeDriveSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapeDriveSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapeDrivesSpectraS3Response struct {
    TapeDriveList TapeDriveList
    Headers *http.Header
}

func (getTapeDrivesSpectraS3Response *GetTapeDrivesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapeDrivesSpectraS3Response.TapeDriveList)
}

func NewGetTapeDrivesSpectraS3Response(webResponse WebResponse) (*GetTapeDrivesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapeDrivesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapeFailuresSpectraS3Response struct {
    DetailedTapeFailureList DetailedTapeFailureList
    Headers *http.Header
}

func (getTapeFailuresSpectraS3Response *GetTapeFailuresSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapeFailuresSpectraS3Response.DetailedTapeFailureList)
}

func NewGetTapeFailuresSpectraS3Response(webResponse WebResponse) (*GetTapeFailuresSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapeFailuresSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapeLibrariesSpectraS3Response struct {
    TapeLibraryList TapeLibraryList
    Headers *http.Header
}

func (getTapeLibrariesSpectraS3Response *GetTapeLibrariesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapeLibrariesSpectraS3Response.TapeLibraryList)
}

func NewGetTapeLibrariesSpectraS3Response(webResponse WebResponse) (*GetTapeLibrariesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapeLibrariesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapeLibrarySpectraS3Response struct {
    TapeLibrary TapeLibrary
    Headers *http.Header
}

func (getTapeLibrarySpectraS3Response *GetTapeLibrarySpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapeLibrarySpectraS3Response.TapeLibrary)
}

func NewGetTapeLibrarySpectraS3Response(webResponse WebResponse) (*GetTapeLibrarySpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapeLibrarySpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapePartitionFailuresSpectraS3Response struct {
    TapePartitionFailureList TapePartitionFailureList
    Headers *http.Header
}

func (getTapePartitionFailuresSpectraS3Response *GetTapePartitionFailuresSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapePartitionFailuresSpectraS3Response.TapePartitionFailureList)
}

func NewGetTapePartitionFailuresSpectraS3Response(webResponse WebResponse) (*GetTapePartitionFailuresSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapePartitionFailuresSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapePartitionSpectraS3Response struct {
    TapePartition TapePartition
    Headers *http.Header
}

func (getTapePartitionSpectraS3Response *GetTapePartitionSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapePartitionSpectraS3Response.TapePartition)
}

func NewGetTapePartitionSpectraS3Response(webResponse WebResponse) (*GetTapePartitionSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapePartitionSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapePartitionWithFullDetailsSpectraS3Response struct {
    DetailedTapePartition DetailedTapePartition
    Headers *http.Header
}

func (getTapePartitionWithFullDetailsSpectraS3Response *GetTapePartitionWithFullDetailsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapePartitionWithFullDetailsSpectraS3Response.DetailedTapePartition)
}

func NewGetTapePartitionWithFullDetailsSpectraS3Response(webResponse WebResponse) (*GetTapePartitionWithFullDetailsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapePartitionWithFullDetailsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapePartitionsSpectraS3Response struct {
    TapePartitionList TapePartitionList
    Headers *http.Header
}

func (getTapePartitionsSpectraS3Response *GetTapePartitionsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapePartitionsSpectraS3Response.TapePartitionList)
}

func NewGetTapePartitionsSpectraS3Response(webResponse WebResponse) (*GetTapePartitionsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapePartitionsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapePartitionsWithFullDetailsSpectraS3Response struct {
    NamedDetailedTapePartitionList NamedDetailedTapePartitionList
    Headers *http.Header
}

func (getTapePartitionsWithFullDetailsSpectraS3Response *GetTapePartitionsWithFullDetailsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapePartitionsWithFullDetailsSpectraS3Response.NamedDetailedTapePartitionList)
}

func NewGetTapePartitionsWithFullDetailsSpectraS3Response(webResponse WebResponse) (*GetTapePartitionsWithFullDetailsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapePartitionsWithFullDetailsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (getTapeSpectraS3Response *GetTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapeSpectraS3Response.Tape)
}

func NewGetTapeSpectraS3Response(webResponse WebResponse) (*GetTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetTapesSpectraS3Response struct {
    TapeList TapeList
    Headers *http.Header
}

func (getTapesSpectraS3Response *GetTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getTapesSpectraS3Response.TapeList)
}

func NewGetTapesSpectraS3Response(webResponse WebResponse) (*GetTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ImportAllTapesSpectraS3Response struct {
    
    Headers *http.Header
}



func NewImportAllTapesSpectraS3Response(webResponse WebResponse) (*ImportAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ImportAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ImportTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (importTapeSpectraS3Response *ImportTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &importTapeSpectraS3Response.Tape)
}

func NewImportTapeSpectraS3Response(webResponse WebResponse) (*ImportTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ImportTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type InspectAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (inspectAllTapesSpectraS3Response *InspectAllTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, inspectAllTapesSpectraS3Response.TapeFailureList)
}

func NewInspectAllTapesSpectraS3Response(webResponse WebResponse) (*InspectAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &InspectAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body InspectAllTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type InspectTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (inspectTapeSpectraS3Response *InspectTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &inspectTapeSpectraS3Response.Tape)
}

func NewInspectTapeSpectraS3Response(webResponse WebResponse) (*InspectTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body InspectTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyAllTapePartitionsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewModifyAllTapePartitionsSpectraS3Response(webResponse WebResponse) (*ModifyAllTapePartitionsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ModifyAllTapePartitionsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyTapeDriveSpectraS3Response struct {
    TapeDrive TapeDrive
    Headers *http.Header
}

func (modifyTapeDriveSpectraS3Response *ModifyTapeDriveSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyTapeDriveSpectraS3Response.TapeDrive)
}

func NewModifyTapeDriveSpectraS3Response(webResponse WebResponse) (*ModifyTapeDriveSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyTapeDriveSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyTapePartitionSpectraS3Response struct {
    TapePartition TapePartition
    Headers *http.Header
}

func (modifyTapePartitionSpectraS3Response *ModifyTapePartitionSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyTapePartitionSpectraS3Response.TapePartition)
}

func NewModifyTapePartitionSpectraS3Response(webResponse WebResponse) (*ModifyTapePartitionSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyTapePartitionSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (modifyTapeSpectraS3Response *ModifyTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyTapeSpectraS3Response.Tape)
}

func NewModifyTapeSpectraS3Response(webResponse WebResponse) (*ModifyTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type OnlineAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (onlineAllTapesSpectraS3Response *OnlineAllTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, onlineAllTapesSpectraS3Response.TapeFailureList)
}

func NewOnlineAllTapesSpectraS3Response(webResponse WebResponse) (*OnlineAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &OnlineAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body OnlineAllTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type OnlineTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (onlineTapeSpectraS3Response *OnlineTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &onlineTapeSpectraS3Response.Tape)
}

func NewOnlineTapeSpectraS3Response(webResponse WebResponse) (*OnlineTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body OnlineTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type RawImportAllTapesSpectraS3Response struct {
    
    Headers *http.Header
}



func NewRawImportAllTapesSpectraS3Response(webResponse WebResponse) (*RawImportAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &RawImportAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type RawImportTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (rawImportTapeSpectraS3Response *RawImportTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &rawImportTapeSpectraS3Response.Tape)
}

func NewRawImportTapeSpectraS3Response(webResponse WebResponse) (*RawImportTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body RawImportTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func (verifyAllTapesSpectraS3Response *VerifyAllTapesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, verifyAllTapesSpectraS3Response.TapeFailureList)
}

func NewVerifyAllTapesSpectraS3Response(webResponse WebResponse) (*VerifyAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &VerifyAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body VerifyAllTapesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyTapeSpectraS3Response struct {
    Tape Tape
    Headers *http.Header
}

func (verifyTapeSpectraS3Response *VerifyTapeSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &verifyTapeSpectraS3Response.Tape)
}

func NewVerifyTapeSpectraS3Response(webResponse WebResponse) (*VerifyTapeSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body VerifyTapeSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ForceTargetEnvironmentRefreshSpectraS3Response struct {
    
    Headers *http.Header
}



func NewForceTargetEnvironmentRefreshSpectraS3Response(webResponse WebResponse) (*ForceTargetEnvironmentRefreshSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ForceTargetEnvironmentRefreshSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutAzureTargetBucketNameSpectraS3Response struct {
    AzureTargetBucketName AzureTargetBucketName
    Headers *http.Header
}

func (putAzureTargetBucketNameSpectraS3Response *PutAzureTargetBucketNameSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putAzureTargetBucketNameSpectraS3Response.AzureTargetBucketName)
}

func NewPutAzureTargetBucketNameSpectraS3Response(webResponse WebResponse) (*PutAzureTargetBucketNameSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutAzureTargetBucketNameSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutAzureTargetReadPreferenceSpectraS3Response struct {
    AzureTargetReadPreference AzureTargetReadPreference
    Headers *http.Header
}

func (putAzureTargetReadPreferenceSpectraS3Response *PutAzureTargetReadPreferenceSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putAzureTargetReadPreferenceSpectraS3Response.AzureTargetReadPreference)
}

func NewPutAzureTargetReadPreferenceSpectraS3Response(webResponse WebResponse) (*PutAzureTargetReadPreferenceSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutAzureTargetReadPreferenceSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteAzureTargetBucketNameSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteAzureTargetBucketNameSpectraS3Response(webResponse WebResponse) (*DeleteAzureTargetBucketNameSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteAzureTargetBucketNameSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteAzureTargetFailureSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteAzureTargetFailureSpectraS3Response(webResponse WebResponse) (*DeleteAzureTargetFailureSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteAzureTargetFailureSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteAzureTargetReadPreferenceSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteAzureTargetReadPreferenceSpectraS3Response(webResponse WebResponse) (*DeleteAzureTargetReadPreferenceSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteAzureTargetReadPreferenceSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteAzureTargetSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteAzureTargetSpectraS3Response(webResponse WebResponse) (*DeleteAzureTargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteAzureTargetSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetAzureTargetBucketNamesSpectraS3Response struct {
    AzureTargetBucketNameList AzureTargetBucketNameList
    Headers *http.Header
}

func (getAzureTargetBucketNamesSpectraS3Response *GetAzureTargetBucketNamesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getAzureTargetBucketNamesSpectraS3Response.AzureTargetBucketNameList)
}

func NewGetAzureTargetBucketNamesSpectraS3Response(webResponse WebResponse) (*GetAzureTargetBucketNamesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetAzureTargetBucketNamesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetAzureTargetFailuresSpectraS3Response struct {
    AzureTargetFailureList AzureTargetFailureList
    Headers *http.Header
}

func (getAzureTargetFailuresSpectraS3Response *GetAzureTargetFailuresSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getAzureTargetFailuresSpectraS3Response.AzureTargetFailureList)
}

func NewGetAzureTargetFailuresSpectraS3Response(webResponse WebResponse) (*GetAzureTargetFailuresSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetAzureTargetFailuresSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetAzureTargetReadPreferenceSpectraS3Response struct {
    AzureTargetReadPreference AzureTargetReadPreference
    Headers *http.Header
}

func (getAzureTargetReadPreferenceSpectraS3Response *GetAzureTargetReadPreferenceSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getAzureTargetReadPreferenceSpectraS3Response.AzureTargetReadPreference)
}

func NewGetAzureTargetReadPreferenceSpectraS3Response(webResponse WebResponse) (*GetAzureTargetReadPreferenceSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetAzureTargetReadPreferenceSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetAzureTargetReadPreferencesSpectraS3Response struct {
    AzureTargetReadPreferenceList AzureTargetReadPreferenceList
    Headers *http.Header
}

func (getAzureTargetReadPreferencesSpectraS3Response *GetAzureTargetReadPreferencesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getAzureTargetReadPreferencesSpectraS3Response.AzureTargetReadPreferenceList)
}

func NewGetAzureTargetReadPreferencesSpectraS3Response(webResponse WebResponse) (*GetAzureTargetReadPreferencesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetAzureTargetReadPreferencesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetAzureTargetSpectraS3Response struct {
    AzureTarget AzureTarget
    Headers *http.Header
}

func (getAzureTargetSpectraS3Response *GetAzureTargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getAzureTargetSpectraS3Response.AzureTarget)
}

func NewGetAzureTargetSpectraS3Response(webResponse WebResponse) (*GetAzureTargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetAzureTargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetAzureTargetsSpectraS3Response struct {
    AzureTargetList AzureTargetList
    Headers *http.Header
}

func (getAzureTargetsSpectraS3Response *GetAzureTargetsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getAzureTargetsSpectraS3Response.AzureTargetList)
}

func NewGetAzureTargetsSpectraS3Response(webResponse WebResponse) (*GetAzureTargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetAzureTargetsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBlobsOnAzureTargetSpectraS3Response struct {
    BulkObjectList BulkObjectList
    Headers *http.Header
}

func (getBlobsOnAzureTargetSpectraS3Response *GetBlobsOnAzureTargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBlobsOnAzureTargetSpectraS3Response.BulkObjectList)
}

func NewGetBlobsOnAzureTargetSpectraS3Response(webResponse WebResponse) (*GetBlobsOnAzureTargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBlobsOnAzureTargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ImportAzureTargetSpectraS3Response struct {
    
    Headers *http.Header
}



func NewImportAzureTargetSpectraS3Response(webResponse WebResponse) (*ImportAzureTargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ImportAzureTargetSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyAllAzureTargetsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewModifyAllAzureTargetsSpectraS3Response(webResponse WebResponse) (*ModifyAllAzureTargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ModifyAllAzureTargetsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyAzureTargetSpectraS3Response struct {
    AzureTarget AzureTarget
    Headers *http.Header
}

func (modifyAzureTargetSpectraS3Response *ModifyAzureTargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyAzureTargetSpectraS3Response.AzureTarget)
}

func NewModifyAzureTargetSpectraS3Response(webResponse WebResponse) (*ModifyAzureTargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyAzureTargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type RegisterAzureTargetSpectraS3Response struct {
    AzureTarget AzureTarget
    Headers *http.Header
}

func (registerAzureTargetSpectraS3Response *RegisterAzureTargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &registerAzureTargetSpectraS3Response.AzureTarget)
}

func NewRegisterAzureTargetSpectraS3Response(webResponse WebResponse) (*RegisterAzureTargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body RegisterAzureTargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyAzureTargetSpectraS3Response struct {
    AzureTarget AzureTarget
    Headers *http.Header
}

func (verifyAzureTargetSpectraS3Response *VerifyAzureTargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &verifyAzureTargetSpectraS3Response.AzureTarget)
}

func NewVerifyAzureTargetSpectraS3Response(webResponse WebResponse) (*VerifyAzureTargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body VerifyAzureTargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutDs3TargetReadPreferenceSpectraS3Response struct {
    Ds3TargetReadPreference Ds3TargetReadPreference
    Headers *http.Header
}

func (putDs3TargetReadPreferenceSpectraS3Response *PutDs3TargetReadPreferenceSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putDs3TargetReadPreferenceSpectraS3Response.Ds3TargetReadPreference)
}

func NewPutDs3TargetReadPreferenceSpectraS3Response(webResponse WebResponse) (*PutDs3TargetReadPreferenceSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutDs3TargetReadPreferenceSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteDs3TargetFailureSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteDs3TargetFailureSpectraS3Response(webResponse WebResponse) (*DeleteDs3TargetFailureSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteDs3TargetFailureSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteDs3TargetReadPreferenceSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteDs3TargetReadPreferenceSpectraS3Response(webResponse WebResponse) (*DeleteDs3TargetReadPreferenceSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteDs3TargetReadPreferenceSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteDs3TargetSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteDs3TargetSpectraS3Response(webResponse WebResponse) (*DeleteDs3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteDs3TargetSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBlobsOnDs3TargetSpectraS3Response struct {
    BulkObjectList BulkObjectList
    Headers *http.Header
}

func (getBlobsOnDs3TargetSpectraS3Response *GetBlobsOnDs3TargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBlobsOnDs3TargetSpectraS3Response.BulkObjectList)
}

func NewGetBlobsOnDs3TargetSpectraS3Response(webResponse WebResponse) (*GetBlobsOnDs3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBlobsOnDs3TargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDs3TargetDataPoliciesSpectraS3Response struct {
    DataPolicyList DataPolicyList
    Headers *http.Header
}

func (getDs3TargetDataPoliciesSpectraS3Response *GetDs3TargetDataPoliciesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDs3TargetDataPoliciesSpectraS3Response.DataPolicyList)
}

func NewGetDs3TargetDataPoliciesSpectraS3Response(webResponse WebResponse) (*GetDs3TargetDataPoliciesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDs3TargetDataPoliciesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDs3TargetFailuresSpectraS3Response struct {
    Ds3TargetFailureList Ds3TargetFailureList
    Headers *http.Header
}

func (getDs3TargetFailuresSpectraS3Response *GetDs3TargetFailuresSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDs3TargetFailuresSpectraS3Response.Ds3TargetFailureList)
}

func NewGetDs3TargetFailuresSpectraS3Response(webResponse WebResponse) (*GetDs3TargetFailuresSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDs3TargetFailuresSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDs3TargetReadPreferenceSpectraS3Response struct {
    Ds3TargetReadPreference Ds3TargetReadPreference
    Headers *http.Header
}

func (getDs3TargetReadPreferenceSpectraS3Response *GetDs3TargetReadPreferenceSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDs3TargetReadPreferenceSpectraS3Response.Ds3TargetReadPreference)
}

func NewGetDs3TargetReadPreferenceSpectraS3Response(webResponse WebResponse) (*GetDs3TargetReadPreferenceSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDs3TargetReadPreferenceSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDs3TargetReadPreferencesSpectraS3Response struct {
    Ds3TargetReadPreferenceList Ds3TargetReadPreferenceList
    Headers *http.Header
}

func (getDs3TargetReadPreferencesSpectraS3Response *GetDs3TargetReadPreferencesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDs3TargetReadPreferencesSpectraS3Response.Ds3TargetReadPreferenceList)
}

func NewGetDs3TargetReadPreferencesSpectraS3Response(webResponse WebResponse) (*GetDs3TargetReadPreferencesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDs3TargetReadPreferencesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDs3TargetSpectraS3Response struct {
    Ds3Target Ds3Target
    Headers *http.Header
}

func (getDs3TargetSpectraS3Response *GetDs3TargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDs3TargetSpectraS3Response.Ds3Target)
}

func NewGetDs3TargetSpectraS3Response(webResponse WebResponse) (*GetDs3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDs3TargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetDs3TargetsSpectraS3Response struct {
    Ds3TargetList Ds3TargetList
    Headers *http.Header
}

func (getDs3TargetsSpectraS3Response *GetDs3TargetsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getDs3TargetsSpectraS3Response.Ds3TargetList)
}

func NewGetDs3TargetsSpectraS3Response(webResponse WebResponse) (*GetDs3TargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetDs3TargetsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyAllDs3TargetsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewModifyAllDs3TargetsSpectraS3Response(webResponse WebResponse) (*ModifyAllDs3TargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ModifyAllDs3TargetsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyDs3TargetSpectraS3Response struct {
    Ds3Target Ds3Target
    Headers *http.Header
}

func (modifyDs3TargetSpectraS3Response *ModifyDs3TargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyDs3TargetSpectraS3Response.Ds3Target)
}

func NewModifyDs3TargetSpectraS3Response(webResponse WebResponse) (*ModifyDs3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyDs3TargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PairBackRegisteredDs3TargetSpectraS3Response struct {
    
    Headers *http.Header
}



func NewPairBackRegisteredDs3TargetSpectraS3Response(webResponse WebResponse) (*PairBackRegisteredDs3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &PairBackRegisteredDs3TargetSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type RegisterDs3TargetSpectraS3Response struct {
    Ds3Target Ds3Target
    Headers *http.Header
}

func (registerDs3TargetSpectraS3Response *RegisterDs3TargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &registerDs3TargetSpectraS3Response.Ds3Target)
}

func NewRegisterDs3TargetSpectraS3Response(webResponse WebResponse) (*RegisterDs3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body RegisterDs3TargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyDs3TargetSpectraS3Response struct {
    Ds3Target Ds3Target
    Headers *http.Header
}

func (verifyDs3TargetSpectraS3Response *VerifyDs3TargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &verifyDs3TargetSpectraS3Response.Ds3Target)
}

func NewVerifyDs3TargetSpectraS3Response(webResponse WebResponse) (*VerifyDs3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body VerifyDs3TargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutS3TargetBucketNameSpectraS3Response struct {
    S3TargetBucketName S3TargetBucketName
    Headers *http.Header
}

func (putS3TargetBucketNameSpectraS3Response *PutS3TargetBucketNameSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putS3TargetBucketNameSpectraS3Response.S3TargetBucketName)
}

func NewPutS3TargetBucketNameSpectraS3Response(webResponse WebResponse) (*PutS3TargetBucketNameSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutS3TargetBucketNameSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type PutS3TargetReadPreferenceSpectraS3Response struct {
    S3TargetReadPreference S3TargetReadPreference
    Headers *http.Header
}

func (putS3TargetReadPreferenceSpectraS3Response *PutS3TargetReadPreferenceSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putS3TargetReadPreferenceSpectraS3Response.S3TargetReadPreference)
}

func NewPutS3TargetReadPreferenceSpectraS3Response(webResponse WebResponse) (*PutS3TargetReadPreferenceSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutS3TargetReadPreferenceSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteS3TargetBucketNameSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteS3TargetBucketNameSpectraS3Response(webResponse WebResponse) (*DeleteS3TargetBucketNameSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteS3TargetBucketNameSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteS3TargetFailureSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteS3TargetFailureSpectraS3Response(webResponse WebResponse) (*DeleteS3TargetFailureSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteS3TargetFailureSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteS3TargetReadPreferenceSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteS3TargetReadPreferenceSpectraS3Response(webResponse WebResponse) (*DeleteS3TargetReadPreferenceSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteS3TargetReadPreferenceSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DeleteS3TargetSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteS3TargetSpectraS3Response(webResponse WebResponse) (*DeleteS3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteS3TargetSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetBlobsOnS3TargetSpectraS3Response struct {
    BulkObjectList BulkObjectList
    Headers *http.Header
}

func (getBlobsOnS3TargetSpectraS3Response *GetBlobsOnS3TargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getBlobsOnS3TargetSpectraS3Response.BulkObjectList)
}

func NewGetBlobsOnS3TargetSpectraS3Response(webResponse WebResponse) (*GetBlobsOnS3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetBlobsOnS3TargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetS3TargetBucketNamesSpectraS3Response struct {
    S3TargetBucketNameList S3TargetBucketNameList
    Headers *http.Header
}

func (getS3TargetBucketNamesSpectraS3Response *GetS3TargetBucketNamesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getS3TargetBucketNamesSpectraS3Response.S3TargetBucketNameList)
}

func NewGetS3TargetBucketNamesSpectraS3Response(webResponse WebResponse) (*GetS3TargetBucketNamesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetS3TargetBucketNamesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetS3TargetFailuresSpectraS3Response struct {
    S3TargetFailureList S3TargetFailureList
    Headers *http.Header
}

func (getS3TargetFailuresSpectraS3Response *GetS3TargetFailuresSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getS3TargetFailuresSpectraS3Response.S3TargetFailureList)
}

func NewGetS3TargetFailuresSpectraS3Response(webResponse WebResponse) (*GetS3TargetFailuresSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetS3TargetFailuresSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetS3TargetReadPreferenceSpectraS3Response struct {
    S3TargetReadPreference S3TargetReadPreference
    Headers *http.Header
}

func (getS3TargetReadPreferenceSpectraS3Response *GetS3TargetReadPreferenceSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getS3TargetReadPreferenceSpectraS3Response.S3TargetReadPreference)
}

func NewGetS3TargetReadPreferenceSpectraS3Response(webResponse WebResponse) (*GetS3TargetReadPreferenceSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetS3TargetReadPreferenceSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetS3TargetReadPreferencesSpectraS3Response struct {
    S3TargetReadPreferenceList S3TargetReadPreferenceList
    Headers *http.Header
}

func (getS3TargetReadPreferencesSpectraS3Response *GetS3TargetReadPreferencesSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getS3TargetReadPreferencesSpectraS3Response.S3TargetReadPreferenceList)
}

func NewGetS3TargetReadPreferencesSpectraS3Response(webResponse WebResponse) (*GetS3TargetReadPreferencesSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetS3TargetReadPreferencesSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetS3TargetSpectraS3Response struct {
    S3Target S3Target
    Headers *http.Header
}

func (getS3TargetSpectraS3Response *GetS3TargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getS3TargetSpectraS3Response.S3Target)
}

func NewGetS3TargetSpectraS3Response(webResponse WebResponse) (*GetS3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetS3TargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetS3TargetsSpectraS3Response struct {
    S3TargetList S3TargetList
    Headers *http.Header
}

func (getS3TargetsSpectraS3Response *GetS3TargetsSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getS3TargetsSpectraS3Response.S3TargetList)
}

func NewGetS3TargetsSpectraS3Response(webResponse WebResponse) (*GetS3TargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetS3TargetsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ImportS3TargetSpectraS3Response struct {
    
    Headers *http.Header
}



func NewImportS3TargetSpectraS3Response(webResponse WebResponse) (*ImportS3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ImportS3TargetSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyAllS3TargetsSpectraS3Response struct {
    
    Headers *http.Header
}



func NewModifyAllS3TargetsSpectraS3Response(webResponse WebResponse) (*ModifyAllS3TargetsSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &ModifyAllS3TargetsSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyS3TargetSpectraS3Response struct {
    S3Target S3Target
    Headers *http.Header
}

func (modifyS3TargetSpectraS3Response *ModifyS3TargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyS3TargetSpectraS3Response.S3Target)
}

func NewModifyS3TargetSpectraS3Response(webResponse WebResponse) (*ModifyS3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyS3TargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type RegisterS3TargetSpectraS3Response struct {
    S3Target S3Target
    Headers *http.Header
}

func (registerS3TargetSpectraS3Response *RegisterS3TargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &registerS3TargetSpectraS3Response.S3Target)
}

func NewRegisterS3TargetSpectraS3Response(webResponse WebResponse) (*RegisterS3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body RegisterS3TargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type VerifyS3TargetSpectraS3Response struct {
    S3Target S3Target
    Headers *http.Header
}

func (verifyS3TargetSpectraS3Response *VerifyS3TargetSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &verifyS3TargetSpectraS3Response.S3Target)
}

func NewVerifyS3TargetSpectraS3Response(webResponse WebResponse) (*VerifyS3TargetSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body VerifyS3TargetSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DelegateCreateUserSpectraS3Response struct {
    SpectraUser SpectraUser
    Headers *http.Header
}

func (delegateCreateUserSpectraS3Response *DelegateCreateUserSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &delegateCreateUserSpectraS3Response.SpectraUser)
}

func NewDelegateCreateUserSpectraS3Response(webResponse WebResponse) (*DelegateCreateUserSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body DelegateCreateUserSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type DelegateDeleteUserSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDelegateDeleteUserSpectraS3Response(webResponse WebResponse) (*DelegateDeleteUserSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DelegateDeleteUserSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetUserSpectraS3Response struct {
    SpectraUser SpectraUser
    Headers *http.Header
}

func (getUserSpectraS3Response *GetUserSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getUserSpectraS3Response.SpectraUser)
}

func NewGetUserSpectraS3Response(webResponse WebResponse) (*GetUserSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetUserSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type GetUsersSpectraS3Response struct {
    SpectraUserList SpectraUserList
    Headers *http.Header
}

func (getUsersSpectraS3Response *GetUsersSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &getUsersSpectraS3Response.SpectraUserList)
}

func NewGetUsersSpectraS3Response(webResponse WebResponse) (*GetUsersSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetUsersSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type ModifyUserSpectraS3Response struct {
    SpectraUser SpectraUser
    Headers *http.Header
}

func (modifyUserSpectraS3Response *ModifyUserSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &modifyUserSpectraS3Response.SpectraUser)
}

func NewModifyUserSpectraS3Response(webResponse WebResponse) (*ModifyUserSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ModifyUserSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

type RegenerateUserSecretKeySpectraS3Response struct {
    SpectraUser SpectraUser
    Headers *http.Header
}

func (regenerateUserSecretKeySpectraS3Response *RegenerateUserSecretKeySpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &regenerateUserSecretKeySpectraS3Response.SpectraUser)
}

func NewRegenerateUserSecretKeySpectraS3Response(webResponse WebResponse) (*RegenerateUserSecretKeySpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body RegenerateUserSecretKeySpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}

