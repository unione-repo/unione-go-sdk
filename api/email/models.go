// Models of the Email section.

package email

import "github.com/unione-repo/unione-go-sdk/common"

// EmailSendRequest
type EmailSendRequest struct {
	// Object containing all message properties.
	Message *MessageObject `json:"message"`
}

// MessageObject — Object containing all message properties.
type MessageObject struct {
	// Contains recipient emails, substitutions(merge tags) and metadata.
	Recipients []MessageObjectRecipientsItem `json:"recipients"`

	// Optional identifier of the template that had been created by
	// [template/set](https://docs.unione.io/en/web-api-ref#template-set) method. If template_id is passed
	// then fields from the template are used instead of missed email/send parameters. E.g. if "body"
	// parameter in email/send is omitted, "body" from template will be used. If "subject" in email/send is
	// omitted, "subject" from template will be used.
	TemplateID *string `json:"template_id,omitempty"`

	// An array of strings. The maximum string length is 100 characters. You can include up to 4 strings
	// which must be unique. No more than 10 000 tags are allowed per project; if you try to add more,
	// email/send will return an error. You may use tags to categorize emails by your own criteria, and
	// they will be passed along by event-dump methods.
	Tags []string `json:"tags,omitempty"`

	// Whether to skip or not appending default unsubscribe footer. 1=skip, 0=append, default is 0. You
	// should [ask support](https://cp.unione.io/en/support) to approve
	// usage of skip_unsubscribe=1. [Read more](https://docs.unione.io/en/unsubscribe-link).
	// Allowed values: `0`, `1`.
	SkipUnsubscribe *int `json:"skip_unsubscribe,omitempty"`

	// The language of the unsubscribe footer and unsubscribe page. Languages supported are: "be", "de",
	// "en", "es", "fr", "it", "pl", "pt", "ru", "ua", "kz".
	GlobalLanguage *string `json:"global_language,omitempty"`

	// The [template engine](https://docs.unione.io/en/template-engines) for handling the
	// substitutions(merge tags), "simple", "velocity", "liquid" or "none". Default value is "simple".
	TemplateEngine common.TemplateEngine `json:"template_engine,omitempty"`

	// Object for passing the substitutions(merge tags) common for all recipients - e.g., company name. If
	// the substitution names are duplicated in object "substitutions", the values of the variables will be
	// taken from the object "substitutions". The substitutions can be used in the following parameters:
	//
	// - body.html
	// - body.plaintext
	// - body.amp
	// - subject
	// - from_name
	// - options.unsubscribe_url
	// A substitution name can consist from latin characeters, numbers and "_" symbol, and should start
	// with the letter.
	GlobalSubstitutions map[string]string `json:"global_substitutions,omitempty"`

	// Object for passing the metadata common for all the recipients, such as "key": "value". Max key
	// quantity: 10. Max key length: 64 symbols. Max value length: 1024 symbols. The metadata is returned
	// by [webhook](https://docs.unione.io/en/web-api-ref#webhook). To group single emails into a campaign,
	// you can use the system key 'campaign_id' with a value in the form of a non-negative decimal integer
	// up to 128 bits or a UUID in the format 'c7703772-453e-4c77-a9ae-9a1b94051b5a'. Incorrect values will
	// be treated as '0'.
	GlobalMetadata map[string]string `json:"global_metadata,omitempty"`

	// Contains HTML/plaintext/AMP parts of the email. Either html or plaintext part is required.
	Body *common.BodyObject `json:"body"`

	// Email subject.
	Subject string `json:"subject"`

	// Sender's email. Required only if template_id parameter is empty.
	FromEmail *string `json:"from_email,omitempty"`

	// Sender's name.
	FromName *string `json:"from_name,omitempty"`

	// Optional Reply-To email (in case it's different to sender's email)
	ReplyTo *string `json:"reply_to,omitempty"`

	// Optional Reply-To name (if reply_to email is specified and you want to display not only this email
	// but also the name)
	ReplyToName *string `json:"reply_to_name,omitempty"`

	// 1 = click tracking is on (default), 0 = click tracking is off. To use track_links = 0, you need to
	// ask support to enable this feature.
	// Allowed values: `0`, `1`.
	TrackLinks *int `json:"track_links,omitempty"`

	// 1 = read tracking is on (default), 0=read tracking is off. To use track_read = 0, you need to ask
	// support to enable this feature.
	// Allowed values: `0`, `1`.
	TrackRead *int `json:"track_read,omitempty"`

	// (allowed values 0 or 1, default 0) If set to 1, the global unavailability list will be ignored. Even
	// if the address was found to be unreachable while sending other UniOne users' emails, or its owner
	// has issued complaints, the message will still be sent. The setting may be ignored for certain
	// addresses.
	// Allowed values: `0`, `1`.
	BypassGlobal *int `json:"bypass_global,omitempty"`

	// (allowed values 0 or 1, default 0) If set to 1, the current user's or project's unavailability list
	// will be ignored. Works only if bypass_global is set to 1.
	// Allowed values: `0`, `1`.
	BypassUnavailable *int `json:"bypass_unavailable,omitempty"`

	// (allowed values 0 or 1, default 0) If set to 1, the current list of unsubscribed addresses for this
	// account or project will be ignored. Works only if bypass_global is set to 1. The setting is
	// available only for users that have been granted the right to omit the unsubscribe link (to request,
	// please contact [support](https://cp.unione.io/en/support)).
	// Allowed values: `0`, `1`.
	BypassUnsubscribed *int `json:"bypass_unsubscribed,omitempty"`

	// (allowed values 0 or 1, default 0) If set to 1, the user's or project's complaint list will be
	// ignored. Works only if bypass_global is set to 1. The setting is available only for users that have
	// been granted the right to omit the unsubscribe link (to request, please contact
	// [support](https://cp.unione.io/en/support)).
	// Allowed values: `0`, `1`.
	BypassComplained *int `json:"bypass_complained,omitempty"`

	// A string of up to 64 characters containing a unique message key. This can be used to prevent
	// occasional message duplicates. If you send another API request with the same message key within the
	// next minute, it will be declined. We can generate a message key for each letter automatically; to
	// enable this option, please contact our [tech support](https://cp.unione.io/en/support).
	IdempotenceKey *string `json:"idempotence_key,omitempty"`

	// Contains email headers, maximum 50. Only headers with “X-” name prefix are accepted, all other are
	// ignored, for example X-UNIONE-Global-Language, X-UNIONE-Template-Engine. Standard headers "To,"
	// "CC," and "BCC" are passed without the "X-." Yet, they are processed in a particular way and, as a
	// result, have a number of restrictions. You can find more details about it [here](/cc-and-bcc). If
	// our support have approved omitting standard unsubscription block for you, you can also pass
	// List-Unsubscribe, List-Subscribe, List-Help, List-Owner, List-Archive, In-Reply-To and References
	// headers.
	Headers *common.HeadersObject `json:"headers,omitempty"`

	// Optional array of attachments
	Attachments []common.AttachmentsArrayItem `json:"attachments,omitempty"`

	// Optional array of inline attachments, e.g. for including images in email instead of loading them
	// from external URL.
	InlineAttachments []common.InlineAttachmentsArrayItem `json:"inline_attachments,omitempty"`

	// Additional message options.
	Options *MessageObjectOptions `json:"options,omitempty"`
}

