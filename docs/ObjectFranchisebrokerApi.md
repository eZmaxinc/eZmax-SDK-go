# eZmaxAPI\ObjectFranchisebrokerApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FranchisebrokerGetAutocompleteV1**](ObjectFranchisebrokerApi.md#FranchisebrokerGetAutocompleteV1) | **Get** /1/object/franchisebroker/getAutocomplete/{sSelector} | Retrieve Franchisebrokers and IDs



## FranchisebrokerGetAutocompleteV1

> CommonGetAutocompleteV1Response FranchisebrokerGetAutocompleteV1(ctx, sSelector).SQuery(sQuery).Execute()

Retrieve Franchisebrokers and IDs



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
    sSelector := "sSelector_example" // string | The type of Franchisebrokers to return
    sQuery := "sQuery_example" // string | Allow to filter on the option value (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectFranchisebrokerApi.FranchisebrokerGetAutocompleteV1(context.Background(), sSelector).SQuery(sQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectFranchisebrokerApi.FranchisebrokerGetAutocompleteV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FranchisebrokerGetAutocompleteV1`: CommonGetAutocompleteV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectFranchisebrokerApi.FranchisebrokerGetAutocompleteV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Franchisebrokers to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiFranchisebrokerGetAutocompleteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sQuery** | **string** | Allow to filter on the option value | 

### Return type

[**CommonGetAutocompleteV1Response**](Common-getAutocomplete-v1-Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

