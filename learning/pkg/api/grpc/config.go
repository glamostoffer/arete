package grpc

type Config struct {
	Address string `validate:"required"`
}
