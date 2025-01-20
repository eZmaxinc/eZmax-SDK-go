# eZmaxAPI\ObjectDiscussionmessageAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscussionmessageCreateObjectV1**](ObjectDiscussionmessageAPI.md#DiscussionmessageCreateObjectV1) | **Post** /1/object/discussionmessage | Create a new Discussionmessage
[**DiscussionmessageDeleteObjectV1**](ObjectDiscussionmessageAPI.md#DiscussionmessageDeleteObjectV1) | **Delete** /1/object/discussionmessage/{pkiDiscussionmessageID} | Delete an existing Discussionmessage
[**DiscussionmessagePatchObjectV1**](ObjectDiscussionmessageAPI.md#DiscussionmessagePatchObjectV1) | **Patch** /1/object/discussionmessage/{pkiDiscussionmessageID} | Patch an existing Discussionmessage



## DiscussionmessageCreateObjectV1

> DiscussionmessageCreateObjectV1Response DiscussionmessageCreateObjectV1(ctx).DiscussionmessageCreateObjectV1Request(discussionmessageCreateObjectV1Request).Execute()

Create a new Discussionmessage



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
	discussionmessageCreateObjectV1Request := *openapiclient.NewDiscussionmessageCreateObjectV1Request([]openapiclient.DiscussionmessageRequestCompound{*openapiclient.NewDiscussionmessageRequestCompound(int32(125), "Hello, this is an example of content in a message")}) // DiscussionmessageCreateObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDiscussionmessageAPI.DiscussionmessageCreateObjectV1(context.Background()).DiscussionmessageCreateObjectV1Request(discussionmessageCreateObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDiscussionmessageAPI.DiscussionmessageCreateObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscussionmessageCreateObjectV1`: DiscussionmessageCreateObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDiscussionmessageAPI.DiscussionmessageCreateObjectV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscussionmessageCreateObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discussionmessageCreateObjectV1Request** | [**DiscussionmessageCreateObjectV1Request**](DiscussionmessageCreateObjectV1Request.md) |  | 

### Return type

[**DiscussionmessageCreateObjectV1Response**](DiscussionmessageCreateObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscussionmessageDeleteObjectV1

> DiscussionmessageDeleteObjectV1Response DiscussionmessageDeleteObjectV1(ctx, pkiDiscussionmessageID).Execute()

Delete an existing Discussionmessage



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
	pkiDiscussionmessageID := int32(56) // int32 | The unique ID of the Discussionmessage

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDiscussionmessageAPI.DiscussionmessageDeleteObjectV1(context.Background(), pkiDiscussionmessageID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDiscussionmessageAPI.DiscussionmessageDeleteObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscussionmessageDeleteObjectV1`: DiscussionmessageDeleteObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDiscussionmessageAPI.DiscussionmessageDeleteObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDiscussionmessageID** | **int32** | The unique ID of the Discussionmessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscussionmessageDeleteObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiscussionmessageDeleteObjectV1Response**](DiscussionmessageDeleteObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscussionmessagePatchObjectV1

> DiscussionmessagePatchObjectV1Response DiscussionmessagePatchObjectV1(ctx, pkiDiscussionmessageID).DiscussionmessagePatchObjectV1Request(discussionmessagePatchObjectV1Request).Execute()

Patch an existing Discussionmessage



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
	pkiDiscussionmessageID := int32(56) // int32 | The unique ID of the Discussionmessage
	discussionmessagePatchObjectV1Request := *openapiclient.NewDiscussionmessagePatchObjectV1Request(*openapiclient.NewDiscussionmessageRequestPatch()) // DiscussionmessagePatchObjectV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDiscussionmessageAPI.DiscussionmessagePatchObjectV1(context.Background(), pkiDiscussionmessageID).DiscussionmessagePatchObjectV1Request(discussionmessagePatchObjectV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDiscussionmessageAPI.DiscussionmessagePatchObjectV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscussionmessagePatchObjectV1`: DiscussionmessagePatchObjectV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDiscussionmessageAPI.DiscussionmessagePatchObjectV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDiscussionmessageID** | **int32** | The unique ID of the Discussionmessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscussionmessagePatchObjectV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **discussionmessagePatchObjectV1Request** | [**DiscussionmessagePatchObjectV1Request**](DiscussionmessagePatchObjectV1Request.md) |  | 

### Return type

[**DiscussionmessagePatchObjectV1Response**](DiscussionmessagePatchObjectV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

