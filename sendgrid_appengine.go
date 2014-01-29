// +build appengine

package sendgrid

import (
	"errors"
	"strings"
	"sync"

	"appengine"
	"appengine/datastore"
	"appengine/delay"
	"appengine/mail"
	"appengine/urlfetch"
)

var globalConfig *Config
var configSync sync.Mutex

var ErrConfig = errors.New("Unable to fetch SendGrid config")

var sendgridDelay = delay.Func("sendgrid", sendMail)

func sendMail(c appengine.Context, m *mail.Message) error {
	if appengine.IsDevAppServer() {
		c.Infof("Would have sent e-mail - %#v", m)
		return nil
	}
	err := loadConfig(c)
	if err != nil {
		return err
	}
	sgmail := Mail{
		subject: m.Subject,
		html:    m.HTMLBody,
		text:    m.Body,
		from:    m.Sender,
		replyto: m.ReplyTo,
	}
	for _, email := range m.To {
		res := strings.Split(email, "<")
		if len(res) == 2 {
			sgmail.to = append(sgmail.to, strings.Trim(res[1], "<>\" "))
			sgmail.toname = append(sgmail.toname, strings.Trim(res[0], "<>\" "))
		} else {
			sgmail.to = append(sgmail.to, email)
			sgmail.toname = append(sgmail.toname, "")
		}
	}

	for _, email := range m.Bcc {
		res := strings.Split(email, "<")
		if len(res) == 2 {
			sgmail.bcc = append(sgmail.bcc, strings.Trim(res[1], "<>\" "))
		} else {
			sgmail.bcc = append(sgmail.bcc, email)
		}
	}
	sgclient := NewSendGridClient(globalConfig.APIUser, globalConfig.APIPassword)
	sgclient.Client = urlfetch.Client(c)
	return sgclient.SendAPI(sgmail)
}

type Config struct {
	APIUser     string
	APIPassword string
}

func loadConfig(c appengine.Context) error {
	configSync.Lock()
	defer configSync.Unlock()
	if globalConfig == nil {
		key := datastore.NewKey(c, "SendGridConfig", "SendGridConfig", 0, nil)
		tempConfig := new(Config)
		err := datastore.Get(c, key, tempConfig)
		if err != nil {
			c.Errorf("Unable to fetch SendGrid config, please populate information in web console - %v", err)
			_, err = datastore.Put(c, key, &Config{
				APIUser:     "default",
				APIPassword: "default",
			})
			// put the default stub entry
			// so it can be updated in the web console
			if err != nil {
				c.Errorf("Error putting stub SendGrid config - %v", err)
			}
			return ErrConfig
		}
		if tempConfig.APIPassword == "default" || tempConfig.APIUser == "default" ||
			tempConfig.APIPassword == "" || tempConfig.APIUser == "" {
			c.Errorf("Found default SendGrid config in the datastore, please populate information in web console")
			return ErrConfig
		}
		globalConfig = tempConfig
	}
	return nil
}

func SendMailDelay(c appengine.Context, m *mail.Message) {
	sendgridDelay.Call(c, m)
}
