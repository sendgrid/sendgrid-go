package sendgrid

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
	"path/filepath"
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

/*
Send will try to use the WebAPI first. If it's a success then a nill will be returned.
Else, the SMTP API will be used as a fail over. If this happens regardless of what is the result
of the SMTP attempt, you will receive an array of errors.
If the length of the array is 1, then SMTP succeeded, else it also failed.
*/
func (sg *SGClient) Send(m Mail) []error {
	if apiError := sg.SendAPI(m); apiError != nil {
		var errors []error
		errors = append(errors, apiError)
		if smtpError := sg.SendSMTP(m); smtpError != nil {
			return append(errors, smtpError)
		} else {
			return errors
		}
	} else {
		return nil //sucess
	}
}

/*
SMTP interface. Still being developed. Use API instead.
*/
func (sg *SGClient) SendSMTP(m Mail) error {
	return smtp.SendMail(sg.smtpUrl+":"+sg.smtpPort, sg.smtpAuth, m.from, m.to, []byte(m.html))
}

/*
Sends email using SG web API
*/
func (sg *SGClient) SendAPI(m Mail) error {
	values := url.Values{}
	values.Set("api_user", sg.apiUser)
	values.Set("api_key", sg.apiPwd)
	values.Set("subject", m.subject)
	values.Set("html", m.html)
	values.Set("text", m.text)
	values.Set("from", m.from)
	headers, e := json.Marshal(m.headers)
	if e != nil {
		return fmt.Errorf("sendgrid.go: Error parsing JSON headers")
	}
	values.Set("headers", string(headers[:]))
	for i := 0; i < len(m.to); i++ {
		values.Set("to[]", m.to[i])
	}
	for i := 0; i < len(m.bcc); i++ {
		values.Set("bcc[]", m.bcc[i])
	}
	for i := 0; i < len(m.toname); i++ {
		values.Set("toname[]", m.toname[i])
	}
	for k, v := range m.files {
		values.Set("files["+k+"]", v)
	}
	if sg.Client == nil {
		sg.Client = http.DefaultClient
	}
	fmt.Print(values)
	r, e := sg.Client.PostForm(sg.apiUrl, values)
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
	files    map[string]string
	headers  map[string]string
}

func NewMail() Mail {
	return Mail{}
}

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

func (m *Mail) AddHeader(header, value string) {
	if m.headers == nil {
		m.headers = make(map[string]string)
	}
	m.headers[header] = value
}

func (m *Mail) AddAttachment(filePath string) error {
	if m.files == nil {
		m.files = make(map[string]string)
	}
	buf, e := ioutil.ReadFile(filePath)
	if e != nil {
		return e
	}
	_, filename := filepath.Split(filePath)
	m.files[filename] = base64.StdEncoding.EncodeToString(buf)
	return nil
}
