package jwt

import "github.com/golang-jwt/jwt"

type JWT string

func NewJWT(claims jwt.Claims, secret string) (JWT, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return JWT(tokenString), nil
}

func (j JWT) String() string {
	return string(j)
}
