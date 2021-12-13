# \UserWorkspaceConnectionAssociationsApi

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserWorkspaceConnectionAssociation**](UserWorkspaceConnectionAssociationsApi.md#CreateUserWorkspaceConnectionAssociation) | **Post** /user/{user_handle}/workspace/{workspace_handle}/conn | Create user workspace connection association
[**DeleteUserWorkspaceConnectionAssociation**](UserWorkspaceConnectionAssociationsApi.md#DeleteUserWorkspaceConnectionAssociation) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Delete user workspace connection association
[**GetUserWorkspaceConnectionAssociation**](UserWorkspaceConnectionAssociationsApi.md#GetUserWorkspaceConnectionAssociation) | **Get** /user/{user_handle}/workspace/{workspace_handle}/conn/{conn_handle} | Get user workspace connection association
[**ListUserWorkspaceConnectionAssociations**](UserWorkspaceConnectionAssociationsApi.md#ListUserWorkspaceConnectionAssociations) | **Get** /user/{user_handle}/workspace/{workspace_handle}/conn | List user workspace connection associations
[**TestUserWorkspaceConnectionAssociation**](UserWorkspaceConnectionAssociationsApi.md#TestUserWorkspaceConnectionAssociation) | **Get** /user/{user_handle}/workspace/{workspace_handle}/conn/{conn_handle}/test | Test user workspace connection



## CreateUserWorkspaceConnectionAssociation

> TypesWorkspaceConn CreateUserWorkspaceConnectionAssociation(ctx, userHandle, workspaceHandle).Request(request).Execute()

Create user workspace connection association



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
    userHandle := "userHandle_example" // string | The handle of the user where we want to create an association.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the connection will be associated.
    request := *openapiclient.NewTypesCreateWorkspaceConnRequest("ConnectionHandle_example") // TypesCreateWorkspaceConnRequest | The request body for the association to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnectionAssociationsApi.CreateUserWorkspaceConnectionAssociation(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnectionAssociationsApi.CreateUserWorkspaceConnectionAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserWorkspaceConnectionAssociation`: TypesWorkspaceConn
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnectionAssociationsApi.CreateUserWorkspaceConnectionAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where we want to create an association. | 
**workspaceHandle** | **string** | The handle of the workspace where the connection will be associated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserWorkspaceConnectionAssociationRequest struct via the builder pattern


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


## DeleteUserWorkspaceConnectionAssociation

> TypesWorkspaceConn DeleteUserWorkspaceConnectionAssociation(ctx, userHandle, workspaceHandle, connHandle).Execute()

Delete user workspace connection association



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose association needs to be deleted.
    connHandle := "connHandle_example" // string | The handle of the conn whose association needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnectionAssociationsApi.DeleteUserWorkspaceConnectionAssociation(context.Background(), userHandle, workspaceHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnectionAssociationsApi.DeleteUserWorkspaceConnectionAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserWorkspaceConnectionAssociation`: TypesWorkspaceConn
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnectionAssociationsApi.DeleteUserWorkspaceConnectionAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user performing the action. | 
**workspaceHandle** | **string** | The handle of the workspace whose association needs to be deleted. | 
**connHandle** | **string** | The handle of the conn whose association needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserWorkspaceConnectionAssociationRequest struct via the builder pattern


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


## GetUserWorkspaceConnectionAssociation

> TypesWorkspaceConn GetUserWorkspaceConnectionAssociation(ctx, userHandle, workspaceHandle, connHandle).Execute()

Get user workspace connection association



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
    userHandle := "userHandle_example" // string | The handle of the user for which you want to get the association.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace wherethe association exist.
    connHandle := "connHandle_example" // string | The handle of the conn whose association details needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnectionAssociationsApi.GetUserWorkspaceConnectionAssociation(context.Background(), userHandle, workspaceHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnectionAssociationsApi.GetUserWorkspaceConnectionAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserWorkspaceConnectionAssociation`: TypesWorkspaceConn
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnectionAssociationsApi.GetUserWorkspaceConnectionAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which you want to get the association. | 
**workspaceHandle** | **string** | The handle of the workspace wherethe association exist. | 
**connHandle** | **string** | The handle of the conn whose association details needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserWorkspaceConnectionAssociationRequest struct via the builder pattern


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


## ListUserWorkspaceConnectionAssociations

> TypesListWorkspaceConnResponse ListUserWorkspaceConnectionAssociations(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspace connection associations



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
    userHandle := "userHandle_example" // string | The handle of the user for which you want to start listing associations.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where you want to list the associations.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnectionAssociationsApi.ListUserWorkspaceConnectionAssociations(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnectionAssociationsApi.ListUserWorkspaceConnectionAssociations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserWorkspaceConnectionAssociations`: TypesListWorkspaceConnResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnectionAssociationsApi.ListUserWorkspaceConnectionAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which you want to start listing associations. | 
**workspaceHandle** | **string** | The handle of the workspace where you want to list the associations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserWorkspaceConnectionAssociationsRequest struct via the builder pattern


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


## TestUserWorkspaceConnectionAssociation

> TypesConnectionTestResponse TestUserWorkspaceConnectionAssociation(ctx, userHandle, workspaceHandle, connHandle).Execute()

Test user workspace connection



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to test the connection in. The connection must already be associated with this workspace.
    connHandle := "connHandle_example" // string | The handle of the connection to test in this workspace. The connection must already be associated with this workspace whose association needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnectionAssociationsApi.TestUserWorkspaceConnectionAssociation(context.Background(), userHandle, workspaceHandle, connHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnectionAssociationsApi.TestUserWorkspaceConnectionAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestUserWorkspaceConnectionAssociation`: TypesConnectionTestResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnectionAssociationsApi.TestUserWorkspaceConnectionAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user performing the action. | 
**workspaceHandle** | **string** | The handle of the workspace to test the connection in. The connection must already be associated with this workspace. | 
**connHandle** | **string** | The handle of the connection to test in this workspace. The connection must already be associated with this workspace whose association needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestUserWorkspaceConnectionAssociationRequest struct via the builder pattern


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

