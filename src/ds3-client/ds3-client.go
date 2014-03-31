package main

import (
    "ds3"
    "ds3/net"
    "net/url"
)

func makeClient() *ds3.Client {
    endpoint, err1 := url.Parse("http://192.168.6.129:8080")
    proxy, err2 := url.Parse("http://localhost:8888")
    if err1 == nil && err2 == nil {
        credentials := net.Credentials{"aGFuc2s=", "hVbJjvJM"}
        return ds3.NewBuilder(endpoint, credentials).WithProxy(proxy).Build()
    }
    return nil //TODO: error handling needs to be better
}

func main() {
    makeClient()
}

