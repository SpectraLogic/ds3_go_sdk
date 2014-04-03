package main

import (
    "log"
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
    cmdErr := commands.RunCommand(buildClientFromArgs(args), args)
    if cmdErr != nil {
        log.Fatal(cmdErr)
    }
}

func buildClientFromArgs(args *commands.Arguments) *ds3.Client {
    return ds3.
        NewBuilder(args.Endpoint, net.Credentials{args.AccessKey, args.SecretKey}).
        WithProxy(args.Proxy).
        Build()
}

