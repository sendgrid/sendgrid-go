/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Templates API
* The Twilio SendGrid Templates API allows you to create and manage email templates to be delivered with SendGrid's sending APIs. The templates you create will be available using a template ID that is passed to our sending APIs as part of the request. Each template may then have multiple versions associated with it. Whichever version is set as \"active\" at the time of the request will be sent to your recipients. This system allows you to update a single template's look and feel entirely without modifying your requests to our Mail Send API. For example, you could have a single template for welcome emails. That welcome template could then have a version for each season of the year that's themed appropriately and marked as active during the appropriate season. The template ID passed to our sending APIs never needs to change; you can just modify which version is active.  This API provides operations to create and manage your templates as well as their versions.  Each user can create up to 300 different templates. Templates are specific to accounts and Subusers. Templates created on a parent account will not be accessible from the Subusers' accounts.
*
* NOTE: This class is auto generated by OpenAPI Generator.
* https://openapi-generator.tech
* Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"net/url"

	"strings"
)

type ActivateTemplateVersionParam struct {
	// The ID of the original template
	TemplateId *string `json:"template_id"`
	// The ID of the template version
	VersionId *string `json:"version_id"`
	// The `on-behalf-of` header allows you to make API calls from a parent account on behalf of the parent's Subusers or customer accounts. You will use the parent account's API key when using this header. When making a call on behalf of a customer account, the property value should be \"account-id\" followed by the customer account's ID (e.g., `on-behalf-of: account-id <account-id>`). When making a call on behalf of a Subuser, the property value should be the Subuser's username (e.g., `on-behalf-of: <subuser-username>`). See [**On Behalf Of**](https://docs.sendgrid.com/api-reference/how-to-use-the-sendgrid-v3-api/on-behalf-of) for more information.
	Onbehalfof *string `json:"on-behalf-of,omitempty"`
}

func (params *ActivateTemplateVersionParam) SetTemplateId(TemplateId string) *ActivateTemplateVersionParam {
	params.TemplateId = &TemplateId
	return params
}
func (params *ActivateTemplateVersionParam) SetVersionId(VersionId string) *ActivateTemplateVersionParam {
	params.VersionId = &VersionId
	return params
}
func (params *ActivateTemplateVersionParam) SetOnbehalfof(Onbehalfof string) *ActivateTemplateVersionParam {
	params.Onbehalfof = &Onbehalfof
	return params
}

// **This endpoint allows you to activate a version of one of your templates.**
func (c *ApiService) ActivateTemplateVersion(params *ActivateTemplateVersionParam) (interface{}, error) {
	path := "/v3/templates/{TemplateId}/versions/{VersionId}/activate"
	if params != nil && params.TemplateId != nil {
		path = strings.Replace(path, "{"+"TemplateId"+"}", *params.TemplateId, -1)
	}
	if params != nil && params.VersionId != nil {
		path = strings.Replace(path, "{"+"VersionId"+"}", *params.VersionId, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Onbehalfof != nil {
		headers["on-behalf-of"] = *params.Onbehalfof
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &TransactionalTemplateVersionOutput{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}