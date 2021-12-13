# \OrgWorkspaces

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgWorkspaces.md#Create) | **Post** /org/{org_handle}/workspace | Create org workspace
[**Delete**](OrgWorkspaces.md#Delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle} | Delete org workspace
[**Get**](OrgWorkspaces.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle} | Get org workspace
[**List**](OrgWorkspaces.md#List) | **Get** /org/{org_handle}/workspace | List org workspaces
[**ListDBLogs**](OrgWorkspaces.md#ListDBLogs) | **Get** /org/{org_handle}/workspace/{workspace_handle}/logs | Org workspace logs
[**OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet**](OrgWorkspaces.md#OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet) | **Get** /org/{org_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query org workspace with extensions
[**OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost**](OrgWorkspaces.md#OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost) | **Post** /org/{org_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query org workspace with extensions
[**OrgOrgHandleWorkspaceWorkspaceHandleQueryGet**](OrgWorkspaces.md#OrgOrgHandleWorkspaceWorkspaceHandleQueryGet) | **Get** /org/{org_handle}/workspace/{workspace_handle}/query | Query org workspace
[**OrgOrgHandleWorkspaceWorkspaceHandleQueryPost**](OrgWorkspaces.md#OrgOrgHandleWorkspaceWorkspaceHandleQueryPost) | **Post** /org/{org_handle}/workspace/{workspace_handle}/query | Query org workspace
[**OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet**](OrgWorkspaces.md#OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema | Get org workspace schemas
[**Update**](OrgWorkspaces.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle} | Update org workspace



## Create

> TypesWorkspace Create(ctx, orgHandle).Request(request).Execute()

Create org workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the organization where we want to create the workspace.
    request := *openapiclient.NewTypesCreateWorkspaceRequest("Handle_example") // TypesCreateWorkspaceRequest | The request body for the workspace to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.Create(context.Background(), orgHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where we want to create the workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**TypesCreateWorkspaceRequest**](TypesCreateWorkspaceRequest.md) | The request body for the workspace to be created. | 

### Return type

[**TypesWorkspace**](TypesWorkspace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> TypesWorkspace Delete(ctx, orgHandle, workspaceHandle).Execute()

Delete org workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the organization where the workspace exist.
    workspaceHandle := "workspaceHandle_example" // string | Provide the handle of the workspace which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.Delete(context.Background(), orgHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the workspace exist. | 
**workspaceHandle** | **string** | Provide the handle of the workspace which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesWorkspace**](TypesWorkspace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> TypesWorkspace Get(ctx, orgHandle, workspaceHandle).Execute()

Get org workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle for an organization where the workspace exist.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.Get(context.Background(), orgHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle for an organization where the workspace exist. | 
**workspaceHandle** | **string** | The handle of the workspace whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesWorkspace**](TypesWorkspace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> TypesListWorkspacesResponse List(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspaces



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the organization for which we want to list the workspace.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.List(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: TypesListWorkspacesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization for which we want to list the workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListWorkspacesResponse**](TypesListWorkspacesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDBLogs

> TypesListLogsResponse ListDBLogs(ctx, orgHandle, workspaceHandle).Execute()

Org workspace logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | Specify the org handle to get the workspace logs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose logs needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.ListDBLogs(context.Background(), orgHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.ListDBLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDBLogs`: TypesListLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.ListDBLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the org handle to get the workspace logs. | 
**workspaceHandle** | **string** | The handle of the workspace whose logs needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDBLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesListLogsResponse**](TypesListLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet

> TypesWorkspaceQueryResult OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet(ctx, orgHandle, workspaceHandle, extensions).Sql(sql).ContentType(contentType).Execute()

Query org workspace with extensions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the org that the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to query.
    extensions := "extensions_example" // string | The content type for the request. E.g. 
    sql := "sql_example" // string | The sql query to perform against the org workspace.
    contentType := "contentType_example" // string | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet(context.Background(), orgHandle, workspaceHandle, extensions).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to query. | 
**extensions** | **string** | The content type for the request. E.g.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sql** | **string** | The sql query to perform against the org workspace. | 
 **contentType** | **string** | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. | 

### Return type

[**TypesWorkspaceQueryResult**](TypesWorkspaceQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, text/markdown

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost

> TypesWorkspaceQueryResult OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost(ctx, orgHandle, workspaceHandle, extensions).Sql(sql).ContentType(contentType).Execute()

Query org workspace with extensions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the org that the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to query.
    extensions := "extensions_example" // string | The content type for the request. E.g. 
    sql := "sql_example" // string | The sql query to perform against the org workspace.
    contentType := "contentType_example" // string | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost(context.Background(), orgHandle, workspaceHandle, extensions).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to query. | 
**extensions** | **string** | The content type for the request. E.g.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sql** | **string** | The sql query to perform against the org workspace. | 
 **contentType** | **string** | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. | 

### Return type

[**TypesWorkspaceQueryResult**](TypesWorkspaceQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, text/markdown

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgOrgHandleWorkspaceWorkspaceHandleQueryGet

> TypesWorkspaceQueryResult OrgOrgHandleWorkspaceWorkspaceHandleQueryGet(ctx, orgHandle, workspaceHandle).Sql(sql).ContentType(contentType).Execute()

Query org workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the org that the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to query.
    sql := "sql_example" // string | The sql query to perform against the org workspace.
    contentType := "contentType_example" // string | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryGet(context.Background(), orgHandle, workspaceHandle).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgOrgHandleWorkspaceWorkspaceHandleQueryGet`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgOrgHandleWorkspaceWorkspaceHandleQueryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sql** | **string** | The sql query to perform against the org workspace. | 
 **contentType** | **string** | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. | 

### Return type

[**TypesWorkspaceQueryResult**](TypesWorkspaceQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, text/markdown

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgOrgHandleWorkspaceWorkspaceHandleQueryPost

> TypesWorkspaceQueryResult OrgOrgHandleWorkspaceWorkspaceHandleQueryPost(ctx, orgHandle, workspaceHandle).Sql(sql).ContentType(contentType).Execute()

Query org workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the org that the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to query.
    sql := "sql_example" // string | The sql query to perform against the org workspace.
    contentType := "contentType_example" // string | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryPost(context.Background(), orgHandle, workspaceHandle).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgOrgHandleWorkspaceWorkspaceHandleQueryPost`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleQueryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgOrgHandleWorkspaceWorkspaceHandleQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sql** | **string** | The sql query to perform against the org workspace. | 
 **contentType** | **string** | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. | 

### Return type

[**TypesWorkspaceQueryResult**](TypesWorkspaceQueryResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/csv, text/markdown
- **Accept**: application/json, text/csv, text/markdown

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet

> QueryWorkspaceSchema OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet(ctx, orgHandle, workspaceHandle).Execute()

Get org workspace schemas



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the org that the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to get the schemas for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet(context.Background(), orgHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet`: QueryWorkspaceSchema
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to get the schemas for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgOrgHandleWorkspaceWorkspaceHandleSchemaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**QueryWorkspaceSchema**](QueryWorkspaceSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> TypesWorkspace Update(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Update org workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the organization where the workspace exist.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace which needs to be updated.
    request := *openapiclient.NewTypesUpdateWorkspaceRequest() // TypesUpdateWorkspaceRequest | The request body of the workspace which needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaces.Update(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaces.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaces.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the workspace exist. | 
**workspaceHandle** | **string** | The handle of the workspace which needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**TypesUpdateWorkspaceRequest**](TypesUpdateWorkspaceRequest.md) | The request body of the workspace which needs to be updated. | 

### Return type

[**TypesWorkspace**](TypesWorkspace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

