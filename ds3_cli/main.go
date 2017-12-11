package main

import (
    "log"
    "spectra/ds3_go_sdk/ds3/buildclient"
    "spectra/ds3_go_sdk/ds3_cli/commands"
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



