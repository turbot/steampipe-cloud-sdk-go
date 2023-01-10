# \UserWorkspacePipelines

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserWorkspacePipelines.md#Create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/pipeline | Create user workspace pipeline
[**Delete**](UserWorkspacePipelines.md#Delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Delete user workspace pipeline
[**Get**](UserWorkspacePipelines.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Get user workspace pipeline
[**List**](UserWorkspacePipelines.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/pipeline | List user workspace pipelines
[**Update**](UserWorkspacePipelines.md#Update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Update user workspace pipeline



## Create

> Pipeline Create(ctx, userHandle, workspaceHandle).Request(request).Execute()

Create user workspace pipeline



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
    userHandle := "userHandle_example" // string | The handle of the user which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where we want to create the pipeline.
    request := *openapiclient.NewCreatePipelineRequest(interface{}(123), interface{}(123), "Pipeline_example", "Title_example") // CreatePipelineRequest | The request body for the pipeline to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspacePipelines.Create(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspacePipelines.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspacePipelines.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where we want to create the pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreatePipelineRequest**](CreatePipelineRequest.md) | The request body for the pipeline to be created. | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Pipeline Delete(ctx, userHandle, workspaceHandle, pipelineId).Execute()

Delete user workspace pipeline



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
    userHandle := "userHandle_example" // string | The handle of the user which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the pipeline exists.
    pipelineId := "pipelineId_example" // string | Provide the id of the pipeline which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspacePipelines.Delete(context.Background(), userHandle, workspaceHandle, pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspacePipelines.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspacePipelines.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the pipeline exists. | 
**pipelineId** | **string** | Provide the id of the pipeline which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Pipeline Get(ctx, userHandle, workspaceHandle, pipelineId).Execute()

Get user workspace pipeline



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
    userHandle := "userHandle_example" // string | The handle of the user which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the pipeline exists.
    pipelineId := "pipelineId_example" // string | The id of the pipeline whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspacePipelines.Get(context.Background(), userHandle, workspaceHandle, pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspacePipelines.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspacePipelines.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the pipeline exists. | 
**pipelineId** | **string** | The id of the pipeline whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListPipelinesResponse List(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspace pipelines



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
    userHandle := "userHandle_example" // string | The handle of the user for which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which we want to list the pipelines.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspacePipelines.List(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspacePipelines.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListPipelinesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspacePipelines.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace for which we want to list the pipelines. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListPipelinesResponse**](ListPipelinesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Pipeline Update(ctx, userHandle, workspaceHandle, pipelineId).Request(request).Execute()

Update user workspace pipeline



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
    userHandle := "userHandle_example" // string | The handle of the user which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the pipeline exists.
    pipelineId := "pipelineId_example" // string | The id of the pipeline which needs to be updated.
    request := *openapiclient.NewUpdatePipelineRequest() // UpdatePipelineRequest | The request body for the pipeline which needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspacePipelines.Update(context.Background(), userHandle, workspaceHandle, pipelineId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspacePipelines.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspacePipelines.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the pipeline exists. | 
**pipelineId** | **string** | The id of the pipeline which needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdatePipelineRequest**](UpdatePipelineRequest.md) | The request body for the pipeline which needs to be updated. | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

