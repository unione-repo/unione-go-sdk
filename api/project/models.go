// Models of the Project section.

package project

import "github.com/unione-repo/unione-go-sdk/common"

// ProjectCreateRequest
type ProjectCreateRequest struct {
	// Object with project properties.
	Project *ProjectObject `json:"project"`
}

// ProjectObject — Object with project properties.
type ProjectObject struct {
	// Project name, unique for user account.
	Name string `json:"name"`

	// ISO-3166 alpha-2 country code. If set, UniOne treats project personal data according to country
	// laws, e.g. GDPR for european countries.
	Country *string `json:"country,omitempty"`

	// Whether email sending is enabled for this project or not.
	// Default: `True`.
	SendEnabled *bool `json:"send_enabled,omitempty"`

	// If it's false then UniOne adds default unsubscribe footer to each email sent with API key of the
	// project. If it's true then appending default unsubscribe footer for this project is avoided and
	// sending with custom unsubscribe url or even without unsubscribe url is permitted for the project.
	// Please note that custom_unsubscribe_url_enabled=true is available only if removing unsubscribe link
	// is approved for the user account by support. More on this
	// [here](https://docs.unione.io/en/unsubscribe-link). When custom_unsubscribe_url_enabled is skipped
	// on creating a project, it's value is taken from user
	CustomUnsubscribeURLEnabled *bool `json:"custom_unsubscribe_url_enabled,omitempty"`

	// A unique domain identifier which will determine the [tracking
	// domain](https://docs.unione.io/en/tracking-domains) or [dedicated IP](/dedicated-ip) pool to be used
	// by default. If the value is not specified, a default backend domain will be assigned for your
	// account or project instead.
	BackendDomainID *int `json:"backend_domain_id,omitempty"`

	// A unique identifier that defines the default unsubscribe page. If absent, the account/project will
	// be assigned the system default unsubscribe page.
	UnsubscribePageID *int `json:"unsubscribe_page_id,omitempty"`
}

// ProjectCreateResponse
type ProjectCreateResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Unqiue project identifier, ASCII string up to 36 characters long.
	ProjectID string `json:"project_id,omitempty"`

	// API key of the project. You can use it instead of the user API key parameter in all methods except
	// project/* methods.
	ProjectAPIKey string `json:"project_api_key,omitempty"`
}

// ProjectUpdateRequest
type ProjectUpdateRequest struct {
	// Unique project identifier. You should identify project to update either with project_id or with
	// project_api_key field. But we recommend usage of project_id as a more secure option.
	ProjectID *string `json:"project_id,omitempty"`

	// Optional project API key. You should identify project to update either with project_id or with
	// project_api_key field. But we recommend usage of project_id as a more secure option.
	ProjectAPIKey *string `json:"project_api_key,omitempty"`

	// Object with project properties.
	Project *ProjectUpdateObject `json:"project"`
}

// ProjectUpdateObject — Object with project properties.
type ProjectUpdateObject struct {
	// Project name, unique for user account.
	Name string `json:"name"`

	// ISO-3166 alpha-2 country code. If set, UniOne treats project personal data according to country
	// laws, e.g. GDPR for european countries.
	Country *string `json:"country,omitempty"`

	// Whether email sending is enabled for this project or not.
	// Default: `True`.
	SendEnabled *bool `json:"send_enabled,omitempty"`

	// If it's false then UniOne adds default unsubscribe footer to each email sent with API key of the
	// project. If it's true then appending default unsubscribe footer for this project is avoided and
	// sending with custom unsubscribe url or even without unsubscribe url is permitted for the project.
	// Please note that custom_unsubscribe_url_enabled=true is available only if removing unsubscribe link
	// is approved for the user account by support. More on this
	// [here](https://docs.unione.io/en/unsubscribe-link). When custom_unsubscribe_url_enabled is skipped
	// on creating a project, it's value is taken from user
	CustomUnsubscribeURLEnabled *bool `json:"custom_unsubscribe_url_enabled,omitempty"`

	// A unique domain identifier which will determine the [tracking
	// domain](https://docs.unione.io/en/tracking-domains) or [dedicated IP](/dedicated-ip) pool to be used
	// by default. If the value is not specified, a default backend domain will be assigned for your
	// account or project instead.
	BackendDomainID *int `json:"backend_domain_id,omitempty"`

	// A unique identifier that defines the default unsubscribe page. If absent, the account/project will
	// be assigned the system default unsubscribe page.
	UnsubscribePageID *int `json:"unsubscribe_page_id,omitempty"`

	// Project's email counter value. 0 value only is accepted (which means resetting the counter to zero),
	// or refreshes automatically depending on the counter mode set with the 'email_counter_mode'
	// parameter. If the request doesn't contain the field or contains null, the limit value doesn't
	// change.
	EmailCounter *int `json:"email_counter,omitempty"`

	// Allows you to limit the number of emails sent by the project. The sending stops when the limit is
	// reached. Any non-negative 64-bit integer is acceptable. Setting to 0 means no limit. If the request
	// doesn't contain the field or contains null, the limit value doesn't change.
	// Default: `0`.
	EmailCounterLimit *int `json:"email_counter_limit,omitempty"`

	// Sets the email counter mode. It allows you to choose whether the counter should reset automatically,
	// or remain permanent. The following values are acceptable: 'default' - the counter resets to zero
	// automatically at the beginning of each accounting period, 'permanent' - the counter never resets
	// automatically. The reset can be done only manually by the user in the request with the passes 0
	// value in the email_counter field. If the request doesn't contain the field or contains a null value,
	// the limit value doesn't change.
	// Default: `default`.
	EmailCounterMode common.EmailCounterMode `json:"email_counter_mode,omitempty"`
}

