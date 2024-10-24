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
	"fmt"
	"net/http"
	"net/url"

	"strings"
)

type AssociateSubuserWithDomainMultipleParam struct {
	// ID of the authenticated domain to associate with the subuser.
	DomainId *int32 `json:"domain_id"`
	//
	AssociateSubuserWithDomainRequest *AssociateSubuserWithDomainRequest `json:"AssociateSubuserWithDomainRequest,omitempty"`
}

func (params *AssociateSubuserWithDomainMultipleParam) SetDomainId(DomainId int32) *AssociateSubuserWithDomainMultipleParam {
	params.DomainId = &DomainId
	return params
}
func (params *AssociateSubuserWithDomainMultipleParam) SetAssociateSubuserWithDomainRequest(AssociateSubuserWithDomainRequest AssociateSubuserWithDomainRequest) *AssociateSubuserWithDomainMultipleParam {
	params.AssociateSubuserWithDomainRequest = &AssociateSubuserWithDomainRequest
	return params
}

// **This endpoint allows you to associate a specific authenticated domain with a subuser. It can be used to associate up to five authenticated domains.**   This functionality allows subusers to send mail using their parent's domain. Authenticated domains can be associated with (i.e. assigned to) subusers from a parent account. To associate an authenticated domain with a subuser, the parent account must first authenticate and validate the domain. The parent may then associate the authenticated domain via the subuser management tools.   A subuser can have up to five associated authenticated domains. To see the domains that have already been associated with this user, you can [use the API to list the domains currently associated with the subuser](https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/list-the-authenticated-domain-associated-with-a-subuser-multiple).   When selecting a domain to send email from, SendGrid checks for domains in the following order and chooses the first one that appears in the hierarchy:  1. Domain assigned by the subuser that matches the email's `From` address domain.  2. The subuser's default domain.  3. Domain assigned by the parent user that matches the `From` address domain.  4. Parent user's default domain.  5. sendgrid.net
func (c *ApiService) AssociateSubuserWithDomainMultiple(params *AssociateSubuserWithDomainMultipleParam) (interface{}, error) {
	path := "/v3/whitelabel/domains/{DomainId}/subuser:add"
	if params != nil && params.DomainId != nil {
		path = strings.Replace(path, "{"+"DomainId"+"}", fmt.Sprint(*params.DomainId), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.AssociateSubuserWithDomainRequest != nil {
		b, err := json.Marshal(*params.AssociateSubuserWithDomainRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		ps := &AuthenticatedDomainSpf{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}