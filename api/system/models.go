// Models of the System section.

package system

import "github.com/unione-repo/unione-go-sdk/common"

// SystemPingResponse
type SystemPingResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Unique user identifier.
	UserID int `json:"user_id,omitempty"`
}

// SystemInfoResponse
type SystemInfoResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Unique user identifier.
	UserID int `json:"user_id,omitempty"`

	// Email of the user.
	Email string `json:"email,omitempty"`

	// Unqiue project identifier, ASCII string up to 36 characters long. Present only if the API key used
	// for request is the project API key.
	ProjectID string `json:"project_id,omitempty"`

	// Project name, unique for user account. Present only if the API key used for request is the project
	// API key.
	ProjectName string `json:"project_name,omitempty"`

	// Object with email counter properties. Present only if the API key used for request is the user API
	// key.
	ProjectAccounting *ProjectAccountingObject `json:"project_accounting,omitempty"`

	// Object with accounting period properties. Present only if the API key used for request is the user
	// API key.
	Accounting *SystemInfoResponseAccounting `json:"accounting,omitempty"`
}

// ProjectAccountingObject — Object with email counter properties. Present only if the API key used for request is the user API
// key.
type ProjectAccountingObject struct {
	// Project's email counter value. It increases on every email sent.
	EmailCounter int `json:"email_counter,omitempty"`

	// Allows you to limit the number of emails sent by the project. The sending stops when the limit is
	// reached. Any non-negative 64-bit integer is acceptable. 0 means no limit.
	EmailCounterLimit int `json:"email_counter_limit,omitempty"`

	// 'default' - the counter resets to zero automatically at the beginning of each accounting period.
	// 'permanent' - the counter never resets automatically.
	// Default: `default`.
	EmailCounterMode common.EmailCounterMode `json:"email_counter_mode,omitempty"`
}

// SystemInfoResponseAccounting — Object with accounting period properties. Present only if the API key used for request is the user
// API key.
type SystemInfoResponseAccounting struct {
	// Date and time of accounting period start in UTC in "YYYY-MM-DD hh:mm:ss" format.
	PeriodStart string `json:"period_start,omitempty"`

	// Date and time of accounting period end in UTC in "YYYY-MM-DD hh:mm:ss" format.
	PeriodEnd string `json:"period_end,omitempty"`

	// Number of emails included into accounting period.
	EmailsIncluded int `json:"emails_included,omitempty"`

	// Number of emails sent during accounting period. Can be bigger than emails_included in case of
	// overage.
	EmailsSent int `json:"emails_sent,omitempty"`
}
