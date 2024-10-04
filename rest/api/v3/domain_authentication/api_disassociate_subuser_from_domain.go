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

type DisassociateSubuserFromDomainParam struct {
	// ID of the authenticated domain to be disassociated with the subuser.
	DomainId *int32 `json:"domain_id"`
	// Username for the subuser to find associated authenticated domain.
	Username *string `json:"username,omitempty"`
}

func (params *DisassociateSubuserFromDomainParam) SetDomainId(DomainId int32) *DisassociateSubuserFromDomainParam {
	params.DomainId = &DomainId
	return params
}
func (params *DisassociateSubuserFromDomainParam) SetUsername(Username string) *DisassociateSubuserFromDomainParam {
	params.Username = &Username
	return params
}

// **This endpoint allows you to disassociate a specific authenticated domain from a subuser, for users with up to five associated domains.**   This functionality allows subusers to send mail using their parent's domain. Authenticated domains can be associated with (i.e. assigned to) subusers kknt, and a subuser can have up to five associated authenticated domains.   You can dissociate an authenticated domain from any subuser that has one or more authenticated domains using this endpoint.
func (c *ApiService) DisassociateSubuserFromDomain(params *DisassociateSubuserFromDomainParam) (interface{}, error) {
	path := "/v3/whitelabel/domains/{DomainId}/subuser"
	if params != nil && params.DomainId != nil {
		path = strings.Replace(path, "{"+"DomainId"+"}", fmt.Sprint(*params.DomainId), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Username != nil {
		data.Set("username", *params.Username)
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		ps := &map[string]interface{}{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
