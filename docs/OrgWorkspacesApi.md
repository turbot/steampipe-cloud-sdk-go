# \OrgWorkspacesApi

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWorkspace**](OrgWorkspacesApi.md#CreateOrgWorkspace) | **Post** /org/{org_handle}/workspace | Create org workspace
[**DeleteOrgWorkspace**](OrgWorkspacesApi.md#DeleteOrgWorkspace) | **Delete** /org/{org_handle}/workspace/{workspace_handle} | Delete org workspace
[**GetOrgWorkspace**](OrgWorkspacesApi.md#GetOrgWorkspace) | **Get** /org/{org_handle}/workspace/{workspace_handle} | Get org workspace
[**ListOrgWorkspaces**](OrgWorkspacesApi.md#ListOrgWorkspaces) | **Get** /org/{org_handle}/workspace | List org workspaces
[**OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet**](OrgWorkspacesApi.md#OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet) | **Get** /org/{org_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query org workspace with extensions
[**OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost**](OrgWorkspacesApi.md#OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost) | **Post** /org/{org_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query org workspace with extensions
[**OrgOrgHandleWorkspaceWorkspaceHandleQueryGet**](OrgWorkspacesApi.md#OrgOrgHandleWorkspaceWorkspaceHandleQueryGet) | **Get** /org/{org_handle}/workspace/{workspace_handle}/query | Query org workspace
[**OrgOrgHandleWorkspaceWorkspaceHandleQueryPost**](OrgWorkspacesApi.md#OrgOrgHandleWorkspaceWorkspaceHandleQueryPost) | **Post** /org/{org_handle}/workspace/{workspace_handle}/query | Query org workspace
[**OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet**](OrgWorkspacesApi.md#OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema | Get org workspace schemas
[**UpdateOrgWorkspace**](OrgWorkspacesApi.md#UpdateOrgWorkspace) | **Patch** /org/{org_handle}/workspace/{workspace_handle} | Update org workspace



## CreateOrgWorkspace

> TypesWorkspace CreateOrgWorkspace(ctx, orgHandle).Request(request).Execute()

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
    resp, r, err := api_client.OrgWorkspacesApi.CreateOrgWorkspace(context.Background(), orgHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacesApi.CreateOrgWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgWorkspace`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacesApi.CreateOrgWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where we want to create the workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgWorkspaceRequest struct via the builder pattern


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


## DeleteOrgWorkspace

> TypesWorkspace DeleteOrgWorkspace(ctx, orgHandle, workspaceHandle).Execute()

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
    resp, r, err := api_client.OrgWorkspacesApi.DeleteOrgWorkspace(context.Background(), orgHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacesApi.DeleteOrgWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgWorkspace`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacesApi.DeleteOrgWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the workspace exist. | 
**workspaceHandle** | **string** | Provide the handle of the workspace which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgWorkspaceRequest struct via the builder pattern


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


## GetOrgWorkspace

> TypesWorkspace GetOrgWorkspace(ctx, orgHandle, workspaceHandle).Execute()

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
    resp, r, err := api_client.OrgWorkspacesApi.GetOrgWorkspace(context.Background(), orgHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacesApi.GetOrgWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgWorkspace`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacesApi.GetOrgWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle for an organization where the workspace exist. | 
**workspaceHandle** | **string** | The handle of the workspace whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgWorkspaceRequest struct via the builder pattern


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


## ListOrgWorkspaces

> TypesListWorkspacesResponse ListOrgWorkspaces(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

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
    resp, r, err := api_client.OrgWorkspacesApi.ListOrgWorkspaces(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacesApi.ListOrgWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgWorkspaces`: TypesListWorkspacesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacesApi.ListOrgWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization for which we want to list the workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWorkspacesRequest struct via the builder pattern


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
    resp, r, err := api_client.OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet(context.Background(), orgHandle, workspaceHandle, extensions).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet`: %v\n", resp)
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
    resp, r, err := api_client.OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost(context.Background(), orgHandle, workspaceHandle, extensions).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost`: %v\n", resp)
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
    resp, r, err := api_client.OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryGet(context.Background(), orgHandle, workspaceHandle).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgOrgHandleWorkspaceWorkspaceHandleQueryGet`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryGet`: %v\n", resp)
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
    resp, r, err := api_client.OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryPost(context.Background(), orgHandle, workspaceHandle).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgOrgHandleWorkspaceWorkspaceHandleQueryPost`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleQueryPost`: %v\n", resp)
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
    resp, r, err := api_client.OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet(context.Background(), orgHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet`: QueryWorkspaceSchema
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacesApi.OrgOrgHandleWorkspaceWorkspaceHandleSchemaGet`: %v\n", resp)
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


## UpdateOrgWorkspace

> TypesWorkspace UpdateOrgWorkspace(ctx, orgHandle, workspaceHandle).Request(request).Execute()

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
    resp, r, err := api_client.OrgWorkspacesApi.UpdateOrgWorkspace(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacesApi.UpdateOrgWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgWorkspace`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacesApi.UpdateOrgWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the workspace exist. | 
**workspaceHandle** | **string** | The handle of the workspace which needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgWorkspaceRequest struct via the builder pattern


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

