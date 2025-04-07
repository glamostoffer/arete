package dto

type SendEmailRequest struct {
	Subject     string
	Recipient   string
	ContentType string
	Body        string
}
