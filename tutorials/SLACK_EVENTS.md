<a name="slack-events"></a>
# Integrations - Slack Events

## Table of Contents

* [Credential Storage](#credentials)
* [Slack Setup](#setup)
* [Setting Up a Webhook](#webhook)
* [Event Triggered Mail](#triggered_mail)
* [Getting Starred Items](#starred)
* [Conclusion](#conclusion)

This tutorial will go over how to setup a simple integration with Slack that allows an email to be triggered based on some defined event. You'll need the following to complete this tutorial.

* SendGrid API Key
* Slack OAuth Token
* Slack account & Valid Slack Workspace
* ngrok (if you do not have an publicly accessible web endpoint)

The following example will explore the basic setup of an endpoint for the [Slack Events API](https://api.slack.com/events-api) that will listen for a specific [event type](https://api.slack.com/events), and then react by using SendGrid's mail send to send a customized message.

Slack is a communal space, where knowledge is shared and conversations are had at sometimes overwhelming levels. Slack offers a few ways of keeping track of things, including starring or pinning messages or files. We can use SendGrid and transactional email to help us maintain alternate copies of information we need from Slack.

We will be listening for a [message event](https://api.slack.com/events/message) with the text `sgstarred|to@example.com`, which we'll treat as a command to trigger an email containing the summary of the last 5 starred items on one's Slack workspace in order to help us track things. The [stars.list](https://api.slack.com/methods/stars.list) endpoint will be used to get these starred items.

<a name="credentials"></a>
### Credential Storage

It's not a good practice to directly insert API keys (or any authorization information) into your script for security reasons. As an alternative, please save your credentials in their respective `.env` files, with the following command:
```
echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > sendgrid.env
echo "export SLACK_OAUTH_TOKEN='YOUR_OAUTH_TOKEN'" > slack.env
source ./sendgrid.env && source ./slack.env
```

<a name="setup"></a>
### Slack Setup

Once you have a functional Slack workspace running, you'll need to:

1. Create a Slack [app](https://api.slack.com/apps), and then click on it for further configuration.

2. On the sidebar under Features select *Event Subscriptions*.

3. Toggle **On** for Enable Events. We will insert a Request URL later.

4. Under Subscribe to Workspace Events, click **Add Workspace Event** and add the [message.channels](https://api.slack.com/events/message.channels) event type.

Once you've setup your new app the following the steps above, you can now install it to any given workspace. Installing the app will also generate an OAuth token you can copy & paste from the UI, into an `.env` file, which we'll be using for an API call.

<a name="setup"></a>
### Setting Up a Webhook

Now that you have an installed Slack app configured with **Event Subscriptions** turned **On**, we need to setup a server and a web endpoint that accepts POST requests. Once we have a functional endpoint, we will insert the endpoint URL into the field *Request URL* on the **Events Subscriptions** page.

Let's setup a basic server that will run on a specified port:

```go
//this import contains all dependencies used in the example.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"net/http"
	"os"
	"strings"
)

//startServer will accept a given port on localhost, and listen for requests.
func startServer(port string) error {
	portNum := fmt.Sprintf(":%v", port)
	fmt.Println("Currently listening on localhost @ port ", portNum)
	err := http.ListenAndServe(portNum, nil)
	return err
}
```

Before you can save a new endpoint URL into the **Request URL** field, you'll be asked to have your endpoint return an HTTP response containing the value of `challenge` for verification purposes.

We'll need to model Slack's Event API JSON responses as well as the verification JSON object as Go structs in order to properly unmarshal them, and decode raw data we'll be receiving from the endpoint.

```go
//SlackPostVerification contains attributes that Slack needs to verify URL works.
type SlackPostVerification struct {
	Type      string `json:"type"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}

//Event encapsulates actual data we want from that posted to the webhook.
type Event struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	Ts      string `json:"event_ts"`
	Type    string `json:"type"`
	User    string `json:"user"`
}

//SlackMessageEvent represents JSON structure of a Slack message event.
type SlackMessageEvent struct {
	SlackPostVerification
	Event Event `json:"event"`
}

//Item is wraps stars.list JSON response.
type Item struct {
    Channel    string  `json:"channel"`
    DateCreate string  `json:"date_create"`
    Message    Message `json:"message"`
		File File `json:"file"`
    Type       string  `json:"type"`
}

//Message represents a subset of key/values from the starred.list endpoint.
type Message struct {
    IsStarred bool   `json:"is_starred"`
    Permalink string `json:"permalink"`
    Text      string `json:"text"`
    Ts        string `json:"ts"`
    Type      string `json:"type"`
    User      string `json:"user"`
}

//File represents a subset of key/values from the starred.list endpoint.
//A File object contains many attributes, but they've been redacted for this example.
type File struct {
    Name               string   `json:"name"`
    Permalink          string   `json:"permalink"`
    Timestamp          int64    `json:"timestamp"`
    UrlPrivateDownload string   `json:"url_private_download"`
}
```

`SlackPostVerification` is used for endpoint verification that must happen when inserting a new request URL in Slack's event subscription settings.

`Event` contains data from Slack's message event.

`Item` is the top-level object that contains either a [Message](https://api.slack.com/docs/messages) or a [File](https://api.slack.com/types/file) that's starred.

`Message` and `File` contain information about Slack's Message and File events respectively.

With these models, we can now build our endpoint. We'll want it to:

1. Only accept POST requests.

2. Decode raw JSON data into useful structs.

```go
//slackEventEndpoint accepts POST requests, which will contain slack events.
func slackEventEndpoint(w http.ResponseWriter, req *http.Request) {
	// only accept POST requests
	if req.Method != http.MethodPost {
		w.WriteHeader(405)
	} else {

		// translate body to JSON
		decoder := json.NewDecoder(req.Body)
		var slackMessage SlackMessageEvent
		if err := decoder.Decode(&slackMessage); err != nil {
			log.Fatal(err)
		}

		// return challenge as part of response.
		var response SlackPostVerification
		response.Challenge = slackMessage.Challenge

		if resp, err := json.Marshal(response); err != nil {
			log.Fatal(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(resp)
		}

		// parse consumed message - send Email based on params.

	}
}

```

With this functional endpoint, it's time to associate it with our server. We'll do this in the `main` function.

```go
func main() {
	http.HandleFunc("/slack_events", slackEventEndpoint) // <your_base_url>/post_events will be our endpoint.
	err := startServer("8080")
	if err != nil {
		log.Println(err)
	}
}
```

Now when our server runs, `/post_events` can receive incoming POST requests.

With all these pieces in place, it's time to insert our request URL into Slack. Since we are assuming a development environment, we will be using [ngrok](https://ngrok.com/) to expose a locally available endpoint.

Please follow installation [instructions](https://ngrok.com/download), and continue this tutorial once you are able to successfully run `./ngrok http <port number>`. Running this command will create a publicly accessible URL that tunnels to your localhost at the specified port, e.g. `http://a123bcde.ngrok.io`.

Just insert `<your_base_url>/post_events` into Slack, and make sure it verifies in the UI (`challenge` value was accepted). Event data should now be sent to your endpoint when a message is submitted on a channel in your Slack workspace.

<a name="starred"></a>
### Getting Starred Items

The last steps we need to complete before we can start sending triggered mail is setting up some helper functions to:

1. Make a GET request to `https://slack.com/api/stars.list`.

2. Process API response and format basic HTML content listing starred events.

For the first step, we want to write a function `slackGetStarred` that will return a request object that will be executed for JSON response data.

```go
//slackGetStarred makes an API call to `https://slack.com/api/stars.list`
func slackGetStarred() *http.Request {
	url := "https://slack.com/api/stars.list?token=" + os.Getenv("SLACK_OAUTH_TOKEN")
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req
}

//executeRequest wraps an HTTP client and returns request response as []bytes.
func executeRequest(r *http.Request) []byte {
	c := &http.Client{Timeout: time.Second * 10}
	resp, err := c.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
```

We will then write a function `populateStarred` that will take a string object, and then append to it formatted information from the `stars.list` endpoint.

```go
//populateStarred gets all starred items, and populates an email body that will be
//attached to the SG message.
func populateStarred(emailBody string) string {
	req := slackGetStarred()
	data := executeRequest(req)
	starred := UnmarshalSlackStarred(data)
	//Add messages
	for _, i := range(starred.Items) {
		item := ""
		messageTS, _ := strconv.ParseFloat(i.Message.Ts, 64)
		if i.Type == "message" {
			item += fmt.Sprintf("<p>Posted on %v</p><br><p><b>%v</br></p></br><p>Link to Post: %v</p></br>", time.Unix(int64(messageTS), 0), i.Message.Text, i.Message.Permalink)
		}
		emailBody += item
	}

	//Add files
	for _, i := range(starred.Items) {
		item := ""
		if i.Type == "file" {
			item += fmt.Sprintf("<p>Posted on %v</p><br><p><b>%v</br></p></br><p>Download Link: %v</p></br>", time.Unix(i.File.Timestamp, 0), i.File.Name, i.File.UrlPrivateDownload)
		}
		emailBody += item
	}
	return emailBody
}
```

<a name="triggered_mail"></a>
### Event Triggered Mail

Now that we've laid out the infrastructure (Slack setup, server, API helpers functions and endpoint), we can start handling incoming Slack events. In this case we are listening for the [message.channels](https://api.slack.com/events/message.channels) event type, and in the previous section we created Go structs to help us decode this event data.

We need to now add code that will:

1. Handle a `message.channels` event and parse out individual text components. We expect the format of a command to be `sgstarred|test@example.com`.

2. Build a mail object using the SendGrid-Go API library. We will use the helper functions we wrote above to format email content.

3. Send the email.

```go
//slackEventEndpoint accepts POST requests, which will contain slack events.
func slackEventEndpoint(w http.ResponseWriter, req *http.Request) {
	// only accept POST requests
	if req.Method != http.MethodPost {
		w.WriteHeader(405)
	} else {

		// translate body to JSON
		decoder := json.NewDecoder(req.Body)
		var slackMessage SlackMessageEvent
		if err := decoder.Decode(&slackMessage); err != nil {
			log.Fatal(err)
		}

		// return challenge as part of response.
		var response SlackPostVerification
		response.Challenge = slackMessage.Challenge
		if resp, err := json.Marshal(response); err != nil {
			log.Fatal(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(resp)
		}
		// parse consumed message - send Email based on params.
		eventParsed := parseMessageEvent(slackMessage.Event)
		if eventParsed[0] == "sgstarred" {
			if len(eventParsed) == 2 {
				m := convertEventMail(eventParsed)
				sendSGMessage(m)
			}
		}
	}
}

//parseMessageEvent reads the text field of `Event`, and splits it into
// ["<cmd>", "<body>", "<from_address>", "<to_address>"]
func parseMessageEvent(e Event) []string {
    eventParsed := strings.Split(e.Text, "|")
    return eventParsed
}

//convertEventMail takes relevant info from Event and populates a mail object.
func convertEventMail(eventParsed []string) *mail.SGMailV3 {
	m := mail.NewV3Mail()
	from := mail.NewEmail("From Slack Channel", "<a_from_address>")
	mailBody := "<p>Hi there!\nHere are the five most recent starred items on Slack:<p></br>"

	withStarred := populateStarred(mailBody)
	content := mail.NewContent("text/html", withStarred)
	to := mail.NewEmail("", eventParsed[1])
	p := mail.NewPersonalization()

	m.SetFrom(from)
	m.AddContent(content)
	p.AddTos(to)
	p.Subject = "Starred Items - From Slack"
	m.AddPersonalizations(p)
	return m
}

//sendSGMessage wraps SendGrid library
func sendSGMessage(m *mail.SGMailV3) {
  request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
  request.Method = "POST"
  request.Body = mail.GetRequestBody(m)
  response, err := sendgrid.API(request)

  if err != nil {
    log.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
  return
}
```

<a name="conclusion"></a>
### Conclusion

Now when we submit a message of the format `sgstarred|test@example.com` in any channel on our Slack workspace, our endpoint will process the event data, make a call to a Slack API endpoint and then send a custom formatted email using SendGrid.

One way to improve upon this example is to create a [Transactional Template](https://sendgrid.com/docs/User_Guide/Transactional_Templates/index.html), and then reference it in our mail send function and use [Substitutions](https://sendgrid.com/docs/API_Reference/SMTP_API/substitution_tags.html).

We covered a basic integration - but there's many different [event types](https://api.slack.com/events/api) that Slack's API offers. The possibilities are plentiful!
