// SuppressionAPI: API methods of the Suppression section.

package suppression

import (
	"context"

	"github.com/unione-repo/unione-go-sdk/client"
	"github.com/unione-repo/unione-go-sdk/common"
)

// SuppressionAPI — Methods for managing suppression lists. An email can be added to suppression list because of
// unsubscription, bounce or complain.
type SuppressionAPI struct {
	client *client.APIClient
}

// NewSuppressionAPI creates the Suppression API.
func NewSuppressionAPI(c *client.APIClient) *SuppressionAPI {
	return &SuppressionAPI{client: c}
}

// Set — Adds an email address to the suppression list. You can always remove this address from the
// suppression list later using the
// [suppression/delete](https://docs.unione.io/en/web-api-ref#suppression-delete) method.
//
// Documentation: https://docs.unione.io/en/web-api-ref#suppression-set
func (a *SuppressionAPI) Set(ctx context.Context, request *SuppressionSetRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/suppression/set.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Get — Gets a reason and date of email suppression.
//
// Documentation: https://docs.unione.io/en/web-api-ref#suppression-get
func (a *SuppressionAPI) Get(ctx context.Context, request *SuppressionGetRequest) (*SuppressionGetResponse, error) {
	response := &SuppressionGetResponse{}
	if err := a.client.Call(ctx, "/suppression/get.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — Returns a suppression list since provided date.
//
// Documentation: https://docs.unione.io/en/web-api-ref#suppression-list
func (a *SuppressionAPI) List(ctx context.Context, request *SuppressionListRequest) (*SuppressionListResponse, error) {
	response := &SuppressionListResponse{}
	if err := a.client.Call(ctx, "/suppression/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Deletes an email from suppression list (only records with is_deletable=true). If there are no
// entries of this email in the suppression list, API error 3003 is returned. If such entries exist but
// no one can be deleted, API error 3004 is returned. If daily limit of deletes exceeded, API error 906
// is returned.
//
// Documentation: https://docs.unione.io/en/web-api-ref#suppression-delete
func (a *SuppressionAPI) Delete(ctx context.Context, request *SuppressionDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/suppression/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
