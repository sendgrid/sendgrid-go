# ValidateReverseDns

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateReverseDns**](ValidateReverseDns.md#ValidateReverseDns) | **Post** /v3/whitelabel/ips/{Id}/validate | Validate a reverse DNS record



## ValidateReverseDns

> ValidateReverseDns200Response ValidateReverseDns(ctx, Idoptional)

Validate a reverse DNS record

**This endpoint allows you to validate a reverse DNS record.**  Always check the `valid` property of the response’s `validation_results.a_record` object. This field will indicate whether it was possible to validate the reverse DNS record. If the `validation_results.a_record.valid` is `false`, this indicates only that Twilio SendGrid could not determine the validity your reverse DNS record — it may still be valid.  If validity couldn’t be determined, you can check the value of `validation_results.a_record.reason` to find out why.  You can retrieve the IDs associated with all your reverse DNS records using the \"Retrieve all reverse DNS records\" endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The ID of the reverse DNS record that you would like to validate.

### Other Parameters

Other parameters are passed through a pointer to a ValidateReverseDnsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Onbehalfof** | **string** | The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.

### Return type

[**ValidateReverseDns200Response**](ValidateReverseDns200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

