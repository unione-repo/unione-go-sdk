// Models of the Webhook section.

package webhook

import "github.com/unione-repo/unione-go-sdk/common"

// WebhookStatus — Webhook status, "active" by default. "disabled" means that webhook has been disabled by the user,
// "stopped" means that webhook has been stopped by the system after 24h of failed calls (with minimum
// of 10 distinct events).
type WebhookStatus string

// WebhookStatus values.
const (
	WebhookStatusActive   WebhookStatus = "active"
	WebhookStatusDisabled WebhookStatus = "disabled"
	WebhookStatusStopped  WebhookStatus = "stopped"
)

// WebhookEventFormat — Notification format. "json_post"(default) or "json_post_gzip".
type WebhookEventFormat string

// WebhookEventFormat values.
const (
	WebhookEventFormatJSONPost     WebhookEventFormat = "json_post"
	WebhookEventFormatJSONPostGzip WebhookEventFormat = "json_post_gzip"
)

// EventName — Type of event data contained in event_data field, either "transactional_email_status" or
// "transactional_spam_block".
type EventName string

// EventName values.
const (
	EventNameTransactionalEmailStatus EventName = "transactional_email_status"
	EventNameTransactionalSpamBlock   EventName = "transactional_spam_block"
)

// DeviceType — Recipient's device type, detected from the user agent. Present for "clicked" and "opened" statuses;
// equals "unknown" if the device type couldn't be detected.
type DeviceType string

// DeviceType values.
const (
	DeviceTypeDesktop DeviceType = "desktop"
	DeviceTypeMobile  DeviceType = "mobile"
	DeviceTypeTablet  DeviceType = "tablet"
	DeviceTypeOther   DeviceType = "other"
	DeviceTypeUnknown DeviceType = "unknown"
)

// TrafficType — Traffic quality for "opened" and "clicked" statuses. "human" means a recognized browser or email
// client without machine signs; "machine" means a bot, link scanner or security proxy; "protected"
// means a mailbox/OS privacy proxy where a human action cannot be confirmed (Apple Mail Privacy
// Protection / iCloud Private Relay, Gmail Image Proxy — Apple MPP is protected, not machine);
// "unknown" means the user agent is missing, broken or not recognized.
type TrafficType string

// TrafficType values.
const (
	TrafficTypeHuman     TrafficType = "human"
	TrafficTypeMachine   TrafficType = "machine"
	TrafficTypeProtected TrafficType = "protected"
	TrafficTypeUnknown   TrafficType = "unknown"
)

// BlockType — Spam block type, either single or multiple sending SMTP. For single sending SMTP block in common
// pool UniOne retries several other SMTPs. This property exists only if
// event_name="transactional_spam_block".
type BlockType string

// BlockType values.
const (
	BlockTypeOneSMTP BlockType = "one_smtp"
	BlockTypeAllSMTP BlockType = "all_smtp"
)

// DomainStatus — Whether it's a block or unblock event. This property exists only if
// event_name="transactional_spam_block".
type DomainStatus string

// DomainStatus values.
const (
	DomainStatusBlocked   DomainStatus = "blocked"
	DomainStatusUnblocked DomainStatus = "unblocked"
)

// WebhookSetRequest
type WebhookSetRequest struct {
	// URL that will receive the notification when an event occurs. The URL should be unique for a user or
	// a project. Only ASCII characters are supported for now in URL, please convert to
	// [Punycode](https://en.wikipedia.org/wiki/Punycode) if you need to use non-ASCII characters.
	URL string `json:"url"`

	// Webhook status, "active" by default. "disabled" means that webhook has been disabled by the user,
	// "stopped" means that webhook has been stopped by the system after 24h of failed calls (with minimum
	// of 10 distinct events).
	// Default: `active`.
	Status WebhookStatus `json:"status,omitempty"`

	// Notification format. "json_post"(default) or "json_post_gzip".
	// Default: `json_post`.
	EventFormat WebhookEventFormat `json:"event_format,omitempty"`

	// Shall detailed delivery info be returned(1) or not(0). When it's 1, the information about SMTP
	// response and internal delivery status is returned for "hard_bounced" and "soft_bounced" email
	// statuses; the user agent, recipient IP address, detected geo fields, device info detected from the
	// user agent (device type and brand, OS name, browser or email client name) and traffic_type (human,
	// machine, protected or unknown) are returned for "opened" and "clicked" statuses; URL is returned for
	// "clicked" status. For "delivered", "soft_bounced" and "hard_bounced" statuses UniOne may also return
	// sender_ip, the IP address of the SMTP server that sent the message, if available.
	// Allowed values: `0`, `1`.
	// Default: `0`.
	DeliveryInfo *int `json:"delivery_info,omitempty"`

	// 0=several event notifications can be reported in single webhook call, 1=only single event can be
	// reported in one call (not recommended).
	// Allowed values: `0`, `1`.
	// Default: `0`.
	SingleEvent *int `json:"single_event,omitempty"`

	// Maximum quantity of permitted parallel queries to your server. The more your server can handle - the
	// better.
	// Default: `10`.
	MaxParallel *int `json:"max_parallel,omitempty"`

	// Object containing events to notify of.
	Events *WebhookEventsObject `json:"events,omitempty"`
}

