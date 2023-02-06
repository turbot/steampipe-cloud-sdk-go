# \UserWorkspaceProcesses

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](UserWorkspaceProcesses.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/process/{process_id} | Get user workspace process
[**List**](UserWorkspaceProcesses.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/process | List user workspace processes
[**Log**](UserWorkspaceProcesses.md#Log) | **Get** /user/{user_handle}/workspace/{workspace_handle}/process/{process_id}/log/{log_file}.{content_type} | List user workspace process logs



## Get

> SpProcess Get(ctx, userHandle, workspaceHandle, processId).Execute()

Get user workspace process



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose process needs to be fetched.
    processId := "processId_example" // string | The id of the process to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceProcesses.Get(context.Background(), userHandle, workspaceHandle, processId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceProcesses.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: SpProcess
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceProcesses.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the workspace exist. | 
**workspaceHandle** | **string** | The handle of the workspace whose process needs to be fetched. | 
**processId** | **string** | The id of the process to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SpProcess**](SpProcess.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListProcessesResponse List(ctx, userHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()

List user workspace processes



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
    userHandle := "userHandle_example" // string | The handle of the user for which you want to start listing processes.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where you want to list the processes.
    where := "where_example" // string | The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API. (optional)
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceProcesses.List(context.Background(), userHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceProcesses.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListProcessesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceProcesses.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which you want to start listing processes. | 
**workspaceHandle** | **string** | The handle of the workspace where you want to list the processes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **where** | **string** | The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API. | 
 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListProcessesResponse**](ListProcessesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Log

> string Log(ctx, userHandle, workspaceHandle, processId, logFile, contentType).Execute()

List user workspace process logs



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where you want to list the process logs.
    processId := "processId_example" // string | The id of the process where you want to list the process logs.
    logFile := "logFile_example" // string | The process logs file to be fetched.
    contentType := "contentType_example" // string | The required content type of the process logs.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceProcesses.Log(context.Background(), userHandle, workspaceHandle, processId, logFile, contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceProcesses.Log``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Log`: string
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceProcesses.Log`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the workspace exist. | 
**workspaceHandle** | **string** | The handle of the workspace where you want to list the process logs. | 
**processId** | **string** | The id of the process where you want to list the process logs. | 
**logFile** | **string** | The process logs file to be fetched. | 
**contentType** | **string** | The required content type of the process logs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/jsonlines+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

