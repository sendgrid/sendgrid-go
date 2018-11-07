# Tutorial Heroku Deployments

## Require

* [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli)
* [godep](https://github.com/tools/godep)

## Initial project with example code

```go
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var sendgridkey = os.Getenv("SENDGRID_API_KEY")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/sendmail/:to", sendmail)

	router.Run(fmt.Sprintf(":%d", port))
}

func sendmail(c *gin.Context) {
	to := c.Param("to")
	res, err := sendmailWithSendGrid(to)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		c.JSON(res.StatusCode, res.Body)
	}

}

func sendmailWithSendGrid(sendTo string) (*rest.Response, error) {
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", sendTo)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sendgridkey)
	return client.Send(message)
}

```

### Save dependencies to your project with godep

```bash
go get github.com/gin-gonic/gin
go get github.com/sendgrid/sendgrid-go
go get github.com/sendgrid/rest
godep save ./...
```

### Define a Procfile

follow document [link.](https://devcenter.heroku.com/articles/getting-started-with-nodejs#define-a-procfile)

The `Procfile` in the example app you deployed looks like this:

```sh
web: sendgrid-go-heroku
```

create app in heroku with heroku-cli run `heroku apps:create sendgrid-go-heroku`
and deploy with simple git command like this:

```sh
git add .
git commit -m 'deploy to heroku'
git push heroku master
```

## If results look like this

```sh
remote: Compressing source files... done.
remote: Building source:
remote:
remote: -----> Go app detected
remote: -----> Checking Godeps/Godeps.json file.
remote: -----> Using go1.8.4
remote:  !!    Installing package '.' (default)
remote:  !!
remote: -----> Running: go install -v -tags heroku .
remote: sendgrid-go-heroku/vendor/github.com/gin-contrib/sse
remote: sendgrid-go-heroku/vendor/github.com/gin-gonic/gin/json
remote: sendgrid-go-heroku/vendor/github.com/golang/protobuf/proto
remote: sendgrid-go-heroku/vendor/github.com/ugorji/go/codec
remote: sendgrid-go-heroku/vendor/gopkg.in/go-playground/validator.v8
remote: sendgrid-go-heroku/vendor/gopkg.in/yaml.v2
remote: sendgrid-go-heroku/vendor/github.com/mattn/go-isatty
remote: sendgrid-go-heroku/vendor/github.com/sendgrid/rest
remote: sendgrid-go-heroku/vendor/github.com/sendgrid/sendgrid-go/helpers/mail
remote: sendgrid-go-heroku/vendor/github.com/sendgrid/sendgrid-go
remote: sendgrid-go-heroku/vendor/github.com/gin-gonic/gin/binding
remote: sendgrid-go-heroku/vendor/github.com/gin-gonic/gin/render
remote: sendgrid-go-heroku/vendor/github.com/gin-gonic/gin
remote: sendgrid-go-heroku
remote: -----> Discovering process types
remote:        Procfile declares types -> web
remote:
remote: -----> Compressing...
remote:        Done: 5.1M
remote: -----> Launching...
remote:        Released v4
remote:        https://sendgrid-go-heroku.herokuapp.com/ deployed to Heroku
remote:
remote: Verifying deploy... done.
```