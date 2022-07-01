# \OrgWorkspaceMembers

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmInvite**](OrgWorkspaceMembers.md#ConfirmInvite) | **Get** /org/{org_handle}/workspace/{workspace_handle}/member/invite/confirm | Confirm Org Workspace Member Invite
[**Delete**](OrgWorkspaceMembers.md#Delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Delete Org Workspace Member
[**Get**](OrgWorkspaceMembers.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Get Org Workspace Member
[**Invite**](OrgWorkspaceMembers.md#Invite) | **Post** /org/{org_handle}/workspace/{workspace_handle}/member/invite | Invite Org Workspace Member
[**ListAccepted**](OrgWorkspaceMembers.md#ListAccepted) | **Get** /org/{org_handle}/workspace/{workspace_handle}/member | List accepted Org Workspace Members
[**ListInvited**](OrgWorkspaceMembers.md#ListInvited) | **Get** /org/{org_handle}/workspace/{workspace_handle}/member/invite | List invited Org Workspace Members
[**RejectInvite**](OrgWorkspaceMembers.md#RejectInvite) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/member/invite | Reject Org Workspace Member Invite
[**Update**](OrgWorkspaceMembers.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/member/{user_handle} | Update Org Workspace Member



## ConfirmInvite

> ConfirmInvite(ctx, orgHandle, workspaceHandle).T(t).Execute()

Confirm Org Workspace Member Invite



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
    t := "t_example" // string | Specify the token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMembers.ConfirmInvite(context.Background(), orgHandle, workspaceHandle).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMembers.ConfirmInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member need to be invited. | 
**workspaceHandle** | **string** | Specify the handle of the workspace where the member need to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t** | **string** | Specify the token. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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


## Invite

> OrgWorkspaceUser Invite(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Invite Org Workspace Member



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
    request := *openapiclient.NewInviteOrgWorkspaceUserRequest("Role_example") // InviteOrgWorkspaceUserRequest | The request body to invite a member to an organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMembers.Invite(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMembers.Invite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Invite`: OrgWorkspaceUser
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMembers.Invite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member need to be invited. | 
**workspaceHandle** | **string** | Specify the handle of the workspace where the member need to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**InviteOrgWorkspaceUserRequest**](InviteOrgWorkspaceUserRequest.md) | The request body to invite a member to an organization. | 

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


## ListAccepted

> ListOrgWorkspaceUsersResponse ListAccepted(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List accepted Org Workspace Members



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
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMembers.ListAccepted(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMembers.ListAccepted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccepted`: ListOrgWorkspaceUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMembers.ListAccepted`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member need to be invited. | 
**workspaceHandle** | **string** | Specify the handle of the workspace where the member need to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAcceptedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListInvited

> ListOrgWorkspaceUsersResponse ListInvited(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List invited Org Workspace Members



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
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMembers.ListInvited(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMembers.ListInvited``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvited`: ListOrgWorkspaceUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceMembers.ListInvited`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member need to be invited. | 
**workspaceHandle** | **string** | Specify the handle of the workspace where the member need to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInvitedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RejectInvite

> RejectInvite(ctx, orgHandle, workspaceHandle).T(t).Execute()

Reject Org Workspace Member Invite



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
    t := "t_example" // string | Specify the token to be rejected.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceMembers.RejectInvite(context.Background(), orgHandle, workspaceHandle).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceMembers.RejectInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member need to be invited. | 
**workspaceHandle** | **string** | Specify the handle of the workspace where the member need to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t** | **string** | Specify the token to be rejected. | 

### Return type

 (empty response body)

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

