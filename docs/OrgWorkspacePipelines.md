# \OrgWorkspacePipelines

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgWorkspacePipelines.md#Create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/pipeline | Create org workspace pipeline
[**Delete**](OrgWorkspacePipelines.md#Delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Delete org workspace pipeline
[**Get**](OrgWorkspacePipelines.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Get org workspace pipeline
[**List**](OrgWorkspacePipelines.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/pipeline | List org workspace pipelines
[**Update**](OrgWorkspacePipelines.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/pipeline/{pipeline_id} | Update org workspace pipeline



## Create

> Pipeline Create(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Create org workspace pipeline



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
    orgHandle := "orgHandle_example" // string | The handle of the organization which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where we want to create the pipeline.
    request := *openapiclient.NewCreatePipelineRequest(interface{}(123), interface{}(123), "Pipeline_example", "Title_example") // CreatePipelineRequest | The request body for the pipeline to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspacePipelines.Create(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacePipelines.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacePipelines.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization which contains the workspace. | 
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

> Pipeline Delete(ctx, orgHandle, workspaceHandle, pipelineId).Execute()

Delete org workspace pipeline



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
    orgHandle := "orgHandle_example" // string | The handle of the organization which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the pipeline exists.
    pipelineId := "pipelineId_example" // string | Provide the id of the pipeline which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspacePipelines.Delete(context.Background(), orgHandle, workspaceHandle, pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacePipelines.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacePipelines.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization which contains the workspace. | 
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

> Pipeline Get(ctx, orgHandle, workspaceHandle, pipelineId).Execute()

Get org workspace pipeline



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
    orgHandle := "orgHandle_example" // string | The handle of the organization which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the pipeline exists.
    pipelineId := "pipelineId_example" // string | The id of the pipeline whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspacePipelines.Get(context.Background(), orgHandle, workspaceHandle, pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacePipelines.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacePipelines.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization which contains the workspace. | 
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

> ListPipelinesResponse List(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspace pipelines



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
    orgHandle := "orgHandle_example" // string | The handle of the organization which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which we want to list the pipelines.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspacePipelines.List(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacePipelines.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListPipelinesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacePipelines.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization which contains the workspace. | 
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

> Pipeline Update(ctx, orgHandle, workspaceHandle, pipelineId).Request(request).Execute()

Update org workspace pipeline



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
    orgHandle := "orgHandle_example" // string | The handle of the organization which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the pipeline exists.
    pipelineId := "pipelineId_example" // string | The id of the pipeline which needs to be updated.
    request := *openapiclient.NewUpdatePipelineRequest() // UpdatePipelineRequest | The request body for the pipeline which needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspacePipelines.Update(context.Background(), orgHandle, workspaceHandle, pipelineId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspacePipelines.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspacePipelines.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization which contains the workspace. | 
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

