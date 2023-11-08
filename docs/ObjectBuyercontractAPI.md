# eZmaxAPI\ObjectBuyercontractAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuyercontractGetCommunicationListV1**](ObjectBuyercontractAPI.md#BuyercontractGetCommunicationListV1) | **Get** /1/object/buyercontract/{pkiBuyercontractID}/getCommunicationList | Retrieve Communication list



## BuyercontractGetCommunicationListV1

> BuyercontractGetCommunicationListV1Response BuyercontractGetCommunicationListV1(ctx, pkiBuyercontractID).Execute()

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
    pkiBuyercontractID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectBuyercontractAPI.BuyercontractGetCommunicationListV1(context.Background(), pkiBuyercontractID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectBuyercontractAPI.BuyercontractGetCommunicationListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuyercontractGetCommunicationListV1`: BuyercontractGetCommunicationListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectBuyercontractAPI.BuyercontractGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiBuyercontractID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuyercontractGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuyercontractGetCommunicationListV1Response**](BuyercontractGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

