# eZmaxAPI\ObjectEzsignfoldersignerassociationApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignfoldersignerassociationCreateObjectV1**](ObjectEzsignfoldersignerassociationApi.md#EzsignfoldersignerassociationCreateObjectV1) | **Post** /1/object/ezsignfoldersignerassociation | Create a new Ezsignfoldersignerassociation
[**EzsignfoldersignerassociationDeleteObjectV1**](ObjectEzsignfoldersignerassociationApi.md#EzsignfoldersignerassociationDeleteObjectV1) | **Delete** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Delete an existing Ezsignfoldersignerassociation
[**EzsignfoldersignerassociationGetChildrenV1**](ObjectEzsignfoldersignerassociationApi.md#EzsignfoldersignerassociationGetChildrenV1) | **Get** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/getChildren | Retrieve an existing Ezsignfoldersignerassociation&#39;s children IDs
[**EzsignfoldersignerassociationGetInPersonLoginUrlV1**](ObjectEzsignfoldersignerassociationApi.md#EzsignfoldersignerassociationGetInPersonLoginUrlV1) | **Get** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID}/getInPersonLoginUrl | Retrieve a Login Url to allow In-Person signing
[**EzsignfoldersignerassociationGetObjectV1**](ObjectEzsignfoldersignerassociationApi.md#EzsignfoldersignerassociationGetObjectV1) | **Get** /1/object/ezsignfoldersignerassociation/{pkiEzsignfoldersignerassociationID} | Retrieve an existing Ezsignfoldersignerassociation



## EzsignfoldersignerassociationCreateObjectV1

> EzsignfoldersignerassociationCreateObjectV1Response EzsignfoldersignerassociationCreateObjectV1(ctx).EzsignfoldersignerassociationCreateObjectV1Request(ezsignfoldersignerassociationCreateObjectV1Request).Execute()

Create a new Ezsignfoldersignerassociation



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
    ezsignfoldersignerassociationCreateObjectV1Request := []openapiclient.EzsignfoldersignerassociationCreateObjectV1Request{*openapiclient.NewEzsignfoldersignerassociationCreateObjectV1Request()} // []EzsignfoldersignerassociationCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationCreateObjectV1(context.Background()).EzsignfoldersignerassociationCreateObjectV1Request(ezsignfoldersignerassociationCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfoldersignerassociationCreateObjectV1`: EzsignfoldersignerassociationCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsignfoldersignerassociationCreateObjectV1Request** | [**[]EzsignfoldersignerassociationCreateObjectV1Request**](EzsignfoldersignerassociationCreateObjectV1Request.md) |  | 

### Return type

[**EzsignfoldersignerassociationCreateObjectV1Response**](EzsignfoldersignerassociationCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationDeleteObjectV1

> EzsignfoldersignerassociationDeleteObjectV1Response EzsignfoldersignerassociationDeleteObjectV1(ctx, pkiEzsignfoldersignerassociationID).Execute()

Delete an existing Ezsignfoldersignerassociation

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
    pkiEzsignfoldersignerassociationID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationDeleteObjectV1(context.Background(), pkiEzsignfoldersignerassociationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationDeleteObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfoldersignerassociationDeleteObjectV1`: EzsignfoldersignerassociationDeleteObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfoldersignerassociationDeleteObjectV1Response**](EzsignfoldersignerassociationDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationGetChildrenV1

> EzsignfoldersignerassociationGetChildrenV1(ctx, pkiEzsignfoldersignerassociationID).Execute()

Retrieve an existing Ezsignfoldersignerassociation's children IDs



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
    pkiEzsignfoldersignerassociationID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationGetChildrenV1(context.Background(), pkiEzsignfoldersignerassociationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationGetChildrenV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationGetChildrenV1Request struct via the builder pattern


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


## EzsignfoldersignerassociationGetInPersonLoginUrlV1

> EzsignfoldersignerassociationGetInPersonLoginUrlV1Response EzsignfoldersignerassociationGetInPersonLoginUrlV1(ctx, pkiEzsignfoldersignerassociationID).Execute()

Retrieve a Login Url to allow In-Person signing



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
    pkiEzsignfoldersignerassociationID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationGetInPersonLoginUrlV1(context.Background(), pkiEzsignfoldersignerassociationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationGetInPersonLoginUrlV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfoldersignerassociationGetInPersonLoginUrlV1`: EzsignfoldersignerassociationGetInPersonLoginUrlV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationGetInPersonLoginUrlV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationGetInPersonLoginUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfoldersignerassociationGetInPersonLoginUrlV1Response**](EzsignfoldersignerassociationGetInPersonLoginUrlV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldersignerassociationGetObjectV1

> EzsignfoldersignerassociationGetObjectV1Response EzsignfoldersignerassociationGetObjectV1(ctx, pkiEzsignfoldersignerassociationID).Execute()

Retrieve an existing Ezsignfoldersignerassociation



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
    pkiEzsignfoldersignerassociationID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationGetObjectV1(context.Background(), pkiEzsignfoldersignerassociationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationGetObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfoldersignerassociationGetObjectV1`: EzsignfoldersignerassociationGetObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldersignerassociationApi.EzsignfoldersignerassociationGetObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfoldersignerassociationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldersignerassociationGetObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfoldersignerassociationGetObjectV1Response**](EzsignfoldersignerassociationGetObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

