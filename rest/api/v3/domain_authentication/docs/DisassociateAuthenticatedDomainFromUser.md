# DisassociateAuthenticatedDomainFromUser

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisassociateAuthenticatedDomainFromUser**](DisassociateAuthenticatedDomainFromUser.md#DisassociateAuthenticatedDomainFromUser) | **Delete** /v3/whitelabel/domains/subuser | Disassociate an authenticated domain from a given user.



## DisassociateAuthenticatedDomainFromUser

> DisassociateAuthenticatedDomainFromUser(ctx, optional)

Disassociate an authenticated domain from a given user.

**This endpoint allows you to disassociate a specific authenticated domain from a subuser.**  Authenticated domains can be associated with (i.e. assigned to) subusers from a parent account. This functionality allows subusers to send mail using their parent's domain. To associate an authenticated domain with a subuser, the parent account must first authenticate and validate the domain. The parent may then associate the authenticated domain via the subuser management tools.   Note that if you used the [`/v3/whitelabel/domains/{domain_id}/subuser:add` endpoint](https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/associate-an-authenticated-domain-with-a-subuser-multiple) to add multiple domains to the subuser, you should use the [`/v3/whitelabel/domains/{domain_id}/subuser` endpoint](https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/disassociate-an-authenticated-domain-from-a-subuser-multiple) to disassociate those domains.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DisassociateAuthenticatedDomainFromUserParams struct


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

