# AssociateSubuserWithDomain

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateSubuserWithDomain**](AssociateSubuserWithDomain.md#AssociateSubuserWithDomain) | **Post** /v3/whitelabel/domains/{DomainId}/subuser | Associate an authenticated domain with a given user.



## AssociateSubuserWithDomain

> AuthenticatedDomainSpf AssociateSubuserWithDomain(ctx, DomainIdoptional)

Associate an authenticated domain with a given user.

**This endpoint allows you to associate a specific authenticated domain with a subuser.** Authenticated domains can be associated with (i.e. assigned to) subusers from a parent account. This functionality allows subusers to send mail using their parent's domain. To associate an authenticated domain with a subuser, the parent account must first authenticate and validate the domain. The parent may then associate the authenticated domain via the subuser management tools.   [You can associate more than one domain with a subuser using the `v3/whitelabel/domains/{domain_id}/subuser:add` endpoint](https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/associate-an-authenticated-domain-with-a-subuser-multiple).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainId** | **int32** | ID of the authenticated domain to associate with the subuser.

### Other Parameters

Other parameters are passed through a pointer to a AssociateSubuserWithDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssociateSubuserWithDomainRequest** | [**AssociateSubuserWithDomainRequest**](AssociateSubuserWithDomainRequest.md) | 

### Return type

[**AuthenticatedDomainSpf**](AuthenticatedDomainSpf.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

