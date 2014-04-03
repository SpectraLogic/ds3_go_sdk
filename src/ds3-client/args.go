package main

import (
    "errors"
    "os"
    "flag"
    "net/url"
    "ds3"
    "ds3/net"
)

type argSet struct {
    endpoint, proxy *url.URL
    accessKey, secretKey string
}

func buildClientFromArgs(args *argSet) *ds3.Client {
    return ds3.
        NewBuilder(args.endpoint, net.Credentials{args.accessKey, args.secretKey}).
        WithProxy(args.proxy).
        Build()
}

func parseArgs() (*argSet, error) {
    // Parse command line arguments.
    endpointParam := flag.String("endpoint", "", "Specifies the url to the DS3 server.")
    accessKeyParam := flag.String("access_key", "", "Specifies the access_key for the DS3 user.")
    secretKeyParam := flag.String("secret_key", "", "Specifies the secret_key for the DS3 user.")
    proxyParam := flag.String("proxy", "", "Specifies the HTTP proxy to route through.")
    flag.Parse()

    // Build arg set.
    return buildArgsFromStrings(
        getParam(*endpointParam, "DS3_ENDPOINT"),
        getParam(*accessKeyParam, "DS3_ACCESS_KEY"),
        getParam(*secretKeyParam, "DS3_SECRET_KEY"),
        getParam(*proxyParam, "DS3_PROXY"),
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

func buildArgsFromStrings(endpoint, accessKey, secretKey, proxy string) (*argSet, error) {
    // Validate required arguments.
    switch {
        case endpoint == "": return nil, errors.New("Must specify an endpoint.")
        case accessKey == "": return nil, errors.New("Must specify an access key.")
        case secretKey == "": return nil, errors.New("Must specify an secret key.")
        default:
    }

    // Set keys.
    args := argSet{
        accessKey: accessKey,
        secretKey: secretKey,
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

