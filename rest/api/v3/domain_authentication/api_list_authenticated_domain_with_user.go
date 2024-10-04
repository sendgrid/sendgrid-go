/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Domain Authentication API
* The Twilio SendGrid Domain Authentication API allows you to manage your authenticated domains and their settings.  Domain Authentication is a required step when setting up your Twilio SendGrid account because it's essential to ensuring the deliverability of your email. Domain Authentication signals trustworthiness to email inbox providers and your recipients by approving SendGrid to send email on behalf of your domain. For more information, see [**How to Set Up Domain Authentication**](https://sendgrid.com/docs/ui/account-and-settings/how-to-set-up-domain-authentication/).  Each user may have a maximum of 3,000 authenticated domains and 3,000 link brandings. This limit is at the user level, meaning each Subuser belonging to a parent account may have its own 3,000 authenticated domains and 3,000 link brandings.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
)

type ListAuthenticatedDomainWithUserParam struct {
	// Username for the subuser to find associated authenticated domain.
	Username *string `json:"username"`
}

func (params *ListAuthenticatedDomainWithUserParam) SetUsername(Username string) *ListAuthenticatedDomainWithUserParam {
	params.Username = &Username
	return params
}

// **This endpoint allows you to retrieve all of the authenticated domains that have been assigned to a specific subuser.**  Authenticated domains can be associated with (i.e. assigned to) subusers from a parent account. This functionality allows subusers to send mail using their parent's domain. To associate an authenticated domain with a subuser, the parent account must first authenticate and validate the domain. The parent may then associate the authenticated domain via the subuser management tools.   Note that if you used the [`/v3/whitelabel/domains/{domain_id}/subuser:add` endpoint]( https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/associate-an-authenticated-domain-with-a-subuser-multiple) to add multiple domains to the subuser, you can use the [`/v3/whitelabel/domains/subuser/all` endpoint](https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/list-the-authenticated-domain-associated-with-a-subuser-multiple) to list those associated domains.
func (c *ApiService) ListAuthenticatedDomainWithUser(params *ListAuthenticatedDomainWithUserParam) (interface{}, error) {
	path := "/v3/whitelabel/domains/subuser"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Username != nil {
		data.Set("username", *params.Username)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &AuthenticatedDomainSpf{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return nil, nil
}
