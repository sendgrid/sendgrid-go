# AssociateSubuserWithDomainMultiple

All URIs are relative to *https://api.sendgrid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateSubuserWithDomainMultiple**](AssociateSubuserWithDomainMultiple.md#AssociateSubuserWithDomainMultiple) | **Post** /v3/whitelabel/domains/{DomainId}/subuser:add | Associate an authenticated domain with a given user, for up to five domains.



## AssociateSubuserWithDomainMultiple

> AuthenticatedDomainSpf AssociateSubuserWithDomainMultiple(ctx, DomainIdoptional)

Associate an authenticated domain with a given user, for up to five domains.

**This endpoint allows you to associate a specific authenticated domain with a subuser. It can be used to associate up to five authenticated domains.**   This functionality allows subusers to send mail using their parent's domain. Authenticated domains can be associated with (i.e. assigned to) subusers from a parent account. To associate an authenticated domain with a subuser, the parent account must first authenticate and validate the domain. The parent may then associate the authenticated domain via the subuser management tools.   A subuser can have up to five associated authenticated domains. To see the domains that have already been associated with this user, you can [use the API to list the domains currently associated with the subuser](https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/list-the-authenticated-domain-associated-with-a-subuser-multiple).   When selecting a domain to send email from, SendGrid checks for domains in the following order and chooses the first one that appears in the hierarchy:  1. Domain assigned by the subuser that matches the email's `From` address domain.  2. The subuser's default domain.  3. Domain assigned by the parent user that matches the `From` address domain.  4. Parent user's default domain.  5. sendgrid.net

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DomainId** | **int32** | ID of the authenticated domain to associate with the subuser.

### Other Parameters

Other parameters are passed through a pointer to a AssociateSubuserWithDomainMultipleParams struct


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

