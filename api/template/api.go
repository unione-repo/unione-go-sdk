// TemplateAPI: API methods of the Template section.

package template

import (
	"context"

	"github.com/unione-repo/unione-go-sdk/client"
	"github.com/unione-repo/unione-go-sdk/common"
)

// TemplateAPI — You have the option of storing a template on the UniOne side and then using it by id, which is
// highly recommended for sending repetitive messages or high volume traffic. UniOne supports two
// different [template engines](https://docs.unione.io/en/template-engines).
type TemplateAPI struct {
	client *client.APIClient
}

// NewTemplateAPI creates the Template API.
func NewTemplateAPI(c *client.APIClient) *TemplateAPI {
	return &TemplateAPI{client: c}
}

// Set — Creates or updates an email template.
//
// Restrictions:
// * The maximum request size is 10 MB.
// * The number of templates is limited to 10000.
//
// Documentation: https://docs.unione.io/en/web-api-ref#template-set
func (a *TemplateAPI) Set(ctx context.Context, request *TemplateSetRequest) (*TemplateSetResponse, error) {
	response := &TemplateSetResponse{}
	if err := a.client.Call(ctx, "/template/set.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Get — Gets template properties by it's id.
//
// Documentation: https://docs.unione.io/en/web-api-ref#template-get
func (a *TemplateAPI) Get(ctx context.Context, request *TemplateGetRequest) (*TemplateGetResponse, error) {
	response := &TemplateGetResponse{}
	if err := a.client.Call(ctx, "/template/get.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — Returns the list of all or some templates of the account.
//
// You can iterate thru large list using "offset" and "limit" parameters. When less than "limit"
// templates are returned then the end of list is reached.
//
// Documentation: https://docs.unione.io/en/web-api-ref#template-list
func (a *TemplateAPI) List(ctx context.Context, request *TemplateListRequest) (*TemplateListResponse, error) {
	response := &TemplateListResponse{}
	if err := a.client.Call(ctx, "/template/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Deletes a template by id.
//
// Documentation: https://docs.unione.io/en/web-api-ref#template-delete
func (a *TemplateAPI) Delete(ctx context.Context, request *TemplateDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/template/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
