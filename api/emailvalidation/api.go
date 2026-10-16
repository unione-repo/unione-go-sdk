// EmailValidationAPI: API methods of the EmailValidation section.

package emailvalidation

import (
	"context"

	"github.com/unione-repo/unione-go-sdk/client"
)

// EmailValidationAPI — Email validation methods. As a result of the check, you’ll get detailed information on email
// validity. It will help eliminate nonexistent emails, decrease the number of complaints, clean up the
// subscribers list and eventually enhance your sender reputation score. For example, you may add
// validation to a subscription form on your site to let users know they’ve just entered an invalid
// address.
type EmailValidationAPI struct {
	client *client.APIClient
}

// NewEmailValidationAPI creates the EmailValidation API.
func NewEmailValidationAPI(c *client.APIClient) *EmailValidationAPI {
	return &EmailValidationAPI{client: c}
}

// Single — A maximum of two concurrent requests are allowed; if you send more requests, an error will be
// returned. Since the verification process should be completed promptly, we've set a time limit on
// email check. If we encounter a slow email server that does not respond in 5 seconds, you may get an
// “unknown” status.
//
// Documentation: https://docs.unione.io/en/web-api-ref#email-validation-single
func (a *EmailValidationAPI) Single(ctx context.Context, request *EmailValidationSingleRequest) (*EmailValidationSingleResponse, error) {
	response := &EmailValidationSingleResponse{}
	if err := a.client.Call(ctx, "/email-validation/single.json", request, response); err != nil {
		return nil, err
	}
	return response, nil
}
