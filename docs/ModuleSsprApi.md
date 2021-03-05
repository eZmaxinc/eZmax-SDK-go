# eZmaxAPI\ModuleSsprApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SsprResetPasswordRequestV1**](ModuleSsprApi.md#SsprResetPasswordRequestV1) | **Post** /1/module/sspr/resetPasswordRequest/ | Reset Password Request
[**SsprResetPasswordV1**](ModuleSsprApi.md#SsprResetPasswordV1) | **Post** /1/module/sspr/resetPassword | Reset Password
[**SsprSendUsernamesV1**](ModuleSsprApi.md#SsprSendUsernamesV1) | **Post** /1/module/sspr/sendUsernames | Send username(s)
[**SsprUnlockAccountRequestV1**](ModuleSsprApi.md#SsprUnlockAccountRequestV1) | **Post** /1/module/sspr/unlockAccountRequest | Unlock Account Request
[**SsprUnlockAccountV1**](ModuleSsprApi.md#SsprUnlockAccountV1) | **Post** /1/module/sspr/unlockAccount | Unlock Account
[**SsprValidateTokenV1**](ModuleSsprApi.md#SsprValidateTokenV1) | **Post** /1/module/sspr/validateToken | Validate Token



## SsprResetPasswordRequestV1

> SsprResetPasswordRequestV1(ctx).SsprResetPasswordRequestV1Request(ssprResetPasswordRequestV1Request).Execute()

Reset Password Request



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
    ssprResetPasswordRequestV1Request := *openapiclient.NewSsprResetPasswordRequestV1Request("demo", int32(2), openapiclient.Field-eUserTypeSSPR("EzsignUser")) // SsprResetPasswordRequestV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModuleSsprApi.SsprResetPasswordRequestV1(context.Background()).SsprResetPasswordRequestV1Request(ssprResetPasswordRequestV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleSsprApi.SsprResetPasswordRequestV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsprResetPasswordRequestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssprResetPasswordRequestV1Request** | [**SsprResetPasswordRequestV1Request**](SsprResetPasswordRequestV1Request.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SsprResetPasswordV1

> SsprResetPasswordV1(ctx).SsprResetPasswordV1Request(ssprResetPasswordV1Request).Execute()

Reset Password



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
    ssprResetPasswordV1Request := *openapiclient.NewSsprResetPasswordV1Request("demo", int32(2), openapiclient.Field-eUserTypeSSPR("EzsignUser"), "012345678901234567890123456789ab", "Qwerty1234!") // SsprResetPasswordV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModuleSsprApi.SsprResetPasswordV1(context.Background()).SsprResetPasswordV1Request(ssprResetPasswordV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleSsprApi.SsprResetPasswordV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsprResetPasswordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssprResetPasswordV1Request** | [**SsprResetPasswordV1Request**](SsprResetPasswordV1Request.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SsprSendUsernamesV1

> SsprSendUsernamesV1(ctx).SsprSendUsernamesV1Request(ssprSendUsernamesV1Request).Execute()

Send username(s)



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
    ssprSendUsernamesV1Request := *openapiclient.NewSsprSendUsernamesV1Request("demo", int32(2), openapiclient.Field-eUserTypeSSPR("EzsignUser"), "example@domain.com") // SsprSendUsernamesV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModuleSsprApi.SsprSendUsernamesV1(context.Background()).SsprSendUsernamesV1Request(ssprSendUsernamesV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleSsprApi.SsprSendUsernamesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsprSendUsernamesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssprSendUsernamesV1Request** | [**SsprSendUsernamesV1Request**](SsprSendUsernamesV1Request.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SsprUnlockAccountRequestV1

> SsprUnlockAccountRequestV1(ctx).SsprUnlockAccountRequestV1Request(ssprUnlockAccountRequestV1Request).Execute()

Unlock Account Request



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
    ssprUnlockAccountRequestV1Request := *openapiclient.NewSsprUnlockAccountRequestV1Request("demo", int32(2), openapiclient.Field-eUserTypeSSPR("EzsignUser")) // SsprUnlockAccountRequestV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModuleSsprApi.SsprUnlockAccountRequestV1(context.Background()).SsprUnlockAccountRequestV1Request(ssprUnlockAccountRequestV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleSsprApi.SsprUnlockAccountRequestV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsprUnlockAccountRequestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssprUnlockAccountRequestV1Request** | [**SsprUnlockAccountRequestV1Request**](SsprUnlockAccountRequestV1Request.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SsprUnlockAccountV1

> SsprUnlockAccountV1(ctx).SsprUnlockAccountV1Request(ssprUnlockAccountV1Request).Execute()

Unlock Account



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
    ssprUnlockAccountV1Request := *openapiclient.NewSsprUnlockAccountV1Request("demo", int32(2), openapiclient.Field-eUserTypeSSPR("EzsignUser"), "012345678901234567890123456789ab") // SsprUnlockAccountV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModuleSsprApi.SsprUnlockAccountV1(context.Background()).SsprUnlockAccountV1Request(ssprUnlockAccountV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleSsprApi.SsprUnlockAccountV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsprUnlockAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssprUnlockAccountV1Request** | [**SsprUnlockAccountV1Request**](SsprUnlockAccountV1Request.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SsprValidateTokenV1

> SsprValidateTokenV1(ctx).SsprValidateTokenV1Request(ssprValidateTokenV1Request).Execute()

Validate Token



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
    ssprValidateTokenV1Request := *openapiclient.NewSsprValidateTokenV1Request("demo", int32(2), openapiclient.Field-eUserTypeSSPR("EzsignUser"), "012345678901234567890123456789ab") // SsprValidateTokenV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ModuleSsprApi.SsprValidateTokenV1(context.Background()).SsprValidateTokenV1Request(ssprValidateTokenV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ModuleSsprApi.SsprValidateTokenV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsprValidateTokenV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssprValidateTokenV1Request** | [**SsprValidateTokenV1Request**](SsprValidateTokenV1Request.md) |  | 

### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

