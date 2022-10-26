# \Auth

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmLogin**](Auth.md#ConfirmLogin) | **Get** /login/confirm | Confirm user login
[**ConfirmSignup**](Auth.md#ConfirmSignup) | **Get** /signup/confirm | Confirm user signup
[**Login**](Auth.md#Login) | **Post** /login | User login
[**LoginTokenCreate**](Auth.md#LoginTokenCreate) | **Post** /login/token | Generate temporary token request
[**LoginTokenDelete**](Auth.md#LoginTokenDelete) | **Get** /login/token/{temporary_token_request_id} | Delete temporary token request
[**LoginTokenUpdate**](Auth.md#LoginTokenUpdate) | **Patch** /login/token/{temporary_token_request_id} | Update temporary token request
[**Logout**](Auth.md#Logout) | **Get** /logout/{provider} | User logout
[**Provider**](Auth.md#Provider) | **Get** /auth/{provider} | Auth Provider
[**ProviderCallback**](Auth.md#ProviderCallback) | **Get** /auth/{provider}/callback | Auth provider callback
[**Signup**](Auth.md#Signup) | **Post** /signup | User signup



## ConfirmLogin

> ConfirmLogin(ctx).T(t).Execute()

Confirm user login



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
    t := "t_example" // string | Specify the token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.ConfirmLogin(context.Background()).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.ConfirmLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmLoginRequest struct via the builder pattern


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


## ConfirmSignup

> ConfirmSignup(ctx).T(t).Execute()

Confirm user signup



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
    t := "t_example" // string | Specify the token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.ConfirmSignup(context.Background()).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.ConfirmSignup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmSignupRequest struct via the builder pattern


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


## Login

> Login(ctx).Request(request).Execute()

User login



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
    request := *openapiclient.NewUserLoginRequest("Email_example") // UserLoginRequest | The request body to login.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.Login(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**UserLoginRequest**](UserLoginRequest.md) | The request body to login. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginTokenCreate

> TemporaryTokenRequest LoginTokenCreate(ctx).Execute()

Generate temporary token request



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.LoginTokenCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.LoginTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoginTokenCreate`: TemporaryTokenRequest
    fmt.Fprintf(os.Stdout, "Response from `Auth.LoginTokenCreate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoginTokenCreateRequest struct via the builder pattern


### Return type

[**TemporaryTokenRequest**](TemporaryTokenRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginTokenDelete

> TemporaryTokenRequest LoginTokenDelete(ctx, temporaryTokenRequestId).Execute()

Delete temporary token request



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
    temporaryTokenRequestId := "temporaryTokenRequestId_example" // string | The ID of the temporary token request to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.LoginTokenDelete(context.Background(), temporaryTokenRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.LoginTokenDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoginTokenDelete`: TemporaryTokenRequest
    fmt.Fprintf(os.Stdout, "Response from `Auth.LoginTokenDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**temporaryTokenRequestId** | **string** | The ID of the temporary token request to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoginTokenDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemporaryTokenRequest**](TemporaryTokenRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginTokenUpdate

> TemporaryTokenRequest LoginTokenUpdate(ctx, temporaryTokenRequestId).Request(request).Execute()

Update temporary token request



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
    temporaryTokenRequestId := "temporaryTokenRequestId_example" // string | The ID of the temporary token request to update.
    request := *openapiclient.NewUpdateTemporaryTokenRequest("State_example") // UpdateTemporaryTokenRequest | The request body containing updates for the temporary token request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.LoginTokenUpdate(context.Background(), temporaryTokenRequestId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.LoginTokenUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoginTokenUpdate`: TemporaryTokenRequest
    fmt.Fprintf(os.Stdout, "Response from `Auth.LoginTokenUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**temporaryTokenRequestId** | **string** | The ID of the temporary token request to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoginTokenUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**UpdateTemporaryTokenRequest**](UpdateTemporaryTokenRequest.md) | The request body containing updates for the temporary token request. | 

### Return type

[**TemporaryTokenRequest**](TemporaryTokenRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Logout(ctx, provider).Execute()

User logout



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
    provider := "provider_example" // string | The authentication provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.Logout(context.Background(), provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The authentication provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## Provider

> Provider(ctx, provider).Execute()

Auth Provider



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
    provider := "provider_example" // string | The authentication provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.Provider(context.Background(), provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.Provider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The authentication provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ProviderCallback

> ProviderCallback(ctx, provider).Execute()

Auth provider callback



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
    provider := "provider_example" // string | The authentication provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.ProviderCallback(context.Background(), provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.ProviderCallback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The authentication provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProviderCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## Signup

> Signup(ctx).Request(request).Execute()

User signup



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
    request := *openapiclient.NewUserSignupRequest("Email_example") // UserSignupRequest | The request body to signup.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Auth.Signup(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Auth.Signup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**UserSignupRequest**](UserSignupRequest.md) | The request body to signup. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

