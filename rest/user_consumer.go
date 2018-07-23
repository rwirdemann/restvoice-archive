package rest

import (
	"regexp"
	"net/http"
)

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


