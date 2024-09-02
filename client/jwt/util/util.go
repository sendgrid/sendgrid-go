package util

const (
	HS256 = "HS256"
	JWT   = "JWT"
	CType = "twilio-fpa;v=1"
)

func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
