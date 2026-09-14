# eZmaxAPI\ObjectAttachmentAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachmentDeleteV1**](ObjectAttachmentAPI.md#AttachmentDeleteV1) | **Post** /1/object/attachment/{pkiAttachmentID}/delete | Delete an existing attachment
[**AttachmentDownloadV1**](ObjectAttachmentAPI.md#AttachmentDownloadV1) | **Get** /1/object/attachment/{pkiAttachmentID}/download | Retrieve the content
[**AttachmentGetAttachmentlogsV1**](ObjectAttachmentAPI.md#AttachmentGetAttachmentlogsV1) | **Get** /1/object/attachment/{pkiAttachmentID}/getAttachmentlogs | Retrieve the Attachmentlogs
[**AttachmentRenameV1**](ObjectAttachmentAPI.md#AttachmentRenameV1) | **Post** /1/object/attachment/{pkiAttachmentID}/rename | Rename an attachment
[**AttachmentRestoreV1**](ObjectAttachmentAPI.md#AttachmentRestoreV1) | **Post** /1/object/attachment/{pkiAttachmentID}/restore | Restore a deleted attachment
[**AttachmentValidateV1**](ObjectAttachmentAPI.md#AttachmentValidateV1) | **Patch** /1/object/attachment/{pkiAttachmentID}/validate | Validate an existing attachment



## AttachmentDeleteV1

> AttachmentDeleteV1Response AttachmentDeleteV1(ctx, pkiAttachmentID).Body(body).Execute()

Delete an existing attachment



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAttachmentAPI.AttachmentDeleteV1(context.Background(), pkiAttachmentID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAttachmentAPI.AttachmentDeleteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentDeleteV1`: AttachmentDeleteV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAttachmentAPI.AttachmentDeleteV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAttachmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentDeleteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**AttachmentDeleteV1Response**](AttachmentDeleteV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## AttachmentRenameV1

> AttachmentRenameV1Response AttachmentRenameV1(ctx, pkiAttachmentID).AttachmentRenameV1Request(attachmentRenameV1Request).Execute()

Rename an attachment



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
	attachmentRenameV1Request := *openapiclient.NewAttachmentRenameV1Request("Document.pdf", "Inscription") // AttachmentRenameV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAttachmentAPI.AttachmentRenameV1(context.Background(), pkiAttachmentID).AttachmentRenameV1Request(attachmentRenameV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAttachmentAPI.AttachmentRenameV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentRenameV1`: AttachmentRenameV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAttachmentAPI.AttachmentRenameV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAttachmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentRenameV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachmentRenameV1Request** | [**AttachmentRenameV1Request**](AttachmentRenameV1Request.md) |  | 

### Return type

[**AttachmentRenameV1Response**](AttachmentRenameV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentRestoreV1

> AttachmentRestoreV1Response AttachmentRestoreV1(ctx, pkiAttachmentID).AttachmentRestoreV1Request(attachmentRestoreV1Request).Execute()

Restore a deleted attachment



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
	attachmentRestoreV1Request := *openapiclient.NewAttachmentRestoreV1Request() // AttachmentRestoreV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAttachmentAPI.AttachmentRestoreV1(context.Background(), pkiAttachmentID).AttachmentRestoreV1Request(attachmentRestoreV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAttachmentAPI.AttachmentRestoreV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentRestoreV1`: AttachmentRestoreV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAttachmentAPI.AttachmentRestoreV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAttachmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentRestoreV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachmentRestoreV1Request** | [**AttachmentRestoreV1Request**](AttachmentRestoreV1Request.md) |  | 

### Return type

[**AttachmentRestoreV1Response**](AttachmentRestoreV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentValidateV1

> AttachmentValidateV1Response AttachmentValidateV1(ctx, pkiAttachmentID).AttachmentValidateV1Request(attachmentValidateV1Request).Execute()

Validate an existing attachment



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
	attachmentValidateV1Request := *openapiclient.NewAttachmentValidateV1Request(openapiclient.Field-eAttachmentVerified("No")) // AttachmentValidateV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectAttachmentAPI.AttachmentValidateV1(context.Background(), pkiAttachmentID).AttachmentValidateV1Request(attachmentValidateV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectAttachmentAPI.AttachmentValidateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentValidateV1`: AttachmentValidateV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectAttachmentAPI.AttachmentValidateV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiAttachmentID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentValidateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attachmentValidateV1Request** | [**AttachmentValidateV1Request**](AttachmentValidateV1Request.md) |  | 

### Return type

[**AttachmentValidateV1Response**](AttachmentValidateV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

