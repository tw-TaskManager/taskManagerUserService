package tokenGenerator

import (
	"github.com/dgrijalva/jwt-go"
	"encoding/base64"
)

const secretKey string = "TM2016"

func Generate(id string) string {
	token := jwt.New(jwt.SigningMethodES256)
	token.Claims["id"] = id;
	decodedId, _ := base64.URLEncoding.DecodeString(secretKey)
	signedToken, _ := token.SignedString(decodedId)
	return signedToken
}