// WebhookEventsObject — Object containing events to notify of.
type WebhookEventsObject struct {
	// If present then spam block events will be reported. Should contain single array element with "*"
	// string.
	SpamBlock []string `json:"spam_block,omitempty"`

	// If present then email status change events will be reported. Contains names of statuses to notify
	// of.
	EmailStatus []common.EmailStatus `json:"email_status,omitempty"`
}

// WebhookSetResponse
type WebhookSetResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Object with all the webhook properties.
	Object *WebhookObject `json:"object,omitempty"`
}

// WebhookObject — Object with all the webhook properties.
type WebhookObject struct {
	// Webhook unique identifier.
	ID int `json:"id,omitempty"`

	// Webhook URL.
	URL string `json:"url,omitempty"`

	// Webhook status, "active" by default. "disabled" means that webhook has been disabled by the user,
	// "stopped" means that webhook has been stopped by the system after 24h of failed calls (with minimum
	// of 10 distinct events).
	// Default: `active`.
	Status WebhookStatus `json:"status,omitempty"`

	// Notification format. "json_post"(default) or "json_post_gzip".
	// Default: `json_post`.
	EventFormat WebhookEventFormat `json:"event_format,omitempty"`

	// Shall detailed delivery info be returned(1) or not(0). When it's 1, the information about SMTP
	// response and internal delivery status is returned for "hard_bounced" and "soft_bounced" email
	// statuses; the user agent, recipient IP address, detected geo fields, device info detected from the
	// user agent (device type and brand, OS name, browser or email client name) and traffic_type (human,
	// machine, protected or unknown) are returned for "opened" and "clicked" statuses; URL is returned for
	// "clicked" status. For "delivered", "soft_bounced" and "hard_bounced" statuses UniOne may also return
	// sender_ip, the IP address of the SMTP server that sent the message, if available.
	// Allowed values: `0`, `1`.
	// Default: `0`.
	DeliveryInfo int `json:"delivery_info,omitempty"`

	// 0=several event notifications can be reported in single webhook call, 1=only single event can be
	// reported in one call (not recommended).
	// Allowed values: `0`, `1`.
	// Default: `0`.
	SingleEvent int `json:"single_event,omitempty"`

	// Maximum quantity of permitted parallel queries to your server. The more your server can handle - the
	// better.
	// Default: `10`.
	MaxParallel int `json:"max_parallel,omitempty"`

	// Webhook properties last update date and time in UTC timezone in "YYYY-MM-DD hh:mm:ss" format.
	UpdatedAt string `json:"updated_at,omitempty"`

	// Object containing events to notify of.
	Events *WebhookEventsObject `json:"events,omitempty"`
}

// WebhookCallback — Webhook callback format description.
type WebhookCallback struct {
	// MD5-hash of the message body, in which the value "auth" is replaced by api key of the user/project;
	// with this field the recipient of the notification can both authenticate and verify the notification
	// integrity
	Auth string `json:"auth,omitempty"`

	// Array with only one element, containing events of a user/project.
	EventsByUser []WebhookCallbackEventsByUserItem `json:"events_by_user,omitempty"`
}

// WebhookCallbackEventsByUserItem
type WebhookCallbackEventsByUserItem struct {
	// Unique user identifier.
	UserID int `json:"user_id,omitempty"`

	// Project identifier, present only if webhook was registered for the project using project API key.
	ProjectID string `json:"project_id,omitempty"`

	// Project name, present only if webhook was registered for the project using project API key.
	ProjectName string `json:"project_name,omitempty"`

	// Array of events reported by webhook.
	Events []WebhookCallbackEventsByUserItemEventsItem `json:"events,omitempty"`
}

