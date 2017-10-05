# Send a Single Email to a Single Recipient

The following code assumes you are storing the API key in an environment variable (recommended).

This is the minimum code needed to send a message.

```go
package main

import (
        "fmt"
        "os"

        "github.com/sendgrid/sendgrid-go"
        "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
        message := &mail.Message{
                To:          []mail.To{mail.To{
                                     "Email": "test@example.com",
                                     "Name": "Example Recipient",
                             }},
                From:        mail.Email{"from@example.com", "Example Sender"},
                Subject:     "Test Email Subject",
                TextContent: "Text Email Content",
                HTMLContent: "<strong>HTML Email Content<strong>",
        }

        client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
        response, err := client.Send(message)
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println(response.StatusCode)
                fmt.Println(response.Body)
                fmt.Println(response.Headers)
        }
}
```

# Send a Single Email to Multiple Recipients

The following code assumes you are storing the API key in an environment variable (recommended).

```go
package main

import (
        "fmt"
        "os"

        "github.com/sendgrid/sendgrid-go"
        "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
        message := &mail.Message{
                To:     []mail.To{
                        mail.To{"Email": "test1@example.com", "Name": "Example Recipient1"},
                        mail.To{"Email": "test2@example.com", "Name": "Example Recipient3"},
                        mail.To{"Email": "test3@example.com", "Name": "Example Recipient3"},
                },
                From:        mail.Email{"from@example.com", "Example Sender"},
                Subject:     "Test Email Subject",
                TextContent: "Text Email Content",
                HTMLContent: "<strong>HTML Email Content<strong>",
        }

        client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
        response, err := client.Send(message)
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println(response.StatusCode)
                fmt.Println(response.Body)
                fmt.Println(response.Headers)
        }
}
```

# Send Multiple Emails to Multiple Recipients

The following code assumes you are storing the API key in an environment variable (recommended).
```go
package main

import (
        "fmt"
        "os"
        "time"

        "github.com/sendgrid/sendgrid-go"
        "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
        message := &mail.Message{
                From: mail.Email{"from@example.com", "Example Sender"},
        }
	
	// mail.To struct encompasses the personalizations for the recipient
	// You can add substitutions, categories, custom arguments, etc. for each recipient to the corresponding mail.To
	// object.
	// Learn more about personalization [here](https://sendgrid.com/docs/API_Reference/api_v3.html)
	// and [here](https://sendgrid.com/docs/Classroom/Send/v3_Mail_Send/personalizations.html)

        to1 := &mail.To{
                "Email": "test1@example.com",
                "Name": "Test Recipient",
                "Substitutions": map[string]string{
                        "-name-":"Alain",
                        "-github-", "http://github.com/test1",
                },
        }
        message.AddTo(to1)

        moreTos := []mail.To{
                mail.To{
                        "Email": "test2@example.com",
                        "Name": "Test Recipient",
                        "Substitutions": map[string]string{
                                "-name-":"Elmer",
                                "-github-", "http://github.com/thinkingserious",
                        },
                        "Subject": "Override global subject",
                },
                mail.To{
                        "Email": "test3@example.com",
                        "Name": "Test Recipient",
                        "Substitutions": map[string]string{
                                "-name-":"Matt",
                                "-github-", "http://github.com/test3",
                        },
                },
        }
        message.AddTos(moreTos)

        message.AddGlobalSubstitutions(map[string]string{"-time-": time.Now().String()})

        message.SetSubject("Hi -name-!")
        message.SetTextContent("Hello -name-, your github is -github-, email sent at -time-")
        message.SetHTMLContent("<strong>Hello -name-, your github is <a href=\"-github-\">here</a></strong> email sent at -time-")

        client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
        response, err := client.Send(message)
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println(response.StatusCode)
                fmt.Println(response.Body)
                fmt.Println(response.Headers)
        }
}
```

# Kitchen Sink - an example with all settings used

