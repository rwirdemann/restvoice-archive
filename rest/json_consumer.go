package rest

import (
	"encoding/json"
	"log"
)

type JSONConsumer struct {
	result interface{}
}

func NewJSONConsumer(result interface{}) JSONConsumer {
	return JSONConsumer{result: result}
}

func (c JSONConsumer) Consume(body interface{}) interface{} {
	if err := json.Unmarshal(body.([]byte), c.result); err != nil {
		log.Println(err)
	}
	return c.result
}