// WebhookCallbackEventsByUserItemEventsItem
type WebhookCallbackEventsByUserItemEventsItem struct {
	// Type of event data contained in event_data field, either "transactional_email_status" or
	// "transactional_spam_block".
	EventName EventName `json:"event_name,omitempty"`

	// Object with different event properties depending on "event_name". Below you can see all the
	// properties, "transactional_email_status"-related first and then "transactional_spam_block"-related.
	EventData *WebhookCallbackEventsByUserItemEventsItemEventData `json:"event_data,omitempty"`
}

// WebhookCallbackEventsByUserItemEventsItemEventData — Object with different event properties depending on "event_name". Below you can see all the
// properties, "transactional_email_status"-related first and then "transactional_spam_block"-related.
type WebhookCallbackEventsByUserItemEventsItemEventData struct {
	// Job identifier returned earlier by [email/send](https://docs.unione.io/en/web-api-ref#email-send)
	// method. Property exists only if event_name="transactional_email_status".
	JobID string `json:"job_id,omitempty"`

	// Metadata passed in [email/send](https://docs.unione.io/en/web-api-ref#email-send) method in
	// recipients.metadata or global_metadata properties. This property exists only if
	// event_name="transactional_email_status".
	Metadata map[string]string `json:"metadata,omitempty"`

	// Recipient's email. This property exists only if event_name="transactional_email_status".
	Email string `json:"email,omitempty"`

	// Email delivery status. This property exists only if event_name="transactional_email_status".
	// Possible values are:
	//
	// __accepted__ — the message has been accepted, but not sent yet.
	// __sent__ — the message has been sent, but not delivered yet.
	// __delivered__ — the message has been delivered. It can be changed to “opened”, “clicked”,
	// “unsubscribed”, “subscribed” or “spam”.
	// __opened__ — the message has been delivered and read. It can be changed to “clicked”,
	// “unsubscribed”, “subscribed” or “spam”.
	// __clicked__ — the message has been delivered, read, and at least one of the links in email has been
	// clicked. It can be changed to “unsubscribed”, “subscribed” or “spam”.
	// __unsubscribed__ — the message has been delivered to the recipient and read, but then the recipient
	// unsubscribed. It can be changed to “subscribed”.
	// __subscribed__ — the message has been delivered to the recipient and read, but then the recipient
	// unsubscribed and subscribed again. It can be changed to “unsubscribed”.
	// __soft_bounced__ — temporary delivery failure. UniOne will try to deliver the message during 48
	// hours. In case of success the status changes to “delivered”, in case of failure to deliver during 48
	// hours message status changes to “hard_bounced”.
	// __hard_bounced__ — failed to deliver the message, there will be no delivery attempts. It's a final
	// status. There are a lot of possible reasons, you can use SMTP response from
	// delivery_info.destination_response field or UniOne internal reason classification from
	// delivery_info.delivery_status field to analyze the reason.
	// __spam__ — the message has been delivered, but it was reported as spam by the recipient. UniOne can
	// receive and process the spam complaint from several domains, including msn.com, outlook.com,
	// hotmail.com, live.com, ukr.net, yahoo.com, aol.com using
	// [FBL](https://en.wikipedia.org/wiki/Feedback_loop_(email)).
	Status common.EmailStatus `json:"status,omitempty"`

	// Event date & time in UTC time zone in "YYYY-MM-DD hh:mm:ss" format. This property exists only if
	// event_name="transactional_email_status".
	EventTime string `json:"event_time,omitempty"`

	// URL for "opened" and "clicked" statuses. This property exists only if
	// event_name="transactional_email_status".
	URL string `json:"url,omitempty"`

	// Object with detailed delivery info.Event date & time in UTC time zone in "YYYY-MM-DD hh:mm:ss"
	// format. This property exists only if webhook has delivery_info proeprty set to 1 and
	// event_name="transactional_email_status".
	DeliveryInfo *WebhookCallbackEventsByUserItemEventsItemEventDataDeliveryInfo `json:"delivery_info,omitempty"`

	// Spam block date & time in UTC time zone in "YYYY-MM-DD hh:mm:ss" format. This property exists only
	// if event_name="transactional_spam_block".
	BlockTime string `json:"block_time,omitempty"`

	// Spam block type, either single or multiple sending SMTP. For single sending SMTP block in common
	// pool UniOne retries several other SMTPs. This property exists only if
	// event_name="transactional_spam_block".
	BlockType BlockType `json:"block_type,omitempty"`

	// Domain that blocked sending. This property exists only if event_name="transactional_spam_block".
	Domain string `json:"domain,omitempty"`

	// Number of sending SMTPs blocked. This property exists only if event_name="transactional_spam_block".
	SMTPBlocksCount int `json:"SMTP_blocks_count,omitempty"`

	// Whether it's a block or unblock event. This property exists only if
	// event_name="transactional_spam_block".
	DomainStatus DomainStatus `json:"domain_status,omitempty"`
}

