# eZmaxAPI\ObjectSalaryAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SalaryBatchDownloadV1**](ObjectSalaryAPI.md#SalaryBatchDownloadV1) | **Post** /1/object/salary/{pkiSalaryID}/batchDownload | Download multiples attachments from a Reconciliation
[**SalaryGetAttachmentsV1**](ObjectSalaryAPI.md#SalaryGetAttachmentsV1) | **Get** /1/object/salary/{pkiSalaryID}/getAttachments | Retrieve Salary&#39;s attachments
[**SalaryImportIntoEDMV1**](ObjectSalaryAPI.md#SalaryImportIntoEDMV1) | **Post** /1/object/salary/{pkiSalaryID}/importIntoEDM | Import attachments into the Salary



## SalaryBatchDownloadV1

> *os.File SalaryBatchDownloadV1(ctx, pkiSalaryID).SalaryBatchDownloadV1Request(salaryBatchDownloadV1Request).Execute()

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
	pkiSalaryID := int32(56) // int32 | 
	salaryBatchDownloadV1Request := *openapiclient.NewSalaryBatchDownloadV1Request([]int32{int32(1)}) // SalaryBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryBatchDownloadV1(context.Background(), pkiSalaryID).SalaryBatchDownloadV1Request(salaryBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSalaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSalaryBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **salaryBatchDownloadV1Request** | [**SalaryBatchDownloadV1Request**](SalaryBatchDownloadV1Request.md) |  | 

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


## SalaryGetAttachmentsV1

> SalaryGetAttachmentsV1Response SalaryGetAttachmentsV1(ctx, pkiSalaryID).Execute()

Retrieve Salary's attachments

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
	pkiSalaryID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryGetAttachmentsV1(context.Background(), pkiSalaryID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryGetAttachmentsV1`: SalaryGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSalaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSalaryGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalaryGetAttachmentsV1Response**](SalaryGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SalaryImportIntoEDMV1

> SalaryImportIntoEDMV1Response SalaryImportIntoEDMV1(ctx, pkiSalaryID).SalaryImportIntoEDMV1Request(salaryImportIntoEDMV1Request).Execute()

Import attachments into the Salary

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
	pkiSalaryID := int32(56) // int32 | 
	salaryImportIntoEDMV1Request := *openapiclient.NewSalaryImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // SalaryImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSalaryAPI.SalaryImportIntoEDMV1(context.Background(), pkiSalaryID).SalaryImportIntoEDMV1Request(salaryImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSalaryAPI.SalaryImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SalaryImportIntoEDMV1`: SalaryImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSalaryAPI.SalaryImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSalaryID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSalaryImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **salaryImportIntoEDMV1Request** | [**SalaryImportIntoEDMV1Request**](SalaryImportIntoEDMV1Request.md) |  | 

### Return type

[**SalaryImportIntoEDMV1Response**](SalaryImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

