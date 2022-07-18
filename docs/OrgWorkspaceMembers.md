# \OrgWorkspaceMembers

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgWorkspaceMembers.md#Create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/member | Create Org Workspace Member
[**Delete**](OrgWorkspaceMembers.md#Delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Delete Org Workspace Member
[**Get**](OrgWorkspaceMembers.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Get Org Workspace Member
[**List**](OrgWorkspaceMembers.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/member | List Organization Workspace Members
[**Update**](OrgWorkspaceMembers.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Update Org Workspace Member



## Create

> OrgWorkspaceUser Create(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Create Org Workspace Member



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
    orgHandle := "orgHandle_example" // string | Specify the handle of the organization where the member need to be invited.
    workspaceHandle := "workspaceHandle_example" // string | Specify the handle of the workspace where the member need to be invited.
    request := *openapiclient.NewCreateOrgWorkspaceUserRequest("Handle_example", "Role_example") // CreateOrgWorkspaceUserRequest | The request body to invite a member to an organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMembers.Create(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMembers.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: OrgWorkspaceUser
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMembers.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member need to be invited. | 
**workspaceHandle** | **string** | Specify the handle of the workspace where the member need to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateOrgWorkspaceUserRequest**](CreateOrgWorkspaceUserRequest.md) | The request body to invite a member to an organization. | 

### Return type

[**OrgWorkspaceUser**](OrgWorkspaceUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> OrgWorkspaceUser Delete(ctx, orgHandle, workspaceHandle, userHandle).Execute()

Delete Org Workspace Member



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
    orgHandle := "orgHandle_example" // string | Specify the handle of the organization where the member need to be invited.
    workspaceHandle := "workspaceHandle_example" // string | Specify the handle of the workspace where the member need to be invited.
    userHandle := "userHandle_example" // string | Specify the handle of the user whose information you want to retrive.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMembers.Delete(context.Background(), orgHandle, workspaceHandle, userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMembers.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: OrgWorkspaceUser
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMembers.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member need to be invited. | 
**workspaceHandle** | **string** | Specify the handle of the workspace where the member need to be invited. | 
**userHandle** | **string** | Specify the handle of the user whose information you want to retrive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrgWorkspaceUser**](OrgWorkspaceUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> OrgWorkspaceUser Get(ctx, orgHandle, workspaceHandle, userHandle).Execute()

Get Org Workspace Member



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
    orgHandle := "orgHandle_example" // string | Specify the handle of the organization where the member need to be invited.
    workspaceHandle := "workspaceHandle_example" // string | Specify the handle of the workspace where the member need to be invited.
    userHandle := "userHandle_example" // string | Specify the handle of the user whose information you want to retrive.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMembers.Get(context.Background(), orgHandle, workspaceHandle, userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMembers.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: OrgWorkspaceUser
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMembers.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member need to be invited. | 
**workspaceHandle** | **string** | Specify the handle of the workspace where the member need to be invited. | 
**userHandle** | **string** | Specify the handle of the user whose information you want to retrive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrgWorkspaceUser**](OrgWorkspaceUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListOrgWorkspaceUsersResponse List(ctx, orgHandle, workspaceHandle).Q(q).Limit(limit).NextToken(nextToken).Execute()

List Organization Workspace Members



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
    orgHandle := "orgHandle_example" // string | Specify the handle of the organization where the member need to be invited.
    workspaceHandle := "workspaceHandle_example" // string | Specify the handle of the workspace where the member need to be invited.
    q := "q_example" // string | Optional free-text search to filter the org workspace members. (optional)
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMembers.List(context.Background(), orgHandle, workspaceHandle).Q(q).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMembers.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListOrgWorkspaceUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMembers.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member need to be invited. | 
**workspaceHandle** | **string** | Specify the handle of the workspace where the member need to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** | Optional free-text search to filter the org workspace members. | 
 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListOrgWorkspaceUsersResponse**](ListOrgWorkspaceUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> OrgWorkspaceUser Update(ctx, orgHandle, workspaceHandle, userHandle).Request(request).Execute()

Update Org Workspace Member



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
    orgHandle := "orgHandle_example" // string | Specify the handle of the organization where the member need to be invited.
    workspaceHandle := "workspaceHandle_example" // string | Specify the handle of the workspace where the member need to be invited.
    userHandle := "userHandle_example" // string | Specify the handle of the user whose information you want to retrive.
    request := *openapiclient.NewUpdateOrgWorkspaceUserRequest("Role_example") // UpdateOrgWorkspaceUserRequest | The request body for the member.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMembers.Update(context.Background(), orgHandle, workspaceHandle, userHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMembers.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: OrgWorkspaceUser
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMembers.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member need to be invited. | 
**workspaceHandle** | **string** | Specify the handle of the workspace where the member need to be invited. | 
**userHandle** | **string** | Specify the handle of the user whose information you want to retrive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateOrgWorkspaceUserRequest**](UpdateOrgWorkspaceUserRequest.md) | The request body for the member. | 

### Return type

[**OrgWorkspaceUser**](OrgWorkspaceUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