// MessageObjectRecipientsItem
type MessageObjectRecipientsItem struct {
	// Recipient email
	Email string `json:"email"`

	// Object to pass the substitutions(merge tags) for the recipient (e.g. recipient name, discount code,
	// password change link, etc. See [Template engines](https://docs.unione.io/en/template-engines)). The
	// substitutions can be used in the following parameters:
	//
	// - body.html
	// - body.plaintext
	// - body.amp
	// - subject
	// - from_name
	// - headers["List-Unsubscribe"]
	// - options.unsubscribe_url
	// A substitution name can consist from latin characters, numbers and "_" symbol, and should start with
	// the letter. There's a special substitution "to_name" which is used to put recipent's name like "Name
	// Surname" to include it to SMTP header "To" in the form "Name Surname <email@example.com>". The
	// "to_name" length is limited to 78 symbols.
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Object for passing the metadata, such as "key": "value". Max key quantity: 10. Max key length: 64
	// symbols. Max value length: 1024 symbols. The metadata is returned by
	// [webhook](https://docs.unione.io/en/web-api-ref#webhook) and
	// [event-dump](https://docs.unione.io/en/web-api-ref#event-dump) methods. If you pass strign up to 128
	// bit with "campaign_id" field name (name is configurable thru support), it will be considered as
	// campaign identifier in statistics.
	Metadata map[string]string `json:"metadata,omitempty"`
}

// MessageObjectOptions — Additional message options.
type MessageObjectOptions struct {
	// Date and time in "YYYY-MM-DD hh:mm:ss" format in the UTC timezone. Allows schedule sending up to 24
	// hours in advance.
	SendAt *string `json:"send_at,omitempty"`

	// Custom unsubscribe link. Read more [here](https://docs.unione.io/en/unsubscribe-link).
	UnsubscribeURL *string `json:"unsubscribe_url,omitempty"`

	// Backend-domain identifier to send message with. If custom_backend_id is not set, default one for
	// your account/project will be used. You can pay for one or more dedicated IPs and use this option.
	CustomBackendID *int `json:"custom_backend_id,omitempty"`

	// SMTP pool identifier to send message with. Usually you don't need to pass this parameter because
	// correct smtp_pool_id is selected automatically based on the default one for your account/project or
	// according to custom_backend_id parameter.
	SMTPPoolID *string `json:"smtp_pool_id,omitempty"`
}

// EmailSendResponse
type EmailSendResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Job identifier, might be useful for investigating errors.
	JobID string `json:"job_id,omitempty"`

	// Array of recipients emails successfully accepted for sending.
	Emails []string `json:"emails,omitempty"`

	// Object with emails rejected for sending as property names and their statuses as property values,
	// e.g.: {"email1@gmail.com": "temporary_unavailable"}. Possible statuses:
	//
	// - __unsubscribed__ - specified email is unsubscribed;
	// - __invalid__ - the email does not exist or has been entered incorrectly;
	// - __duplicate__ - the email already exists in the request, email duplicating is prevented;
	// - __temporary_unavailable__ - the email address is unavailable. This means that over the next three days sending to this address will return an error. Email may be temporarily unavailable due to several reasons, e.g.:
	//
	// - a previous email has been rejected by the recipient's server for spam;
	// - the recipient's mailbox is full or is not used;
	// - recipient's domain does not accept mail;
	// - sending server was rejected due to blacklisting;
	//
	// - __permanent_unavailable__ - the email address is permanently unavailable or unsubscribed globally.
	// - __complained__ -  the recipient reported spam in the previous emails;
	// - __blocked__ -  sending to the email is prohibited by administration of UniOne.
	// We may added new statuses in the future.
	FailedEmails map[string]string `json:"failed_emails,omitempty"`
}

// EmailSubscribeRequest
type EmailSubscribeRequest struct {
	// Sender’s email on a confirmed domain. (Field alias "email_address_from" is supported, too).
	FromEmail string `json:"from_email"`

	// Sender’s name. (Field alias "name_from" is supported, too).
	FromName string `json:"from_name"`

	// Recipient’s email. (Field alias "email_address_to" is supported, too).
	ToEmail string `json:"to_email"`
}
