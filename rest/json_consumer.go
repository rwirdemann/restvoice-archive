package rest

import (
	"encoding/json"
)

type JSONConsumer struct {
	result interface{}
}

func NewJSONConsumer(result interface{}) JSONConsumer {
	return JSONConsumer{result: result}
}

func (c JSONConsumer) Consume(body interface{}) interface{} {
	json.Unmarshal(body.([]byte), c.result)
	return c.result
}
