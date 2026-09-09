# eZmaxAPI\ObjectPurchaseAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PurchaseBatchDownloadV1**](ObjectPurchaseAPI.md#PurchaseBatchDownloadV1) | **Post** /1/object/purchase/{pkiPurchaseID}/batchDownload | Download multiples attachments from a Purchase
[**PurchaseGetAttachmentsV1**](ObjectPurchaseAPI.md#PurchaseGetAttachmentsV1) | **Get** /1/object/purchase/{pkiPurchaseID}/getAttachments | Retrieve Purchase&#39;s attachments
[**PurchaseImportIntoEDMV1**](ObjectPurchaseAPI.md#PurchaseImportIntoEDMV1) | **Post** /1/object/purchase/{pkiPurchaseID}/importIntoEDM | Import attachments into the Purchase



## PurchaseBatchDownloadV1

> *os.File PurchaseBatchDownloadV1(ctx, pkiPurchaseID).PurchaseBatchDownloadV1Request(purchaseBatchDownloadV1Request).Execute()

Download multiples attachments from a Purchase

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
	pkiPurchaseID := int32(56) // int32 | 
	purchaseBatchDownloadV1Request := *openapiclient.NewPurchaseBatchDownloadV1Request([]int32{int32(1)}) // PurchaseBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseBatchDownloadV1(context.Background(), pkiPurchaseID).PurchaseBatchDownloadV1Request(purchaseBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseBatchDownloadV1Request** | [**PurchaseBatchDownloadV1Request**](PurchaseBatchDownloadV1Request.md) |  | 

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


## PurchaseGetAttachmentsV1

> PurchaseGetAttachmentsV1Response PurchaseGetAttachmentsV1(ctx, pkiPurchaseID).Execute()

Retrieve Purchase's attachments

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
	pkiPurchaseID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseGetAttachmentsV1(context.Background(), pkiPurchaseID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseGetAttachmentsV1`: PurchaseGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurchaseGetAttachmentsV1Response**](PurchaseGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseImportIntoEDMV1

> PurchaseImportIntoEDMV1Response PurchaseImportIntoEDMV1(ctx, pkiPurchaseID).PurchaseImportIntoEDMV1Request(purchaseImportIntoEDMV1Request).Execute()

Import attachments into the Purchase

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
	pkiPurchaseID := int32(56) // int32 | 
	purchaseImportIntoEDMV1Request := *openapiclient.NewPurchaseImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // PurchaseImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPurchaseAPI.PurchaseImportIntoEDMV1(context.Background(), pkiPurchaseID).PurchaseImportIntoEDMV1Request(purchaseImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPurchaseAPI.PurchaseImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseImportIntoEDMV1`: PurchaseImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPurchaseAPI.PurchaseImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPurchaseID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseImportIntoEDMV1Request** | [**PurchaseImportIntoEDMV1Request**](PurchaseImportIntoEDMV1Request.md) |  | 

### Return type

[**PurchaseImportIntoEDMV1Response**](PurchaseImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

