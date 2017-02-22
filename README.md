Spectra S3 Go SDK
-----------------

[![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/SpectraLogic/ds3_go_sdk/blob/master/LICENSE.md)

This is not an officially supported SDK.  This was created as an experiment for a much earlier version of the
Spectra S3 API.  It is not compatible with the most recent version of Spectra S3.

Installing with GB
------------------

Install the latest version of the Go SDK.  You can clone it with the following command:
```bash
git clone https://github.com/SpectraLogic/ds3_go_sdk.git
```

Install the latest version of [GB](https://getgb.io/docs/install/). Make sure to compiled the GB project.

Compile the Go SDK using GB with the commands:
```bash
cd ds3_go_sdk
gb build all
```

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
