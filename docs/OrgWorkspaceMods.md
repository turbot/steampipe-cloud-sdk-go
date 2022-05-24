# \OrgWorkspaceMods

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgWorkspaceMod**](OrgWorkspaceMods.md#GetOrgWorkspaceMod) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Get organization workspace installed mod
[**InstallOrgWorkspaceMod**](OrgWorkspaceMods.md#InstallOrgWorkspaceMod) | **Post** /org/{org_handle}/workspace/{workspace_handle}/mod | Install a mod to an organization workspace
[**ListOrgWorkspaceMods**](OrgWorkspaceMods.md#ListOrgWorkspaceMods) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod | List organization workspace installed mods
[**UninstallOrgWorkspaceMod**](OrgWorkspaceMods.md#UninstallOrgWorkspaceMod) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Uninstall mod from organization workspace.
[**UpdateOrgWorkspaceMod**](OrgWorkspaceMods.md#UpdateOrgWorkspaceMod) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Update a mod in an organization workspace



## GetOrgWorkspaceMod

> WorkspaceMod GetOrgWorkspaceMod(ctx, orgHandle, workspaceHandle, modAlias).Execute()

Get organization workspace installed mod



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
    orgHandle := "orgHandle_example" // string | The handle of an organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where mod was installed.
    modAlias := "modAlias_example" // string | The mod alias or mod ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMods.GetOrgWorkspaceMod(context.Background(), orgHandle, workspaceHandle, modAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMods.GetOrgWorkspaceMod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgWorkspaceMod`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMods.GetOrgWorkspaceMod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where mod was installed. | 
**modAlias** | **string** | The mod alias or mod ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgWorkspaceModRequest struct via the builder pattern


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


## InstallOrgWorkspaceMod

> WorkspaceMod InstallOrgWorkspaceMod(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Install a mod to an organization workspace



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
    orgHandle := "orgHandle_example" // string | The handle of an organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod will be installed.
    request := *openapiclient.NewCreateWorkspaceModRequest("Path_example") // CreateWorkspaceModRequest | The request body to create a notification rule for this organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMods.InstallOrgWorkspaceMod(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMods.InstallOrgWorkspaceMod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallOrgWorkspaceMod`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMods.InstallOrgWorkspaceMod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod will be installed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallOrgWorkspaceModRequest struct via the builder pattern


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


## ListOrgWorkspaceMods

> ListWorkspaceModsResponse ListOrgWorkspaceMods(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List organization workspace installed mods



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
    orgHandle := "orgHandle_example" // string | The handle of an organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where mods were installed
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMods.ListOrgWorkspaceMods(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMods.ListOrgWorkspaceMods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgWorkspaceMods`: ListWorkspaceModsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMods.ListOrgWorkspaceMods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where mods were installed | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWorkspaceModsRequest struct via the builder pattern


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


## UninstallOrgWorkspaceMod

> WorkspaceMod UninstallOrgWorkspaceMod(ctx, orgHandle, workspaceHandle, modAlias).Execute()

Uninstall mod from organization workspace.



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
    orgHandle := "orgHandle_example" // string | The handle of an organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod was installed.
    modAlias := "modAlias_example" // string | The mod alias or mod ID to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMods.UninstallOrgWorkspaceMod(context.Background(), orgHandle, workspaceHandle, modAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMods.UninstallOrgWorkspaceMod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UninstallOrgWorkspaceMod`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMods.UninstallOrgWorkspaceMod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod was installed. | 
**modAlias** | **string** | The mod alias or mod ID to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallOrgWorkspaceModRequest struct via the builder pattern


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


## UpdateOrgWorkspaceMod

> WorkspaceMod UpdateOrgWorkspaceMod(ctx, orgHandle, workspaceHandle, modAlias).Request(request).Execute()

Update a mod in an organization workspace



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
    orgHandle := "orgHandle_example" // string | The handle of an organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod will be updated.
    modAlias := "modAlias_example" // string | The mod alias or mod ID to update.
    request := *openapiclient.NewUpdateWorkspaceModRequest("Constraint_example") // UpdateWorkspaceModRequest | The request body to update a mod for this workspace.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMods.UpdateOrgWorkspaceMod(context.Background(), orgHandle, workspaceHandle, modAlias).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMods.UpdateOrgWorkspaceMod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgWorkspaceMod`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMods.UpdateOrgWorkspaceMod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod will be updated. | 
**modAlias** | **string** | The mod alias or mod ID to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgWorkspaceModRequest struct via the builder pattern


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

