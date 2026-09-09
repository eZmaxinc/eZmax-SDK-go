# eZmaxAPI\ObjectDepositAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DepositBatchDownloadV1**](ObjectDepositAPI.md#DepositBatchDownloadV1) | **Post** /1/object/deposit/{pkiDepositID}/batchDownload | Download multiples attachments from a Deposit
[**DepositGetAttachmentsV1**](ObjectDepositAPI.md#DepositGetAttachmentsV1) | **Get** /1/object/deposit/{pkiDepositID}/getAttachments | Retrieve Deposit&#39;s attachments
[**DepositImportIntoEDMV1**](ObjectDepositAPI.md#DepositImportIntoEDMV1) | **Post** /1/object/deposit/{pkiDepositID}/importIntoEDM | Import attachments into the Deposit



## DepositBatchDownloadV1

> *os.File DepositBatchDownloadV1(ctx, pkiDepositID).DepositBatchDownloadV1Request(depositBatchDownloadV1Request).Execute()

Download multiples attachments from a Deposit

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
	pkiDepositID := int32(56) // int32 | 
	depositBatchDownloadV1Request := *openapiclient.NewDepositBatchDownloadV1Request([]int32{int32(1)}) // DepositBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDepositAPI.DepositBatchDownloadV1(context.Background(), pkiDepositID).DepositBatchDownloadV1Request(depositBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDepositAPI.DepositBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DepositBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectDepositAPI.DepositBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDepositID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDepositBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depositBatchDownloadV1Request** | [**DepositBatchDownloadV1Request**](DepositBatchDownloadV1Request.md) |  | 

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


## DepositGetAttachmentsV1

> DepositGetAttachmentsV1Response DepositGetAttachmentsV1(ctx, pkiDepositID).Execute()

Retrieve Deposit's attachments

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
	pkiDepositID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDepositAPI.DepositGetAttachmentsV1(context.Background(), pkiDepositID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDepositAPI.DepositGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DepositGetAttachmentsV1`: DepositGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDepositAPI.DepositGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDepositID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDepositGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DepositGetAttachmentsV1Response**](DepositGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositImportIntoEDMV1

> DepositImportIntoEDMV1Response DepositImportIntoEDMV1(ctx, pkiDepositID).DepositImportIntoEDMV1Request(depositImportIntoEDMV1Request).Execute()

Import attachments into the Deposit

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
	pkiDepositID := int32(56) // int32 | 
	depositImportIntoEDMV1Request := *openapiclient.NewDepositImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // DepositImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDepositAPI.DepositImportIntoEDMV1(context.Background(), pkiDepositID).DepositImportIntoEDMV1Request(depositImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDepositAPI.DepositImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DepositImportIntoEDMV1`: DepositImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDepositAPI.DepositImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDepositID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDepositImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depositImportIntoEDMV1Request** | [**DepositImportIntoEDMV1Request**](DepositImportIntoEDMV1Request.md) |  | 

### Return type

[**DepositImportIntoEDMV1Response**](DepositImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

