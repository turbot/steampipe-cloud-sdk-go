# \OrgsApi

All URIs are relative to *https://cloud.steampipe.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrg**](OrgsApi.md#CreateOrg) | **Post** /org | Create org
[**DeleteOrg**](OrgsApi.md#DeleteOrg) | **Delete** /org/{org_handle} | Delete org
[**GetOrg**](OrgsApi.md#GetOrg) | **Get** /org/{org_handle} | Get org
[**ListOrgAuditLogs**](OrgsApi.md#ListOrgAuditLogs) | **Get** /org/{org_handle}/audit | Org audit logs
[**ListOrgWorkspaceAuditLogs**](OrgsApi.md#ListOrgWorkspaceAuditLogs) | **Get** /org/{org_handle}/workspace/{workspace_handle}/audit | Org workspace audit logs
[**ListOrgWorkspaceLogs**](OrgsApi.md#ListOrgWorkspaceLogs) | **Get** /org/{org_handle}/workspace/{workspace_handle}/logs | Org workspace logs
[**ListOrgs**](OrgsApi.md#ListOrgs) | **Get** /org | List orgs
[**OrgQuota**](OrgsApi.md#OrgQuota) | **Get** /org/{org_handle}/quota | Org quota
[**UpdateOrg**](OrgsApi.md#UpdateOrg) | **Patch** /org/{org_handle} | Update org



## CreateOrg

> TypesOrg CreateOrg(ctx).Request(request).Execute()

Create org



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
    request := *openapiclient.NewTypesCreateOrgRequest("Handle_example") // TypesCreateOrgRequest | The request body to create the organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgsApi.CreateOrg(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.CreateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrg`: TypesOrg
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.CreateOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**TypesCreateOrgRequest**](TypesCreateOrgRequest.md) | The request body to create the organization. | 

### Return type

[**TypesOrg**](TypesOrg.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrg

> TypesOrg DeleteOrg(ctx, orgHandle).Execute()

Delete org



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
    orgHandle := "orgHandle_example" // string | Specify the handle of the organization which need to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgsApi.DeleteOrg(context.Background(), orgHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.DeleteOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrg`: TypesOrg
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.DeleteOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization which need to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesOrg**](TypesOrg.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrg

> TypesOrg GetOrg(ctx, orgHandle).Execute()

Get org



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
    orgHandle := "orgHandle_example" // string | Specify the handle of an organization whose information you want to retrive.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgsApi.GetOrg(context.Background(), orgHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.GetOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrg`: TypesOrg
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.GetOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of an organization whose information you want to retrive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesOrg**](TypesOrg.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgAuditLogs

> TypesListAuditLogsResponse ListOrgAuditLogs(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

Org audit logs



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
    orgHandle := "orgHandle_example" // string | Specify the org handle to get the audit logs.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgsApi.ListOrgAuditLogs(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.ListOrgAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgAuditLogs`: TypesListAuditLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.ListOrgAuditLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the org handle to get the audit logs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAuditLogsRequest struct via the builder pattern


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


## ListOrgWorkspaceAuditLogs

> TypesListAuditLogsResponse ListOrgWorkspaceAuditLogs(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

Org workspace audit logs



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
    orgHandle := "orgHandle_example" // string | Specify the org handle to get the audit logs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose logs needs to be fetched.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | An optional token returned from a prior request. When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgsApi.ListOrgWorkspaceAuditLogs(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.ListOrgWorkspaceAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgWorkspaceAuditLogs`: TypesListAuditLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.ListOrgWorkspaceAuditLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the org handle to get the audit logs. | 
**workspaceHandle** | **string** | The handle of the workspace whose logs needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWorkspaceAuditLogsRequest struct via the builder pattern


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


## ListOrgWorkspaceLogs

> TypesListLogsResponse ListOrgWorkspaceLogs(ctx, orgHandle, workspaceHandle).Execute()

Org workspace logs



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
    orgHandle := "orgHandle_example" // string | Specify the org handle to get the workspace logs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose logs needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgsApi.ListOrgWorkspaceLogs(context.Background(), orgHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.ListOrgWorkspaceLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgWorkspaceLogs`: TypesListLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.ListOrgWorkspaceLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the org handle to get the workspace logs. | 
**workspaceHandle** | **string** | The handle of the workspace whose logs needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWorkspaceLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TypesListLogsResponse**](TypesListLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgs

> TypesListOrgsResponse ListOrgs(ctx).Limit(limit).NextToken(nextToken).Execute()

List orgs



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
    resp, r, err := api_client.OrgsApi.ListOrgs(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.ListOrgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgs`: TypesListOrgsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.ListOrgs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**TypesListOrgsResponse**](TypesListOrgsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgQuota

> TypesOrgQuota OrgQuota(ctx, orgHandle).Execute()

Org quota



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
    orgHandle := "orgHandle_example" // string | Specify the org handle to get the quota details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgsApi.OrgQuota(context.Background(), orgHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.OrgQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgQuota`: TypesOrgQuota
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.OrgQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the org handle to get the quota details. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TypesOrgQuota**](TypesOrgQuota.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrg

> TypesOrg UpdateOrg(ctx, orgHandle).Request(request).Execute()

Update org



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
    orgHandle := "orgHandle_example" // string | Specify the handle of the organization which need to be updated.
    request := *openapiclient.NewTypesUpdateOrgRequest() // TypesUpdateOrgRequest | The request body for the organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgsApi.UpdateOrg(context.Background(), orgHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.UpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrg`: TypesOrg
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.UpdateOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of the organization which need to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**TypesUpdateOrgRequest**](TypesUpdateOrgRequest.md) | The request body for the organization. | 

### Return type

[**TypesOrg**](TypesOrg.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

