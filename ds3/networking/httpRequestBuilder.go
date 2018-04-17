package networking

import (
    "io"
    "net/http"
    "net/url"
    "strings"
    "spectra/ds3_go_sdk/ds3/models"
    "sort"
)

type HttpRequestBuilder struct {
    reader io.Reader
    contentLength *int64
    queryParams *url.Values
    headers *http.Header
    signatureFields signatureFields
    checksumType models.ChecksumType
}

func NewHttpRequestBuilder() *HttpRequestBuilder {
    return &HttpRequestBuilder{
        queryParams:&url.Values{},
        headers:&http.Header{},
        checksumType:models.NONE,
    }
}

// Internally converts reader with size decorator to limit reader to ensure size is respected
func (builder *HttpRequestBuilder) WithReader(stream models.ReaderWithSizeDecorator) *HttpRequestBuilder {
    streamSize, _ := stream.Size()
    builder.reader = io.LimitReader(stream, streamSize)
    builder.contentLength = &streamSize
    return builder
}

// Internally converts reader with size decorator to limit reader to ensure size is respected
// and adds the closer functionality to the limit reader. The send network will automatically
// close the reader when finished.
func (builder *HttpRequestBuilder) WithReadCloser(stream models.ReadCloserWithSizeDecorator) *HttpRequestBuilder {
    streamSize, _ := stream.Size()
    builder.reader = NewLimitReadCloser(stream)
    builder.contentLength = &streamSize
    return builder
}

func (builder *HttpRequestBuilder) WithHttpVerb(verb string) *HttpRequestBuilder {
    builder.signatureFields.Verb = verb
    return builder
}

func (builder *HttpRequestBuilder) WithPath(path string) *HttpRequestBuilder {
    builder.signatureFields.Path = path
    return builder
}

func (builder *HttpRequestBuilder) WithHeader(key string, value string) *HttpRequestBuilder {
    builder.headers.Add(key, value)
    return builder
}

func (builder *HttpRequestBuilder) WithHeaders(headers map[string]string) *HttpRequestBuilder {
    for key, value := range headers {
        builder.WithHeader(key, value)
    }
    return builder
}

func (builder *HttpRequestBuilder) WithQueryParam(key string, value string) *HttpRequestBuilder {
    builder.queryParams.Set(key, value)
    return builder
}

func (builder *HttpRequestBuilder) WithOptionalQueryParam(key string, value *string) *HttpRequestBuilder {
    if value == nil {
        return builder
    }
    builder.queryParams.Set(key, *value)
    return builder
}

func (builder *HttpRequestBuilder) WithOptionalVoidQueryParam(key string, value bool) *HttpRequestBuilder {
    if value {
        builder.queryParams.Set(key, "")
    }
    return builder
}

func (builder *HttpRequestBuilder) WithChecksum(checksum models.Checksum) *HttpRequestBuilder {
    builder.signatureFields.ContentHash = checksum.ContentHash
    builder.checksumType = checksum.Type
    return builder
}

func (builder *HttpRequestBuilder) WithContentType(contentType string) *HttpRequestBuilder {
    builder.signatureFields.ContentType = contentType
    return builder
}

func (builder *HttpRequestBuilder) Build(conn *ConnectionInfo) (*http.Request, error) {
    httpRequest, err := http.NewRequest(builder.signatureFields.Verb, builder.buildUrl(conn), builder.reader)
    if err != nil {
        return nil, err
    }

    if builder.contentLength != nil {
        httpRequest.ContentLength = *builder.contentLength

        // Special casing for content length == 0.  Go won't include the content length header
        // if the length is 0, but BlackPearl needs the content length header to be there and be 0
        // to create a folder.
        // See https://github.com/golang/go/issues/20257
        if *builder.contentLength == 0 {
            httpRequest.Body = http.NoBody
        }
    }

    builder.signatureFields.Date = getCurrentTime()

    builder.maybeAddAmazonCanonicalHeaders()

    authHeaderVal := builder.signatureFields.BuildAuthHeaderValue(conn.Credentials)

    // Set the http request headers such as authorization and date.
    return builder.addHttpRequestHeaders(httpRequest, authHeaderVal)
}

func (builder *HttpRequestBuilder) buildUrl(conn *ConnectionInfo) string {
    var httpUrl = *conn.Endpoint
    httpUrl.Path = builder.signatureFields.Path
    httpUrl.RawQuery = encodeQueryParams(builder.queryParams)
    return httpUrl.String()
}

func (builder *HttpRequestBuilder) maybeAddAmazonCanonicalHeaders() {
	headerKeys := make([]string, 0)

	for key, value := range *builder.headers {
		lowerCaseKey := strings.ToLower(key)
		if strings.HasPrefix(lowerCaseKey, models.AMZ_META_HEADER) && len(value) > 0 {
			headerKeys = append(headerKeys, key)
		}
	}

	if len(headerKeys) == 0 {
		return
	}

	sort.Strings(headerKeys)

	var stringBuilder strings.Builder

	var httpHeaders map[string][]string = *builder.headers

	for _, headerKey := range headerKeys {
		lowerCaseKey := strings.ToLower(headerKey)
		headerValue := httpHeaders[headerKey]

		if len(headerValue) > 0 {
			stringBuilder.WriteString(lowerCaseKey)
			stringBuilder.WriteString(":")
			stringBuilder.WriteString(strings.Join(headerValue, ","))
			stringBuilder.WriteString("\n")
		}
	}

	canonicalAmazonHeaders := stringBuilder.String()

	if len(canonicalAmazonHeaders) > 0 {
		builder.signatureFields.CanonicalizedAmzHeaders = stringBuilder.String()
	}
}

func (builder *HttpRequestBuilder) addHttpRequestHeaders(httpRequest *http.Request, authHeader string) (*http.Request, error) {
    httpRequest.Header.Add("Date", builder.signatureFields.Date)
    httpRequest.Header.Add("Authorization", authHeader)

    if builder.checksumType != models.NONE {
        checksumKey, err := getChecksumHeaderKey(builder.checksumType)
        if err != nil {
            return nil, err
        }
        httpRequest.Header.Add(checksumKey, builder.signatureFields.ContentHash)
    }

    // Copy the headers from the Ds3Request object.
    for key, val := range *builder.headers {
        httpRequest.Header.Add(key, val[0])
    }
    return httpRequest, nil
}

// Percent encodes query parameters and constructs encoded string.
// Spaces are percent encoded as '%20'
func encodeQueryParams(queryParams *url.Values) string {
    // url.Encode encodes spaces as plus (+), so after urlEncode we replace plus (+) signs
    // with percent encoding for spaces (%20)
    return strings.Replace(queryParams.Encode(), "+", "%20", -1)
}

