# Go API client for 

The Twilio SendGrid Teammates API allows you to add, manage, and remove Teammates, or user accounts, from your SendGrid account. Teammates function like user accounts on the SendGrid account, allowing you to invite additional users to your account with scoped access. You can think of Teammates as SendGrid's approach to enabling [role-based access control](https://en.wikipedia.org/wiki/Role-based_access_control) for your SendGrid account. For even more control over the access to your account, see the [Twilio SendGrid SSO API](https://docs.sendgrid.com/api-reference/single-sign-on-teammates/), which enables SSO-authenticated Teammates to be managed through a SAML 2.0 identity provider.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/sendgrid-oai](https://github.com/twilio/sendgrid-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 
- Build date: 2024-10-24T13:26:07.890176Z[Etc/UTC]
- Build package: com.sendgrid.oai.go.SendgridGoGenerator
For more information, please visit [https://support.sendgrid.com/hc/en-us](https://support.sendgrid.com/hc/en-us)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.sendgrid.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DeletePendingTeammate* | [**DeletePendingTeammate**](docs/DeletePendingTeammate.md#deletependingteammate) | **Delete** /v3/teammates/pending/{Token} | Delete pending teammate
*DeleteTeammate* | [**DeleteTeammate**](docs/DeleteTeammate.md#deleteteammate) | **Delete** /v3/teammates/{Username} | Delete teammate
*GetTeammate* | [**GetTeammate**](docs/GetTeammate.md#getteammate) | **Get** /v3/teammates/{Username} | Retrieve specific teammate
*InviteTeammate* | [**InviteTeammate**](docs/InviteTeammate.md#inviteteammate) | **Post** /v3/teammates | Invite teammate
*ListPendingTeammate* | [**ListPendingTeammate**](docs/ListPendingTeammate.md#listpendingteammate) | **Get** /v3/teammates/pending | Retrieve all pending teammates
*ListSubuserByTemplate* | [**ListSubuserByTemplate**](docs/ListSubuserByTemplate.md#listsubuserbytemplate) | **Get** /v3/teammates/{TeammateName}/subuser_access | Get Teammate Subuser Access
*ListTeammate* | [**ListTeammate**](docs/ListTeammate.md#listteammate) | **Get** /v3/teammates | Retrieve all teammates
*ResendTeammateInvite* | [**ResendTeammateInvite**](docs/ResendTeammateInvite.md#resendteammateinvite) | **Post** /v3/teammates/pending/{Token}/resend | Resend teammate invite
*UpdateTeammate* | [**UpdateTeammate**](docs/UpdateTeammate.md#updateteammate) | **Patch** /v3/teammates/{Username} | Update teammate&#39;s permissions


## Documentation For Models

 - [GetTeammate200Response](GetTeammate200Response.md)
 - [InviteTeammate201Response](InviteTeammate201Response.md)
 - [InviteTeammate400Response](InviteTeammate400Response.md)
 - [InviteTeammate400ResponseErrorsInner](InviteTeammate400ResponseErrorsInner.md)
 - [InviteTeammateRequest](InviteTeammateRequest.md)
 - [ListPendingTeammate200Response](ListPendingTeammate200Response.md)
 - [ListPendingTeammate200ResponseResultInner](ListPendingTeammate200ResponseResultInner.md)
 - [ListSubuserByTemplate200Response](ListSubuserByTemplate200Response.md)
 - [ListSubuserByTemplate200ResponseMetadata](ListSubuserByTemplate200ResponseMetadata.md)
 - [ListSubuserByTemplate200ResponseMetadataNextParams](ListSubuserByTemplate200ResponseMetadataNextParams.md)
 - [ListSubuserByTemplate200ResponseSubuserAccessInner](ListSubuserByTemplate200ResponseSubuserAccessInner.md)
 - [ListSubuserByTemplate400Response](ListSubuserByTemplate400Response.md)
 - [ListSubuserByTemplate400ResponseErrorsInner](ListSubuserByTemplate400ResponseErrorsInner.md)
 - [ListTeammate200Response](ListTeammate200Response.md)
 - [ListTeammate200ResponseResultInner](ListTeammate200ResponseResultInner.md)
 - [PermissionType](PermissionType.md)
 - [ResendTeammateInvite200Response](ResendTeammateInvite200Response.md)
 - [UpdateTeammate200Response](UpdateTeammate200Response.md)
 - [UpdateTeammateRequest](UpdateTeammateRequest.md)
 - [UserType](UserType.md)
 - [UserType1](UserType1.md)
 - [UserType2](UserType2.md)


## Documentation For Authorization



## BearerAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```
