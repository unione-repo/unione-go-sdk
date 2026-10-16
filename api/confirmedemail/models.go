// Models of the ConfirmedEmail section.

package confirmedemail

import "github.com/unione-repo/unione-go-sdk/common"

// ConfirmedEmailRequest
type ConfirmedEmailRequest struct {
	// Email address to add or delete.
	Email string `json:"email"`
}

// ConfirmedEmailCreateResponse
type ConfirmedEmailCreateResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	ConfirmedEmail *ConfirmedEmail `json:"confirmed_email,omitempty"`
}

// ConfirmedEmail
type ConfirmedEmail struct {
	Email string `json:"email,omitempty"`

	// Confirmation status of the email address.
	Status common.VerificationRecordStatus `json:"status,omitempty"`
}

// ConfirmedEmailListResponse
type ConfirmedEmailListResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	ConfirmedEmails []ConfirmedEmail `json:"confirmed_emails,omitempty"`
}

// ConfirmedEmailDeleteResponse
type ConfirmedEmailDeleteResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`
}
