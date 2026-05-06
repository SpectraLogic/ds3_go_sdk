package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/SpectraLogic/ds3_go_sdk/ds3/buildclient"
    "github.com/SpectraLogic/ds3_go_sdk/ds3_cli/commands"
)

func main() {
    // Cancel the operation if the user interrupts (Ctrl+C) or the process is asked to terminate.
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

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
    if cmdErr := commands.RunCommand(ctx, client, args); cmdErr != nil {
        log.Fatal(cmdErr)
    }
}
