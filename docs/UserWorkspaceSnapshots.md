# \UserWorkspaceSnapshots

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserWorkspaceSnapshots.md#Create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/snapshot | Create user workspace snapshot
[**Delete**](UserWorkspaceSnapshots.md#Delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Delete user workspace snapshot
[**Download**](UserWorkspaceSnapshots.md#Download) | **Get** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id}.{content_type} | Download user workspace snapshot
[**Get**](UserWorkspaceSnapshots.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Get user workspace snapshot
[**List**](UserWorkspaceSnapshots.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/snapshot | List user workspace snapshots
[**Update**](UserWorkspaceSnapshots.md#Update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Update user workspace snapshot



## Create

> WorkspaceSnapshot Create(ctx, userHandle, workspaceHandle).Request(request).Execute()

Create user workspace snapshot



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
    userHandle := "userHandle_example" // string | The handle of the user to create the workspace snapshot for.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the user workspace to create the snapshot for.
    request := *openapiclient.NewCreateWorkspaceSnapshotRequest(*openapiclient.NewWorkspaceSnapshotData("EndTime_example", *openapiclient.NewWorkspaceSnapshotDataLayout("Name_example", "PanelType_example"), map[string]interface{}(123), "SchemaVersion_example", "StartTime_example")) // CreateWorkspaceSnapshotRequest | The request body for the user workspace snapshot to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSnapshots.Create(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSnapshots.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: WorkspaceSnapshot
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSnapshots.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to create the workspace snapshot for. | 
**workspaceHandle** | **string** | The handle of the user workspace to create the snapshot for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateWorkspaceSnapshotRequest**](CreateWorkspaceSnapshotRequest.md) | The request body for the user workspace snapshot to be created. | 

### Return type

[**WorkspaceSnapshot**](WorkspaceSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> WorkspaceSnapshot Delete(ctx, userHandle, workspaceHandle, snapshotId).Execute()

Delete user workspace snapshot



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
    userHandle := "userHandle_example" // string | The handle of the user that the workspace snapshot belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the user workspace that the snapshot belongs to.
    snapshotId := "snapshotId_example" // string | The handle of the snapshot which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSnapshots.Delete(context.Background(), userHandle, workspaceHandle, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSnapshots.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: WorkspaceSnapshot
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSnapshots.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace snapshot belongs to. | 
**workspaceHandle** | **string** | The handle of the user workspace that the snapshot belongs to. | 
**snapshotId** | **string** | The handle of the snapshot which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceSnapshot**](WorkspaceSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Download

> WorkspaceSnapshotData Download(ctx, userHandle, workspaceHandle, snapshotId, contentType).Execute()

Download user workspace snapshot



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
    userHandle := "userHandle_example" // string | The handle of the user that the workspace snapshot belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the user workspace that the snapshot belongs to.
    snapshotId := "snapshotId_example" // string | The Id of the snapshot to be downloaded.
    contentType := "contentType_example" // string | The type of content to the downloaded.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSnapshots.Download(context.Background(), userHandle, workspaceHandle, snapshotId, contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSnapshots.Download``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Download`: WorkspaceSnapshotData
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSnapshots.Download`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace snapshot belongs to. | 
**workspaceHandle** | **string** | The handle of the user workspace that the snapshot belongs to. | 
**snapshotId** | **string** | The Id of the snapshot to be downloaded. | 
**contentType** | **string** | The type of content to the downloaded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**WorkspaceSnapshotData**](WorkspaceSnapshotData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> WorkspaceSnapshot Get(ctx, userHandle, workspaceHandle, snapshotId).Execute()

Get user workspace snapshot



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
    userHandle := "userHandle_example" // string | The handle of the user that the workspace snapshot belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the user workspace that the snapshot belongs to.
    snapshotId := "snapshotId_example" // string | The handle of the snapshot whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSnapshots.Get(context.Background(), userHandle, workspaceHandle, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSnapshots.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceSnapshot
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSnapshots.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace snapshot belongs to. | 
**workspaceHandle** | **string** | The handle of the user workspace that the snapshot belongs to. | 
**snapshotId** | **string** | The handle of the snapshot whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceSnapshot**](WorkspaceSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListWorkspaceSnapshotsResponse List(ctx, userHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()

List user workspace snapshots



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
    userHandle := "userHandle_example" // string | The handle of the user to list the workspace snapshots for.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to list snapshots for.
    where := "where_example" // string | The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API. (optional)
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSnapshots.List(context.Background(), userHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSnapshots.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListWorkspaceSnapshotsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSnapshots.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to list the workspace snapshots for. | 
**workspaceHandle** | **string** | The handle of the workspace to list snapshots for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **where** | **string** | The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API. | 
 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListWorkspaceSnapshotsResponse**](ListWorkspaceSnapshotsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> WorkspaceSnapshot Update(ctx, userHandle, workspaceHandle, snapshotId).Request(request).Execute()

Update user workspace snapshot



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
    userHandle := "userHandle_example" // string | The handle of the user that the workspace snapshot belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the user workspace that the snapshot belongs to.
    snapshotId := "snapshotId_example" // string | The handle of the snapshot to update.
    request := *openapiclient.NewUpdateWorkspaceSnapshotRequest() // UpdateWorkspaceSnapshotRequest | The request body for the user workspace snapshot to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceSnapshots.Update(context.Background(), userHandle, workspaceHandle, snapshotId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceSnapshots.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: WorkspaceSnapshot
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceSnapshots.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that the workspace snapshot belongs to. | 
**workspaceHandle** | **string** | The handle of the user workspace that the snapshot belongs to. | 
**snapshotId** | **string** | The handle of the snapshot to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateWorkspaceSnapshotRequest**](UpdateWorkspaceSnapshotRequest.md) | The request body for the user workspace snapshot to be updated. | 

### Return type

[**WorkspaceSnapshot**](WorkspaceSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

