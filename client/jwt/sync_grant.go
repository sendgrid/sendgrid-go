package jwt

import "fmt"

type SyncGrant struct {
	ServiceSid string `json:"service_sid"`
	EndpointID string `json:"endpoint_id"`
}

func (syncGrant *SyncGrant) Key() string {
	return "data_sync"
}

func (syncGrant *SyncGrant) ToPayload() map[string]interface{} {
	grant := make(map[string]interface{})
	if syncGrant.ServiceSid != "" {
		grant["service_sid"] = syncGrant.ServiceSid
	}
	if syncGrant.EndpointID != "" {
		grant["endpoint_id"] = syncGrant.EndpointID
	}

	return grant
}

func (syncGrant *SyncGrant) ToString() string {
	return fmt.Sprintf("<%s %s>", "SyncGrant", syncGrant.ToPayload())
}
