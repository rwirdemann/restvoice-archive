package rest

import (
	"net/http"
)

type QueryVariableConsumer struct {
	queryVarName string
}

func NewQueryVariableConsumer(queryVarName string) QueryVariableConsumer {
	return QueryVariableConsumer{queryVarName: queryVarName}
}

func (c QueryVariableConsumer) Consume(request interface{}) interface{} {
	r := request.(*http.Request)
	q := r.URL.Query()
	if v, ok := q[c.queryVarName]; ok {
		return v[0]
	}
	return nil
}
