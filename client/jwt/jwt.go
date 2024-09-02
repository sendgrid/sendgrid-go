package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	. "github.com/twilio/twilio-go/client/jwt/util"
)

type TokenType string

type Jwt struct {
	// The secret used to encode the JWT token.
	SecretKey string
	// The issuer of the JWT token.
	Issuer string
	// The subject of this JWT, omitted from the payload by default.
	Subject string
	// The algorithm used to encode the JWT token, defaults to 'HS256'.
	Algorithm string
	// Time in seconds before the JWT token is invalid. Defaults to now.
	Nbf float64
	// Time to live of the JWT in seconds; defaults to 1 hour.
	Ttl float64
	// Time in seconds since epoch this JWT is valid for. Override ttl if provided.
	ValidUntil float64

	DecodedHeaders map[string]interface{}
	DecodedPayload map[string]interface{}
}

func (token *Jwt) generatePayload(payload map[string]interface{}) map[string]interface{} {
	now := time.Now().Unix()

	payload["iss"] = token.Issuer
	payload["exp"] = float64(now) + token.Ttl

	if token.Nbf != 0 {
		payload["nbf"] = token.Nbf
	} else {
		payload["nbf"] = float64(now)
	}

	if token.ValidUntil != 0 {
		payload["exp"] = token.ValidUntil
	}
	if token.Subject != "" {
		payload["sub"] = token.Subject
	}

	return payload
}

func (token *Jwt) FromJwt(jwtStr string, key string) (*Jwt, error) {
	verifyToken := true
	if key == "" {
		verifyToken = false
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	decodedToken, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(key), nil
	})

	if decodedToken != nil {
		if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok {
			if verifyToken && !decodedToken.Valid {
				return nil, errors.New("token could not be validated")
			}
			jwtToken := Jwt{
				SecretKey: key,
				Issuer:    claims["iss"].(string),
				Algorithm: decodedToken.Header["alg"].(string),
			}

			if val, ok := claims["sub"].(string); ok {
				jwtToken.Subject = val
			}

			if val, ok := claims["exp"].(float64); ok {
				jwtToken.ValidUntil = val
			}

			nbf, err := strconv.ParseFloat(fmt.Sprintf("%v", claims["nbf"]), 64)
			if err == nil {
				jwtToken.Nbf = nbf
			}

			jwtToken.DecodedHeaders = decodedToken.Header
			jwtToken.DecodedPayload = claims
			return &jwtToken, nil
		} else {
			return nil, err
		}
	}

	return nil, errors.New("error decoding JWT token")
}

func (token *Jwt) Payload() map[string]interface{} {
	if token.DecodedPayload == nil {
		token.DecodedPayload = token.generatePayload(map[string]interface{}{})
	}

	return token.DecodedPayload
}

func (token *Jwt) Headers() map[string]interface{} {
	if token.DecodedHeaders == nil {
		token.DecodedHeaders = token.generateHeaders()
	}

	return token.DecodedHeaders
}

func (token *Jwt) generateHeaders() map[string]interface{} {
	headers := make(map[string]interface{})
	headers["alg"] = HS256
	headers["typ"] = JWT
	return headers
}

// Encode this JWT struct into a string.
// algorithm - algorithm used to encode the JWT that overrides the default
// ttl - specify ttl to override the default
func (token *Jwt) ToJwt(generateHeaders, generatePayload func() map[string]interface{}) (string, error) {
	if token.SecretKey == "" {
		panic("JWT does not have a signing key configured.")
	}

	headers := generateHeaders()
	payload := generatePayload()
	if signedToken, err := SignTokenWithHMAC(headers, payload, token.SecretKey); err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}

func SignTokenWithHMAC(headers, payload map[string]interface{}, secret string) (string, error) {
	claims := jwt.MapClaims{}

	for k, v := range payload {
		claims[k] = v
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	for hk, hv := range headers {
		jwtToken.Header[hk] = hv
	}

	if tokenString, err := jwtToken.SignedString([]byte(secret)); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}
