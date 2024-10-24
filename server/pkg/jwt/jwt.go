package jwt

import "github.com/golang-jwt/jwt/v4"

type Jwt struct {
	signingKey []byte
}

func NewJwt(signingKey string) *Jwt {
	return &Jwt{signingKey: []byte(signingKey)}
}

func (j *Jwt) CreateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signingKey)
}
