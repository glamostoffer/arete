package email

type Config struct {
	SMTPHost string `validate:"required"`
	SMTPPort int    `validate:"required"`
	Login    string `validate:"required"`
	Password string `validate:"required"`
}
