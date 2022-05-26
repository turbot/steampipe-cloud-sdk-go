# \UserWorkspaceModVariables

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSetting**](UserWorkspaceModVariables.md#CreateSetting) | **Post** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | Create a setting for a mod variable in a user workspace
[**DeleteSetting**](UserWorkspaceModVariables.md#DeleteSetting) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Delete setting for a mod variable in a user workspace
[**List**](UserWorkspaceModVariables.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | List variables for a user workspace mod
[**UpdateSetting**](UserWorkspaceModVariables.md#UpdateSetting) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Update setting for a mod variable in a user workspace



## CreateSetting

> WorkspaceModVariable CreateSetting(ctx, userHandle, workspaceHandle, modAlias).Request(request).Execute()

Create a setting for a mod variable in a user workspace



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod was installed
    modAlias := "modAlias_example" // string | The mod alias or mod ID for which the variable setting is to be created
    request := *openapiclient.NewCreateWorkspaceModVariableSettingRequest("Name_example", map[string]interface{}(123)) // CreateWorkspaceModVariableSettingRequest | The request body to create setting for mod variable in the user workspace.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceModVariables.CreateSetting(context.Background(), userHandle, workspaceHandle, modAlias).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceModVariables.CreateSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSetting`: WorkspaceModVariable
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceModVariables.CreateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod was installed | 
**modAlias** | **string** | The mod alias or mod ID for which the variable setting is to be created | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**CreateWorkspaceModVariableSettingRequest**](CreateWorkspaceModVariableSettingRequest.md) | The request body to create setting for mod variable in the user workspace. | 

### Return type

[**WorkspaceModVariable**](WorkspaceModVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSetting

> WorkspaceModVariable DeleteSetting(ctx, userHandle, workspaceHandle, modAlias, variableName).Execute()

Delete setting for a mod variable in a user workspace



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod was installed
    modAlias := "modAlias_example" // string | The mod alias or mod ID for which the variable setting is to be deleted
    variableName := "variableName_example" // string | The name of the variable for which setting is to be deleted

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceModVariables.DeleteSetting(context.Background(), userHandle, workspaceHandle, modAlias, variableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceModVariables.DeleteSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSetting`: WorkspaceModVariable
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceModVariables.DeleteSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod was installed | 
**modAlias** | **string** | The mod alias or mod ID for which the variable setting is to be deleted | 
**variableName** | **string** | The name of the variable for which setting is to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**WorkspaceModVariable**](WorkspaceModVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListWorkspaceModVariablesResponse List(ctx, userHandle, workspaceHandle, modAlias).Limit(limit).NextToken(nextToken).Execute()

List variables for a user workspace mod



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
    modAlias := "modAlias_example" // string | The mod alias or mod ID for which we want the variables to be listed
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceModVariables.List(context.Background(), userHandle, workspaceHandle, modAlias).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceModVariables.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListWorkspaceModVariablesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceModVariables.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where mods were installed | 
**modAlias** | **string** | The mod alias or mod ID for which we want the variables to be listed | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListWorkspaceModVariablesResponse**](ListWorkspaceModVariablesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSetting

> WorkspaceModVariable UpdateSetting(ctx, userHandle, workspaceHandle, modAlias, variableName).Request(request).Execute()

Update setting for a mod variable in a user workspace



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod was installed
    modAlias := "modAlias_example" // string | The mod alias or mod ID which contains the variable.
    variableName := "variableName_example" // string | The name of the variable for which setting is to be updated
    request := *openapiclient.NewUpdateWorkspaceModVariableSettingRequest(map[string]interface{}(123)) // UpdateWorkspaceModVariableSettingRequest | The request body to update setting for mod variable in the user workspace.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceModVariables.UpdateSetting(context.Background(), userHandle, workspaceHandle, modAlias, variableName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceModVariables.UpdateSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSetting`: WorkspaceModVariable
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceModVariables.UpdateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod was installed | 
**modAlias** | **string** | The mod alias or mod ID which contains the variable. | 
**variableName** | **string** | The name of the variable for which setting is to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | [**UpdateWorkspaceModVariableSettingRequest**](UpdateWorkspaceModVariableSettingRequest.md) | The request body to update setting for mod variable in the user workspace. | 

### Return type

[**WorkspaceModVariable**](WorkspaceModVariable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

