package jwt

type BaseGrant interface {
	Key() string
	ToPayload() map[string]interface{}
	ToString() string
}
