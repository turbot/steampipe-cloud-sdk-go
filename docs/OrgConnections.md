# \OrgConnections

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgConnections.md#Create) | **Post** /org/{org_handle}/connection | Create org connection
[**CreateDeprecated**](OrgConnections.md#CreateDeprecated) | **Post** /org/{org_handle}/conn | Create org connection
[**Delete**](OrgConnections.md#Delete) | **Delete** /org/{org_handle}/connection/{connection_handle} | Delete org connection
[**DeleteDeprecated**](OrgConnections.md#DeleteDeprecated) | **Delete** /org/{org_handle}/conn/{conn_handle} | Delete org connection
[**Get**](OrgConnections.md#Get) | **Get** /org/{org_handle}/connection/{connection_handle} | Get org connection
[**GetDeprecated**](OrgConnections.md#GetDeprecated) | **Get** /org/{org_handle}/conn/{conn_handle} | Get org connection
[**List**](OrgConnections.md#List) | **Get** /org/{org_handle}/connection | List org connections
[**ListDeprecated**](OrgConnections.md#ListDeprecated) | **Get** /org/{org_handle}/conn | List org connections
[**ListWorkspaces**](OrgConnections.md#ListWorkspaces) | **Get** /org/{org_handle}/connection/{connection_handle}/workspace | List org connection workspaces
[**ListWorkspacesDeprecated**](OrgConnections.md#ListWorkspacesDeprecated) | **Get** /org/{org_handle}/conn/{conn_handle}/workspace | List org connection workspaces
[**Test**](OrgConnections.md#Test) | **Post** /org/{org_handle}/connection/{connection_handle}/test | Test org connection
[**TestDeprecated**](OrgConnections.md#TestDeprecated) | **Post** /org/{org_handle}/conn/{conn_handle}/test | Test org connection
[**Update**](OrgConnections.md#Update) | **Patch** /org/{org_handle}/connection/{connection_handle} | Update org connection
[**UpdateDeprecated**](OrgConnections.md#UpdateDeprecated) | **Patch** /org/{org_handle}/conn/{conn_handle} | Update org connection



## Create

> Connection Create(ctx, orgHandle).Request(request).Mode(mode).Execute()

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
    request := *openapiclient.NewCreateConnectionRequest("Handle_example", "Plugin_example") // CreateConnectionRequest | The request body for the connection to be created.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.Create(context.Background(), orgHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where we want to create the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreateConnectionRequest**](CreateConnectionRequest.md) | The request body for the connection to be created. | 
 **mode** | **string** | The mode of this request | 

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


## CreateDeprecated

> Connection CreateDeprecated(ctx, orgHandle).Request(request).Mode(mode).Execute()

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
    request := *openapiclient.NewCreateConnectionRequest("Handle_example", "Plugin_example") // CreateConnectionRequest | The request body for the connection to be created.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.CreateDeprecated(context.Background(), orgHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.CreateDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeprecated`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.CreateDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where we want to create the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreateConnectionRequest**](CreateConnectionRequest.md) | The request body for the connection to be created. | 
 **mode** | **string** | The mode of this request | 

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

> Connection Delete(ctx, orgHandle, connectionHandle).Execute()

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
    orgHandle := "orgHandle_example" // string | The handle of the organization where the connection exists.
    connectionHandle := "connectionHandle_example" // string | The handle of the connection that needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.Delete(context.Background(), orgHandle, connectionHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the connection exists. | 
**connectionHandle** | **string** | The handle of the connection that needs to be deleted. | 

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


## DeleteDeprecated

> Connection DeleteDeprecated(ctx, orgHandle, connHandle).Execute()

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
    connHandle := "connHandle_example" // string | The handle of the connection that needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.DeleteDeprecated(context.Background(), orgHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.DeleteDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDeprecated`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.DeleteDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the connection exist. | 
**connHandle** | **string** | The handle of the connection that needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeprecatedRequest struct via the builder pattern


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

> Connection Get(ctx, orgHandle, connectionHandle).Execute()

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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.Get(context.Background(), orgHandle, connectionHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization where the connection exists. | 
**connectionHandle** | **string** | The handle of the connection whose detail needs to be fetched. | 

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


## GetDeprecated

> Connection GetDeprecated(ctx, orgHandle, connHandle).Execute()

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
    resp, r, err := api_client.OrgConnections.GetDeprecated(context.Background(), orgHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.GetDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeprecated`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.GetDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization where the connection exists. | 
**connHandle** | **string** | The handle of the connection whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeprecatedRequest struct via the builder pattern


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

> ListConnectionsResponse List(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

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
    orgHandle := "orgHandle_example" // string | The handle of the organization for which we want to list connections.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.List(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization for which we want to list connections. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

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


## ListDeprecated

> ListConnectionsResponse ListDeprecated(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to list connections.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.ListDeprecated(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.ListDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeprecated`: ListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.ListDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to list connections. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

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


## ListWorkspaces

> ListWorkspaceConnectionAssociationsResponse ListWorkspaces(ctx, orgHandle, connectionHandle).Limit(limit).NextToken(nextToken).Execute()

List org connection workspaces



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
    orgHandle := "orgHandle_example" // string | The handle of the organization where the connection exists.
    connectionHandle := "connectionHandle_example" // string | The handle of the connection for which we want to list workspaces.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.ListWorkspaces(context.Background(), orgHandle, connectionHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.ListWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspaces`: ListWorkspaceConnectionAssociationsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.ListWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the connection exists. | 
**connectionHandle** | **string** | The handle of the connection for which we want to list workspaces. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListWorkspaceConnectionAssociationsResponse**](ListWorkspaceConnectionAssociationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspacesDeprecated

> ListWorkspaceConnResponse ListWorkspacesDeprecated(ctx, orgHandle, connHandle).Limit(limit).NextToken(nextToken).Execute()

List org connection workspaces



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
    orgHandle := "orgHandle_example" // string | The handle of the organization where the connection exists.
    connHandle := "connHandle_example" // string | The handle of the connection for which we want to list workspaces.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.ListWorkspacesDeprecated(context.Background(), orgHandle, connHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.ListWorkspacesDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspacesDeprecated`: ListWorkspaceConnResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.ListWorkspacesDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the connection exists. | 
**connHandle** | **string** | The handle of the connection for which we want to list workspaces. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspacesDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListWorkspaceConnResponse**](ListWorkspaceConnResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Test

> ConnectionTestResult Test(ctx, orgHandle, connectionHandle).Request(request).Execute()

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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection to be tested. For connections that are not yet created, use underscore `_` as the handle, else pass the handle of the existing connection.
    request := *openapiclient.NewTestConnectionRequest("Plugin_example") // TestConnectionRequest | The request body for the connection to be tested.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.Test(context.Background(), orgHandle, connectionHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.Test``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Test`: ConnectionTestResult
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.Test`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org performing the action. | 
**connectionHandle** | **string** | The handle of the connection to be tested. For connections that are not yet created, use underscore &#x60;_&#x60; as the handle, else pass the handle of the existing connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**TestConnectionRequest**](TestConnectionRequest.md) | The request body for the connection to be tested. | 

### Return type

[**ConnectionTestResult**](ConnectionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestDeprecated

> ConnectionTestResult TestDeprecated(ctx, orgHandle, connHandle).Request(request).Execute()

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
    orgHandle := "orgHandle_example" // string | The handle of the org where the connection exists / intends to be created.
    connHandle := "connHandle_example" // string | The handle of the connection to be tested. For connections that are not yet created, use underscore `_` as the handle, else pass the handle of the existing connection.
    request := *openapiclient.NewTestConnectionRequest("Plugin_example") // TestConnectionRequest | The request body for the connection to be tested.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.TestDeprecated(context.Background(), orgHandle, connHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.TestDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestDeprecated`: ConnectionTestResult
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.TestDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org where the connection exists / intends to be created. | 
**connHandle** | **string** | The handle of the connection to be tested. For connections that are not yet created, use underscore &#x60;_&#x60; as the handle, else pass the handle of the existing connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**TestConnectionRequest**](TestConnectionRequest.md) | The request body for the connection to be tested. | 

### Return type

[**ConnectionTestResult**](ConnectionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Connection Update(ctx, orgHandle, connectionHandle).Request(request).Mode(mode).Execute()

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
    orgHandle := "orgHandle_example" // string | The handle of the organization where the connection exists.
    connectionHandle := "connectionHandle_example" // string | The handle of the connection which needs to be updated.
    request := *openapiclient.NewUpdateConnectionRequest() // UpdateConnectionRequest | The request body of the connection which needs to be updated.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.Update(context.Background(), orgHandle, connectionHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the connection exists. | 
**connectionHandle** | **string** | The handle of the connection which needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateConnectionRequest**](UpdateConnectionRequest.md) | The request body of the connection which needs to be updated. | 
 **mode** | **string** | The mode of this request | 

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


## UpdateDeprecated

> Connection UpdateDeprecated(ctx, orgHandle, connHandle).Request(request).Mode(mode).Execute()

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
    request := *openapiclient.NewUpdateConnectionRequest() // UpdateConnectionRequest | The request body for the connection which needs to be updated.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnections.UpdateDeprecated(context.Background(), orgHandle, connHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnections.UpdateDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeprecated`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnections.UpdateDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the connection exist. | 
**connHandle** | **string** | The handle of the connection which needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateConnectionRequest**](UpdateConnectionRequest.md) | The request body for the connection which needs to be updated. | 
 **mode** | **string** | The mode of this request | 

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

