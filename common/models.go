// Models of the Common section.

package common

import (
	"encoding/json"
	"strings"
)

// TemplateEngine — The [template engine](https://docs.unione.io/en/template-engines) for handling the
// substitutions(merge tags), "simple", "velocity", "liquid" or "none". Default value is "simple".
type TemplateEngine string

// TemplateEngine values.
const (
	TemplateEngineSimple   TemplateEngine = "simple"
	TemplateEngineVelocity TemplateEngine = "velocity"
	TemplateEngineLiquid   TemplateEngine = "liquid"
	TemplateEngineNone     TemplateEngine = "none"
)

// EmailStatus
type EmailStatus string

// EmailStatus values.
const (
	EmailStatusSent         EmailStatus = "sent"
	EmailStatusDelivered    EmailStatus = "delivered"
	EmailStatusOpened       EmailStatus = "opened"
	EmailStatusClicked      EmailStatus = "clicked"
	EmailStatusUnsubscribed EmailStatus = "unsubscribed"
	EmailStatusSubscribed   EmailStatus = "subscribed"
	EmailStatusSoftBounced  EmailStatus = "soft_bounced"
	EmailStatusHardBounced  EmailStatus = "hard_bounced"
	EmailStatusSpam         EmailStatus = "spam"
)

// VerificationRecordStatus — Only domains with "confirmed" verification record are allowed as sender domains.
type VerificationRecordStatus string

// VerificationRecordStatus values.
const (
	VerificationRecordStatusConfirmed   VerificationRecordStatus = "confirmed"
	VerificationRecordStatusUnconfirmed VerificationRecordStatus = "unconfirmed"
)

// EmailCounterMode — Sets the email counter mode. It allows you to choose whether the counter should reset automatically,
// or remain permanent. The following values are acceptable: 'default' - the counter resets to zero
// automatically at the beginning of each accounting period, 'permanent' - the counter never resets
// automatically. The reset can be done only manually by the user in the request with the passes 0
// value in the email_counter field. If the request doesn't contain the field or contains a null value,
// the limit value doesn't change.
type EmailCounterMode string

// EmailCounterMode values.
const (
	EmailCounterModeDefault   EmailCounterMode = "default"
	EmailCounterModePermanent EmailCounterMode = "permanent"
)

// BodyObject — Contains HTML/plaintext/AMP parts of the email. Either html or plaintext part is required.
type BodyObject struct {
	// HTML part of the email body
	HTML *string `json:"html,omitempty"`

	// Plaintext part of the email body.
	Plaintext *string `json:"plaintext,omitempty"`

	// Optional [AMP](http://amp.dev/documentation/guides-and-tutorials/learn/email_fundamentals/) part of
	// the email body.
	Amp *string `json:"amp,omitempty"`
}

// HeadersObject — Contains email headers, maximum 50. Only headers with “X-” name prefix are accepted, all other are
// ignored, for example X-UNIONE-Global-Language, X-UNIONE-Template-Engine. Standard headers "To,"
// "CC," and "BCC" are passed without the "X-." Yet, they are processed in a particular way and, as a
// result, have a number of restrictions. You can find more details about it [here](/cc-and-bcc). If
// our support have approved omitting standard unsubscription block for you, you can also pass
// List-Unsubscribe, List-Subscribe, List-Help, List-Owner, List-Archive, In-Reply-To and References
// headers.
//
// Keys beyond the named fields (any header, e.g. X-MyHeader, CC, List-Unsubscribe) go to Extra: they
// are sent next to the named fields, a named field wins over an Extra key with the same name; unknown
// keys of a response land in Extra.
type HeadersObject struct {
	// The language of the unsubscribe footer and unsubscribe page. Languages supported are: "be", "de",
	// "en", "es", "fr", "it", "pl", "pt", "ru", "ua", "kz".
	XUNIONEGlobalLanguage *string `json:"X-UNIONE-Global-Language,omitempty"`

	// The template engine for handling the substitutions(merge tags), “simple”, “velocity”, “liquid” or
	// “none”. Default value is “simple”. Has a priority over the value allowed in the “template_engine”
	// parameter.
	XUNIONETemplateEngine *string `json:"X-UNIONE-Template-Engine,omitempty"`

	// Extra holds the keys beyond the named fields, e.g. {"X-MyHeader": "some data"}.
	Extra map[string]string `json:"-"`
}

// headersObjectNamed are the JSON keys of the named fields of HeadersObject, lower case (encoding/json matches them so).
var headersObjectNamed = map[string]bool{
	"x-unione-global-language": true,
	"x-unione-template-engine": true,
}

// MarshalJSON writes the named fields and the Extra keys; a named field wins over an Extra key.
func (m HeadersObject) MarshalJSON() ([]byte, error) {
	type plain HeadersObject
	data, err := json.Marshal(plain(m))
	if err != nil || len(m.Extra) == 0 {
		return data, err
	}
	merged := map[string]json.RawMessage{}
	for key, value := range m.Extra {
		raw, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		merged[key] = raw
	}
	named := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &named); err != nil {
		return nil, err
	}
	for key, value := range named {
		merged[key] = value
	}
	return json.Marshal(merged)
}

// UnmarshalJSON reads the named fields; the other keys go to Extra.
func (m *HeadersObject) UnmarshalJSON(data []byte) error {
	type plain HeadersObject
	var fields plain
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	all := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &all); err != nil {
		return err
	}
	*m = HeadersObject(fields)
	m.Extra = nil
	for key, raw := range all {
		if headersObjectNamed[strings.ToLower(key)] {
			continue
		}
		var value string
		if json.Unmarshal(raw, &value) != nil {
			value = string(raw) // a number or an object: its JSON text
		}
		if m.Extra == nil {
			m.Extra = map[string]string{}
		}
		m.Extra[key] = value
	}
	return nil
}

// AttachmentsArrayItem
type AttachmentsArrayItem struct {
	// Attachment type, see [MIME](https://en.wikipedia.org/wiki/MIME). If unsure, use
	// "application/octet-stream".
	Type string `json:"type"`

	// Attachment name in the format: "name.extension". When multiple attachments are passed, their names
	// must be unique. The symbol ‘/’ is not allowed in attachment names.
	Name string `json:"name"`

	// File contents in [base64](https://en.wikipedia.org/wiki/Base64). Maximum file size 7MB (9786710
	// bytes in base64).
	Content string `json:"content"`
}

// InlineAttachmentsArrayItem
type InlineAttachmentsArrayItem struct {
	// Attachment type, see [MIME](https://en.wikipedia.org/wiki/MIME). If unsure, use
	// "application/octet-stream".
	Type string `json:"type"`

	// Attachment content id. For example, "name"="IMAGECID1" should be referenced in HTML like <img
	// src="cid:IMAGECID1">.
	Name string `json:"name"`

	// File contents in [base64](https://en.wikipedia.org/wiki/Base64). Maximum file size 7MB (9786710
	// bytes in base64).
	Content string `json:"content"`
}

// SuccessResponse
type SuccessResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`
}
