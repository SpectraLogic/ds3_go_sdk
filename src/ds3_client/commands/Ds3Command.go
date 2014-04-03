package commands

import (
    "fmt"
    "ds3"
)

func RunCommand(client *ds3.Client, args *Arguments) error {
    switch args.Command {
        case "get_service": return getService(client)
        case "get_bucket": fallthrough
        case "get_object": fallthrough
        case "put_bucket": fallthrough
        case "put_object": fallthrough
        case "delete_bucket": fallthrough
        case "delete_object": fallthrough
        case "get_bulk": fallthrough
        case "put_bulk": fallthrough
        default:
            return fmt.Errorf("Unsupported command: '%s'", args.Command)
    }
}

