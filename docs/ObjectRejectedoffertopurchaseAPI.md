# eZmaxAPI\ObjectRejectedoffertopurchaseAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RejectedoffertopurchaseGetCommunicationListV1**](ObjectRejectedoffertopurchaseAPI.md#RejectedoffertopurchaseGetCommunicationListV1) | **Get** /1/object/rejectedoffertopurchase/{pkiRejectedoffertopurchaseID}/getCommunicationList | Retrieve Communication list



## RejectedoffertopurchaseGetCommunicationListV1

> RejectedoffertopurchaseGetCommunicationListV1Response RejectedoffertopurchaseGetCommunicationListV1(ctx, pkiRejectedoffertopurchaseID).Execute()

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
    pkiRejectedoffertopurchaseID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationListV1(context.Background(), pkiRejectedoffertopurchaseID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectedoffertopurchaseGetCommunicationListV1`: RejectedoffertopurchaseGetCommunicationListV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectRejectedoffertopurchaseAPI.RejectedoffertopurchaseGetCommunicationListV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiRejectedoffertopurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectedoffertopurchaseGetCommunicationListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RejectedoffertopurchaseGetCommunicationListV1Response**](RejectedoffertopurchaseGetCommunicationListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

