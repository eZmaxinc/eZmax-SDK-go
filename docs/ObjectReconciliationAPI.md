# eZmaxAPI\ObjectReconciliationAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReconciliationBatchDownloadV1**](ObjectReconciliationAPI.md#ReconciliationBatchDownloadV1) | **Post** /1/object/reconciliation/{pkiReconciliationID}/batchDownload | Download multiples attachments from a Reconciliation
[**ReconciliationGetAttachmentsV1**](ObjectReconciliationAPI.md#ReconciliationGetAttachmentsV1) | **Get** /1/object/reconciliation/{pkiReconciliationID}/getAttachments | Retrieve Reconciliation&#39;s attachments
[**ReconciliationImportIntoEDMV1**](ObjectReconciliationAPI.md#ReconciliationImportIntoEDMV1) | **Post** /1/object/reconciliation/{pkiReconciliationID}/importIntoEDM | Import attachments into the Reconciliation



## ReconciliationBatchDownloadV1

> *os.File ReconciliationBatchDownloadV1(ctx, pkiReconciliationID).ReconciliationBatchDownloadV1Request(reconciliationBatchDownloadV1Request).Execute()

Download multiples attachments from a Reconciliation

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
	pkiReconciliationID := int32(56) // int32 | 
	reconciliationBatchDownloadV1Request := *openapiclient.NewReconciliationBatchDownloadV1Request([]int32{int32(1)}) // ReconciliationBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectReconciliationAPI.ReconciliationBatchDownloadV1(context.Background(), pkiReconciliationID).ReconciliationBatchDownloadV1Request(reconciliationBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectReconciliationAPI.ReconciliationBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReconciliationBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectReconciliationAPI.ReconciliationBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiReconciliationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReconciliationBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reconciliationBatchDownloadV1Request** | [**ReconciliationBatchDownloadV1Request**](ReconciliationBatchDownloadV1Request.md) |  | 

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


## ReconciliationGetAttachmentsV1

> ReconciliationGetAttachmentsV1Response ReconciliationGetAttachmentsV1(ctx, pkiReconciliationID).Execute()

Retrieve Reconciliation's attachments

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
	pkiReconciliationID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectReconciliationAPI.ReconciliationGetAttachmentsV1(context.Background(), pkiReconciliationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectReconciliationAPI.ReconciliationGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReconciliationGetAttachmentsV1`: ReconciliationGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectReconciliationAPI.ReconciliationGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiReconciliationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReconciliationGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReconciliationGetAttachmentsV1Response**](ReconciliationGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReconciliationImportIntoEDMV1

> ReconciliationImportIntoEDMV1Response ReconciliationImportIntoEDMV1(ctx, pkiReconciliationID).ReconciliationImportIntoEDMV1Request(reconciliationImportIntoEDMV1Request).Execute()

Import attachments into the Reconciliation

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
	pkiReconciliationID := int32(56) // int32 | 
	reconciliationImportIntoEDMV1Request := *openapiclient.NewReconciliationImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // ReconciliationImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectReconciliationAPI.ReconciliationImportIntoEDMV1(context.Background(), pkiReconciliationID).ReconciliationImportIntoEDMV1Request(reconciliationImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectReconciliationAPI.ReconciliationImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReconciliationImportIntoEDMV1`: ReconciliationImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectReconciliationAPI.ReconciliationImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiReconciliationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReconciliationImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reconciliationImportIntoEDMV1Request** | [**ReconciliationImportIntoEDMV1Request**](ReconciliationImportIntoEDMV1Request.md) |  | 

### Return type

[**ReconciliationImportIntoEDMV1Response**](ReconciliationImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

