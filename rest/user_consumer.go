package rest

import (
	"regexp"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"crypto/rsa"
	"io/ioutil"
	"log"
)

const publicKeyFilePath = "keycloak_key.pub"

var publicKey *rsa.PublicKey

func init() {
	if f, err := ioutil.ReadFile(publicKeyFilePath); err == nil {
		if publicKey, err = jwt.ParseRSAPublicKeyFromPEM(f); err != nil {
			log.Fatalf("Could not parse public key from pem file")
		}
	} else {
		log.Fatalf("Could not open public key file: %s", publicKeyFilePath)
	}
}

type UserConsumer struct {
}

func NewUserConsumer() *UserConsumer {
	return &UserConsumer{}
}

func (UserConsumer) Consume(i interface{}) interface{} {
	return "1234"
}

func extractJwtFromHeader(header http.Header) (jwt string) {
	var jwtRegex = regexp.MustCompile(`^Bearer (\S+)$`)

	if val, ok := header["Authorization"]; ok {
		for _, value := range val {
			if result := jwtRegex.FindStringSubmatch(value); result != nil {
				jwt = result[1]
				return
			}
		}
	}

	return
}

func claim(token string, key string) string {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err == nil {
		if claims, ok := t.Claims.(jwt.MapClaims); ok {
			if claims[key] != nil {
				return claims[key].(string)
			}
		}
	}

	return ""
}
