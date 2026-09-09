# eZmaxAPI\ObjectExternalbrokerAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalbrokerBatchDownloadV1**](ObjectExternalbrokerAPI.md#ExternalbrokerBatchDownloadV1) | **Post** /1/object/externalbroker/{pkiExternalbrokerID}/batchDownload | Download multiples attachments from an Externalbroker
[**ExternalbrokerGetAttachmentsV1**](ObjectExternalbrokerAPI.md#ExternalbrokerGetAttachmentsV1) | **Get** /1/object/externalbroker/{pkiExternalbrokerID}/getAttachments | Retrieve Externalbroker&#39;s attachments
[**ExternalbrokerImportIntoEDMV1**](ObjectExternalbrokerAPI.md#ExternalbrokerImportIntoEDMV1) | **Post** /1/object/externalbroker/{pkiExternalbrokerID}/importIntoEDM | Import attachments into the Externalbroker



## ExternalbrokerBatchDownloadV1

> *os.File ExternalbrokerBatchDownloadV1(ctx, pkiExternalbrokerID).ExternalbrokerBatchDownloadV1Request(externalbrokerBatchDownloadV1Request).Execute()

Download multiples attachments from an Externalbroker

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
	pkiExternalbrokerID := int32(56) // int32 | 
	externalbrokerBatchDownloadV1Request := *openapiclient.NewExternalbrokerBatchDownloadV1Request([]int32{int32(1)}) // ExternalbrokerBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerBatchDownloadV1(context.Background(), pkiExternalbrokerID).ExternalbrokerBatchDownloadV1Request(externalbrokerBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalbrokerBatchDownloadV1Request** | [**ExternalbrokerBatchDownloadV1Request**](ExternalbrokerBatchDownloadV1Request.md) |  | 

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


## ExternalbrokerGetAttachmentsV1

> ExternalbrokerGetAttachmentsV1Response ExternalbrokerGetAttachmentsV1(ctx, pkiExternalbrokerID).Execute()

Retrieve Externalbroker's attachments

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
	pkiExternalbrokerID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerGetAttachmentsV1(context.Background(), pkiExternalbrokerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerGetAttachmentsV1`: ExternalbrokerGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalbrokerGetAttachmentsV1Response**](ExternalbrokerGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalbrokerImportIntoEDMV1

> ExternalbrokerImportIntoEDMV1Response ExternalbrokerImportIntoEDMV1(ctx, pkiExternalbrokerID).ExternalbrokerImportIntoEDMV1Request(externalbrokerImportIntoEDMV1Request).Execute()

Import attachments into the Externalbroker



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
	pkiExternalbrokerID := int32(56) // int32 | 
	externalbrokerImportIntoEDMV1Request := *openapiclient.NewExternalbrokerImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // ExternalbrokerImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectExternalbrokerAPI.ExternalbrokerImportIntoEDMV1(context.Background(), pkiExternalbrokerID).ExternalbrokerImportIntoEDMV1Request(externalbrokerImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectExternalbrokerAPI.ExternalbrokerImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalbrokerImportIntoEDMV1`: ExternalbrokerImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectExternalbrokerAPI.ExternalbrokerImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiExternalbrokerID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalbrokerImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalbrokerImportIntoEDMV1Request** | [**ExternalbrokerImportIntoEDMV1Request**](ExternalbrokerImportIntoEDMV1Request.md) |  | 

### Return type

[**ExternalbrokerImportIntoEDMV1Response**](ExternalbrokerImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

