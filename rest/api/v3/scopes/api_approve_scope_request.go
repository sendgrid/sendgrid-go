/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Scopes API
* The Twilio SendGrid Scopes API allows you to retrieve the scopes or permissions available to a user, see the user's attempts to access your SendGrid account, and, if necessary, deny an access request.
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

type ApproveScopeRequestParam struct {
	// The ID of the request that you want to approve.
	RequestId *string `json:"request_id"`
}

func (params *ApproveScopeRequestParam) SetRequestId(RequestId string) *ApproveScopeRequestParam {
	params.RequestId = &RequestId
	return params
}

// **This endpoint allows you to approve an access attempt.**  **Note:** Only teammate admins may approve another teammate’s access request.
func (c *ApiService) ApproveScopeRequest(params *ApproveScopeRequestParam) (interface{}, error) {
	path := "/v3/scopes/requests/{RequestId}/approve"
	if params != nil && params.RequestId != nil {
		path = strings.Replace(path, "{"+"RequestId"+"}", *params.RequestId, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Patch(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		ps := &ApproveScopeRequest200Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 401 {
		ps := &map[string]interface{}{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &DenyScopeRequest404Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{Status: resp.Status, StatusCode: resp.StatusCode, Body: resp.Body}, nil
}
