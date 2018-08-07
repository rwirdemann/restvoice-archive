package rest

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"crypto/rsa"
	"io/ioutil"
	"log"
	"regexp"
	"os"
)

const publicKeyFilePath = "keycloak_key.pub"

var publicKey *rsa.PublicKey

var kf []byte

func init() {
	var err error
	if kf, err = ioutil.ReadFile(publicKeyFilePath); err == nil {
		if publicKey, err = jwt.ParseRSAPublicKeyFromPEM(kf); err != nil {
			log.Fatalf("Could not parse public key from pem file")
		}
	} else {
		log.Fatalf("Could not open public key file: %s", publicKeyFilePath)
	}
}

func BasicAuth(r *http.Request) bool {
	if os.Getenv("AUTHENTICATION") != "basic" {
		return true
	}

	if username, password, ok := r.BasicAuth(); ok {
		if username == "go" && password == "time" {
			return true
		}
	}
	return false
}

func JwtAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Skip if jwt authentication is disabled
		if os.Getenv("AUTHENTICATION") != "jwt" {
			next.ServeHTTP(w, r)
			return
		}

		token := extractJwtFromHeader(r.Header)
		if verifyJWT(token) {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("WWW-Authenticate", "Bearer realm=\"restvoice.org\"")
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func verifyJWT(token string) bool {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	return err == nil && t.Valid
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

