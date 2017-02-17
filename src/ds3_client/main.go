package main

import (
    "log"
    "ds3_client/buildclient"
    "ds3_client/commands"
)

func main() {
    // Parse the arguments.
    args, argsErr := commands.ParseArgs()
    if argsErr != nil {
        log.Fatal(argsErr)
    }

    // Build the client.
    client, clientErr := buildclient.FromArgs(args)
    if clientErr != nil {
        log.Fatal(clientErr)
    }

    // Run the command
    if cmdErr := commands.RunCommand(client, args); cmdErr != nil {
        log.Fatal(cmdErr)
    }
}



