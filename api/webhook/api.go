// WebhookAPI: API methods of the Webhook section.

package webhook

import (
	"context"
	"encoding/json"

	"github.com/unione-repo/unione-go-sdk/client"
	"github.com/unione-repo/unione-go-sdk/common"
)

// WebhookAPI — UniOne lets you recieve notifications about email status changes or spam blocks. You provide an URL
// and every time a watched event occurs we will notify you by submitting a request with JSON data to
// this URL.
//
// You can configure the format of webhook calls (JSON or gzipped JSON), which events you want to be
// notified of, whether or not you let us groups several events in one request, do you want to receive
// extra details like SMTP answer or recipient's user agent, etc.
//
// Please note that your webhook handler should be able to process at least 5 parallel requests, up to
// 100 events each (but you can configure the maximum number of parallel connections up to 100, which
// is recommended).
//
// If the URL of your webhook handler is not available (the response is not HTTP 200 OK within 3
// seconds timeout), the system will repeat requests each 10 minutes during 24 hours, passing extra
// parameter retry_count. If there were at least 10 events during last 24 hours and all the webhook
// requests about these events have failed then we consider webhook handler broken and change it's
// status to inactive automatically. A notification about that is sent to an email [configured in
// control panel](https://cp.unione.io/en/user/alerting/technical/).
//
// The format of webhook request and details of each event are described
// [here](https://docs.unione.io/en/web-api-ref#callback-format).
type WebhookAPI struct {
	client *client.APIClient
}

// NewWebhookAPI creates the Webhook API.
func NewWebhookAPI(c *client.APIClient) *WebhookAPI {
	return &WebhookAPI{client: c}
}

// Set — Sets or edits a [webhook](https://docs.unione.io/en/web-api-ref#webhook), i.e. an event notification
// handler.
//
// Documentation: https://docs.unione.io/en/web-api-ref#webhook-set
func (a *WebhookAPI) Set(ctx context.Context, request *WebhookSetRequest) (*WebhookSetResponse, error) {
	response := &WebhookSetResponse{}
	if err := a.client.Call(ctx, "/webhook/set.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Get — Gets properties of a [webhook](https://docs.unione.io/en/web-api-ref#webhook).
//
// Documentation: https://docs.unione.io/en/web-api-ref#webhook-get
func (a *WebhookAPI) Get(ctx context.Context, request *WebhookGetRequest) (*WebhookGetResponse, error) {
	response := &WebhookGetResponse{}
	if err := a.client.Call(ctx, "/webhook/get.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — List all or some [webhooks](https://docs.unione.io/en/web-api-ref#webhook) (event notification
// handlers) of a user or a project.
//
// Documentation: https://docs.unione.io/en/web-api-ref#webhook-list
func (a *WebhookAPI) List(ctx context.Context, request *WebhookListRequest) (*WebhookListResponse, error) {
	response := &WebhookListResponse{}
	if err := a.client.Call(ctx, "/webhook/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Deletes an event notification handler
//
// You can also temporarily deactivate webhook without deleting it by setting it's status to "disabled"
// in [webhook/set](https://docs.unione.io/en/web-api-ref#webhook-set) method.
//
// Documentation: https://docs.unione.io/en/web-api-ref#webhook-delete
func (a *WebhookAPI) Delete(ctx context.Context, request *WebhookDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/webhook/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ParseCallback — Webhook callback format description.
func ParseCallback(body []byte) (*WebhookCallback, error) {
	callback := &WebhookCallback{}
	if err := json.Unmarshal(body, callback); err != nil {
		return nil, err
	}
	return callback, nil
}
