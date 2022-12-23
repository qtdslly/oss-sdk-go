// Code generated by smithy-go-codegen DO NOT EDIT.

// Package signer provides the API client, operations, and parameter types for AWS
// Signer.
//
// AWS Signer is a fully managed code signing service to help you ensure the trust
// and integrity of your code. AWS Signer supports the following applications: With
// code signing for AWS Lambda, you can sign AWS Lambda deployment packages.
// Integrated support is provided for Amazon S3, Amazon CloudWatch, and AWS
// CloudTrail. In order to sign code, you create a signing profile and then use
// Signer to sign Lambda zip files in S3. With code signing for IoT, you can sign
// code for any IoT device that is supported by oss. IoT code signing is available
// for Amazon FreeRTOS (http://docs.aws.amazon.com/freertos/latest/userguide/) and
// AWS IoT Device Management
// (http://docs.aws.amazon.com/iot/latest/developerguide/), and is integrated with
// AWS Certificate Manager (ACM)
// (http://docs.aws.amazon.com/acm/latest/userguide/). In order to sign code, you
// import a third-party code signing certificate using ACM, and use that to sign
// updates in Amazon FreeRTOS and AWS IoT Device Management. For more information
// about AWS Signer, see the AWS Signer Developer Guide
// (http://docs.aws.amazon.com/signer/latest/developerguide/Welcome.html).
package signer