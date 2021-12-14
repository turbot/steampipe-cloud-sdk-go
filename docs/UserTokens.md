# \UserTokens

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserTokens.md#Create) | **Post** /user/{user_handle}/token | Create token
[**Delete**](UserTokens.md#Delete) | **Delete** /user/{user_handle}/token/{token_id} | Delete token
[**Get**](UserTokens.md#Get) | **Get** /user/{user_handle}/token/{token_id} | Get token
[**List**](UserTokens.md#List) | **Get** /user/{user_handle}/token | List tokens
[**Update**](UserTokens.md#Update) | **Patch** /user/{user_handle}/token/{token_id} | Update token



## Create

> Token Create(ctx, userHandle).Execute()

Create token



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
    userHandle := "userHandle_example" // string | The handle of the user where the token will be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserTokens.Create(context.Background(), userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokens.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Token
    fmt.Fprintf(os.Stdout, "Response from `UserTokens.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the token will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Token Delete(ctx, userHandle, tokenId).Execute()

Delete token



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
    userHandle := "userHandle_example" // string | The handle of the user where the token exist.
    tokenId := "tokenId_example" // string | The identifier of the token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserTokens.Delete(context.Background(), userHandle, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokens.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Token
    fmt.Fprintf(os.Stdout, "Response from `UserTokens.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the token exist. | 
**tokenId** | **string** | The identifier of the token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Token Get(ctx, tokenId, userHandle).Execute()

Get token



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
    tokenId := "tokenId_example" // string | The identifier of the token.
    userHandle := "userHandle_example" // string | The handle of the user where the token exist.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserTokens.Get(context.Background(), tokenId, userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokens.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Token
    fmt.Fprintf(os.Stdout, "Response from `UserTokens.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | The identifier of the token. | 
**userHandle** | **string** | The handle of the user where the token exist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListTokensResponse List(ctx, userHandle).Limit(limit).NextToken(nextToken).Execute()

List tokens



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
    userHandle := "userHandle_example" // string | The handle of the user.
    limit := int32(56) // int32 | Pagination limit (optional) (default to 20)
    nextToken := "nextToken_example" // string | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserTokens.List(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokens.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListTokensResponse
    fmt.Fprintf(os.Stdout, "Response from `UserTokens.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Pagination limit | [default to 20]
 **nextToken** | **string** | When a list is truncated this element specifies the last part of the list, as well as the value to use for the part-number-marker request parameter in a subsequent request. | 

### Return type

[**ListTokensResponse**](ListTokensResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Token Update(ctx, userHandle, tokenId).Request(request).Execute()

Update token



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
    userHandle := "userHandle_example" // string | The handle of the user where the token exist.
    tokenId := "tokenId_example" // string | The identifier of the token.
    request := *openapiclient.NewUpdateTokenRequest("Status_example") // UpdateTokenRequest | The request body for the token which needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserTokens.Update(context.Background(), userHandle, tokenId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserTokens.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Token
    fmt.Fprintf(os.Stdout, "Response from `UserTokens.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the token exist. | 
**tokenId** | **string** | The identifier of the token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateTokenRequest**](UpdateTokenRequest.md) | The request body for the token which needs to be updated. | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

