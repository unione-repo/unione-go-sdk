# UniOne Go SDK

[![Go](https://pkg.go.dev/badge/github.com/unione-repo/unione-go-sdk.svg)](https://pkg.go.dev/github.com/unione-repo/unione-go-sdk) [![License: MIT](https://img.shields.io/badge/license-MIT-blue)](LICENSE)

Official Go library for the [UniOne](https://unione.io) transactional email API.

[API reference](https://docs.unione.io/en/web-api-ref) · Docs for LLMs: [llms.txt](https://docs.unione.io/llms.txt)

> **On Go net/http?** Install [`github.com/unione-repo/unione-nethttp`](https://github.com/unione-repo/unione-nethttp) instead: this SDK plus configuration, a mailer and a webhook route.

## Install

```bash
go get github.com/unione-repo/unione-go-sdk
```

Requirements: Go 1.18+. No third-party dependencies.

Get an API key in the [control panel](https://cp.unione.io/en/settings/security/api) and keep it in the environment, never in code:

```bash
export UNIONE_API_KEY="your-api-key"
```

## Send an email

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/email"
	"github.com/unione-repo/unione-go-sdk/common"
)

func main() {
	client := unione.NewClient(os.Getenv("UNIONE_API_KEY"))
	ctx := context.Background()

	response, err := client.Email.Send(ctx, &email.EmailSendRequest{
		Message: &email.MessageObject{
			Recipients: []email.MessageObjectRecipientsItem{email.RecipientFromAddress("Alex <alex@example.com>")},
			FromEmail:  unione.String("shop@example.com"),
			FromName:   unione.String("Shop"),
			Subject:    "Your order",
			Body:       common.HTMLBody("<p>Hello, {{to_name}}! Your order has shipped.</p>"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.JobID)
}
```

The sender address must be on a domain verified in your account. No domain yet? Send from the [sandbox domain](#no-domain-yet-the-test-sandbox-domain) in three steps.

## Contents

- [Recipes](#recipes): personalization · attachments · CC and BCC · scheduling · safe retries · errors · webhooks
- [No domain yet: the test sandbox domain](#no-domain-yet-the-test-sandbox-domain)
- [Going to production](#going-to-production)
- [Reference](#reference)
- [Help](#help)

## Recipes

Every recipe continues the example above: the client and its imports are there.

### Personalize for each recipient

Up to 500 recipients per call, each gets a separate copy with its own substitutions. Metadata of a recipient comes back in its webhook events.

```go
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients: []email.MessageObjectRecipientsItem{
			email.RecipientFromAddress("Alex <alex@example.com>", map[string]string{"order": "1001"}).
				WithMetadata(map[string]string{"user_id": "42"}),
			email.RecipientFromAddress("Maria <maria@example.com>", map[string]string{"order": "1002"}).
				WithMetadata(map[string]string{"user_id": "43"}),
		},
		FromEmail: unione.String("shop@example.com"),
		Subject:   "Order #{{order}}",
		Body:      common.HTMLBody("<p>{{to_name}}, order #{{order}} has shipped.</p>"),
	},
}); err != nil {
	log.Fatal(err)
}
```

### Attach a file

```go
import (
	"encoding/base64"
	"os"
)

content, err := os.ReadFile("invoice.pdf")
if err != nil {
	log.Fatal(err)
}
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients: []email.MessageObjectRecipientsItem{email.RecipientFromAddress("alex@example.com")},
		FromEmail:  unione.String("shop@example.com"),
		Subject:    "Invoice",
		Body:       common.HTMLBody("<p>Your invoice is attached.</p>"),
		Attachments: []common.AttachmentsArrayItem{
			{Type: "application/pdf", Name: "invoice.pdf", Content: base64.StdEncoding.EncodeToString(content)},
		},
	},
}); err != nil {
	log.Fatal(err)
}
```

### CC and BCC

Every address goes to `recipients` (each gets a copy); the visible ones also go to the headers. `audit@example.com` is in neither header: a BCC. The rules and limits: [CC and BCC](https://docs.unione.io/en/cc-and-bcc).

```go
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients: []email.MessageObjectRecipientsItem{
			email.RecipientFromAddress("alex@example.com"),
			email.RecipientFromAddress("boss@example.com"),
			email.RecipientFromAddress("audit@example.com"),
		},
		FromEmail: unione.String("shop@example.com"),
		Subject:   "Weekly report",
		Body:      common.HTMLBody("<p>The report is attached.</p>"),
		Headers: &common.HeadersObject{
			Extra: map[string]string{"To": "alex@example.com", "CC": "boss@example.com"},
		},
	},
}); err != nil {
	log.Fatal(err)
}
```

### Schedule

`send_at` is UTC, up to 24 hours ahead.

```go
import "time"

