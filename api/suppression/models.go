// Models of the Suppression section.

package suppression

// SuppressionSetRequest
type SuppressionSetRequest struct {
	// Email to add in suppression list.
	Email string `json:"email"`

	// Cause of email being suppressed. One of:
	//
	// - __unsubscribed__ - email is unsubscribed;
	// - __temporary_unavailable__ - the email address is unavailable. This means that over the next three days sending to this address will return an error. Email may be temporarily unavailable due to several reasons, e.g.:
	//
	// - a previous email has been rejected by the recipient's server for spam;
	// - the recipient's mailbox is full or is not used;
	// - recipient's domain does not accept mail;
	// - sending server was rejected due to blacklisting;
	//
	// - __permanent_unavailable__ - the email address is permanently unavailable due to multiple hard bounces;
	// - __complained__ -  the recipient reported spam in the previous emails;
	Cause string `json:"cause"`

	// When suppression was created, in UTC timezone in "YYYY-MM-DD hh:mm:ss" format.
	Created *string `json:"created,omitempty"`
}

// SuppressionGetRequest
type SuppressionGetRequest struct {
	// Email to get suppression details for.
	Email string `json:"email"`

	// If a user has projects functionality enabled, he/she can pass all_projects=true to search in all
	// projects' data.
	// Default: `False`.
	AllProjects *bool `json:"all_projects,omitempty"`
}

// SuppressionGetResponse
type SuppressionGetResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// The email for which suppression details were requested.
	Email string `json:"email,omitempty"`

	// Array of suppression objects.
	Suppressions []SuppressionGetObject `json:"suppressions,omitempty"`
}

// SuppressionGetObject — Object with suppression details.
type SuppressionGetObject struct {
	// Unqiue project identifier, ASCII string up to 36 characters long.
	ProjectID string `json:"project_id,omitempty"`

	// Cause of email being suppressed. One of:
	//
	// - __unsubscribed__ - email is unsubscribed;
	// - __temporary_unavailable__ - the email address is unavailable. This means that over the next three days sending to this address will return an error. Email may be temporarily unavailable due to several reasons, e.g.:
	//
	// - a previous email has been rejected by the recipient's server for spam;
	// - the recipient's mailbox is full or is not used;
	// - recipient's domain does not accept mail;
	// - sending server was rejected due to blacklisting;
	//
	// - __permanent_unavailable__ - the email address is permanently unavailable due to multiple hard bounces;
	// - __complained__ -  the recipient reported spam in the previous emails;
	// - __blocked__ -  sending to the email is prohibited by administration of UniOne.
	// We may add some new causes in the future.
	Cause string `json:"cause,omitempty"`

	// Source of email being suppressed. One of:
	//
	// - __user__ - suppressed by user with    [suppression/set](https://docs.unione.io/en/web-api-ref#suppression-set);
	// - __system__ - sending to the email is prohibited by system, for example due to multiple hard bounces;
	// - __subscriber__ - the recipient reported spam or unsubscribed in the previous emails.
	Source string `json:"source,omitempty"`

	// Is it possible to delete this suppression by calling suppression/delete method.
	IsDeletable bool `json:"is_deletable,omitempty"`

	// When suppression was created, in UTC timezone in "YYYY-MM-DD hh:mm:ss" format.
	Created string `json:"created,omitempty"`
}

// SuppressionListRequest
type SuppressionListRequest struct {
	// Cause of email being suppressed. One of:
	//
	// - __unsubscribed__ - email is unsubscribed;
	// - __temporary_unavailable__ - the email address is unavailable. This means that over the next three days sending to this address will return an error. Email may be temporarily unavailable due to several reasons, e.g.:
	//
	// - a previous email has been rejected by the recipient's server for spam;
	// - the recipient's mailbox is full or is not used;
	// - recipient's domain does not accept mail;
	// - sending server was rejected due to blacklisting;
	//
	// - __permanent_unavailable__ - the email address is permanently unavailable due to multiple hard bounces;
	// - __complained__ -  the recipient reported spam in the previous emails;
	// - __blocked__ -  sending to the email is prohibited by administration of UniOne.
	// We may add some new causes in the future.
	Cause *string `json:"cause,omitempty"`

	// Source of email being suppressed. One of:
	//
	// - __user__ - suppressed by user with    [suppression/set](https://docs.unione.io/en/web-api-ref#suppression-set);
	// - __system__ - sending to the email is prohibited by system, for example due to multiple hard bounces;
	// - __subscriber__ - the recipient reported spam or unsubscribed in the previous emails.
	Source *string `json:"source,omitempty"`

	// Date in the format YYYY-MM-DD hh:mm:ss to get suppression list from the "start_time" to the present
	// day. Ignored if "cursor" is not empty.
	StartTime *string `json:"start_time,omitempty"`

	// The parameter indicates from which position the selection is to be started. Must be empty or omitted
	// for the first data chunk. In order to get subsequent chunks, you must set the "cursor" parameter in
	// your request, using the value received in response to the previous request.
	Cursor *string `json:"cursor,omitempty"`

	// Limits the number of records to be returned at one time, default is 50.
	Limit *int `json:"limit,omitempty"`
}

// SuppressionListResponse
type SuppressionListResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Array of suppression objects.
	Suppressions []SuppressionListObject `json:"suppressions,omitempty"`

	// The parameter indicates from which position the selection is to be started. Must be empty or omitted
	// for the first data chunk. In order to get subsequent chunks, you must set the "cursor" parameter in
	// your request, using the value received in response to the previous request.
	Cursor string `json:"cursor,omitempty"`
}

// SuppressionListObject — Object with suppression details.
type SuppressionListObject struct {
	// Email to get suppression details for.
	Email string `json:"email,omitempty"`

	// Cause of email being suppressed. One of:
	//
	// - __unsubscribed__ - email is unsubscribed;
	// - __temporary_unavailable__ - the email address is unavailable. This means that over the next three days sending to this address will return an error. Email may be temporarily unavailable due to several reasons, e.g.:
	//
	// - a previous email has been rejected by the recipient's server for spam;
	// - the recipient's mailbox is full or is not used;
	// - recipient's domain does not accept mail;
	// - sending server was rejected due to blacklisting;
	//
	// - __permanent_unavailable__ - the email address is permanently unavailable due to multiple hard bounces;
	// - __complained__ -  the recipient reported spam in the previous emails;
	// - __blocked__ -  sending to the email is prohibited by administration of UniOne.
	// We may add some new causes in the future.
	Cause string `json:"cause,omitempty"`

	// Source of email being suppressed. One of:
	//
	// - __user__ - suppressed by user with    [suppression/set](https://docs.unione.io/en/web-api-ref#suppression-set);
	// - __system__ - sending to the email is prohibited by system, for example due to multiple hard bounces;
	// - __subscriber__ - the recipient reported spam or unsubscribed in the previous emails.
	Source string `json:"source,omitempty"`

	// Is it possible to delete this suppression by calling suppression/delete method.
	IsDeletable bool `json:"is_deletable,omitempty"`

	// When suppression was created, in UTC timezone in "YYYY-MM-DD hh:mm:ss" format.
	Created string `json:"created,omitempty"`
}

// SuppressionDeleteRequest
type SuppressionDeleteRequest struct {
	// An email to delete from suppression list.
	Email string `json:"email"`
}
