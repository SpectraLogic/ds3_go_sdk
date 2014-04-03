package main

import (
    "fmt"
    "log"
    "ds3"
    "ds3/net"
    "ds3/models"
    "ds3-client/commands"
)

func main() {
    // Parse the arguments.
    args, argsErr := parseArgs()
    if argsErr != nil {
        log.Fatal(argsErr)
    }

    // Build the client.
    cmdErr := commands.RunCommand(buildClientFromArgs(args), args)
    if cmdErr != nil {
        log.Fatal(cmdErr)
    }

    getServiceResponse, err := client.GetService(models.NewGetServiceRequest())
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Success:")
        fmt.Println(getServiceResponse)
    }
}

func buildClientFromArgs(args *argSet) *ds3.Client {
    return ds3.
        NewBuilder(args.endpoint, net.Credentials{args.accessKey, args.secretKey}).
        WithProxy(args.proxy).
        Build()
}

