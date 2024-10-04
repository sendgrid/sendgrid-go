package jwt

import "fmt"

type IpMessagingGrant struct {
	ServiceSid        string `json:"service_sid"`
	EndpointID        string `json:"endpoint_id"`
	DeploymentRoleSid string `json:"deployment_role_sid"`
	PushCredentialSid string `json:"push_credential_sid"`
}

func (ipMessagingGrant *IpMessagingGrant) Key() string {
	return "ip_messaging"
}

func (ipMessagingGrant *IpMessagingGrant) ToPayload() map[string]interface{} {
	grant := make(map[string]interface{})
	if ipMessagingGrant.ServiceSid != "" {
		grant["service_sid"] = ipMessagingGrant.ServiceSid
	}
	if ipMessagingGrant.EndpointID != "" {
		grant["endpoint_id"] = ipMessagingGrant.EndpointID
	}
	if ipMessagingGrant.DeploymentRoleSid != "" {
		grant["deployment_role_sid"] = ipMessagingGrant.DeploymentRoleSid
	}
	if ipMessagingGrant.PushCredentialSid != "" {
		grant["push_credential_sid"] = ipMessagingGrant.PushCredentialSid
	}

	return grant
}

func (ipMessagingGrant *IpMessagingGrant) ToString() string {
	return fmt.Sprintf("<%s %s>", "IpMessagingGrant", ipMessagingGrant.ToPayload())
}
