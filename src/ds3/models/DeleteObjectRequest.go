package models

import "net/url"

type DeleteObjectRequest struct {
    bucketName, objectName string
}

func NewDeleteObjectRequest(bucketName, objectName string) *DeleteObjectRequest {
    return &DeleteObjectRequest{bucketName, objectName}
}

func (DeleteObjectRequest) Verb() HttpVerb {
    return DELETE
}

func (self DeleteObjectRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (self DeleteObjectRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (DeleteObjectRequest) GetContentStream() SizedReadCloser {
    return nil
}


