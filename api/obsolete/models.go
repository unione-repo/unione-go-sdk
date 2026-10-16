// Models of the Obsolete section.

package obsolete

// ObsoleteUnsubscribedSetMessage — Unsubscription process details:
//
// - "unsubscribed" if email has just been unsubscribed,
// - "was inactive" if it was inactive due some reason,
// - "was unsubscribed" if it already was unsubscribed before the call,
// - "added and unsubscribed" if it's a new email without status history and it was added and then unsubscribed immediately.
type ObsoleteUnsubscribedSetMessage string

// ObsoleteUnsubscribedSetMessage values.
const (
	ObsoleteUnsubscribedSetMessageUnsubscribed         ObsoleteUnsubscribedSetMessage = "unsubscribed"
	ObsoleteUnsubscribedSetMessageWasInactive          ObsoleteUnsubscribedSetMessage = "was inactive"
	ObsoleteUnsubscribedSetMessageWasUnsubscribed      ObsoleteUnsubscribedSetMessage = "was unsubscribed"
	ObsoleteUnsubscribedSetMessageAddedAndUnsubscribed ObsoleteUnsubscribedSetMessage = "added and unsubscribed"
)

// ObsoleteUnsubscribedSetRequest
type ObsoleteUnsubscribedSetRequest struct {
	// Email to be unsubscribed.
	Address string `json:"address"`
}

// ObsoleteUnsubscribedSetResponse
type ObsoleteUnsubscribedSetResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// The unsubscribed email.
	Address string `json:"address,omitempty"`

	// Unsubscription process details:
	//
	// - "unsubscribed" if email has just been unsubscribed,
	// - "was inactive" if it was inactive due some reason,
	// - "was unsubscribed" if it already was unsubscribed before the call,
	// - "added and unsubscribed" if it's a new email without status history and it was added and then unsubscribed immediately.
	Message ObsoleteUnsubscribedSetMessage `json:"message,omitempty"`
}

// ObsoleteUnsubscribedCheckRequest
type ObsoleteUnsubscribedCheckRequest struct {
	// Email to check.
	Address string `json:"address"`
}

// ObsoleteUnsubscribedCheckResponse
type ObsoleteUnsubscribedCheckResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Email that was checked.
	Address string `json:"address,omitempty"`

	// true if unsubscribed, false if not.
	IsUnsubscribed bool `json:"is_unsubscribed,omitempty"`
}

// ObsoleteUnsubscribedListRequest
type ObsoleteUnsubscribedListRequest struct {
	// Date in the format YYYY-MM-DD to get all unsubscribed emails from the "date_from" to the present
	// day. If the parameter is not specified, the default period is today (from 00:00 UTC).
	DateFrom *string `json:"date_from,omitempty"`
}

// ObsoleteUnsubscribedListResponse
type ObsoleteUnsubscribedListResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Array of object describing unsubscribed emails.
	Unsubscribed []ObsoleteUnsubscribedListResponseUnsubscribedItem `json:"unsubscribed,omitempty"`
}

// ObsoleteUnsubscribedListResponseUnsubscribedItem
type ObsoleteUnsubscribedListResponseUnsubscribedItem struct {
	// Unsubscribed email.
	Address string `json:"address,omitempty"`

	// Date and time when a recipient unsubscribed, in UTC timezone and "YYYY-MM-DD hh:mm:ss" format.
	UnsubscribedOn string `json:"unsubscribed_on,omitempty"`
}
