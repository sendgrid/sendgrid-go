package sendgrid

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
)

type SGClient struct {
	apiUser  string
	apiPwd   string
	apiUrl   string
	smtpUrl  string
	smtpPort string
	smtpAuth smtp.Auth
	// Client is the HTTP transport to use when making requests.
	// It will default to http.DefaultClient if nil.
	Client *http.Client
}

/*
apiUser - SG username
apiPwd - SG password
*/
func NewSendGridClient(apiUser, apiPwd string) SGClient {
	smtpUrl := "smtp.sendgrid.net"
	smtpPort := "587"
	apiUrl := "https://sendgrid.com/api/mail.send.json?"
	smtpAuth := smtp.PlainAuth("", apiUser, apiPwd, smtpUrl)
	return SGClient{
		apiUser:  apiUser,
		apiPwd:   apiPwd,
		apiUrl:   apiUrl,
		smtpUrl:  smtpUrl,
		smtpPort: smtpPort,
		smtpAuth: smtpAuth,
	}
}

func (sg *SGClient) Send(m Mail) error {
	if sg.SendAPI(m) != nil {
		return sg.SendSMTP(m)
	} else {
		return nil //sucess
	}
}

func (sg *SGClient) SendSMTP(m Mail) error {
	return smtp.SendMail(sg.smtpUrl+":"+sg.smtpPort, sg.smtpAuth, m.from, m.to, []byte(m.html))
}

func (sg *SGClient) SendAPI(m Mail) error {
	var reqUrl bytes.Buffer
	reqUrl.WriteString(sg.apiUrl)
	values := url.Values{}
	values.Set("api_user", sg.apiUser)
	values.Set("api_key", sg.apiPwd)
	values.Set("subject", m.subject)
	values.Set("html", m.html)
	values.Set("text", m.text)
	values.Set("from", m.from)
	for i := 0; i < len(m.to); i++ {
		values.Set("to[]", m.to[i])
	}
	for i := 0; i < len(m.bcc); i++ {
		values.Set("bcc[]", m.bcc[i])
	}
	for i := 0; i < len(m.toname); i++ {
		values.Set("toname[]", m.toname[i])
	}
	reqUrl.WriteString(values.Encode())
	if sg.Client == nil {
		sg.Client = http.DefaultClient
	}
	r, e := sg.Client.Get(reqUrl.String())
	defer r.Body.Close()
	if r.StatusCode == 200 && e == nil {
		return nil
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		return fmt.Errorf("sendgrid.go: code:%d error:%v body:%s", r.StatusCode, e, body)
	}
}

type Mail struct {
	to       []string
	toname   []string
	subject  string
	html     string
	text     string
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

func (m *Mail) AddText(text string) {
	m.text = text
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
