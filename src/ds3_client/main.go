package main

import (
    "log"
    "errors"
    "net/url"
    "ds3"
    "ds3/net"
    "ds3_client/commands"
)

func main() {
    // Parse the arguments.
    args, argsErr := commands.ParseArgs()
    if argsErr != nil {
        log.Fatal(argsErr)
    }

    // Build the client.
    client, clientErr := buildClientFromArgs(args)
    if clientErr != nil {
        log.Fatal(clientErr)
    }

    // Run the command
    if cmdErr := commands.RunCommand(client, args); cmdErr != nil {
        log.Fatal(cmdErr)
    }
}

func buildClientFromArgs(args *commands.Arguments) (*ds3.Client, error) {
    // Parse endpoint.
    endpointUrl, endpointErr := url.Parse(args.Endpoint)
    if endpointErr != nil {
        return nil, errors.New("The endpoint format was invalid.")
    }

    // Parse proxy.
    var proxyUrlPtr *url.URL
    if args.Proxy != "" {
        var proxyErr error
        proxyUrlPtr, proxyErr = url.Parse(args.Proxy)
        if proxyErr != nil {
            return nil, errors.New("The proxy format was invalid.")
        }
    }

    // Create the client.
    client := ds3.
        NewBuilder(endpointUrl, net.Credentials{args.AccessKey, args.SecretKey}).
        WithProxy(proxyUrlPtr).
        Build()
    return client, nil
}

