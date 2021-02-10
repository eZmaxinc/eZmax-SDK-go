# eZmaxAPI\ObjectEzsignfolderApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignfolderCreateObjectV1**](ObjectEzsignfolderApi.md#EzsignfolderCreateObjectV1) | **Post** /1/object/ezsignfolder | Create a new Ezsignfolder
[**EzsignfolderDeleteObjectV1**](ObjectEzsignfolderApi.md#EzsignfolderDeleteObjectV1) | **Delete** /1/object/ezsignfolder/{pkiEzsignfolderID} | Delete an existing Ezsignfolder
[**EzsignfolderEditObjectV1**](ObjectEzsignfolderApi.md#EzsignfolderEditObjectV1) | **Put** /1/object/ezsignfolder/{pkiEzsignfolderID} | Modify an existing Ezsignfolder
[**EzsignfolderGetChildrenV1**](ObjectEzsignfolderApi.md#EzsignfolderGetChildrenV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getChildren | Retrieve an existing Ezsignfolder&#39;s children IDs
[**EzsignfolderGetObjectV1**](ObjectEzsignfolderApi.md#EzsignfolderGetObjectV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID} | Retrieve an existing Ezsignfolder
[**EzsignfolderSendV1**](ObjectEzsignfolderApi.md#EzsignfolderSendV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/send | Send the Ezsignfolder to the signatories for signature



## EzsignfolderCreateObjectV1

> EzsignfolderCreateObjectV1Response EzsignfolderCreateObjectV1(ctx).EzsignfolderCreateObjectV1Request(ezsignfolderCreateObjectV1Request).Execute()

Create a new Ezsignfolder



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
    ezsignfolderCreateObjectV1Request := []openapiclient.EzsignfolderCreateObjectV1Request{*openapiclient.NewEzsignfolderCreateObjectV1Request()} // []EzsignfolderCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfolderApi.EzsignfolderCreateObjectV1(context.Background()).EzsignfolderCreateObjectV1Request(ezsignfolderCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderCreateObjectV1`: EzsignfolderCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderApi.EzsignfolderCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignfolderCreateObjectV1Request** | [**[]EzsignfolderCreateObjectV1Request**](ezsignfolder-createObject-v1-Request.md) |  | 

### Return type

[**EzsignfolderCreateObjectV1Response**](ezsignfolder-createObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderDeleteObjectV1

> EzsignfolderDeleteObjectV1Response EzsignfolderDeleteObjectV1(ctx, pkiEzsignfolderID).Execute()

Delete an existing Ezsignfolder

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
    pkiEzsignfolderID := int32(56) // int32 | The unique ID of the Ezsignfolder

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfolderApi.EzsignfolderDeleteObjectV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderDeleteObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderDeleteObjectV1`: EzsignfolderDeleteObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderApi.EzsignfolderDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderDeleteObjectV1Response**](ezsignfolder-deleteObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderEditObjectV1

> EzsignfolderEditObjectV1Response EzsignfolderEditObjectV1(ctx, pkiEzsignfolderID).EzsignfolderEditObjectV1Request(ezsignfolderEditObjectV1Request).Execute()

Modify an existing Ezsignfolder

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
    pkiEzsignfolderID := int32(56) // int32 | The unique ID of the Ezsignfolder
    ezsignfolderEditObjectV1Request := *openapiclient.NewEzsignfolderEditObjectV1Request() // EzsignfolderEditObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfolderApi.EzsignfolderEditObjectV1(context.Background(), pkiEzsignfolderID).EzsignfolderEditObjectV1Request(ezsignfolderEditObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderEditObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderEditObjectV1`: EzsignfolderEditObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderApi.EzsignfolderEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderEditObjectV1Request** | [**EzsignfolderEditObjectV1Request**](EzsignfolderEditObjectV1Request.md) |  | 

### Return type

[**EzsignfolderEditObjectV1Response**](ezsignfolder-editObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetChildrenV1

> EzsignfolderGetChildrenV1(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder's children IDs

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
    pkiEzsignfolderID := int32(56) // int32 | The unique ID of the Ezsignfolder

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfolderApi.EzsignfolderGetChildrenV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderGetChildrenV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetChildrenV1Request struct via the builder pattern


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


## EzsignfolderGetObjectV1

> EzsignfolderGetObjectV1Response EzsignfolderGetObjectV1(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder

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
    pkiEzsignfolderID := int32(56) // int32 | The unique ID of the Ezsignfolder

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfolderApi.EzsignfolderGetObjectV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderGetObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetObjectV1`: EzsignfolderGetObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderApi.EzsignfolderGetObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetObjectV1Response**](ezsignfolder-getObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderSendV1

> EzsignfolderSendV1Response EzsignfolderSendV1(ctx, pkiEzsignfolderID).EzsignfolderSendV1Request(ezsignfolderSendV1Request).Execute()

Send the Ezsignfolder to the signatories for signature

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
    pkiEzsignfolderID := int32(56) // int32 | The unique ID of the Ezsignfolder
    ezsignfolderSendV1Request := *openapiclient.NewEzsignfolderSendV1Request("TExtraMessage_example") // EzsignfolderSendV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfolderApi.EzsignfolderSendV1(context.Background(), pkiEzsignfolderID).EzsignfolderSendV1Request(ezsignfolderSendV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderSendV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderSendV1`: EzsignfolderSendV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderApi.EzsignfolderSendV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** | The unique ID of the Ezsignfolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderSendV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderSendV1Request** | [**EzsignfolderSendV1Request**](EzsignfolderSendV1Request.md) |  | 

### Return type

[**EzsignfolderSendV1Response**](ezsignfolder-send-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

