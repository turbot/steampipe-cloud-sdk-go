# \UserWorkspaceSchemas

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Attach**](UserWorkspaceSchemas.md#Attach) | **Post** /user/{user_handle}/workspace/{workspace_handle}/schema | Attach a schema to a user workspace
[**Detach**](UserWorkspaceSchemas.md#Detach) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/schema/{schema_name} | Detach a schema from a user workspace
[**Get**](UserWorkspaceSchemas.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema/{schema_name} | Get user workspace schema
[**Get_0**](UserWorkspaceSchemas.md#Get_0) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema/{schema_name}/table | List user workspace schema tables
[**List**](UserWorkspaceSchemas.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/schema | List user workspace schemas



## Attach

> WorkspaceSchema Attach(ctx, userHandle, workspaceHandle).Request(request).Execute()

Attach a schema to a user workspace



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the schema will be attached.
    request := *openapiclient.NewAttachWorkspaceSchemaRequest("ConnectionHandle_example") // AttachWorkspaceSchemaRequest | The request body for the schema to be attached.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSchemas.Attach(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSchemas.Attach``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Attach`: WorkspaceSchema
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSchemas.Attach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace where the schema will be attached. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**AttachWorkspaceSchemaRequest**](AttachWorkspaceSchemaRequest.md) | The request body for the schema to be attached. | 

### Return type

[**WorkspaceSchema**](WorkspaceSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Detach

> WorkspaceSchema Detach(ctx, userHandle, workspaceHandle, schemaName).Execute()

Detach a schema from a user workspace



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace from which the schema will be detached.
    schemaName := "schemaName_example" // string | The name of the schema that will be detached.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSchemas.Detach(context.Background(), userHandle, workspaceHandle, schemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSchemas.Detach``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Detach`: WorkspaceSchema
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSchemas.Detach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace from which the schema will be detached. | 
**schemaName** | **string** | The name of the schema that will be detached. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceSchema**](WorkspaceSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> WorkspaceSchema Get(ctx, userHandle, workspaceHandle, schemaName).Execute()

Get user workspace schema



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the schema exists.
    schemaName := "schemaName_example" // string | The name of the schema whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSchemas.Get(context.Background(), userHandle, workspaceHandle, schemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSchemas.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceSchema
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSchemas.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace where the schema exists. | 
**schemaName** | **string** | The name of the schema whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceSchema**](WorkspaceSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get_0

> ListWorkspaceSchemaTableResponse Get_0(ctx, userHandle, workspaceHandle, schemaName).Execute()

List user workspace schema tables



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the schema exists.
    schemaName := "schemaName_example" // string | The name of the schema whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSchemas.Get_0(context.Background(), userHandle, workspaceHandle, schemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSchemas.Get_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get_0`: ListWorkspaceSchemaTableResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSchemas.Get_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace where the schema exists. | 
**schemaName** | **string** | The name of the schema whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ListWorkspaceSchemaTableResponse**](ListWorkspaceSchemaTableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListWorkspaceSchemaResponse List(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspace schemas



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
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSchemas.List(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSchemas.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListWorkspaceSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSchemas.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace to get the schemas for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListWorkspaceSchemaResponse**](ListWorkspaceSchemaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

