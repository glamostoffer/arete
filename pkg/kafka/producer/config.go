package producer

type Config struct {
	Address string `validate:"required"`
	Topic   string `validate:"required"`
}
