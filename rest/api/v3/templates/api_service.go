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
	sendgrid "github.com/sendgrid/sendgrid-go/v4/client"
)

type ApiService struct {
	baseURL        string
	requestHandler *sendgrid.RequestHandler
}

func NewApiService(requestHandler *sendgrid.RequestHandler) *ApiService {
	return &ApiService{
		requestHandler: requestHandler,
		baseURL:        "https://api.sendgrid.com",
	}
}

func NewApiServiceWithClient(client sendgrid.BaseClient) *ApiService {
	return NewApiService(sendgrid.NewRequestHandler(client))
}
