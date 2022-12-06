# \OrgWorkspaceSnapshots

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgWorkspaceSnapshots.md#Create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/snapshot | Create org workspace snapshot
[**Delete**](OrgWorkspaceSnapshots.md#Delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Delete org workspace snapshot
[**Download**](OrgWorkspaceSnapshots.md#Download) | **Get** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id}.{content_type} | Download org workspace snapshot
[**Get**](OrgWorkspaceSnapshots.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Get org workspace snapshot
[**List**](OrgWorkspaceSnapshots.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/snapshot | List org workspace snapshots
[**Update**](OrgWorkspaceSnapshots.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/snapshot/{snapshot_id} | Update org workspace snapshot



## Create

> WorkspaceSnapshot Create(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Create org workspace snapshot



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
    orgHandle := "orgHandle_example" // string | The handle of the org to create the workspace snapshot for.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace to create the snapshot for.
    request := *openapiclient.NewCreateWorkspaceSnapshotRequest(*openapiclient.NewWorkspaceSnapshotData("EndTime_example", map[string]interface{}(123), map[string]interface{}(123), "SchemaVersion_example", "StartTime_example")) // CreateWorkspaceSnapshotRequest | The request body for the org workspace snapshot to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSnapshots.Create(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSnapshots.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: WorkspaceSnapshot
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSnapshots.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org to create the workspace snapshot for. | 
**workspaceHandle** | **string** | The handle of the org workspace to create the snapshot for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateWorkspaceSnapshotRequest**](CreateWorkspaceSnapshotRequest.md) | The request body for the org workspace snapshot to be created. | 

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

> WorkspaceSnapshot Delete(ctx, orgHandle, workspaceHandle, snapshotId).Execute()

Delete org workspace snapshot



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
    orgHandle := "orgHandle_example" // string | The handle of the org that the workspace snapshot belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace that the snapshot belongs to.
    snapshotId := "snapshotId_example" // string | The handle of the snapshot which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSnapshots.Delete(context.Background(), orgHandle, workspaceHandle, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSnapshots.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: WorkspaceSnapshot
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSnapshots.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace snapshot belongs to. | 
**workspaceHandle** | **string** | The handle of the org workspace that the snapshot belongs to. | 
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

> WorkspaceSnapshotData Download(ctx, orgHandle, workspaceHandle, snapshotId, contentType).Execute()

Download org workspace snapshot



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
    orgHandle := "orgHandle_example" // string | The handle of the org that the workspace snapshot belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace that the snapshot belongs to.
    snapshotId := "snapshotId_example" // string | The Id of the snapshot to be downloaded.
    contentType := "contentType_example" // string | The type of content to the downloaded.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSnapshots.Download(context.Background(), orgHandle, workspaceHandle, snapshotId, contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSnapshots.Download``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Download`: WorkspaceSnapshotData
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSnapshots.Download`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace snapshot belongs to. | 
**workspaceHandle** | **string** | The handle of the org workspace that the snapshot belongs to. | 
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

> WorkspaceSnapshot Get(ctx, orgHandle, workspaceHandle, snapshotId).Execute()

Get org workspace snapshot



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
    orgHandle := "orgHandle_example" // string | The handle of the org that the workspace snapshot belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace that the snapshot belongs to.
    snapshotId := "snapshotId_example" // string | The handle of the snapshot whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSnapshots.Get(context.Background(), orgHandle, workspaceHandle, snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSnapshots.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceSnapshot
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSnapshots.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace snapshot belongs to. | 
**workspaceHandle** | **string** | The handle of the org workspace that the snapshot belongs to. | 
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

> ListWorkspaceSnapshotsResponse List(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspace snapshots



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
    orgHandle := "orgHandle_example" // string | The handle of the org to list the workspace snapshots for.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace to list snapshots for.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSnapshots.List(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSnapshots.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListWorkspaceSnapshotsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSnapshots.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org to list the workspace snapshots for. | 
**workspaceHandle** | **string** | The handle of the org workspace to list snapshots for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> WorkspaceSnapshot Update(ctx, orgHandle, workspaceHandle, snapshotId).Request(request).Execute()

Update org workspace snapshot



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
    orgHandle := "orgHandle_example" // string | The handle of the org that the workspace snapshot belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace that the snapshot belongs to.
    snapshotId := "snapshotId_example" // string | The handle of the snapshot to update.
    request := *openapiclient.NewUpdateWorkspaceSnapshotRequest() // UpdateWorkspaceSnapshotRequest | The request body for the org workspace snapshot to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceSnapshots.Update(context.Background(), orgHandle, workspaceHandle, snapshotId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceSnapshots.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: WorkspaceSnapshot
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceSnapshots.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org that the workspace snapshot belongs to. | 
**workspaceHandle** | **string** | The handle of the org workspace that the snapshot belongs to. | 
**snapshotId** | **string** | The handle of the snapshot to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateWorkspaceSnapshotRequest**](UpdateWorkspaceSnapshotRequest.md) | The request body for the org workspace snapshot to be updated. | 

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

