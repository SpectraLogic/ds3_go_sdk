package models

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockLogger struct {
	warningMessages []string
}

func (l *mockLogger) Infof(format string, args ...interface{})  {}
func (l *mockLogger) Debugf(format string, args ...interface{}) {}
func (l *mockLogger) Errorf(format string, args ...interface{}) {}
func (l *mockLogger) Warningf(format string, args ...interface{}) {
	l.warningMessages = append(l.warningMessages, fmt.Sprintf(format, args...))
}

type testWebResponse struct {
	body       io.ReadCloser
	statusCode int
}

func (r *testWebResponse) StatusCode() int {
	return r.statusCode
}

func (r *testWebResponse) Body() io.ReadCloser {
	return r.body
}

func (r *testWebResponse) Header() *http.Header {
	return &http.Header{}
}

func TestLoggerPropagation(t *testing.T) {
	// XML with an unknown tag "UnknownTag" for the Bucket model
	responseContent := "<Bucket><Name>test-bucket</Name><UnknownTag>some-value</UnknownTag></Bucket>"
	responseReadCloser := &nopCloser{Reader: strings.NewReader(responseContent)}
	webResponse := &testWebResponse{body: responseReadCloser, statusCode: 200}

	logger := &mockLogger{}
	var bucket Bucket

	// parseResponsePayload will call bucket.parse(root, aggErr, logger)
	err := parseResponsePayload(webResponse, &bucket, logger)

	if err != nil {
		t.Fatalf("expected no error during parsing, but got: %v", err)
	}

	// Verify that the warning was logged
	expectedWarning := "unable to parse unknown xml tag 'UnknownTag' while parsing Bucket."
	found := false
	for _, msg := range logger.warningMessages {
		if msg == expectedWarning {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected warning message '%s' was not logged. Logged warnings: %v", expectedWarning, logger.warningMessages)
	}

	// Verify that the known tag was still parsed correctly
	if bucket.Name == nil || *bucket.Name != "test-bucket" {
		t.Errorf("expected bucket name to be 'test-bucket', but got '%v'", bucket.Name)
	}
}

func TestStringSliceLoggerPropagation(t *testing.T) {
	xmlNodes := []XmlNode{
		{XMLName: struct {
			Space string
			Local string
		}{Local: "ExpectedTag"}, Content: []byte("value1")},
		{XMLName: struct {
			Space string
			Local string
		}{Local: "UnexpectedTag"}, Content: []byte("value2")},
	}

	logger := &mockLogger{}
	var aggErr AggregateError

	result := parseStringSlice("ExpectedTag", xmlNodes, &aggErr, logger)

	if len(result) != 1 || result[0] != "value1" {
		t.Errorf("expected slice result ['value1'], but got %v", result)
	}

	expectedWarning := "Discovered unexpected xml tag 'UnexpectedTag' when expected tag 'ExpectedTag' when parsing string slice."
	found := false
	for _, msg := range logger.warningMessages {
		if msg == expectedWarning {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected warning message '%s' was not logged. Logged warnings: %v", expectedWarning, logger.warningMessages)
	}
}

func TestResponseConstructorLoggerPropagation(t *testing.T) {
	responseContent := "<Bucket><Name>test-bucket</Name><UnknownTag>some-value</UnknownTag></Bucket>"
	responseReadCloser := &nopCloser{Reader: strings.NewReader(responseContent)}
	webResponse := &testWebResponse{body: responseReadCloser, statusCode: 201}

	logger := &mockLogger{}

	// NewPutBucketSpectraS3Response is a generated constructor that should take a logger
	// and pass it down to the Bucket model's parse method.
	_, err := NewPutBucketSpectraS3Response(webResponse, logger)

	if err != nil {
		t.Fatalf("expected no error during response construction, but got: %v", err)
	}

	// Verify that the warning from the underlying model's parse method was logged
	expectedWarning := "unable to parse unknown xml tag 'UnknownTag' while parsing Bucket."
	found := false
	for _, msg := range logger.warningMessages {
		if msg == expectedWarning {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected warning message '%s' was not logged via response constructor. Logged warnings: %v", expectedWarning, logger.warningMessages)
	}
}