sendAt := time.Now().UTC().Add(time.Hour).Format("2006-01-02 15:04:05") // UTC, "YYYY-MM-DD hh:mm:ss"
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients: []email.MessageObjectRecipientsItem{email.RecipientFromAddress("alex@example.com")},
		FromEmail:  unione.String("shop@example.com"),
		Subject:    "Reminder",
		Body:       common.TextBody("See you tomorrow!"),
		Options:    &email.MessageObjectOptions{SendAt: unione.String(sendAt)},
	},
}); err != nil {
	log.Fatal(err)
}
```

### Retry safely

A repeat with the same `idempotence_key` within a minute is not sent twice: retry network errors, timeouts, 429 and 5xx with it.

```go
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients:     []email.MessageObjectRecipientsItem{email.RecipientFromAddress("alex@example.com")},
		FromEmail:      unione.String("shop@example.com"),
		Subject:        "Reset your password",
		Body:           common.HTMLBody("<p>Use this link to reset your password.</p>"),
		IdempotenceKey: unione.String("password-reset-8f3a"),
	},
}); err != nil {
	log.Fatal(err)
}
```

### Handle errors

Recipients the API drops (unsubscribed, invalid, suppressed) are not an error: they are in `failed_emails` of the response, the rest still get the message.

```go
import "errors"

if _, err := client.Email.Send(ctx, request); err != nil {
	var apiErr *unione.APIError
	if errors.As(err, &apiErr) {
		// HTTP status (0: network error or timeout), API error code, message
		log.Println(apiErr.HTTPStatus, apiErr.ErrorCode(), apiErr.Message)
	}
}
```

### Receive delivery events

Register your HTTPS endpoint once (`webhook/set`). Then check the signature of every notification (`auth`: MD5 of the raw body with the `auth` value replaced by your API key; a wrong one gets 401) and process the events.

```go
import (
	"crypto/md5"
	"net/http"
	"os"
	"strings"

	"github.com/unione-repo/unione-go-sdk/api/webhook"
)

// handleWebhook: the body of the POST request as received; returns the HTTP status to answer
func handleWebhook(body []byte) int {
	callback, err := webhook.ParseCallback(body)
	if err != nil {
		return http.StatusBadRequest
	}
	// auth: MD5 of the raw body with the auth value replaced by your API key
	auth := callback.Auth
	signed := strings.Replace(string(body), auth, os.Getenv("UNIONE_API_KEY"), 1)
	if auth == "" || fmt.Sprintf("%x", md5.Sum([]byte(signed))) != auth {
		return http.StatusUnauthorized
	}

	for _, user := range callback.EventsByUser {
		for _, event := range user.Events {
			// email, status (delivered, opened, hard_bounced, spam, ...), job_id, metadata
			log.Printf("%+v", event.EventData)
		}
	}
	return http.StatusOK
}
```

> **Answer 200 within 3 seconds** and process the events in the background (a queue): otherwise UniOne repeats the request every 10 minutes for 24 hours (`retry_count` grows), and a handler that failed every request for at least 10 events of a day is turned inactive. The handler must take at least **5 parallel requests of up to 100 events each**; deduplicate, a repeat delivers the same events again.

## No domain yet: the test sandbox domain

Until your domain is verified, send from the **sandbox domain**: a test domain every account has, already verified, with DKIM and SPF set. It delivers only to addresses you confirm (the confirmation link comes by email, valid 48 hours), up to 100 emails a day. The sandbox domain is deleted 31 days after it was created, but you can get it again: the next `domain/get-sandbox-domain` call creates a new one (with a new name, so read it from the call, do not hardcode it).

```go
import "github.com/unione-repo/unione-go-sdk/api/confirmedemail"

