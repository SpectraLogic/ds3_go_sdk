package ds3

//import (
//    "testing"
//    "reflect"
//    "net/url"
//    "ds3/models"
//)
//
//type mockedNet struct {
//    // Test manager.
//    t *testing.T
//
//    // Request data to validate.
//    verb models.HttpVerb
//    path string
//    queryParams *url.Values
//
//    // Response data to return.
//    statusCode int
//    response string
//}
//
//func (self *mockedNet) Invoke(request models.Ds3Request) (*http.Response, error) {
//    verb := request.Verb()
//    if verb != self.verb {
//        self.t.Errorf("Expected verb '%s' but got '%s'.", self.verb, verb)
//    }
//
//    path := request.Path()
//    if path != self.path {
//        self.t.Errorf("Expected path '%s' but got '%s'.", self.path, path)
//    }
//
//    queryParams := request.QueryParams()
//    if reflect.DeepEqual(queryParams, self.queryParams) != true {
//        self.t.Errorf(
//            "Expected query params '%s' but got '%s'.",
//            self.queryParams.Encode(),
//            queryParams.Encode()
//        )
//    }
//}
//
//func TestGetBadService(t *testing.T) {
//    client := 
//}
//
