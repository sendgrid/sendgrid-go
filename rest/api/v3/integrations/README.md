# Go API client for 

Integrations allows you to connect your SendGrid applications with third-party services and forward SendGrid email events to those external applications. Currently, Integrations supports event forwarding to [Segment](https://segment.com/docs). You can use this API to create, delete, view, and update your Integrations.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/sendgrid-oai](https://github.com/twilio/sendgrid-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 
- Build date: 2024-10-24T13:26:01.758803Z[Etc/UTC]
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
*AddIntegration* | [**AddIntegration**](docs/AddIntegration.md#addintegration) | **Post** /v3/marketing/integrations | CreateIntegration
*DeleteIntegration* | [**DeleteIntegration**](docs/DeleteIntegration.md#deleteintegration) | **Delete** /v3/marketing/integrations | DeleteBulkIntegration
*FindIntegrationById* | [**FindIntegrationById**](docs/FindIntegrationById.md#findintegrationbyid) | **Get** /v3/marketing/integrations/{Id} | GetIntegration
*GetIntegrationsByUser* | [**GetIntegrationsByUser**](docs/GetIntegrationsByUser.md#getintegrationsbyuser) | **Get** /v3/marketing/integrations | ListIntegration
*UpdateIntegration* | [**UpdateIntegration**](docs/UpdateIntegration.md#updateintegration) | **Patch** /v3/marketing/integrations/{Id} | UpdateIntegration


## Documentation For Models

 - [AddIntegration400Response](AddIntegration400Response.md)
 - [DeleteIntegration400Response](DeleteIntegration400Response.md)
 - [DeleteIntegration404Response](DeleteIntegration404Response.md)
 - [Destination](Destination.md)
 - [Destination1](Destination1.md)
 - [Destination2](Destination2.md)
 - [Destination3](Destination3.md)
 - [DestinationRegion](DestinationRegion.md)
 - [DestinationRegion1](DestinationRegion1.md)
 - [DestinationRegion2](DestinationRegion2.md)
 - [Forbidden](Forbidden.md)
 - [GetIntegrationsByUser200Response](GetIntegrationsByUser200Response.md)
 - [GetIntegrationsByUser403Response](GetIntegrationsByUser403Response.md)
 - [GetIntegrationsByUser500Response](GetIntegrationsByUser500Response.md)
 - [Id](Id.md)
 - [Integration](Integration.md)
 - [IntegrationFilters](IntegrationFilters.md)
 - [IntegrationInput](IntegrationInput.md)
 - [IntegrationInputFilters](IntegrationInputFilters.md)
 - [IntegrationInputProperties](IntegrationInputProperties.md)
 - [IntegrationNotFound](IntegrationNotFound.md)
 - [IntegrationPatch](IntegrationPatch.md)
 - [IntegrationPatchFilters](IntegrationPatchFilters.md)
 - [IntegrationPatchProperties](IntegrationPatchProperties.md)
 - [IntegrationProperties](IntegrationProperties.md)
 - [InternalError](InternalError.md)
 - [InvalidDeleteRequest](InvalidDeleteRequest.md)
 - [InvalidRequest](InvalidRequest.md)
 - [Items](Items.md)
 - [Items1](Items1.md)
 - [Items2](Items2.md)


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
