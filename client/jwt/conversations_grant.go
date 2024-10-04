package jwt

import "fmt"

type ConversationsGrant struct {
	ConfigurationProfileSid string `json:"configuration_profile_sid"`
}

func (conversationsGrant *ConversationsGrant) Key() string {
	return "rtc"
}

func (conversationsGrant *ConversationsGrant) ToPayload() map[string]interface{} {
	grant := make(map[string]interface{})

	if conversationsGrant.ConfigurationProfileSid != "" {
		grant["configuration_profile_sid"] = conversationsGrant.ConfigurationProfileSid
	}

	return grant
}

func (conversationsGrant *ConversationsGrant) ToString() string {
	return fmt.Sprintf("<%s %s>", "ConversationsGrant", conversationsGrant.ToPayload())
}
