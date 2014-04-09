package models

import (
    "net/url"
    "encoding/xml"
    "ds3/net"
)

type CompleteMultipartRequest struct {
    bucketName, objectName string
    uploadId string
    parts []Part
}

type CompleteMultipartUpload struct {
    Parts []Part
}

type Part struct {
    PartNumber int
    ETag string
}

func NewCompleteMultipartRequest(bucketName, objectName, uploadId string, parts []Part) *CompleteMultipartRequest {
    return &CompleteMultipartRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
        parts: parts,
    }
}

func (CompleteMultipartRequest) Verb() net.HttpVerb {
    return net.POST
}

func (self CompleteMultipartRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (self CompleteMultipartRequest) QueryParams() *url.Values {
    return &url.Values{"uploadId": []string{self.uploadId}}
}

func (self CompleteMultipartRequest) GetContentStream() net.SizedReadCloser {
    // Create an xml document from the entity.
    xmlBytes, err := xml.Marshal(CompleteMultipartUpload{self.parts})
    if err != nil {
        panic(err)
    }

    // Create a SizedReadCloser which the network layer expects.
    return net.BuildSizedReadCloser(xmlBytes)
}

