# Template

[Back to the list of sections](../README.md)

You have the option of storing a template on the UniOne side and then using it by id, which is highly recommended for sending repetitive messages or high volume traffic. UniOne supports two different [template engines](https://docs.unione.io/en/template-engines).

`template.TemplateAPI`, `client.Template`

| Method | Endpoint | Description |
|---|---|---|
| [`Set`](#set) | `POST /template/set.json` | Creates or updates an email template. |
| [`Get`](#get) | `POST /template/get.json` | Gets template properties by it's id. |
| [`List`](#list) | `POST /template/list.json` | Returns the list of all or some templates of the account. |
| [`Delete`](#delete) | `POST /template/delete.json` | Deletes a template by id. |

## Set

`POST /template/set.json` · [Documentation](https://docs.unione.io/en/web-api-ref#template-set)

Creates or updates an email template.

Restrictions:
* The maximum request size is 10 MB.
* The number of templates is limited to 10000.

### Request: `TemplateSetRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `Template` | `template.TemplateObject` | yes | Object with all the template properties. |

### Response: `TemplateSetResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `Template` | `template.TemplateObject` | Object with all the template properties. |

### Example

```go
package main

import (
	"context"
	"fmt"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/template"
	"github.com/unione-repo/unione-go-sdk/common"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.Template.Set(context.Background(), &template.TemplateSetRequest{
		Template: &template.TemplateObject{
			Name:           "name",
			EditorType:     template.EditorTypeHTML,
			TemplateEngine: common.TemplateEngineVelocity,
			Subject:        unione.String("UniOne test email "),
			FromEmail:      "user@example.com",
			FromName:       unione.String("John Smith"),
			ReplyTo:        unione.String("email@example.com"),
			ReplyToName:    unione.String("John Smith"),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Template)
}
```

## Get

`POST /template/get.json` · [Documentation](https://docs.unione.io/en/web-api-ref#template-get)

Gets template properties by it's id.

### Request: `TemplateGetRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `ID` | `string` | yes | Template id. |

### Response: `TemplateGetResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `Template` | `template.TemplateObject` | Object with all the template properties. |

### Example

```go
package main

import (
	"context"
	"fmt"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/template"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.Template.Get(context.Background(), &template.TemplateGetRequest{
		ID: "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Template)
}
```

## List

`POST /template/list.json` · [Documentation](https://docs.unione.io/en/web-api-ref#template-list)

Returns the list of all or some templates of the account.

You can iterate thru large list using "offset" and "limit" parameters. When less than "limit" templates are returned then the end of list is reached.

### Request: `TemplateListRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `Limit` | `int` |  | Limits the number of templates returned, default is 50. |
| `Offset` | `int` |  | Index of the first template in the list, starting with 0. |

### Response: `TemplateListResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `Templates` | `[]template.TemplateListObject` | Array of template objects. |

### Example

```go
package main

import (
	"context"
	"fmt"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/template"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.Template.List(context.Background(), &template.TemplateListRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Templates)
}
```

## Delete

`POST /template/delete.json` · [Documentation](https://docs.unione.io/en/web-api-ref#template-delete)

Deletes a template by id.

### Request: `TemplateDeleteRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `ID` | `string` | yes | Template id. |

### Response: `SuccessResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |

### Example

```go
package main

import (
	"context"
	"fmt"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/template"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.Template.Delete(context.Background(), &template.TemplateDeleteRequest{
		ID: "00000000-0000-0000-0000-000000000000",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
```
