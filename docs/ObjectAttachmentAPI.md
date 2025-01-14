# eZmaxAPI\ObjectAttachmentAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachmentDownloadV1**](ObjectAttachmentAPI.md#AttachmentDownloadV1) | **Get** /1/object/attachment/{pkiAttachmentID}/download | Retrieve the content
[**AttachmentGetAttachmentlogsV1**](ObjectAttachmentAPI.md#AttachmentGetAttachmentlogsV1) | **Get** /1/object/attachment/{pkiAttachmentID}/getAttachmentlogs | Retrieve the Attachmentlogs



## AttachmentDownloadV1

> AttachmentDownloadV1(ctx, pkiAttachmentID).Execute()

Retrieve the content



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
	pkiAttachmentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectAttachmentAPI.AttachmentDownloadV1(context.Background(), pkiAttachmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAttachmentAPI.AttachmentDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAttachmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Authorization](../README.md#Authorization), [Presigned](../README.md#Presigned)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentGetAttachmentlogsV1

> AttachmentGetAttachmentlogsV1Response AttachmentGetAttachmentlogsV1(ctx, pkiAttachmentID).Execute()

Retrieve the Attachmentlogs



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
	pkiAttachmentID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAttachmentAPI.AttachmentGetAttachmentlogsV1(context.Background(), pkiAttachmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAttachmentAPI.AttachmentGetAttachmentlogsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentGetAttachmentlogsV1`: AttachmentGetAttachmentlogsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAttachmentAPI.AttachmentGetAttachmentlogsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAttachmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentGetAttachmentlogsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttachmentGetAttachmentlogsV1Response**](AttachmentGetAttachmentlogsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

