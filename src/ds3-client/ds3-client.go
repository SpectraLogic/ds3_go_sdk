package main

import (
    "ds3"
    "ds3/net"
    "ds3/models"
    "net/url"
    "fmt"
)

func main() {
    //proxy, err1 := url.Parse("http://192.168.56.1:8888")
    var proxy *url.URL
    var err1 error
    endpoint, err2 := url.Parse("http://192.168.56.102:8080")
    if err1 == nil && err2 == nil {
        credentials := net.Credentials{"aGFuc2s=", "hVbJjvJM"}
        client := ds3.
            NewBuilder(endpoint, credentials).
            WithProxy(proxy).
            Build()
        //TODO: error handling
        getServiceResponse, err := client.GetService(models.NewGetServiceRequest())
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("Success:")
            fmt.Println(getServiceResponse)
        }
    }
}

