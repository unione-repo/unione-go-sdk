// SystemAPI: API methods of the System section.

package system

import (
	"context"

	"github.com/unione-repo/unione-go-sdk/client"
)

// SystemAPI — System operations.
type SystemAPI struct {
	client *client.APIClient
}

// NewSystemAPI creates the System API.
func NewSystemAPI(c *client.APIClient) *SystemAPI {
	return &SystemAPI{client: c}
}

// Ping — Checks API key validity and returns user ID.
//
// Documentation: https://docs.unione.io/en/web-api-ref#system-ping
func (a *SystemAPI) Ping(ctx context.Context) (*SystemPingResponse, error) {
	response := &SystemPingResponse{}
	if err := a.client.Call(ctx, "/system/ping.json", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// Info — Gets user or project info by API key.
//
// Documentation: https://docs.unione.io/en/web-api-ref#system-info
func (a *SystemAPI) Info(ctx context.Context) (*SystemInfoResponse, error) {
	response := &SystemInfoResponse{}
	if err := a.client.Call(ctx, "/system/info.json", nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
