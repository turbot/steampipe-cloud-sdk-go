# \UserWorkspaceAggregators

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserWorkspaceAggregators.md#Create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/aggregator | Create an aggregator for a user workspace
[**Delete**](UserWorkspaceAggregators.md#Delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Delete an aggregator for a user workspace
[**Get**](UserWorkspaceAggregators.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Get an aggregator for a user workspace
[**GetConnection**](UserWorkspaceAggregators.md#GetConnection) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle}/connection/{connection_handle} | Get a connection in the scope of an aggregator for a user workspace
[**List**](UserWorkspaceAggregators.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator | List aggregators for a user workspace
[**ListConnections**](UserWorkspaceAggregators.md#ListConnections) | **Get** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle}/connection | List connections in the scope of an aggregator for a user workspace
[**Update**](UserWorkspaceAggregators.md#Update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/aggregator/{aggregator_handle} | Update an aggregator for a user workspace



## Create

> WorkspaceAggregator Create(ctx, userHandle, workspaceHandle).Request(request).Mode(mode).Execute()

Create an aggregator for a user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the aggregator will be created.
    request := *openapiclient.NewCreateWorkspaceAggregatorRequest([]string{"Connections_example"}, "Handle_example", "Plugin_example") // CreateWorkspaceAggregatorRequest | The request body for the aggregator to be created.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceAggregators.Create(context.Background(), userHandle, workspaceHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceAggregators.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: WorkspaceAggregator
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceAggregators.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace where the aggregator will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateWorkspaceAggregatorRequest**](CreateWorkspaceAggregatorRequest.md) | The request body for the aggregator to be created. | 
 **mode** | **string** | The mode of this request | 

### Return type

[**WorkspaceAggregator**](WorkspaceAggregator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> WorkspaceAggregator Delete(ctx, userHandle, workspaceHandle, aggregatorHandle).Execute()

Delete an aggregator for a user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the aggregator exists.
    aggregatorHandle := "aggregatorHandle_example" // string | The handle of the aggregator that needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceAggregators.Delete(context.Background(), userHandle, workspaceHandle, aggregatorHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceAggregators.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: WorkspaceAggregator
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceAggregators.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace where the aggregator exists. | 
**aggregatorHandle** | **string** | The handle of the aggregator that needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceAggregator**](WorkspaceAggregator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> WorkspaceAggregator Get(ctx, userHandle, workspaceHandle, aggregatorHandle).Execute()

Get an aggregator for a user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the aggregator exists.
    aggregatorHandle := "aggregatorHandle_example" // string | The handle of the aggregator whose details need to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceAggregators.Get(context.Background(), userHandle, workspaceHandle, aggregatorHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceAggregators.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceAggregator
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceAggregators.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace where the aggregator exists. | 
**aggregatorHandle** | **string** | The handle of the aggregator whose details need to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceAggregator**](WorkspaceAggregator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnection

> WorkspaceConnection GetConnection(ctx, userHandle, workspaceHandle, aggregatorHandle, connectionHandle).Execute()

Get a connection in the scope of an aggregator for a user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the aggregator exists.
    aggregatorHandle := "aggregatorHandle_example" // string | The handle of the aggregator which the connection is a part of.
    connectionHandle := "connectionHandle_example" // string | The handle of the connection whose details need to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceAggregators.GetConnection(context.Background(), userHandle, workspaceHandle, aggregatorHandle, connectionHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceAggregators.GetConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnection`: WorkspaceConnection
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceAggregators.GetConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace where the aggregator exists. | 
**aggregatorHandle** | **string** | The handle of the aggregator which the connection is a part of. | 
**connectionHandle** | **string** | The handle of the connection whose details need to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**WorkspaceConnection**](WorkspaceConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListWorkspaceAggregatorsResponse List(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List aggregators for a user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which we will be listing aggregators.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceAggregators.List(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceAggregators.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListWorkspaceAggregatorsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceAggregators.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace for which we will be listing aggregators. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListWorkspaceAggregatorsResponse**](ListWorkspaceAggregatorsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnections

> ListWorkspaceConnectionsResponse ListConnections(ctx, userHandle, workspaceHandle, aggregatorHandle).Limit(limit).NextToken(nextToken).Execute()

List connections in the scope of an aggregator for a user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the aggregator exists.
    aggregatorHandle := "aggregatorHandle_example" // string | The handle of the aggregator whose connections need to be fetched.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceAggregators.ListConnections(context.Background(), userHandle, workspaceHandle, aggregatorHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceAggregators.ListConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnections`: ListWorkspaceConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceAggregators.ListConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace where the aggregator exists. | 
**aggregatorHandle** | **string** | The handle of the aggregator whose connections need to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListWorkspaceConnectionsResponse**](ListWorkspaceConnectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> WorkspaceAggregator Update(ctx, userHandle, workspaceHandle, aggregatorHandle).Request(request).Mode(mode).Execute()

Update an aggregator for a user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the aggregator exists.
    aggregatorHandle := "aggregatorHandle_example" // string | The handle of the aggregator whose details need to be updated.
    request := *openapiclient.NewUpdateWorkspaceAggregatorRequest() // UpdateWorkspaceAggregatorRequest | The request body for the aggregator which needs to be updated.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceAggregators.Update(context.Background(), userHandle, workspaceHandle, aggregatorHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceAggregators.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: WorkspaceAggregator
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceAggregators.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace where the aggregator exists. | 
**aggregatorHandle** | **string** | The handle of the aggregator whose details need to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateWorkspaceAggregatorRequest**](UpdateWorkspaceAggregatorRequest.md) | The request body for the aggregator which needs to be updated. | 
 **mode** | **string** | The mode of this request | 

### Return type

[**WorkspaceAggregator**](WorkspaceAggregator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

