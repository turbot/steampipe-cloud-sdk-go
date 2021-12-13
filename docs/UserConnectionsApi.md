# \UserConnectionsApi

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserConnection**](UserConnectionsApi.md#CreateUserConnection) | **Post** /user/{user_handle}/conn | Create user connection
[**DeleteUserConnection**](UserConnectionsApi.md#DeleteUserConnection) | **Delete** /user/{user_handle}/conn/{conn_handle} | Delete user connection
[**GetUserConnection**](UserConnectionsApi.md#GetUserConnection) | **Get** /user/{user_handle}/conn/{conn_handle} | Get user connection
[**ListActorConnections**](UserConnectionsApi.md#ListActorConnections) | **Get** /actor/conn | List actor connections
[**ListUserConnections**](UserConnectionsApi.md#ListUserConnections) | **Get** /user/{user_handle}/conn | List user connections
[**TestUserConnection**](UserConnectionsApi.md#TestUserConnection) | **Post** /user/{user_handle}/conn/{conn_handle}/test | Test user connection
[**UpdateUserConnection**](UserConnectionsApi.md#UpdateUserConnection) | **Patch** /user/{user_handle}/conn/{conn_handle} | Update user connection



## CreateUserConnection

> TypesConnection CreateUserConnection(ctx, userHandle).Request(request).Execute()

Create user connection



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
    userHandle := "userHandle_example" // string | The handle of the user where we want to create the connection.
    request := *openapiclient.NewTypesCreateConnectionRequest("Handle_example", "Plugin_example") // TypesCreateConnectionRequest | The request body for the connection to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserConnectionsApi.CreateUserConnection(context.Background(), userHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnectionsApi.CreateUserConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserConnection`: TypesConnection
    fmt.Fprintf(os.Stdout, "Response from `UserConnectionsApi.CreateUserConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where we want to create the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**TypesCreateConnectionRequest**](TypesCreateConnectionRequest.md) | The request body for the connection to be created. | 

### Return type

[**TypesConnection**](TypesConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserConnection

> TypesConnection DeleteUserConnection(ctx, userHandle, connHandle).Execute()

Delete user connection



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
    userHandle := "userHandle_example" // string | The handle of the user where the connection exist.
    connHandle := "connHandle_example" // string | Provide the handle of the connection which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserConnectionsApi.DeleteUserConnection(context.Background(), userHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnectionsApi.DeleteUserConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserConnection`: TypesConnection
    fmt.Fprintf(os.Stdout, "Response from `UserConnectionsApi.DeleteUserConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the connection exist. | 
**connHandle** | **string** | Provide the handle of the connection which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesConnection**](TypesConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserConnection

> TypesConnection GetUserConnection(ctx, userHandle, connHandle).Execute()

Get user connection



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
    userHandle := "userHandle_example" // string | The handle for the user where the connection exist.
    connHandle := "connHandle_example" // string | The handle of the connection whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserConnectionsApi.GetUserConnection(context.Background(), userHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnectionsApi.GetUserConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserConnection`: TypesConnection
    fmt.Fprintf(os.Stdout, "Response from `UserConnectionsApi.GetUserConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle for the user where the connection exist. | 
**connHandle** | **string** | The handle of the connection whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesConnection**](TypesConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActorConnections

> TypesListConnectionsResponse ListActorConnections(ctx).Limit(limit).NextToken(nextToken).Execute()

List actor connections



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
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserConnectionsApi.ListActorConnections(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnectionsApi.ListActorConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActorConnections`: TypesListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserConnectionsApi.ListActorConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListActorConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListConnectionsResponse**](TypesListConnectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserConnections

> TypesListConnectionsResponse ListUserConnections(ctx, userHandle).Limit(limit).NextToken(nextToken).Execute()

List user connections



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
    userHandle := "userHandle_example" // string | The handle of the user for which we want to list the connection.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserConnectionsApi.ListUserConnections(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnectionsApi.ListUserConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserConnections`: TypesListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserConnectionsApi.ListUserConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which we want to list the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListConnectionsResponse**](TypesListConnectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUserConnection

> TypesConnectionTestResult TestUserConnection(ctx, userHandle, connHandle).Execute()

Test user connection



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
    userHandle := "userHandle_example" // string | The handle of the user performing the action.
    connHandle := "connHandle_example" // string | The handle of the connection to be tested. For connections that are not yet created, use underscore `_` as the handle, else pass the handle of the existing connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserConnectionsApi.TestUserConnection(context.Background(), userHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnectionsApi.TestUserConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestUserConnection`: TypesConnectionTestResult
    fmt.Fprintf(os.Stdout, "Response from `UserConnectionsApi.TestUserConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user performing the action. | 
**connHandle** | **string** | The handle of the connection to be tested. For connections that are not yet created, use underscore &#x60;_&#x60; as the handle, else pass the handle of the existing connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestUserConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesConnectionTestResult**](TypesConnectionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserConnection

> TypesConnection UpdateUserConnection(ctx, userHandle, connHandle).Request(request).Execute()

Update user connection



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
    userHandle := "userHandle_example" // string | The handle of the user where the connection exist.
    connHandle := "connHandle_example" // string | The handle of the connection which needs to be updated.
    request := *openapiclient.NewTypesUpdateConnectionRequest() // TypesUpdateConnectionRequest | The request body for the connection which needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserConnectionsApi.UpdateUserConnection(context.Background(), userHandle, connHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnectionsApi.UpdateUserConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserConnection`: TypesConnection
    fmt.Fprintf(os.Stdout, "Response from `UserConnectionsApi.UpdateUserConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the connection exist. | 
**connHandle** | **string** | The handle of the connection which needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**TypesUpdateConnectionRequest**](TypesUpdateConnectionRequest.md) | The request body for the connection which needs to be updated. | 

### Return type

[**TypesConnection**](TypesConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

