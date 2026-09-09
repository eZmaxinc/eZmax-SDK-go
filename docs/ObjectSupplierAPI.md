# eZmaxAPI\ObjectSupplierAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SupplierBatchDownloadV1**](ObjectSupplierAPI.md#SupplierBatchDownloadV1) | **Post** /1/object/supplier/{pkiSupplierID}/batchDownload | Download multiples attachments from a Supplier
[**SupplierGetAttachmentsV1**](ObjectSupplierAPI.md#SupplierGetAttachmentsV1) | **Get** /1/object/supplier/{pkiSupplierID}/getAttachments | Retrieve Supplier&#39;s attachments
[**SupplierGetListV1**](ObjectSupplierAPI.md#SupplierGetListV1) | **Get** /1/object/supplier/getList | Retrieve Supplier list
[**SupplierImportIntoEDMV1**](ObjectSupplierAPI.md#SupplierImportIntoEDMV1) | **Post** /1/object/supplier/{pkiSupplierID}/importIntoEDM | Import attachments into the Supplier



## SupplierBatchDownloadV1

> *os.File SupplierBatchDownloadV1(ctx, pkiSupplierID).SupplierBatchDownloadV1Request(supplierBatchDownloadV1Request).Execute()

Download multiples attachments from a Supplier

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
	pkiSupplierID := int32(56) // int32 | 
	supplierBatchDownloadV1Request := *openapiclient.NewSupplierBatchDownloadV1Request([]int32{int32(1)}) // SupplierBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSupplierAPI.SupplierBatchDownloadV1(context.Background(), pkiSupplierID).SupplierBatchDownloadV1Request(supplierBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSupplierAPI.SupplierBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplierBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectSupplierAPI.SupplierBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSupplierID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupplierBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supplierBatchDownloadV1Request** | [**SupplierBatchDownloadV1Request**](SupplierBatchDownloadV1Request.md) |  | 

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


## SupplierGetAttachmentsV1

> SupplierGetAttachmentsV1Response SupplierGetAttachmentsV1(ctx, pkiSupplierID).Execute()

Retrieve Supplier's attachments

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
	pkiSupplierID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSupplierAPI.SupplierGetAttachmentsV1(context.Background(), pkiSupplierID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSupplierAPI.SupplierGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplierGetAttachmentsV1`: SupplierGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSupplierAPI.SupplierGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSupplierID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupplierGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupplierGetAttachmentsV1Response**](SupplierGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupplierGetListV1

> SupplierGetListV1Response SupplierGetListV1(ctx).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()

Retrieve Supplier list



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
	eOrderBy := "eOrderBy_example" // string | Specify how you want the results to be sorted (optional)
	iRowMax := int32(56) // int32 |  (optional)
	iRowOffset := int32(56) // int32 |  (optional) (default to 0)
	acceptLanguage := openapiclient.Header-Accept-Language("*") // HeaderAcceptLanguage |  (optional)
	sFilter := "sFilter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSupplierAPI.SupplierGetListV1(context.Background()).EOrderBy(eOrderBy).IRowMax(iRowMax).IRowOffset(iRowOffset).AcceptLanguage(acceptLanguage).SFilter(sFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSupplierAPI.SupplierGetListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplierGetListV1`: SupplierGetListV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSupplierAPI.SupplierGetListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSupplierGetListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eOrderBy** | **string** | Specify how you want the results to be sorted | 
 **iRowMax** | **int32** |  | 
 **iRowOffset** | **int32** |  | [default to 0]
 **acceptLanguage** | [**HeaderAcceptLanguage**](HeaderAcceptLanguage.md) |  | 
 **sFilter** | **string** |  | 

### Return type

[**SupplierGetListV1Response**](SupplierGetListV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupplierImportIntoEDMV1

> SupplierImportIntoEDMV1Response SupplierImportIntoEDMV1(ctx, pkiSupplierID).SupplierImportIntoEDMV1Request(supplierImportIntoEDMV1Request).Execute()

Import attachments into the Supplier



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
	pkiSupplierID := int32(56) // int32 | 
	supplierImportIntoEDMV1Request := *openapiclient.NewSupplierImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // SupplierImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectSupplierAPI.SupplierImportIntoEDMV1(context.Background(), pkiSupplierID).SupplierImportIntoEDMV1Request(supplierImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectSupplierAPI.SupplierImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SupplierImportIntoEDMV1`: SupplierImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectSupplierAPI.SupplierImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiSupplierID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSupplierImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supplierImportIntoEDMV1Request** | [**SupplierImportIntoEDMV1Request**](SupplierImportIntoEDMV1Request.md) |  | 

### Return type

[**SupplierImportIntoEDMV1Response**](SupplierImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

