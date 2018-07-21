package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PathVariableConsumer struct {
	pathVarName string
}

func NewPathVariableConsumer(pathVarName string) PathVariableConsumer {
	return PathVariableConsumer{pathVarName: pathVarName}
}

func (c PathVariableConsumer) Consume(request interface{}) interface{} {
	vars := mux.Vars(request.(*http.Request))
	s := vars[c.pathVarName]
	i, _ := strconv.Atoi(s)
	return i
}
