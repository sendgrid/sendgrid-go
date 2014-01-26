package sendgrid

import (
	"github.com/elbuo8/smtpmail"
	"github.com/sendgrid/smtpapi-go"
	"io/ioutil"
	"net/mail"
	"path/filepath"
)

type SGMail struct {
	smtpmail.Mail
	smtpapi.SMTPAPIHeader
}

func NewMail() SGMail {
	return SGMail{}
}

func (m *SGMail) AddAttachment(filePath string) error {
	bytes, e := ioutil.ReadFile(filePath)
	if e != nil {
		return e
	}
	_, filename := filepath.Split(filePath)
	m.AddAttachmentStream(filename, bytes)
	return nil
}

func (m *SGMail) AddAttachmentStream(filename string, stream []byte) {
	if m.Files == nil {
		m.Files = make(map[string]string)
	}
	m.Files[filename] = string(stream)
}

func (m *SGMail) AddTo(email string) {
	m.Mail.AddTo(email)
}

func (m *SGMail) AddRecipient(email *mail.Address) {
	m.Mail.AddRecipient(email)
}
