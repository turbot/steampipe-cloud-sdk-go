# \OrgWorkspaceNotifications

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWorkspaceNotificationRule**](OrgWorkspaceNotifications.md#CreateOrgWorkspaceNotificationRule) | **Post** /org/{org_handle}/workspace/{workspace_handle}/notification_rule | Create organization workspace notification rule
[**DeleteOrgWorkspaceNotificationRule**](OrgWorkspaceNotifications.md#DeleteOrgWorkspaceNotificationRule) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/notification_rule/{notification_rule_id} | Delete organization workspace notification rule
[**GetOrgWorkspaceNotificationRule**](OrgWorkspaceNotifications.md#GetOrgWorkspaceNotificationRule) | **Get** /org/{org_handle}/workspace/{workspace_handle}/notification_rule/{notification_rule_id} | Get organization workspace notification rule
[**ListOrgWorkspaceNotificationRules**](OrgWorkspaceNotifications.md#ListOrgWorkspaceNotificationRules) | **Get** /org/{org_handle}/workspace/{workspace_handle}/notification_rule | List organization workspace notification rules



## CreateOrgWorkspaceNotificationRule

> NotificationRule CreateOrgWorkspaceNotificationRule(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Create organization workspace notification rule



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the notification rule will be created.
    request := *openapiclient.NewCreateWorkspaceNotificationRequest([]string{"Events_example"}, "Title_example") // CreateWorkspaceNotificationRequest | The request body to create a notification rule for this organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceNotifications.CreateOrgWorkspaceNotificationRule(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceNotifications.CreateOrgWorkspaceNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgWorkspaceNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceNotifications.CreateOrgWorkspaceNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the notification rule will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgWorkspaceNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateWorkspaceNotificationRequest**](CreateWorkspaceNotificationRequest.md) | The request body to create a notification rule for this organization. | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgWorkspaceNotificationRule

> NotificationRule DeleteOrgWorkspaceNotificationRule(ctx, orgHandle, workspaceHandle, notificationRuleId).Execute()

Delete organization workspace notification rule



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the notification exists.
    notificationRuleId := "notificationRuleId_example" // string | The notification rule id to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceNotifications.DeleteOrgWorkspaceNotificationRule(context.Background(), orgHandle, workspaceHandle, notificationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceNotifications.DeleteOrgWorkspaceNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgWorkspaceNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceNotifications.DeleteOrgWorkspaceNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the notification exists. | 
**notificationRuleId** | **string** | The notification rule id to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgWorkspaceNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgWorkspaceNotificationRule

> NotificationRule GetOrgWorkspaceNotificationRule(ctx, orgHandle, workspaceHandle, notificationRuleId).Execute()

Get organization workspace notification rule



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the notification exists.
    notificationRuleId := "notificationRuleId_example" // string | The notification rule id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceNotifications.GetOrgWorkspaceNotificationRule(context.Background(), orgHandle, workspaceHandle, notificationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceNotifications.GetOrgWorkspaceNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgWorkspaceNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceNotifications.GetOrgWorkspaceNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the notification exists. | 
**notificationRuleId** | **string** | The notification rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgWorkspaceNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgWorkspaceNotificationRules

> ListNotificationRulesResponse ListOrgWorkspaceNotificationRules(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List organization workspace notification rules



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the notification exists.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceNotifications.ListOrgWorkspaceNotificationRules(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceNotifications.ListOrgWorkspaceNotificationRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgWorkspaceNotificationRules`: ListNotificationRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceNotifications.ListOrgWorkspaceNotificationRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the notification exists. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWorkspaceNotificationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListNotificationRulesResponse**](ListNotificationRulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

