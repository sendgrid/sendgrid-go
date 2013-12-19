// Package sendgrid provides a simple interface to interact with the SendGrid API
// Special thanks to this gist -> https://gist.github.com/rmulley/6603544
package sendgrid

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SGClient will contain the credentials and default values
type SGClient struct {
	apiUser    string
	apiPwd     string
	apiMail     string
	Client     *http.Client
}

// NewSendGridClient will return a new SGClient.
func NewSendGridClient(apiUser, apiPwd string) SGClient {
	apiMail := "https://api.sendgrid.com/api/mail.send.json?"
	return SGClient{
		apiUser:    apiUser,
		apiPwd:     apiPwd,
		apiMail:     apiMail,
	}
}

// SendAPI will send mail using SG web API
func (sg *SGClient) Send(m SGMail) error {
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
	r, e := sg.Client.PostForm(sg.apiMail, values)
	defer r.Body.Close()
	if r.StatusCode == 200 && e == nil {
		return nil
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		return fmt.Errorf("sendgrid.go: code:%d error:%v body:%s", r.StatusCode, e, body)
	}
}

