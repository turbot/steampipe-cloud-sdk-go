# \OrgWorkspaceConnectionAssociations

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgWorkspaceConnectionAssociations.md#Create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/conn | Create org workspace connection association
[**Delete**](OrgWorkspaceConnectionAssociations.md#Delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Delete org workspace connection association
[**Get**](OrgWorkspaceConnectionAssociations.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Get org workspace connection association
[**List**](OrgWorkspaceConnectionAssociations.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/conn | List org workspace connection associations



## Create

> WorkspaceConn Create(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Create org workspace connection association



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
    orgHandle := "orgHandle_example" // string | The handle of the organization where we want to create an association.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the connection will be associated.
    request := *openapiclient.NewCreateWorkspaceConnRequest("ConnectionHandle_example") // CreateWorkspaceConnRequest | The request body for the association to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConnectionAssociations.Create(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionAssociations.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: WorkspaceConn
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionAssociations.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where we want to create an association. | 
**workspaceHandle** | **string** | The handle of the workspace where the connection will be associated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateWorkspaceConnRequest**](CreateWorkspaceConnRequest.md) | The request body for the association to be created. | 

### Return type

[**WorkspaceConn**](WorkspaceConn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> WorkspaceConn Delete(ctx, orgHandle, workspaceHandle, connHandle).Execute()

Delete org workspace connection association



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
    orgHandle := "orgHandle_example" // string | The handle of an organization where we want to delete the association.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose association needs to be deleted.
    connHandle := "connHandle_example" // string | The handle of the conn whose association needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConnectionAssociations.Delete(context.Background(), orgHandle, workspaceHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionAssociations.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: WorkspaceConn
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionAssociations.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization where we want to delete the association. | 
**workspaceHandle** | **string** | The handle of the workspace whose association needs to be deleted. | 
**connHandle** | **string** | The handle of the conn whose association needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceConn**](WorkspaceConn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> WorkspaceConn Get(ctx, orgHandle, workspaceHandle, connHandle).Execute()

Get org workspace connection association



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which you want to get the association.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the connection exist.
    connHandle := "connHandle_example" // string | The handle of the conn whose association details needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConnectionAssociations.Get(context.Background(), orgHandle, workspaceHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionAssociations.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceConn
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionAssociations.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which you want to get the association. | 
**workspaceHandle** | **string** | The handle of the workspace where the connection exist. | 
**connHandle** | **string** | The handle of the conn whose association details needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceConn**](WorkspaceConn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListWorkspaceConnResponse List(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspace connection associations



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
    orgHandle := "orgHandle_example" // string | The handle of an organization for which you want to list the associations.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where you want to list the associations.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConnectionAssociations.List(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionAssociations.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListWorkspaceConnResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionAssociations.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization for which you want to list the associations. | 
**workspaceHandle** | **string** | The handle of the workspace where you want to list the associations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

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