// ProjectUpdateResponse
type ProjectUpdateResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// API key of the project. Returned only if the project to update was identified by project_api_key.
	// Not returned in case of updating by project_id.
	ProjectAPIKey string `json:"project_api_key,omitempty"`
}

// ProjectListRequest
type ProjectListRequest struct {
	// Optional possibility to specify project_id to include only this project to list.
	ProjectID *string `json:"project_id,omitempty"`

	// Optional possibility to specify project_api_key to include only this project to list.
	ProjectAPIKey *string `json:"project_api_key,omitempty"`
}

// ProjectListResponse
type ProjectListResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Array of objects where each object represents one project.
	Projects []ProjectListObject `json:"projects,omitempty"`
}

// ProjectListObject
type ProjectListObject struct {
	// Unqiue project identifier, ASCII string up to 36 characters long.
	ID string `json:"id,omitempty"`

	// API key of the project. You can use it instead of the user API key parameter in all methods except
	// project/* methods.
	APIKey string `json:"api_key,omitempty"`

	// Project name, unique for user account.
	Name string `json:"name,omitempty"`

	// ISO-3166 alpha-2 country code. If set, UniOne treats project personal data according to country
	// laws, e.g. GDPR for european countries.
	Country string `json:"country,omitempty"`

	// Date and time of project creation in UTC timezone in "YYYY-MM-DD hh:mm:ss" format.
	RegTime string `json:"reg_time,omitempty"`

	// Whether email sending is enabled for this project or not.
	// Default: `True`.
	SendEnabled bool `json:"send_enabled,omitempty"`

	// If it's false then UniOne adds default unsubscribe footer to each email sent with API key of the
	// project. If it's true then appending default unsubscribe footer for this project is avoided and
	// sending with custom unsubscribe url or even without unsubscribe url is permitted for the project.
	// Please note that custom_unsubscribe_url_enabled=true is available only if removing unsubscribe link
	// is approved for the user account by support. More on this
	// [here](https://docs.unione.io/en/unsubscribe-link). When custom_unsubscribe_url_enabled is skipped
	// on creating a project, it's value is taken from user
	CustomUnsubscribeURLEnabled bool `json:"custom_unsubscribe_url_enabled,omitempty"`

	// A unique domain identifier which will determine the [tracking
	// domain](https://docs.unione.io/en/tracking-domains) or [dedicated IP](/dedicated-ip) pool to be used
	// by default. If the value is not specified, a default backend domain will be assigned for your
	// account or project instead.
	BackendDomainID int `json:"backend_domain_id,omitempty"`

	// Project's email counter value. It increases on every email sent.
	EmailCounter int `json:"email_counter,omitempty"`

	// Allows you to limit the number of emails sent by the project. The sending stops when the limit is
	// reached. Any non-negative 64-bit integer is acceptable. 0 means no limit.
	EmailCounterLimit int `json:"email_counter_limit,omitempty"`

	// 'default' - the counter resets to zero automatically at the beginning of each accounting period.
	// 'permanent' - the counter never resets automatically.
	// Default: `default`.
	EmailCounterMode common.EmailCounterMode `json:"email_counter_mode,omitempty"`

	// A unique identifier that defines the default unsubscribe page. If absent, the account/project will
	// be assigned the system default unsubscribe page.
	UnsubscribePageID int `json:"unsubscribe_page_id,omitempty"`
}

// ProjectDeleteRequest
type ProjectDeleteRequest struct {
	// Unique project identifier. You should identify project to delete either with project_id or with
	// project_api_key field. We recommend usage of project_id as a more secure option.
	ProjectID *string `json:"project_id,omitempty"`

	// Optional project API key. You should identify project to delete either with project_id or with
	// project_api_key field. We recommend usage of project_id as a more secure option.
	ProjectAPIKey *string `json:"project_api_key,omitempty"`
}
