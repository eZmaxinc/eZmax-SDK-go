# eZmaxAPI\ObjectEzsignfolderApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignfolderCreateObjectV1**](ObjectEzsignfolderApi.md#EzsignfolderCreateObjectV1) | **Post** /1/object/ezsignfolder | Create a new Ezsignfolder
[**EzsignfolderDeleteObjectV1**](ObjectEzsignfolderApi.md#EzsignfolderDeleteObjectV1) | **Delete** /1/object/ezsignfolder/{pkiEzsignfolderID} | Delete an existing Ezsignfolder
[**EzsignfolderEditObjectV1**](ObjectEzsignfolderApi.md#EzsignfolderEditObjectV1) | **Put** /1/object/ezsignfolder/{pkiEzsignfolderID} | Edit an Ezsignfolder
[**EzsignfolderGetEzsigndocumentsV1**](ObjectEzsignfolderApi.md#EzsignfolderGetEzsigndocumentsV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getEzsigndocuments | Retrieve an existing Ezsignfolder&#39;s Ezsigndocuments
[**EzsignfolderGetEzsignfoldersignerassociationsV1**](ObjectEzsignfolderApi.md#EzsignfolderGetEzsignfoldersignerassociationsV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getEzsignfoldersignerassociations | Retrieve an existing Ezsignfolder&#39;s Ezsignfoldersignerassociations
[**EzsignfolderGetFormsDataV1**](ObjectEzsignfolderApi.md#EzsignfolderGetFormsDataV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID}/getFormsData | Retrieve an existing Ezsignfolder&#39;s forms data
[**EzsignfolderGetListV1**](ObjectEzsignfolderApi.md#EzsignfolderGetListV1) | **Get** /1/object/ezsignfolder/getList | Retrieve Ezsignfolder list
[**EzsignfolderGetObjectV1**](ObjectEzsignfolderApi.md#EzsignfolderGetObjectV1) | **Get** /1/object/ezsignfolder/{pkiEzsignfolderID} | Retrieve an existing Ezsignfolder
[**EzsignfolderSendV1**](ObjectEzsignfolderApi.md#EzsignfolderSendV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/send | Send the Ezsignfolder to the signatories for signature
[**EzsignfolderUnsendV1**](ObjectEzsignfolderApi.md#EzsignfolderUnsendV1) | **Post** /1/object/ezsignfolder/{pkiEzsignfolderID}/unsend | Unsend the Ezsignfolder



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderApi.EzsignfolderCreateObjectV1(context.Background()).EzsignfolderCreateObjectV1Request(ezsignfolderCreateObjectV1Request).Execute()
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
 **ezsignfolderCreateObjectV1Request** | [**[]EzsignfolderCreateObjectV1Request**](EzsignfolderCreateObjectV1Request.md) |  | 

### Return type

[**EzsignfolderCreateObjectV1Response**](EzsignfolderCreateObjectV1Response.md)

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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderApi.EzsignfolderDeleteObjectV1(context.Background(), pkiEzsignfolderID).Execute()
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
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderDeleteObjectV1Response**](EzsignfolderDeleteObjectV1Response.md)

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

Edit an Ezsignfolder

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
    pkiEzsignfolderID := int32(56) // int32 | 
    ezsignfolderEditObjectV1Request := *openapiclient.NewEzsignfolderEditObjectV1Request(*openapiclient.NewEzsignfolderRequestCompound(int32(5), int32(92), "Test eZsign Folder", "An extra notes we can add to the ezsign folder", openapiclient.Field-eEzsignfolderSendreminderfrequency("None"))) // EzsignfolderEditObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderApi.EzsignfolderEditObjectV1(context.Background(), pkiEzsignfolderID).EzsignfolderEditObjectV1Request(ezsignfolderEditObjectV1Request).Execute()
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
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderEditObjectV1Request** | [**EzsignfolderEditObjectV1Request**](EzsignfolderEditObjectV1Request.md) |  | 

### Return type

[**EzsignfolderEditObjectV1Response**](EzsignfolderEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetEzsigndocumentsV1

> EzsignfolderGetEzsigndocumentsV1Response EzsignfolderGetEzsigndocumentsV1(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder's Ezsigndocuments

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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderApi.EzsignfolderGetEzsigndocumentsV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderGetEzsigndocumentsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetEzsigndocumentsV1`: EzsignfolderGetEzsigndocumentsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderApi.EzsignfolderGetEzsigndocumentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetEzsigndocumentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetEzsigndocumentsV1Response**](EzsignfolderGetEzsigndocumentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetEzsignfoldersignerassociationsV1

> EzsignfolderGetEzsignfoldersignerassociationsV1Response EzsignfolderGetEzsignfoldersignerassociationsV1(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder's Ezsignfoldersignerassociations

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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderApi.EzsignfolderGetEzsignfoldersignerassociationsV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderGetEzsignfoldersignerassociationsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetEzsignfoldersignerassociationsV1`: EzsignfolderGetEzsignfoldersignerassociationsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderApi.EzsignfolderGetEzsignfoldersignerassociationsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetEzsignfoldersignerassociationsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetEzsignfoldersignerassociationsV1Response**](EzsignfolderGetEzsignfoldersignerassociationsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetFormsDataV1

> EzsignfolderGetFormsDataV1Response EzsignfolderGetFormsDataV1(ctx, pkiEzsignfolderID).Execute()

Retrieve an existing Ezsignfolder's forms data

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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderApi.EzsignfolderGetFormsDataV1(context.Background(), pkiEzsignfolderID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderGetFormsDataV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetFormsDataV1`: EzsignfolderGetFormsDataV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderApi.EzsignfolderGetFormsDataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetFormsDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetFormsDataV1Response**](EzsignfolderGetFormsDataV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderGetListV1

> EzsignfolderGetListV1Response EzsignfolderGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsignfolder list



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
    eOrderBy := "eOrderBy_example" // string | Specify how you want the results to be sorted (optional)
    iRowMax := int32(56) // int32 |  (optional)
    iRowOffset := int32(56) // int32 |  (optional)
    acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)
    sFilter := "sFilter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderApi.EzsignfolderGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderGetListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderGetListV1`: EzsignfolderGetListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderApi.EzsignfolderGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsignfolderGetListV1Response**](EzsignfolderGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

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
    pkiEzsignfolderID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderApi.EzsignfolderGetObjectV1(context.Background(), pkiEzsignfolderID).Execute()
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
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderGetObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignfolderGetObjectV1Response**](EzsignfolderGetObjectV1Response.md)

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
    pkiEzsignfolderID := int32(56) // int32 | 
    ezsignfolderSendV1Request := *openapiclient.NewEzsignfolderSendV1Request("TExtraMessage_example") // EzsignfolderSendV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderApi.EzsignfolderSendV1(context.Background(), pkiEzsignfolderID).EzsignfolderSendV1Request(ezsignfolderSendV1Request).Execute()
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
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderSendV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsignfolderSendV1Request** | [**EzsignfolderSendV1Request**](EzsignfolderSendV1Request.md) |  | 

### Return type

[**EzsignfolderSendV1Response**](EzsignfolderSendV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfolderUnsendV1

> EzsignfolderUnsendV1Response EzsignfolderUnsendV1(ctx, pkiEzsignfolderID).Body(body).Execute()

Unsend the Ezsignfolder



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
    pkiEzsignfolderID := int32(56) // int32 | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignfolderApi.EzsignfolderUnsendV1(context.Background(), pkiEzsignfolderID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfolderApi.EzsignfolderUnsendV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfolderUnsendV1`: EzsignfolderUnsendV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfolderApi.EzsignfolderUnsendV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignfolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfolderUnsendV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**EzsignfolderUnsendV1Response**](EzsignfolderUnsendV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

