package models

import (
    "net/url"
    "net/http"
    "encoding/xml"
    "ds3/networking"
)

type CompleteMultipartRequest struct {
    bucketName string
    objectName string
    uploadId string
    parts []Part
}

type CompleteMultipartUpload struct {
    Parts []Part `xml:"Part"`
}

type Part struct {
    PartNumber int
    ETag string
}

func NewCompleteMultipartRequest(bucketName string, objectName string, uploadId string, parts []Part) *CompleteMultipartRequest {
    return &CompleteMultipartRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
        parts: parts,
    }
}

func (CompleteMultipartRequest) Verb() networking.HttpVerb {
    return networking.POST
}

func (completeMultipartRequest *CompleteMultipartRequest) Path() string {
    return "/" + completeMultipartRequest.bucketName + "/" + completeMultipartRequest.objectName
}

func (completeMultipartRequest *CompleteMultipartRequest) QueryParams() *url.Values {
    return &url.Values{"uploadId": []string{completeMultipartRequest.uploadId}}
}

func (CompleteMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (completeMultipartRequest *CompleteMultipartRequest) GetContentStream() networking.SizedReadCloser {
    // Create an xml document from the entity.
    xmlBytes, err := xml.Marshal(CompleteMultipartUpload{completeMultipartRequest.parts})
    if err != nil {
        panic(err)
    }

    // Create a SizedReadCloser which the network layer expects.
    return networking.BuildSizedReadCloser(xmlBytes)
}

func (CompleteMultipartRequest) GetChecksum() string {
    return ""
}

func (CompleteMultipartRequest) GetChecksumType() networking.ChecksumType {
    return networking.NONE
}