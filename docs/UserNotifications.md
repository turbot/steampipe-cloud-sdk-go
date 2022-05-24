# \UserNotifications

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserNotificationRule**](UserNotifications.md#CreateUserNotificationRule) | **Post** /user/{user_handle}/notification_rule | Create user notification rule
[**DeleteUserNotificationRule**](UserNotifications.md#DeleteUserNotificationRule) | **Delete** /user/{user_handle}/notification_rule/{notification_rule_id} | Delete user notification rule
[**GetUserNotificationRule**](UserNotifications.md#GetUserNotificationRule) | **Get** /user/{user_handle}/notification_rule/{notification_rule_id} | Get user notification rule
[**ListUserNotificationRules**](UserNotifications.md#ListUserNotificationRules) | **Get** /user/{user_handle}/notification_rule | List user notification rules



## CreateUserNotificationRule

> NotificationRule CreateUserNotificationRule(ctx, userHandle).Request(request).Execute()

Create user notification rule



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
    userHandle := "userHandle_example" // string | Specify the handle of the user whose notification rule will be created against.
    request := *openapiclient.NewCreateUserNotificationRequestNoSender([]string{"Events_example"}, "Title_example") // CreateUserNotificationRequestNoSender | The request body to create a notification rule for this user.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotifications.CreateUserNotificationRule(context.Background(), userHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotifications.CreateUserNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `UserNotifications.CreateUserNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user whose notification rule will be created against. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreateUserNotificationRequestNoSender**](CreateUserNotificationRequestNoSender.md) | The request body to create a notification rule for this user. | 

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


## DeleteUserNotificationRule

> NotificationRule DeleteUserNotificationRule(ctx, userHandle, notificationRuleId).Execute()

Delete user notification rule



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
    userHandle := "userHandle_example" // string | Specify the handle of the user where the notification exists
    notificationRuleId := "notificationRuleId_example" // string | The notification rule id to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotifications.DeleteUserNotificationRule(context.Background(), userHandle, notificationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotifications.DeleteUserNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `UserNotifications.DeleteUserNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user where the notification exists | 
**notificationRuleId** | **string** | The notification rule id to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserNotificationRuleRequest struct via the builder pattern


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


## GetUserNotificationRule

> NotificationRule GetUserNotificationRule(ctx, userHandle, notificationRuleId).Execute()

Get user notification rule



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
    userHandle := "userHandle_example" // string | Specify the handle of the user where the notification exists
    notificationRuleId := "notificationRuleId_example" // string | The notification rule id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotifications.GetUserNotificationRule(context.Background(), userHandle, notificationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotifications.GetUserNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `UserNotifications.GetUserNotificationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user where the notification exists | 
**notificationRuleId** | **string** | The notification rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNotificationRuleRequest struct via the builder pattern


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


## ListUserNotificationRules

> ListNotificationRulesResponse ListUserNotificationRules(ctx, userHandle).Limit(limit).NextToken(nextToken).Execute()

List user notification rules



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
    userHandle := "userHandle_example" // string | Specify the handle of the user where the notification exists
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotifications.ListUserNotificationRules(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotifications.ListUserNotificationRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserNotificationRules`: ListNotificationRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserNotifications.ListUserNotificationRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user where the notification exists | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserNotificationRulesRequest struct via the builder pattern


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