// WebhookCallbackEventsByUserItemEventsItemEventDataDeliveryInfo — Object with detailed delivery info.Event date & time in UTC time zone in "YYYY-MM-DD hh:mm:ss"
// format. This property exists only if webhook has delivery_info proeprty set to 1 and
// event_name="transactional_email_status".
type WebhookCallbackEventsByUserItemEventsItemEventDataDeliveryInfo struct {
	// UniOne internal detailed delivery status. Some possible values are:
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
	DeliveryStatus string `json:"delivery_status,omitempty"`

	// SMTP response.
	DestinationResponse string `json:"destination_response,omitempty"`

	// User agent of recipient. Present only if detected for "clicked" and "opened" statuses.
	UserAgent string `json:"user_agent,omitempty"`

	// Recipient's device type, detected from the user agent. Present for "clicked" and "opened" statuses;
	// equals "unknown" if the device type couldn't be detected.
	DeviceType DeviceType `json:"device_type,omitempty"`

	// Recipient's device brand, detected from the user agent. Present for "clicked" and "opened" statuses
	// if detected. For device_type "mobile" and "tablet" it's always present and equals "unknown" when the
	// brand couldn't be detected.
	DeviceBrand string `json:"device_brand,omitempty"`

	// Name of the operating system of the recipient's device, detected from the user agent. Present for
	// "clicked" and "opened" statuses; equals "unknown" if the OS couldn't be detected.
	OsName string `json:"os_name,omitempty"`

	// Name of the browser the email was opened in, detected from the user agent. Present for "clicked" and
	// "opened" statuses only if a browser was detected, so it never comes together with email_client_name.
	BrowserName string `json:"browser_name,omitempty"`

	// Name of the email client application the email was opened in, detected from the user agent. Present
	// for "clicked" and "opened" statuses only if an email client was detected, so it never comes together
	// with browser_name.
	EmailClientName string `json:"email_client_name,omitempty"`

	// Traffic quality for "opened" and "clicked" statuses. "human" means a recognized browser or email
	// client without machine signs; "machine" means a bot, link scanner or security proxy; "protected"
	// means a mailbox/OS privacy proxy where a human action cannot be confirmed (Apple Mail Privacy
	// Protection / iCloud Private Relay, Gmail Image Proxy — Apple MPP is protected, not machine);
	// "unknown" means the user agent is missing, broken or not recognized.
	TrafficType TrafficType `json:"traffic_type,omitempty"`

	// Recipient's IP address. Present only if detected for "clicked" and "opened" statuses.
	IP string `json:"ip,omitempty"`

	// ISO 3166-1 alpha-2 country code detected from recipient IP. Present for "clicked" and "opened"
	// statuses.
	CountryCode string `json:"country_code,omitempty"`

	// Country detected from recipient IP. Present for "clicked" and "opened" statuses.
	Country string `json:"country,omitempty"`

	// City detected from recipient IP. Present for "clicked" and "opened" statuses.
	City string `json:"city,omitempty"`

	// IP address of the SMTP server that sent the message. Present only for "delivered", "soft_bounced"
	// and "hard_bounced" statuses, if available.
	SenderIP string `json:"sender_ip,omitempty"`
}

// WebhookGetRequest
type WebhookGetRequest struct {
	// Webhook URL.
	URL string `json:"url"`
}

// WebhookGetResponse
type WebhookGetResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Object with all the webhook properties.
	Object *WebhookObject `json:"object,omitempty"`
}

// WebhookListRequest
type WebhookListRequest struct {
	// Limits the number of webhooks returned. 50 is recommended.
	Limit *int `json:"limit,omitempty"`

	// Index of the first webhook in the list, starting with 0.
	// Default: `0`.
	Offset *int `json:"offset,omitempty"`
}

// WebhookListResponse
type WebhookListResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// Array of webhooks objects.
	Objects []WebhookObject `json:"objects,omitempty"`
}

// WebhookDeleteRequest
type WebhookDeleteRequest struct {
	// Webhook URL.
	URL string `json:"url"`
}