// 1. Your sandbox domain: created on the first call, the same one afterwards
sandbox, err := client.Domain.GetSandboxDomain(ctx)
if err != nil {
	log.Fatal(err)
}

// 2. Confirm your inbox: UniOne sends a link, click it within 48 hours
confirm := &confirmedemail.ConfirmedEmailRequest{Email: "you@example.com"}
if _, err := client.ConfirmedEmail.Create(ctx, confirm); err != nil {
	log.Fatal(err)
}

// 3. Send from any address on the sandbox domain to the confirmed one
if _, err := client.Email.Send(ctx, &email.EmailSendRequest{
	Message: &email.MessageObject{
		Recipients: []email.MessageObjectRecipientsItem{email.RecipientFromAddress("you@example.com")},
		FromEmail:  unione.String("noreply@" + sandbox.Domain),
		Subject:    "Hello from UniOne",
		Body:       common.HTMLBody("<p>It works!</p>"),
	},
}); err != nil {
	log.Fatal(err)
}
```

<details>
<summary>Sandbox limits</summary>

- Sends only to confirmed addresses: up to 10 per account.
- Up to 100 emails a day.
- The sandbox domain expires after 31 days; call `domain/get-sandbox-domain` again to get a new one.
- Works only with the account API key, not a project key.
- Not listed in `domain/list`; the other `domain/*` methods do not apply to it.

The sandbox is for development and testing. Do not use a sandbox domain as the sender in production.

</details>

<details>
<summary>The same with curl</summary>

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

</details>

## Going to production

**1. Verify your sending domain.** `domain/get-dns-records` returns the verification record, the DKIM key and the DMARC target; add the records at your DNS provider, then `domain/validate-verification-record` and `domain/validate-dkim`.

| Type | Host | Value | Where it comes from |
|---|---|---|---|
| TXT | `@` | `unione-validate-hash=…` | `verification-record` of the response, as is |
| TXT | `<selector>._domainkey` (the selector is on the domain page of the control panel, e.g. `us`) | `v=DKIM1; k=rsa; p=<dkim>` | `dkim` of the response after `v=DKIM1; k=rsa; p=` |
| TXT | `@` | `v=spf1 include:spf.unione.io ~all` | not in the response; add the include to your SPF record if you have one |
| CNAME | `_dmarc` | `<domain>.dmarc.…` | `dmarc` of the response (UniOne publishes the DMARC policy); keep your own DMARC TXT record instead if the domain has one |

**2. Make sending reliable.**

| | |
|---|---|
| Retries | `429` and `5xx` with backoff and the same `idempotence_key`; fix other errors instead of retrying |
| Limits | 500 recipients and 10 MB per request, attachments included: base64 makes a file about a third larger |
| Body | always `html` or `plaintext` (a template can supply it) |
| Delivery status | a webhook, not polling |
| Queue | send from a background job, not inside the web request |
| Config | `from_email` from config, so the sandbox → production switch needs no code change |
| Defaults | keep the unsubscribe footer and tracking; changing them needs UniOne support |

## Reference

### Sections

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

### API host

The second argument selects the API host (datacenter) or a full base URL.

```go
client := unione.NewClient("YOUR_API_KEY", "us1.unione.io")
client = unione.NewClient("YOUR_API_KEY", unione.HostUS1)
```

### Error handling

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

### HTTP client, headers, retries

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

### Extra keys of a model

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

## Help

- [Integration guides](https://docs.unione.io/en/integrations)
- AI coding agents: [llms.txt](https://docs.unione.io/llms.txt), the [MCP server](https://unione.io/en/agents), skills: `npx skills add https://github.com/unione-repo/unione-skills.git`
- [Support](https://cp.unione.io/en/support) · [Issues](https://github.com/unione-repo/unione-go-sdk/issues)
- Contributing: [AGENTS.md](AGENTS.md) — the SDK is generated from the OpenAPI spec.
