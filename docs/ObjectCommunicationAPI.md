# eZmaxAPI\ObjectCommunicationAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationGetObjectV2**](ObjectCommunicationAPI.md#CommunicationGetObjectV2) | **Get** /2/object/communication/{pkiCommunicationID} | Retrieve an existing Communication



## CommunicationGetObjectV2

> CommunicationGetObjectV2Response CommunicationGetObjectV2(ctx, pkiCommunicationID).Execute()

Retrieve an existing Communication



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
    pkiCommunicationID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectCommunicationAPI.CommunicationGetObjectV2(context.Background(), pkiCommunicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectCommunicationAPI.CommunicationGetObjectV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationGetObjectV2`: CommunicationGetObjectV2Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectCommunicationAPI.CommunicationGetObjectV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiCommunicationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationGetObjectV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommunicationGetObjectV2Response**](CommunicationGetObjectV2Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

