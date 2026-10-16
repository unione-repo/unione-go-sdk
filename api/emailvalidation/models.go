// Models of the EmailValidation section.

package emailvalidation

// EmailValidationSingleRequest
type EmailValidationSingleRequest struct {
	// Email address to be checked.
	Email string `json:"email"`
}

// EmailValidationSingleResponse
type EmailValidationSingleResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Email address to be checked.
	Email string `json:"email,omitempty"`

	// Possible statuses:
	//
	// - "valid" - the address is valid,
	// - "invalid" - the address is invalid,
	// - "suspicious" - the address looks suspicious,
	// - "unknown" - could not perform validation, the domain’s mail server has not responded within the time limit.
	Result string `json:"result,omitempty"`

	// Possible statuses:
	//
	// - "no_mx_record" - no MX record found for the target domain,
	// - "syntax_error" - the address syntax is invalid,
	// - "possible_typo" - the address is likely to have a typo,
	// - "mailbox_not_found" - the address does not exist,
	// - "global_suppression" - the address has been marked as unreachable due to multiple previous delivery errors,
	// - "disposable" - this is a disposable one-time email address which is usually valid for a few minutes only,
	// - "role" - the address is not likely to belong to an actual person, but rather to a certain business staff role,
	// - "abuse" - the address is known to be a source of a large number of complaints, sometimes issued automatically,
	// - "spamtrap" - this email is a spam trap, it is published openly but never used for actual emails. Sending messages to such addresses has a huge negative impact on reputation score,
	// - "smtp_connection_failed" - the domain’s SMTP server does not respond; the address may contain a typo.
	Cause string `json:"cause,omitempty"`

	// Validity score, from 0 to 100 (0 – invalid, 100 – valid).
	Validity int `json:"validity,omitempty"`

	// Local part (everything preceding the @ sign).
	LocalPart string `json:"local_part,omitempty"`

	// Domain name part.
	Domain string `json:"domain,omitempty"`

	// True if the address’ domain does have an MX record, false if does not.
	MxFound bool `json:"mx_found,omitempty"`

	// Preferred MX record for the domain.
	MxRecord string `json:"mx_record,omitempty"`

	// For addresses which are likely to have typing errors (cause=possible_typo), a suggested variant with
	// a fixed typo.
	DidYouMean string `json:"did_you_mean,omitempty"`

	// Email check date and time, YYYY-MM-DD hh:mm:ss UTC.
	ProcessedAt string `json:"processed_at,omitempty"`
}
