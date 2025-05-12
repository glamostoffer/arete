package consumer

type Config struct {
	Brokers []string `validate:"required"`
	Topic   string   `validate:"required"`
}
