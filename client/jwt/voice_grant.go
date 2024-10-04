package jwt

import "fmt"

type VoiceGrant struct {
	Incoming          Incoming `json:"incoming"`
	Outgoing          Outgoing `json:"outgoing"`
	PushCredentialSid string   `json:"push_credential_sid"`
	EndpointID        string   `json:"endpoint_id"`
}

type Incoming struct {
	Allow bool `json:"allow"`
}

type Outgoing struct {
	ApplicationSid    string                 `json:"application_sid"`
	ApplicationParams map[string]interface{} `json:"params"`
}

func (voiceGrant *VoiceGrant) Key() string {
	return "voice"
}

func (voiceGrant *VoiceGrant) ToPayload() map[string]interface{} {
	grant := make(map[string]interface{})
	if voiceGrant.Incoming.Allow {
		grant["incoming"] = map[string]interface{}{
			"allow": true,
		}
	}
	if voiceGrant.Outgoing.ApplicationSid != "" {
		grant["outgoing"] = map[string]interface{}{
			"application_sid": voiceGrant.Outgoing.ApplicationSid,
		}
		if len(voiceGrant.Outgoing.ApplicationParams) != 0 {
			grant["outgoing"].(map[string]interface{})["params"] = voiceGrant.Outgoing.ApplicationParams
		}
	}

	if voiceGrant.PushCredentialSid != "" {
		grant["push_credential_sid"] = voiceGrant.PushCredentialSid
	}
	if voiceGrant.EndpointID != "" {
		grant["endpoint_id"] = voiceGrant.EndpointID
	}

	return grant
}

func (voiceGrant *VoiceGrant) ToString() string {
	return fmt.Sprintf("<%s %s>", "VoiceGrant", voiceGrant.ToPayload())
}
