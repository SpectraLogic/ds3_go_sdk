package buildclient

import (
    "ds3_cli/commands"
    "ds3"
    "net/url"
    "errors"
    "ds3/networking"
    "os"
)

const endpointEnv = "DS3_ENDPOINT"
const accessKeyEnv = "DS3_ACCESS_KEY"
const secretKeyEnv = "DS3_SECRET_KEY"
const proxyEnv = "DS3_PROXY"

// Creates a client from cli arguments
func FromArgs(args *commands.Arguments) (*ds3.Client, error) {
    return createClient(args.Endpoint, args.AccessKey, args.SecretKey, args.Proxy)
}

// Creates a client from environment variables
func FromEnv() (*ds3.Client, error) {
    //Retrieve the environment variables
    endpoint := os.Getenv(endpointEnv)
    accessKey := os.Getenv(accessKeyEnv)
    secretKey := os.Getenv(secretKeyEnv)
    proxy := os.Getenv(proxyEnv)

    // Validate required arguments and create client if all required values are present
    switch {
    case endpoint == "": return nil, errors.New("Must specify an endpoint.")
    case accessKey == "": return nil, errors.New("Must specify an access key.")
    case secretKey == "": return nil, errors.New("Must specify an secret key.")
    default: return createClient(endpoint, accessKey, secretKey, proxy)
    }
}

// Creates a client
func createClient(endpoint, accessKey, secretKey, proxy string) (*ds3.Client, error) {
    // Parse endpoint.
    endpointUrl, endpointErr := url.Parse(endpoint)
    if endpointErr != nil {
        return nil, errors.New("The endpoint format was invalid.")
    }

    // Parse proxy.
    var proxyUrlPtr *url.URL
    if proxy != "" {
        var proxyErr error
        proxyUrlPtr, proxyErr = url.Parse(proxy)
        if proxyErr != nil {
            return nil, errors.New("The proxy format was invalid.")
        }
    }

    // Create the client.
    client := ds3.
    NewClientBuilder(endpointUrl, networking.Credentials{AccessId: accessKey, Key: secretKey}).
        WithProxy(proxyUrlPtr).
        BuildClient()
    return client, nil
}
