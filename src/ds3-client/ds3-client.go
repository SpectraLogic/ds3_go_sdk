package main

import (
    "os"
    "fmt"
    "ds3/models"
)

func main() {
    args, err := parseArgs()
    if err != nil {
        fmt.Errorf("%s", err)
        os.Exit(1)
    }

    client := buildClientFromArgs(args)
    //TODO: error handling
    getServiceResponse, err := client.GetService(models.NewGetServiceRequest())
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Success:")
        fmt.Println(getServiceResponse)
    }
}

