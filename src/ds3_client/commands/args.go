package commands

import (
    "errors"
    "os"
    "flag"
)

type Arguments struct {
    Endpoint, Proxy string
    AccessKey, SecretKey string
    Command string
    Bucket string
    KeyPrefix string
}

func ParseArgs() (*Arguments, error) {
    // Parse command line arguments.
    endpointParam := flag.String("endpoint", "", "Specifies the url to the DS3 server.")
    accessKeyParam := flag.String("access_key", "", "Specifies the access_key for the DS3 user.")
    secretKeyParam := flag.String("secret_key", "", "Specifies the secret_key for the DS3 user.")
    proxyParam := flag.String("proxy", "", "Specifies the HTTP proxy to route through.")
    commandParam := flag.String("command", "", "The HTTP call to execute.")
    bucketParam := flag.String("bucket", "", "The name of the bucket to constrict the request to.")
    keyPrefixParam := flag.String("prefix", "", "The key prefix by which to constrain the results.")
    flag.Parse()

    // Build the arguments object.
    args := Arguments{
        Endpoint: getParam(*endpointParam, "DS3_ENDPOINT"),
        AccessKey: getParam(*accessKeyParam, "DS3_ACCESS_KEY"),
        SecretKey: getParam(*secretKeyParam, "DS3_SECRET_KEY"),
        Proxy: getParam(*proxyParam, "DS3_PROXY"),
        Command: *commandParam,
        Bucket: *bucketParam,
        KeyPrefix: *keyPrefixParam,
    }

    // Validate required arguments.
    switch {
        case args.Endpoint == "": return nil, errors.New("Must specify an endpoint.")
        case args.AccessKey == "": return nil, errors.New("Must specify an access key.")
        case args.SecretKey == "": return nil, errors.New("Must specify an secret key.")
        case args.Command == "": return nil, errors.New("Must specify a command.")
        default: return &args, nil
    }
}

func getParam(param, envName string) string {
    env := os.Getenv(envName)
    switch {
        case param != "": return param
        case env != "": return env
        default: return ""
    }
}

