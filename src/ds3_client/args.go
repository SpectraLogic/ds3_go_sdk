package main

import (
    "errors"
    "os"
    "flag"
    "net/url"
)

type argSet struct {
    endpoint, proxy *url.URL
    accessKey, secretKey string
    command string
}

func parseArgs() (*argSet, error) {
    // Parse command line arguments.
    endpointParam := flag.String("endpoint", "", "Specifies the url to the DS3 server.")
    accessKeyParam := flag.String("access_key", "", "Specifies the access_key for the DS3 user.")
    secretKeyParam := flag.String("secret_key", "", "Specifies the secret_key for the DS3 user.")
    proxyParam := flag.String("proxy", "", "Specifies the HTTP proxy to route through.")
    commandParam := flag.String("command", "", "The HTTP call to execute.")
    flag.Parse()

    // Build arg set.
    return buildArgsFromStrings(
        getParam(*endpointParam, "DS3_ENDPOINT"),
        getParam(*accessKeyParam, "DS3_ACCESS_KEY"),
        getParam(*secretKeyParam, "DS3_SECRET_KEY"),
        getParam(*proxyParam, "DS3_PROXY"),
        *commandParam,
    )
}

func getParam(param, envName string) string {
    env := os.Getenv(envName)
    switch {
        case param != "": return param
        case env != "": return env
        default: return ""
    }
}

func buildArgsFromStrings(endpoint, accessKey, secretKey, proxy, command string) (*argSet, error) {
    // Validate required arguments.
    switch {
        case endpoint == "": return nil, errors.New("Must specify an endpoint.")
        case accessKey == "": return nil, errors.New("Must specify an access key.")
        case secretKey == "": return nil, errors.New("Must specify an secret key.")
        case command == "": return nil, errors.New("Must specify a command.")
        default:
    }

    // Set keys.
    args := argSet{
        accessKey: accessKey,
        secretKey: secretKey,
        command: command,
    }

    // Set endpoint.
    endpointUrl, endpointErr := url.Parse(endpoint)
    if endpointErr == nil {
        args.endpoint = endpointUrl
    } else {
        return nil, errors.New("The endpoint format was invalid.")
    }

    // Set proxy.
    if proxy != "" {
        proxyUrl, proxyErr := url.Parse(proxy)
        if proxyErr == nil {
            args.proxy = proxyUrl
        } else {
            return nil, errors.New("The proxy format was invalid.")
        }
    }

    // Return args.
    return &args, nil
}
