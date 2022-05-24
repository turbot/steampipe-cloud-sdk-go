# \OrgWorkspaceModVariables

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWorkspaceModVariableSetting**](OrgWorkspaceModVariables.md#CreateOrgWorkspaceModVariableSetting) | **Post** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | Create a setting for a mod variable in an organization workspace
[**DeleteOrgWorkspaceModVariableSetting**](OrgWorkspaceModVariables.md#DeleteOrgWorkspaceModVariableSetting) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Delete setting for a mod variable in an organization workspace
[**ListOrgWorkspaceModVariables**](OrgWorkspaceModVariables.md#ListOrgWorkspaceModVariables) | **Get** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable | List variables in an organization workspace mod
[**UpdateOrgWorkspaceModVariableSetting**](OrgWorkspaceModVariables.md#UpdateOrgWorkspaceModVariableSetting) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/mod/{mod_alias}/variable/{variable_name} | Update setting for a mod variable in an organization workspace



## CreateOrgWorkspaceModVariableSetting

> WorkspaceModVariable CreateOrgWorkspaceModVariableSetting(ctx, orgHandle, workspaceHandle, modAlias).Request(request).Execute()

Create a setting for a mod variable in an organization workspace



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod was installed.
    modAlias := "modAlias_example" // string | The mod alias or mod ID for which the variable setting is to be created.
    request := *openapiclient.NewCreateWorkspaceModVariableSettingRequest("Name_example", map[string]interface{}(123)) // CreateWorkspaceModVariableSettingRequest | The request body to create setting for mod variable in the organization workspace.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceModVariables.CreateOrgWorkspaceModVariableSetting(context.Background(), orgHandle, workspaceHandle, modAlias).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceModVariables.CreateOrgWorkspaceModVariableSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgWorkspaceModVariableSetting`: WorkspaceModVariable
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceModVariables.CreateOrgWorkspaceModVariableSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod was installed. | 
**modAlias** | **string** | The mod alias or mod ID for which the variable setting is to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgWorkspaceModVariableSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**CreateWorkspaceModVariableSettingRequest**](CreateWorkspaceModVariableSettingRequest.md) | The request body to create setting for mod variable in the organization workspace. | 

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


## DeleteOrgWorkspaceModVariableSetting

> WorkspaceModVariable DeleteOrgWorkspaceModVariableSetting(ctx, orgHandle, workspaceHandle, modAlias, variableName).Execute()

Delete setting for a mod variable in an organization workspace



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod was installed.
    modAlias := "modAlias_example" // string | The mod alias or mod ID which contains the variable.
    variableName := "variableName_example" // string | The name of the variable for which setting is to be delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceModVariables.DeleteOrgWorkspaceModVariableSetting(context.Background(), orgHandle, workspaceHandle, modAlias, variableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceModVariables.DeleteOrgWorkspaceModVariableSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgWorkspaceModVariableSetting`: WorkspaceModVariable
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceModVariables.DeleteOrgWorkspaceModVariableSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod was installed. | 
**modAlias** | **string** | The mod alias or mod ID which contains the variable. | 
**variableName** | **string** | The name of the variable for which setting is to be delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgWorkspaceModVariableSettingRequest struct via the builder pattern


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


## ListOrgWorkspaceModVariables

> ListWorkspaceModVariablesResponse ListOrgWorkspaceModVariables(ctx, orgHandle, workspaceHandle, modAlias).Limit(limit).NextToken(nextToken).Execute()

List variables in an organization workspace mod



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod was installed.
    modAlias := "modAlias_example" // string | The mod alias or mod ID for which we want the variables to be listed.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceModVariables.ListOrgWorkspaceModVariables(context.Background(), orgHandle, workspaceHandle, modAlias).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceModVariables.ListOrgWorkspaceModVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgWorkspaceModVariables`: ListWorkspaceModVariablesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceModVariables.ListOrgWorkspaceModVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod was installed. | 
**modAlias** | **string** | The mod alias or mod ID for which we want the variables to be listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWorkspaceModVariablesRequest struct via the builder pattern


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


## UpdateOrgWorkspaceModVariableSetting

> WorkspaceModVariable UpdateOrgWorkspaceModVariableSetting(ctx, orgHandle, workspaceHandle, modAlias, variableName).Request(request).Execute()

Update setting for a mod variable in an organization workspace



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod was installed.
    modAlias := "modAlias_example" // string | The mod alias or mod ID which contains the variable.
    variableName := "variableName_example" // string | The name of the variable for which setting is to be updated.
    request := *openapiclient.NewUpdateWorkspaceModVariableSettingRequest(map[string]interface{}(123)) // UpdateWorkspaceModVariableSettingRequest | The request body to update setting for mod variable in the organization workspace.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceModVariables.UpdateOrgWorkspaceModVariableSetting(context.Background(), orgHandle, workspaceHandle, modAlias, variableName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceModVariables.UpdateOrgWorkspaceModVariableSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgWorkspaceModVariableSetting`: WorkspaceModVariable
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceModVariables.UpdateOrgWorkspaceModVariableSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod was installed. | 
**modAlias** | **string** | The mod alias or mod ID which contains the variable. | 
**variableName** | **string** | The name of the variable for which setting is to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgWorkspaceModVariableSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | [**UpdateWorkspaceModVariableSettingRequest**](UpdateWorkspaceModVariableSettingRequest.md) | The request body to update setting for mod variable in the organization workspace. | 

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

