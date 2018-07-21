package rest

import (
	"net/http"
)

type QueryVariableConsumer struct {
	queryVarName string
}

func NewQueryVariableConsumer(pathVarName string) PathVariableConsumer {
	return PathVariableConsumer{pathVarName: pathVarName}
}

func (c QueryVariableConsumer) Consume(request interface{}) interface{} {
	r := request.(*http.Request)
	return r.URL.Query()[c.queryVarName]
}
