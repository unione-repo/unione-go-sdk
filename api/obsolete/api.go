// ObsoleteAPI: API methods of the Obsolete section.

package obsolete

import (
	"context"

	"github.com/unione-repo/unione-go-sdk/client"
)

// ObsoleteAPI — Obsolete methods dealing with unsubscribes. We recommend using the more flexible
// [suppression](https://docs.unione.io/en/web-api-ref#suppression) methods instead. Note that you
// can't resubscribe an unsubscribed recipient using WebAPI, but you can send him an email with
// subscription link using email/subscribe method.
type ObsoleteAPI struct {
	client *client.APIClient
}

// NewObsoleteAPI creates the Obsolete API.
func NewObsoleteAPI(c *client.APIClient) *ObsoleteAPI {
	return &ObsoleteAPI{client: c}
}

// UnsubscribedSet — Registers a new unsubscribed email.
//
// After registering an email as unsubscribed all future attempts to send a message to this email will
// be rejected. Please note that there's no way to remove an email from unsubscribed thru the API. But
// you can send an email with a link to resubscribe using
// [email/subscribe](https://docs.unione.io/en/web-api-ref#email-subscribe) method or remove
// "unsubscribed" status manually on [Email Search page](https://cp.unione.io/en/user/email) in control
// panel (but only a limited number of unsubscribed email per day can be removed).
//
// Documentation: https://docs.unione.io/en/web-api-ref#unsubscribed-set
func (a *ObsoleteAPI) UnsubscribedSet(ctx context.Context, request *ObsoleteUnsubscribedSetRequest) (*ObsoleteUnsubscribedSetResponse, error) {
	response := &ObsoleteUnsubscribedSetResponse{}
	if err := a.client.Call(ctx, "/unsubscribed/set.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// UnsubscribedCheck — Checks if an email is unsubscribed.
//
// Documentation: https://docs.unione.io/en/web-api-ref#unsubscribed-check
func (a *ObsoleteAPI) UnsubscribedCheck(ctx context.Context, request *ObsoleteUnsubscribedCheckRequest) (*ObsoleteUnsubscribedCheckResponse, error) {
	response := &ObsoleteUnsubscribedCheckResponse{}
	if err := a.client.Call(ctx, "/unsubscribed/check.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// UnsubscribedList — Returns a list of all unsubscribed emails since provided date.
//
// Documentation: https://docs.unione.io/en/web-api-ref#unsubscribed-list
func (a *ObsoleteAPI) UnsubscribedList(ctx context.Context, request *ObsoleteUnsubscribedListRequest) (*ObsoleteUnsubscribedListResponse, error) {
	response := &ObsoleteUnsubscribedListResponse{}
	if err := a.client.Call(ctx, "/unsubscribed/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
