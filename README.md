Spectra S3 Go SDK
-----------------

[![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/SpectraLogic/ds3_go_sdk/blob/master/LICENSE.md)

An SDK conforming to Spectra S3 for Golang 1.10

Contact Us
----------

Join us at our [Google Groups](https://groups.google.com/d/forum/spectralogicds3-sdks) forum to ask questions, or see frequently asked questions.


Contributing
------------
If you would like to contribute to the source code, sign the 
[Contributors Agreement](https://developer.spectralogic.com/contributors-agreement/).  For an overview of how we use 
Github, please review our [Github Workflow](https://github.com/SpectraLogic/spectralogic.github.com/wiki/Github-Workflow).

Documentation
-------------
The latest documentation is located at [Go SDK Documentation](https://spectralogic.github.io/ds3_go_sdk/index.html).

Installing with GB
------------------

The Go SDK must be installed under `$GOPATH/src/github.com/SpectraLogic`, where `$GOPATH` denotes the location of the Go workspace, and `src` denotes the location of source files within the workspace.
For information regarding Go workspace, see [Golang Workspaces](https://golang.org/doc/code.html#Workspaces).

Install the latest version of the Go SDK in the `src/github.com/SpectraLogic` folder.  You can clone it with the following command:
```bash
cd $GOPATH/src/github.com/SpectraLogic
git clone https://github.com/SpectraLogic/ds3_go_sdk.git
```

Install the latest version of [GB](https://getgb.io/docs/install/). Make sure to compiled the GB project.

Compile the Go SDK using GB with the commands:
```bash
cd ds3_go_sdk
gb build all
```

Client
---------
In the ds3_go_sdk you create a `Client` instance through the setting of the following environment variables and using `buildclient.FromEnv()`.

* `DS3_ENDPOINT` - The URL to the DS3 Endpoint
* `DS3_ACCESS_KEY` - The DS3 access key
* `DS3_SECRET_KEY` - The DS3 secret key
* `http_proxy` - If set, the `Client` instance will proxy through this URL

Examples
--------
All examples are listed in the [samples](samples/) module. All samples can be run from [samples main](samples/samples.go).

* [How to use get service to list buckets](samples/functions/getServiceSample.go)
* [How to create a bucket](samples/functions/getBucketSample.go)
* [How to get a single object using a naked S3 get](samples/functions/getObjectSample.go)
* [How to use bulk put to send multiple files to the BP efficiently](samples/functions/putBulkSample.go)
* [How to use bulk get to retrieve multiple files from the BP efficiently](samples/functions/getBulkSample.go)

Running Tests with GB
---------------------

There are unit tests in the `ds3` package, and integration tests in the `ds3_integration` package. To run tests, cd to
the Go SDK's main folder `ds3_go_sdk`.

In order to run the integration tests the following environment variables must be used to configure which DS3 appliance 
to run the tests against: `DS3_ENDPOINT`, `DS3_ACCESS_KEY`, and `DS3_SECRET_KEY`.

Run all tests (unit and integration) with test output:
```bash
gb test -v
```

Run unit tests with test output:
```bash
gb test -v ds3
```

Run integration tests with test output:
```bash
gb test -v ds3_integration
```
