# Project

[Back to the list of sections](../README.md)

Projects allow you to group several independent domains, webhooks, templates, and unsubscription lists together. It's very useful if you have your own clients and manage their data from a single UniOne account. Each project receives its own API key, and all API calls with that API key only deal with the data of that project. This way, you can assign a separate project for each client and be sure that they won't conflict; e.g. unsubscribing from one project won't affect another.

Projects methods are disabled by default due to antispam measures. To enable them, [contact support](https://cp.unione.io/en/support).

Please note that sending domains are confirmed and used on per-project basis by default. But it's possible to allow projects to send emails from domains confirmed for main account - just [ask support](https://cp.unione.io/en/support) to enable it.

`project.ProjectAPI`, `client.Project`

| Method | Endpoint | Description |
|---|---|---|
| [`Create`](#create) | `POST /project/create.json` | Creates a new project. |
| [`Update`](#update) | `POST /project/update.json` | Updates a project. |
| [`List`](#list) | `POST /project/list.json` | Lists projects. |
| [`Delete`](#delete) | `POST /project/delete.json` | Deletes a project. |

## Create

`POST /project/create.json` · [Documentation](https://docs.unione.io/en/web-api-ref#project-create)

Creates a new project.

### Request: `ProjectCreateRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `Project` | `project.ProjectObject` | yes | Object with project properties. |

### Response: `ProjectCreateResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `ProjectID` | `string` | Unqiue project identifier, ASCII string up to 36 characters long. |
| `ProjectAPIKey` | `string` | API key of the project. You can use it instead of the user API key parameter in all methods except project/* methods. |

### Example

```go
package main

import (
	"context"
	"fmt"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/project"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.Project.Create(context.Background(), &project.ProjectCreateRequest{
		Project: &project.ProjectObject{
			Name:              "Project 1A",
			Country:           unione.String("ES"),
			BackendDomainID:   unione.Int(1234),
			UnsubscribePageID: unione.Int(1),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.ProjectID)
}
```

## Update

`POST /project/update.json` · [Documentation](https://docs.unione.io/en/web-api-ref#project-update)

Updates a project.

### Request: `ProjectUpdateRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `ProjectID` | `string` |  | Unique project identifier. You should identify project to update either with project_id or with project_api_key field. But we recommend usage of project_id as a more secure option. |
| `ProjectAPIKey` | `string` |  | Optional project API key. You should identify project to update either with project_id or with project_api_key field. But we recommend usage of project_id as a more secure option. |
| `Project` | `project.ProjectUpdateObject` | yes | Object with project properties. |

### Response: `ProjectUpdateResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `ProjectAPIKey` | `string` | API key of the project. Returned only if the project to update was identified by project_api_key. Not returned in case of updating by project_id. |

### Example

```go
package main

import (
	"context"
	"fmt"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/project"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.Project.Update(context.Background(), &project.ProjectUpdateRequest{
		ProjectID:     unione.String("1234567890"),
		ProjectAPIKey: unione.String("your api_key"),
		Project: &project.ProjectUpdateObject{
			Name:              "Project 1A",
			Country:           unione.String("ES"),
			BackendDomainID:   unione.Int(1234),
			UnsubscribePageID: unione.Int(1),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.ProjectAPIKey)
}
```

## List

`POST /project/list.json` · [Documentation](https://docs.unione.io/en/web-api-ref#project-list)

Lists projects.

### Request: `ProjectListRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `ProjectID` | `string` |  | Optional possibility to specify project_id to include only this project to list. |
| `ProjectAPIKey` | `string` |  | Optional possibility to specify project_api_key to include only this project to list. |

### Response: `ProjectListResponse`

| Field | Type | Description |
|---|---|---|
| `Status` | `string` | "success" string (`success`) |
| `Projects` | `[]project.ProjectListObject` | Array of objects where each object represents one project. |

### Example

```go
package main

import (
	"context"
	"fmt"

	unione "github.com/unione-repo/unione-go-sdk"
	"github.com/unione-repo/unione-go-sdk/api/project"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.Project.List(context.Background(), &project.ProjectListRequest{
		ProjectID:     unione.String("1234567890"),
		ProjectAPIKey: unione.String("your api_key"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Projects)
}
```

## Delete

`POST /project/delete.json` · [Documentation](https://docs.unione.io/en/web-api-ref#project-delete)

Deletes a project.

The project will be moved to the "trash bin" before being permanently deleted. You can restore the project by contacting support.

### Request: `ProjectDeleteRequest`

| Field | Type | Required | Description |
|---|---|---|---|
| `ProjectID` | `string` |  | Unique project identifier. You should identify project to delete either with project_id or with project_api_key field. We recommend usage of project_id as a more secure option. |
| `ProjectAPIKey` | `string` |  | Optional project API key. You should identify project to delete either with project_id or with project_api_key field. We recommend usage of project_id as a more secure option. |

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
	"github.com/unione-repo/unione-go-sdk/api/project"
)

func main() {
	client := unione.NewClient("YOUR_API_KEY")

	response, err := client.Project.Delete(context.Background(), &project.ProjectDeleteRequest{
		ProjectID:     unione.String("1234567890"),
		ProjectAPIKey: unione.String("your api_key"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
```
