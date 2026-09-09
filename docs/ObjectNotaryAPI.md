# eZmaxAPI\ObjectNotaryAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotaryBatchDownloadV1**](ObjectNotaryAPI.md#NotaryBatchDownloadV1) | **Post** /1/object/notary/{pkiNotaryID}/batchDownload | Download multiples attachments from a Notary
[**NotaryGetAttachmentsV1**](ObjectNotaryAPI.md#NotaryGetAttachmentsV1) | **Get** /1/object/notary/{pkiNotaryID}/getAttachments | Retrieve Notary&#39;s attachments
[**NotaryImportIntoEDMV1**](ObjectNotaryAPI.md#NotaryImportIntoEDMV1) | **Post** /1/object/notary/{pkiNotaryID}/importIntoEDM | Import attachments into the Notary



## NotaryBatchDownloadV1

> *os.File NotaryBatchDownloadV1(ctx, pkiNotaryID).NotaryBatchDownloadV1Request(notaryBatchDownloadV1Request).Execute()

Download multiples attachments from a Notary

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
	pkiNotaryID := int32(56) // int32 | 
	notaryBatchDownloadV1Request := *openapiclient.NewNotaryBatchDownloadV1Request([]int32{int32(1)}) // NotaryBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryBatchDownloadV1(context.Background(), pkiNotaryID).NotaryBatchDownloadV1Request(notaryBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotaryBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notaryBatchDownloadV1Request** | [**NotaryBatchDownloadV1Request**](NotaryBatchDownloadV1Request.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip, text/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotaryGetAttachmentsV1

> NotaryGetAttachmentsV1Response NotaryGetAttachmentsV1(ctx, pkiNotaryID).Execute()

Retrieve Notary's attachments

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
	pkiNotaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryGetAttachmentsV1(context.Background(), pkiNotaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryGetAttachmentsV1`: NotaryGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotaryGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotaryGetAttachmentsV1Response**](NotaryGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotaryImportIntoEDMV1

> NotaryImportIntoEDMV1Response NotaryImportIntoEDMV1(ctx, pkiNotaryID).NotaryImportIntoEDMV1Request(notaryImportIntoEDMV1Request).Execute()

Import attachments into the Notary

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
	pkiNotaryID := int32(56) // int32 | 
	notaryImportIntoEDMV1Request := *openapiclient.NewNotaryImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // NotaryImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectNotaryAPI.NotaryImportIntoEDMV1(context.Background(), pkiNotaryID).NotaryImportIntoEDMV1Request(notaryImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectNotaryAPI.NotaryImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotaryImportIntoEDMV1`: NotaryImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectNotaryAPI.NotaryImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiNotaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotaryImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notaryImportIntoEDMV1Request** | [**NotaryImportIntoEDMV1Request**](NotaryImportIntoEDMV1Request.md) |  | 

### Return type

[**NotaryImportIntoEDMV1Response**](NotaryImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

