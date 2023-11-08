# eZmaxAPI\ObjectInvoiceAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceGetCommunicationListV1**](ObjectInvoiceAPI.md#InvoiceGetCommunicationListV1) | **Get** /1/object/invoice/{pkiInvoiceID}/getCommunicationList | Retrieve Communication list



## InvoiceGetCommunicationListV1

> InvoiceGetCommunicationListV1Response InvoiceGetCommunicationListV1(ctx, pkiInvoiceID).Execute()

Retrieve Communication list



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
    pkiInvoiceID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectInvoiceAPI.InvoiceGetCommunicationListV1(context.Background(), pkiInvoiceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectInvoiceAPI.InvoiceGetCommunicationListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvoiceGetCommunicationListV1`: InvoiceGetCommunicationListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectInvoiceAPI.InvoiceGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiInvoiceID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceGetCommunicationListV1Response**](InvoiceGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

