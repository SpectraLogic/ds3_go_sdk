package networking

import (
    "fmt"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_utils/ds3Testing"
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
    ds3Testing.AssertNilError(t, err)

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
        ds3Testing.AssertString(t, fmt.Sprintf("encoding symbol %s", testCase.input), expected, encodedUrl)
    }
}
