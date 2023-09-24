# eZmaxAPI\ObjectEzsigntemplatedocumentAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsigntemplatedocumentCreateObjectV1**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentCreateObjectV1) | **Post** /1/object/ezsigntemplatedocument | Create a new Ezsigntemplatedocument
[**EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1) | **Put** /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/editEzsigntemplateformfieldgroups | Edit multiple Ezsigntemplateformfieldgroups
[**EzsigntemplatedocumentEditEzsigntemplatesignaturesV1**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentEditEzsigntemplatesignaturesV1) | **Put** /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/editEzsigntemplatesignatures | Edit multiple Ezsigntemplatesignatures
[**EzsigntemplatedocumentEditObjectV1**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentEditObjectV1) | **Put** /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID} | Edit an existing Ezsigntemplatedocument
[**EzsigntemplatedocumentFlattenV1**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentFlattenV1) | **Post** /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/flatten | Flatten
[**EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1) | **Get** /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/getEzsigntemplatedocumentpages | Retrieve an existing Ezsigntemplatedocument&#39;s Ezsigntemplatedocumentpages
[**EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1) | **Get** /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/getEzsigntemplateformfieldgroups | Retrieve an existing Ezsigntemplatedocument&#39;s Ezsigntemplateformfieldgroups
[**EzsigntemplatedocumentGetEzsigntemplatesignaturesV1**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentGetEzsigntemplatesignaturesV1) | **Get** /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/getEzsigntemplatesignatures | Retrieve an existing Ezsigntemplatedocument&#39;s Ezsigntemplatesignatures
[**EzsigntemplatedocumentGetObjectV2**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentGetObjectV2) | **Get** /2/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID} | Retrieve an existing Ezsigntemplatedocument
[**EzsigntemplatedocumentGetWordsPositionsV1**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentGetWordsPositionsV1) | **Post** /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/getWordsPositions | Retrieve positions X,Y of given words from a Ezsigntemplatedocument
[**EzsigntemplatedocumentPatchObjectV1**](ObjectEzsigntemplatedocumentAPI.md#EzsigntemplatedocumentPatchObjectV1) | **Patch** /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID} | Patch an existing Ezsigntemplatedocument



## EzsigntemplatedocumentCreateObjectV1

> EzsigntemplatedocumentCreateObjectV1Response EzsigntemplatedocumentCreateObjectV1(ctx).EzsigntemplatedocumentCreateObjectV1Request(ezsigntemplatedocumentCreateObjectV1Request).Execute()

Create a new Ezsigntemplatedocument



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
    ezsigntemplatedocumentCreateObjectV1Request := *openapiclient.NewEzsigntemplatedocumentCreateObjectV1Request([]openapiclient.EzsigntemplatedocumentRequestCompound{*openapiclient.NewEzsigntemplatedocumentRequestCompound(int32(36), "Standard Contract", "EEzsigntemplatedocumentSource_example")}) // EzsigntemplatedocumentCreateObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentCreateObjectV1(context.Background()).EzsigntemplatedocumentCreateObjectV1Request(ezsigntemplatedocumentCreateObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentCreateObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentCreateObjectV1`: EzsigntemplatedocumentCreateObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ezsigntemplatedocumentCreateObjectV1Request** | [**EzsigntemplatedocumentCreateObjectV1Request**](EzsigntemplatedocumentCreateObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatedocumentCreateObjectV1Response**](EzsigntemplatedocumentCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1

> EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1(ctx, pkiEzsigntemplatedocumentID).EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request(ezsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request).Execute()

Edit multiple Ezsigntemplateformfieldgroups



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
    pkiEzsigntemplatedocumentID := int32(56) // int32 | 
    ezsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request := *openapiclient.NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request([]openapiclient.EzsigntemplateformfieldgroupRequestCompound{*openapiclient.NewEzsigntemplateformfieldgroupRequestCompound(int32(133), openapiclient.Field-eEzsigntemplateformfieldgroupType("Text"), openapiclient.Field-eEzsigntemplateformfieldgroupSignerrequirement("All"), "Allergies", int32(1), "Foo", int32(1), int32(2), false, []openapiclient.EzsigntemplateformfieldgroupsignerRequestCompound{*openapiclient.NewEzsigntemplateformfieldgroupsignerRequestCompound(int32(9))}, []openapiclient.EzsigntemplateformfieldRequestCompound{*openapiclient.NewEzsigntemplateformfieldRequestCompound(int32(1), "Peanuts", int32(200), int32(300), int32(102), int32(22))})}) // EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1(context.Background(), pkiEzsigntemplatedocumentID).EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request(ezsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1`: EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request** | [**EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request**](EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Request.md) |  | 

### Return type

[**EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response**](EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentEditEzsigntemplatesignaturesV1

> EzsigntemplatedocumentEditEzsigntemplatesignaturesV1Response EzsigntemplatedocumentEditEzsigntemplatesignaturesV1(ctx, pkiEzsigntemplatedocumentID).EzsigntemplatedocumentEditEzsigntemplatesignaturesV1Request(ezsigntemplatedocumentEditEzsigntemplatesignaturesV1Request).Execute()

Edit multiple Ezsigntemplatesignatures



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
    pkiEzsigntemplatedocumentID := int32(56) // int32 | 
    ezsigntemplatedocumentEditEzsigntemplatesignaturesV1Request := *openapiclient.NewEzsigntemplatedocumentEditEzsigntemplatesignaturesV1Request([]openapiclient.EzsigntemplatesignatureRequestCompound{*openapiclient.NewEzsigntemplatesignatureRequestCompound(int32(133), int32(9), int32(1), int32(200), int32(300), int32(1), openapiclient.Field-eEzsigntemplatesignatureType("Acknowledgement"))}) // EzsigntemplatedocumentEditEzsigntemplatesignaturesV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentEditEzsigntemplatesignaturesV1(context.Background(), pkiEzsigntemplatedocumentID).EzsigntemplatedocumentEditEzsigntemplatesignaturesV1Request(ezsigntemplatedocumentEditEzsigntemplatesignaturesV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentEditEzsigntemplatesignaturesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentEditEzsigntemplatesignaturesV1`: EzsigntemplatedocumentEditEzsigntemplatesignaturesV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentEditEzsigntemplatesignaturesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentEditEzsigntemplatesignaturesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatedocumentEditEzsigntemplatesignaturesV1Request** | [**EzsigntemplatedocumentEditEzsigntemplatesignaturesV1Request**](EzsigntemplatedocumentEditEzsigntemplatesignaturesV1Request.md) |  | 

### Return type

[**EzsigntemplatedocumentEditEzsigntemplatesignaturesV1Response**](EzsigntemplatedocumentEditEzsigntemplatesignaturesV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentEditObjectV1

> EzsigntemplatedocumentEditObjectV1Response EzsigntemplatedocumentEditObjectV1(ctx, pkiEzsigntemplatedocumentID).EzsigntemplatedocumentEditObjectV1Request(ezsigntemplatedocumentEditObjectV1Request).Execute()

Edit an existing Ezsigntemplatedocument



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
    pkiEzsigntemplatedocumentID := int32(56) // int32 | 
    ezsigntemplatedocumentEditObjectV1Request := *openapiclient.NewEzsigntemplatedocumentEditObjectV1Request(*openapiclient.NewEzsigntemplatedocumentRequestCompound(int32(36), "Standard Contract", "EEzsigntemplatedocumentSource_example")) // EzsigntemplatedocumentEditObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentEditObjectV1(context.Background(), pkiEzsigntemplatedocumentID).EzsigntemplatedocumentEditObjectV1Request(ezsigntemplatedocumentEditObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentEditObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentEditObjectV1`: EzsigntemplatedocumentEditObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentEditObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentEditObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatedocumentEditObjectV1Request** | [**EzsigntemplatedocumentEditObjectV1Request**](EzsigntemplatedocumentEditObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatedocumentEditObjectV1Response**](EzsigntemplatedocumentEditObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentFlattenV1

> EzsigntemplatedocumentFlattenV1Response EzsigntemplatedocumentFlattenV1(ctx, pkiEzsigntemplatedocumentID).Body(body).Execute()

Flatten



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
    pkiEzsigntemplatedocumentID := int32(56) // int32 | 
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentFlattenV1(context.Background(), pkiEzsigntemplatedocumentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentFlattenV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentFlattenV1`: EzsigntemplatedocumentFlattenV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentFlattenV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentFlattenV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**EzsigntemplatedocumentFlattenV1Response**](EzsigntemplatedocumentFlattenV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1

> EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1Response EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1(ctx, pkiEzsigntemplatedocumentID).Execute()

Retrieve an existing Ezsigntemplatedocument's Ezsigntemplatedocumentpages



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
    pkiEzsigntemplatedocumentID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1(context.Background(), pkiEzsigntemplatedocumentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1`: EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1Response**](EzsigntemplatedocumentGetEzsigntemplatedocumentpagesV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1

> EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1Response EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1(ctx, pkiEzsigntemplatedocumentID).Execute()

Retrieve an existing Ezsigntemplatedocument's Ezsigntemplateformfieldgroups



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
    pkiEzsigntemplatedocumentID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1(context.Background(), pkiEzsigntemplatedocumentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1`: EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1Response**](EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentGetEzsigntemplatesignaturesV1

> EzsigntemplatedocumentGetEzsigntemplatesignaturesV1Response EzsigntemplatedocumentGetEzsigntemplatesignaturesV1(ctx, pkiEzsigntemplatedocumentID).Execute()

Retrieve an existing Ezsigntemplatedocument's Ezsigntemplatesignatures



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
    pkiEzsigntemplatedocumentID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetEzsigntemplatesignaturesV1(context.Background(), pkiEzsigntemplatedocumentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetEzsigntemplatesignaturesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentGetEzsigntemplatesignaturesV1`: EzsigntemplatedocumentGetEzsigntemplatesignaturesV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetEzsigntemplatesignaturesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentGetEzsigntemplatesignaturesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatedocumentGetEzsigntemplatesignaturesV1Response**](EzsigntemplatedocumentGetEzsigntemplatesignaturesV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentGetObjectV2

> EzsigntemplatedocumentGetObjectV2Response EzsigntemplatedocumentGetObjectV2(ctx, pkiEzsigntemplatedocumentID).Execute()

Retrieve an existing Ezsigntemplatedocument



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
    pkiEzsigntemplatedocumentID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetObjectV2(context.Background(), pkiEzsigntemplatedocumentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentGetObjectV2`: EzsigntemplatedocumentGetObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsigntemplatedocumentGetObjectV2Response**](EzsigntemplatedocumentGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentGetWordsPositionsV1

> EzsigntemplatedocumentGetWordsPositionsV1Response EzsigntemplatedocumentGetWordsPositionsV1(ctx, pkiEzsigntemplatedocumentID).EzsigntemplatedocumentGetWordsPositionsV1Request(ezsigntemplatedocumentGetWordsPositionsV1Request).Execute()

Retrieve positions X,Y of given words from a Ezsigntemplatedocument



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
    pkiEzsigntemplatedocumentID := int32(56) // int32 | 
    ezsigntemplatedocumentGetWordsPositionsV1Request := *openapiclient.NewEzsigntemplatedocumentGetWordsPositionsV1Request("EGet_example", false) // EzsigntemplatedocumentGetWordsPositionsV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetWordsPositionsV1(context.Background(), pkiEzsigntemplatedocumentID).EzsigntemplatedocumentGetWordsPositionsV1Request(ezsigntemplatedocumentGetWordsPositionsV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetWordsPositionsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentGetWordsPositionsV1`: EzsigntemplatedocumentGetWordsPositionsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentGetWordsPositionsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentGetWordsPositionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatedocumentGetWordsPositionsV1Request** | [**EzsigntemplatedocumentGetWordsPositionsV1Request**](EzsigntemplatedocumentGetWordsPositionsV1Request.md) |  | 

### Return type

[**EzsigntemplatedocumentGetWordsPositionsV1Response**](EzsigntemplatedocumentGetWordsPositionsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsigntemplatedocumentPatchObjectV1

> EzsigntemplatedocumentPatchObjectV1Response EzsigntemplatedocumentPatchObjectV1(ctx, pkiEzsigntemplatedocumentID).EzsigntemplatedocumentPatchObjectV1Request(ezsigntemplatedocumentPatchObjectV1Request).Execute()

Patch an existing Ezsigntemplatedocument



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
    pkiEzsigntemplatedocumentID := int32(56) // int32 | 
    ezsigntemplatedocumentPatchObjectV1Request := *openapiclient.NewEzsigntemplatedocumentPatchObjectV1Request(*openapiclient.NewEzsigntemplatedocumentRequestPatch()) // EzsigntemplatedocumentPatchObjectV1Request | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentPatchObjectV1(context.Background(), pkiEzsigntemplatedocumentID).EzsigntemplatedocumentPatchObjectV1Request(ezsigntemplatedocumentPatchObjectV1Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentPatchObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsigntemplatedocumentPatchObjectV1`: EzsigntemplatedocumentPatchObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsigntemplatedocumentAPI.EzsigntemplatedocumentPatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsigntemplatedocumentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsigntemplatedocumentPatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ezsigntemplatedocumentPatchObjectV1Request** | [**EzsigntemplatedocumentPatchObjectV1Request**](EzsigntemplatedocumentPatchObjectV1Request.md) |  | 

### Return type

[**EzsigntemplatedocumentPatchObjectV1Response**](EzsigntemplatedocumentPatchObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

