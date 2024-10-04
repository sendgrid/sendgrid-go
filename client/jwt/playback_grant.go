package jwt

import "fmt"

type PlaybackGrant map[string]interface{}

func (playbackGrant *PlaybackGrant) Key() string {
	return "player"
}

func (playbackGrant *PlaybackGrant) ToPayload() map[string]interface{} {
	return *playbackGrant
}

func (playbackGrant *PlaybackGrant) ToString() string {
	return fmt.Sprintf("<%s %s>", "PlaybackGrant", playbackGrant.ToPayload())
}
