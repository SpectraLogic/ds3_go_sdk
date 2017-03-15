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
    queryParams *url.Values
}

type CompleteMultipartUpload struct {
    Parts []Part `xml:"Part"`
}

type Part struct {
    PartNumber int
    ETag string
}

func NewCompleteMultipartRequest(bucketName string, objectName string, uploadId string, parts []Part) *CompleteMultipartRequest {
    queryParams := &url.Values{}
    queryParams.Set("uploadId", uploadId)

    return &CompleteMultipartRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
        parts: parts,
        queryParams: queryParams,
    }
}

func (CompleteMultipartRequest) Verb() networking.HttpVerb {
    return networking.POST
}

func (completeMultipartRequest *CompleteMultipartRequest) Path() string {
    return "/" + completeMultipartRequest.bucketName + "/" + completeMultipartRequest.objectName
}

func (completeMultipartRequest *CompleteMultipartRequest) QueryParams() *url.Values {
    return completeMultipartRequest.queryParams
}

func (CompleteMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (completeMultipartRequest *CompleteMultipartRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    // Create an xml document from the entity.
    xmlBytes, err := xml.Marshal(CompleteMultipartUpload{completeMultipartRequest.parts})
    if err != nil {
        panic(err)
    }

    // Create a ReaderWithSizeDecorator which the network layer expects.
    return networking.BuildSizedReadCloser(xmlBytes)
}

func (CompleteMultipartRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}