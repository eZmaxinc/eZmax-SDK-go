# eZmaxAPI\ObjectPaymentpreparationAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentpreparationBatchDownloadV1**](ObjectPaymentpreparationAPI.md#PaymentpreparationBatchDownloadV1) | **Post** /1/object/paymentpreparation/{pkiPaymentpreparationID}/batchDownload | Download multiples attachments from an Paymentpreparation
[**PaymentpreparationGetAttachmentsV1**](ObjectPaymentpreparationAPI.md#PaymentpreparationGetAttachmentsV1) | **Get** /1/object/paymentpreparation/{pkiPaymentpreparationID}/getAttachments | Retrieve Paymentpreparation&#39;s attachments
[**PaymentpreparationImportIntoEDMV1**](ObjectPaymentpreparationAPI.md#PaymentpreparationImportIntoEDMV1) | **Post** /1/object/paymentpreparation/{pkiPaymentpreparationID}/importIntoEDM | Import attachments into the Paymentpreparation



## PaymentpreparationBatchDownloadV1

> *os.File PaymentpreparationBatchDownloadV1(ctx, pkiPaymentpreparationID).PaymentpreparationBatchDownloadV1Request(paymentpreparationBatchDownloadV1Request).Execute()

Download multiples attachments from an Paymentpreparation

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
	pkiPaymentpreparationID := int32(56) // int32 | 
	paymentpreparationBatchDownloadV1Request := *openapiclient.NewPaymentpreparationBatchDownloadV1Request([]int32{int32(1)}) // PaymentpreparationBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymentpreparationAPI.PaymentpreparationBatchDownloadV1(context.Background(), pkiPaymentpreparationID).PaymentpreparationBatchDownloadV1Request(paymentpreparationBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymentpreparationAPI.PaymentpreparationBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentpreparationBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymentpreparationAPI.PaymentpreparationBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPaymentpreparationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentpreparationBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentpreparationBatchDownloadV1Request** | [**PaymentpreparationBatchDownloadV1Request**](PaymentpreparationBatchDownloadV1Request.md) |  | 

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


## PaymentpreparationGetAttachmentsV1

> PaymentpreparationGetAttachmentsV1Response PaymentpreparationGetAttachmentsV1(ctx, pkiPaymentpreparationID).Execute()

Retrieve Paymentpreparation's attachments

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
	pkiPaymentpreparationID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymentpreparationAPI.PaymentpreparationGetAttachmentsV1(context.Background(), pkiPaymentpreparationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymentpreparationAPI.PaymentpreparationGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentpreparationGetAttachmentsV1`: PaymentpreparationGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymentpreparationAPI.PaymentpreparationGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPaymentpreparationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentpreparationGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentpreparationGetAttachmentsV1Response**](PaymentpreparationGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentpreparationImportIntoEDMV1

> PaymentpreparationImportIntoEDMV1Response PaymentpreparationImportIntoEDMV1(ctx, pkiPaymentpreparationID).PaymentpreparationImportIntoEDMV1Request(paymentpreparationImportIntoEDMV1Request).Execute()

Import attachments into the Paymentpreparation

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
	pkiPaymentpreparationID := int32(56) // int32 | 
	paymentpreparationImportIntoEDMV1Request := *openapiclient.NewPaymentpreparationImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // PaymentpreparationImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectPaymentpreparationAPI.PaymentpreparationImportIntoEDMV1(context.Background(), pkiPaymentpreparationID).PaymentpreparationImportIntoEDMV1Request(paymentpreparationImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectPaymentpreparationAPI.PaymentpreparationImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentpreparationImportIntoEDMV1`: PaymentpreparationImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectPaymentpreparationAPI.PaymentpreparationImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiPaymentpreparationID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentpreparationImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentpreparationImportIntoEDMV1Request** | [**PaymentpreparationImportIntoEDMV1Request**](PaymentpreparationImportIntoEDMV1Request.md) |  | 

### Return type

[**PaymentpreparationImportIntoEDMV1Response**](PaymentpreparationImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

