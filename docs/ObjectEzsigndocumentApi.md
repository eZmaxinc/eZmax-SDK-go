# eZmaxAPI\ObjectEzsigndocumentApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigndocumentApplyEzsigntemplateV1**](ObjectEzsigndocumentApi.md#EzsigndocumentApplyEzsigntemplateV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/applyezsigntemplate | Apply an Ezsign Template to the Ezsigndocument.
[**EzsigndocumentApplyEzsigntemplateV2**](ObjectEzsigndocumentApi.md#EzsigndocumentApplyEzsigntemplateV2) | **Post** /2/object/ezsigndocument/{pkiEzsigndocumentID}/applyEzsigntemplate | Apply an Ezsign Template to the Ezsigndocument.
[**EzsigndocumentCreateObjectV1**](ObjectEzsigndocumentApi.md#EzsigndocumentCreateObjectV1) | **Post** /1/object/ezsigndocument | Create a new Ezsigndocument
[**EzsigndocumentDeleteObjectV1**](ObjectEzsigndocumentApi.md#EzsigndocumentDeleteObjectV1) | **Delete** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Delete an existing Ezsigndocument
[**EzsigndocumentEditEzsignsignaturesV1**](ObjectEzsigndocumentApi.md#EzsigndocumentEditEzsignsignaturesV1) | **Put** /1/object/ezsigndocument/{pkiEzsigndocumentID}/editEzsignsignatures | Edit multiple ezsignsignatures
[**EzsigndocumentGetDownloadUrlV1**](ObjectEzsigndocumentApi.md#EzsigndocumentGetDownloadUrlV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getDownloadUrl/{eDocumentType} | Retrieve a URL to download documents.
[**EzsigndocumentGetEzsignpagesV1**](ObjectEzsigndocumentApi.md#EzsigndocumentGetEzsignpagesV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getEzsignpages | Retrieve an existing Ezsigndocument&#39;s Ezsignpages
[**EzsigndocumentGetFormDataV1**](ObjectEzsigndocumentApi.md#EzsigndocumentGetFormDataV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getFormData | Retrieve an existing Ezsigndocument&#39;s Form Data
[**EzsigndocumentGetObjectV1**](ObjectEzsigndocumentApi.md#EzsigndocumentGetObjectV1) | **Get** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Retrieve an existing Ezsigndocument
[**EzsigndocumentGetWordsPositionsV1**](ObjectEzsigndocumentApi.md#EzsigndocumentGetWordsPositionsV1) | **Post** /1/object/ezsigndocument/{pkiEzsigndocumentID}/getWordsPositions | Retrieve positions X,Y of given words from a Ezsigndocument
[**EzsigndocumentPatchObjectV1**](ObjectEzsigndocumentApi.md#EzsigndocumentPatchObjectV1) | **Patch** /1/object/ezsigndocument/{pkiEzsigndocumentID} | Patch an existing Ezsigndocument



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
    pkiEzsigndocumentID := int32(56) // int32 | 
    ezsigndocumentApplyEzsigntemplateV1Request := *openapiclient.NewEzsigndocumentApplyEzsigntemplateV1Request(int32(36), []string{"John"}, []int32{int32(20)}) // EzsigndocumentApplyEzsigntemplateV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentApplyEzsigntemplateV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateV1Request(ezsigndocumentApplyEzsigntemplateV1Request).Execute()
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
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentApplyEzsigntemplateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentApplyEzsigntemplateV1Request** | [**EzsigndocumentApplyEzsigntemplateV1Request**](EzsigndocumentApplyEzsigntemplateV1Request.md) |  | 

### Return type

[**EzsigndocumentApplyEzsigntemplateV1Response**](EzsigndocumentApplyEzsigntemplateV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentApplyEzsigntemplateV2

> EzsigndocumentApplyEzsigntemplateV2Response EzsigndocumentApplyEzsigntemplateV2(ctx, pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateV2Request(ezsigndocumentApplyEzsigntemplateV2Request).Execute()

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
    pkiEzsigndocumentID := int32(56) // int32 | 
    ezsigndocumentApplyEzsigntemplateV2Request := *openapiclient.NewEzsigndocumentApplyEzsigntemplateV2Request(int32(36), []string{"John"}, []int32{int32(20)}) // EzsigndocumentApplyEzsigntemplateV2Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentApplyEzsigntemplateV2(context.Background(), pkiEzsigndocumentID).EzsigndocumentApplyEzsigntemplateV2Request(ezsigndocumentApplyEzsigntemplateV2Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentApplyEzsigntemplateV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentApplyEzsigntemplateV2`: EzsigndocumentApplyEzsigntemplateV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentApplyEzsigntemplateV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentApplyEzsigntemplateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentApplyEzsigntemplateV2Request** | [**EzsigndocumentApplyEzsigntemplateV2Request**](EzsigndocumentApplyEzsigntemplateV2Request.md) |  | 

### Return type

[**EzsigndocumentApplyEzsigntemplateV2Response**](EzsigndocumentApplyEzsigntemplateV2Response.md)

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentCreateObjectV1(context.Background()).EzsigndocumentCreateObjectV1Request(ezsigndocumentCreateObjectV1Request).Execute()
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
 **ezsigndocumentCreateObjectV1Request** | [**[]EzsigndocumentCreateObjectV1Request**](EzsigndocumentCreateObjectV1Request.md) |  | 

### Return type

[**EzsigndocumentCreateObjectV1Response**](EzsigndocumentCreateObjectV1Response.md)

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
    pkiEzsigndocumentID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentDeleteObjectV1(context.Background(), pkiEzsigndocumentID).Execute()
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
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentDeleteObjectV1Response**](EzsigndocumentDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentEditEzsignsignaturesV1

> EzsigndocumentEditEzsignsignaturesV1Response EzsigndocumentEditEzsignsignaturesV1(ctx, pkiEzsigndocumentID).EzsignsignatureRequestCompound(ezsignsignatureRequestCompound).Execute()

Edit multiple ezsignsignatures



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
    pkiEzsigndocumentID := int32(56) // int32 | 
    ezsignsignatureRequestCompound := []openapiclient.EzsignsignatureRequestCompound{*openapiclient.NewEzsignsignatureRequestCompound(int32(20), int32(1), int32(200), int32(300), int32(1), openapiclient.Field-eEzsignsignatureType("Acknowledgement"), int32(97))} // []EzsignsignatureRequestCompound | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentEditEzsignsignaturesV1(context.Background(), pkiEzsigndocumentID).EzsignsignatureRequestCompound(ezsignsignatureRequestCompound).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentEditEzsignsignaturesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentEditEzsignsignaturesV1`: EzsigndocumentEditEzsignsignaturesV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentEditEzsignsignaturesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentEditEzsignsignaturesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignsignatureRequestCompound** | [**[]EzsignsignatureRequestCompound**](EzsignsignatureRequestCompound.md) |  | 

### Return type

[**EzsigndocumentEditEzsignsignaturesV1Response**](EzsigndocumentEditEzsignsignaturesV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
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
    pkiEzsigndocumentID := int32(56) // int32 | 
    eDocumentType := "eDocumentType_example" // string | The type of document to retrieve.  1. **Initial** Is the initial document before any signature were applied. 2. **Signed** Is the final document once all signatures were applied. 3. **Proofdocument** Is the evidence report. 4. **Proof** Is the complete evidence archive including all of the above and more. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentGetDownloadUrlV1(context.Background(), pkiEzsigndocumentID, eDocumentType).Execute()
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
**pkiEzsigndocumentID** | **int32** |  | 
**eDocumentType** | **string** | The type of document to retrieve.  1. **Initial** Is the initial document before any signature were applied. 2. **Signed** Is the final document once all signatures were applied. 3. **Proofdocument** Is the evidence report. 4. **Proof** Is the complete evidence archive including all of the above and more.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetDownloadUrlV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EzsigndocumentGetDownloadUrlV1Response**](EzsigndocumentGetDownloadUrlV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetEzsignpagesV1

> EzsigndocumentGetEzsignpagesV1Response EzsigndocumentGetEzsignpagesV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument's Ezsignpages

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
    pkiEzsigndocumentID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentGetEzsignpagesV1(context.Background(), pkiEzsigndocumentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentGetEzsignpagesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentGetEzsignpagesV1`: EzsigndocumentGetEzsignpagesV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentGetEzsignpagesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetEzsignpagesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetEzsignpagesV1Response**](EzsigndocumentGetEzsignpagesV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetFormDataV1

> EzsigndocumentGetFormDataV1Response EzsigndocumentGetFormDataV1(ctx, pkiEzsigndocumentID).Execute()

Retrieve an existing Ezsigndocument's Form Data

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
    pkiEzsigndocumentID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentGetFormDataV1(context.Background(), pkiEzsigndocumentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentGetFormDataV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentGetFormDataV1`: EzsigndocumentGetFormDataV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentGetFormDataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetFormDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetFormDataV1Response**](EzsigndocumentGetFormDataV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/zip, text/csv

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
    pkiEzsigndocumentID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentGetObjectV1(context.Background(), pkiEzsigndocumentID).Execute()
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
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigndocumentGetObjectV1Response**](EzsigndocumentGetObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentGetWordsPositionsV1

> EzsigndocumentGetWordsPositionsV1Response EzsigndocumentGetWordsPositionsV1(ctx, pkiEzsigndocumentID).EzsigndocumentGetWordsPositionsV1Request(ezsigndocumentGetWordsPositionsV1Request).Execute()

Retrieve positions X,Y of given words from a Ezsigndocument

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
    pkiEzsigndocumentID := int32(56) // int32 | 
    ezsigndocumentGetWordsPositionsV1Request := *openapiclient.NewEzsigndocumentGetWordsPositionsV1Request("EGet_example", false) // EzsigndocumentGetWordsPositionsV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentGetWordsPositionsV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentGetWordsPositionsV1Request(ezsigndocumentGetWordsPositionsV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentGetWordsPositionsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentGetWordsPositionsV1`: EzsigndocumentGetWordsPositionsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentGetWordsPositionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentGetWordsPositionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentGetWordsPositionsV1Request** | [**EzsigndocumentGetWordsPositionsV1Request**](EzsigndocumentGetWordsPositionsV1Request.md) |  | 

### Return type

[**EzsigndocumentGetWordsPositionsV1Response**](EzsigndocumentGetWordsPositionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigndocumentPatchObjectV1

> EzsigndocumentPatchObjectV1Response EzsigndocumentPatchObjectV1(ctx, pkiEzsigndocumentID).EzsigndocumentPatchObjectV1Request(ezsigndocumentPatchObjectV1Request).Execute()

Patch an existing Ezsigndocument

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
    pkiEzsigndocumentID := int32(56) // int32 | 
    ezsigndocumentPatchObjectV1Request := *openapiclient.NewEzsigndocumentPatchObjectV1Request(*openapiclient.NewEzsigndocumentRequestPatch()) // EzsigndocumentPatchObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigndocumentApi.EzsigndocumentPatchObjectV1(context.Background(), pkiEzsigndocumentID).EzsigndocumentPatchObjectV1Request(ezsigndocumentPatchObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigndocumentApi.EzsigndocumentPatchObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigndocumentPatchObjectV1`: EzsigndocumentPatchObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigndocumentApi.EzsigndocumentPatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigndocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigndocumentPatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigndocumentPatchObjectV1Request** | [**EzsigndocumentPatchObjectV1Request**](EzsigndocumentPatchObjectV1Request.md) |  | 

### Return type

[**EzsigndocumentPatchObjectV1Response**](EzsigndocumentPatchObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

