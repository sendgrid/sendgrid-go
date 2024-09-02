package jwt

import "fmt"

type ChatGrant struct {
	ServiceSid        string `json:"service_sid"`
	EndpointID        string `json:"endpoint_id"`
	DeploymentRoleSid string `json:"deployment_role_sid"`
	PushCredentialSid string `json:"push_credential_sid"`
}

func (chatGrant *ChatGrant) Key() string {
	return "chat"
}

func (chatGrant *ChatGrant) ToPayload() map[string]interface{} {
	grant := make(map[string]interface{})
	if chatGrant.ServiceSid != "" {
		grant["service_sid"] = chatGrant.ServiceSid
	}
	if chatGrant.EndpointID != "" {
		grant["endpoint_id"] = chatGrant.EndpointID
	}
	if chatGrant.DeploymentRoleSid != "" {
		grant["deployment_role_sid"] = chatGrant.DeploymentRoleSid
	}
	if chatGrant.PushCredentialSid != "" {
		grant["push_credential_sid"] = chatGrant.PushCredentialSid
	}

	return grant
}

func (chatGrant *ChatGrant) ToString() string {
	return fmt.Sprintf("<%s %s>", "ChatGrant", chatGrant.ToPayload())
}
