# ListVerifiedSenderDomain

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVerifiedSenderDomain**](ListVerifiedSenderDomain.md#ListVerifiedSenderDomain) | **Get** /v3/verified_senders/domains | Domain Warn List



## ListVerifiedSenderDomain

> ListVerifiedSenderDomain200Response ListVerifiedSenderDomain(ctx, )

Domain Warn List

**This endpoint returns a list of domains known to implement DMARC and categorizes them by failure type — hard failure or soft failure**.  Domains listed as hard failures will not deliver mail when used as a [Sender Identity](https://sendgrid.com/docs/for-developers/sending-email/sender-identity/) due to the domain's DMARC policy settings.  For example, using a `yahoo.com` email address as a Sender Identity will likely result in the rejection of your mail. For more information about DMARC, see [Everything about DMARC](https://sendgrid.com/docs/ui/sending-email/dmarc/).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVerifiedSenderDomainParams struct


### Return type

[**ListVerifiedSenderDomain200Response**](ListVerifiedSenderDomain200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