The following code assumes you are storing the API key in an environment variable (recommended).
```go
package main

import (
        "fmt"
        "os"

        "github.com/sendgrid/sendgrid-go"
        "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
        message := &mail.Message{}

        to1 := &mail.To{mail.To{"Email": "test1@example.com", "Name": "Test Recipient 1"}}
        to2 := &mail.To{mail.To{"Email": "test2@example.com", "Name": "Test Recipient 2"}}

        to1.AddCC(&mail.Email{"cc1@example.com", "Test Recipient"})
        ccs := []mail.Email{
                mail.Email{"cc2@example.com", "Test Recipient"}.
                mail.Email{"cc3@example.com", "Test Recipient"},
        }
        to2.AddCCs(ccs)

        to1.AddBCC(&mail.Email{"bcc1@example.com", "Test Recipient"})
        bccs := []mail.Email{
                mail.Email{"bcc2@example.com", "Test Recipient"}.
                mail.Email{"bcc3@example.com", "Test Recipient"},
        }
        to2.AddBCCs(bccs)

        to1.AddHeader("X-Hdr1", "Test1")
        hdrs := map[string]string{
                "X-Hdr2", "Test2",
                "X-Hdr3", "Test3",
        }
        to2.AddHeaders(hdrs)

        to1.AddSubstitution("%City1%", "Denver")
        to1.AddSubstitution("%name1%": "Name 1")
        substitutions := map[string]string{
                "%City2%", "Orange",
                "%name2%": "Name 2",
        }
        to2.AddSubstitutions(substitutions)

        to1.AddCustomArg("Message-Category", "Marketing")
        customArgs := map[string]string{
                "Campaign-ID": "Mkt-123",
                "Location": "US-West",
        }
        to2.AddCustomArgs(customArgs)

        to1.SetSendAt(1461775051)

        to1.SetSubject("Override subject for Rcpt 1")

        message.AddTo(to1)
        message.AddTo(to2)

        // If you need to add more [Personalizations](https://sendgrid.com/docs/Classroom/Send/v3_Mail_Send/personalizations.html),
        // Here is an example of adding another Personalization by passing in a personalization index

        // Global message level methods
        message.SetFrom(&mail.Email{"from@example.com", "Example Sender"})

        message.SetReplyTo(&mail.Email{"replyto@example.com", "Reply To"})

        message.SetSubject("Sending Email is fun!")

        message.AddContent(&mail.Content{
                "Type": "text/plain",
                "Value": "Text content",
        })
        message.AddContent(&mail.Content{
                "Type": "text/html",
                "Value": "<strong> HTML content</strong>",
        })

        message.AddAttachment(&mail.Attachment{
                "filename": "balance_001.pdf",
                "content": "base64 encoded string",
                "type": "application/pdf",
                "disposition": "attachment",
        })

        message.SetTemplateId("13b8f94f-bcae-4ec6-b752-70d6cb59f932")

        message.AddSection("%section1", "Substitution for Section 1 Tag")
        sections := map[string]string{
                "%section2%", "Substitution for Section 2 Tag",
                "%section3%", "Substitution for Section 3 Tag",
        }
        message.AddSections(sections)

        message.AddCategory("customer")
        categories := []string{"new_account", "aws"}
        message.AddCategories(categories")

        message.AddCustomArg("campaign", "welcome")
        globalCustomArgs := map[string]string{
                "sequence2", "2",
                "sequence3", "3",
        }
        message.AddCustomArgs(globalCustomArgs)

        asmGroupIds := []int{1, 4, 5}
        message.SetAsm(3, asmGroupIds)

        message.SetSendAt(1461775051)

        message.SetIpPoolName("23")

        // mail settings
        message.SetBccSetting(true, "test@example.com")
        message.SetBypassListManagement(true)
        message.SetFooterSetting(true, "Some Footer HTML", "Some Footer Text")
        message.SetSandBoxMode(true)
        message.SetSpamCheck(true, 1, "https://gotchya.example.com")

        // tracking settings
        message.SetClickTracking(true, false)
        message.SetOpenTracking(true, "Optional tag to replace with the open image in the body of the message")
        message.SetSubscriptionTracking(true,
                                        "HTML to insert into the text / html portion of the message",
                                        "text to insert into the text/plain portion of the message",
                                        "substitution tag")

        message.SetGoogleAnalytics(true, "some campaign", "some content", "some medium", "some source", "some term")

        client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
        response, err := client.Send(message)
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println(response.StatusCode)
                fmt.Println(response.Body)
                fmt.Println(response.Headers)
        }
}

```

# Attachments

The following code assumes you are storing the API key in an environment variable (recommended).

```go
package main

import (
        "fmt"
        "os"

        "github.com/sendgrid/sendgrid-go"
        "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
        message := &mail.Message{
                To:          []mail.To{
                                     mail.To{"Email": "test@example.com", "Name": "Example Recipient"},
                             },
                From:        mail.Email{"from@example.com", "Example Sender"},
                Subject:     "Test Email Subject",
                TextContent: "Text Email Content",
                HTMLContent: "<strong>HTML Email Content<strong>",
        }

        message.AddAttachment(&mail.Attachment{
                "filename": "yosemite.jpg",
                "content": "base64 encoded string",
                "type": "image/jpg",
                "disposition": "inline",
                "content_id": "image_1",
        })

        attachments := []mail.Attachment{
                mail.Attachment{
                        "filename": "el_capitan.jpg",
                        "content": "base64 encoded string 2",
                        "type": "image/jpg",
                        "disposition": "inline",
                        "content_id": "image_2",
                },
                mail.Attachment{
                        "filename": "sierra.jpg",
                        "content": "base64 encoded string 3",
                        "type": "image/jpg",
                        "disposition": "inline",
                        "content_id": "image_3",
                },
        }
        message.AddAttachments(attachments)

        client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
        response, err := client.Send(message)
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println(response.StatusCode)
                fmt.Println(response.Body)
                fmt.Println(response.Headers)
        }
}

```

# Transactional Templates

The following code assumes you are storing the API key in an environment variable (recommended).

For this example, we assume you have created a [transactional template](https://sendgrid.com/docs/User_Guide/Transactional_Templates/index.html). Following is the template content we used for testing.

Template ID (replace with your own):

```text
13b8f94f-bcae-4ec6-b752-70d6cb59f932
```

Email Subject:

```text
<%subject%>
```

Template Body:

```html
<html>
<head>
	<title></title>
</head>
<body>
Hello -name-,
<br /><br/>
I'm glad you are trying out the template feature!
<br /><br/>
<%body%>
<br /><br/>
I hope you are having a great day in -city- :)
<br /><br/>
</body>
</html>
```

```go
package main

import (
        "fmt"
        "os"

        "github.com/sendgrid/sendgrid-go"
        "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
        message := &mail.Message{
                To:   []mail.To{mail.To{
                              "Email": "test@example.com",
                              "Name": "Example Recipient",
                              "Substitutions": map[string]string{
                                      "-name-", "Example User",
                              },
                      }},
                From: mail.Email{"from@example.com", "Example Sender"},
        }

        message.SetTemplateId("13b8f94f-bcae-4ec6-b752-70d6cb59f932")
        message.AddSubstitution("-city-", "Orange")
        message.AddSubstitution("%subject%", "A message with Template")

        client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
        response, err := client.Send(message)
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println(response.StatusCode)
                fmt.Println(response.Body)
                fmt.Println(response.Headers)
        }
}

```
