# \OrgNotifications

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNotificationRule**](OrgNotifications.md#CreateOrgNotificationRule) | **Post** /org/{org_handle}/notification_rule | Create organization notification rule
[**DeleteOrgNotificationRule**](OrgNotifications.md#DeleteOrgNotificationRule) | **Delete** /user/{org_handle}/notification_rule/{notification_rule_id} | Delete organization notification rule
[**GetOrgNotificationRule**](OrgNotifications.md#GetOrgNotificationRule) | **Get** /user/{org_handle}/notification_rule/{notification_rule_id} | Get organization notification rule
[**ListOrgNotificationRules**](OrgNotifications.md#ListOrgNotificationRules) | **Get** /user/{org_handle}/notification_rule | List organization notification rules



## CreateOrgNotificationRule

> NotificationRule CreateOrgNotificationRule(ctx, orgHandle).Request(request).Execute()

Create organization notification rule



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
    orgHandle := "orgHandle_example" // string | Specify the handle of an organization where the notification rule to be created.
    request := *openapiclient.NewCreateOrgNotificationRequest([]string{"Events_example"}, "Title_example") // CreateOrgNotificationRequest | The request body to create a notification rule for this organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgNotifications.CreateOrgNotificationRule(context.Background(), orgHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgNotifications.CreateOrgNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `OrgNotifications.CreateOrgNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of an organization where the notification rule to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreateOrgNotificationRequest**](CreateOrgNotificationRequest.md) | The request body to create a notification rule for this organization. | 

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


## DeleteOrgNotificationRule

> NotificationRule DeleteOrgNotificationRule(ctx, orgHandle, notificationRuleId).Execute()

Delete organization notification rule



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
    orgHandle := "orgHandle_example" // string | Specify the handle of an organization where the notification rule exists.
    notificationRuleId := "notificationRuleId_example" // string | The notification rule id to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgNotifications.DeleteOrgNotificationRule(context.Background(), orgHandle, notificationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgNotifications.DeleteOrgNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `OrgNotifications.DeleteOrgNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of an organization where the notification rule exists. | 
**notificationRuleId** | **string** | The notification rule id to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgNotificationRuleRequest struct via the builder pattern


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


## GetOrgNotificationRule

> NotificationRule GetOrgNotificationRule(ctx, orgHandle, notificationRuleId).Execute()

Get organization notification rule



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
    orgHandle := "orgHandle_example" // string | Specify the handle of an organization where the notification rule exists.
    notificationRuleId := "notificationRuleId_example" // string | The notification rule id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgNotifications.GetOrgNotificationRule(context.Background(), orgHandle, notificationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgNotifications.GetOrgNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `OrgNotifications.GetOrgNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of an organization where the notification rule exists. | 
**notificationRuleId** | **string** | The notification rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgNotificationRuleRequest struct via the builder pattern


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


## ListOrgNotificationRules

> ListNotificationRulesResponse ListOrgNotificationRules(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

List organization notification rules



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
    orgHandle := "orgHandle_example" // string | Specify the handle of an organization where the notification rules exist.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgNotifications.ListOrgNotificationRules(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgNotifications.ListOrgNotificationRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgNotificationRules`: ListNotificationRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgNotifications.ListOrgNotificationRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | Specify the handle of an organization where the notification rules exist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgNotificationRulesRequest struct via the builder pattern


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

