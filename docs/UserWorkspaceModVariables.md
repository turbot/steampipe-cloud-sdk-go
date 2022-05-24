# \UserWorkspaceModVariables

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserWorkspaceModVariableSetting**](UserWorkspaceModVariables.md#CreateUserWorkspaceModVariableSetting) | **Post** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | Create a setting for a mod variable in a user workspace
[**DeleteUserWorkspaceModVariableSetting**](UserWorkspaceModVariables.md#DeleteUserWorkspaceModVariableSetting) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Delete setting for a mod variable in a user workspace
[**ListUserWorkspaceModVariables**](UserWorkspaceModVariables.md#ListUserWorkspaceModVariables) | **Get** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | List variables for a user workspace mod
[**UpdateUserWorkspaceModVariableSetting**](UserWorkspaceModVariables.md#UpdateUserWorkspaceModVariableSetting) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Update setting for a mod variable in a user workspace



## CreateUserWorkspaceModVariableSetting

> WorkspaceModVariable CreateUserWorkspaceModVariableSetting(ctx, userHandle, workspaceHandle, modAlias).Request(request).Execute()

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
    resp, r, err := api_client.UserWorkspaceModVariables.CreateUserWorkspaceModVariableSetting(context.Background(), userHandle, workspaceHandle, modAlias).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceModVariables.CreateUserWorkspaceModVariableSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserWorkspaceModVariableSetting`: WorkspaceModVariable
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceModVariables.CreateUserWorkspaceModVariableSetting`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateUserWorkspaceModVariableSettingRequest struct via the builder pattern


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


## DeleteUserWorkspaceModVariableSetting

> WorkspaceModVariable DeleteUserWorkspaceModVariableSetting(ctx, userHandle, workspaceHandle, modAlias, variableName).Execute()

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
    resp, r, err := api_client.UserWorkspaceModVariables.DeleteUserWorkspaceModVariableSetting(context.Background(), userHandle, workspaceHandle, modAlias, variableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceModVariables.DeleteUserWorkspaceModVariableSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserWorkspaceModVariableSetting`: WorkspaceModVariable
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceModVariables.DeleteUserWorkspaceModVariableSetting`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteUserWorkspaceModVariableSettingRequest struct via the builder pattern


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


## ListUserWorkspaceModVariables

> ListWorkspaceModVariablesResponse ListUserWorkspaceModVariables(ctx, userHandle, workspaceHandle, modAlias).Limit(limit).NextToken(nextToken).Execute()

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
    resp, r, err := api_client.UserWorkspaceModVariables.ListUserWorkspaceModVariables(context.Background(), userHandle, workspaceHandle, modAlias).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceModVariables.ListUserWorkspaceModVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserWorkspaceModVariables`: ListWorkspaceModVariablesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceModVariables.ListUserWorkspaceModVariables`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListUserWorkspaceModVariablesRequest struct via the builder pattern


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


## UpdateUserWorkspaceModVariableSetting

> WorkspaceModVariable UpdateUserWorkspaceModVariableSetting(ctx, userHandle, workspaceHandle, modAlias, variableName).Request(request).Execute()

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
    resp, r, err := api_client.UserWorkspaceModVariables.UpdateUserWorkspaceModVariableSetting(context.Background(), userHandle, workspaceHandle, modAlias, variableName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceModVariables.UpdateUserWorkspaceModVariableSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserWorkspaceModVariableSetting`: WorkspaceModVariable
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceModVariables.UpdateUserWorkspaceModVariableSetting`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateUserWorkspaceModVariableSettingRequest struct via the builder pattern


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

