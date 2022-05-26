# \OrgWorkspaceMods

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](OrgWorkspaceMods.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Get organization workspace installed mod
[**Install**](OrgWorkspaceMods.md#Install) | **Post** /org/{org_handle}/workspace/{workspace_handle}/mod | Install a mod to an organization workspace
[**List**](OrgWorkspaceMods.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod | List organization workspace installed mods
[**Uninstall**](OrgWorkspaceMods.md#Uninstall) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Uninstall mod from organization workspace.
[**Update**](OrgWorkspaceMods.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias} | Update a mod in an organization workspace



## Get

> WorkspaceMod Get(ctx, orgHandle, workspaceHandle, modAlias).Execute()

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
    resp, r, err := api_client.OrgWorkspaceMods.Get(context.Background(), orgHandle, workspaceHandle, modAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMods.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMods.Get`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


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


## Install

> WorkspaceMod Install(ctx, orgHandle, workspaceHandle).Request(request).Execute()

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
    request := *openapiclient.NewCreateWorkspaceModRequest("Path_example") // CreateWorkspaceModRequest | The request body to install a mod in the mentioned workspace for this organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMods.Install(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMods.Install``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Install`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMods.Install`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod will be installed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateWorkspaceModRequest**](CreateWorkspaceModRequest.md) | The request body to install a mod in the mentioned workspace for this organization. | 

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


## List

> ListWorkspaceModsResponse List(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

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
    resp, r, err := api_client.OrgWorkspaceMods.List(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMods.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListWorkspaceModsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMods.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where mods were installed | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


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


## Uninstall

> WorkspaceMod Uninstall(ctx, orgHandle, workspaceHandle, modAlias).Execute()

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
    resp, r, err := api_client.OrgWorkspaceMods.Uninstall(context.Background(), orgHandle, workspaceHandle, modAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMods.Uninstall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Uninstall`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMods.Uninstall`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUninstallRequest struct via the builder pattern


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


## Update

> WorkspaceMod Update(ctx, orgHandle, workspaceHandle, modAlias).Request(request).Execute()

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
    resp, r, err := api_client.OrgWorkspaceMods.Update(context.Background(), orgHandle, workspaceHandle, modAlias).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMods.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMods.Update`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


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

