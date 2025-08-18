# EmailDnsRecord

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailDnsRecord**](EmailDnsRecord.md#EmailDnsRecord) | **Post** /v3/whitelabel/dns/email | Email DNS records to a co-worker



## EmailDnsRecord

> EmailDnsRecord(ctx, optional)

Email DNS records to a co-worker

**This endpoint is used to share DNS records with a colleagues**  Use this endpoint to send SendGrid-generated DNS record information to a co-worker so they can enter it into your DNS provider to validate your domain and link branding.   What type of records are sent will depend on whether you have chosen Automated Security or not. When using Automated Security, SendGrid provides you with three CNAME records. If you turn Automated Security off, you are instead given TXT and MX records.  If you pass a `link_id` to this endpoint, the generated email will supply the DNS records necessary to complete [Link Branding](https://sendgrid.com/docs/ui/account-and-settings/how-to-set-up-link-branding/) setup. If you pass a `domain_id` to this endpoint, the generated email will supply the DNS records needed to complete [Domain Authentication](https://sendgrid.com/docs/ui/account-and-settings/how-to-set-up-domain-authentication/). Passing both IDs will generate an email with the records needed to complete both setup steps.  You can retrieve all your domain IDs from the returned `id` fields for each domain using the \"List all authenticated domains\" endpoint. You can retrieve all of your link IDs using the \"Retrieve all branded links\" endpoint.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a EmailDnsRecordParams struct


Name | Type | Description
------------- | ------------- | -------------
**EmailDnsRecordRequest** | [**EmailDnsRecordRequest**](EmailDnsRecordRequest.md) | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

