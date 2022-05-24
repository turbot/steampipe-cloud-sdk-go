# \UserWorkspaceMods

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserWorkspaceMod**](UserWorkspaceMods.md#GetUserWorkspaceMod) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Get user workspace installed mod
[**InstallUserWorkspaceMods**](UserWorkspaceMods.md#InstallUserWorkspaceMods) | **Post** /user/{user_handle}/workspace/{workspace_handle}/mod | Install a mod to a user&#39;s workspace
[**ListUserWorkspaceMods**](UserWorkspaceMods.md#ListUserWorkspaceMods) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod | List user workspace installed mods
[**UninstallUserWorkspaceMod**](UserWorkspaceMods.md#UninstallUserWorkspaceMod) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Uninstall mod from a user&#39;s workspace.
[**UpdateUserWorkspaceMod**](UserWorkspaceMods.md#UpdateUserWorkspaceMod) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Update a mod in a user&#39;s workspace



## GetUserWorkspaceMod

> WorkspaceMod GetUserWorkspaceMod(ctx, userHandle, workspaceHandle, modAlias).Execute()

Get user workspace installed mod



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
    userHandle := "userHandle_example" // string | The handle of the user who owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | Provide the handle of the workspace where mod was installed.
    modAlias := "modAlias_example" // string | The mod alias or mod ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceMods.GetUserWorkspaceMod(context.Background(), userHandle, workspaceHandle, modAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceMods.GetUserWorkspaceMod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserWorkspaceMod`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceMods.GetUserWorkspaceMod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | Provide the handle of the workspace where mod was installed. | 
**modAlias** | **string** | The mod alias or mod ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserWorkspaceModRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceMod**](WorkspaceMod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallUserWorkspaceMods

> WorkspaceMod InstallUserWorkspaceMods(ctx, userHandle, workspaceHandle).Request(request).Execute()

Install a mod to a user's workspace



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
    userHandle := "userHandle_example" // string | The handle of the user who owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod will be installed.
    request := *openapiclient.NewCreateWorkspaceModRequest("Path_example") // CreateWorkspaceModRequest | The request body to create a notification rule for this organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceMods.InstallUserWorkspaceMods(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceMods.InstallUserWorkspaceMods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallUserWorkspaceMods`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceMods.InstallUserWorkspaceMods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod will be installed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallUserWorkspaceModsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateWorkspaceModRequest**](CreateWorkspaceModRequest.md) | The request body to create a notification rule for this organization. | 

### Return type

[**WorkspaceMod**](WorkspaceMod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserWorkspaceMods

> ListWorkspaceModsResponse ListUserWorkspaceMods(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspace installed mods



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
    userHandle := "userHandle_example" // string | The handle of the user who owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where mods were installed
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceMods.ListUserWorkspaceMods(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceMods.ListUserWorkspaceMods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserWorkspaceMods`: ListWorkspaceModsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceMods.ListUserWorkspaceMods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where mods were installed | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserWorkspaceModsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListWorkspaceModsResponse**](ListWorkspaceModsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallUserWorkspaceMod

> WorkspaceMod UninstallUserWorkspaceMod(ctx, userHandle, workspaceHandle, modAlias).Execute()

Uninstall mod from a user's workspace.



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
    userHandle := "userHandle_example" // string | The handle of the user who owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod was installed.
    modAlias := "modAlias_example" // string | The mod alias or mod ID to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceMods.UninstallUserWorkspaceMod(context.Background(), userHandle, workspaceHandle, modAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceMods.UninstallUserWorkspaceMod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UninstallUserWorkspaceMod`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceMods.UninstallUserWorkspaceMod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod was installed. | 
**modAlias** | **string** | The mod alias or mod ID to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallUserWorkspaceModRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceMod**](WorkspaceMod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserWorkspaceMod

> WorkspaceMod UpdateUserWorkspaceMod(ctx, userHandle, workspaceHandle, modAlias).Request(request).Execute()

Update a mod in a user's workspace



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
    userHandle := "userHandle_example" // string | The handle of the user who owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod will be updated.
    modAlias := "modAlias_example" // string | The mod alias or mod ID to update.
    request := *openapiclient.NewUpdateWorkspaceModRequest("Constraint_example") // UpdateWorkspaceModRequest | The request body to update a mod for this workspace.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceMods.UpdateUserWorkspaceMod(context.Background(), userHandle, workspaceHandle, modAlias).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceMods.UpdateUserWorkspaceMod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserWorkspaceMod`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceMods.UpdateUserWorkspaceMod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod will be updated. | 
**modAlias** | **string** | The mod alias or mod ID to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserWorkspaceModRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateWorkspaceModRequest**](UpdateWorkspaceModRequest.md) | The request body to update a mod for this workspace. | 

### Return type

[**WorkspaceMod**](WorkspaceMod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

