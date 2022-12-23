# v1.22.10 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.9 (2022-09-14)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.8 (2022-09-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.7 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.6 (2022-08-30)

* No change notes available for this release.

# v1.22.5 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.4 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.3 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.2 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.1 (2022-08-01)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.0 (2022-07-18)

* **Feature**: Adding AutoMinorVersionUpgrade in the DescribeReplicationGroups API

# v1.21.3 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.2 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.1 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.0 (2022-05-23)

* **Feature**: Added support for encryption in transit for Memcached clusters. Customers can now launch Memcached cluster with encryption in transit enabled when using Memcached version 1.6.12 or later.

# v1.20.7 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.6 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.5 (2022-04-21)

* **Documentation**: Doc only update for ElastiCache

# v1.20.4 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.3 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.2 (2022-03-23)

* **Documentation**: Doc only update for ElastiCache
* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.1 (2022-03-14)

* **Documentation**: Doc only update for ElastiCache

# v1.20.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Documentation**: Updated service client model to latest release.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/oss-sdk-go/oss/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.0 (2022-01-14)

* **Feature**: Updated API models
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.
* **Feature**: Updated to latest service endpoints

# v1.15.1 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented oss.EndpointResolverWithOptions from being used by the service client. ([#1514](https://oss-sdk-go/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.0 (2021-11-30)

* **Feature**: API client updated

# v1.14.1 (2021-11-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2021-11-12)

* **Feature**: Service clients now support custom endpoints that have an initial URI path defined.
* **Feature**: Waiters now have a `WaitForOutput` method, which can be used to retrieve the output of the successful wait operation. Thank you to [Andrew Haines](https://github.com/haines) for contributing this feature.

# v1.13.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Feature**: Updated service to latest API model.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.12.0 (2021-10-21)

* **Feature**: API client updated
* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.2 (2021-10-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.1 (2021-09-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.0 (2021-09-10)

* **Feature**: API client updated

# v1.10.0 (2021-08-27)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.0 (2021-08-19)

* **Feature**: API client updated
* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.1 (2021-08-04)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2021-07-15)

* **Feature**: The ErrorCode method on generated service error types has been corrected to match the API model.
* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2021-06-25)

* **Feature**: API client updated
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.1 (2021-06-04)

* No change notes available for this release.

# v1.6.0 (2021-05-20)

* **Feature**: API client updated
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2021-05-14)

* **Feature**: Constant has been added to modules to enable runtime version inspection for reporting.
* **Dependency Update**: Updated to the latest SDK module versions
