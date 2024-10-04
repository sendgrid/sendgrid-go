/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Integrations API
* Integrations allows you to connect your SendGrid applications with third-party services and forward SendGrid email events to those external applications. Currently, Integrations supports event forwarding to [Segment](https://segment.com/docs). You can use this API to create, delete, view, and update your Integrations.
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

type UpdateIntegrationParam struct {
	// The ID of the Integration you would like to update.
	Id *string `json:"id"`
	//
	IntegrationPatch *IntegrationPatch `json:"IntegrationPatch"`
}

func (params *UpdateIntegrationParam) SetId(Id string) *UpdateIntegrationParam {
	params.Id = &Id
	return params
}
func (params *UpdateIntegrationParam) SetIntegrationPatch(IntegrationPatch IntegrationPatch) *UpdateIntegrationParam {
	params.IntegrationPatch = &IntegrationPatch
	return params
}

// This endpoint updates an existing Integration.
func (c *ApiService) UpdateIntegration(params *UpdateIntegrationParam) (interface{}, error) {
	path := "/v3/marketing/integrations/{Id}"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", *params.Id, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.IntegrationPatch != nil {
		b, err := json.Marshal(*params.IntegrationPatch)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Patch(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &Integration{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &AddIntegration400Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 403 {
		ps := &GetIntegrationsByUser403Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &DeleteIntegration404Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 500 {
		ps := &GetIntegrationsByUser500Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
