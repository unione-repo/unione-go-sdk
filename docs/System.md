# System

[Back to the list of sections](../README.md)

System operations.

`system.SystemAPI`, `client.System`

| Method | Endpoint | Description |
|---|---|---|
| [`Ping`](#ping) | `POST /system/ping.json` | Checks API key validity and returns user ID. |
| [`Info`](#info) | `POST /system/info.json` | Gets user or project info by API key. |

## Ping

`POST /system/ping.json` · [Documentation](https://docs.unione.io/en/web-api-ref#system-ping)

Checks API key validity and returns user ID.

### Request

The method has no parameters.

### Response: `SystemPingResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `UserID` | `int` | Unique user identifier. |

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

	response, err := client.System.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.UserID)
}
```

## Info

`POST /system/info.json` · [Documentation](https://docs.unione.io/en/web-api-ref#system-info)

Gets user or project info by API key.

### Request

The method has no parameters.

### Response: `SystemInfoResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `UserID` | `int` | Unique user identifier. |
| `Email` | `string` | Email of the user. |
| `ProjectID` | `string` | Unqiue project identifier, ASCII string up to 36 characters long. Present only if the API key used for request is the project API key. |
| `ProjectName` | `string` | Project name, unique for user account. Present only if the API key used for request is the project API key. |
| `ProjectAccounting` | `system.ProjectAccountingObject` | Object with email counter properties. Present only if the API key used for request is the user API key. |
| `Accounting` | `system.SystemInfoResponseAccounting` | Object with accounting period properties. Present only if the API key used for request is the user API key. |

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

	response, err := client.System.Info(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.UserID)
}
```
