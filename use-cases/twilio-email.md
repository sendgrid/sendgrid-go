First, follow the [Twilio Setup](twilio-setup.md) guide for creating a Twilio account and setting up environment variables with the proper credentials.

Then, initialize the Twilio Email Client.

```go
mailClient := NewTwilioEmailSendClient(os.Getenv("TWILIO_API_KEY"), os.Getenv("TWILIO_API_SECRET"))

// or

mailClient := NewTwilioEmailSendClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"))
```

This sets the client to use Twilio Auth and the Twilio Email API.
