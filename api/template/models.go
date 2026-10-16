// Models of the Template section.

package template

import "github.com/unione-repo/unione-go-sdk/common"

// EditorType — Editor type, "html" or "visual". Default is "html". You can create a template with "visual" editor
// type only in web interface.
type EditorType string

// EditorType values.
const (
	EditorTypeHTML   EditorType = "html"
	EditorTypeVisual EditorType = "visual"
)

// TemplateSetRequest — Object with all the properties of the template.
type TemplateSetRequest struct {
	// Object with all the template properties.
	Template *TemplateObject `json:"template"`
}

// TemplateObject — Object with all the template properties.
type TemplateObject struct {
	// Template unique identifier. If you omit the "id" field then new template will be created. If you
	// specify the "id" field, the system will try to find a template with the specified id and edit it.
	ID *string `json:"id,omitempty"`

	// Template name.
	Name string `json:"name"`

	// Editor type, "html" or "visual". Default is "html". You can create a template with "visual" editor
	// type only in web interface.
	EditorType EditorType `json:"editor_type,omitempty"`

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
	Body *common.BodyObject `json:"body,omitempty"`

	// Email subject.
	Subject *string `json:"subject,omitempty"`

	// Sender's email.
	FromEmail string `json:"from_email"`

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
}

// TemplateSetResponse
type TemplateSetResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Object with all the template properties.
	Template *TemplateObject `json:"template,omitempty"`
}

// TemplateGetRequest
type TemplateGetRequest struct {
	// Template id.
	ID string `json:"id"`
}

// TemplateGetResponse
type TemplateGetResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Object with all the template properties.
	Template *TemplateObject `json:"template,omitempty"`
}

// TemplateListRequest
type TemplateListRequest struct {
	// Limits the number of templates returned, default is 50.
	// Default: `50`.
	Limit *int `json:"limit,omitempty"`

	// Index of the first template in the list, starting with 0.
	// Default: `0`.
	Offset *int `json:"offset,omitempty"`
}

// TemplateListResponse
type TemplateListResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Array of template objects.
	Templates []TemplateListObject `json:"templates,omitempty"`
}

// TemplateListObject
type TemplateListObject struct {
	// Template unique identifier.
	ID string `json:"id,omitempty"`

	// Template name.
	Name string `json:"name,omitempty"`

	// Editor type, "html" or "visual". Default is "html". You can create a template with "visual" editor
	// type only in web interface.
	EditorType EditorType `json:"editor_type,omitempty"`

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
	Body *common.BodyObject `json:"body,omitempty"`

	// Email subject.
	Subject string `json:"subject,omitempty"`

	// Sender's email.
	FromEmail string `json:"from_email,omitempty"`

	// Sender's name.
	FromName string `json:"from_name,omitempty"`

	// Optional Reply-To email (in case it's different to sender's email)
	ReplyTo string `json:"reply_to,omitempty"`

	// Optional Reply-To name (if reply_to email is specified and you want to display not only this email
	// but also the name)
	ReplyToName string `json:"reply_to_name,omitempty"`

	// 1 = click tracking is on (default), 0 = click tracking is off. To use track_links = 0, you need to
	// ask support to enable this feature.
	// Allowed values: `0`, `1`.
	TrackLinks int `json:"track_links,omitempty"`

	// 1 = read tracking is on (default), 0=read tracking is off. To use track_read = 0, you need to ask
	// support to enable this feature.
	// Allowed values: `0`, `1`.
	TrackRead int `json:"track_read,omitempty"`

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

	// Template creation date and time in UTC timezone in format "YYYY-MM-DD hh:mm:ss"
	Created string `json:"created,omitempty"`

	// Unique user identifier.
	UserID int `json:"user_id,omitempty"`

	// Unqiue project identifier, ASCII string up to 36 characters long.
	ProjectID string `json:"project_id,omitempty"`

	// Project name, unique for user account.
	ProjectName string `json:"project_name,omitempty"`
}

// TemplateDeleteRequest
type TemplateDeleteRequest struct {
	// Template id.
	ID string `json:"id"`
}
