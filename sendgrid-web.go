package sendgridgo

import (
	//"errors"
	"net/http"
	//"net/url"
	//"encoding/json"
	"fmt"
	"net/mail"
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
	url := fmt.Sprintf("%sapi_user=%s&api_key=%s&to=%s&subject=%s&html=%s&from=%s", sg.apiUrl, sg.apiUser, sg.apiPwd, m.to, m.subject, m.html, m.from)
	fmt.Println(url)
	r, e := http.Get(url)
	fmt.Println(r, e)
	return nil
}

type Mail struct {
	to       string
	toname   string
	subject  string
	html     string
	from     string
	bcc      string
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
	m.to = email
}

func (m *Mail) AddToName(name string) {
	m.toname = name
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
	m.bcc = email
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
