# \OrgConnectionsApi

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgConnection**](OrgConnectionsApi.md#CreateOrgConnection) | **Post** /org/{org_handle}/conn | Create org connection
[**DeleteOrgConnection**](OrgConnectionsApi.md#DeleteOrgConnection) | **Delete** /org/{org_handle}/conn/{conn_handle} | Delete org connection
[**GetOrgConnection**](OrgConnectionsApi.md#GetOrgConnection) | **Get** /org/{org_handle}/conn/{conn_handle} | Get org connection
[**ListOrgConnections**](OrgConnectionsApi.md#ListOrgConnections) | **Get** /org/{org_handle}/conn | List org connections
[**TestOrgConnection**](OrgConnectionsApi.md#TestOrgConnection) | **Post** /org/{org_handle}/conn/{conn_handle}/test | Test org connection
[**UpdateOrgConnection**](OrgConnectionsApi.md#UpdateOrgConnection) | **Patch** /org/{org_handle}/conn/{conn_handle} | Update org connection



## CreateOrgConnection

> TypesConnection CreateOrgConnection(ctx, orgHandle).Request(request).Execute()

Create org connection



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
    orgHandle := "orgHandle_example" // string | The handle of the organization where we want to create the connection.
    request := *openapiclient.NewTypesCreateConnectionRequest("Handle_example", "Plugin_example") // TypesCreateConnectionRequest | The request body for the connection to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionsApi.CreateOrgConnection(context.Background(), orgHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionsApi.CreateOrgConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgConnection`: TypesConnection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionsApi.CreateOrgConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where we want to create the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgConnectionRequest struct via the builder pattern


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


## DeleteOrgConnection

> TypesConnection DeleteOrgConnection(ctx, orgHandle, connHandle).Execute()

Delete org connection



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
    orgHandle := "orgHandle_example" // string | The handle of the organization where the connection exist.
    connHandle := "connHandle_example" // string | Provide the handle of the connection which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionsApi.DeleteOrgConnection(context.Background(), orgHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionsApi.DeleteOrgConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgConnection`: TypesConnection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionsApi.DeleteOrgConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the connection exist. | 
**connHandle** | **string** | Provide the handle of the connection which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgConnectionRequest struct via the builder pattern


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


## GetOrgConnection

> TypesConnection GetOrgConnection(ctx, orgHandle, connHandle).Execute()

Get org connection



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
    orgHandle := "orgHandle_example" // string | The handle of an organization where the connection exists.
    connHandle := "connHandle_example" // string | The handle of the connection whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionsApi.GetOrgConnection(context.Background(), orgHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionsApi.GetOrgConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgConnection`: TypesConnection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionsApi.GetOrgConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization where the connection exists. | 
**connHandle** | **string** | The handle of the connection whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgConnectionRequest struct via the builder pattern


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


## ListOrgConnections

> TypesListConnectionsResponse ListOrgConnections(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

List org connections



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
    orgHandle := "orgHandle_example" // string | The handle of the organization for which we want to list the connections.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionsApi.ListOrgConnections(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionsApi.ListOrgConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgConnections`: TypesListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionsApi.ListOrgConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization for which we want to list the connections. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgConnectionsRequest struct via the builder pattern


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


## TestOrgConnection

> TypesConnectionTestResult TestOrgConnection(ctx, orgHandle, connHandle).Execute()

Test org connection



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
    orgHandle := "orgHandle_example" // string | The handle of the org performing the action.
    connHandle := "connHandle_example" // string | The handle of the connection to be tested. For connections that are not yet created, use underscore `_` as the handle, else pass the handle of the existing connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionsApi.TestOrgConnection(context.Background(), orgHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionsApi.TestOrgConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestOrgConnection`: TypesConnectionTestResult
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionsApi.TestOrgConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org performing the action. | 
**connHandle** | **string** | The handle of the connection to be tested. For connections that are not yet created, use underscore &#x60;_&#x60; as the handle, else pass the handle of the existing connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestOrgConnectionRequest struct via the builder pattern


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


## UpdateOrgConnection

> TypesConnection UpdateOrgConnection(ctx, orgHandle, connHandle).Request(request).Execute()

Update org connection



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
    orgHandle := "orgHandle_example" // string | The handle of the organization where the connection exist.
    connHandle := "connHandle_example" // string | The handle of the connection which needs to be updated.
    request := *openapiclient.NewTypesUpdateConnectionRequest() // TypesUpdateConnectionRequest | The request body of the connection which needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionsApi.UpdateOrgConnection(context.Background(), orgHandle, connHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionsApi.UpdateOrgConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgConnection`: TypesConnection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionsApi.UpdateOrgConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the connection exist. | 
**connHandle** | **string** | The handle of the connection which needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**TypesUpdateConnectionRequest**](TypesUpdateConnectionRequest.md) | The request body of the connection which needs to be updated. | 

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

