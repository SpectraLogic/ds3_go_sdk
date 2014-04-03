package commands

import (
    "ds3"
    "ds3_client"
)

func RunCommand(client *ds3.Client, args *ds3_client.Arguments) error {
    switch args.command {
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
            return fmt.Errorf("Unsupported command: '%s'", args.command)
    }
}

