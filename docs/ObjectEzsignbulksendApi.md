# eZmaxAPI\ObjectEzsignbulksendApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignbulksendGetListV1**](ObjectEzsignbulksendApi.md#EzsignbulksendGetListV1) | **Get** /1/object/ezsignbulksend/getList | Retrieve Ezsignbulksend list
[**EzsignbulksendGetObjectV1**](ObjectEzsignbulksendApi.md#EzsignbulksendGetObjectV1) | **Get** /1/object/ezsignbulksend/{pkiEzsignbulksendID} | Retrieve an existing Ezsignbulksend



## EzsignbulksendGetListV1

> EzsignbulksendGetListV1Response EzsignbulksendGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsignbulksend list



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
    resp, r, err := apiClient.ObjectEzsignbulksendApi.EzsignbulksendGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendApi.EzsignbulksendGetListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignbulksendGetListV1`: EzsignbulksendGetListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendApi.EzsignbulksendGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsignbulksendGetListV1Response**](EzsignbulksendGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignbulksendGetObjectV1

> EzsignbulksendGetObjectV1Response EzsignbulksendGetObjectV1(ctx, pkiEzsignbulksendID).Execute()

Retrieve an existing Ezsignbulksend



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
    pkiEzsignbulksendID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectEzsignbulksendApi.EzsignbulksendGetObjectV1(context.Background(), pkiEzsignbulksendID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignbulksendApi.EzsignbulksendGetObjectV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignbulksendGetObjectV1`: EzsignbulksendGetObjectV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignbulksendApi.EzsignbulksendGetObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiEzsignbulksendID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignbulksendGetObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EzsignbulksendGetObjectV1Response**](EzsignbulksendGetObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

