package jwt

import "fmt"

type VideoGrant struct {
	Room string `json:"room"`
}

func (videoGrant *VideoGrant) Key() string {
	return "video"
}

func (videoGrant *VideoGrant) ToPayload() map[string]interface{} {
	grant := make(map[string]interface{})

	if videoGrant.Room != "" {
		grant["room"] = videoGrant.Room
	}

	return grant
}

func (videoGrant *VideoGrant) ToString() string {
	return fmt.Sprintf("<%s %s>", "VideoGrant", videoGrant.ToPayload())
}
