package sendgrid

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
)

type SGClient struct {
	apiUser  string
	apiPwd   string
	apiUrl   string
	smptUrl  string
	smtpPort string
	smtpAuth smtp.Auth
}

/*
apiUser - SG username
apiPwd - SG password
*/
func NewSendGridClient(apiUser, apiPwd string) SGClient {
	smptUrl := "smtp.sendgrid.net"
	smtpPort := "587"
	apiUrl := "https://sendgrid.com/api/mail.send.json?"
	auth := smtp.PlainAuth("", apiUser, apiPwd, smptUrl)
	return SGClient{apiUser, apiPwd, apiUrl, smptUrl, smtpPort, auth}
}

func (sg *SGClient) Send(m Mail) error {
	if sg.SendAPI(m) != nil {
		return sg.SendSMTP(m)
	} else {
		return nil //sucess
	}
}

func (sg *SGClient) SendSMTP(m Mail) error {
	return smtp.SendMail(sg.smptUrl+":"+sg.smtpPort, sg.smtpAuth, m.from, m.to, []byte(m.html))
}

func (sg *SGClient) SendAPI(m Mail) error {
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
		reqUrl.WriteString(values.Encode())
		r, e := http.Get(reqUrl.String())
		defer r.Body.Close()
		if r.StatusCode == 200 && e == nil {
			return nil
		} else {
			body, _ := ioutil.ReadAll(r.Body)
			return errors.New(string(body))
		}

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
