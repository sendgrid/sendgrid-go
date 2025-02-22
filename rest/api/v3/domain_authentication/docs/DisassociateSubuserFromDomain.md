# DisassociateSubuserFromDomain

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisassociateSubuserFromDomain**](DisassociateSubuserFromDomain.md#DisassociateSubuserFromDomain) | **Delete** /v3/whitelabel/domains/{DomainId}/subuser | Disassociate an authenticated domain from a given user for users with up to five associated domains.



## DisassociateSubuserFromDomain

> DisassociateSubuserFromDomain(ctx, DomainIdoptional)

Disassociate an authenticated domain from a given user for users with up to five associated domains.

**This endpoint allows you to disassociate a specific authenticated domain from a subuser, for users with up to five associated domains.**   This functionality allows subusers to send mail using their parent's domain. Authenticated domains can be associated with (i.e. assigned to) subusers kknt, and a subuser can have up to five associated authenticated domains.   You can dissociate an authenticated domain from any subuser that has one or more authenticated domains using this endpoint.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainId** | **int32** | ID of the authenticated domain to be disassociated with the subuser.

### Other Parameters

Other parameters are passed through a pointer to a DisassociateSubuserFromDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**Username** | **string** | Username for the subuser to find associated authenticated domain.

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

