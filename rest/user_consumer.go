package rest

type UserConsumer struct {
}

func NewUserConsumer() *UserConsumer {
	return &UserConsumer{}
}

func (UserConsumer) Consume(i interface{}) interface{} {
	return "1234"
}
