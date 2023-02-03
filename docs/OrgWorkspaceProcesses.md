# \OrgWorkspaceProcesses

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](OrgWorkspaceProcesses.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/process/{process_id} | Get org workspace process
[**List**](OrgWorkspaceProcesses.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/process | List org workspace processes
[**Log**](OrgWorkspaceProcesses.md#Log) | **Get** /org/{org_handle}/workspace/{workspace_handle}/process/{process_id}/log/{log_file}.{content_type} | List org workspace process logs



## Get

> SpProcess Get(ctx, orgHandle, workspaceHandle, processId).Execute()

Get org workspace process



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
    orgHandle := "orgHandle_example" // string | The handle of the org where the workspace exist.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose process needs to be fetched.
    processId := "processId_example" // string | The id of the process to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceProcesses.Get(context.Background(), orgHandle, workspaceHandle, processId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceProcesses.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: SpProcess
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceProcesses.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org where the workspace exist. | 
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

> ListProcessesResponse List(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspace processes



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
    orgHandle := "orgHandle_example" // string | The handle of an organization for which you want to list the processes.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where you want to list the processes.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceProcesses.List(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceProcesses.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListProcessesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceProcesses.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization for which you want to list the processes. | 
**workspaceHandle** | **string** | The handle of the workspace where you want to list the processes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> string Log(ctx, orgHandle, workspaceHandle, processId, logFile, contentType).Execute()

List org workspace process logs



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which you want to list process logs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where you want to list the process logs.
    processId := "processId_example" // string | The id of the process where you want to list the process logs.
    logFile := "logFile_example" // string | The process logs file to be fetched.
    contentType := "contentType_example" // string | The required content type of the process logs.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceProcesses.Log(context.Background(), orgHandle, workspaceHandle, processId, logFile, contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceProcesses.Log``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Log`: string
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceProcesses.Log`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which you want to list process logs. | 
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

