# Send a Single Email to a Single Recipient

The following code assumes you are storing the API key in an environment variable (recommended). 

This is the minimum code needed to send an email.

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
    from := mail.NewFrom("test@example.com", "Example User")
    to := mail.NewTo("test@example.com", "Example User")
    subject := mail.NewSubject("Sending with Twilio SendGrid is Fun")
    plainTextContent := mail.NewPlainTextContent("and easy to do anywhere, even with Go")
    htmlContent := mail.NewHtmlContent("<strong>and easy to do anywhere, even with Go</strong>")
    email := mail.NewMessage(from,
                             to,
                             subject, 
                             plainTextContent,
                             htmlContent)
    sendgrid := sendgrid.NewClient(os.Getenv("SENDGRID_API_KEY"))
    response, err := sendgrid.Send(email)
    if err != nil {
        log.Println(err)
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
    "log"
    "os"

    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
    from := mail.NewFrom("test@example.com", "Example User")
    to := []*mail.Tos{
        mail.NewTo("test@example1.com", "Example User 1"),
        mail.NewTo("test@example2.com", "Example User 2"),
    }
    subject := mail.NewSubject("Sending with Twilio SendGrid is Fun")
    plainTextContent := mail.NewPlainTextContent("and easy to do anywhere, even with Go")
    htmlContent := mail.NewHtmlContent("<strong>and easy to do anywhere, even with Go</strong>")
    email := mail.NewMessage(from,
                             to...,
                             subject, 
                             plainTextContent,
                             htmlContent)
    sendgrid := sendgrid.NewClient(os.Getenv("SENDGRID_API_KEY"))
    response, err := sendgrid.Send(email)
    if err != nil {
        log.Println(err)
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
    "log"
    "os"

    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
    from := mail.NewFrom("test@example.com", "Example User")
    tos := []*mail.Tos{
        mail.NewTo("test1@example.com",
                   "Example User 1",
                   map[string]string{
                       "-name-", "Alain",
                       "-github-", "http://github.com/ninsuo",
                   }),
        mail.NewTo("test2@example.com",
                   "Example User 2",
                   map[string]string{
                        "-name-", "Elmer",
                        "-github-", "http://github.com/thinkingserious",
                    }),),
    }
    // Alternatively, you can pass in a map of subjects OR add a subject as a parameter to the `To` object
    subject := mail.NewSubject("Hi -name-!")
    globalSubstitution = mail.NewGlobalSubstitution("-time-", "<Current Time>")
    plainTextContent := mail.NewPlainTextContent("and easy to do anywhere, even with Go")
    htmlContent := mail.NewHtmlContent("<strong>and easy to do anywhere, even with Go</strong>")
    email := mail.NewMessage(from,
                             to...,
                             subject, // or subject...
                             plainTextContent,
                             htmlContent,
                             globalSubstitution) // or globalSubstition...
    sendgrid := sendgrid.NewClient(os.Getenv("SENDGRID_API_KEY"))
    response, err := sendgrid.Send(email)
    if err != nil {
        log.Println(err)
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
    "log"
    "os"

    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
    from := mail.NewFrom("test@example.com", "Example User")
    to := mail.NewTo("test@example.com", "Example User")
    subject := mail.NewSubject("Sending with Twilio SendGrid is Fun")
    plainTextContent := mail.NewPlainTextContent("and easy to do anywhere, even with Go")
    htmlContent := mail.NewHtmlContent("<strong>and easy to do anywhere, even with Go</strong>")
    email := mail.NewMessage(from,
                             to,
                             subject, 
                             plainTextContent,
                             htmlContent)
    // For a detailed description of each of these settings, please see the [documentation](https://sendgrid.com/docs/API_Reference/api_v3.html).

    email.AddTo("test1@example.com", "Example User1")
    to := []*mail.Tos{
        mail.NewTo("test2@example.com", "Example User2"),
        mail.NewTo("test3@example.com", "Example User3"),
    }
    email.AddTos(cc...)

    email.AddCcs("test1@example4.com", "Example User4")
    cc := []*mail.Ccs{
        mail.NewCcs("test5@example.com", "Example User5"),
        mail.NewCcs("test6@example.com", "Example User6"),
    }
    email.AddCcs(cc...)

    email.AddBcc("test1@example7.com", "Example User7")
    bcc := []*mail.Bccs{
        mail.NewBcc("test8@example.com", "Example User8"),
        mail.NewBcc("test9@example.com", "Example User9"),
    }
    email.AddBccs(bcc...)

    email.AddHeader("X-Test1", "Test1")
    email.AddHeader("X-Test2", "Test2")
    headers := []*mail.Headers{
        mail.NewHeader("X-Test3", "Test3"),
        mail.NewHeader("X-Test4", "Test4"),
    }
    email.AddHeaders(headers)

    email.AddSubstitution("%name1%", "Example Name 1")
    email.AddSubstitution("%city1%", "Denver")
    substitutions := []*mail.Substitutions{
        mail.NewSubstitution("%name2%", "Example Name 2"),
        mail.NewSubstitution("%city2%", "Orange"),
    }
    email.AddSubstitutions(substitutions)

    email.AddCustomArg("marketing1", "false")
    email.AddCustomArg("transactional1", "true")
    customArgs := []*mail.CustomArgs{
        mail.NewCustomArg("marketing2", "true"),
        mail.NewCustomArg("transactional2", "false"),
    }
    email.AddCustomArgs(customArgs)

    email.SetSendAt(1461775051)

    // If you need to add more [Personalizations](https://sendgrid.com/docs/Classroom/Send/v3_Mail_Send/personalizations.html), here is an example of adding another Personalization by passing in a personalization index

    email.AddTo("test10@example.com", "Example User10", 1)
    to := []*mail.Tos{
        mail.NewTo("test11@example.com", "Example User11"),
        mail.NewTo("test12@example.com", "Example User12"),
    }
    email.AddCcs(cc..., 1)

    email.AddCcs("test13@example.com", "Example User13", 1)
    cc := []*mail.Ccs{
        mail.NewCcs("test14@example.com", "Example User14"),
        mail.NewCcs("test15@example.com", "Example User15"),
    }
    email.AddCcs(cc..., 1)

    email.AddBcc("test16@example.com", "Example User16", 1)
    bcc := []*mail.Bccs{
        mail.NewBcc("test17@example.com", "Example User17"),
        mail.NewBcc("test18@example.com", "Example User18"),
    }
    email.AddBccs(bcc..., 1)

    email.AddHeader("X-Test5", "Test5", 1)
    email.AddHeader("X-Test6", "Test6", 1)
    headers1 := []*mail.Headers{
        mail.NewHeader("X-Test7", "Test7"),
        mail.NewHeader("X-Test8", "Test8"),
    }
    email.AddHeaders(headers1, 1)

    email.AddSubstitution("%name3%", "Example Name 3", 1)
    email.AddSubstitution("%city3%", "Redwood City", 1)
    substitutions1 := []*mail.Substitutions{
        mail.NewSubstitution("%name4%", "Example Name 4"),
        mail.NewSubstitution("%city4%", "London"),
    }
    email.AddSubstitutions(substitutions1, 1)

    email.AddCustomArg("marketing3", "true", 1)
    email.AddCustomArg("transactional3", "false", 1)
    customArgs1 := []*mail.CustomArgs{
        mail.NewCustomArg("marketing4", "false"),
        mail.NewCustomArg("transactional4", "true"),
    }
    email.AddCustomArgs(customArgs1, 1)

    email.SetSendAt(1461775051, 1)

    email.SetSubject("this subject overrides the Global Subject", 1)

    // The values below this comment are global to entire message

    email.SetFrom("test@example.com", "Example User 0")

    email.SetGlobalSubject("Sending with Twilio SendGrid is Fun")

    email.AddContent(sendgrid.MimeType.Text, "and easy to do anywhere, even with Go")
    email.AddContent(sendgrid.MimeType.Html, "<strong>and easy to do anywhere, even with Go</strong>")
    contents := []*mail.Contents{
        mail.NewContent("text/calendar", "Party Time!!"),
        mail.NewContent("text/calendar2", "Party Time2!!"),
    }
    email.AddContents(contents)

    email.AddAttachment("balance_001.pdf",
                        "base64 encoded string",
                        "application/pdf",
                        "attachment",
                        "Balance Sheet")
    attachments := []*mail.Attachments{
        mail.NewAttachment("banner.png",
                           "base64 encoded string",
                           "image/png",
                           "inline",
                           "Banner"),
        mail.NewAttachment("banner2.png",
                           "base64 encoded string",
                           "image/png",
                           "inline",
                           "Banner2"),
    }
    email.AddAttachments(attachments)

    email.SetTemplateId("13b8f94f-bcae-4ec6-b752-70d6cb59f932")

    email.AddGlobalHeader("X-Day", "Monday")
    globalHeaders := []*mail.Headers{
        mail.NewHeader("X-Test3", "Test3"),
        mail.NewHeader("X-Test4", "Test4"),
    }
    email.AddGlobalHeaders(globalHeaders)

    email.AddSection("%section1", "Substitution for Section 1 Tag")
    sections := []*mail.Sections{
        mail.NewSection("%section2%", "Substitution for Section 2 Tag"),
        mail.NewSection("%section3%", "Substitution for Section 3 Tag"),
    }
    email.AddSections(sections)

    email.AddCategory("customer")
    categories := []*mail.Categories{
        mail.NewCategory("vip"),
        mail.NewCategory("new_account"),
    }
    email.AddCategories(categories)

    email.AddGlobalCustomArg("campaign", "welcome")
    globalCustomArgs := []*mail.CustomArgs{
        mail.NewCustomArg("sequence2", "2"),
        mail.NewCustomArg("sequence3", "3"),
    }
    email.AddGlobalCustomArgs(globalCustomArgs)

    asmGroups := []*mail.AsmGroups{
        mail.NewAsmGroup(1),
        mail.NewAsmGroup(4),
        mail.NewAsmGroup(5),
    }
    email.SetAsm(3, asmGroups)

    email.SetGlobalSendAt(1461775051)

    email.SetIpPoolName("23")

    // This must be a valid [batch ID](https://sendgrid.com/docs/API_Reference/SMTP_API/scheduling_parameters.html)
    //email.SetBatchId("some_batch_id")

    email.SetBccSetting(true, "test@example.com")

    email.SetBypassListManagement(true)

    email.SetFooterSetting(true, "Some Footer HTML", "Some Footer Text")

    email.SetSandBoxMode(true)

    email.SetSpamCheck(true, 1, "https://gotchya.example.com")

    email.SetClickTracking(true, false)

    email.SetOpenTracking(true, "Optional tag to replace with the open image in the body of the message")

    email.SetSubscriptionTracking(true,
                                  "HTML to insert into the text / html portion of the message",
                                  "text to insert into the text/plain portion of the message",
                                  "substitution tag")

    email.SetGoogleAnalytics(true,
                             "some campaign",
                             "some content",
                             "some medium",
                             "some source",
                             "some term")

    email.SetReplyTo("test+reply@example.com", "Reply To Me")      

    sendgrid := sendgrid.NewClient(os.Getenv("SENDGRID_API_KEY"))
    response, err := sendgrid.Send(email)
    if err != nil {
        log.Println(err)
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
    "log"
    "os"

    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
    from := mail.NewFrom("test@example.com", "Example User")
    to := mail.NewTo("test@example.com", "Example User")
    subject := mail.NewSubject("Sending with Twilio SendGrid is Fun")
    plainTextContent := mail.NewPlainTextContent("and easy to do anywhere, even with Go")
    htmlContent := mail.NewHtmlContent("<strong>and easy to do anywhere, even with Go</strong>")
    email := mail.NewMessage(from,
                             to,
                             subject, 
                             plainTextContent,
                             htmlContent)
    email.AddAttachment("balance_001.pdf",
                        "base64 encoded string",
                        "application/pdf",
                        "attachment",
                        "Balance Sheet")

    sendgrid := sendgrid.NewClient(os.Getenv("SENDGRID_API_KEY"))
    response, err := sendgrid.Send(email)
    if err != nil {
        log.Println(err)
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
    "log"
    "os"

    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
    from := mail.NewFrom("test@example.com", "Example User")
    to := mail.NewTo("test@example.com", "Example User")
    subject := mail.NewSubject("Sending with Twilio SendGrid is Fun")
    plainTextContent := mail.NewPlainTextContent("and easy to do anywhere, even with Go")
    htmlContent := mail.NewHtmlContent("<strong>and easy to do anywhere, even with Go</strong>")
    email := mail.NewMessage(from,
                             to,
                             subject, 
                             plainTextContent,
                             htmlContent)
    // See `Send Multiple Emails to Multiple Recipients` for additional methods for adding substitutions
    email.AddSubstitution("-name-", "Example User")
    email.AddSubstitution("-city-", "Denver")
    email.SetTemplateId("13b8f94f-bcae-4ec6-b752-70d6cb59f932")
    sendgrid := sendgrid.NewClient(os.Getenv("SENDGRID_API_KEY"))
    response, err := sendgrid.Send(email)
    if err != nil {
        log.Println(err)
    } else {
        fmt.Println(response.StatusCode)
        fmt.Println(response.Body)
        fmt.Println(response.Headers)
    }
}
```
