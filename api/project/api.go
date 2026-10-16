// ProjectAPI: API methods of the Project section.

package project

import (
	"context"

	"github.com/unione-repo/unione-go-sdk/client"
	"github.com/unione-repo/unione-go-sdk/common"
)

// ProjectAPI — Projects allow you to group several independent domains, webhooks, templates, and unsubscription
// lists together. It's very useful if you have your own clients and manage their data from a single
// UniOne account. Each project receives its own API key, and all API calls with that API key only deal
// with the data of that project. This way, you can assign a separate project for each client and be
// sure that they won't conflict; e.g. unsubscribing from one project won't affect another.
//
// Projects methods are disabled by default due to antispam measures. To enable them, [contact
// support](https://cp.unione.io/en/support).
//
// Please note that sending domains are confirmed and used on per-project basis by default. But it's
// possible to allow projects to send emails from domains confirmed for main account - just [ask
// support](https://cp.unione.io/en/support) to enable it.
type ProjectAPI struct {
	client *client.APIClient
}

// NewProjectAPI creates the Project API.
func NewProjectAPI(c *client.APIClient) *ProjectAPI {
	return &ProjectAPI{client: c}
}

// Create — Creates a new project.
//
// Documentation: https://docs.unione.io/en/web-api-ref#project-create
func (a *ProjectAPI) Create(ctx context.Context, request *ProjectCreateRequest) (*ProjectCreateResponse, error) {
	response := &ProjectCreateResponse{}
	if err := a.client.Call(ctx, "/project/create.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Update — Updates a project.
//
// Documentation: https://docs.unione.io/en/web-api-ref#project-update
func (a *ProjectAPI) Update(ctx context.Context, request *ProjectUpdateRequest) (*ProjectUpdateResponse, error) {
	response := &ProjectUpdateResponse{}
	if err := a.client.Call(ctx, "/project/update.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// List — Lists projects.
//
// Documentation: https://docs.unione.io/en/web-api-ref#project-list
func (a *ProjectAPI) List(ctx context.Context, request *ProjectListRequest) (*ProjectListResponse, error) {
	response := &ProjectListResponse{}
	if err := a.client.Call(ctx, "/project/list.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Deletes a project.
//
// The project will be moved to the "trash bin" before being permanently deleted. You can restore the
// project by contacting support.
//
// Documentation: https://docs.unione.io/en/web-api-ref#project-delete
func (a *ProjectAPI) Delete(ctx context.Context, request *ProjectDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/project/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
