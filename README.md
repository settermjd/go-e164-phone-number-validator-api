# A Go-based E.164 Phone Number Validator API

![Main branch Go Workflow](https://github.com/settermjd/go-e164-phone-number-validator-api/actions/workflows/go.yml/badge.svg)

This is a small project that shows how to validate a phone number as being in [E.164 format](https://www.twilio.com/docs/glossary/what-e164) or not, using Go and its [net/http](https://pkg.go.dev/net/http) package (among others).

It isn't mean to be a production-ready application, rather one that just shows the essentials of what's required to validate a phone number as being in E.164 format, or not, wrapped in a small API.

## Usage

To use the application, first clone it locally, then start it by running the following command.

```bash
go run main.go
```

Then, using curl, or another HTTP client, send a POST request to the default route, with one form parameter, `phone_number`; for example:

```bash
curl -i --form "phone_number=+61 176 9876 1234" http://localhost:4000
```

The application doesn't check if the phone number is active, just if it's formatted in E.164 format.
You'll then get back a JSON response informing you whether the phone number is formatted correctly, or not.

## Have an issue?

If so, please create a new [issue](https://github.com/settermjd/go-e164-phone-number-validator-api/issues/new/choose) or PR.