package jsonwebtoken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/halorium/networks-example/stamps"
)

type Identity struct {
	IdentityID string   `json:"identity-id"`
	Actions    []string `json:"actions"`
}

type Claims struct {
	jwt.StandardClaims
	// Audience  string `json:"aud,omitempty"`
	// ExpiresAt int64  `json:"exp,omitempty"`
	// Id        string `json:"jti,omitempty"`
	// IssuedAt  int64  `json:"iat,omitempty"`
	// Issuer    string `json:"iss,omitempty"`
	// NotBefore int64  `json:"nbf,omitempty"`
	// Subject   string `json:"sub,omitempty"`
	DelegatorIdentity Identity `json:"delegator-identity,omitempty"`
	DelegatedIdentity Identity `json:"delegated-identity,omitempty"`
}

func NewClaims() *Claims {
	claims := &Claims{}
	now := stamps.New()
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = now.Add(time.Hour).Unix()

	return claims
}
