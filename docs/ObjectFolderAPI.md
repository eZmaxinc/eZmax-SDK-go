# eZmaxAPI\ObjectFolderAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FolderBatchDownloadV1**](ObjectFolderAPI.md#FolderBatchDownloadV1) | **Post** /1/object/folder/{pkiFolderID}/batchDownload | Download multiples attachments from an Folder
[**FolderGetAttachmentsV1**](ObjectFolderAPI.md#FolderGetAttachmentsV1) | **Get** /1/object/folder/{pkiFolderID}/getAttachments | Retrieve Folder&#39;s attachments
[**FolderImportIntoEDMV1**](ObjectFolderAPI.md#FolderImportIntoEDMV1) | **Post** /1/object/folder/{pkiFolderID}/importIntoEDM | Import attachments into the Folder



## FolderBatchDownloadV1

> *os.File FolderBatchDownloadV1(ctx, pkiFolderID).FolderBatchDownloadV1Request(folderBatchDownloadV1Request).Execute()

Download multiples attachments from an Folder

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
	pkiFolderID := int32(56) // int32 | 
	folderBatchDownloadV1Request := *openapiclient.NewFolderBatchDownloadV1Request([]int32{int32(1)}) // FolderBatchDownloadV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFolderAPI.FolderBatchDownloadV1(context.Background(), pkiFolderID).FolderBatchDownloadV1Request(folderBatchDownloadV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFolderAPI.FolderBatchDownloadV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FolderBatchDownloadV1`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ObjectFolderAPI.FolderBatchDownloadV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiFolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderBatchDownloadV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderBatchDownloadV1Request** | [**FolderBatchDownloadV1Request**](FolderBatchDownloadV1Request.md) |  | 

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


## FolderGetAttachmentsV1

> FolderGetAttachmentsV1Response FolderGetAttachmentsV1(ctx, pkiFolderID).Execute()

Retrieve Folder's attachments

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
	pkiFolderID := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFolderAPI.FolderGetAttachmentsV1(context.Background(), pkiFolderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFolderAPI.FolderGetAttachmentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FolderGetAttachmentsV1`: FolderGetAttachmentsV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectFolderAPI.FolderGetAttachmentsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiFolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderGetAttachmentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderGetAttachmentsV1Response**](FolderGetAttachmentsV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FolderImportIntoEDMV1

> FolderImportIntoEDMV1Response FolderImportIntoEDMV1(ctx, pkiFolderID).FolderImportIntoEDMV1Request(folderImportIntoEDMV1Request).Execute()

Import attachments into the Folder

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
	pkiFolderID := int32(56) // int32 | 
	folderImportIntoEDMV1Request := *openapiclient.NewFolderImportIntoEDMV1Request([]openapiclient.CustomAttachmentImportIntoEDMRequest{*openapiclient.NewCustomAttachmentImportIntoEDMRequest("EAttachmentSource_example", "Document.pdf", "Inscription", openapiclient.Field-eAttachmentPrivacy("All"))}) // FolderImportIntoEDMV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectFolderAPI.FolderImportIntoEDMV1(context.Background(), pkiFolderID).FolderImportIntoEDMV1Request(folderImportIntoEDMV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectFolderAPI.FolderImportIntoEDMV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FolderImportIntoEDMV1`: FolderImportIntoEDMV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectFolderAPI.FolderImportIntoEDMV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pkiFolderID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFolderImportIntoEDMV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderImportIntoEDMV1Request** | [**FolderImportIntoEDMV1Request**](FolderImportIntoEDMV1Request.md) |  | 

### Return type

[**FolderImportIntoEDMV1Response**](FolderImportIntoEDMV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

