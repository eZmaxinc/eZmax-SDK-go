# eZmaxAPI\ObjectEzsigndocumentApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigndocumentApplyEzsigntemplateV1**](ObjectEzsigndocumentApi.md#EzsigndocumentApplyEzsigntemplateV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/applyezsigntemplate | Apply an Ezsign Template to the Ezsigndocument.
[**EzsigndocumentCreateObjectV1**](ObjectEzsigndocumentApi.md#EzsigndocumentCreateObjectV1) | **Post** /1/object/ezsigndocument | Create a new Ezsigndocument
[**EzsigndocumentDeleteObjectV1**](ObjectEzsigndocumentApi.md#EzsigndocumentDeleteObjectV1) | **Delete** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Delete an existing Ezsigndocument
[**EzsigndocumentGetChildrenV1**](ObjectEzsigndocumentApi.md#EzsigndocumentGetChildrenV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getChildren | Retrieve an existing Ezsigndocument&#39;s children IDs
[**EzsigndocumentGetDownloadUrlV1**](ObjectEzsigndocumentApi.md#EzsigndocumentGetDownloadUrlV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getDownloadUrl/{eDocumentType} | Retrieve a URL to download documents.
[**EzsigndocumentGetObjectV1**](ObjectEzsigndocumentApi.md#EzsigndocumentGetObjectV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Retrieve an existing Ezsigndocument



## EzsigndocumentApplyEzsigntemplateV1

> EzsigndocumentApplyEzsigntemplateV1Response EzsigndocumentApplyEzsigntemplateV1(ctx, pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateV1Request(ezsigndocumentApplyEzsigntemplateV1Request).Execute()

Apply an Ezsign Template to the Ezsigndocument.



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
    pkiEzsigndocumentID := int32(56) // int32 | The unique ID of the Ezsigndocument
    ezsigndocumentApplyEzsigntemplateV1Request := *openapiclient.NewEzsigndocumentApplyEzsigntemplateV1Request(int32(36), []string{"John"}, []int32{int32(20)}) // EzsigndocumentApplyEzsigntemplateV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsigndocumentApi.EzsigndocumentApplyEzsigntemplateV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateV1Request(ezsigndocumentApplyEzsigntemplateV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentApplyEzsigntemplateV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentApplyEzsigntemplateV1`: EzsigndocumentApplyEzsigntemplateV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentApplyEzsigntemplateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentApplyEzsigntemplateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentApplyEzsigntemplateV1Request** | [**EzsigndocumentApplyEzsigntemplateV1Request**](EzsigndocumentApplyEzsigntemplateV1Request.md) |  | 

### Return type

[**EzsigndocumentApplyEzsigntemplateV1Response**](ezsigndocument-applyEzsigntemplate-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentCreateObjectV1

> EzsigndocumentCreateObjectV1Response EzsigndocumentCreateObjectV1(ctx).EzsigndocumentCreateObjectV1Request(ezsigndocumentCreateObjectV1Request).Execute()

Create a new Ezsigndocument



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
    ezsigndocumentCreateObjectV1Request := []openapiclient.EzsigndocumentCreateObjectV1Request{*openapiclient.NewEzsigndocumentCreateObjectV1Request()} // []EzsigndocumentCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsigndocumentApi.EzsigndocumentCreateObjectV1(context.Background()).EzsigndocumentCreateObjectV1Request(ezsigndocumentCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentCreateObjectV1`: EzsigndocumentCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigndocumentCreateObjectV1Request** | [**[]EzsigndocumentCreateObjectV1Request**](ezsigndocument-createObject-v1-Request.md) |  | 

### Return type

[**EzsigndocumentCreateObjectV1Response**](ezsigndocument-createObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentDeleteObjectV1

> EzsigndocumentDeleteObjectV1Response EzsigndocumentDeleteObjectV1(ctx, pkiEzsigndocumentID).Execute()

Delete an existing Ezsigndocument

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
    pkiEzsigndocumentID := int32(56) // int32 | The unique ID of the Ezsigndocument

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsigndocumentApi.EzsigndocumentDeleteObjectV1(context.Background(), pkiEzsigndocumentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentDeleteObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentDeleteObjectV1`: EzsigndocumentDeleteObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentDeleteObjectV1Response**](ezsigndocument-deleteObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetChildrenV1

> EzsigndocumentGetChildrenV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument's children IDs

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
    pkiEzsigndocumentID := int32(56) // int32 | The unique ID of the Ezsigndocument

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsigndocumentApi.EzsigndocumentGetChildrenV1(context.Background(), pkiEzsigndocumentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentGetChildrenV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetChildrenV1Request struct via the builder pattern


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


## EzsigndocumentGetDownloadUrlV1

> EzsigndocumentGetDownloadUrlV1Response EzsigndocumentGetDownloadUrlV1(ctx, pkiEzsigndocumentID, eDocumentType).Execute()

Retrieve a URL to download documents.



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
    pkiEzsigndocumentID := int32(56) // int32 | The unique ID of the Ezsigndocument
    eDocumentType := "eDocumentType_example" // string | The type of document to retrieve.  1. **Initial** Is the initial document before any signature were applied. 2. **Signed** Is the final document once all signatures were applied. 3. **Proofdocument** Is the evidence report. 4. **Proof** Is the complete evidence archive including all of the above and more. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsigndocumentApi.EzsigndocumentGetDownloadUrlV1(context.Background(), pkiEzsigndocumentID, eDocumentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentGetDownloadUrlV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentGetDownloadUrlV1`: EzsigndocumentGetDownloadUrlV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentGetDownloadUrlV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**eDocumentType** | **string** | The type of document to retrieve.  1. **Initial** Is the initial document before any signature were applied. 2. **Signed** Is the final document once all signatures were applied. 3. **Proofdocument** Is the evidence report. 4. **Proof** Is the complete evidence archive including all of the above and more.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetDownloadUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EzsigndocumentGetDownloadUrlV1Response**](ezsigndocument-getDownloadUrl-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetObjectV1

> EzsigndocumentGetObjectV1Response EzsigndocumentGetObjectV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument

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
    pkiEzsigndocumentID := int32(56) // int32 | The unique ID of the Ezsigndocument

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsigndocumentApi.EzsigndocumentGetObjectV1(context.Background(), pkiEzsigndocumentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentGetObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentGetObjectV1`: EzsigndocumentGetObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentGetObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetObjectV1Response**](ezsigndocument-getObject-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

