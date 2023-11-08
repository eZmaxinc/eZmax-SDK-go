# eZmaxAPI\ObjectApikeyAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApikeyCreateObjectV2**](ObjectApikeyAPI.md#ApikeyCreateObjectV2) | **Post** /2/object/apikey | Create a new Apikey
[**ApikeyEditObjectV1**](ObjectApikeyAPI.md#ApikeyEditObjectV1) | **Put** /1/object/apikey/{pkiApikeyID} | Edit an existing Apikey
[**ApikeyEditPermissionsV1**](ObjectApikeyAPI.md#ApikeyEditPermissionsV1) | **Put** /1/object/apikey/{pkiApikeyID}/editPermissions | Edit multiple Permissions
[**ApikeyGetCorsV1**](ObjectApikeyAPI.md#ApikeyGetCorsV1) | **Get** /1/object/apikey/{pkiApikeyID}/getCors | Retrieve an existing Apikey&#39;s cors
[**ApikeyGetListV1**](ObjectApikeyAPI.md#ApikeyGetListV1) | **Get** /1/object/apikey/getList | Retrieve Apikey list
[**ApikeyGetObjectV2**](ObjectApikeyAPI.md#ApikeyGetObjectV2) | **Get** /2/object/apikey/{pkiApikeyID} | Retrieve an existing Apikey
[**ApikeyGetPermissionsV1**](ObjectApikeyAPI.md#ApikeyGetPermissionsV1) | **Get** /1/object/apikey/{pkiApikeyID}/getPermissions | Retrieve an existing Apikey&#39;s Permissions
[**ApikeyGetSubnetsV1**](ObjectApikeyAPI.md#ApikeyGetSubnetsV1) | **Get** /1/object/apikey/{pkiApikeyID}/getSubnets | Retrieve an existing Apikey&#39;s subnets
[**ApikeyRegenerateV1**](ObjectApikeyAPI.md#ApikeyRegenerateV1) | **Post** /1/object/apikey/{pkiApikeyID}/regenerate | Regenerate the Apikey



## ApikeyCreateObjectV2

> ApikeyCreateObjectV2Response ApikeyCreateObjectV2(ctx).ApikeyCreateObjectV2Request(apikeyCreateObjectV2Request).Execute()

Create a new Apikey



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
    apikeyCreateObjectV2Request := *openapiclient.NewApikeyCreateObjectV2Request([]openapiclient.ApikeyRequestCompound{*openapiclient.NewApikeyRequestCompound(int32(70), *openapiclient.NewMultilingualApikeyDescription())}) // ApikeyCreateObjectV2Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectApikeyAPI.ApikeyCreateObjectV2(context.Background()).ApikeyCreateObjectV2Request(apikeyCreateObjectV2Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectApikeyAPI.ApikeyCreateObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApikeyCreateObjectV2`: ApikeyCreateObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectApikeyAPI.ApikeyCreateObjectV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApikeyCreateObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apikeyCreateObjectV2Request** | [**ApikeyCreateObjectV2Request**](ApikeyCreateObjectV2Request.md) |  | 

### Return type

[**ApikeyCreateObjectV2Response**](ApikeyCreateObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApikeyEditObjectV1

> ApikeyEditObjectV1Response ApikeyEditObjectV1(ctx, pkiApikeyID).ApikeyEditObjectV1Request(apikeyEditObjectV1Request).Execute()

Edit an existing Apikey



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
    pkiApikeyID := int32(56) // int32 | The unique ID of the Apikey
    apikeyEditObjectV1Request := *openapiclient.NewApikeyEditObjectV1Request(*openapiclient.NewApikeyRequestCompound(int32(70), *openapiclient.NewMultilingualApikeyDescription())) // ApikeyEditObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectApikeyAPI.ApikeyEditObjectV1(context.Background(), pkiApikeyID).ApikeyEditObjectV1Request(apikeyEditObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectApikeyAPI.ApikeyEditObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApikeyEditObjectV1`: ApikeyEditObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectApikeyAPI.ApikeyEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiApikeyID** | **int32** | The unique ID of the Apikey | 

### Other Parameters

Other parameters are passed through a pointer to a apiApikeyEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apikeyEditObjectV1Request** | [**ApikeyEditObjectV1Request**](ApikeyEditObjectV1Request.md) |  | 

### Return type

[**ApikeyEditObjectV1Response**](ApikeyEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApikeyEditPermissionsV1

> ApikeyEditPermissionsV1Response ApikeyEditPermissionsV1(ctx, pkiApikeyID).ApikeyEditPermissionsV1Request(apikeyEditPermissionsV1Request).Execute()

Edit multiple Permissions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
    pkiApikeyID := int32(56) // int32 | 
    apikeyEditPermissionsV1Request := *openapiclient.NewApikeyEditPermissionsV1Request([]openapiclient.PermissionRequestCompound{*openapiclient.NewPermissionRequestCompound(int32(53))}) // ApikeyEditPermissionsV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectApikeyAPI.ApikeyEditPermissionsV1(context.Background(), pkiApikeyID).ApikeyEditPermissionsV1Request(apikeyEditPermissionsV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectApikeyAPI.ApikeyEditPermissionsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApikeyEditPermissionsV1`: ApikeyEditPermissionsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectApikeyAPI.ApikeyEditPermissionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiApikeyID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApikeyEditPermissionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apikeyEditPermissionsV1Request** | [**ApikeyEditPermissionsV1Request**](ApikeyEditPermissionsV1Request.md) |  | 

### Return type

[**ApikeyEditPermissionsV1Response**](ApikeyEditPermissionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApikeyGetCorsV1

> ApikeyGetCorsV1Response ApikeyGetCorsV1(ctx, pkiApikeyID).Execute()

Retrieve an existing Apikey's cors

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
    pkiApikeyID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectApikeyAPI.ApikeyGetCorsV1(context.Background(), pkiApikeyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectApikeyAPI.ApikeyGetCorsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApikeyGetCorsV1`: ApikeyGetCorsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectApikeyAPI.ApikeyGetCorsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiApikeyID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApikeyGetCorsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApikeyGetCorsV1Response**](ApikeyGetCorsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApikeyGetListV1

> ApikeyGetListV1Response ApikeyGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Apikey list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
    eOrderBy := "eOrderBy_example" // string | Specify how you want the results to be sorted (optional)
    iRowMax := int32(56) // int32 |  (optional)
    iRowOffset := int32(56) // int32 |  (optional) (default to 0)
    acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)
    sFilter := "sFilter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectApikeyAPI.ApikeyGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectApikeyAPI.ApikeyGetListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApikeyGetListV1`: ApikeyGetListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectApikeyAPI.ApikeyGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApikeyGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**ApikeyGetListV1Response**](ApikeyGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApikeyGetObjectV2

> ApikeyGetObjectV2Response ApikeyGetObjectV2(ctx, pkiApikeyID).Execute()

Retrieve an existing Apikey



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
    pkiApikeyID := int32(56) // int32 | The unique ID of the Apikey

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectApikeyAPI.ApikeyGetObjectV2(context.Background(), pkiApikeyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectApikeyAPI.ApikeyGetObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApikeyGetObjectV2`: ApikeyGetObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectApikeyAPI.ApikeyGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiApikeyID** | **int32** | The unique ID of the Apikey | 

### Other Parameters

Other parameters are passed through a pointer to a apiApikeyGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApikeyGetObjectV2Response**](ApikeyGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApikeyGetPermissionsV1

> ApikeyGetPermissionsV1Response ApikeyGetPermissionsV1(ctx, pkiApikeyID).Execute()

Retrieve an existing Apikey's Permissions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
    pkiApikeyID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectApikeyAPI.ApikeyGetPermissionsV1(context.Background(), pkiApikeyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectApikeyAPI.ApikeyGetPermissionsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApikeyGetPermissionsV1`: ApikeyGetPermissionsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectApikeyAPI.ApikeyGetPermissionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiApikeyID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApikeyGetPermissionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApikeyGetPermissionsV1Response**](ApikeyGetPermissionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApikeyGetSubnetsV1

> ApikeyGetSubnetsV1Response ApikeyGetSubnetsV1(ctx, pkiApikeyID).Execute()

Retrieve an existing Apikey's subnets

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
    pkiApikeyID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectApikeyAPI.ApikeyGetSubnetsV1(context.Background(), pkiApikeyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectApikeyAPI.ApikeyGetSubnetsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApikeyGetSubnetsV1`: ApikeyGetSubnetsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectApikeyAPI.ApikeyGetSubnetsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiApikeyID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApikeyGetSubnetsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApikeyGetSubnetsV1Response**](ApikeyGetSubnetsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApikeyRegenerateV1

> ApikeyRegenerateV1Response ApikeyRegenerateV1(ctx, pkiApikeyID).ApikeyRegenerateV1Request(apikeyRegenerateV1Request).Execute()

Regenerate the Apikey



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
    pkiApikeyID := int32(56) // int32 | 
    apikeyRegenerateV1Request := *openapiclient.NewApikeyRegenerateV1Request() // ApikeyRegenerateV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectApikeyAPI.ApikeyRegenerateV1(context.Background(), pkiApikeyID).ApikeyRegenerateV1Request(apikeyRegenerateV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectApikeyAPI.ApikeyRegenerateV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApikeyRegenerateV1`: ApikeyRegenerateV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectApikeyAPI.ApikeyRegenerateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiApikeyID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApikeyRegenerateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apikeyRegenerateV1Request** | [**ApikeyRegenerateV1Request**](ApikeyRegenerateV1Request.md) |  | 

### Return type

[**ApikeyRegenerateV1Response**](ApikeyRegenerateV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

