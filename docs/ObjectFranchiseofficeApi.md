# eZmaxAPI\ObjectFranchiseofficeApi

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FranchiseofficeGetAutocompleteV1**](ObjectFranchiseofficeApi.md#FranchiseofficeGetAutocompleteV1) | **Get** /1/object/franchiseoffice/getAutocomplete/{sSelector} | Retrieve Franchiseoffices and IDs



## FranchiseofficeGetAutocompleteV1

> CommonGetAutocompleteV1Response FranchiseofficeGetAutocompleteV1(ctx, sSelector).SQuery(sQuery).Execute()

Retrieve Franchiseoffices and IDs



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
    sSelector := "sSelector_example" // string | The type of Franchiseoffices to return
    sQuery := "sQuery_example" // string | Allow to filter on the option value (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectFranchiseofficeApi.FranchiseofficeGetAutocompleteV1(context.Background(), sSelector).SQuery(sQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectFranchiseofficeApi.FranchiseofficeGetAutocompleteV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FranchiseofficeGetAutocompleteV1`: CommonGetAutocompleteV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectFranchiseofficeApi.FranchiseofficeGetAutocompleteV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sSelector** | **string** | The type of Franchiseoffices to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiFranchiseofficeGetAutocompleteV1Request struct via the builder pattern


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

