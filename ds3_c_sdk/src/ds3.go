package ds3_c

//import "ds3"
import "fmt"

func Get() string {
    fmt.Println("bar")
    return "foo"
}

//func new_client(endpoint, accessId, key string) *ds3.Client {
//    return ds3_new_client_extended(endpoint, accessId, key, nil, 5)
//}
//
//func new_client_extended(endpoint, accessId, key, proxy string, redirectRetryCount int) *ds3.Client {
//    client := ds3.
//        NewBuilder(goEndpoint, Credentials{ accessId, key }).
//        WithProxy(proxy).
//        WithRedirectRetryCount(redirectRetryCount).
//        Build()
//    return unsafe.Pointer(&client)
//}

