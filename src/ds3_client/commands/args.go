package commands

import (
    "errors"
    "os"
    "flag"
)

// Represents the parsed command line arguments that we may be interested in.
type Arguments struct {
    Endpoint, Proxy string
    AccessKey, SecretKey string
    Command string
    Bucket string
    Key string
    KeyPrefix string
    MaxKeys int
}

func ParseArgs() (*Arguments, error) {
    // Parse command line arguments.
    endpointParam := flag.String("endpoint", "", "Specifies the url to the DS3 server.")
    accessKeyParam := flag.String("access_key", "", "Specifies the access_key for the DS3 user.")
    secretKeyParam := flag.String("secret_key", "", "Specifies the secret_key for the DS3 user.")
    proxyParam := flag.String("proxy", "", "Specifies the HTTP proxy to route through.")
    commandParam := flag.String("command", "", "The HTTP call to execute.")
    bucketParam := flag.String("bucket", "", "The name of the bucket to constrict the request to.")
    keyParam := flag.String("key", "", "The key for the object to get.")
    keyPrefixParam := flag.String("prefix", "", "The key prefix by which to constrain the results.")
    maxKeysParam := flag.Int("max-keys", 0, "The maximum number of objects to return.")
    flag.Parse()

    // Build the arguments object.
    args := Arguments{
        Endpoint: paramOrEnv(*endpointParam, "DS3_ENDPOINT"),
        AccessKey: paramOrEnv(*accessKeyParam, "DS3_ACCESS_KEY"),
        SecretKey: paramOrEnv(*secretKeyParam, "DS3_SECRET_KEY"),
        Proxy: paramOrEnv(*proxyParam, "DS3_PROXY"),
        Command: *commandParam,
        Bucket: *bucketParam,
        Key: *keyParam,
        KeyPrefix: *keyPrefixParam,
        MaxKeys: *maxKeysParam,
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

func paramOrEnv(param, envName string) string {
    env := os.Getenv(envName)
    switch {
        case param != "": return param
        case env != "": return env
        default: return ""
    }
}

