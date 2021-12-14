# \UserConnections

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserConnections.md#Create) | **Post** /user/{user_handle}/conn | Create user connection
[**Delete**](UserConnections.md#Delete) | **Delete** /user/{user_handle}/conn/{conn_handle} | Delete user connection
[**Get**](UserConnections.md#Get) | **Get** /user/{user_handle}/conn/{conn_handle} | Get user connection
[**List**](UserConnections.md#List) | **Get** /user/{user_handle}/conn | List user connections
[**Test**](UserConnections.md#Test) | **Post** /user/{user_handle}/conn/{conn_handle}/test | Test user connection
[**Update**](UserConnections.md#Update) | **Patch** /user/{user_handle}/conn/{conn_handle} | Update user connection



## Create

> Connection Create(ctx, userHandle).Request(request).Execute()

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
    request := *openapiclient.NewCreateConnectionRequest("Handle_example", "Plugin_example") // CreateConnectionRequest | The request body for the connection to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserConnections.Create(context.Background(), userHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnections.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Connection
    fmt.Fprintf(os.Stdout, "Response from `UserConnections.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where we want to create the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreateConnectionRequest**](CreateConnectionRequest.md) | The request body for the connection to be created. | 

### Return type

[**Connection**](Connection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Connection Delete(ctx, userHandle, connHandle).Execute()

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
    resp, r, err := api_client.UserConnections.Delete(context.Background(), userHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnections.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Connection
    fmt.Fprintf(os.Stdout, "Response from `UserConnections.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the connection exist. | 
**connHandle** | **string** | Provide the handle of the connection which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Connection**](Connection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Connection Get(ctx, userHandle, connHandle).Execute()

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
    resp, r, err := api_client.UserConnections.Get(context.Background(), userHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnections.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Connection
    fmt.Fprintf(os.Stdout, "Response from `UserConnections.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle for the user where the connection exist. | 
**connHandle** | **string** | The handle of the connection whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Connection**](Connection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListConnectionsResponse List(ctx, userHandle).Limit(limit).NextToken(nextToken).Execute()

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
    resp, r, err := api_client.UserConnections.List(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnections.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserConnections.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which we want to list the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**ListConnectionsResponse**](ListConnectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Test

> ConnectionTestResult Test(ctx, userHandle, connHandle).Execute()

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
    resp, r, err := api_client.UserConnections.Test(context.Background(), userHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnections.Test``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Test`: ConnectionTestResult
    fmt.Fprintf(os.Stdout, "Response from `UserConnections.Test`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user performing the action. | 
**connHandle** | **string** | The handle of the connection to be tested. For connections that are not yet created, use underscore &#x60;_&#x60; as the handle, else pass the handle of the existing connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectionTestResult**](ConnectionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Connection Update(ctx, userHandle, connHandle).Request(request).Execute()

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
    request := *openapiclient.NewUpdateConnectionRequest() // UpdateConnectionRequest | The request body for the connection which needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserConnections.Update(context.Background(), userHandle, connHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConnections.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Connection
    fmt.Fprintf(os.Stdout, "Response from `UserConnections.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the connection exist. | 
**connHandle** | **string** | The handle of the connection which needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateConnectionRequest**](UpdateConnectionRequest.md) | The request body for the connection which needs to be updated. | 

### Return type

[**Connection**](Connection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

