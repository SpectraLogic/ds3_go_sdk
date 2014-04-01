package main

import (
    "ds3"
    "ds3/net"
    "ds3/models"
    "net/url"
    "fmt"
)

func main() {
    endpoint, err1 := url.Parse("http://192.168.6.129:8080")
    proxy, err2 := url.Parse("http://localhost:8888")
    if err1 == nil && err2 == nil {
        credentials := net.Credentials{"aGFuc2s=", "hVbJjvJM"}
        client := ds3.
            NewBuilder(endpoint, credentials).
            WithProxy(proxy).
            Build()
        //TODO: error handling
        getServiceResponse, _ := client.GetService(models.NewGetServiceRequest())
        fmt.Println(getServiceResponse)
    }
}

