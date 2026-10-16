// EmailAPI: API methods of the Email section.

package email

import (
	"context"

	"github.com/unione-repo/unione-go-sdk/client"
	"github.com/unione-repo/unione-go-sdk/common"
)

// EmailAPI — This group of methods is related with email sending.
type EmailAPI struct {
	client *client.APIClient
}

// NewEmailAPI creates the Email API.
func NewEmailAPI(c *client.APIClient) *EmailAPI {
	return &EmailAPI{client: c}
}

// Send — Sends an email or a bunch of emails
//
// While sending you can provide substitutions (merge tags), use a template, turn on read or click
// tracking, etc.
//
// Restrictions:
// * maximum request size is 10 MB
// * maximum number of recipients is 500
// * The symbol '/' is not allowed in attachment names
//
// If the list of recipients contains invalid, non-existent, duplicate, or unsubscribed emails, those
// emails will be returned in the output parameter failed_emails, and the email will be sent to all
// other valid emails addresses. If all the emails addresses fall into failed_emails, sending will not
// be carried out and API error 204 will be returned.
//
// Documentation: https://docs.unione.io/en/web-api-ref#email-send
func (a *EmailAPI) Send(ctx context.Context, request *EmailSendRequest) (*EmailSendResponse, error) {
	response := &EmailSendResponse{}
	if err := a.client.Call(ctx, "/email/send.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Subscribe — Sends an email with (re)subscribe link.
//
// This method allows subscription renewal by a former subscriber who has previously unsubscribed, or
// who is blocked because of complaints. This method sends an email with a new subscription link to a
// former subscriber. Once the recipient clicks the link, their subscription status is restored.
// There is a default restriction that limits this method to being used once a day. If you need to
// resubscribe more often, please ask support about it.
//
// Documentation: https://docs.unione.io/en/web-api-ref#email-subscribe
func (a *EmailAPI) Subscribe(ctx context.Context, request *EmailSubscribeRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/email/subscribe.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
