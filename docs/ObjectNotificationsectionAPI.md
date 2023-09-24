# eZmaxAPI\ObjectNotificationsectionAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationsectionGetNotificationtestsV1**](ObjectNotificationsectionAPI.md#NotificationsectionGetNotificationtestsV1) | **Get** /1/object/notificationsection/{pkiNotificationsectionID}/getNotificationtests | Retrieve an existing Notificationsection&#39;s Notificationtests



## NotificationsectionGetNotificationtestsV1

> NotificationsectionGetNotificationtestsV1Response NotificationsectionGetNotificationtestsV1(ctx, pkiNotificationsectionID).BShowHidden(bShowHidden).Execute()

Retrieve an existing Notificationsection's Notificationtests



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
    pkiNotificationsectionID := int32(56) // int32 | 
    bShowHidden := true // bool | Whether or not to return the hidden Notificationtests

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectNotificationsectionAPI.NotificationsectionGetNotificationtestsV1(context.Background(), pkiNotificationsectionID).BShowHidden(bShowHidden).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotificationsectionAPI.NotificationsectionGetNotificationtestsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationsectionGetNotificationtestsV1`: NotificationsectionGetNotificationtestsV1Response
    fmt.Fprintf(os.Stdout, "Response from `ObjectNotificationsectionAPI.NotificationsectionGetNotificationtestsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotificationsectionID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationsectionGetNotificationtestsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bShowHidden** | **bool** | Whether or not to return the hidden Notificationtests | 

### Return type

[**NotificationsectionGetNotificationtestsV1Response**](NotificationsectionGetNotificationtestsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

