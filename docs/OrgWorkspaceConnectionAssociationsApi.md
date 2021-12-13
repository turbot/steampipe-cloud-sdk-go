# \OrgWorkspaceConnectionAssociationsApi

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWorkspaceConnectionAssociation**](OrgWorkspaceConnectionAssociationsApi.md#CreateOrgWorkspaceConnectionAssociation) | **Post** /org/{org_handle}/workspace/{workspace_handle}/conn | Create org workspace connection association
[**DeleteOrgWorkspaceConnectionAssociation**](OrgWorkspaceConnectionAssociationsApi.md#DeleteOrgWorkspaceConnectionAssociation) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Delete org workspace connection association
[**GetOrgWorkspaceConnectionAssociation**](OrgWorkspaceConnectionAssociationsApi.md#GetOrgWorkspaceConnectionAssociation) | **Get** /org/{org_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Get org workspace connection association
[**ListOrgWorkspaceConnectionAssociations**](OrgWorkspaceConnectionAssociationsApi.md#ListOrgWorkspaceConnectionAssociations) | **Get** /org/{org_handle}/workspace/{workspace_handle}/conn | List org workspace connection associations
[**TestOrgWorkspaceConnectionAssociation**](OrgWorkspaceConnectionAssociationsApi.md#TestOrgWorkspaceConnectionAssociation) | **Get** /org/{org_handle}/workspace/{workspace_handle}/conn/{conn_handle}/test | Test org workspace connection



## CreateOrgWorkspaceConnectionAssociation

> TypesWorkspaceConn CreateOrgWorkspaceConnectionAssociation(ctx, orgHandle, workspaceHandle).Request(request).Execute()

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
    request := *openapiclient.NewTypesCreateWorkspaceConnRequest("ConnectionHandle_example") // TypesCreateWorkspaceConnRequest | The request body for the association to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConnectionAssociationsApi.CreateOrgWorkspaceConnectionAssociation(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionAssociationsApi.CreateOrgWorkspaceConnectionAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgWorkspaceConnectionAssociation`: TypesWorkspaceConn
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionAssociationsApi.CreateOrgWorkspaceConnectionAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where we want to create an association. | 
**workspaceHandle** | **string** | The handle of the workspace where the connection will be associated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgWorkspaceConnectionAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**TypesCreateWorkspaceConnRequest**](TypesCreateWorkspaceConnRequest.md) | The request body for the association to be created. | 

### Return type

[**TypesWorkspaceConn**](TypesWorkspaceConn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgWorkspaceConnectionAssociation

> TypesWorkspaceConn DeleteOrgWorkspaceConnectionAssociation(ctx, orgHandle, workspaceHandle, connHandle).Execute()

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
    resp, r, err := api_client.OrgWorkspaceConnectionAssociationsApi.DeleteOrgWorkspaceConnectionAssociation(context.Background(), orgHandle, workspaceHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionAssociationsApi.DeleteOrgWorkspaceConnectionAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgWorkspaceConnectionAssociation`: TypesWorkspaceConn
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionAssociationsApi.DeleteOrgWorkspaceConnectionAssociation`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteOrgWorkspaceConnectionAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TypesWorkspaceConn**](TypesWorkspaceConn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgWorkspaceConnectionAssociation

> TypesWorkspaceConn GetOrgWorkspaceConnectionAssociation(ctx, orgHandle, workspaceHandle, connHandle).Execute()

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
    resp, r, err := api_client.OrgWorkspaceConnectionAssociationsApi.GetOrgWorkspaceConnectionAssociation(context.Background(), orgHandle, workspaceHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionAssociationsApi.GetOrgWorkspaceConnectionAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgWorkspaceConnectionAssociation`: TypesWorkspaceConn
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionAssociationsApi.GetOrgWorkspaceConnectionAssociation`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOrgWorkspaceConnectionAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TypesWorkspaceConn**](TypesWorkspaceConn.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgWorkspaceConnectionAssociations

> TypesListWorkspaceConnResponse ListOrgWorkspaceConnectionAssociations(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

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
    resp, r, err := api_client.OrgWorkspaceConnectionAssociationsApi.ListOrgWorkspaceConnectionAssociations(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionAssociationsApi.ListOrgWorkspaceConnectionAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgWorkspaceConnectionAssociations`: TypesListWorkspaceConnResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionAssociationsApi.ListOrgWorkspaceConnectionAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization for which you want to list the associations. | 
**workspaceHandle** | **string** | The handle of the workspace where you want to list the associations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWorkspaceConnectionAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListWorkspaceConnResponse**](TypesListWorkspaceConnResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestOrgWorkspaceConnectionAssociation

> TypesConnectionTestResponse TestOrgWorkspaceConnectionAssociation(ctx, orgHandle, workspaceHandle, connHandle).Execute()

Test org workspace connection



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
    resp, r, err := api_client.OrgWorkspaceConnectionAssociationsApi.TestOrgWorkspaceConnectionAssociation(context.Background(), orgHandle, workspaceHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionAssociationsApi.TestOrgWorkspaceConnectionAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestOrgWorkspaceConnectionAssociation`: TypesConnectionTestResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionAssociationsApi.TestOrgWorkspaceConnectionAssociation`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiTestOrgWorkspaceConnectionAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TypesConnectionTestResponse**](TypesConnectionTestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

