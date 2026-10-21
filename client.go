// Package unione is the UniOne Web API client.
package unione

import (
	"github.com/unione-repo/unione-go-sdk/api/confirmedemail"
	"github.com/unione-repo/unione-go-sdk/api/domain"
	"github.com/unione-repo/unione-go-sdk/api/email"
	"github.com/unione-repo/unione-go-sdk/api/emailvalidation"
	"github.com/unione-repo/unione-go-sdk/api/eventdump"
	"github.com/unione-repo/unione-go-sdk/api/obsolete"
	"github.com/unione-repo/unione-go-sdk/api/project"
	"github.com/unione-repo/unione-go-sdk/api/suppression"
	"github.com/unione-repo/unione-go-sdk/api/system"
	"github.com/unione-repo/unione-go-sdk/api/tag"
	"github.com/unione-repo/unione-go-sdk/api/template"
	"github.com/unione-repo/unione-go-sdk/api/webhook"
	"github.com/unione-repo/unione-go-sdk/client"
)

// Client is the entry point of the UniOne API: one field per section.
type Client struct {
	// APIClient is the low level transport.
	APIClient *client.APIClient

	// Email — This group of methods is related with email sending.
	Email *email.EmailAPI

	// EmailValidation — Email validation methods. As a result of the check, you’ll get detailed information on email
	// validity. It will help eliminate nonexistent emails, decrease the number of complaints, clean up the
	// subscribers list and eventually enhance your sender reputation score. For example, you may add
	// validation to a subscription form on your site to let users know they’ve just entered an invalid
	// address.
	EmailValidation *emailvalidation.EmailValidationAPI

	// ConfirmedEmail — Manage email addresses confirmed for sandbox-domain sending. These methods are available only with a
	// main account API key.
	ConfirmedEmail *confirmedemail.ConfirmedEmailAPI

	// Template — You have the option of storing a template on the UniOne side and then using it by id, which is
	// highly recommended for sending repetitive messages or high volume traffic. UniOne supports two
	// different [template engines](https://docs.unione.io/en/template-engines).
	Template *template.TemplateAPI

	// Webhook — UniOne lets you recieve notifications about email status changes or spam blocks. You provide an URL
	// and every time a watched event occurs we will notify you by submitting a request with JSON data to
	// this URL.
	Webhook *webhook.WebhookAPI

	// Suppression — Methods for managing suppression lists. An email can be added to suppression list because of
	// unsubscription, bounce or complain.
	Suppression *suppression.SuppressionAPI

	// Domain — You can use these methods to setup sender domains and check their statuses. They might be useful in
	// case you want to automate adding domains for your own clients.
	Domain *domain.DomainAPI

	// EventDump — These methods are used to obtain a list of events for a given period of time. The data is stored for
	// up to 45 days, depending on your tariff, i. e. you may fetch all event data starting from 00:00:00
	// 31 days ago and up to current time. For custom tarifs with big sending volumes event dumps are
	// available only for the last 7 days.
	EventDump *eventdump.EventDumpAPI

	// Tag — This family of methods is used to manage user-defined tags. Tags are used to categorize letters
	// based on your own criteria, e. g. to create a distinction between marketing and transactional
	// letters or group together all notifications for a particular order type. UniOne allows you to assign
	// tags to all outgoing emails, and the tags you’ve set may then be used to filter and analyze delivery
	// statistics.
	Tag *tag.TagAPI

	// Project — Projects allow you to group several independent domains, webhooks, templates, and unsubscription
	// lists together. It's very useful if you have your own clients and manage their data from a single
	// UniOne account. Each project receives its own API key, and all API calls with that API key only deal
	// with the data of that project. This way, you can assign a separate project for each client and be
	// sure that they won't conflict; e.g. unsubscribing from one project won't affect another.
	Project *project.ProjectAPI

	// System — System operations.
	System *system.SystemAPI

	// Obsolete — Obsolete methods dealing with unsubscribes. We recommend using the more flexible
	// [suppression](https://docs.unione.io/en/web-api-ref#suppression) methods instead. Note that you
	// can't resubscribe an unsubscribed recipient using WebAPI, but you can send him an email with
	// subscription link using email/subscribe method.
	Obsolete *obsolete.ObsoleteAPI
}

// API hosts for NewClient, e.g. NewClient("YOUR_API_KEY", HostUS1).
const (
	HostAPI = client.HostAPI
	HostEU1 = client.HostEU1
	HostUS1 = client.HostUS1
)

// NewClient returns a client with the default configuration. The optional host is an API host
// (HostUS1, "us1.unione.io") or a full base URL.
func NewClient(apiKey string, host ...string) *Client {
	config := client.NewConfiguration(apiKey)
	if len(host) > 0 && host[0] != "" {
		config.BaseURL = client.BaseURLFor(host[0])
	}
	return NewClientWithConfig(config)
}

// NewClientWithConfig returns a client with a custom configuration.
func NewClientWithConfig(config *client.Configuration) *Client {
	c := client.NewAPIClient(config)
	return &Client{
		APIClient:       c,
		Email:           email.NewEmailAPI(c),
		EmailValidation: emailvalidation.NewEmailValidationAPI(c),
		ConfirmedEmail:  confirmedemail.NewConfirmedEmailAPI(c),
		Template:        template.NewTemplateAPI(c),
		Webhook:         webhook.NewWebhookAPI(c),
		Suppression:     suppression.NewSuppressionAPI(c),
		Domain:          domain.NewDomainAPI(c),
		EventDump:       eventdump.NewEventDumpAPI(c),
		Tag:             tag.NewTagAPI(c),
		Project:         project.NewProjectAPI(c),
		System:          system.NewSystemAPI(c),
		Obsolete:        obsolete.NewObsoleteAPI(c),
	}
}

// APIError is the error of every method (client.APIError): errors.As(err, &apiErr).
type APIError = client.APIError

// String returns a pointer to v, for optional fields of the requests.
func String(v string) *string {
	return &v
}

// Int returns a pointer to v, for optional fields of the requests.
func Int(v int) *int {
	return &v
}

// Float64 returns a pointer to v, for optional fields of the requests.
func Float64(v float64) *float64 {
	return &v
}

// Bool returns a pointer to v, for optional fields of the requests.
func Bool(v bool) *bool {
	return &v
}
