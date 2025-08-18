# ListAllAuthenticatedDomainWithUser

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllAuthenticatedDomainWithUser**](ListAllAuthenticatedDomainWithUser.md#ListAllAuthenticatedDomainWithUser) | **Get** /v3/whitelabel/domains/subuser/all | List all the authenticated domains associated with the given user.



## ListAllAuthenticatedDomainWithUser

> []ListAllAuthenticatedDomainWithUser200ResponseInner ListAllAuthenticatedDomainWithUser(ctx, Username)

List all the authenticated domains associated with the given user.

**This endpoint allows you to retrieve all of the authenticated domains that have been assigned to a specific subuser.**   This functionality allows subusers to send mail using their parent's domain. Authenticated domains can be associated with (i.e. assigned to) subusers from a parent account, and a subuser can have up to five associated domains.   To associate an authenticated domain with a subuser, the parent account must first authenticate and validate the domain. The parent may then associate the authenticated domain via the subuser management tools.   When selecting a domain to send email from, SendGrid checks for domains in the following order and chooses the first one that appears in the hierarchy:  1. Domain assigned by the subuser that matches the email's `From` address domain.  2. The subuser's default domain.  3. Domain assigned by the parent user that matches the `From` address domain.  4. Parent user's default domain.  5. sendgrid.net

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAllAuthenticatedDomainWithUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**[]ListAllAuthenticatedDomainWithUser200ResponseInner**](ListAllAuthenticatedDomainWithUser200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

