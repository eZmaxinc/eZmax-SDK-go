# eZmaxAPI\ModuleAuthenticateApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateAuthenticateV2**](ModuleAuthenticateApi.md#AuthenticateAuthenticateV2) | **Post** /2/module/authenticate/authenticate/ezsignuser/{eSessionType} | Authenticate a user



## AuthenticateAuthenticateV2

> AuthenticateAuthenticateV2Response AuthenticateAuthenticateV2(ctx, eSessionType).AuthenticateAuthenticateV2Request(authenticateAuthenticateV2Request).Execute()

Authenticate a user



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
    eSessionType := "eSessionType_example" // string | 
    authenticateAuthenticateV2Request := *openapiclient.NewAuthenticateAuthenticateV2Request("demo", "Qwerty1234!") // AuthenticateAuthenticateV2Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModuleAuthenticateApi.AuthenticateAuthenticateV2(context.Background(), eSessionType).AuthenticateAuthenticateV2Request(authenticateAuthenticateV2Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleAuthenticateApi.AuthenticateAuthenticateV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateAuthenticateV2`: AuthenticateAuthenticateV2Response
    fmt.Fprintf(os.Stdout, "Response from `ModuleAuthenticateApi.AuthenticateAuthenticateV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eSessionType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateAuthenticateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticateAuthenticateV2Request** | [**AuthenticateAuthenticateV2Request**](AuthenticateAuthenticateV2Request.md) |  | 

### Return type

[**AuthenticateAuthenticateV2Response**](AuthenticateAuthenticateV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

