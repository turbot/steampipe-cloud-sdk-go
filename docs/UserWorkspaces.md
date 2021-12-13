# \UserWorkspaces

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserWorkspaces.md#Create) | **Post** /user/{user_handle}/workspace | Create user workspace
[**Delete**](UserWorkspaces.md#Delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle} | Delete user workspace
[**Get**](UserWorkspaces.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle} | Get user workspace
[**List**](UserWorkspaces.md#List) | **Get** /user/{user_handle}/workspace | List user workspaces
[**ListDBLogs**](UserWorkspaces.md#ListDBLogs) | **Get** /user/{user_handle}/workspace/{workspace_handle}/logs | User workspace logs
[**Update**](UserWorkspaces.md#Update) | **Patch** /user/{user_handle}/workspace/{workspace_handle} | Update user workspace
[**UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet**](UserWorkspaces.md#UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet) | **Get** /user/{user_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query user workspace with extensions
[**UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost**](UserWorkspaces.md#UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost) | **Post** /user/{user_handle}/workspace/{workspace_handle}/query/data.{extensions} | Query user workspace with extensions
[**UserUserHandleWorkspaceWorkspaceHandleQueryGet**](UserWorkspaces.md#UserUserHandleWorkspaceWorkspaceHandleQueryGet) | **Get** /user/{user_handle}/workspace/{workspace_handle}/query | Query user workspace
[**UserUserHandleWorkspaceWorkspaceHandleQueryPost**](UserWorkspaces.md#UserUserHandleWorkspaceWorkspaceHandleQueryPost) | **Post** /user/{user_handle}/workspace/{workspace_handle}/query | Query user workspace
[**UserUserHandleWorkspaceWorkspaceHandleSchemaGet**](UserWorkspaces.md#UserUserHandleWorkspaceWorkspaceHandleSchemaGet) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema | Get user workspace schemas



## Create

> TypesWorkspace Create(ctx, userHandle).Request(request).Execute()

Create user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user where we want to create the workspace.
    request := *openapiclient.NewTypesCreateWorkspaceRequest("Handle_example") // TypesCreateWorkspaceRequest | The request body for the workspace to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.Create(context.Background(), userHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where we want to create the workspace. | 

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

> TypesWorkspace Delete(ctx, userHandle, workspaceHandle).Execute()

Delete user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user where the workspace exist.
    workspaceHandle := "workspaceHandle_example" // string | Provide the handle of the workspace which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.Delete(context.Background(), userHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the workspace exist. | 
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

> TypesWorkspace Get(ctx, userHandle, workspaceHandle).Execute()

Get user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user where the workspace exist.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.Get(context.Background(), userHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the workspace exist. | 
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

> TypesListWorkspacesResponse List(ctx, userHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspaces



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
    userHandle := "userHandle_example" // string | The handle of the user for which we want to list the workspace.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.List(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: TypesListWorkspacesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which we want to list the workspace. | 

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

> TypesListLogsResponse ListDBLogs(ctx, userHandle, workspaceHandle).Execute()

User workspace logs



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
    userHandle := "userHandle_example" // string | Specify the user handle to get the workspace logs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose logs needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.ListDBLogs(context.Background(), userHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.ListDBLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDBLogs`: TypesListLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.ListDBLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the user handle to get the workspace logs. | 
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


## Update

> TypesWorkspace Update(ctx, userHandle, workspaceHandle).Request(request).Execute()

Update user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user where the workspace exist.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace which needs to be updated.
    request := *openapiclient.NewTypesUpdateWorkspaceRequest() // TypesUpdateWorkspaceRequest | The request body for the workspace which needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.Update(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: TypesWorkspace
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the workspace exist. | 
**workspaceHandle** | **string** | The handle of the workspace which needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**TypesUpdateWorkspaceRequest**](TypesUpdateWorkspaceRequest.md) | The request body for the workspace which needs to be updated. | 

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


## UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet

> TypesWorkspaceQueryResult UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet(ctx, userHandle, workspaceHandle, extensions).Sql(sql).ContentType(contentType).Execute()

Query user workspace with extensions



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
    userHandle := "userHandle_example" // string | The handle of the user that the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to query.
    extensions := "extensions_example" // string | The content type for the request. E.g. 
    sql := "sql_example" // string | The sql query to perform against the user workspace.
    contentType := "contentType_example" // string | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet(context.Background(), userHandle, workspaceHandle, extensions).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to query. | 
**extensions** | **string** | The content type for the request. E.g.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sql** | **string** | The sql query to perform against the user workspace. | 
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


## UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost

> TypesWorkspaceQueryResult UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost(ctx, userHandle, workspaceHandle, extensions).Sql(sql).ContentType(contentType).Execute()

Query user workspace with extensions



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
    userHandle := "userHandle_example" // string | The handle of the user that the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to query.
    extensions := "extensions_example" // string | The content type for the request. E.g. 
    sql := "sql_example" // string | The sql query to perform against the user workspace.
    contentType := "contentType_example" // string | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost(context.Background(), userHandle, workspaceHandle, extensions).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to query. | 
**extensions** | **string** | The content type for the request. E.g.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserHandleWorkspaceWorkspaceHandleQueryDataExtensionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sql** | **string** | The sql query to perform against the user workspace. | 
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


## UserUserHandleWorkspaceWorkspaceHandleQueryGet

> TypesWorkspaceQueryResult UserUserHandleWorkspaceWorkspaceHandleQueryGet(ctx, userHandle, workspaceHandle).Sql(sql).ContentType(contentType).Execute()

Query user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user that the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to query.
    sql := "sql_example" // string | The sql query to perform against the user workspace.
    contentType := "contentType_example" // string | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryGet(context.Background(), userHandle, workspaceHandle).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserUserHandleWorkspaceWorkspaceHandleQueryGet`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserHandleWorkspaceWorkspaceHandleQueryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sql** | **string** | The sql query to perform against the user workspace. | 
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


## UserUserHandleWorkspaceWorkspaceHandleQueryPost

> TypesWorkspaceQueryResult UserUserHandleWorkspaceWorkspaceHandleQueryPost(ctx, userHandle, workspaceHandle).Sql(sql).ContentType(contentType).Execute()

Query user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user that the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to query.
    sql := "sql_example" // string | The sql query to perform against the user workspace.
    contentType := "contentType_example" // string | The required content type for the response. Defaults to application/json. Supported values are json, application/json, csv, text/csv, md and text/markdown. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryPost(context.Background(), userHandle, workspaceHandle).Sql(sql).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserUserHandleWorkspaceWorkspaceHandleQueryPost`: TypesWorkspaceQueryResult
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleQueryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserHandleWorkspaceWorkspaceHandleQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sql** | **string** | The sql query to perform against the user workspace. | 
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


## UserUserHandleWorkspaceWorkspaceHandleSchemaGet

> QueryWorkspaceSchema UserUserHandleWorkspaceWorkspaceHandleSchemaGet(ctx, userHandle, workspaceHandle).Execute()

Get user workspace schemas



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
    userHandle := "userHandle_example" // string | The handle of the user that the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to get the schemas for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleSchemaGet(context.Background(), userHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleSchemaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserUserHandleWorkspaceWorkspaceHandleSchemaGet`: QueryWorkspaceSchema
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaces.UserUserHandleWorkspaceWorkspaceHandleSchemaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to get the schemas for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUserHandleWorkspaceWorkspaceHandleSchemaGetRequest struct via the builder pattern


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

