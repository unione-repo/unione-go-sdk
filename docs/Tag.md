# Tag

[Back to the list of sections](../README.md)

This family of methods is used to manage user-defined tags. Tags are used to categorize letters based on your own criteria, e. g. to create a distinction between marketing and transactional letters or group together all notifications for a particular order type. UniOne allows you to assign tags to all outgoing emails, and the tags you’ve set may then be used to filter and analyze delivery statistics.

`tag.TagAPI`, `client.Tag`

| Method | Endpoint | Description |
|---|---|---|
| [`List`](#list) | `POST /tag/list.json` | This method returns the full list of user-defined tags. Does not receive any parameters. |
| [`Delete`](#delete) | `POST /tag/delete.json` | Removes tag. |

## List

`POST /tag/list.json` · [Documentation](https://docs.unione.io/en/web-api-ref#tag-list)

This method returns the full list of user-defined tags. Does not receive any parameters.

### Request

The method has no parameters.

### Response: `TagListResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `Tags` | `[]tag.TagsResponseObject` | An array of objects containing data for user-defined tags. |

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

	response, err := client.Tag.List(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Tags)
}
```

## Delete

`POST /tag/delete.json` · [Documentation](https://docs.unione.io/en/web-api-ref#tag-delete)

Removes tag.

### Request: `TagDeleteRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `TagID` | `int` | yes | Unique tag id. |

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
	"github.com/unione-repo/unione-go-sdk/api/tag"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.Tag.Delete(context.Background(), &tag.TagDeleteRequest{
		TagID: 54,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
```
