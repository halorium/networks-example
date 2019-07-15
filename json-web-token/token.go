package jsonwebtoken

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/halorium/networks-example/flaw"
)

var InJwtSigningKey []byte

func init() {
	if os.Getenv("IN_JWT_SIGNING_KEY") == "" {
		panic(flaw.New("IN_JWT_SIGNING_KEY is required").Wrap("cannot init"))
	}

	InJwtSigningKey = []byte(os.Getenv("IN_JWT_SIGNING_KEY"))
}

type Token struct {
	*jwt.Token
	SignedToken string `json:"signed-token"`
}

func New(claims *Claims) (*Token, *flaw.Error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(InJwtSigningKey)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot New")
	}

	return &Token{
		token,
		signedToken,
	}, nil
}
