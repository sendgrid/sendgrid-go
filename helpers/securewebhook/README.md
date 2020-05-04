**This helper allows you to quickly and easily enable/disable SecureWebhook feature or get the public key through Twilio SendGrid.**

## Dependencies

- [rest](https://github.com/sendgrid/rest)

# Quick Start

Run the [example](https://github.com/sendgrid/sendgrid-go/blob/master/examples/securewebhook/securewebhook.go) (make sure you have set your environment variable to include your SENDGRID_API_KEY).
```bash
go run examples/securewebhook/securewebhook.go
```

## Usage

- See the [example](https://github.com/sendgrid/sendgrid-go/blob/master/examples/securewebhook/securewebhook.go) for a complete working example.
- [Documentation](https://sendgrid.com/docs/for-developers/tracking-events/)

## Test

```bash
go test ./... -v
```

or

```bash
cd helpers/securewebhook
go test -v
```
