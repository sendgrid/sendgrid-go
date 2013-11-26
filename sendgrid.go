// Package sendgrid provides a simple interface to interact with the SendGrid API
// Special thanks to this gist -> https://gist.github.com/rmulley/6603544
package sendgrid

import (
	"fmt"
	"github.com/elbuo8/smtpmail"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SGClient will contain the credentials and default values
type SGClient struct {
	apiUser    string
	apiPwd     string
	apiUrl     string
	Client     *http.Client
	smtpClient smtpmail.SMTPClient
}

// NewSendGridClient will return a new SGClient.
func NewSendGridClient(apiUser, apiPwd string) SGClient {
	smtpUrl := "smtp.sendgrid.net"
	smtpPort := "587"
	apiUrl := "https://api.sendgrid.com/api/mail.send.json?"
	return SGClient{
		apiUser:    apiUser,
		apiPwd:     apiPwd,
		apiUrl:     apiUrl,
		smtpClient: smtpmail.NewSMTPClient(apiUser, apiPwd, smtpUrl, smtpPort),
	}
}

// Send will try to use the WebAPI first. If it's a success then a nill will be returned.
// Else, the SMTP API will be used as a fail over. If this happens regardless of what is the result
// of the SMTP attempt, you will receive an array of errors.
// If the length of the array is 1, then SMTP succeeded, else it also failed.
func (sg *SGClient) Send(m SGMail) []error {
	if apiError := sg.SendAPI(m); apiError != nil {
		var errors []error
		errors = append(errors, apiError)
		if smtpError := sg.SendSMTP(m); smtpError != nil {
			return append(errors, smtpError)
		} else {
			return errors
		}
	} else {
		return nil
	}
}

// SendAPI will send mail using SG web API
func (sg *SGClient) SendAPI(m SGMail) error {
	if e := m.SetHeaders(); e != nil {
		return e
	}
	values := url.Values{}
	values.Set("api_user", sg.apiUser)
	values.Set("api_key", sg.apiPwd)
	values.Set("subject", m.mail.Subject)
	values.Set("html", m.mail.HTML)
	values.Set("text", m.mail.Text)
	values.Set("from", m.mail.From)
	values.Set("headers", m.mail.Headers)
	for i := 0; i < len(m.mail.To); i++ {
		values.Set("to[]", m.mail.To[i])
	}
	for i := 0; i < len(m.mail.Bcc); i++ {
		values.Set("bcc[]", m.mail.Bcc[i])
	}
	for i := 0; i < len(m.mail.ToName); i++ {
		values.Set("toname[]", m.mail.ToName[i])
	}
	for k, v := range m.mail.Files {
		values.Set("files["+k+"]", v)
	}
	if sg.Client == nil {
		sg.Client = http.DefaultClient
	}
	r, e := sg.Client.PostForm(sg.apiUrl, values)
	defer r.Body.Close()
	if r.StatusCode == 200 && e == nil {
		return nil
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		return fmt.Errorf("sendgrid.go: code:%d error:%v body:%s", r.StatusCode, e, body)
	}
}

// SendSMTP will send mail using SG smtp API
func (sg *SGClient) SendSMTP(m SGMail) error {
	if e := m.SetHeaders(); e != nil {
		return e
	}
	return sg.smtpClient.SendSMTP(m.mail)
}
