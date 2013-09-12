package sendgridgo

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type SGClient struct {
	apiUser string
	apiPwd  string
	apiUrl  string
}

func NewSendGridClient(apiUser, apiPwd string) SGClient {
	return SGClient{apiUser, apiPwd, "https://sendgrid.com/api/mail.send.json?"}
}

func (sg *SGClient) Send(m Mail) error {
	if total := len(m.to); total == len(m.toname) || total == 1 {
		var reqUrl bytes.Buffer
		reqUrl.WriteString(sg.apiUrl)
		values := url.Values{}
		values.Set("api_user", sg.apiUser)
		values.Set("api_key", sg.apiPwd)
		values.Set("subject", m.subject)
		values.Set("html", m.html)
		values.Set("from", m.from)
		switch {
		case total == 1:
			values.Set("to", m.to[0])
			if len(m.toname) == 1 {
				values.Set("toname", m.toname[0])
			}
		case total > 1:
			for i := 0; i < total; i++ {
				values.Set("to[]", m.to[i])
				values.Set("toname[]", m.toname[i])
			}
		}
		fmt.Println(reqUrl.String(), values)
		reqUrl.WriteString(values.Encode())
		fmt.Println(reqUrl)
		r, e := http.Get(reqUrl.String())
		fmt.Println(r, e)
		return nil
	}
	return errors.New(`The total names of receipients must be equal
	to the email addresses. Unless there is only one receipient.`)
}

type Mail struct {
	to       []string
	toname   []string
	subject  string
	html     string
	from     string
	bcc      []string
	fromname string
	replyto  string
	date     string
	//still missing some stuff
}

func NewMail() Mail {
	return Mail{}
}

/*
TODO: Validate email addressed with RegExp.
*/
func (m *Mail) AddTo(email string) {
	m.to = append(m.to, email)
}

func (m *Mail) AddToName(name string) {
	m.toname = append(m.toname, name)
}

func (m *Mail) AddSubject(s string) {
	m.subject = s
}

func (m *Mail) AddHTML(html string) {
	m.html = html
}

func (m *Mail) AddFrom(from string) {
	m.from = from
}

func (m *Mail) AddBCC(email string) {
	m.bcc = append(m.bcc, email)
}

func (m *Mail) AddFromName(name string) {
	m.fromname = name
}

func (m *Mail) AddReplyTo(reply string) {
	m.replyto = reply
}

func (m *Mail) AddDate(date string) {
	m.date = date
}
