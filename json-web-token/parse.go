package jsonwebtoken

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/halorium/networks-example/flaw"
)

func Parse(signedString string) (*Token, *flaw.Error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		signedString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, flaw.New("invalid signing method").Wrap("cannot Parse")
			}

			return InJwtSigningKey, nil
		},
	)

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot Parse")
	}

	err = claims.Valid()

	if err != nil {
		return nil, flaw.From(err).Wrap("cannot Parse")
	}

	if !token.Valid {
		return nil, flaw.New("invalid token").Wrap("cannot Parse")
	}

	return New(claims)
}
