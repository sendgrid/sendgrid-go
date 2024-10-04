# SendMail

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendMail**](SendMail.md#SendMail) | **Post** /v3/mail/send | Send Email with Twilio SendGrid.



## SendMail

> SendMail(ctx, optional)

Send Email with Twilio SendGrid.

*The Mail Send operation allows you to send email over SendGrid's v3 Web API*  For an overview of this API, including its features and limitations, please see the [Mail Send API overview page](https://www.twilio.com/docs/sendgrid/api-reference/mail-send)  The overview page also includes links to SendGrid's Email API quickstarts and helper libraries to get you working with this endpoint even faster. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a SendMailParams struct


Name | Type | Description
------------- | ------------- | -------------
**ContentEncoding** | [**ContentEncoding**](ContentEncodingContentEncoding.md) | Use this header when sending a gzip compressed mail body. Mail body compression is available to some high volume accounts. Submit a [request to support](https://support.sendgrid.com/hc/en-us) to have gzip enabled on your account.
**SendMailRequest** | [**SendMailRequest**](SendMailRequest.md) | 

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

