# \Identities

All URIs are relative to *https://cloud.steampipe.io/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](Identities.md#Get) | **Get** /identity/search | Get identity



## Get

> Identity Get(ctx, identityHandle).Execute()

Get identity



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
    identityHandle := "identityHandle_example" // string | Specify the handle of the identity whose information you want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Identities.Get(context.Background(), identityHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Identities.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Identity
    fmt.Fprintf(os.Stdout, "Response from `Identities.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityHandle** | **string** | Specify the handle of the identity whose information you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Identity**](Identity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

