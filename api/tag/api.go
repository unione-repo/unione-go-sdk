// TagAPI: API methods of the Tag section.

package tag

import (
	"context"

	"github.com/unione-repo/unione-go-sdk/client"
	"github.com/unione-repo/unione-go-sdk/common"
)

// TagAPI — This family of methods is used to manage user-defined tags. Tags are used to categorize letters
// based on your own criteria, e. g. to create a distinction between marketing and transactional
// letters or group together all notifications for a particular order type. UniOne allows you to assign
// tags to all outgoing emails, and the tags you’ve set may then be used to filter and analyze delivery
// statistics.
type TagAPI struct {
	client *client.APIClient
}

// NewTagAPI creates the Tag API.
func NewTagAPI(c *client.APIClient) *TagAPI {
	return &TagAPI{client: c}
}

// List — This method returns the full list of user-defined tags. Does not receive any parameters.
//
// Documentation: https://docs.unione.io/en/web-api-ref#tag-list
func (a *TagAPI) List(ctx context.Context) (*TagListResponse, error) {
	response := &TagListResponse{}
	if err := a.client.Call(ctx, "/tag/list.json", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete — Removes tag.
//
// Documentation: https://docs.unione.io/en/web-api-ref#tag-delete
func (a *TagAPI) Delete(ctx context.Context, request *TagDeleteRequest) (*common.SuccessResponse, error) {
	response := &common.SuccessResponse{}
	if err := a.client.Call(ctx, "/tag/delete.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
