# eZmaxAPI\ObjectEzsignfoldertypeApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EzsignfoldertypeGetAutocompleteV1**](ObjectEzsignfoldertypeApi.md#EzsignfoldertypeGetAutocompleteV1) | **Get** /1/object/ezsignfoldertype/getAutocomplete/{sSelector}/ | Retrieve Ezsignfoldertypes and IDs
[**EzsignfoldertypeGetListV1**](ObjectEzsignfoldertypeApi.md#EzsignfoldertypeGetListV1) | **Get** /1/object/ezsignfoldertype/getList | Retrieve Ezsignfoldertype list



## EzsignfoldertypeGetAutocompleteV1

> CommonGetAutocompleteV1Response EzsignfoldertypeGetAutocompleteV1(ctx, sSelector).AcceptLanguage(acceptLanguage).SQuery(sQuery).Execute()

Retrieve Ezsignfoldertypes and IDs



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
    sSelector := "sSelector_example" // string | The type of Ezsignfoldertypes to return
    acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)
    sQuery := "sQuery_example" // string | Allow to filter the returned results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfoldertypeApi.EzsignfoldertypeGetAutocompleteV1(context.Background(), sSelector).AcceptLanguage(acceptLanguage).SQuery(sQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldertypeApi.EzsignfoldertypeGetAutocompleteV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfoldertypeGetAutocompleteV1`: CommonGetAutocompleteV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldertypeApi.EzsignfoldertypeGetAutocompleteV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Ezsignfoldertypes to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldertypeGetAutocompleteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sQuery** | **string** | Allow to filter the returned results | 

### Return type

[**CommonGetAutocompleteV1Response**](CommonGetAutocompleteV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EzsignfoldertypeGetListV1

> EzsignfoldertypeGetListV1Response EzsignfoldertypeGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Ezsignfoldertype list



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectEzsignfoldertypeApi.EzsignfoldertypeGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectEzsignfoldertypeApi.EzsignfoldertypeGetListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EzsignfoldertypeGetListV1`: EzsignfoldertypeGetListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectEzsignfoldertypeApi.EzsignfoldertypeGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEzsignfoldertypeGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | 
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**EzsignfoldertypeGetListV1Response**](EzsignfoldertypeGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

