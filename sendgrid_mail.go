package sendgrid

import (
	"github.com/elbuo8/smtpmail"
	"github.com/sendgrid/smtpapi-go"
)

type SGMail struct {
	smtpmail.Mail
	smtpapi.SMTPAPIHeader
}

func NewMail() SGMail {
	return SGMail{}
}
