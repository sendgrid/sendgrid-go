# ListAuthenticatedDomainWithUser

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAuthenticatedDomainWithUser**](ListAuthenticatedDomainWithUser.md#ListAuthenticatedDomainWithUser) | **Get** /v3/whitelabel/domains/subuser | List the authenticated domain associated with the given user.



## ListAuthenticatedDomainWithUser

> AuthenticatedDomainSpf ListAuthenticatedDomainWithUser(ctx, Username)

List the authenticated domain associated with the given user.

**This endpoint allows you to retrieve all of the authenticated domains that have been assigned to a specific subuser.**  Authenticated domains can be associated with (i.e. assigned to) subusers from a parent account. This functionality allows subusers to send mail using their parent's domain. To associate an authenticated domain with a subuser, the parent account must first authenticate and validate the domain. The parent may then associate the authenticated domain via the subuser management tools.   Note that if you used the [`/v3/whitelabel/domains/{domain_id}/subuser:add` endpoint]( https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/associate-an-authenticated-domain-with-a-subuser-multiple) to add multiple domains to the subuser, you can use the [`/v3/whitelabel/domains/subuser/all` endpoint](https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/list-the-authenticated-domain-associated-with-a-subuser-multiple) to list those associated domains.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAuthenticatedDomainWithUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AuthenticatedDomainSpf**](AuthenticatedDomainSpf.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

