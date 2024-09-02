package taskrouter

type Policy struct {
	Url         string                 `json:"url,omitempty"`
	Method      string                 `json:"method,omitempty"`
	Allow       bool                   `json:"allow,omitempty"`
	PostFilter  map[string]interface{} `json:"post_filter"`
	QueryFilter map[string]interface{} `json:"query_filter"`
}

func GeneratePolicy(url string, method string, allow bool, postFilters map[string]interface{}, queryFilters map[string]interface{}) Policy {
	return Policy{
		Url:         url,
		Method:      method,
		Allow:       allow,
		PostFilter:  postFilters,
		QueryFilter: queryFilters,
	}
}

func (policy *Policy) Payload() map[string]interface{} {
	payload := map[string]interface{}{
		"url":          policy.Url,
		"method":       policy.Method,
		"allow":        policy.Allow,
		"post_filter":  map[string]interface{}{},
		"query_filter": map[string]interface{}{},
	}
	if policy.QueryFilter != nil {
		payload["query_filter"] = policy.QueryFilter
	}
	if policy.PostFilter != nil {
		payload["post_filter"] = policy.PostFilter
	}
	return payload
}
