# eZmaxAPI\ObjectOfficetaxreportAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficetaxreportBatchDownloadV1**](ObjectOfficetaxreportAPI.md#OfficetaxreportBatchDownloadV1) | **Post** /1/object/officetaxreport/{pkiOfficetaxreportID}/batchDownload | Download multiples attachments from an Officetaxreport
[**OfficetaxreportGetAttachmentsV1**](ObjectOfficetaxreportAPI.md#OfficetaxreportGetAttachmentsV1) | **Get** /1/object/officetaxreport/{pkiOfficetaxreportID}/getAttachments | Retrieve Officetaxreport&#39;s attachments
[**OfficetaxreportImportIntoEDMV1**](ObjectOfficetaxreportAPI.md#OfficetaxreportImportIntoEDMV1) | **Post** /1/object/officetaxreport/{pkiOfficetaxreportID}/importIntoEDM | Import attachments into the Officetaxreport



## OfficetaxreportBatchDownloadV1

> *os.File OfficetaxreportBatchDownloadV1(ctx, pkiOfficetaxreportID).OfficetaxreportBatchDownloadV1Request(officetaxreportBatchDownloadV1Request).Execute()

Download multiples attachments from an Officetaxreport

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
	pkiOfficetaxreportID := int32(56) // int32 | 
	officetaxreportBatchDownloadV1Request := *openapiclient.NewOfficetaxreportBatchDownloadV1Request([]int32{int32(1)}) // OfficetaxreportBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectOfficetaxreportAPI.OfficetaxreportBatchDownloadV1(context.Background(), pkiOfficetaxreportID).OfficetaxreportBatchDownloadV1Request(officetaxreportBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectOfficetaxreportAPI.OfficetaxreportBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficetaxreportBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectOfficetaxreportAPI.OfficetaxreportBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiOfficetaxreportID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficetaxreportBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **officetaxreportBatchDownloadV1Request** | [**OfficetaxreportBatchDownloadV1Request**](OfficetaxreportBatchDownloadV1Request.md) |  | 

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


## OfficetaxreportGetAttachmentsV1

> OfficetaxreportGetAttachmentsV1Response OfficetaxreportGetAttachmentsV1(ctx, pkiOfficetaxreportID).Execute()

Retrieve Officetaxreport's attachments

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
	pkiOfficetaxreportID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectOfficetaxreportAPI.OfficetaxreportGetAttachmentsV1(context.Background(), pkiOfficetaxreportID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectOfficetaxreportAPI.OfficetaxreportGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficetaxreportGetAttachmentsV1`: OfficetaxreportGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectOfficetaxreportAPI.OfficetaxreportGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiOfficetaxreportID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficetaxreportGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OfficetaxreportGetAttachmentsV1Response**](OfficetaxreportGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficetaxreportImportIntoEDMV1

> OfficetaxreportImportIntoEDMV1Response OfficetaxreportImportIntoEDMV1(ctx, pkiOfficetaxreportID).OfficetaxreportImportIntoEDMV1Request(officetaxreportImportIntoEDMV1Request).Execute()

Import attachments into the Officetaxreport

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
	pkiOfficetaxreportID := int32(56) // int32 | 
	officetaxreportImportIntoEDMV1Request := *openapiclient.NewOfficetaxreportImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // OfficetaxreportImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectOfficetaxreportAPI.OfficetaxreportImportIntoEDMV1(context.Background(), pkiOfficetaxreportID).OfficetaxreportImportIntoEDMV1Request(officetaxreportImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectOfficetaxreportAPI.OfficetaxreportImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficetaxreportImportIntoEDMV1`: OfficetaxreportImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectOfficetaxreportAPI.OfficetaxreportImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiOfficetaxreportID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficetaxreportImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **officetaxreportImportIntoEDMV1Request** | [**OfficetaxreportImportIntoEDMV1Request**](OfficetaxreportImportIntoEDMV1Request.md) |  | 

### Return type

[**OfficetaxreportImportIntoEDMV1Response**](OfficetaxreportImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

