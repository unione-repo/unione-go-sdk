// Models of the EventDump section.

package eventdump

import "github.com/unione-repo/unione-go-sdk/common"

// EventDumpCreateRequest
type EventDumpCreateRequest struct {
	// Date and time in YYYY-MM-DD hh:mm:ss format, specifying period start time. Data is stored for up to
	// 32 days, depending on your tariff.
	StartTime string `json:"start_time"`

	// Date and time in YYYY-MM-DD hh:mm:ss format, specifying period end time (non-inclusive).
	EndTime *string `json:"end_time,omitempty"`

	// Maximum number of events returned (default is 50). If this value is over 100 000, several files will
	// be created, each having 100 000 events maximum.
	Limit *int `json:"limit,omitempty"`

	// For accounts with projects enabled, the value all_projects=true allows to fetch data for all
	// existing projects.
	// Default: `False`.
	AllProjects *bool `json:"all_projects,omitempty"`

	// An object with the properties of the event dump filter.
	Filter *FilterObject `json:"filter,omitempty"`

	// An array of strings containing the names of columns you need in your report, in your preferred
	// order. If the 'dump_fields' array is missing or empty, a default field list and order will be
	// applied. Allowed values, in their default order, are - 'project_id', 'event_time', 'job_id',
	// 'from_email', 'email', 'status', 'delivery_status', 'metadata', 'destination_response',
	// 'user_agent', 'url', 'ip', 'tags'. For maximum security, the 'subject' field is not included by
	// default, but you can add it by explicitly including in the 'dump_fields' array.
	DumpFields []string `json:"dump_fields,omitempty"`

	// Allows you to obtain aggregated statistics for a given time period. The two allowed values are:
	//
	// - __day_status__ - fetch aggregated statistics (the filter value must be missing or empty, as filters are not allowed with this option);
	// - __an empty string__ - used by default, event_dump/create works as if the aggregate parameter is omitted.
	Aggregate *string `json:"aggregate,omitempty"`

	// Field delimiter, can be set to either  ',' or ';', defaults to  ','.
	Delimiter *string `json:"delimiter,omitempty"`

	// File format, either “csv” (default) or “csv_gzip”.
	// Default: `csv`.
	Format *string `json:"format,omitempty"`
}

// FilterObject — An object with the properties of the event dump filter.
type FilterObject struct {
	// Job identifier returned earlier by [email/send](https://docs.unione.io/en/web-api-ref#email-send)
	// method.
	JobID *string `json:"job_id,omitempty"`

	// Email status. Possible values are:
	//
	// __accepted__ — the message has been accepted, but not sent yet.
	// __sent__ — the message has been sent, but not delivered yet.
	// __delivered__ — the message has been delivered.
	// __opened__ — the message has been delivered and read.
	// __clicked__ — the message has been delivered, read, and at least one of the links in email has been
	// clicked.
	// __unsubscribed__ — the message has been delivered to the recipient and read, but then the recipient
	// unsubscribed.
	// __subscribed__ — the message has been delivered to the recipient and read, but then the recipient
	// unsubscribed and subscribed again.
	// __soft_bounced__ — temporary delivery failure. UniOne will try to deliver the message during 48
	// hours.
	// __hard_bounced__ — failed to deliver the message, there will be no delivery attempts.
	// __spam__ — the message has been delivered, but it was reported as spam by the recipient.
	Status common.EmailStatus `json:"status,omitempty"`

	// UniOne internal detailed delivery status, you can pass several values at once, comma-separated. Some
	// possible values are:
	//
	// __err_user_unknown__ -- email address doesn't exist;
	// __err_user_inactive__ -- email address isn't used anymore;
	// __err_will_retry__ -- email was temporary rejected, delivery will be retried later;
	// __err_mailbox_discarded__ -- email address was active earlier, but now it's deleted;
	// __err_mailbox_full__ -- mailbox is full;
	// __err_spam_rejected__ -- email was rejected as a spam;
	// __err_blacklisted__ -- email was rejected because sender IP or sender domain is found in a black
	// list;
	// __err_too_large__ -- email size is over limit, according to recipient's server;
	// __err_unsubscribed__ -- email was unsubscribed;
	// __err_unreachable__ -- numerous delivery failures lead to mark this email address as permanently
	// unavailable;
	// __err_skip_letter__ -- sending has been canceled because the email address is temporarily
	// unavailable;
	// __err_domain_inactive__ -- the domain does not accept mail or does not exist;
	// __err_destination_misconfigured__ -- the domain does not accept mail due to incorrect settings on
	// the recipient’s side, and the server’s response contains information about a cause that can be fixed
	// (for example, an inoperative blacklist is used, etc.);
	// __err_delivery_failed__ -- delivery failed due to other reasons;
	// __err_spam_skipped__ -- sending has been canceled because the campaign has been blocked as a spam;
	// __err_lost__ -- the message was not sent due to inconsistency of its parts or it has been lost
	// because of the failure on our side. The sender needs to re-send the letter on his own, since the
	// original has not been saved;
	//
	// We reserve the right to change or add internal delivery statuses as we need, so the following list
	// is provided as an example only, without any description: ok_sent, ok_delivered, ok_read,
	// ok_link_visited, ok_unsubscribed, ok_resubscribed, ok_spam_folder, ok_fbl, not_sent,
	// skip_dup_unreachable, skip_dup_temp_unreachable, skip_dup_mailbox_full, err_spam_removed,
	// err_resend, err_unknown, err_retry_letter, err_src_invalid, err_dest_invalid, err_not_available,
	// err_internal, err_no_dns, err_no_smtp, err_giveup.
	DeliveryStatus *string `json:"delivery_status,omitempty"`

	// Recipient email address
	Email *string `json:"email,omitempty"`

	// Sender email address
	EmailFrom *string `json:"email_from,omitempty"`

	// Recipient domain
	Domain *string `json:"domain,omitempty"`

	// Campaign identifier, unsigned decimal integer or UUID up to 128-bit, passed in metadata with name
	// "campaign_id" (name is configurable thru support).
	CampaignID *string `json:"campaign_id,omitempty"`
}

