# eZmaxAPI\ObjectDisclosureAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisclosureBatchDownloadV1**](ObjectDisclosureAPI.md#DisclosureBatchDownloadV1) | **Post** /1/object/disclosure/{pkiDisclosureID}/batchDownload | Download multiples attachments from a Disclosure
[**DisclosureGetAttachmentsV1**](ObjectDisclosureAPI.md#DisclosureGetAttachmentsV1) | **Get** /1/object/disclosure/{pkiDisclosureID}/getAttachments | Retrieve Disclosure&#39;s attachments
[**DisclosureImportIntoEDMV1**](ObjectDisclosureAPI.md#DisclosureImportIntoEDMV1) | **Post** /1/object/disclosure/{pkiDisclosureID}/importIntoEDM | Import attachments into the Disclosure



## DisclosureBatchDownloadV1

> *os.File DisclosureBatchDownloadV1(ctx, pkiDisclosureID).DisclosureBatchDownloadV1Request(disclosureBatchDownloadV1Request).Execute()

Download multiples attachments from a Disclosure

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
	pkiDisclosureID := int32(56) // int32 | 
	disclosureBatchDownloadV1Request := *openapiclient.NewDisclosureBatchDownloadV1Request([]int32{int32(1)}) // DisclosureBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDisclosureAPI.DisclosureBatchDownloadV1(context.Background(), pkiDisclosureID).DisclosureBatchDownloadV1Request(disclosureBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDisclosureAPI.DisclosureBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisclosureBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectDisclosureAPI.DisclosureBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDisclosureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisclosureBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disclosureBatchDownloadV1Request** | [**DisclosureBatchDownloadV1Request**](DisclosureBatchDownloadV1Request.md) |  | 

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


## DisclosureGetAttachmentsV1

> DisclosureGetAttachmentsV1Response DisclosureGetAttachmentsV1(ctx, pkiDisclosureID).Execute()

Retrieve Disclosure's attachments

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
	pkiDisclosureID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDisclosureAPI.DisclosureGetAttachmentsV1(context.Background(), pkiDisclosureID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDisclosureAPI.DisclosureGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisclosureGetAttachmentsV1`: DisclosureGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDisclosureAPI.DisclosureGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDisclosureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisclosureGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisclosureGetAttachmentsV1Response**](DisclosureGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisclosureImportIntoEDMV1

> DisclosureImportIntoEDMV1Response DisclosureImportIntoEDMV1(ctx, pkiDisclosureID).DisclosureImportIntoEDMV1Request(disclosureImportIntoEDMV1Request).Execute()

Import attachments into the Disclosure

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
	pkiDisclosureID := int32(56) // int32 | 
	disclosureImportIntoEDMV1Request := *openapiclient.NewDisclosureImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // DisclosureImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectDisclosureAPI.DisclosureImportIntoEDMV1(context.Background(), pkiDisclosureID).DisclosureImportIntoEDMV1Request(disclosureImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDisclosureAPI.DisclosureImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisclosureImportIntoEDMV1`: DisclosureImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectDisclosureAPI.DisclosureImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiDisclosureID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisclosureImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disclosureImportIntoEDMV1Request** | [**DisclosureImportIntoEDMV1Request**](DisclosureImportIntoEDMV1Request.md) |  | 

### Return type

[**DisclosureImportIntoEDMV1Response**](DisclosureImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

