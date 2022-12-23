# v1.30.1 (2022-09-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.30.0 (2022-09-14)

* **Feature**: Fixed a bug in the API client generation which caused some operation parameters to be incorrectly generated as value types instead of pointer types. The service API always required these affected parameters to be nilable. This fixes the SDK client to match the expectations of the the service API.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.29.0 (2022-09-02)

* **Feature**: This release adds search APIs for Routing Profiles and Queues, which can be used to search for those resources within a Connect Instance.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.2 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.1 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.28.0 (2022-08-19)

* **Feature**: This release adds SearchSecurityProfiles API which can be used to search for Security Profile resources within a Connect Instance.

# v1.27.6 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.5 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.4 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.3 (2022-08-01)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.2 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.1 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.27.0 (2022-06-17)

* **Feature**: This release updates these APIs: UpdateInstanceAttribute, DescribeInstanceAttribute and ListInstanceAttributes. You can use it to programmatically enable/disable High volume outbound communications using attribute type HIGH_VOLUME_OUTBOUND on the specified Amazon Connect instance.

# v1.26.1 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.26.0 (2022-06-06)

* **Feature**: This release adds a new API, GetCurrentUserData, which returns real-time details about users' current activity.

# v1.25.0 (2022-06-02)

* **Feature**: This release adds the following features: 1) New APIs to manage (create, list, update) task template resources, 2) Updates to startTaskContact API to support task templates, and 3) new TransferContact API to programmatically transfer in-progress tasks via a contact flow.

# v1.24.1 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.24.0 (2022-04-28)

* **Feature**: This release introduces an API for changing the current agent status of a user in Connect.

# v1.23.0 (2022-04-25)

* **Feature**: This release adds SearchUsers API which can be used to search for users with a Connect Instance
* **Dependency Update**: Updated to the latest SDK module versions

# v1.22.0 (2022-04-20)

* **Feature**: This release adds APIs to search, claim, release, list, update, and describe phone numbers. You can also use them to associate and disassociate contact flows to phone numbers.

# v1.21.0 (2022-04-01)

* **Feature**: This release updates these APIs: UpdateInstanceAttribute, DescribeInstanceAttribute and ListInstanceAttributes. You can use it to programmatically enable/disable multi-party conferencing using attribute type MULTI_PARTY_CONFERENCING on the specified Amazon Connect instance.

# v1.20.3 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.2 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.1 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.0 (2022-03-11)

* **Feature**: This release adds support for enabling Rich Messaging when starting a new chat session via the StartChatContact API. Rich Messaging enables the following formatting options: bold, italics, hyperlinks, bulleted lists, and numbered lists.

# v1.19.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Feature**: Updated service client model to latest release.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/oss-sdk-go/oss/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.17.0 (2022-01-28)

* **Feature**: Updated to latest API model.

# v1.16.0 (2022-01-14)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.

# v1.13.1 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented oss.EndpointResolverWithOptions from being used by the service client. ([#1514](https://oss-sdk-go/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.13.0 (2021-11-30)

* **Feature**: API client updated

# v1.12.0 (2021-11-19)

* **Feature**: API client updated
* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.0 (2021-11-12)

* **Feature**: Updated service to latest API model.

# v1.10.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Feature**: Updated service to latest API model.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.0 (2021-10-21)

* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.1 (2021-10-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2021-09-30)

* **Feature**: API client updated

# v1.7.1 (2021-09-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2021-08-27)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.1 (2021-08-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.0 (2021-08-12)

* **Feature**: API client updated

# v1.5.2 (2021-08-04)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2021-07-15)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2021-06-25)

* **Feature**: API client updated
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.1 (2021-05-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2021-05-14)

* **Feature**: Constant has been added to modules to enable runtime version inspection for reporting.
* **Feature**: Updated to latest service API model.
* **Dependency Update**: Updated to the latest SDK module versions