// EventDumpCreateResponse
type EventDumpCreateResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Dump identifier, used to call the
	// [event-dump/get](https://docs.unione.io/en/web-api-ref#event-dump-get) method.
	DumpID string `json:"dump_id,omitempty"`
}

// EventDumpGetRequest
type EventDumpGetRequest struct {
	// Dump identifier received by the method
	// [event-dump/create](https://docs.unione.io/en/web-api-ref#event-dump-create).
	DumpID string `json:"dump_id"`
}

// EventDumpGetResponse
type EventDumpGetResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// An object with the properties of the event dump.
	EventDump *EventDumpObject `json:"event_dump,omitempty"`
}

// EventDumpObject — An object with the properties of the event dump.
type EventDumpObject struct {
	// Dump ID received by the method
	// [event-dump/create](https://docs.unione.io/en/web-api-ref#event-dump-create).
	DumpID string `json:"dump_id,omitempty"`

	// Task status, possible values are:
	//
	// - __queued__ - the task is in the queue, processing is not started yet;
	// - __in_process__ - the request is being processed;
	// - __ready__ - the task is finished, dump file is ready for download;
	// - __failed__ -  processing failed due to an internal error.
	DumpStatus string `json:"dump_status,omitempty"`

	// An array of objects, each representing a file, ready for download. You may start downloading
	// everything listed in “files” even if "dump_status" is still "in_process". If there are no events,
	// according to the specified parameters, an empty array will be returned.
	Files []EventDumpFilesObject `json:"files,omitempty"`
}

// EventDumpFilesObject — An object with the properties of the event dump file.
type EventDumpFilesObject struct {
	// Use this unique URL to download the requested dump file (either .csv or .csv.gz, depending on your
	// request).
	URL string `json:"url,omitempty"`

	// File size.
	Size int `json:"size,omitempty"`
}

// EventDumpListResponse
type EventDumpListResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// An array of objects describing event dumps. Each object is similar to 'event_dump' object returned
	// by the [event-dump/get](https://docs.unione.io/en/web-api-ref#event-dump-get) method.
	EventDumps []EventDumpObject `json:"event_dumps,omitempty"`
}

// EventDumpDeleteRequest
type EventDumpDeleteRequest struct {
	// Dump identifier, obtained via the
	// [event-dump/create](https://docs.unione.io/en/web-api-ref#event-dump-create) method.
	DumpID string `json:"dump_id"`
}
