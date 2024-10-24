/*
* This code was generated by
*
* SENDGRID-OAI-GENERATOR
*
* Twilio SendGrid Marketing Campaigns Designs
* The Twilio SendGrid Designs API offers the ability to manage assets stored in the Twilio SendGrid [Design Library](https://mc.sendgrid.com/design-library/my-designs).  The Design Library is a feature-rich email layout tool and media repository. You can [build designs for all your marketing email needs](https://sendgrid.com/docs/ui/sending-email/working-with-marketing-campaigns-email-designs/), including Single Sends and Automations.  You can also duplicate and then modify one of the pre-built designs provided by Twilio SendGrid to get you started.  The Designs API provides a REST-like interface for creating new designs, retrieving a list of existing designs, duplicating or updating a design, and deleting a design.
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

type UpdateDesignParam struct {
	// The ID of the Design you want to duplicate.
	Id *string `json:"id"`
	//
	UpdateDesignRequest *UpdateDesignRequest `json:"UpdateDesignRequest,omitempty"`
}

func (params *UpdateDesignParam) SetId(Id string) *UpdateDesignParam {
	params.Id = &Id
	return params
}
func (params *UpdateDesignParam) SetUpdateDesignRequest(UpdateDesignRequest UpdateDesignRequest) *UpdateDesignParam {
	params.UpdateDesignRequest = &UpdateDesignRequest
	return params
}

// **This endpoint allows you to edit a design**.  The Design API supports PATCH requests, which allow you to make partial updates to a single design. Passing data to a specific field will update only the data stored in that field; all other fields will be unaltered.  For example, updating a design's name requires that you make a PATCH request to this endpoint with data specified for the `name` field only.  ``` {     \"name\": \"<Updated Name>\" } ```
func (c *ApiService) UpdateDesign(params *UpdateDesignParam) (interface{}, error) {
	path := "/v3/designs/{Id}"
	if params != nil && params.Id != nil {
		path = strings.Replace(path, "{"+"Id"+"}", *params.Id, -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.UpdateDesignRequest != nil {
		b, err := json.Marshal(*params.UpdateDesignRequest)
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
		ps := &DesignOutput{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 400 {
		ps := &ApiErrors{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	if resp.StatusCode == 404 {
		ps := &ApiErrors{}
		if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
			return nil, err
		}

		return ps, err
	}
	return http.Response{StatusCode: resp.StatusCode, Body: resp.Body, Header: resp.Header}, nil
}