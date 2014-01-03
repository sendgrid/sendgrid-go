package sendgrid

import (
	"github.com/elbuo8/smtpmail"
	"github.com/sendgrid/smtpapi-go"
	"io/ioutil"
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
	if m.Files == nil {
		m.Files = make(map[string]string)
	}
	file, e := ioutil.ReadFile(filePath)
	if e != nil {
		return e
	}
	_, filename := filepath.Split(filePath)
	m.Files[filename] = string(file)
	return nil
}

