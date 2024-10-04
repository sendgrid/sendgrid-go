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

type DenyScopeRequestParam struct {
	// The ID of the request that you want to deny.
	RequestId *string `json:"request_id"`
}

func (params *DenyScopeRequestParam) SetRequestId(RequestId string) *DenyScopeRequestParam {
	params.RequestId = &RequestId
	return params
}

// **This endpoint allows you to deny an attempt to access your account.**  **Note:** Only teammate admins may delete a teammate's access request.
func (c *ApiService) DenyScopeRequest(params *DenyScopeRequestParam) (interface{}, error) {
	path := "/v3/scopes/requests/{RequestId}"
	if params != nil && params.RequestId != nil {
		path = strings.Replace(path, "{"+"RequestId"+"}", *params.RequestId, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		ps := &DenyScopeRequest404Response{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}
