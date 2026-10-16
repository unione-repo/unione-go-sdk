# UniOne Go SDK 1.1.2

Web API 1.90

Official Go library for the [UniOne](https://unione.io) transactional email API.

Quickstart guide: [Send emails with Go](https://docs.unione.io/en/send-with-go) · Docs for LLMs: [llms.txt](https://docs.unione.io/llms.txt) · [API reference](https://docs.unione.io/en/web-api-ref)

## Install

```bash
go get github.com/unione-repo/unione-go-sdk
```

Requirements: Go 1.18+. No third-party dependencies.

## Get an API key

1. Sign up at [cp.unione.io](https://cp.unione.io/en/user/registration/).
2. Copy your API key from [Settings → Security → API](https://cp.unione.io/en/settings/security/api).
3. Put it in an environment variable:

```bash
export UNIONE_API_KEY="your-api-key"
```

> **Server-side only.** Use this SDK in backend code: API routes, server functions, workers.
> Never ship the API key to a browser or a mobile app, and never commit it.

## Send your first email — no DNS setup

Every account can send test email right away from a **sandbox domain**: pre-verified, DKIM and SPF already set.
It delivers only to addresses you confirm.

<!-- checked examples -->

**1. Get your sandbox domain** (`domain/get-sandbox-domain`). The call creates the domain on first use and returns the same one afterwards.

```go
package main

import (
	"context"
	"fmt"
	"os"

	unione "github.com/unione-repo/unione-go-sdk"
)

func main() {
	client := unione.NewClient(os.Getenv("UNIONE_API_KEY"))

	response, err := client.Domain.GetSandboxDomain(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Domain)
}
```

**2. Confirm your inbox** (`confirmed-email/create`). UniOne sends a confirmation email; click the link in it within 48 hours.

```go
package main

import (
	"context"
	"fmt"
	"os"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/confirmedemail"
)

func main() {
	client := unione.NewClient(os.Getenv("UNIONE_API_KEY"))

	response, err := client.ConfirmedEmail.Create(context.Background(), &confirmedemail.ConfirmedEmailRequest{
		Email: "you@example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.ConfirmedEmail)
}
```

Check the status with `confirmed-email/list` — your address should show `"status": "confirmed"`.

**3. Send the email** (`email/send`) from any address on the sandbox domain to the confirmed address.

```go
package main

import (
	"context"
	"fmt"
	"os"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/email"
	"github.com/unione-repo/unione-go-sdk/common"
)

func main() {
	client := unione.NewClient(os.Getenv("UNIONE_API_KEY"))

	response, err := client.Email.Send(context.Background(), &email.EmailSendRequest{
		Message: &email.MessageObject{
			Recipients: []email.MessageObjectRecipientsItem{
				{
					Email: "you@example.com",
				},
			},
			Body: &common.BodyObject{
				HTML: unione.String("<p>It works!</p>"),
			},
			Subject:   "Hello from UniOne",
			FromEmail: unione.String("noreply@<sandbox-domain>"),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.JobID)
}
```

<!-- /checked examples -->

### What you get back

```json
{
  "status": "success",
  "job_id": "1ZymBc-00041N-9X",
  "emails": [
    "you@example.com"
  ]
}
```

On error the API returns `status: "error"`, a `message` and a numeric `code`:

```json
{
  "status": "error",
  "code": 102,
  "message": "Error ID:83274F98-C7D5-11F1-B221-6214FF8707B1. Can not decode key wrongkey"
}
```

### Same flow with curl

```bash
BASE=https://api.unione.io/en/transactional/api/v1

curl -s -X POST $BASE/domain/get-sandbox-domain.json \
  -H "Content-Type: application/json" -H "X-API-KEY: $UNIONE_API_KEY" -d '{}'

curl -s -X POST $BASE/confirmed-email/create.json \
  -H "Content-Type: application/json" -H "X-API-KEY: $UNIONE_API_KEY" \
  -d '{"email": "you@example.com"}'

curl -s -X POST $BASE/email/send.json \
  -H "Content-Type: application/json" -H "X-API-KEY: $UNIONE_API_KEY" \
  -d '{
        "message": {
          "recipients": [
            {
              "email": "you@example.com"
            }
          ],
          "from_email": "noreply@<sandbox-domain>",
          "subject": "Hello from UniOne",
          "body": {
            "html": "<p>It works!</p>"
          }
        }
      }'
```

## Sandbox limits

- Sends only to confirmed addresses: up to 10 per account.
- Up to 100 emails a day.
- The sandbox domain expires after 31 days; call `domain/get-sandbox-domain` again to get a new one.
- Works only with the account API key, not a project key.
- Not listed in `domain/list`; the other `domain/*` methods do not apply to it.

The sandbox is for development and testing. Do not use a sandbox domain as the sender in production.

## Before production: verify your domain

1. `domain/get-dns-records` — returns the values for your DNS records.
2. Add these records at your DNS provider (the domain page of the control panel shows them too):

   | Type | Host | Value |
   |---|---|---|
   | TXT | `@` | the `verification-record` value (`unione-validate-hash=…`) |
   | TXT | `<selector>._domainkey` (the selector is on the domain page of the control panel, e.g. `us`) | `v=DKIM1; k=rsa; p=<dkim>` |
   | TXT | `@` | `v=spf1 include:spf.unione.io ~all` |
   | CNAME | `_dmarc` | the `dmarc` value |

3. `domain/validate-verification-record`, then `domain/validate-dkim`.
4. `domain/list` — wait until the verification record is confirmed and DKIM is active.
5. Read `from_email` from your config, so switching from the sandbox needs no code change.

## Sending reliably

- **Retries:** set `message.idempotence_key` on sends that may be retried, so a repeated request does not send twice.
- **Limits:** up to 500 recipients and 10 MB per request; each attachment up to 7 MB.
- **Body:** always include `body.html` or `body.plaintext`.
- **Delivery status:** register a webhook (`webhook/set`) with a fast, idempotent HTTPS handler instead of polling.
- **Defaults:** keep the default unsubscribe footer and tracking; changing them requires UniOne support.
- **Errors:** retry `429` and `5xx` with backoff; do not retry other errors unchanged.

## Links

- [Send emails with Go](https://docs.unione.io/en/send-with-go) and other [integration guides](https://docs.unione.io/en/integrations/)
- [Web API reference](https://docs.unione.io/en/web-api-ref)
- [llms.txt](https://docs.unione.io/llms.txt) for AI coding agents
- [UniOne MCP server and agent setup](https://unione.io/en/agents)
- [Support](https://cp.unione.io/en/support)

## For SDK contributors

See [AGENTS.md](AGENTS.md): how this SDK is generated from the OpenAPI spec, built, tested and released.

---

# SDK reference

## API host

The second argument selects the API host (datacenter) or a full base URL.

```go
client := unione.NewClient("YOUR_API_KEY", "us1.unione.io")
client = unione.NewClient("YOUR_API_KEY", unione.HostUS1)
```

## Error handling

Every method raises `*client.APIError` when the API returns an error; it carries the [API error code](https://docs.unione.io/en/web-api-ref#api-errors) and the HTTP status.

A transport error (timeout, DNS, TLS, canceled context) is wrapped: `errors.Is(err, context.DeadlineExceeded)` and `errors.As(err, &netErr)` work through `*client.APIError` (`Err`, `Unwrap()`).

```go
response, err := client.Email.Send(ctx, request)
var apiErr *client.APIError
var netErr net.Error
switch {
case errors.Is(err, context.DeadlineExceeded):
	// the timeout of ctx or of the HTTP client
case errors.As(err, &netErr):
	// network error
case errors.As(err, &apiErr) && apiErr.Code != nil:
	fmt.Println(apiErr.HTTPStatus, *apiErr.Code, apiErr.Message)
}
```

## HTTP client, headers, retries

`Configuration.HTTPClient` sends the requests: its `Transport` (an `http.RoundTripper`) is the hook for retries, logging, metrics and proxies. `Configuration.Headers` adds HTTP headers to every request.

```go
type retryTransport struct{ next http.RoundTripper }

func (t retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Println(req.Method, req.URL.Path)
	// requests are POST with a JSON body: GetBody gives it again for a retry
	return t.next.RoundTrip(req)
}

config := client.NewConfiguration("YOUR_API_KEY")
config.HTTPClient = &http.Client{Timeout: 30 * time.Second, Transport: retryTransport{http.DefaultTransport}}
config.Headers = map[string]string{"X-Request-Source": "billing"}
c := unione.NewClientWithConfig(config)
```

## Extra keys of a model

`HeadersObject` takes keys beyond its named fields (any `X-*` header, `CC`, `List-Unsubscribe`, ...) in its `Extra` map: they are sent next to the named fields, a named field wins over an `Extra` key with the same name; unknown keys of a response land in `Extra`.

```go
headers := &common.HeadersObject{
	XUNIONEGlobalLanguage: unione.String("en"),
	Extra: map[string]string{
		"X-MyHeader":       "some data",
		"CC":               "copy@example.com",
		"List-Unsubscribe": "<mailto:unsubscribe@example.com?subject=unsubscribe>",
	},
}
```

## Sections

| Section | Methods | Description |
|---|---|---|
| [Email](docs/Email.md) | `Send`, `Subscribe` | This group of methods is related with email sending. |
| [EmailValidation](docs/EmailValidation.md) | `Single` | Email validation methods. As a result of the check, you’ll get detailed information on email validity. It will help eliminate nonexistent emails, decrease the number of complaints, clean up the subscribers list and eventually enhance your sender reputation score. For example, you may add validation to a subscription form on your site to let users know they’ve just entered an invalid address. |
| [ConfirmedEmail](docs/ConfirmedEmail.md) | `Create`, `List`, `Delete` | Manage email addresses confirmed for sandbox-domain sending. These methods are available only with a main account API key. |
| [Template](docs/Template.md) | `Set`, `Get`, `List`, `Delete` | You have the option of storing a template on the UniOne side and then using it by id, which is highly recommended for sending repetitive messages or high volume traffic. UniOne supports two different [template engines](https://docs.unione.io/en/template-engines). |
| [Webhook](docs/Webhook.md) | `Set`, `Get`, `List`, `Delete` | UniOne lets you recieve notifications about email status changes or spam blocks. You provide an URL and every time a watched event occurs we will notify you by submitting a request with JSON data to this URL. |
| [Suppression](docs/Suppression.md) | `Set`, `Get`, `List`, `Delete` | Methods for managing suppression lists. An email can be added to suppression list because of unsubscription, bounce or complain. |
| [Domain](docs/Domain.md) | `GetDNSRecords`, `ValidateVerificationRecord`, `ValidateDKIM`, `List`, `GetSandboxDomain`, `Delete` | You can use these methods to setup sender domains and check their statuses. They might be useful in case you want to automate adding domains for your own clients. |
| [EventDump](docs/EventDump.md) | `Create`, `Get`, `List`, `Delete` | These methods are used to obtain a list of events for a given period of time. The data is stored for up to 45 days, depending on your tariff, i. e. you may fetch all event data starting from 00:00:00 31 days ago and up to current time. For custom tarifs with big sending volumes event dumps are available only for the last 7 days. |
| [Tag](docs/Tag.md) | `List`, `Delete` | This family of methods is used to manage user-defined tags. Tags are used to categorize letters based on your own criteria, e. g. to create a distinction between marketing and transactional letters or group together all notifications for a particular order type. UniOne allows you to assign tags to all outgoing emails, and the tags you’ve set may then be used to filter and analyze delivery statistics. |
| [Project](docs/Project.md) | `Create`, `Update`, `List`, `Delete` | Projects allow you to group several independent domains, webhooks, templates, and unsubscription lists together. It's very useful if you have your own clients and manage their data from a single UniOne account. Each project receives its own API key, and all API calls with that API key only deal with the data of that project. This way, you can assign a separate project for each client and be sure that they won't conflict; e.g. unsubscribing from one project won't affect another. |
| [System](docs/System.md) | `Ping`, `Info` | System operations. |
| [Obsolete](docs/Obsolete.md) | `UnsubscribedSet`, `UnsubscribedCheck`, `UnsubscribedList` | Obsolete methods dealing with unsubscribes. We recommend using the more flexible [suppression](https://docs.unione.io/en/web-api-ref#suppression) methods instead. Note that you can't resubscribe an unsubscribed recipient using WebAPI, but you can send him an email with subscription link using email/subscribe method. |
