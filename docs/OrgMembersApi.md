# \OrgMembersApi

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmOrgMemberInvite**](OrgMembersApi.md#ConfirmOrgMemberInvite) | **Get** /org/{org_handle}/member/invite/confirm | Confirm org member invite
[**DeleteOrgMember**](OrgMembersApi.md#DeleteOrgMember) | **Delete** /org/{org_handle}/member/{user_handle} | Delete org member
[**DeleteOrgMemberInvite**](OrgMembersApi.md#DeleteOrgMemberInvite) | **Delete** /org/{org_handle}/member/invite | Delete org member invite
[**GetOrgMember**](OrgMembersApi.md#GetOrgMember) | **Get** /org/{org_handle}/member/{user_handle} | Get org member
[**InviteOrgMember**](OrgMembersApi.md#InviteOrgMember) | **Post** /org/{org_handle}/member/invite | Invite org member
[**ListAcceptedOrgMembers**](OrgMembersApi.md#ListAcceptedOrgMembers) | **Get** /org/{org_handle}/member | List accepted org members
[**ListInvitedOrgMembers**](OrgMembersApi.md#ListInvitedOrgMembers) | **Get** /org/{org_handle}/member/invite | List invited org members
[**UpdateOrgMember**](OrgMembersApi.md#UpdateOrgMember) | **Patch** /org/{org_handle}/member/{user_handle} | Update org member



## ConfirmOrgMemberInvite

> ConfirmOrgMemberInvite(ctx, orgHandle).T(t).Execute()

Confirm org member invite



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
    orgHandle := "orgHandle_example" // string | Specify the handle of an organization where the member need to be invited.
    t := "t_example" // string | Specify the token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgMembersApi.ConfirmOrgMemberInvite(context.Background(), orgHandle).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgMembersApi.ConfirmOrgMemberInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of an organization where the member need to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmOrgMemberInviteRequest struct via the builder pattern


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


## DeleteOrgMember

> TypesOrgUser DeleteOrgMember(ctx, orgHandle, userHandle).Execute()

Delete org member



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
    orgHandle := "orgHandle_example" // string | Specify the handle of the organization where the member exists.
    userHandle := "userHandle_example" // string | Specify the handle of the user which need to be removed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgMembersApi.DeleteOrgMember(context.Background(), orgHandle, userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgMembersApi.DeleteOrgMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgMember`: TypesOrgUser
    fmt.Fprintf(os.Stdout, "Response from `OrgMembersApi.DeleteOrgMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member exists. | 
**userHandle** | **string** | Specify the handle of the user which need to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesOrgUser**](TypesOrgUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgMemberInvite

> DeleteOrgMemberInvite(ctx, orgHandle).T(t).Execute()

Delete org member invite



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
    orgHandle := "orgHandle_example" // string | Specify the organization handle.
    t := "t_example" // string | Specify the token to be rejected.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgMembersApi.DeleteOrgMemberInvite(context.Background(), orgHandle).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgMembersApi.DeleteOrgMemberInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the organization handle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMemberInviteRequest struct via the builder pattern


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


## GetOrgMember

> TypesOrgUser GetOrgMember(ctx, orgHandle, userHandle).Execute()

Get org member



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
    orgHandle := "orgHandle_example" // string | Specify the organization handle where the member is associated.
    userHandle := "userHandle_example" // string | Specify the handle of the user whose information you want to retrive.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgMembersApi.GetOrgMember(context.Background(), orgHandle, userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgMembersApi.GetOrgMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgMember`: TypesOrgUser
    fmt.Fprintf(os.Stdout, "Response from `OrgMembersApi.GetOrgMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the organization handle where the member is associated. | 
**userHandle** | **string** | Specify the handle of the user whose information you want to retrive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesOrgUser**](TypesOrgUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteOrgMember

> TypesOrgUser InviteOrgMember(ctx, orgHandle).Request(request).Execute()

Invite org member



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
    orgHandle := "orgHandle_example" // string | Specify the handle of an organization where the member need to be invited.
    request := *openapiclient.NewTypesInviteOrgUserRequest("Role_example") // TypesInviteOrgUserRequest | The request body to invite a member to an organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgMembersApi.InviteOrgMember(context.Background(), orgHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgMembersApi.InviteOrgMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteOrgMember`: TypesOrgUser
    fmt.Fprintf(os.Stdout, "Response from `OrgMembersApi.InviteOrgMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of an organization where the member need to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteOrgMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**TypesInviteOrgUserRequest**](TypesInviteOrgUserRequest.md) | The request body to invite a member to an organization. | 

### Return type

[**TypesOrgUser**](TypesOrgUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAcceptedOrgMembers

> TypesListOrgUsersResponse ListAcceptedOrgMembers(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

List accepted org members



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
    orgHandle := "orgHandle_example" // string | Specify the organization handle.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgMembersApi.ListAcceptedOrgMembers(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgMembersApi.ListAcceptedOrgMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAcceptedOrgMembers`: TypesListOrgUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgMembersApi.ListAcceptedOrgMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the organization handle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAcceptedOrgMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListOrgUsersResponse**](TypesListOrgUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvitedOrgMembers

> TypesListOrgUsersResponse ListInvitedOrgMembers(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

List invited org members



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
    orgHandle := "orgHandle_example" // string | Specify the organization handle.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgMembersApi.ListInvitedOrgMembers(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgMembersApi.ListInvitedOrgMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvitedOrgMembers`: TypesListOrgUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgMembersApi.ListInvitedOrgMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the organization handle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInvitedOrgMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListOrgUsersResponse**](TypesListOrgUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgMember

> TypesOrgUser UpdateOrgMember(ctx, orgHandle, userHandle).Request(request).Execute()

Update org member



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
    orgHandle := "orgHandle_example" // string | Specify the handle of the organization where the member exists.
    userHandle := "userHandle_example" // string | Specify the handle of the user whose role need to be updated.
    request := *openapiclient.NewTypesUpdateOrgUserRequest("Role_example") // TypesUpdateOrgUserRequest | The request body for the member.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgMembersApi.UpdateOrgMember(context.Background(), orgHandle, userHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgMembersApi.UpdateOrgMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgMember`: TypesOrgUser
    fmt.Fprintf(os.Stdout, "Response from `OrgMembersApi.UpdateOrgMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization where the member exists. | 
**userHandle** | **string** | Specify the handle of the user whose role need to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**TypesUpdateOrgUserRequest**](TypesUpdateOrgUserRequest.md) | The request body for the member. | 

### Return type

[**TypesOrgUser**](TypesOrgUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

