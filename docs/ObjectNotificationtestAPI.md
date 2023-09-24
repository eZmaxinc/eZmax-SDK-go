# eZmaxAPI\ObjectNotificationtestAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationtestGetElementsV1**](ObjectNotificationtestAPI.md#NotificationtestGetElementsV1) | **Get** /1/object/notificationtest/{pkiNotificationtestID}/getElements | Retrieve an existing Notificationtest&#39;s Elements



## NotificationtestGetElementsV1

> NotificationtestGetElementsV1Response NotificationtestGetElementsV1(ctx, pkiNotificationtestID).Execute()

Retrieve an existing Notificationtest's Elements



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
    pkiNotificationtestID := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectNotificationtestAPI.NotificationtestGetElementsV1(context.Background(), pkiNotificationtestID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotificationtestAPI.NotificationtestGetElementsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationtestGetElementsV1`: NotificationtestGetElementsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectNotificationtestAPI.NotificationtestGetElementsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotificationtestID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationtestGetElementsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationtestGetElementsV1Response**](NotificationtestGetElementsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

