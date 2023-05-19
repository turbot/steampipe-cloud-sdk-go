# \OrgWorkspaceSchemas

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Attach**](OrgWorkspaceSchemas.md#Attach) | **Post** /org/{org_handle}/workspace/{workspace_handle}/schema | Attach a schema to an org workspace
[**Detach**](OrgWorkspaceSchemas.md#Detach) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/schema/{schema_name} | Detach a schema from an org workspace
[**Get**](OrgWorkspaceSchemas.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema/{schema_name} | Get org workspace schema
[**Get_0**](OrgWorkspaceSchemas.md#Get_0) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema/{schema_name}/table | List org workspace schema tables
[**List**](OrgWorkspaceSchemas.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/schema | List org workspace schemas



## Attach

> WorkspaceSchema Attach(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Attach a schema to an org workspace



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the schema will be attached.
    request := *openapiclient.NewAttachWorkspaceSchemaRequest("ConnectionHandle_example") // AttachWorkspaceSchemaRequest | The request body for the schema to be attached.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSchemas.Attach(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSchemas.Attach``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Attach`: WorkspaceSchema
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSchemas.Attach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace belongs to. | 
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

> WorkspaceSchema Detach(ctx, orgHandle, workspaceHandle, schemaName).Execute()

Detach a schema from an org workspace



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace from which the schema will be detached.
    schemaName := "schemaName_example" // string | The name of the schema that will be detached.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSchemas.Detach(context.Background(), orgHandle, workspaceHandle, schemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSchemas.Detach``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Detach`: WorkspaceSchema
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSchemas.Detach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace belongs to. | 
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

> WorkspaceSchema Get(ctx, orgHandle, workspaceHandle, schemaName).Execute()

Get org workspace schema



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the schema exists.
    schemaName := "schemaName_example" // string | The name of the schema whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSchemas.Get(context.Background(), orgHandle, workspaceHandle, schemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSchemas.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceSchema
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSchemas.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace belongs to. | 
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

> ListWorkspaceSchemaTableResponse Get_0(ctx, orgHandle, workspaceHandle, schemaName).Execute()

List org workspace schema tables



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the schema exists.
    schemaName := "schemaName_example" // string | The name of the schema whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSchemas.Get_0(context.Background(), orgHandle, workspaceHandle, schemaName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSchemas.Get_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get_0`: ListWorkspaceSchemaTableResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSchemas.Get_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace belongs to. | 
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

> ListWorkspaceSchemaResponse List(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspace schemas



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
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSchemas.List(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSchemas.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListWorkspaceSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSchemas.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace belongs to. | 
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

