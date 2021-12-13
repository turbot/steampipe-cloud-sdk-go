# \Users

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Users.md#Create) | **Post** /user | Create user
[**CreateDBPassword**](Users.md#CreateDBPassword) | **Post** /user/{user_handle}/password | Create user password
[**Delete**](Users.md#Delete) | **Delete** /user/{user_handle} | Delete user
[**Get**](Users.md#Get) | **Get** /user/{user_handle} | Get user
[**GetDBPassword**](Users.md#GetDBPassword) | **Get** /user/{user_handle}/password | Get user password
[**GetQuota**](Users.md#GetQuota) | **Get** /user/{user_handle}/quota | User quota
[**List**](Users.md#List) | **Get** /user | List users
[**ListAuditLogs**](Users.md#ListAuditLogs) | **Get** /user/{user_handle}/audit | User audit logs
[**ListOrgInvites**](Users.md#ListOrgInvites) | **Get** /user/{user_handle}/org/invite | List org invited users
[**ListOrgs**](Users.md#ListOrgs) | **Get** /user/{user_handle}/org | List org users
[**ListWorkspaceAuditLogs**](Users.md#ListWorkspaceAuditLogs) | **Get** /user/{user_handle}/workspace/{workspace_handle}/audit | User workspace audit logs
[**Search**](Users.md#Search) | **Get** /user/search | Search users
[**Update**](Users.md#Update) | **Patch** /user/{user_handle} | Update user



## Create

> TypesUser Create(ctx).Request(request).Execute()

Create user



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
    request := *openapiclient.NewTypesCreateUserRequest("Email_example", "Handle_example") // TypesCreateUserRequest | The request body to create the user.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.Create(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: TypesUser
    fmt.Fprintf(os.Stdout, "Response from `Users.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**TypesCreateUserRequest**](TypesCreateUserRequest.md) | The request body to create the user. | 

### Return type

[**TypesUser**](TypesUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDBPassword

> TypesUserDatabasePassword CreateDBPassword(ctx, userHandle).Request(request).Execute()

Create user password



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
    userHandle := "userHandle_example" // string | Specify the handle of the user whose password need to be created or rotated.
    request := *openapiclient.NewTypesCreateUserPasswordRequest() // TypesCreateUserPasswordRequest | The request body to create or rotate the password.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.CreateDBPassword(context.Background(), userHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.CreateDBPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDBPassword`: TypesUserDatabasePassword
    fmt.Fprintf(os.Stdout, "Response from `Users.CreateDBPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user whose password need to be created or rotated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDBPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**TypesCreateUserPasswordRequest**](TypesCreateUserPasswordRequest.md) | The request body to create or rotate the password. | 

### Return type

[**TypesUserDatabasePassword**](TypesUserDatabasePassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> TypesUser Delete(ctx, userHandle).Execute()

Delete user



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
    userHandle := "userHandle_example" // string | Specify the handle of the user which need to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.Delete(context.Background(), userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: TypesUser
    fmt.Fprintf(os.Stdout, "Response from `Users.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user which need to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesUser**](TypesUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> TypesUser Get(ctx, userHandle).Execute()

Get user



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
    userHandle := "userHandle_example" // string | Specify the handle of the user whose information you want to retrive.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.Get(context.Background(), userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: TypesUser
    fmt.Fprintf(os.Stdout, "Response from `Users.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user whose information you want to retrive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesUser**](TypesUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDBPassword

> TypesUserDatabasePassword GetDBPassword(ctx, userHandle).Execute()

Get user password



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
    userHandle := "userHandle_example" // string | Specify the handle of the user whose password need to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.GetDBPassword(context.Background(), userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.GetDBPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDBPassword`: TypesUserDatabasePassword
    fmt.Fprintf(os.Stdout, "Response from `Users.GetDBPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user whose password need to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDBPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesUserDatabasePassword**](TypesUserDatabasePassword.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuota

> TypesUserQuota GetQuota(ctx, userHandle).Execute()

User quota



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
    userHandle := "userHandle_example" // string | Specify the user handle to get the quota details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.GetQuota(context.Background(), userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.GetQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuota`: TypesUserQuota
    fmt.Fprintf(os.Stdout, "Response from `Users.GetQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the user handle to get the quota details. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesUserQuota**](TypesUserQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> TypesListUsersResponse List(ctx).Limit(limit).NextToken(nextToken).Execute()

List users



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
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.List(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: TypesListUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `Users.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListUsersResponse**](TypesListUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuditLogs

> TypesListAuditLogsResponse ListAuditLogs(ctx, userHandle).Limit(limit).NextToken(nextToken).Execute()

User audit logs



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
    userHandle := "userHandle_example" // string | Specify the user handle to get the audit logs.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.ListAuditLogs(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.ListAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuditLogs`: TypesListAuditLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `Users.ListAuditLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the user handle to get the audit logs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListAuditLogsResponse**](TypesListAuditLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgInvites

> TypesListUserOrgsResponse ListOrgInvites(ctx, userHandle).Limit(limit).NextToken(nextToken).Execute()

List org invited users



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
    userHandle := "userHandle_example" // string | Pass the handle of the user.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.ListOrgInvites(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.ListOrgInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgInvites`: TypesListUserOrgsResponse
    fmt.Fprintf(os.Stdout, "Response from `Users.ListOrgInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Pass the handle of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListUserOrgsResponse**](TypesListUserOrgsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgs

> TypesListUserOrgsResponse ListOrgs(ctx, userHandle).Limit(limit).NextToken(nextToken).Execute()

List org users



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
    userHandle := "userHandle_example" // string | The handle of the user for which we want to list the org.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.ListOrgs(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.ListOrgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgs`: TypesListUserOrgsResponse
    fmt.Fprintf(os.Stdout, "Response from `Users.ListOrgs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which we want to list the org. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListUserOrgsResponse**](TypesListUserOrgsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceAuditLogs

> TypesListAuditLogsResponse ListWorkspaceAuditLogs(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

User workspace audit logs



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
    userHandle := "userHandle_example" // string | Specify the user handle to get the audit logs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose logs needs to be fetched.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.ListWorkspaceAuditLogs(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.ListWorkspaceAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspaceAuditLogs`: TypesListAuditLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `Users.ListWorkspaceAuditLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the user handle to get the audit logs. | 
**workspaceHandle** | **string** | The handle of the workspace whose logs needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListAuditLogsResponse**](TypesListAuditLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> TypesSearchUsersResponse Search(ctx).Q(q).Limit(limit).NextToken(nextToken).Execute()

Search users



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
    q := "q_example" // string | Specify the search string.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.Search(context.Background()).Q(q).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: TypesSearchUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `Users.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Specify the search string. | 
 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesSearchUsersResponse**](TypesSearchUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> TypesUser Update(ctx, userHandle).Request(request).Execute()

Update user



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
    userHandle := "userHandle_example" // string | Specify the handle of the user which need to be updated.
    request := *openapiclient.NewTypesUpdateUserRequest() // TypesUpdateUserRequest | The request body for the user.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Users.Update(context.Background(), userHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Users.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: TypesUser
    fmt.Fprintf(os.Stdout, "Response from `Users.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user which need to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**TypesUpdateUserRequest**](TypesUpdateUserRequest.md) | The request body for the user. | 

### Return type

[**TypesUser**](TypesUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

