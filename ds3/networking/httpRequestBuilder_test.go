package networking

import (
    "fmt"
    "github.com/SpectraLogic/core_go/test_helpers"
    "net/url"
    "testing"
)

func TestBuildUrlEscapingPath(t *testing.T) {
    testCases := []struct {
        input string
        encoded string
    }{
        {input:"-", encoded:"-"},
        {input:".", encoded:"."},
        {input:"_", encoded:"_"},
        {input:"~", encoded:"~"},
        {input:"$", encoded:"$"},
        {input:",", encoded:","},
        {input:"&", encoded:"&"},
        {input:"=", encoded:"="},
        {input:"@", encoded:"@"},
        {input:":", encoded:":"},
        {input:"/", encoded:"/"},
        {input:";", encoded:"%3B"},
        {input:"+", encoded:"%2B"},
        {input:"?", encoded:"%3F"},
        {input:" ", encoded:"%20"},
        {input:"%", encoded:"%25"},
        {input:"#", encoded:"%23"},
        {input:"'", encoded:"%27"},
    }

    const bucketName = "myBucket"

    endPoint, err := url.Parse("http://sm2u-1-10g.eng.sldomain.com")
    test_helpers.FatalOnError(t, err, "unable to parse endpoint")

    connectionInfo := &ConnectionInfo{
        Endpoint: endPoint,
    }

    for _, testCase := range testCases {
        objectName := fmt.Sprintf("object%swithsymbol", testCase.input)

        requestBuilder := NewHttpRequestBuilder().
            WithHttpVerb("PUT").
            WithPath("/" + bucketName + "/" + objectName)

        encodedUrl := requestBuilder.buildUrl(connectionInfo)

        expectedObjectName := fmt.Sprintf("object%swithsymbol", testCase.encoded)
        expected := fmt.Sprintf("%s/%s/%s", endPoint.String(), bucketName, expectedObjectName)
        test_helpers.AssertEqual(t, expected, encodedUrl, fmt.Sprintf("encoding symbol %s", testCase.input))
    }
}
