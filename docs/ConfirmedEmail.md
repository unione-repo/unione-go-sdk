# ConfirmedEmail

[Back to the list of sections](../README.md)

Manage email addresses confirmed for sandbox-domain sending. These methods are available only with a main account API key.

`confirmedemail.ConfirmedEmailAPI`, `client.ConfirmedEmail`

| Method | Endpoint | Description |
|---|---|---|
| [`Create`](#create) | `POST /confirmed-email/create.json` | Adds an email address and sends a confirmation link. |
| [`List`](#list) | `POST /confirmed-email/list.json` | Lists confirmed and unconfirmed email addresses for the account. |
| [`Delete`](#delete) | `POST /confirmed-email/delete.json` | Deletes an email address from the account's confirmed email list. |

## Create

`POST /confirmed-email/create.json` · [Documentation](https://docs.unione.io/en/web-api-ref#confirmed-email-create)

Adds an email address and sends a confirmation link.

### Request: `ConfirmedEmailRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `Email` | `string` | yes | Email address to add or delete. |

### Response: `ConfirmedEmailCreateResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `ConfirmedEmail` | `confirmedemail.ConfirmedEmail` |  |

### Example

```go
package main

import (
	"context"
	"fmt"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/confirmedemail"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.ConfirmedEmail.Create(context.Background(), &confirmedemail.ConfirmedEmailRequest{
		Email: "dev@example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.ConfirmedEmail)
}
```

## List

`POST /confirmed-email/list.json` · [Documentation](https://docs.unione.io/en/web-api-ref#confirmed-email-list)

Lists confirmed and unconfirmed email addresses for the account.

### Request

The method has no parameters.

### Response: `ConfirmedEmailListResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `ConfirmedEmails` | `[]confirmedemail.ConfirmedEmail` |  |

### Example

```go
package main

import (
	"context"
	"fmt"

	unione "github.com/unione-repo/unione-go-sdk"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.ConfirmedEmail.List(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.ConfirmedEmails)
}
```

## Delete

`POST /confirmed-email/delete.json` · [Documentation](https://docs.unione.io/en/web-api-ref#confirmed-email-delete)

Deletes an email address from the account's confirmed email list.

### Request: `ConfirmedEmailRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `Email` | `string` | yes | Email address to add or delete. |

### Response: `ConfirmedEmailDeleteResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `Message` | `string` |  |

### Example

```go
package main

import (
	"context"
	"fmt"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/confirmedemail"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.ConfirmedEmail.Delete(context.Background(), &confirmedemail.ConfirmedEmailRequest{
		Email: "dev@example.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Message)
}
```
