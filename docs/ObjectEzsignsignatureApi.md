# eZmaxAPI\ObjectEzsignsignatureApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignsignatureCreateObjectV1**](ObjectEzsignsignatureApi.md#EzsignsignatureCreateObjectV1) | **Post** /1/object/ezsignsignature | Create a new Ezsignsignature
[**EzsignsignatureDeleteObjectV1**](ObjectEzsignsignatureApi.md#EzsignsignatureDeleteObjectV1) | **Delete** /1/object/ezsignsignature/{pkiEzsignsignatureID} | Delete an existing Ezsignsignature
[**EzsignsignatureEditObjectV1**](ObjectEzsignsignatureApi.md#EzsignsignatureEditObjectV1) | **Put** /1/object/ezsignsignature/{pkiEzsignsignatureID} | Modify an existing Ezsignsignature
[**EzsignsignatureGetObjectGetChildrenV1**](ObjectEzsignsignatureApi.md#EzsignsignatureGetObjectGetChildrenV1) | **Get** /1/object/ezsignsignature/{pkiEzsignsignatureID}/getChildren | Retrieve an existing Ezsignsignature&#39;s children IDs
[**EzsignsignatureGetObjectV1**](ObjectEzsignsignatureApi.md#EzsignsignatureGetObjectV1) | **Get** /1/object/ezsignsignature/{pkiEzsignsignatureID} | Retrieve an existing Ezsignsignature



## EzsignsignatureCreateObjectV1

> EzsignsignatureCreateObjectV1Response EzsignsignatureCreateObjectV1(ctx).EzsignsignatureCreateObjectV1Request(ezsignsignatureCreateObjectV1Request).Execute()

Create a new Ezsignsignature



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
    ezsignsignatureCreateObjectV1Request := []openapiclient.EzsignsignatureCreateObjectV1Request{*openapiclient.NewEzsignsignatureCreateObjectV1Request()} // []EzsignsignatureCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignsignatureApi.EzsignsignatureCreateObjectV1(context.Background()).EzsignsignatureCreateObjectV1Request(ezsignsignatureCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureApi.EzsignsignatureCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignsignatureCreateObjectV1`: EzsignsignatureCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureApi.EzsignsignatureCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignsignatureCreateObjectV1Request** | [**[]EzsignsignatureCreateObjectV1Request**](ezsignsignature-createObject-v1-Request.md) |  | 

### Return type

[**EzsignsignatureCreateObjectV1Response**](ezsignsignature-createObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignatureDeleteObjectV1

> EzsignsignatureDeleteObjectV1Response EzsignsignatureDeleteObjectV1(ctx, pkiEzsignsignatureID).Execute()

Delete an existing Ezsignsignature

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
    pkiEzsignsignatureID := int32(56) // int32 | The unique ID of the Ezsignsignature

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignsignatureApi.EzsignsignatureDeleteObjectV1(context.Background(), pkiEzsignsignatureID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureApi.EzsignsignatureDeleteObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignsignatureDeleteObjectV1`: EzsignsignatureDeleteObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureApi.EzsignsignatureDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignatureID** | **int32** | The unique ID of the Ezsignsignature | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignsignatureDeleteObjectV1Response**](ezsignsignature-deleteObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignatureEditObjectV1

> EzsignsignatureEditObjectV1Response EzsignsignatureEditObjectV1(ctx, pkiEzsignsignatureID).EzsignsignatureEditObjectV1Request(ezsignsignatureEditObjectV1Request).Execute()

Modify an existing Ezsignsignature

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
    pkiEzsignsignatureID := int32(56) // int32 | The unique ID of the Ezsignsignature
    ezsignsignatureEditObjectV1Request := *openapiclient.NewEzsignsignatureEditObjectV1Request() // EzsignsignatureEditObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignsignatureApi.EzsignsignatureEditObjectV1(context.Background(), pkiEzsignsignatureID).EzsignsignatureEditObjectV1Request(ezsignsignatureEditObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureApi.EzsignsignatureEditObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignsignatureEditObjectV1`: EzsignsignatureEditObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureApi.EzsignsignatureEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignatureID** | **int32** | The unique ID of the Ezsignsignature | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignsignatureEditObjectV1Request** | [**EzsignsignatureEditObjectV1Request**](EzsignsignatureEditObjectV1Request.md) |  | 

### Return type

[**EzsignsignatureEditObjectV1Response**](ezsignsignature-editObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignatureGetObjectGetChildrenV1

> EzsignsignatureGetObjectGetChildrenV1(ctx, pkiEzsignsignatureID).Execute()

Retrieve an existing Ezsignsignature's children IDs

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
    pkiEzsignsignatureID := int32(56) // int32 | The unique ID of the Ezsignsignature

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignsignatureApi.EzsignsignatureGetObjectGetChildrenV1(context.Background(), pkiEzsignsignatureID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureApi.EzsignsignatureGetObjectGetChildrenV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignatureID** | **int32** | The unique ID of the Ezsignsignature | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureGetObjectGetChildrenV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignsignatureGetObjectV1

> EzsignsignatureGetObjectV1Response EzsignsignatureGetObjectV1(ctx, pkiEzsignsignatureID).Execute()

Retrieve an existing Ezsignsignature

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
    pkiEzsignsignatureID := int32(56) // int32 | The unique ID of the Ezsignsignature

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignsignatureApi.EzsignsignatureGetObjectV1(context.Background(), pkiEzsignsignatureID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignsignatureApi.EzsignsignatureGetObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignsignatureGetObjectV1`: EzsignsignatureGetObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignsignatureApi.EzsignsignatureGetObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignsignatureID** | **int32** | The unique ID of the Ezsignsignature | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignsignatureGetObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignsignatureGetObjectV1Response**](ezsignsignature-getObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
