package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type GetJobToReplicateSpectraS3Request struct {
    jobId string
    queryParams *url.Values
}

func NewGetJobToReplicateSpectraS3Request(jobId string) *GetJobToReplicateSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("replicate", "")

    return &GetJobToReplicateSpectraS3Request{
        jobId: jobId,
        queryParams: queryParams,
    }
}

func (GetJobToReplicateSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getJobToReplicateSpectraS3Request *GetJobToReplicateSpectraS3Request) Path() string {
    return "/_rest_/job/" + getJobToReplicateSpectraS3Request.jobId
}

func (getJobToReplicateSpectraS3Request *GetJobToReplicateSpectraS3Request) QueryParams() *url.Values {
    return getJobToReplicateSpectraS3Request.queryParams
}

func (GetJobToReplicateSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetJobToReplicateSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}

func (GetJobToReplicateSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}